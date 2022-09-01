package middleware

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"sls/internal/entity/errorEntity"
	"sls/internal/service/tokenService"
	"strings"
)

var (
	authorizationKeyHeader  = "Authorization"
	AuthorizationTypeBearer = "bearer"
)

type middleware struct {
	tokenSrv tokenService.TokenSrv
}

func NewMiddleware(tokenSrv tokenService.TokenSrv) *middleware {
	return &middleware{tokenSrv: tokenSrv}
}

func (m *middleware) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("content-type", "application/json")
		authorizationHeader := request.Header.Get(authorizationKeyHeader)
		if len(authorizationHeader) == 0 {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(errorEntity.NewErrorRes(http.StatusUnauthorized, "No Token"))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(errorEntity.NewErrorRes(http.StatusUnauthorized, "Invalid Header Format"))
			return
		}

		AuthorizationType := strings.ToLower(fields[0])
		if AuthorizationType != AuthorizationTypeBearer {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(errorEntity.NewErrorRes(http.StatusUnauthorized,
				"Invalid Authentication Format"))
			return
		}
		accessToken := fields[1]
		user, err := m.tokenSrv.VerifyToken(accessToken)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(errorEntity.NewErrorRes(http.StatusUnauthorized,
				"Token not verified: "+err.Error()))
			return
		}
		newId := user.Id.String()
		ctx := context.WithValue(request.Context(), "id", newId)
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
func (m *middleware) MapToId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		urlId := chi.URLParam(request, "id")
		ctxId := request.Context().Value("id")
		if urlId != ctxId {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(errorEntity.NewErrorRes(http.StatusUnauthorized,
				"You are not allowed to modify this resource"))
			return
		}
		next.ServeHTTP(writer, request)
	})
}

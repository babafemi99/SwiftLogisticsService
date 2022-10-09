FROM golang:1.19.1-alpine3.16 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o sls-main main.go


FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/sls-main /app
COPY --from=builder /app/internal ./internal
#COPY --from=builder /app .
EXPOSE 8080
CMD ["/app/sls-main"]

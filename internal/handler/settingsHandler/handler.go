package settingsHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sls/internal/entity/adminEntity"
	"sls/internal/service/settingsService"
	"strings"
)

type settingsHandler struct {
	srv settingsService.SettingsService
}

func (a *settingsHandler) GetSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	settings, err := a.srv.GetSettings()
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting settings: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(settings)
}

func NewSettingsHandler(srv settingsService.SettingsService) *settingsHandler {
	return &settingsHandler{srv: srv}
}

func (a *settingsHandler) EditSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var req adminEntity.AdminSettings
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "wrong input", http.StatusBadRequest)
	}

	id := strings.Split(r.URL.Path, "/")[3]

	settings, err := a.srv.EditSettings(id, &req)
	if err != nil {
		http.Error(w, fmt.Sprintf("error editing settings: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(settings)
}

func (a *settingsHandler) CreateSetting(w http.ResponseWriter, r *http.Request) {
	log.Println("inside here")
	w.Header().Set("content-type", "application/json")
	var req adminEntity.AdminSettings
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("inside here 3")
		http.Error(w, "wrong input", http.StatusInternalServerError)
	}
	log.Println("inside here2")

	settings, err := a.srv.CreateSettings(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating settings: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(settings)
}

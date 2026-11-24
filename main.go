package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"keka/incident"
	"keka/notification"
	"net/http"
	"os"
)

func main() {
	mockNotifService := notification.NewMockCallService()
	incidentServic := incident.NewService(mockNotifService)

	handler := NewHandler(incidentServic)

	println("starting server")
	mux := http.NewServeMux()
	mux.HandleFunc("POST /alert", handler.Alert)

	err := http.ListenAndServe("localhost:4614", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}

func NewHandler(svc *incident.Service) Handler {
	return Handler{svc: svc}
}

type Handler struct {
	svc *incident.Service
}

func (h *Handler) Alert(w http.ResponseWriter, r *http.Request) {
	var req incident.Alert
	if !decode(w, r, &req) {
		return
	}

	res, err := h.svc.Alert(r.Context(), &req)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, res)
}

func decode(w http.ResponseWriter, r *http.Request, dst interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON body"})
		return false
	}
	return true
}

func writeJSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	switch {
	//case errors.Is(err, ErrNotFound):
	//	status = http.StatusNotFound
	}
	writeJSON(w, status, map[string]string{"error": err.Error()})
}

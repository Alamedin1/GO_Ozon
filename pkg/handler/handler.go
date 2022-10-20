package handler

import (
	"net/http"
)

type Handler struct {
}

func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", handler)
	return mux
}

func (h *Handler) shortURLHandler(w http.ResponseWriter, r *http.Request) {

}

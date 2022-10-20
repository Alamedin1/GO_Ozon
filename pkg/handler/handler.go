package handler

import "net/http"

type Handler struct {
}

func (h *Handler) shortURLHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		h.notFound(w)
		return
	}

	if req.Method == http.MethodGet {

	}

}

func (h *Handler) serverError(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (h *Handler) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (h *Handler) notFound(w http.ResponseWriter) {
	h.clientError(w, http.StatusNotFound)
}

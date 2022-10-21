package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sample-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.shortURLHandler)
	return mux
}

func (h *Handler) shortURLHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		h.notFound(w)
		return
	}

	if req.Method == http.MethodGet {
		shortUrl := req.URL.Query().Get("url")
		if shortUrl == "" {
			h.clientError(w, http.StatusBadRequest)
			return
		}

		url, err := h.services.UrlGet(shortUrl)
		if err != nil {
			h.serverError(w)
			return
		}
		if url == "" {
			h.clientError(w, http.StatusBadRequest)
			return
		}

		w.Write([]byte(url))
		return
	}

	//Post, сохраняет оригинальный URL, возвращать сокращённый
	if req.Method == http.MethodPost {
		if req.Header.Get("Content-Type") != "text/plain; charset=utf-8" {
			h.clientError(w, http.StatusUnsupportedMediaType)
			return
		}

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			h.serverError(w)
			return
		}

		url := string(body)
		if url == "" {
			h.clientError(w, http.StatusBadRequest)
			return
		}

		shortUrl, err := h.services.UrlPost(url)

		if err != nil {
			if shortUrl != "" {
				h.clientError(w, http.StatusBadRequest)
				w.Write([]byte(fmt.Sprintf("For URL: %s short url already exists: %s\n", url, shortUrl)))
				return
			}
			h.serverError(w)
			return
		}

		w.Write([]byte(shortUrl))
		return
	}

	w.Header().Set("Allow", http.MethodPost+", "+http.MethodGet)
	h.clientError(w, http.StatusMethodNotAllowed)

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

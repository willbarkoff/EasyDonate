package spa

import (
	"io/ioutil"
	"net/http"
)

type Handler struct {
	IndexPath  string
	FileSystem http.FileSystem
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := h.FileSystem.Open(r.URL.Path)
	if err != nil {
		file, err := h.FileSystem.Open(h.IndexPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write(bytes)
	} else {
		http.FileServer(h.FileSystem).ServeHTTP(w, r)
	}
}

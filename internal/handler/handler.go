package handler

import (
	"fmt"
	"github.com/dsbasko/yandex-go-advanced-shortener/internal/storage"
	"io"
	"net/http"
	"strings"
)

var stor = storage.NewStorage()

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("error. response write error: %#v\n", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		if len(body) == 0 {
			fmt.Println("error. empty body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		shortURL := stor.Add(string(body))

		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte("http://" + r.Host + "/" + shortURL))
		if err != nil {
			fmt.Printf("error. response write error: %#v\n", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		fmt.Printf("info. create short url: %#v\n", shortURL)
		return
	} else if r.Method == http.MethodGet {
		path := strings.Split(r.URL.Path, "/")
		if len(path) != 2 {
			fmt.Printf("error. wrong path: %#v\n", r.URL.Path)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fullURL := stor.Get(path[1])
		if fullURL == "" {
			fmt.Printf("error. not found small url: %#v\n", path[1])
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Printf("info. redirect to: %#v\n", fullURL)
		http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)
	return mux
}

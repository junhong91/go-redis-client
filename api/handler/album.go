package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/junhong91/clean-architecture-go-redis/pkg/album"
	"github.com/junhong91/clean-architecture-go-redis/pkg/entity"
)

func albumAdd(service album.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding album"
		var a *entity.Album
		err := json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		a.ID, err = service.Store(a)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(a); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func albumFind(service album.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading album"
		vars := mux.Vars(r)
		id := vars["id"]
		album, err := service.Find(entity.StringToID(id))
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNoAlbum {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if album == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		if err := json.NewEncoder(w).Encode(album); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func MakeBookmarkHandlers(r *mux.Router, n negroni.Negroni, service album.UseCase) {
	r.Handle("/v1/album", n.With(
		negroni.Wrap(albumAdd(service)),
	)).Methods("POST", "OPTIONS").Name("albumAdd")

	r.Handle("/v1/album/{id}", n.With(
		negroni.Wrap(albumFind(service)),
	)).Methods("GET", "OPTIONS").Name("albumFind")
}

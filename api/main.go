package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/junhong91/clean-architecture-go-redis/api/handler"
	"github.com/junhong91/clean-architecture-go-redis/config"
	"github.com/junhong91/clean-architecture-go-redis/pkg/album"
)

func main() {
	rPool := &redis.Pool{
		MaxIdle:     config.REDIS_CONNECTION_POOL,
		IdleTimeout: config.REDIS_IDLE_TIMEOUT * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", config.REDIS_HOST)
		},
	}
	defer rPool.Close()

	r := mux.NewRouter()

	albumRepo := album.NewRedisRepository(rPool)
	albumService := album.NewService(albumRepo)

	n := negroni.New() //include some default middlewares

	//album
	handler.MakeBookmarkHandlers(r, *n, albumService)
	http.Handle("/", r)

	logger := log.New(os.Stderr, "logger:", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}

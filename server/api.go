package server

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/krarjun90/sample-go-api/repository"
	"github.com/lib/pq"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/krarjun90/sample-go-api/config"
	"github.com/krarjun90/sample-go-api/handler"
)

func StartApiServer() {
	fmt.Printf("Starting server at port : %v", config.AppPort())

	mux := mux.NewRouter()
	mux.HandleFunc("/ping", handler.DefaultHandler).Methods(http.MethodGet)

	connString, err := pq.ParseURL(config.DatabaseUrl())
	if err != nil {
		panic("invalid connection string, err: " + err.Error())
	}
	db := sqlx.MustConnect("postgres", connString)
	trackRepo := repository.NewTrackRepository(db)
	trackHandler := handler.NewTrackHandler(trackRepo)
	mux.HandleFunc("/tracks", trackHandler.GetAllTracks).Methods(http.MethodGet)
	mux.HandleFunc("/tracks/{id:[0-9]+}", trackHandler.GetTrackById).Methods(http.MethodGet)
	mux.HandleFunc("/tracks/{id:[0-9]+}", trackHandler.DeleteTrackById).Methods(http.MethodDelete)
	mux.HandleFunc("/tracks", trackHandler.AddTrack).Methods(http.MethodPost)

	n := negroni.Classic()
	n.UseHandler(mux)
	_ = http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(config.AppPort())), n)
}

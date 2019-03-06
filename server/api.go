package server

import (
	"fmt"
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


	n := negroni.Classic()
	n.UseHandler(mux)
	_ = http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(config.AppPort())), n)
}

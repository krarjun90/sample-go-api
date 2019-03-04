package handler

import (
	"fmt"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "pong!")
}

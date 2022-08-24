package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Handler struct {
	userLog *log.Logger
}

func NewHandler(userLog *log.Logger) *Handler {
	return &Handler{userLog}
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, r http.Request) {
	user_input, err := ioutil.ReadAll(r.Body)
	fmt.Fprintf(rw, "hi %s", string(user_input))

	if err != nil {
		http.Error(rw, "Wow mali", http.StatusBadRequest)
		return
	}
}

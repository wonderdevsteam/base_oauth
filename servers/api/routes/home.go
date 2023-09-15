package routes

import (
	"net/http"
)

func (ep *Endpoints) Home(w http.ResponseWriter, r *http.Request) {
	resp := ep.usecases.Home()
	w.Write([]byte(resp))
}

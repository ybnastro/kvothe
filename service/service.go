package service

import "net/http"

//HandlePing is the handler for a ping endpoint
func HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG!"))
}

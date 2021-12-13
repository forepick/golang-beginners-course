package main

import (
	"net/http"
)

func main(){

	requestIds := make(map[string]chan string)

	http.HandleFunc("/test/service", func(w http.ResponseWriter, r *http.Request) {

		ids, ok := r.URL.Query()["id"]

		if !ok {
			w.WriteHeader(400)
			return
		}

		id := ids[0]

		_, ok = requestIds[id]
		if !ok {
			requestIds[id] = make(chan string)
		}
		c := requestIds[id]

		<- c

		w.Write([]byte("Response"))
	})

	http.HandleFunc("/test/callback", func(w http.ResponseWriter, r *http.Request) {

		ids, ok := r.URL.Query()["id"]
		if !ok {
			w.WriteHeader(400)
			return
		}

		id := ids[0]

		_, ok = requestIds[id]
		if !ok {
			w.WriteHeader(404)
			w.Write([]byte("Request not found"))
			return
		}
		c := requestIds[id]
		c <- "Message from callback"

		delete(requestIds, id)

		w.Write([]byte("OK"))
	})

	http.ListenAndServe(":8080", nil)
}

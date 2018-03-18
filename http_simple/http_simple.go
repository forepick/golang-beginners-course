package main

import (
	"net/http"
	"encoding/json"
)

type (

	MyResponse struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
	}
)

func handleTest(w http.ResponseWriter, r *http.Request){

	names, ok := r.URL.Query()["fname"]

	fname := "Roy"
	if ok {
		fname = names[0]
	}

	respObj := &MyResponse{FirstName:fname, LastName:"Pearl"}
	respSer, err := json.Marshal(respObj)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respSer)
}


func main(){

	http.HandleFunc("/test", handleTest);
	http.ListenAndServe(":8080", nil)
}

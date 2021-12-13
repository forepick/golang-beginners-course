package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (

	MyResponse struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
	}
)

func handlePost(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(400)
		return
	}

	fmt.Fprintf(w, "Post from website! r.PostFrom = %s\n", r.Form.Get("name"))

	fname := "default"
	if val := r.Form.Get("name"); val != ""{
		fname = val
	}

	w.Write([]byte(fmt.Sprintf("fname: %s\n", fname)))
}


func handleGet(w http.ResponseWriter, r *http.Request){

	fname := "Roy"
	if names, ok := r.URL.Query()["fname"]; ok {
		fname = names[0]
	}

	lname := "Pearl"
	if names, ok := r.URL.Query()["lname"]; ok {
		lname = names[0]
	}

	respObj := &MyResponse{FirstName:fname, LastName:lname}
	respSer, err := json.Marshal(respObj)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respSer)
}

func handleJsonPost(w http.ResponseWriter, r *http.Request){

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	bodyObj := make(map[string]interface{})
	err = json.Unmarshal(body, &bodyObj)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", bodyObj)

	if fname, exists := bodyObj["name"].(string); exists {
		respObj := &MyResponse{FirstName: fname}
		resp, err := json.Marshal(respObj)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	} else {
		w.Write([]byte("nothing to report"))
	}




}

func main(){

	http.HandleFunc("/testGet", handleGet)
	http.HandleFunc("/testPost", handlePost)
	http.HandleFunc("/testJsonPost", handleJsonPost)

	http.ListenAndServe(":8080", nil)

}
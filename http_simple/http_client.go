package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)


func main(){
	client := &http.Client{}
	resp, err := client.Get("https://www.inn.co.il")

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))


}

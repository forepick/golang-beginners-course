package main

import (
	"expvar"
	"net/http"
)


func getCounter() func() interface {} {
	i := 0
	return func () interface {} {
		i++
		return i
	}
}


func main(){
	// accessible via /debug/vars

	test_var := &expvar.Int{}
	test_var.Set(14)
	expvar.Publish("number_of_iterations", test_var)

	request_counter := expvar.Func(getCounter())
	expvar.Publish("total_monitoring_requests", request_counter)

	http.ListenAndServe(":8080", nil)
}

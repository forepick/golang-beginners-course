package main

import (
	"net/http"
	"time"
)

type (
	MyHandler struct {
		rateBucket chan time.Time
		burstSize int
		maxRps float64
	}


)

func NewMyHandler(burstSize int, maxRps float64) *MyHandler{
	handler := &MyHandler{}
	handler.burstSize = burstSize
	handler.maxRps = maxRps
	handler.rateBucket = make(chan time.Time, burstSize)

	for i := 0; i < burstSize; i++ {
		handler.rateBucket <- time.Now()
	}

	fillCycle := time.Duration(1 / maxRps) * time.Second
	rateTick := time.Tick(fillCycle)
	go func(){
		for t := range rateTick {
			handler.rateBucket <- t
		}
	}()

	return handler
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	select {
	case <-h.rateBucket:
		w.WriteHeader(200)
		w.Write([]byte("Got ya!"))
	default:
		w.WriteHeader(429)
	}

}

func main() {
	handler := NewMyHandler(3, 0.2)
	http.ListenAndServe(":8080", handler)
}

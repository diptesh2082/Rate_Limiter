package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/time/rate"
)

func rateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	limiter := rate.NewLimiter(2 , 4)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if !limiter.Allow(){
			Message := Message{
				Status: "Request Failled",
				Body: "You are already in your capacity",
			}
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&Message)
			return
		}else{
			next(w,r)
		}
	})
}

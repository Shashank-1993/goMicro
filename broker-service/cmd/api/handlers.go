package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct{
	Error bool `json:"error"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
}

func (app Config) Broker(w http.ResponseWriter, r *http.Request){

responsePayload := jsonResponse{
	Error: false,
	Message: "Hit the broker",
}

out, _:= json.MarshalIndent(responsePayload, "", "\t")

w.Header().Set("Content-Type","application/json")
w.WriteHeader(http.StatusAccepted)
w.Write(out)
}
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type MockResponse struct {
	IsActive bool   `json:"active"`
	ID       string `json:"id"`
}

func startServer() {

	rt := mux.NewRouter()

	rt.HandleFunc("/isValid/{ID}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		ID := vars["ID"]

		resp := MockResponse{true, ID}
		jsonResp, err := json.Marshal(resp)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
		fmt.Fprintf(w, "")
	})

	fmt.Print("Server started ", os.Getenv("IP"), ":", os.Getenv("PORT"))

	http.ListenAndServe(os.Getenv("IP")+":"+os.Getenv("PORT"), rt)
}

func main() {
	startServer()
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func generateRandomStatus() Status {
	return Status{
		Water: rand.Intn(100) + 1,
		Wind:  rand.Intn(100) + 1,
	}
}

func updateStatusFile() {
	for {
		status := generateRandomStatus()
		statusJSON, _ := json.Marshal(status)
		err := ioutil.WriteFile("status.json", statusJSON, 0644)
		if err != nil {
			fmt.Println("Error writing to status.json:", err)
		}
		time.Sleep(15 * time.Second)
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	statusJSON, err := ioutil.ReadFile("status.json")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(statusJSON)
}

func main() {
	go updateStatusFile()

	http.HandleFunc("/status", statusHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

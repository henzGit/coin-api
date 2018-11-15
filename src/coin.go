package main

import (
	"encoding/json"
	"net/http"
)

func handleCoin(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Try handle coin2")
}


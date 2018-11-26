package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BidData struct {
	Coin string `json:"coin"`
	Amount float32 `json:"amount"`
	Price float32 `json:"price"`
	Currency string `json:"currency"`
	UserId int `json:"userId"`
}

func handleCoin(w http.ResponseWriter, req *http.Request) {
	var bidData BidData

	err := json.NewDecoder(req.Body).Decode(&bidData)
	fmt.Println(bidData)

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(bidData)
}


package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		user := User{ Name: "Taro",  Age:  20 }
		res, err := json.Marshal(user)
		if err != nil {
			log.Print("😇json marshal error.")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	case http.MethodPost:
		if req.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			log.Print("🙃application/json")
		}

		length, err := strconv.Atoi(req.Header.Get("Content-Length"))
		if err != nil {
			log.Print("😇Content-LengthがnilだからStatusInternalServerError")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			log.Print("🙃lengthは", length, "バイトです。")
		}

		var user User
		buffer := make([]byte, length)

		_, err = req.Body.Read(buffer)

		if err != nil && err != io.EOF {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(buffer, &user)
		if err != nil {
			log.Print("😇Json Unmarshal Error ")
		}

		fmt.Fprintf(w, "%v\n", user)
		w.WriteHeader(http.StatusOK)
	default:
		fmt.Fprintf(w, "Default ")
		w.WriteHeader(http.StatusOK)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Todo struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	IsDone bool   `json:"isDone"`
}

var todos []Todo = []Todo{
	Todo{0, "昨日に手を振る", true},
	Todo{1, "明日の手を取る", false},
	Todo{2, "さよならとありがとうを言う", false},
	Todo{3, "君の手を取る", true},
}

func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		fmt.Println(err)
	}
}

func TodoCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		fmt.Println(err)
	}
	todos = append(todos, todo)
}

func main() {
	router := httprouter.New()
	router.GET("/todos", TodoIndex)
	router.POST("/todos", TodoCreate)

	fmt.Println(http.ListenAndServe(":8080", router))
}

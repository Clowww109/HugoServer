package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/middleware"
	"net/http"

	"github.com/go-chi/chi"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var people []User

func main() {
	r := chi.NewRouter()

	people = append(people, User{"Max", 55})

	r.Use(middleware.Logger)

	r.Route("/hello", func(r chi.Router) {
		r.Get("/1", handleRoute1)
	})
	r.Route("/sayBye", func(r chi.Router) {
		r.Get("/b", handleRoute2)
	})
	r.Route("/users", func(r chi.Router) {
		r.Get("/list", handleRoute3)
		r.Post("/add", handleRoute3)
		r.Delete("/del", handleRoute3)
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func handleRoute1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Say hello!")
}

func handleRoute2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Say goodBye!")
}

func handleRoute3(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path == "/users/list" {
			data, err := json.Marshal(people)
			if err != nil {
				fmt.Println(err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		} else {
			fmt.Fprintf(w, "Error, wrong GET")
		}
	case "POST":
		if r.URL.Path == "/users/add" {
			fmt.Println(r.URL.Path)
			var user User
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				fmt.Println(err)
				return
			}
			people = append(people, user)
			res, _ := json.Marshal(people)
			fmt.Fprintf(w, string(res))
		}

	case "DELETE":
		if r.URL.Path == "/users/del" {
			var user User
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				fmt.Println(err)
				return
			}
			for i := 0; i < len(people); i++ {
				if user.Name == people[i].Name {
					people = append(people[:i], people[i+1:]...)
					break
				}
			}
			res, _ := json.Marshal(people)
			fmt.Fprintf(w, string(res))
		}
	}
}

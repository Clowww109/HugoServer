package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	r := chi.NewRouter()

	rp := NewReverseProxy("hugo", ":1313")

	r.Use(rp.ReverseProxy)

	r.Get("/api/", handleApiRoute)

	fmt.Println("Сервер стартует")

	go WorkerTask1()
	go WorkerTask2()
	go WorkerTask3()

	http.ListenAndServe(":8080", r)

}

type ReverseProxy struct {
	host string
	port string
}

func NewReverseProxy(host, port string) *ReverseProxy {
	return &ReverseProxy{
		host: host,
		port: port,
	}
}

func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Это верное решение или подразумевалось что-то другое?

		if !strings.HasPrefix(r.URL.Path, "/api/") {
			// Все остальные запросы перенаправляются на http://hugo:1313
			proxy := httputil.NewSingleHostReverseProxy(&url.URL{
				Scheme: "http",
				Host:   rp.host + rp.port,
			})
			//редирект, чтобы не висло с ошибкой
			http.Redirect(w, r, "http://hugo:1313/", http.StatusMovedPermanently)
			proxy.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})

}

func handleApiRoute(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from API"))
	if err != nil {
		fmt.Println(err)
		return
	}

}

func WorkerTask1() {
	for {
		res := GetTask1Page()
		err := os.WriteFile("/app/static/tasks/_index.md", []byte(res), 0644)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(5 * time.Second)
	}
}

func WorkerTask2() {
	counter := 5
	for {
		res := GetTask2Page(counter)
		err := os.WriteFile("/app/static/tasks/binary.md", []byte(res), 0644)
		if err != nil {
			log.Println(err)
		}
		counter++
		time.Sleep(5 * time.Second)
		if counter == 100 {
			counter = 5
		}
	}
}

func WorkerTask3() {
	for {
		res := GetTask3page()
		err := os.WriteFile("/app/static/tasks/graph.md", []byte(res), 0644)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(5 * time.Second)
	}
}

package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

func main() {
	r := chi.NewRouter()

	rp := NewReverseProxy("hugo", ":1313")

	r.Use(rp.ReverseProxy)

	r.Get("/api/", handleApiRoute)

	http.ListenAndServe("localhost:8080", r)
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
		fmt.Println(r.URL.Path)
		if r.URL.Path != "/api/" {
			fmt.Println(r.URL.Path)
			newProxyUrl := url.URL{
				Scheme: "http",
				Host:   net.JoinHostPort(rp.host, rp.port)}
			rp := httputil.NewSingleHostReverseProxy(&newProxyUrl)
			rp.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

const content = ``

func WorkerTest() {
	t := time.NewTicker(1 * time.Second)
	var b byte = 0
	for {
		select {
		case <-t.C:
			err := os.WriteFile("/app/static/_index.md", []byte(fmt.Sprintf(content, b)), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}

func handleApiRoute(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from API"))
	if err != nil {
		fmt.Println(err)
		return
	}

}

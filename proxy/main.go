package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	r := chi.NewRouter()

	rp := NewReverseProxy("localhost", ":1313")

	r.Use(rp.ReverseProxy)

	r.Get("/api/", handleApiRoute)

	fmt.Println("Сервер стартует")

	//WorkerTest()
	http.ListenAndServe(":8080", r)
	//http.ListenAndServe(":1313", r)

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
			http.Redirect(w, r, "http://localhost:1313/", http.StatusMovedPermanently)
			proxy.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})

}

const content = ``

//func WorkerTest() {
//	t := time.NewTicker(1 * time.Second)
//	var b byte = 0
//	for {
//		select {
//		case <-t.C:
//			err := os.WriteFile("/app/static/_index.md", []byte(fmt.Sprintf(content, b)), 0644)
//			if err != nil {
//				log.Println(err)
//			}
//			b++
//		}
//	}
//}

func handleApiRoute(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from API"))
	if err != nil {
		fmt.Println(err)
		return
	}

}

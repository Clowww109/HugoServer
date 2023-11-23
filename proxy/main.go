package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"io"
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

////func WorkerTest() {
//	// Команда для выполнения внутри контейнера
//	filePath := "/app/static/tasks/_index.md"
//	file, err := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//
//	oldData, err := io.ReadAll(file)
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	// Обрезаем файл до нулевой длины
//	err = file.Truncate(0)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	str := strings.Split(string(oldData), "\n")
//	counter := 1
//	for {
//		for i, s := range str {
//			//время
//			if strings.Contains(s, "Текущее время:") {
//				fmt.Println(s)
//				currentTimeStr := fmt.Sprintf("Текущее время: %d-%d-%d %d:%d:%d",
//					time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
//				str[i] = currentTimeStr
//			}
//			//счетчик
//			if strings.Contains(s, "Счетчик:") {
//				fmt.Println(s)
//				currentCounter := fmt.Sprintf("Счетчик: %d", counter)
//				str[i] = currentCounter
//				counter++
//			}
//		}
//		newData := strings.Join(str, "\n")
//
//		t := time.NewTicker(5 * time.Second)
//
//		for {
//			select {
//			case <-t.C:
//				err := os.WriteFile("/app/static/_index.md", []byte(newData), 0644)
//				if err != nil {
//					log.Println(err)
//				}
//			}
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

func WorkerTask1() {
	// Команда для выполнения внутри контейнера
	filePath := "/app/static/tasks/_index.md"
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	oldData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	str := strings.Split(string(oldData), "\n")
	counter := 1
	for {
		for i, s := range str {
			//время
			if strings.Contains(s, "Текущее время:") {
				fmt.Println(s)
				currentTimeStr := fmt.Sprintf("Текущее время: %d-%d-%d %d:%d:%d",
					time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
				str[i] = currentTimeStr
			}
			//счетчик
			if strings.Contains(s, "Счетчик:") {
				fmt.Println(s)
				currentCounter := fmt.Sprintf("Счетчик: %d", counter)
				str[i] = currentCounter
				counter++
			}
		}
		newData := strings.Join(str, "\n")
		err := os.WriteFile("/app/static/tasks/_index.md", []byte(newData), 0644)
		if err != nil {
			log.Println(err)
			if counter == 150 {
				break
			}
		}
		time.Sleep(5 * time.Second)
	}
}

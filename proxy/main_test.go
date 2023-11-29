package main

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewReverseProxy(t *testing.T) {
	main()
	type args struct {
		host string
		port string
	}
	tests := []struct {
		name string
		args args
		want ReverseProxy
	}{
		{"Test new Reverser Proxy", args{host: "localhost", port: ":8080"}, ReverseProxy{
			host: "localhost",
			port: ":8080",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReverseProxy(tt.args.host, tt.args.port); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("NewReverseProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseProxy_ReverseProxy(t *testing.T) {

	fakeHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	rp := &ReverseProxy{
		host: "hugo",
		port: ":1313",
	}

	handler := rp.ReverseProxy(fakeHandler)

	req := httptest.NewRequest("GET", "/path", nil)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	if result.StatusCode != http.StatusMovedPermanently {
		t.Errorf("Expected status code %d, but got %d", http.StatusMovedPermanently, result.StatusCode)
	}

	location := result.Header.Get("Location")
	expectedLocation := "http://localhost:1313/"
	if location != expectedLocation {
		t.Errorf("Expected Location header %q, but got %q", expectedLocation, location)
	}

	// Проверка, что ответ содержит ожидаемый код
	code := rec.Code
	fmt.Println(rec.Code)
	expectedCode := 301
	if code != expectedCode {
		t.Errorf("Expected body %v, but got %v", expectedCode, code)
	}
}

func Test_handleApiRoute(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		name string
	}{
		{"Test handleApi Handler"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("Get", "/api/", nil)
			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()
			handleApiRoute(recorder, req)
			if status := recorder.Code; status != http.StatusOK {
				t.Errorf("Ожидался статус %v, получен %v", http.StatusOK, status)
			}
			if str, _ := recorder.Body.ReadString('\n'); str != "Hello from API" {
				t.Errorf("Получен неверный вывод: %v", str)
			}
		})
	}
}

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var counterIndex int

func GetTask1Page() string {
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
	return ReplData(oldData)
}
func ReplData(oldData []byte) string {
	str := strings.Split(string(oldData), "\n")
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
			currentCounter := fmt.Sprintf("Счетчик: %d", counterIndex)
			str[i] = currentCounter
			counterIndex++
		}
	}
	newData := strings.Join(str, "\n")

	return newData
}

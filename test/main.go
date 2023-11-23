package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	// Команда для выполнения внутри контейнера
	filePath := "test/_index.md"
	file, err := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
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
		err := os.WriteFile("test/_index.md", []byte(newData), 0644)
		if err != nil {
			log.Println(err)
			if counter == 150 {
				break
			}
		}
		time.Sleep(5 * time.Second)
	}
}

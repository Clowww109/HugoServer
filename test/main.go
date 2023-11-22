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
	// Обрезаем файл до нулевой длины
	err = file.Truncate(0)
	if err != nil {
		log.Fatal(err)
	}

	str := strings.Split(string(oldData), "\n")
	counter := 1
	//time
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

	_, err = file.WriteString(newData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newData)
	// Сбрасываем буфер и сохраняем изменения на диск
	err = file.Sync()
	if err != nil {
		log.Fatal(err)
	}
}

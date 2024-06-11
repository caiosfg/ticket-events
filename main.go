package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8888", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func multiThreadExample() { //Trhead 1
	canal := make(chan string)

	go func() {
		canal <- "Olá, canal" //Trhead 2
	}()

	fmt.Println(<-canal) //Trhead1
}

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", echo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func echo(writer http.ResponseWriter, request *http.Request) {
	bytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.Write([]byte("echo error"))
		return
	}
	n, err := writer.Write(bytes)
	if err != nil || n != len(bytes) {
		log.Println(err, "write len:", n)
	}
}

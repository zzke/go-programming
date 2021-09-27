package main

import (
	. "go-programming/rpc/v2/api"
	"log"
)

//func main() {
//	client, err := rpc.Dial("tcp", "localhost:1234")
//	if err != nil {
//		log.Fatal("dialing:", err)
//	}
//	var reply string
//	err = client.Call(HelloServiceName+".Hello", "hello", &reply)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
}

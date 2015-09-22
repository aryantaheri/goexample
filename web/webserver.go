package main

import (
	"io"
	"net/http"
	"fmt"
	"log"
	"time"
)

type Hello struct {
	
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello from ServeHTTP")
}


type WebStruct struct {
	Greeting string
	Punct string
	Who string
}

func (web WebStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if web.Greeting != "" {
		fmt.Fprint(w, web.Greeting, "\n")
	}
	if web.Punct != "" {
		fmt.Fprint(w, web.Punct, " time is: ", time.Now())
	}
	if web.Who != "" {
		fmt.Fprint(w, web.Who, " you're: ", r.RemoteAddr)
	}
}


func serveHTTPTester() {
	http.Handle("/greeting", &WebStruct{Greeting: "Hello there, greeting", Punct: "Punct?"})
	http.Handle("/punct", WebStruct{Punct: "Hello there, Punct"})
	http.Handle("/who", WebStruct{Who: "Hello there, Who"})

	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}


func webHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.RemoteAddr + r.Host)
}

func webHandlerTester() {
	http.HandleFunc("/x", webHandler)
	http.ListenAndServe(":8000", nil)
	
}



func main() {
//	fmt.Println("webHandlerTester")
//	webHandlerTester()
	
	fmt.Println("serveHTTPTester")
	serveHTTPTester()
}

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Hello struct {
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello from ServeHTTP")
}

type WebStruct struct {
	Greeting string
	Punct    string
	Who      string
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
	io.WriteString(w, r.RemoteAddr+r.Host)
}

func webHandlerTester() {
	http.HandleFunc("/x", webHandler)
	http.ListenAndServe(":8000", nil)

}

func fetcher() {
//	url := "http://wiki.ux.uis.no/foswiki/Info/WebHome"
	url := "http://www.google.com"
	response, err := http.Get(url)
	fmt.Printf("Fetch	URL:%v\n	Reponse: %v\n	Err: %v\n", url, response, err)

	body, err := ioutil.ReadAll(response.Body)
	fmt.Printf("Body:	%q\nError:	%v\n", body, err)

	doc, err := html.Parse(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		fmt.Println("helloooooo1", n)
		
		switch n.Type {
			case html.TextNode:
				fmt.Println("NodeType: TextNode ", n.Type) 
			case html.DocumentNode:
				fmt.Println("NodeType: DocumentNode ", n.Type) 
			case html.ElementNode:
				fmt.Println("NodeType: ElementNode ", n.Type, n.Data, n.DataAtom, n.Namespace, n.Attr) 
			case html.CommentNode:
				fmt.Println("NodeType: CommentNode ", n.Type) 
			case html.DoctypeNode:
				fmt.Println("NodeType: DoctypeNode ", n.Type)
			case html.ErrorNode:
				fmt.Println("NodeType: ErrorNode ", n.Type) 
		}
		
		if n.Type == html.TextNode  {
//			for _, a := range n.Attr {
//				if a.Key == "href" {
					fmt.Printf("	helloooooo2 %T %v\n", n, n)
//					break
//				}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

//	tokenizer := html.NewTokenizer(response.Body)
//	fmt.Println("Tokenizer: ", tokenizer)
//
//	for {
//		tokenType := tokenizer.Next()
//		token := tokenizer.Token()
//		fmt.Println("	tokenType: ", tokenType)
//		fmt.Println("	token: ", token)
//		switch tokenType {
//		case html.ErrorToken:
//			fmt.Println("	ErrorToken: ", tokenizer.Err())
//			return
//		case html.TextToken:
//			fmt.Println("	TextToken: ", tokenizer.Text())
//		case html.StartTagToken:
//			tagName, _ := tokenizer.TagName()
//			fmt.Println("	StartTagToken: ", tagName)
//		default:
//
//		}
//	}
}

func main() {
	//	fmt.Println("webHandlerTester")
	//	webHandlerTester()

	//	fmt.Println("serveHTTPTester")
	//	serveHTTPTester()
	fetcher()
}

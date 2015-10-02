package main

import (
	"fmt"
	"strings"
	"sort"
	"net/http"
	"io/ioutil"
	"golang.org/x/net/html"
	"log"
	"net/url"
	"unicode"
)


type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body []string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
//		return
	}
	wordCount(body)
//	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	words.wordList = make([]string, 0)
	words.wordMap = make(map[string]int)
//	url := "http://wiki.ux.uis.no/foswiki/Info/WebHome"
	url := "http://www.vg.no/"
//	url := "http://www.ux.uis.no/~aryan/testPage.html"
	Crawl(url, 4, fetcher)
	sort.Sort(sort.Reverse(&words))
	fmt.Println("Words extracted from ", url)
	printWords(words, 1000)
}

func printWords(words Words, top int) {
	for i, word := range words.wordList {
		if word != "" && i < top {
			fmt.Printf("Rank:%v  %q #%v\n", i, word, words.wordMap[word])
		}
	}
}
// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

// returns body, []urls, error
func (f fakeFetcher) Fetch(url string) ([]string, []string, error) {
	if !strings.HasPrefix(url, "http") {
		return nil, nil, fmt.Errorf("Protocol not supported for URL: %v ", url)

	}
	response, err := http.Get(url)
	fmt.Printf("Fetch	URL:%v\n	Reponse: %q\n	Err: %v\n", url, response, err)
	if response.StatusCode != http.StatusOK {
		log.Println("Return StatusCode is not OK: ", response.Status)
		return nil, nil, fmt.Errorf("Return StatusCode for URL: %v is not OK: %s", url, response.Status)
	}
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	fmt.Printf("READALL:\n	Body:	%+q\n	Err: %v\n", body, err)
	if err != nil {
		log.Println(err)
	}
	
//	wordCount(string(body))
	node, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		log.Println(err)
	}
	
	text, urls, err := parseNode(node, url, false)
	
	fmt.Println("TEXT: ", text)
	fmt.Println("URLS: ", urls)
	fmt.Println("ERROR: ", err)
	return text, urls, err
}


func parseNode(n *html.Node, base string, includeText bool) ([]string, []string, error) {
		urls := make([]string, 0)
		text := make([]string, 0)
		
		baseUrl, err := url.Parse(base)
		if err != nil {
			log.Println(err)
		}
		switch n.Type {
			case html.TextNode:
				fmt.Println("NodeType: TextNode ")
//				fmt.Println("NodeType: TextNode ", n.Type, n.Data, n.DataAtom, n.Attr) 
			case html.DocumentNode:
				fmt.Println("NodeType: DocumentNode ", n.Data, n.Attr) 
			case html.ElementNode:
				fmt.Println("NodeType: ElementNode ", n.Data, n.Attr) 
				if strings.ToLower(n.Data) == "body" {
					includeText = true
				} else if strings.ToLower(n.Data) == "p" || strings.ToLower(n.Data) == "b" ||
							strings.ToLower(n.Data) == "em" || strings.ToLower(n.Data) == "i" ||
							strings.ToLower(n.Data) == "small" || strings.ToLower(n.Data) == "strong" ||
							strings.ToLower(n.Data) == "sub" || strings.ToLower(n.Data) == "sup" ||
							strings.ToLower(n.Data) == "ins" || strings.ToLower(n.Data) == "del" ||
							strings.ToLower(n.Data) == "mark" {
					includeText = true
				} else {
					includeText = false
				}
//			case html.CommentNode:
//				fmt.Println("NodeType: CommentNode ", n.Type) 
//			case html.DoctypeNode:
//				fmt.Println("NodeType: DoctypeNode ", n.Type)
//			case html.ErrorNode:
//				fmt.Println("NodeType: ErrorNode ", n.Type) 
//			default :
//				fmt.Println("NodeType: UNKNOW ", n.Type) 
		}
		if n.Type == html.ElementNode  {
			for _, a := range n.Attr {
				if a.Key == "href" {
					l := strings.TrimSpace(a.Val)
					lo, _ := url.Parse(l)
					if lo == nil {
						continue
					}
					if !lo.IsAbs() {
						lo = baseUrl.ResolveReference(lo)
					}
//					fmt.Printf("	Link: %v\n", lo)
					urls = append(urls, lo.String())
					break
				}
			}
		} else if n.Type == html.TextNode && includeText {
			f := func(c rune) bool { 
    			return !unicode.IsLetter(c)
    		}
			
			data := strings.TrimFunc(n.Data, f)
//			data := strings.Trim(n.Data, "1234567890`~!@#$%^&*()_+-=][\\';}{|\":/.,<>? – « »")
//			data = strings.TrimSpace(data)
			if data != "" {
				fmt.Printf("	Text: %v\n", string(data))
				text = append(text, data)
			}
		}
		
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			childText, childUrls, childError := parseNode(c, base, includeText)
			text = append(text, childText...)
			urls = append(urls, childUrls...) 
			if childError != nil {
				log.Fatal(childError)
			}
		}
		
		return text, urls, nil
		
}

type Words struct {
	wordMap map[string]int
	wordList []string
}

var words Words

func (words *Words) Len() int {
	return len(words.wordList)
}

func (words *Words) Less(i, j int) bool {
	return words.wordMap[words.wordList[i]] < words.wordMap[words.wordList[j]]
} 

func (words *Words) Swap(i, j int)  {
	words.wordList[i], words.wordList[j] = words.wordList[j], words.wordList[i] 
} 

func wordCount(bodies []string) {
	for _, body := range bodies {		
		fields := strings.Fields(body)
//		fmt.Println("wordCount")
		for _, field := range fields {
//			fmt.Printf("	field %d %s\n", i, field)
			word := cleanWord(field)
			if _, exist := words.wordMap[word]; !exist {
				words.wordList = append(words.wordList, word)
			}
			words.wordMap[word] += 1 
		}
	}
}


func cleanWord(word string) string {
	word = strings.Trim(word, " ~!@#$%^&*()_+}{|\":?></.,][';=-0987654321`\\")
	return strings.ToLower(word)
}
// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

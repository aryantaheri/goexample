package main

import (
	"fmt"
	"strings"
	"sort"
	"net/http"
	"io/ioutil"
	"io"
	"golang.org/x/net/html"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
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
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	words.wordList = make([]string, 100)
	words.wordMap = make(map[string]int)
	Crawl("http://wiki.ux.uis.no/foswiki/Info/WebHome", 4, fetcher)
	sort.Sort(sort.Reverse(&words))
	printWords(words)
}

func printWords(words Words) {
	for i, word := range words.wordList {
		if word != "" {
			fmt.Printf("Rank:%v  %v #%v\n", i, word, words.wordMap[word])
		}
	}
}
// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	response, err := http.Get(url)
	fmt.Printf("Fetch	URL:%v\n	Reponse: %v\n	Err: %v\n", url, response, err)
	if err != nil {
		// handle error
	}
//	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	fmt.Printf("Body:\n	%q\n	%v\n", body, err)
	
//	wordCount(string(body))
	
	urls := getUrls(response.Body)
	fmt.Printf("URLS:\n%v\n", urls)
	return string(body), urls, nil
//	if res, ok := f[url]; ok {
//		return res.body, res.urls, nil
//	}
//	return "", nil, fmt.Errorf("not found: %s", url)
}

func getUrls(body io.ReadCloser) []string {
	page := html.NewTokenizer(body)
	
	urls := make([]string, 0)
	fmt.Println("getUrls")
	for {
		tokenType := page.Next()
		fmt.Println("getUrls, tokenType: ", tokenType)
		switch tokenType {
			case html.TextToken:
				wordCount(string(page.Text()))
			case html.StartTagToken:
				token := page.Token()
				if token.DataAtom.String() == "a" {
					for _, attr := range token.Attr {
						if attr.Key == "href" {
							urls = append(urls, attr.Val)
						}
					}
				}
			case html.ErrorToken:
				fmt.Println("ErrorToken: ", page.Err())
				return urls
				
		}
	}
	return urls
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

func wordCount(body string) {
	fields := strings.Fields(body)
	fmt.Println("wordCount")
	for i, field := range fields {
		fmt.Printf("	field %d %s\n", i, field)
		word := cleanWord(field)
		if _, exist := words.wordMap[word]; !exist {
			words.wordList = append(words.wordList, word)
		}
		words.wordMap[word] += 1 
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

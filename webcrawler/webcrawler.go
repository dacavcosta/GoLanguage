package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var storage map[string]bool

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, fetcher Fetcher, Urls chan []string) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("found: %s %q\n", url, body)
	}	
    Urls <- urls
}

func Crawler(url string, depth int, fetcher Fetcher){
	Urls := make(chan []string)
    go Crawl(url, fetcher, Urls)
    lvl := 1
    storage[url] = true
    for i := 0; i < depth; i++ {
        for lvl > 0 {
            lvl--
            next := <- Urls
            for _, url := range next {
                if _, done := storage[url] ; !done {
                    storage[url] = true
                    lvl++
                    go Crawl(url, fetcher, Urls)
                }
            }
        }
    }
    return
}

func main() {
	storage = make(map[string]bool)
    Crawler("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
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

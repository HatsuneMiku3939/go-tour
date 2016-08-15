package main

import (
    "fmt"
)

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, done chan int) {
    defer func() { done <- 1 }()
    if depth <= 0 {
        return
    }

    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Printf("(%d) %s\n", depth, err)
        return
    }

    fmt.Printf("(%d) found: %s %q\n", depth, url, body)

    child_dones := make([] chan int, len(urls))
    for i, u := range urls {
        child_dones[i] = make(chan int)
        go Crawl(u, depth-1, fetcher, child_dones[i])
    }

    for _, child_done := range child_dones {
      <- child_done
    }
    return
}

func main() {
    done := make(chan int)
    defer func() { <- done }()

    addUrlMap("http://golang.org/")
    go Crawl("http://golang.org/", 4, fetcher, done)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
    res, ok := (*f)[url]
    if !ok {
        return "", nil, fmt.Errorf("not found: %s", url)
    }

    res.urls = Filter(res.urls, func(url string) bool {
        _, exist := urlMap[url]
        return !exist
    })

    Each(res.urls, func(url string) {
        addUrlMap(url)
    })

    return res.body, res.urls, nil
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
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

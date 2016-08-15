package main

import (
    "sync"
)

var urlMap map[string]int
var urlMapLock sync.Mutex

func init() {
  urlMap = make(map[string]int)
}

func addUrlMap(url string) {
    urlMapLock.Lock()
    defer urlMapLock.Unlock()

    urlMap[url] = urlMap[url] + 1
}


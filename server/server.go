package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())
var uuid = RandStringBytesMaskImprSrc(10)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2048)
	r.ParseForm()
	log.Println(r.Form)
	log.Println(r.PostForm)
	log.Println(r.MultipartForm)
	fmt.Fprintf(w, "Hello from %s\n", uuid)
}

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":25001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

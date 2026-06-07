// Copyright 2011. No rights reserved.

package main

import (
	"github.com/chris-piekarski/go-misc/myhttp"
	"log"
	"net/http"
)

func main() {
	//var iterations *int = flag.Int("i", 10, "number of iterations to run")
	//flag.Parse()

	myhttp.SetFileToServe("myhttp/data.html")
	http.HandleFunc("/", myhttp.MyHandler)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal(err)
	}
}

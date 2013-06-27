// Copyright 2011. No rights reserved.

package main

import (
	"http"
	"log"
	"misc/myHttp"
)

func main() {
	//var iterations *int = flag.Int("i", 10, "number of iterations to run")
	//flag.Parse()

	myHttp.SetFileToServe("./myHttp/data.html")
	http.HandleFunc("/", myHttp.MyHandler)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal(err.String())
	}
}




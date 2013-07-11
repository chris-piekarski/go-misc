package main

import (
	"time"
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Giver struct {
	Me string
	NotMe string
	Extra string
}

func processPending() {
	result := Giver{}
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	
	for {
		fmt.Print("Checking for pending queries\n")
	
		c := session.DB("pending").C("people")
		err := c.Find(bson.M{}).One(&result)
		if err == nil {
			fmt.Printf("Processing %s for %s",result.NotMe, result.Me)
			err = c.Remove(bson.M{"me":result.Me})
			if err != nil {
				fmt.Print(err)
			}
		}
		
		time.Sleep(5 * time.Second)
	}
}

func storeEntry(me string, notMe string, extra string) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

    c := session.DB("pending").C("people")
    err = c.Insert(&Giver{me, notMe, ""})
	if err != nil {
		panic(err)
	}
	fmt.Print("World")
}

func giveHandler(w http.ResponseWriter, r *http.Request) {
    me := r.FormValue("me")
    notMe := r.FormValue("notMe")
	g := &Giver{Me : me, NotMe : notMe, Extra : ""}
	
	t, _ := template.ParseFiles("start.html")
    t.Execute(w, g)
    
    storeEntry(me, notMe, "")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadFile("mag.html")
	
    w.Write(body)	
}

func main() {
    http.HandleFunc("/give", rootHandler)
    http.HandleFunc("/give/new", giveHandler)
    
    go processPending()
    fmt.Print("Starting server\n")
	http.ListenAndServe(":8080", nil)
}
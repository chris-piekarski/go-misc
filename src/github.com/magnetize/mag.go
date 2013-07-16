package main

import (
	"time"
	"fmt"
	//"io/ioutil"
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

//https://api.fullcontact.com/v2/person.json?email=bart@fullcontact.com&apiKey=a42d9db67d03a3c7
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
			err = c.Remove(bson.M{"me":result.Me, "notme":result.NotMe})
			if err != nil {
				fmt.Print(err)
			}
		}
		
		time.Sleep(10 * time.Second)
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

func giveHandler(w http.ResponseWriter, r *http.Request) *Giver {

    me := r.FormValue("me")
    notMe := r.FormValue("notMe")
	g := &Giver{Me : me, NotMe : notMe, Extra : ""}
	
	//t, _ := template.ParseFiles("start.html")
    //t.Execute(w, g)
    
    storeEntry(me, notMe, "")
    return g
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	g := &Giver{Me : "", NotMe : "", Extra : ""}

	if r.Method  == "POST" {
		fmt.Print("Got a POST")
		g = giveHandler(w, r)
	}

	//body, _ := ioutil.ReadFile("mag.html")
	//w.Write(body)	
	
	t, _ := template.ParseFiles("mag.html")
    t.Execute(w, g)
}

func main() {
    http.HandleFunc("/give", rootHandler)
    //http.HandleFunc("/give/new", giveHandler)
    
    go processPending()
    fmt.Print("Starting server\n")
	http.ListenAndServe(":8080", nil)
}
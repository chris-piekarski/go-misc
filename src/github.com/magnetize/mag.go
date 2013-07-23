package main

import (
	"encoding/json"
	"fmt"
	//"bytes"
	"html/template"
	"io/ioutil"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
	"net/smtp"
	"time"
)

type Giver struct {
	Me    string
	NotMe string
	Extra string
}

type NotMeDetails struct {
	NotMe string
	Time  string
}

type GiverList struct {
	Me    string
	NotMe []NotMeDetails
}

type FCStatus202 struct {
	Message   string
	Status    int
	RequestId string
	NotMe     string
}

type FCStatus404 struct {
	Message string
	Status  int
	NotMe   string
	Me      string
}

type FCContactInfo struct {
	FamilyName     string
	FullName       string
	GivenName      string
	Websites       []map[string]string
	Chats          []map[string]string
}

type FCInfo struct {
	Email            string
	Status           int
	Likelihood       float32
	RequestId        string
	Photos           []map[string]string
	ContactInfo      FCContactInfo
	Organizations  []map[string]string
	Demographics   map[string]string
	SocialProfiles []map[string]interface{}
	DigitalFootprint map[string]interface{}
}

func sendEmail(message string, subject string, recipient string) {
	fmt.Printf("Sending email to %s with subject %s\n", recipient, subject)
	auth := smtp.PlainAuth("", "magnetize@cpiekarski.com", "","smtp.gmail.com")
	sub := "Subject:"+subject+"\r\n\r\n"
	fullBody := sub + message
    err := smtp.SendMail("smtp.gmail.com:587", auth,
		"magnetize@cpiekarski.com", []string{recipient}, []byte(fullBody))
    if err != nil {
        fmt.Print(err)
    }
}

func getFullContact(notMe string) *http.Response {
	apiKey := "&apiKey=a42d9db67d03a3c7"
	url := "https://api.fullcontact.com/v2/person.json?email="
	r, _ := http.Get(url + notMe + apiKey)
	return r
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
			fmt.Printf("Processing %s for %s\n", result.NotMe, result.Me)

			response := getFullContact(result.NotMe)

			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			
			fmt.Print(len(contents))

			if response.StatusCode == 200 {
				var fci FCInfo
				err = json.Unmarshal(contents, &fci)
				fci.Email = result.NotMe
				fmt.Printf("%s likelihood %f\n", response.Status, fci.Likelihood)
				storeFCEntry(fci, result.NotMe)
				storeUserEntry(result.Me, result.NotMe)
				message := "We received a giving request from you for, "+result.NotMe+
					".\r\nWe'll be in touch soon!" 
				sendEmail(message, "New Giving Request", "chris@cpiekarski.com")
			} else if response.StatusCode == 202 {
				var fcs FCStatus202
				err = json.Unmarshal(contents, &fcs)
				fcs.NotMe = result.NotMe
				fmt.Printf("%s likelihood %f\n", response.Status, fcs.Message)
				storeQueuedEntry(fcs)
			} else {
			
				var fcs FCStatus404
				err = json.Unmarshal(contents, &fcs)
				fcs.NotMe = result.NotMe
				fcs.Me = result.Me
				fmt.Printf("%s %s\n", response.Status, fcs.Message)
				storeFailedEntry(fcs)
			}

			_, err = c.RemoveAll(bson.M{"me": result.Me, "notme": result.NotMe})
			if err != nil {
				fmt.Print(err)
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func storeUserEntry(me string, notMe string) {
	result := GiverList{}
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("users").C("people")
	err = c.Find(bson.M{"me": me}).One(&result)
	newGiverDetails := NotMeDetails{notMe, time.Now().String()}
	if err == nil {
		result.NotMe = append(result.NotMe, newGiverDetails)

		err = c.Update(bson.M{"me": me}, result)
	} else {
		a := []NotMeDetails{newGiverDetails}
		c.Insert(&GiverList{me, a})
	}

	if err != nil {
		panic(err)
	}
}

func storeFCEntry(fce FCInfo, notMe string) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	result := FCInfo{}
	c := session.DB("contacts").C("people")
	err = c.Find(bson.M{"email" : notMe}).One(&result)
	
	if err != nil {
		err = c.Insert(&fce)
		if err != nil {
			panic(err)
		}
	}
}

func storeFailedEntry(q FCStatus404) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("pending").C("failed")
	err = c.Insert(&q)
	if err != nil {
		panic(err)
	}
}

func storeQueuedEntry(q FCStatus202) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("pending").C("queued")
	err = c.Insert(&q)
	if err != nil {
		panic(err)
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
}

func giveHandler(w http.ResponseWriter, r *http.Request) *Giver {

	me := r.FormValue("me")
	notMe := r.FormValue("notMe")
	g := &Giver{Me: me, NotMe: notMe, Extra: ""}

	storeEntry(me, notMe, "")
	return g
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	g := &Giver{Me: "", NotMe: "", Extra: ""}

	if r.Method == "POST" {
		fmt.Print("Got a POST\n")
		g = giveHandler(w, r)
	}

	t, _ := template.ParseFiles("mag.html")
	t.Execute(w, g)
}

func main() {
	http.HandleFunc("/give", rootHandler)

	go processPending()
	fmt.Print("Starting server\n")
	http.ListenAndServe(":8080", nil)
}

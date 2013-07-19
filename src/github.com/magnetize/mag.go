package main

import (
	"encoding/json"
	"fmt"
	"bytes"
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
	Organizations  []map[string]string
	Demographics   map[string]string
	SocialProfiles []map[string]string
}

type FCInfo struct {
	Email            string
	Status           int
	Likelihood       float32
	RequestId        string
	Photos           []map[string]string
	ContactInfo      FCContactInfo
	DigitalFootprint map[string]string
	Scores           []map[string]string
}

func sendEmail(message string, recipient string) {
	fmt.Print("Sending email...\n")
	auth := smtp.PlainAuth("", "magnetize@cpiekarski.com", "","smtp.gmail.com")
    err := smtp.SendMail("smtp.gmail.com:587", auth,
		"magnetize@cpiekarski.com", []string{recipient}, []byte("Subject: Mag Test\r\n\r\n"+message))
    if err != nil {
        fmt.Print(err)
    }
    
    fmt.Print("DONE WITH EMAIL\n")
}


func sendEmail2(message string, recipient string) {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial("smtp.gmail.com:25")
	if err != nil {
		fmt.Print(err)
	}
	// Set the sender and recipient.
	c.Mail("magnitize@cpiekarski.com")
	c.Rcpt(recipient)
	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		fmt.Print(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString(message)
	if _, err = buf.WriteTo(wc); err != nil {
		fmt.Print(err)
	}
}

func getFullContact(notMe string) *http.Response {
	apiKey := "&apiKey=a42d9db67d03a3c7"
	url := "https://api.fullcontact.com/v2/person.json?email="
	r, _ := http.Get(url + notMe + apiKey)
	return r
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
			fmt.Printf("Processing %s for %s\n", result.NotMe, result.Me)

			response := getFullContact(result.NotMe)

			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)

			if response.StatusCode == 200 {
				var fci FCInfo
				err = json.Unmarshal(contents, &fci)
				fci.Email = result.NotMe
				fmt.Printf("%s likelihood %f\n", response.Status, fci.Likelihood)
				storeFCEntry(fci, result.NotMe)
				storeUserEntry(result.Me, result.NotMe)
				sendEmail("Hello, World!", "chris@cpiekarski.com")
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

	c := session.DB("contacts").C("people")
	err = c.Insert(&fce)
	if err != nil {
		panic(err)
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

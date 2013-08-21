package main

import (
	"encoding/json"
	"github.com/op/go-logging"
	"fmt"
	"html/template"
	"io/ioutil"
	//"labix.org/v2/mgo"
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
	AddedKeywords []string
	DeletedKeywords []string
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
	Organizations	[]map[string]string
	Demographics	map[string]string
	SocialProfiles	[]map[string]interface{}
	DigitalFootprint map[string]interface{}
}

var log = logging.MustGetLogger("package.mag")

func sendEmail(message string, subject string, recipient string) {
	log.Debug("Sending email to %s with subject %s\n", recipient, subject)
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
	apiKey := "&apiKey="
	url := "https://api.fullcontact.com/v2/person.json?email="
	r, _ := http.Get(url + notMe + apiKey)
	return r
}

func getFootprintTopics(notMe FCInfo) string {
	x := ""
	y := notMe.DigitalFootprint //.(map[string]interface{})
	//notMe.DigitalFootprint["topics"].Array()
	for k,v := range y {
		switch vv := v.(type) {
        case string:
            fmt.Println(k, "is string", vv)
        case int:
            fmt.Println(k, "is int", vv)
        case []interface{}:
            fmt.Println(k, "is an array:")
            if k == "topics" {
	            for i, u := range vv {
	                fmt.Println(i, u)
	//                switch vvv := u.(type) {
	//		        case interface{}:
			           f:=u.(map[string]interface{})
			           x = x+f["value"].(string)+"\r\n"
	//		        }
	            }
	         }
        default:
            fmt.Println(k, "is of a type I don't know how to handle")
        }
	}
	fmt.Print(x)
	return x
}

func getUserKeywords(me string, notMe string) string {
	addedKeys := ""
	result := GiverList{}
	session := getMongoSession()
	defer session.Close()

	c := session.DB("users").C("people")
	err := c.Find(bson.M{"me": me}).One(&result)
	
	if err == nil {
		for _, v := range result.NotMe {
			log.Debug("%s", v.NotMe)
			if (v.NotMe == notMe) {
				log.Info("Found %s entry for %s at time %s", notMe, me, v.Time)
				//TODO: remove DeletedKeywords from list
				for _, vv := range v.AddedKeywords {
					addedKeys += vv
					addedKeys += "\r\n"
				}
			}
		}
	}
	return addedKeys
}

func getKnownData(me string, notMe string) string {
	session := getMongoSession()
	defer session.Close()


	result := FCInfo{}
	c := session.DB("fullcontact").C("contacts")
	_ = c.Find(bson.M{"email" : notMe}).One(&result)
	
	stuff := result.Demographics["locationGeneral"]
	stuff = stuff + "\r\n" + getFootprintTopics(result)
	stuff = stuff + "\r\n" + getUserKeywords(me, notMe)
	return stuff
}

func processPostProcess() {
	result := Giver{}
	session := getMongoSession()
	defer session.Close()

	for {

		c := session.DB("pending").C("postprocess")
		err := c.Find(bson.M{}).One(&result)
		if err == nil {
			log.Info("Post Processing %s for %s\n", result.NotMe, result.Me)
		
	
			freeData := getKnownData(result.Me, result.NotMe)
	
			message := "We received a giving request from you for "+result.NotMe+
						".\r\nHere is what you know so far:\r\n"+freeData+"\r\n"+
						"Click here to add more data: http://magnetize.me/update"
			sendEmail(message, "New Giving Request", "chris@cpiekarski.com")
	
			time.Sleep(1000)
		}
		
		_, err = c.RemoveAll(bson.M{"me": result.Me, "notme": result.NotMe})
		if err != nil {
			fmt.Print(err)
		}
	}
}

func processPending() {
	result := Giver{}
	session := getMongoSession()
	defer session.Close()

	for {
		fmt.Print("Checking for pending requests\n")

		c := session.DB("pending").C("new")
		err := c.Find(bson.M{}).One(&result)
		if err == nil {
			log.Info("Processing %s for %s\n", result.NotMe, result.Me)

			response := getFullContact(result.NotMe)
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			
			log.Debug("Length of response contents %i", len(contents))

			if response.StatusCode == 200 {
				var fci FCInfo
				fci.Email = result.NotMe
				log.Info("%s likelihood %f\n", response.Status, fci.Likelihood)
				storeFCEntry(fci, result.NotMe)
				storeUserEntry(result.Me, result.NotMe)
				storePostProcess(result.Me, result.NotMe)
				
			} else if response.StatusCode == 202 {
				var fcs FCStatus202
				err = json.Unmarshal(contents, &fcs)
				fcs.NotMe = result.NotMe
				log.Info("%s likelihood %f queued for future search", response.Status, fcs.Message)
				storeQueuedEntry(fcs)
			} else {
			
				var fcs FCStatus404
				err = json.Unmarshal(contents, &fcs)
				fcs.NotMe = result.NotMe
				fcs.Me = result.Me
				fmt.Printf("%s %s", response.Status, fcs.Message)
				storeFailedEntry(fcs)
			}

			_, err = c.RemoveAll(bson.M{"me": result.Me, "notme": result.NotMe})
			if err != nil {
				log.Error("%s", err)
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func storePostProcess(me string, notMe string) {
	session := getMongoSession()
	defer session.Close()

	c := session.DB("pending").C("postprocess")
	
	err := c.Insert(&Giver{me, notMe, ""})
	if err != nil {
		panic(err)
	}
}

func storeUserEntry(me string, notMe string) {
	result := GiverList{}
	session := getMongoSession()
	defer session.Close()

	c := session.DB("users").C("people")
	err := c.Find(bson.M{"me": me}).One(&result)
	newGiverDetails := NotMeDetails{notMe, time.Now().String(), nil, nil}
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
	session := getMongoSession()
	defer session.Close()

	result := FCInfo{}
	c := session.DB("fullcontact").C("contacts")
	err := c.Find(bson.M{"email" : notMe}).One(&result)
	
	if err != nil {
		err = c.Insert(&fce)
		if err != nil {
			log.Error("Storing new entry returned %i", err)
		}
	}
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		me := r.FormValue("me")
		notMe := r.FormValue("notMe")
		time := r.FormValue("time")
		
		log.Debug("Edit handler (POST): %s is giving %s at %s", me, notMe, time)
	} else {
		log.Warning("No GET for give handler")
	}
	
	g := &Giver{Me: "x", NotMe: "y", Extra: "z"}
	
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, g)
}

func giveHandler(w http.ResponseWriter, r *http.Request) {

	g := &Giver{Me: "", NotMe: "", Extra: ""}

	if r.Method == "POST" {
		
		me := r.FormValue("me")
		notMe := r.FormValue("notMe")
		
		log.Debug("Give handler (POST): %s is giving %s", me, notMe)
		
		g.Me = me
		g.NotMe = notMe 

		storeNewRequest(me, notMe, "")
	} else {
		log.Warning("No GET for give handler")
	}

	t, _ := template.ParseFiles("mag.html")
	t.Execute(w, g)
}

func addHandles() {
	log.Info("Adding mag handles")
	http.HandleFunc("/give", giveHandler)
	http.HandleFunc("/edit", editHandler)
}

func main() {
	addHandles()
	go processPending()
	go processPostProcess()
	log.Debug("Starting server\n")
	http.ListenAndServe(":8080", nil)
}

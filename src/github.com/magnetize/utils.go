package main

import (
	"labix.org/v2/mgo"
	"net/smtp"
	
)

func countPendingNew() int {
	session := getMongoSession()
	defer session.Close()
	
	c := session.DB("pending").C("new")	
	x, _ := c.Count()
	return x
}

func clearPendingNew(s *mgo.Session) {
	c := s.DB("pending").C("new")	
	c.DropCollection()
}

func clearPendingFailed(s *mgo.Session) {
	c := s.DB("pending").C("failed")	
	c.DropCollection()
}

func clearPendingQueued(s *mgo.Session) {
	c := s.DB("pending").C("queued")	
	c.DropCollection()
}

func clearAllPending() {
	session := getMongoSession()
	defer session.Close()

	clearPendingNew(session)
	clearPendingFailed(session)
	clearPendingQueued(session)
}

func getMongoSession() *mgo.Session {
	session, err := mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	return session
}

func storeFailedEntry(q FCStatus404) {
	session := getMongoSession()
	defer session.Close()

	c := session.DB("pending").C("failed")
	err := c.Insert(&q)
	if err != nil {
		log.Error("Storing failed entry returned %i", err)
	}
}

func storeQueuedEntry(q FCStatus202) {
	session := getMongoSession()
	defer session.Close()

	c := session.DB("pending").C("queued")
	err := c.Insert(&q)
	if err != nil {
		log.Error("Storing queued entry returned %i", err)
	}
}

func storeNewRequest(me string, notMe string, extra string) {
	session := getMongoSession()
	c := session.DB("pending").C("new")
	defer session.Close()
	
	err := c.Insert(&Giver{me, notMe, ""})
	if err != nil {
		log.Warning("Store New Request Error %s", err)
	}
}

func sendEmail(message string, subject string, recipient string) {
	log.Debug("Sending email to %s with subject %s\n", recipient, subject)
	auth := smtp.PlainAuth("", "magnetize@cpiekarski.com", "magnetize","smtp.gmail.com")
	sub := "Subject:"+subject+"\r\n\r\n"
	fullBody := sub + message
    err := smtp.SendMail("smtp.gmail.com:587", auth,
		"magnetize@cpiekarski.com", []string{recipient}, []byte(fullBody))
    if err != nil {
        log.Error("%s", err)
    }
}
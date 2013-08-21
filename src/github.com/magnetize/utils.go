package main

import (
	"labix.org/v2/mgo"
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
package main

import (

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
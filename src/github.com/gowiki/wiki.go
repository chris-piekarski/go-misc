package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
	"regexp"
	"errors"
)

type Page struct {
    Title string
    Body  []byte
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
    	return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")
const lenPath = len("/view/")

func getTitle(w http.ResponseWriter, r *http.Request) (title string, err error) {
    title = r.URL.Path[lenPath:]
    err = nil
    if !titleValidator.MatchString(title) {
        http.NotFound(w, r)
        err = errors.New("Invalid Page Title")
    }
    return title, err
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	title, err := getTitle(w, r)
    if err != nil {
        return
    }
    
    p, err := loadPage(title)
    //handle unknow pages
    if err != nil {
    	http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
	t, _ := template.ParseFiles("view.html")
    t.Execute(w, p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
    if err != nil {
        return
    }
    
	p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    t, _ := template.ParseFiles("edit.html")
    t.Execute(w, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
    if err != nil {
        return
    }    
    
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err = p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl + ".html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func mainTwo() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))
}


//http://golang.org/doc/articles/wiki/
func main() {
	//look into Template caching
	//look into using a closure to reduce code
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
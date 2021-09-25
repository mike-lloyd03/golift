package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
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

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home")
}

func handleWorkout(w http.ResponseWriter, r *http.Request) {
	title := "Workouts"
	p, err := loadPage(title)
	if err != nil {
		log.Fatal(err)
	}
	t, _ := template.ParseFiles("templates/workouts.html")
	t.Execute(w, p)
}

func main() {
	// http.HandleFunc("/", handleHome)
	// http.HandleFunc("/workout", handleWorkout)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	DbConnect()
}

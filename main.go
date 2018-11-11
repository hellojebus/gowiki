package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

func removePage(title string) {
	var err = os.Remove(title + ".txt")
	if err != nil {
		return
	}
	fmt.Println("deleted file " + title + ".txt")
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a test page")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
	removePage("TestPage")
}

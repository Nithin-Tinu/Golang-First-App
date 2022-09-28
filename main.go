package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)
type Details struct{
	Name string
	DOB string
}
func DisplayDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method=="POST"{
		r.ParseForm()
		hello :=Details{strings.Join( r.Form["Name"],""),strings.Join(r.Form["DOB"],"")}
		// fmt.Println(r.Method)
		// fmt.Println(hello)
		outline := template.Must(template.ParseFiles("templates/DisplayDetails.html"))
		if err := outline.ExecuteTemplate(w, "DisplayDetails.html", hello); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
	
}

func fillDetails(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        t, _ := template.ParseFiles("templates/fillDetails.html")
        t.Execute(w, nil)
    }
}
func main() {
    http.HandleFunc("/DisplayDetails", DisplayDetails) 
    http.HandleFunc("/fillDetails", fillDetails)
	fmt.Println("Started")
    err := http.ListenAndServe(":5001", nil) 
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
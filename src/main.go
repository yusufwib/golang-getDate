package main

import (
	"fmt"
	"time"
)
import "html/template"
import "net/http"

type times struct {
	current 	string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var data1 = times{
			current:	time.Now().Format("2006-01-02 15:04:05 PM -07:00 Jan Mon MST"),
		}

		var data = map[string]string{
			"Times":    data1.current,
		}

		var t, err = template.ParseFiles("view.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}


		t.Execute(w, data)
		fmt.Print(data)
	})

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
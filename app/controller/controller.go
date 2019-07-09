package controller

import (
	"DateTimes/app/models"
	"fmt"
	"net/http"
	"time"
)

import "html/template"

func GetDate(w http.ResponseWriter, r *http.Request) {

	var date = models.Datetimes{
		Current: time.Now().Format("2006-01-02 15:04:05 PM -07:00 Jan Mon MST"),
	}

	var data = map[string]string{
		"Times":    date.Current,
	}

	var t, err = template.ParseFiles("app/views/view.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	t.Execute(w, data)
}

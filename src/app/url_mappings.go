package app

import (
	"github.com/lindamanf/gcnote/controllers/calendars"
	"github.com/lindamanf/gcnote/controllers/test"
)

func mapUrls() {
	router.GET("/hello", test.Test)
	router.GET("/calendars", calendars.GetCalendars)
}

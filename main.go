package main

import (
	"fmt"
	"net/http"

	"testproject/handler"
	"testproject/service"
	"github.com/goji/httpauth"
	"github.com/pressly/chi"
)

func main() {
	var (
		r             = chi.NewRouter()
		kannelService = service.NewKannelService("http://149.56.98.254:13013", "chandramouli", "Test12", service.UTF8)
		hndl          = handler.Handler{
			Kannel: kannelService,
		}
	)

	r.Use(httpauth.SimpleBasicAuth("test", "test"))

	r.Post("/outbound/sms", hndl.OutboundSmsPost)

	fmt.Println("Stating the server...")

	http.ListenAndServe(":3000", r)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/alvinashiatey/app5/pkg/config"
	"github.com/alvinashiatey/app5/pkg/handlers"
	"github.com/alvinashiatey/app5/pkg/render"
)

const port = ":8080"
var app config.AppConfig
var session *scs.SessionManager


func sever (){
	app.InProduction = false


	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil{
		fmt.Println("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println("Starting sever on port:", port)
	srv := &http.Server{
		Addr: port,
		Handler: routes(&app),
	}
	if err = srv.ListenAndServe(); err != nil{
		log.Fatal(err)
	}
}

func main() {
	sever()
}

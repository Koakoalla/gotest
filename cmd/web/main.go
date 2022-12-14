package main

import (
	"flag"
	"log"
	"text/template"
)

const version = "1.0.0"
const cssVersion = "1"

type config scruct {
	port int 
	env string
	api string 
	db struct {
		dsn string
	}
	stripe struct {
		secret string
		key string
	}
}

type application struct {
	config config
	infoLog *log.Logger
	errorLog *log.Logger
	templateCache map[string]*template.Template
	version string
}


func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application enviornment {development|production} ")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL to api ")

	flag.Parse()

	
}
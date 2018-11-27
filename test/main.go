package main

import (
	"log"
	"reflect"

	"github.com/houzhongjian/goini"
)

func main() {
	if err := goini.Init("app.conf"); err != nil {
		log.Printf("%+v\n", err)
		return
	}

	log.Println("db_name:", goini.Get("db_name"))
	log.Println("typeof:", reflect.TypeOf(goini.Get("db_name")))

	log.Println("db_host:", goini.Get("db_host"))
	log.Println("typeof:", reflect.TypeOf(goini.Get("db_host")))

	log.Println("db_port:", goini.Get("db_port"))
	log.Println("typeof:", reflect.TypeOf(goini.Get("db_port")))

	log.Println("appid:", goini.Get("appid"))
	log.Println("typeof:", reflect.TypeOf(goini.Get("appid")))

	log.Println("isUpload:", goini.Get("isUpload"))
	log.Println("typeof:", reflect.TypeOf(goini.Get("isUpload")))

	log.Println("page_size:", goini.Get("page_size"))
	log.Println("typeof:", reflect.TypeOf(goini.Get("page_size")))

}

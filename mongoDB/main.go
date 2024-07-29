package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AVVKavvk/mongoWithGO/routers"
)

func main() {
	fmt.Println("staring building netflix");
	
	r := routers.Router();

	
	log.Fatal(http.ListenAndServe(":5060",r));
	fmt.Println("server starting at port 5060")
}
package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT string = ":8000"


func routhandler(w http.ResponseWriter , r *http.Request) {
	fmt.Fprintf(w , "Api Route")
}

func main()  {
	log.Println("Http Server Starting.")


	// Serving Static File
	http.Handle("/" , http.FileServer(http.Dir("./static")))

	// Simple Rout Handler
	http.HandleFunc("/api" , routhandler)

	log.Println("Server Listen On Port" , PORT)
	err := http.ListenAndServe(PORT , nil)

	if err != nil {
		panic(err)
	}
	
}
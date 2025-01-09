package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todoingo/repositories/db"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	mongoConnectionString := os.Getenv("MONGODB_URI")

	if mongoConnectionString == "" {
		log.Fatal("MONGODB_URI environment variable not set")
	}

	db.GetMongoClient(mongoConnectionString)

	r := setupRouter()

	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

package main
 

 
import (
	"fmt"
	"movies_api/handlers"
	"log"
)
func main() {



	server := handlers.NewServer()
	fmt.Println("Server is running on port 8080")

	log.Fatal(server.ListenAndServe())


}
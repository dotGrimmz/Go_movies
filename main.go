package main
 

 
import (
	"fmt"
	"movies_api/handlers"
	"log"
)
func main() {


	fmt.Println("Server is running on port 8080")

	server := handlers.NewServer()
	log.Fatal(server.ListenAndServe())


}
package main

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)



type article struct {
	Title string 'json:"Title"'
	desc string 'json:"Title"'
	content string 'json:"Title"' 
}

var Articles []Article

// func homepage(w http.ResponseWriter, r* http.Request){
// 	fmt.Fprintf(w,"welcome to the homepage!")
// 	fmt.Println("endpoint hit homePage")
// }

func jay(w http.ResponseWriter, r *http.Request){
	fmt.Println("hi this is jay")
	fmt.Fprintf(w,"hey whatsup")

}
func handleRequests(){
	// //http.HandleFunc("/", homepage)
	// log.Fatal(http.ListenAndServe(":10000", nil))

	http.HandleFunc("/jay",jay)
	log.Fatal(http.ListenAndServe(":10000",nil))

}


func main(){
	handleRequests()
}









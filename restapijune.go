package main 


import (
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type article struct {
	id 	    string `json: "id"`
	Title   string `json:"Title"`
	desc    string `json:"desc"`
	content string `json:"content"`
}

func returnallarticles(w http.ResponseWriter , r *http.Request){
	fmt.Println("Endpoint Hit : returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func createnewarticle(w http.ResponseWriter, r * http.Request){
	reqbody , _ := ioutil.ReadAll(r.Body)

	fmt.Printf(w, "%+v" , string( reqBody))





}

func returnarticle(w http.ResponseWriter, r *http.Request){
	var := mux.vars(r)
	key := vars["id"]

	for _,article := range Articles {
		if article.id == key {
			json.NewEncoder(w).Encode(article)
		}
	}

}

func handleRequests(){
	myrouter := mux.NewRouter().StrictSlash(true)
	
	myrouter.HandleFunc("/", homePage)
	myrouter.HandleFunc("/all", returnAllArticles)
	myrouter.Handlefunc(/article/{id}, returnarticle)


	log.Fatal(http.ListenAndServe(":10000", myRouter))
}


func main(){
	Articles = []article{
		article{Title : "jay", Desc: "Article Description", content:"Article content"}
		article{Title : "batra", Desc: "Article presentation", content:"Article content"}
		
}
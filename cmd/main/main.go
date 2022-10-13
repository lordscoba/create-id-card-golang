package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lordscoba/create-id-card-golang/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterCardRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9070", r))
}

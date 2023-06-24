package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nghiavan0610/go-mysql-gorm/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":9010", r))
}
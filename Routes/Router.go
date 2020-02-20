package Routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/meftunca/gowiki/Routes/Book"
)

func RunApp() {

	router := mux.NewRouter()
	// Books Routes

	router.HandleFunc("/books", Book.Get).Methods("GET")
	router.HandleFunc("/book/{id}", Book.GetOne).Methods("GET")
	router.HandleFunc("/books", Book.Post).Methods("POST")
	router.HandleFunc("/book/{id}", Book.Delete).Methods("DELETE")
	router.HandleFunc("/book/{id}", Book.Put).Methods("PUT")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Print(err)
	}
}

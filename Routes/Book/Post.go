package Book

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/meftunca/gowiki/Structs"
)

var db *sql.DB
var err error

func Post(w http.ResponseWriter, r *http.Request) {
	book := Structs.POSTBookStruct{}

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bookJson, err := json.Marshal(book)

	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bookJson)
	// Do something with the Person struct...
	// db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// stmt, err := db.Prepare("INSERT INTO books(name,authorName,publisher,pageCount,description,language,isbn,releaseDate) VALUES(?,?,?,?,?,?,?,?)")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// _, err = stmt.Exec(title)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Fprintf(w, "New post was created")
	// defer db.Close()
}

func Put(w http.ResponseWriter, r *http.Request) {

}

func Get(w http.ResponseWriter, r *http.Request) {

}

func GetOne(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}

package main

import (
	"database/sql"  // Database SQL package to perform Query
	"log"           // Show messages on console
	"net/http"      // Manages URLs and Web Server
	"text/template" // Manage templates

	_ "github.com/go-sql-driver/mysql" // Driver Mysql to Go
)

//Struct used to display data in the template
type Names struct {
	Id    int
	Name  string
	Email string
}

// Function dbConn, opens database connection
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "crudgo"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// The tmpl variable renders all templates in the 'tmpl' folder regardless of extension
var tmpl = template.Must(template.ParseGlob("tmpl/*"))

// Function used to render the Index file
func Index(w http.ResponseWriter, r *http.Request) {
	// Open the database connection using the function dbConn()
	db := dbConn()
	// Performs database query and handles errors
	selDB, err := db.Query("SELECT * FROM names ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	// Build the struct to be used in the template
	n := Names{}

	// Build an array to store the values of the struct
	res := []Names{}

	// Performs the repetition structure taking all values from the bank
	for selDB.Next() {
		// Store values in variables
		var id int
		var name, email string

		// Do Scan the SELECT
		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		// Send results to struct
		n.Id = id
		n.Name = name
		n.Email = email

		// Join Struct with Array
		res = append(res, n)
	}

	// Opens the Index page and displays all registered on the screen
	tmpl.ExecuteTemplate(w, "Index", res)

	// Close connection with db
	defer db.Close()
}

// Show function displays only one result
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	// Get the ID from the URL parameter
	nId := r.URL.Query().Get("id")

	// Use the ID to query and handle errors
	selDB, err := db.Query("SELECT * FROM names WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	// Mount the strcut to be used in the template
	n := Names{}

	// Performs the repetition structure taking all values from the bank
	for selDB.Next() {
		// Store values in variables
		var id int
		var name, email string

		// Scan the SELECT
		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		// Send results to struct
		n.Id = id
		n.Name = name
		n.Email = email
	}

	// Show the template
	tmpl.ExecuteTemplate(w, "Show", n)

	// Close connection
	defer db.Close()

}

// Function New just displays the form for entering new data
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

// Function Edit, edit the data
func Edit(w http.ResponseWriter, r *http.Request) {
	// Open database connection
	db := dbConn()

	// Get the ID from the URL parameter
	nId := r.URL.Query().Get("id")

	selDB, err := db.Query("SELECT * FROM names WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	// Mount the struct to be used in the template
	n := Names{}

	// Performs the repetition structure taking all values from the bank
	for selDB.Next() {
		//Store values in variables
		var id int
		var name, email string

		// Do it Scan on SELECT
		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		// Send results to struct
		n.Id = id
		n.Name = name
		n.Email = email
	}

	// Shows the template with filled-in form for editing
	tmpl.ExecuteTemplate(w, "Edit", n)

	// Closes the database connection
	defer db.Close()
}

// Insert function, inserts values into the database
func Insert(w http.ResponseWriter, r *http.Request) {

	// Open the database connection using the function: dbConn()
	db := dbConn()

	// Check the METHOD of the form passed
	if r.Method == "POST" {

		// Get the form fields
		name := r.FormValue("name")
		email := r.FormValue("email")

		// Prepare SQL and check errors
		insForm, err := db.Prepare("INSERT INTO names(name, email) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}

		// Inserts values from the form with SQL handled and checks for errors
		insForm.Exec(name, email)

		// Displays a log with the values entered in the form
		log.Println("INSERT: Name: " + name + " | E-mail: " + email)
	}

	// Closes the connection of the dbConn()
	defer db.Close()

	//Return HOME
	http.Redirect(w, r, "/", 301)
}

// Function Update, updates values in the database
func Update(w http.ResponseWriter, r *http.Request) {

	// Open the database connection using the function: dbConn()
	db := dbConn()

	// 	Check the METHOD of the form passed
	if r.Method == "POST" {

		// Get the form fields
		name := r.FormValue("name")
		email := r.FormValue("email")
		id := r.FormValue("uid")

		// Prepare SQL and check errors
		insForm, err := db.Prepare("UPDATE names SET name=?, email=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}

		// Insert form values with SQL handled and check for errors
		insForm.Exec(name, email, id)

		// Displays a log with the values entered in the form
		log.Println("UPDATE: Name: " + name + " |E-mail: " + email)
	}

	// Close connection dbConn()
	defer db.Close()

	// Return HOME
	http.Redirect(w, r, "/", 301)
}

// Function Delete, delete values from the database
func Delete(w http.ResponseWriter, r *http.Request) {

	// Open database connection using function: dbConn()
	db := dbConn()

	nId := r.URL.Query().Get("id")

	// Prepare SQL and check errors
	delForm, err := db.Prepare("DELETE FROM names WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	// Insert form values with SQL handled and check for errors
	delForm.Exec(nId)

	// Displays a log with the values entered in the form
	log.Println("DELETE")

	// Close connection to dbConn()
	defer db.Close()

	// Return to HOME
	http.Redirect(w, r, "/", 301)
}

func main() {

	// Displays message that server has started
	log.Println("Server started on: http://localhost:9000")

	// Manage URLs
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)

	// Actions
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	//Start the server on port 9000
	http.ListenAndServe(":9000", nil)
}

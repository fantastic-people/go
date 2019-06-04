package main

import (
    //"database/sql"
    //_ "github.com/mattn/go-sqlite3"
    "fmt"
    "log"
    "net/http"
    "strconv"
)

type User struct {
    ID   int
    Name string
}

func (user User) String() string {
    format := `ID : %d Name: %s`
    return fmt.Sprintf(format, user.ID, user.Name)
}
func saveUser(w http.ResponseWriter, req *http.Request) {
    userId, _ := strconv.Atoi(req.FormValue("Id"))
    name := req.FormValue("name")
    out := strconv.Itoa(userId) + "-" + name
    u := &User{ID: userId, Name: name}
    log.Println(out)
    fmt.Fprintf(w, u.String())
}

func addUser(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello</h1><form action='user' method='post'><input type=text name=Id /><input type=text name=name /><input type=submit /></form>")
}
func add1User(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello  this is  second page !</h1><form action='user' method='post'><input type=text name=Id /><input type=text name=name /><input type=submit /></form>")
}
func main() {
    http.HandleFunc("/", addUser)
	http.HandleFunc("/user", saveUser)
	http.HandleFunc("/user1", add1User)
    err := http.ListenAndServe(":8880", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}
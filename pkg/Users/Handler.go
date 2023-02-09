package Users

import (
	database "FirstApp/pkg/Database"
	"fmt"
	"io/ioutil"
	"strings"

	// "fmt"
	"encoding/json"
	"net/http"
	// "io/ioutil"
)

var (
	user User
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	readUser()
	json.NewEncoder(w).Encode("you got user")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&user)
	addUser(&user)

	println("Just added user1 to database")
	json.NewEncoder(w).Encode("you just added user 1 to the cassandra db")
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	deleteUser()
	json.NewEncoder(w).Encode("you deleted user")
}
func updateHandler(w http.ResponseWriter, r *http.Request) {
	// StudentID := mux.Vars(r)["email"]

	url := r.URL.String()
	v := strings.Split(url, "/")
	// fmt.Printf("slice: %v\n", v)
	// fmt.Println(v)
	// fmt.Println(url)
	StudentID := v[len(v)-1]
	// fmt.Println("iddd", StudentID)
	var UpdateStudent User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data properly")
	}
	json.Unmarshal(reqBody, &UpdateStudent)
	// fmt.Println(UpdateStudent.FirstName)
	// fmt.Println(UpdateStudent.LastName)
	// fmt.Println(UpdateStudent.Email)
	// fmt.Println(StudentID)
	query := `UPDATE users SET first_name = ?, last_name = ? WHERE email = ?;`
	database.ExecuteUpdateQuery(query, UpdateStudent.FirstName, UpdateStudent.LastName, StudentID)
	fmt.Fprintf(w, "updated successfully")
	println("Just updated user")
	json.NewEncoder(w).Encode("you just updates user in cassandra db")
}

func deleteUser() {
	query := `TRUNCATE users;`
	database.ExecuteInsertQuery(query, user.FirstName, user.LastName, user.Email)
}

func addUser(user *User) {
	query := `INSERT INTO users(first_name,last_name,email) VALUES (?,?,?)`
	database.ExecuteInsertQuery(query, user.FirstName, user.LastName, user.Email)
}

func readUser() {
	var students []User
	student := User{}
	projections := make([]interface{}, 0)
	projections = append(projections, &student.FirstName, &student.LastName, &student.Email)
	processStudentRecord := func() {
		students = append(students, student)
	}
	readQuery := database.Query{
		Query:       "SELECT * from Users;",
		Projections: projections,
		ProcessRow:  processStudentRecord,
	}

	readQuery.ExecuteSelectQuery()

	Conv, _ := json.MarshalIndent(students, "", " ")
	fmt.Println(string(Conv))
}

//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"html/template"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"os"
//
//	"github.com/gorilla/mux"
//)
//
//type Todo struct {
//	ID          string `json:"id"`
//	Title       string `json:"title"`
//	Description string `json:"description"`
//	Created     string `json:"created"`
//}
//
//type listTodo []Todo
//
//var list = listTodo{
//	{
//		ID:          "1",
//		Title:       "Golang",
//		Description: "learn",
//		Created:     "27/10/2021",
//	},
//	{
//		ID:          "2",
//		Title:       "Golang",
//		Description: "learn",
//		Created:     "27/10/2021",
//	},
//	{
//		ID:          "3",
//		Title:       "Golang",
//		Description: "learn",
//		Created:     "27/10/2021",
//	},
//	{
//		ID:          "4",
//		Title:       "Golang",
//		Description: "learn",
//		Created:     "27/10/2021",
//	},
//	{
//		ID:          "5",
//		Title:       "Golang",
//		Description: "learn",
//		Created:     "27/10/2021",
//	},
//	{
//		ID:          "6",
//		Title:       "Golang",
//		Description: "learn",
//		Created:     "27/10/2021",
//	},
//}
//
//func homeLink(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Welcome home!")
//}
//
//func getById(w http.ResponseWriter, r *http.Request) {
//	todoID := mux.Vars(r)["id"]
//
//	for _, toDo := range list {
//		if toDo.ID == todoID {
//			json.NewEncoder(w).Encode(toDo)
//		}
//	}
//}
//
//func getAllTodos(w http.ResponseWriter, r *http.Request) {
//	//readFile()
//	json.NewEncoder(w).Encode(list)
//}
//func createTodo(w http.ResponseWriter, r *http.Request) {
//	var newTodo Todo
//	reqBody, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Fprintf(w, "ERROR")
//	}
//
//	json.Unmarshal(reqBody, &newTodo)
//	list = append(list, newTodo)
//	w.WriteHeader(http.StatusCreated)
//
//	json.NewEncoder(w).Encode(newTodo)
//	writeFIle(list)
//}
//
//func updateTodos(w http.ResponseWriter, r *http.Request) {
//	todoID := mux.Vars(r)["id"]
//	var newTodo Todo
//
//	reqBody, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Fprintf(w, "ERROR")
//	}
//	json.Unmarshal(reqBody, &newTodo)
//
//	for i, oldTodo := range list {
//		if oldTodo.ID == todoID {
//			oldTodo.Title = newTodo.Title
//			oldTodo.Description = newTodo.Description
//			oldTodo.Created = newTodo.Created
//			list = append(list[:i], oldTodo)
//			json.NewEncoder(w).Encode(oldTodo)
//		}
//	}
//	writeFIle(list)
//}
//
//func deleteTodo(w http.ResponseWriter, r *http.Request) {
//	todoID := mux.Vars(r)["id"]
//
//	for i, oldTodo := range list {
//		if oldTodo.ID == todoID {
//			list = append(list[:i], list[i+1:]...)
//			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", todoID)
//		}
//	}
//	writeFIle(list)
//}
//func renderPage(w http.ResponseWriter, r *http.Request) {
//	tmpl, err := template.ParseFiles("./template/todos.html")
//	if err != nil {
//		fmt.Fprintf(w, "ERROR")
//	}
//	tmpl.Execute(w, list)
//}
//func writeFIle(list []Todo) {
//	jsonString, _ := json.Marshal(list)
//	ioutil.WriteFile("todos.json", jsonString, os.ModePerm)
//}
//
//// Kiểm tra file tồi tại
//func Exists(name string) bool {
//	if _, err := os.Stat(name); err != nil {
//		if os.IsNotExist(err) {
//			return false
//		}
//	}
//	return true
//}
//
////Read File js
//func readFile() bool {
//	var check bool
//	path := "/home/thuongpv/GolandProjects/Lab_Go/todos.json"
//	if !Exists(path) {
//		f, err := os.Create(path)
//		defer f.Close()
//		if err != nil {
//			log.Fatalln("failed to open file", err)
//		}
//		log.Printf("Tao File")
//	}
//	if Exists(path) {
//		var list1 []Todo
//		data, err1 := ioutil.ReadFile(path)
//		if err1 != nil {
//			log.Printf("File reading error", err1)
//		}
//		if len(data) > 0 {
//			bytes := []byte(data)
//			err2 := json.Unmarshal(bytes, &list)
//			if err2 != nil {
//				log.Printf("Lỗi ReadFIle %v", err2)
//				check = false
//			}
//			fmt.Println(list)
//		}
//		if list != nil {
//			list = list1
//			check = true
//		}
//	}
//	return check
//}
//func main() {
//	router := mux.NewRouter().StrictSlash(true)
//	router.HandleFunc("/", homeLink)
//	router.HandleFunc("/page", renderPage).Methods("GET")
//	router.HandleFunc("/todo", createTodo).Methods("POST")
//	router.HandleFunc("/todos", getAllTodos).Methods("GET")
//	router.HandleFunc("/todos/{id}", getById).Methods("GET")
//	router.HandleFunc("/todos/update/{id}", updateTodos).Methods("PATCH")
//	router.HandleFunc("/todos/delete/{id}", deleteTodo).Methods("DELETE")
//	log.Fatal(http.ListenAndServe(":8080", router))
//
//}

package main

import(
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "strconv"
  "math/rand"
)

type Todos struct {
  Id string `json:"id"`
  Note string `json:"note"`
}

var todos []Todos 

func getTodos(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(todos)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var todo Todos 
  _ = json.NewDecoder(r.Body).Decode(&todo)
  todo.Id = strconv.Itoa(rand.Intn(100000000))
  todos = append(todos, todo)
  json.NewEncoder(w).Encode(todos)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for idx, todo := range todos {
    if todo.Id == params["id"] {
      todos = append(todos[:idx], todos[idx+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(todos)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var updatedNote Todos 
  _ = json.NewDecoder(r.Body).Decode(&updatedNote)
  for idx, todo := range todos {
    if todo.Id == updatedNote.Id {
      todos[idx].Note = updatedNote.Note
    }
  }
  json.NewEncoder(w).Encode(todos)
}

func main() {
  r := mux.NewRouter()
 
  todo := Todos {
    Id: "1",
    Note: "first note",
  }

  todos = append(todos, todo)

  r.HandleFunc("/getTodos", getTodos).Methods("GET")
  r.HandleFunc("/addTodo", addTodo).Methods("POST")
  r.HandleFunc("/deleteTodo/{id}", deleteTodo).Methods("DELETE")
  r.HandleFunc("/updateTodo", updateTodo).Methods("PUT")

  fmt.Println("Starting server at 8045\n")
  log.Fatal(http.ListenAndServe(":8045", r))

}


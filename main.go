package main

import( "fmt"
  "time"
  "bufio"
  "os"
  "encoding/json"
)

type Todo struct {
  Id string `json:"id"`
  Title string `json:"title"`
  Description string `json:"description"`
  CreatedAt time.Time `json:"createdAt"`
  DueDate time.Time `json:"dueDate"`
  Done bool `json:"done"`
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getInputData (inputLabel string) (string, error){
  reader := bufio.NewReader(os.Stdin)
  fmt.Print(inputLabel, ": ")
  input, err := reader.ReadString('\n')
  return input, err
}

func addTodosToFile(todo Todo){  
  jsonData, err := json.Marshal(todo)
  check(err)
  
  err = os.WriteFile("test.json", jsonData, os.ModePerm)
 	check(err)
}

func createTodo(title string, description string, createdAt time.Time, dueDate time.Time, done bool) Todo {
    return Todo{ 
        "1",
        title,
        description, 
        createdAt, 
        dueDate,
        done, 
  }
}

func main(){
  /*todos := []Todo{
    {"1", "my todo 1", "aasdsad", time.Now(),time.Now(), false},
    {"2", "my todo 2", "aasdsa2", time.Now(),time.Now(), true},
  }*/

  title, _        := getInputData("title")
  description, _  := getInputData("description")
  done            := false
  createdAt       := time.Now()
  dueDate         := time.Now()
  tt := createTodo(title, description, createdAt, dueDate, done)
  
  fmt.Println(tt)

  addTodosToFile(tt)

}

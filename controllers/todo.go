package controllers

import (
	"Todo-List/models"
	"fmt"

	//"database/sql"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var (
	gid int = 1
	//item      string
	//completed bool
	view = template.Must(template.ParseFiles("./views/index.html"))
	//database *sql.DB = config.Database()
	todos []models.Todo
)

func Show(w http.ResponseWriter, r *http.Request) {
	//statement, err := database.Query("SELECT * FROM todos")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//var todos []models.Todo

	//for statement.Next() {
	//	if err = statement.Scan(&id, &item, &completed); err != nil {
	//		log.Fatal(err)
	//	}
	//	todo := models.Todo{
	//		Id:        id,
	//		Item:      item,
	//		Completed: completed,
	//	}
	//	todos = append(todos, todo)
	//}

	data := models.View{
		Todos: todos,
	}
	_ = view.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")

	//if _, err := database.Query(`INSERT INTO TODO (item) VALUES (?)`, item); err != nil {
	//	log.Fatal(err)
	//}

	todo := models.Todo{
		Id:        gid,
		Item:      item,
		Completed: false,
	}
	gid++
	todos = append(todos, todo)

	http.Redirect(w, r, "/", http.StatusFound)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Println(id)

	//if _, err := database.Query("DELETE FROM TODO WHERE id = ?", id); err != nil {
	//	log.Fatal(err)
	//}

	if iid, err := strconv.Atoi(id); err == nil {
		for i, todo := range todos {
			if todo.Id == iid {
				todos = append(todos[:i], todos[i+1:]...)
			}
		}
	} else {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func Complete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	//if _, err := database.Query("UPDATE TODO SET completed = TRUE where id = ?", id); err != nil {
	//	log.Fatal(err)
	//}

	if iid, err := strconv.Atoi(id); err == nil {
		for i, todo := range todos {
			if todo.Id == iid {
				item := todo.Item
				todos = append(todos[:i], todos[i+1:]...)
				todos = append(todos, models.Todo{
					Id:        todo.Id,
					Item:      item,
					Completed: true,
				})

			}
		}
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

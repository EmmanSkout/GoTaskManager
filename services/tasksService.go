package services

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	models "github.com/EmmanSkout/TaskManager/models"
	repository "github.com/EmmanSkout/TaskManager/repositories"
)

var Tasks []models.Task
var tmpl *template.Template

func parseTemplate() {
	var err error
	tmpl, err = template.ParseFiles("docs/templates/task.html")
	if err != nil {
		log.Fatal(err)
	}
}
func executeTemplate(w http.ResponseWriter) {
	log.Printf("Executing template")
	err := tmpl.Execute(w, Tasks)
	if err != nil {
		log.Printf("Error executing task template, %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleModify(w http.ResponseWriter, r *http.Request) {
	ID := r.FormValue("ID")
	log.Printf("Modifying task template with ID: %s", ID)
	intId, _ := strconv.Atoi(ID)

	for i, task := range Tasks {
		if task.ID == intId {
			task.Name = r.FormValue("Name")
			task.Description = r.FormValue("Description")
			task.Complete = r.FormValue("Complete") == "on"
			task.Date = r.FormValue("Date")

			Tasks[i] = task
			repository.ModifyTask(task)
			break
		}
	}

	executeTemplate(w)
}

func initializeTasks() {
	repository.InitializeClient()
	Tasks = repository.GetTasks()
}

func HandleLoad(w http.ResponseWriter, r *http.Request) {
	initializeTasks()
	parseTemplate()
	executeTemplate(w)
}

func HandleAdd(w http.ResponseWriter, r *http.Request) {
	log.Printf("Adding new task")
	task := models.Task{
		ID:          rand.Int(),
		Name:        "Name",
		Description: "Description",
		Date:        time.Now().Format("2006-01-02"),
		Complete:    true,
	}
	repository.AddTask(task)
	initializeTasks()
	executeTemplate(w)

}

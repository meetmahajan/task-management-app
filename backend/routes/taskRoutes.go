package routes

import (
	"task-management-backend/controllers"

	"github.com/gorilla/mux"
)

func RegisterTaskRoutes(r *mux.Router) {
	r.HandleFunc("/api/tasks", controllers.GetTasks).Methods("GET")
	r.HandleFunc("/api/tasks/{id}", controllers.GetTask).Methods("GET")
	r.HandleFunc("/api/tasks", controllers.CreateTask).Methods("POST")
	r.HandleFunc("/api/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	r.HandleFunc("/api/tasks/{id}", controllers.DeleteTask).Methods("DELETE")
}

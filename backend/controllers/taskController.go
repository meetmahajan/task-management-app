package controllers

import (
	"encoding/json"
	"net/http"
	"task-management-backend/config"
	"task-management-backend/models"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	config.DB.Find(&tasks)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	config.DB.First(&task, params["id"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	config.DB.Create(&task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	config.DB.First(&task, params["id"])
	json.NewDecoder(r.Body).Decode(&task)
	config.DB.Save(&task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	config.DB.Delete(&task, params["id"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Task deleted successfully")
}

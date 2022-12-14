package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jedmarasigan/muscle_memory_main/backend/pkg/models"
	"github.com/jedmarasigan/muscle_memory_main/backend/pkg/utils"
)

var NewWorkout models.Workout

func GetWorkout(w http.ResponseWriter, r *http.Request) {
	newWorkouts := models.GetAllWorkouts()
	res, _ := json.Marshal(newWorkouts)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetWorkoutById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	workoutId := vars["workoutId"]
	id, err := strconv.ParseInt(workoutId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing in GetWorkoutById in controllers")
	}
	workoutDetails, _ := models.GetWorkoutById(id)
	res, _ := json.Marshal(workoutDetails)
	w.Header().Set("Content-Type", "pkglication/json") //Describes the content to send
	w.WriteHeader(http.StatusOK)                       //Sends status 200
	w.Write(res)
}

func CreateWorkout(w http.ResponseWriter, r *http.Request) {
	CreateWorkout := &models.Workout{}
	utils.ParseBody(r, CreateWorkout)        //Parse the json sent by user into something the DB can understand
	workout := CreateWorkout.CreateWorkout() //Refers to the CreateWorkout function defined in mm-models.go
	res, _ := json.Marshal(workout)          //Converts record sent by DB into json
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteWorkout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	workoutId := vars["workoutId"]
	id, err := strconv.ParseInt(workoutId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing in DeleteWorkout in controllers")
	}
	workout := models.DeleteWorkout(id) //Refers to the DeleteWorkout function defined in mm-models.go
	res, _ := json.Marshal(workout)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateWorkout(w http.ResponseWriter, r *http.Request) {
	var updateWorkout = &models.Workout{}
	utils.ParseBody(r, updateWorkout)
	vars := mux.Vars(r)
	workoutId := vars["workoutId"]
	id, err := strconv.ParseInt(workoutId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing in UpdateWorkout in controllers")
	}
	workoutDetails, db := models.GetWorkoutById(id)

	//Assuming the new data sent by user contains all three
	if updateWorkout.Name != "" {
		workoutDetails.Name = updateWorkout.Name
	}
	if updateWorkout.Reps != 0 {
		workoutDetails.Reps = updateWorkout.Reps
	}
	if updateWorkout.Weight >= 0 {
		workoutDetails.Weight = updateWorkout.Weight
	}

	db.Save(&workoutDetails)
	res, _ := json.Marshal(workoutDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

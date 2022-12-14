package routes

import (
	"github.com/gorilla/mux"
	"github.com/jedmarasigan/muscle_memory_main/backend/pkg/controllers"
)

var RegisterWorkoutRoutes = func(router *mux.Router) {
	router.HandleFunc("/workout/", controllers.CreateWorkout).Methods("POST")
	router.HandleFunc("/workout/", controllers.GetWorkout).Methods("GET")
	router.HandleFunc("/workout/{workoutId}", controllers.GetWorkoutById).Methods("GET")
	router.HandleFunc("/workout/{workoutId}", controllers.UpdateWorkout).Methods("PUT")
	router.HandleFunc("/workout/{workoutId}", controllers.DeleteWorkout).Methods("DELETE")
}

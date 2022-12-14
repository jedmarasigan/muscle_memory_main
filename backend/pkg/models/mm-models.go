package models

import (
	"github.com/jedmarasigan/muscle_memory_main/backend/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Person struct {
	gorm.Model
	Name       string
	BodyWeight float32
}

type Workout struct {
	gorm.Model
	Name   string
	Reps   int
	Weight float32
}

func init() {
	db = config.GetDB()
	db.AutoMigrate(&Workout{})
	db.AutoMigrate(&Person{})
}

func (w *Workout) CreateWorkout() *Workout {
	db.Create(&w)
	return w
}

func GetAllWorkouts() []Workout {
	var Workout []Workout
	db.Find(&Workout)
	return Workout
}

func GetWorkoutById(id int64) (*Workout, *gorm.DB) {
	var getWorkout Workout
	db := db.Where("ID=?", id).Find(&getWorkout)
	return &getWorkout, db
}

func DeleteWorkout(id int64) Workout {
	var workout Workout
	db.Where("ID=?", id).Delete(workout)
	return workout
}

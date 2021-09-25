package main

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Workout table should look like this:
// workoutID | user | name | workout_start   | workout_end     | exercises
// --------- | ---- | ---- | --------------- | --------------- | ----------
// 1         | Mike | A    | 20210924 081500 | 20210924 084500 | []Exercise

type Workout struct {
	gorm.Model
	User          string
	Name          string
	Workout_start time.Time
	Workout_end   time.Time
	// Exercises     []Exercise
}

type Exercise struct {
	name       string
	weight     int
	targetSets int
	targetReps int
}

func DbConnect() {
	db, err := gorm.Open(sqlite.Open("golift.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db.")
	}

	db.AutoMigrate(&Workout{})

	// newWorkout := Workout{user: "Mike", name: "A", workout_start: time.Now()}
	// db.Create(&newWorkout)
	db.Create(&Workout{User: "Mike", Name: "A", Workout_start: time.Now()})

	// var firstWorkout Workout
	// db.First(&firstWorkout, 1)
	// fmt.Println(firstWorkout)
}

# GoLift

## Purpose
I need a simple workout tracker.

## Functionality
This app should be able to:
- Create workouts by specifying which exercises are involved
- Set how many sets and reps will be done for each exercise
- Keep track of past performance

## Minimum Viable Product
- Be able to write workouts to the database
    - Workouts include workout_id, username, start_datetime, list of exercises with sets, reps, and weight completed
    - Ability to write workouts to the db
    - Ability to show weight progress over time

## Database Layout
The database will need the following tables:

### Users
Basic profile information I guess
- UID
- Username
- Email
- Custom exercises (probs just an array of strings)

### Exercises
Just a list of default exercises

### Workouts
- UID (from Users table)
- Workouts (array of structs)

## Structs
Set:
    targetReps: int
    completedReps: int

Exercise:
    Name: string
    Weight: int
    Sets:
        - array of sets

Workout:
    Name: string
    Exercises:
        - array of exercises
    Notes: string

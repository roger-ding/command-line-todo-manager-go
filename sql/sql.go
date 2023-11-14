package sql

import (
	"database/sql"
	"fmt"
  "errors"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Set global constants for database and table name
const (
  databaseLocal = "local_development"
  tableLocal = "todo_manager"
)

type TodoTask struct {
  Id int
  Task string
  Status string
}

// Initialize mysql setup 
func Initialize() {
  cfg := mysql.Config{
    User: "root", Passwd: "", Addr: "127.0.0.1:3306",
  }
  
  var err error
  db, err = sql.Open("mysql", cfg.FormatDSN())
  if err != nil {
    panic(err)
  }
  
  pingDatabase()
  createDBContent()
}

// Ensures database is up and reachable
func pingDatabase() {
  err := db.Ping()
  if err != nil {
    panic(err)
  }
}

// General SQL exec command
func dbExec(sqlCommand string) {
  _, err := db.Exec(sqlCommand)
  if err != nil {
    panic(err)
  }
}

// Initial step for creating database and table 
func createDBContent() {
  dbExec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", databaseLocal))
  dbExec(fmt.Sprintf("USE %s", databaseLocal))
  dbExec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (task varchar(255), status varchar(20))", tableLocal))
}

// Query all rows from local table and return as list of todo-tasks
func GetAllTasks() []TodoTask {
  rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", tableLocal))
  if err != nil {
    panic(err)
  }
  defer rows.Close()

  var taskList []TodoTask
  var task TodoTask
  var idCount = 1
  for rows.Next() {
    task.Id = idCount
    err := rows.Scan(&task.Task, &task.Status)
    if err != nil {
      panic(err)
    }
    taskList = append(taskList, task)
    idCount++
  }
  return taskList
}

// Retrieve single task from list of todo-tasks
func GetSingleTask(taskList []TodoTask, id int) (*TodoTask, error) {
  var taskReturn *TodoTask = nil
  for _, task := range taskList {
    if id == task.Id {
      taskReturn = &task
      break
    }
  }
  if taskReturn == nil {
    return nil, errors.New(fmt.Sprintf("Failed to get task with id %d - Please check that is a valid id!", id))
  }
  return taskReturn, nil
}

// Add task row 
func AddTask(task string, status string) {
  insertString := fmt.Sprintf("INSERT INTO %s (task, status) VALUES ('%s', '%s')", tableLocal, task, status)
  dbExec(insertString)
}

// Delete task row
func DeleteTask(task TodoTask) {
  deleteString := fmt.Sprintf("DELETE FROM %s WHERE task = '%s'", tableLocal, task.Task)
  dbExec(deleteString)
}

// Update task column status 
func UpdateTaskStatus(task TodoTask, status string) {
  updateString := fmt.Sprintf("UPDATE %s SET status = '%s' WHERE task = '%s'", tableLocal, status, task.Task)
  dbExec(updateString)
}

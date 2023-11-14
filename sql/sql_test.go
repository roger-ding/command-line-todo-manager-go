package sql

import (
	"log"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSQLSuite(t *testing.T) {
  t.Run("INIT", testInit)
  t.Run("ADD", testAdd)
  t.Run("GET", testGetTask)
  t.Run("UPDATE", testUpdate)
  t.Run("DELETE", testDelete)
}

func testInit(t *testing.T) {
  Initialize()
  log.Println("Test - Connection to database successful!")
}

func testAdd(t *testing.T) {
  AddTask("test task 1", "NOT DONE")
  AddTask("test task 2", "DONE")
  log.Println("Test - Task add successful!")
}

func testGetTask(t *testing.T) {
  allTasks := GetAllTasks()
  for _, task := range allTasks {
    if task.Task == "test task 1" {
      assert.Equal(t, task.Status, "NOT DONE")
    }
    if task.Task == "test task 2" {
      assert.Equal(t, task.Status, "DONE")
    }
  }
  log.Println("Test - Task get successful!")
}

func testUpdate(t *testing.T) {
  allTasks := GetAllTasks()
  for _, task := range allTasks {
    if task.Task == "test task 1" {
      UpdateTaskStatus(task, "DONE")
    }
    if task.Task == "test task 2" {
      UpdateTaskStatus(task, "NOT_STATUS_MESSAGE")
    }
  }
  
  allTasks2 := GetAllTasks()
  for _, task := range allTasks2 {
    if task.Task == "test task 1" {
      assert.Equal(t, task.Status, "DONE")
    }
    if task.Task == "test task 2" {
      assert.Equal(t, task.Status, "NOT_STATUS_MESSAGE")
    }
  }
  log.Println("Test - Task get all successful!")
}

func testDelete(t *testing.T) {
  allTasks := GetAllTasks()
  for _, task := range allTasks {
    if task.Task == "test task 1" {
      DeleteTask(task)
    }
    if task.Task == "test task 2" {
      DeleteTask(task)
    }
  }
  log.Println("Test - Task delete successful!")
}

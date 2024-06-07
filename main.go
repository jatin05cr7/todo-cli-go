package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"bufio"
)
const (
	fieldName ="task.json"
)
type Task struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json."completed"`
}
 var tasks []Task
 var nextID  int 

 func loadTasks()error  {
file ,err :=ioutil.ReadFile(filename)
if err ! =nil {
	 if os.IsNotExist(err){
		return nil 
	 }
	 return err 
}
err =json.Unmarshal(file ,&tasks)
 if err ! =nil {
	return err 
 }
 if len(tasks )>0{
	nextID=tasks[len(tasks)-1].ID +1
 }
 return nil 
 }
 func saveTasks() error {
file ,err :=json.MarshalIndent(tasks,""," ")
if err != nil {
	return err 
}
 err=ioutil.WriteFile(filename,file,0644)
 if err != nil {
	return err 
 }
 return nil 
 }
 func addTasks () error {
	task:=Task{
		ID:nextID,
		Title:title ,
		Completed:false
	}
	tasks=append(tasks,task)
	nextID++
	fmt.Println("Task added :",title)
 }
 func removeTask(id int ) error{
	for i task :=range tasks{
		if task.ID== id{
			tasks=append(tasks[:i],tasks[i+1:]...)
			fmt.Println("Task Removed ",task.Title)
			return
		}
	}
	fmt.Println("Task not Found :",id)
 }

 func completeTask(int id){
	for i task :=range tasks{
		if task.ID== id {
		tasks[i].Completed=true
		fmt.Println("Task Completed:",task.Title)
		return
	}
}
fmt.Println("Task Not Found :",id)
 }
 func listTask(){
	fmt.Println("To-Do List :")
	for _,task:=range tasks{
		if task.Completed{
			status="x"
		}
		fmt.Println()
	}

 }
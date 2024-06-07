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
	fileName ="task.json"
)
type Task struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json."completed"`
}
 var tasks []Task
 var nextID  int 

 func loadTasks()error  {
file ,err :=ioutil.ReadFile(fileName)
if err != nil {
	 if os.IsNotExist(err){
		return nil 
	 }
	 return err 
}
err =json.Unmarshal(file ,&tasks)
 if err !=nil {
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
 err=ioutil.WriteFile(fileName,file,0644)
 if err != nil {
	return err 
 }
 return nil 
 }
 func addTasks (title string)  {
	task:=Task{
		ID:nextID,
		Title:title ,
		Completed:false,
	}
	tasks=append(tasks,task)
	nextID++
	fmt.Println("Task added :",title)
	
 }
 func removeTask(id int ) {
	for i ,task :=range tasks{
		if task.ID== id{
			tasks=append(tasks[:i],tasks[i+1:]...)
			fmt.Println("Task Removed ",task.Title)
			return
		}
	}
	fmt.Println("Task not Found :",id)
 }

 func completeTask(id int){
	for i ,task :=range tasks{
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
		status:=" "
		if task.Completed{
			status="x"
		}
		fmt.Println("[%s] %d:%s\n",status,task.ID,task.Title)
	}

 }
 func main()  {
	err:=loadTasks()
	if err != nil {
		fmt.Println("Error for loading tasks :",err)
		return
	}
	defer func ()  {
		err:=saveTasks()
		if err !=nil {
			fmt.Println("Error Saving Tasks :",err)
		}
		
	}()
	scanner:=bufio.NewScanner(os.Stdin)
	
	for{
		fmt.Println("\nTo-Do List Application")
		fmt.Println("1. Add Task")
		fmt.Println("2. Remove Task")
		fmt.Println("3. Complete Task")
		fmt.Println("4. List  Tasks")
		fmt.Println("5.Exit")
		fmt.Println("Enter Your choice")

		scanner.Scan()
		choice:=scanner.Text()
		switch choice {
		case "1":
			fmt.Print("Enter Task title:")
			scanner.Scan()
			title:=scanner.Text()
			addTasks(title)
		case "2":
			fmt.Print("Enter task ID to remove :")
			scanner.Scan()
			id,_:=strconv.Atoi(scanner.Text())
			removeTask(id)
		case "3":
			fmt.Print("Enter task ID to mark as completed :")
			scanner.Scan()
			id , _:=strconv.Atoi(scanner.Text())
			completeTask(id)
		case "4":
			listTask()
		case "5":
			fmt.Print("Exiting Application ...")
			return 
		default:
			fmt.Println("Invalid choice . Please try again.")

			
		}
	}
	
 }
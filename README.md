# Command line TODO manager using go

## First install mysql
Installation guide -> [link](https://dev.mysql.com/doc/mysql-getting-started/en/)

## Install necessary go modules
```
rogerding@CPU command-line-todo-manager-go % pwd
$HOME/VSCode/GIT/command-line-todo-manager-go
rogerding@CPU command-line-todo-manager-go % go mod tidy
```

## Start up mysql server 
1. Open terminal
```
rogerding@CPU ~ % sudo /usr/local/mysql/support-files/mysql.server start
Starting MySQL
. SUCCESS!
```
2. Confirm mysql is accessible
```
rogerding@CPU ~ % mysql -u root -p -h localhost
Enter password:
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 8
Server version: 8.0.35 MySQL Community Server - GPL

Copyright (c) 2000, 2023, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql>
```

## Running command line todo manager go application
1. Add tasks to the todo list:
```
rogerding@CPU command-line-todo-manager-go % go run main.go add "laundry"
Added the following task to todo list: 'laundry'
rogerding@CPU command-line-todo-manager-go % go run main.go add "oil change"
Added the following task to todo list: 'oil change'
rogerding@CPU command-line-todo-manager-go % go run main.go add "grocery shopping"
Added the following task to todo list: 'grocery shopping'
```
2. View tasks in todo list: 
```
rogerding@CPU command-line-todo-manager-go % go run main.go list
☐ 1 laundry
☐ 2 oil change
☐ 3 grocery shopping
```
3. Update task in todo list:
```
rogerding@CPU command-line-todo-manager-go % go run main.go update 2 -c
Updated the following task in todo list to completed status: 'oil change'
rogerding@CPU command-line-todo-manager-go % go run main.go list       
☐ 1 laundry
☑ 2 oil change
☐ 3 grocery shopping
```
4. View tasks method can also specify only incomplete or only completed tasks:
```
rogerding@CPU command-line-todo-manager-go % go run main.go list -i
☐ 1 laundry
☐ 3 grocery shopping
rogerding@CPU command-line-todo-manager-go % go run main.go list -c
☑ 2 oil change
```
5. Remove task from todo list:
```
rogerding@CPU command-line-todo-manager-go % go run main.go remove 2
Removed the following task from todo list: 'oil change'
rogerding@CPU command-line-todo-manager-go % go run main.go list    
☐ 1 laundry
☐ 2 grocery shopping
```

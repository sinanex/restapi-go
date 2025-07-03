package main

import routs "go-todo-app/routes"

func main() {
	r := routs.SetupRouter()
	r.Run(":8000")
}

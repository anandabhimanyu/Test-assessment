package main

import (
	"testapi/Routes"
)

func main() {
	routes := Routes.SetupRouter()
	routes.Run()
}

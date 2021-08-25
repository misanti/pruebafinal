package main

import (
	"net/http"
	"pruebafinal/controllers"
)

func main() {
	controllers.RegisterControllers()
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		return 
	}

}

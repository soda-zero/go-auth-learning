package main

import (
	"go-auth/model"
	"go-auth/routes"
)

func main() {
    model.Setup()
    routes.Setup()
}

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

type EndPoint struct {
	Path         string `json:"path"`
	Method       string `json:"method"`
	ResponseFile string `json:"responseFile"`
}

type Service struct {
	Port      string     `json:"port"`
	EndPoints []EndPoint `json:"endpoints"`
}

func main() {
	app := fiber.New()

	//open service.json in current directory
	//service.json contains the service definition

	file, err := os.Open("service.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var data Service
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	println("Service Port: ", data.Port)

	for _, endpoint := range data.EndPoints {
		if endpoint.Method == "GET" {
			app.Get(endpoint.Path, func(c *fiber.Ctx) error {
				return c.SendFile(endpoint.ResponseFile)
			})
		} else if endpoint.Method == "POST" {
			app.Post(endpoint.Path, func(c *fiber.Ctx) error {
				return c.SendFile(endpoint.ResponseFile)
			})
		}
	}

	app.Listen(":" + data.Port)
}

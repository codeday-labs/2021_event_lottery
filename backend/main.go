package main

import (
    "fmt"
    "github.com/codeday-labs/2021_event_lottery/database"
    "github.com/codeday-labs/2021_event_lottery/routes"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "log"
)

func main() {
    database.Connect()

    app := fiber.New()
    app.Use(cors.New())
    
    routes.Setup(app)

    fmt.Println("Listening on localhost:4001", )
    
    log.Fatal(app.Listen(":4001"))
}
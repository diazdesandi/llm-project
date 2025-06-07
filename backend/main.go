package main

import (
  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
  "log"
  "os"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env")
  }

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  r := gin.Default()
  routes.SetupRoutes(r)
  r.Run(":" + port)
}
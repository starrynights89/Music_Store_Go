package main

import (
  "log"

	"rest"
)

func main() {
  log.Println("Main log....")
  rest.RunAPI(":9090")
}
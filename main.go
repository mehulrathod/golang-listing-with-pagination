package main

// only need mysql OR sqlite
// both are included here for reference
import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	Routers "testTask/routers"
)

func main() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	r := Routers.SetupRouter()

	port := os.Getenv("port")

	// For run on requested port
	if len(os.Args) > 1 {
		reqPort := os.Args[1]
		if reqPort != "" {
			port = reqPort
		}
	}

	if port == "" {
		port = "8080" //localhost
	}
	type Job interface {
		Run()
	}

	r.Run(":" + port)

}


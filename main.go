package main

import (
	"fmt"

	"github.com/MohiuddinAkib/my_first_goproj/config"
	"github.com/MohiuddinAkib/my_first_goproj/dotenv"
)

func main() {
	// Load environment variables
	dotenv.Load()
	// Load config
	config.Load()

	fmt.Println(config.Get(`
		app.PORT
	`))
}

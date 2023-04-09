package main

import "linsync/cmd"
import "github.com/joho/godotenv"

func main() {
	godotenv.Load(".env")
	cmd.Execute()
}

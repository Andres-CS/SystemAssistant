package main

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
)

func getEnvVar(key string) string{
	// Load .env file to the system.
	err:= godotenv.Load("../.env");
	if err != nil{
		log.Fatalf("Error loading .env file");
	}
	// Get env var.
	return os.Getenv(key);
}

func main(){
	fmt.Println("Hello World");
	aliasPrefix := getEnvVar("ALIAS_PRE_FIX");
	aliasPath := getEnvVar("ALIAS_PATH");

	files, err := os.ReadDir(aliasPath);
	if err != nil{ 
		log.Fatalf("Error readin files in dir"); 
	}

	for _, file := range files{
		fmt.Println(file.Name());
	}
}

package main

import (
	"os"
	"log"
	"reflect"
)


func main () {
	log.Print("Test")
    env := os.Getenv("USER")
	log.Println(env)
	

	fp, err := os.Create("testfile")
	if err != nil {
		log.Fatal(reflect.TypeOf(err))
		log.Fatal(err)
	}

	log.Println(reflect.TypeOf(fp))
}

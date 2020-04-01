package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("terraform", "output", "-json")
	cmd.Dir = "/Users/dshelestovski/Projects/bnc/pvp/.deploy"

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Printf("all: %q\n", out.String())

	var test interface{}
	err = json.Unmarshal(out.Bytes(), &test)
	if err != nil {
		log.Fatal(err)
	}

	root := test.(map[string]interface{})
	taskDefArn := root["task_def_arn"].(map[string]interface{})
	fmt.Println(taskDefArn["value"])

}

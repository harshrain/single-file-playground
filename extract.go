package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// m, _ := regexp.Compile("[0-9]+.[0-9]+")
	// fmt.Println(m.FindString("The time is 8.0978654s"))

	// cmd := exec.Command("papertrail", "'error tap gateway timed out'", "--min-time", "'2022-02-10 00:00:00'", "--max-time", "'2022-02-16 00:00:00'", ">", "testPapertrail.txt")
	// fmt.Println(cmd.Args)
	// fmt.Println(cmd.Run())

	// cmd := exec.Command("/Users/harsh/rain", "user:get", "-e", "harsh@rain.bh", "-c", "/Users/harsh/production.toml")

	// cmd := exec.Command("rain", "user:get", "-e", "harsh@rain.bh", "-c", "/Users/harsh/production.toml")

	//Note : String argument like email address will go withoud additional quotes

	cmd := exec.Command("./rain", "user:get", "-e", "harsh@rain.bh", "-c", "production.toml")

	//Use cmd.Dir to target a directory for which this command will be executed
	cmd.Dir = "/Users/harsh/"

	fmt.Println(cmd.Args)

	// var out bytes.Buffer
	// var stderr bytes.Buffer
	// cmd.Stdout = &out
	// cmd.Stderr = &stderr
	// err := cmd.Run()

	// if err != nil {
	// 	fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	// 	return
	// }
	// fmt.Println("Result: " + out.String())

	//Get the output of the command
	out, err := cmd.Output()

	if err != nil {
		fmt.Println("Something went wrong : ", err)
	}

	fmt.Println(string(out))

}

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// m, _ := regexp.Compile("[0-9]+.[0-9]+")
	// fmt.Println(m.FindString("The time is 8.0978654s"))
	cmd := exec.Command("papertrail", "'error tap gateway timed out'", "--min-time", "'2022-02-10 00:00:00'", "--max-time", "'2022-02-16 00:00:00'", ">", "testPapertrail.txt")
	fmt.Println(cmd.Args)
	fmt.Println(cmd.Run())
	// err := cmd.Run()
	// fmt.Println(err)
}

package main

import (
	"fmt"
	"os/exec"
	"regexp"
)

func main() {

	//Command with file write is not working. Wirkaroung - use os library and write it to file using write function
	// cmd := exec.Command("papertrail", "error tap gateway timed out", "--min-time", "2022-02-10 00:00:00", "--max-time", "2022-02-16 00:00:00", ">", "testPapertrail.txt")
	cmd := exec.Command("papertrail", "error tap gateway timed out", "--min-time", "2022-02-10 00:00:00", "--max-time", "2022-02-16 00:00:00")
	fmt.Println(cmd.Args)

	//If you want to store retrieved data in file then un-comment following code to create file
	// file, er := os.Create("sample.txt")
	// if er != nil {
	// 	panic(er)
	// }
	// defer file.Close()

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Something went wrong : ", err)
	}

	//Un-comment following code to write retrieved data into the created file
	// _, errr := file.Write(out)
	// if err != nil {
	// 	fmt.Println("Something went wrong : ", errr)
	// }

	retrievedData := string(out)

	//Setting a pattern to be searched in retrieved data
	m, _ := regexp.Compile("[0-9]+\\.[0-9]+")

	result := m.FindAllString(retrievedData, -1)

	if result != nil {
		fmt.Println(result)
	}
}

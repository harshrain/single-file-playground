package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
)

func main() {

	//Command with file write is not working. Wirkaroung - use os library and write it to file using write function
	// cmd := exec.Command("papertrail", "error tap gateway timed out", "--min-time", "2022-02-10 00:00:00", "--max-time", "2022-02-16 00:00:00", ">", "testPapertrail.txt")
	cmd := exec.Command("papertrail", "'received response for https://api.tap.company/v2/charges'", "--min-time", "2022-02-19 00:00:00", "--max-time", "2022-02-20 00:00:00")
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

	timesInString := m.FindAllString(retrievedData, -1)

	fmt.Println("Fist : ", timesInString[0])
	fmt.Println("second : ", timesInString[1])
	fmt.Println("third : ", timesInString[2])
	fmt.Println("fourth : ", timesInString[3])

	fmt.Println("Length of timesInString : ", len(timesInString))

	if timesInString != nil {

		fmt.Println("Total : ", len(timesInString))
		// sort.(result)
		var timesInFloat = []float64{}
		sum := 0.0
		for _, timeInString := range timesInString {
			timeInFloat, err := strconv.ParseFloat(timeInString, 64)
			if err != nil {
				fmt.Println("Something went wrong! Failed to covert string into int")
			}
			sum += timeInFloat
			timesInFloat = append(timesInFloat, timeInFloat)
		}

		sort.Float64s(timesInFloat)
		fmt.Println("Minimum : ", timesInFloat[0])
		fmt.Println("Maximum : ", timesInFloat[len(timesInFloat)-1])
		fmt.Println("Avg : ", sum/float64(len(timesInFloat)))
		temp := float64((95 * len(timesInFloat)) / 100)
		fmt.Println("95th Percentile : ", timesInFloat[int(temp)])
		// fmt.Println(result)
	}
}

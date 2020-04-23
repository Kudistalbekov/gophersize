package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type answer struct {
	q string
	a string
}

func main() {
	logerror := log.New(os.Stdout, "Error:", log.Ldate)
	loginfo := log.New(os.Stdout, "Info:", log.Ldate)

	givenfile := flag.String("filename", "quiz.csv", "file for csv questions and answers")
	timelimit := flag.Int("time", 10, "time for quiz")
	flag.Parse()
	c := make(chan string)
	timechannel := time.After(time.Duration(*timelimit) * time.Second)

	loginfo.Println("Program is starting")
	currentdir, err := os.Getwd()
	if err != nil {
		logerror.Fatal(err)
	}

	csvfile, err := os.Open(currentdir + "/" + *givenfile)
	if err != nil {
		logerror.Fatal(err)
	}

	reader := csv.NewReader(csvfile)
	data, err := reader.ReadAll()
	if err != nil {
		logerror.Fatal(err)
	}

	answer := parse(data)
	var correct int
	for index, val := range answer {
		fmt.Printf("Problem # %v : %v = ", index, val.q)
		go func() {
			var useranswer string
			_, err = fmt.Scanf("%s", &useranswer)
			if err != nil {
				logerror.Println(err)
			}
			c <- useranswer
		}()

		select {
		case <-timechannel:
			loginfo.Println("Time is out!")
			fmt.Printf("correct:%v\n", correct)
			return
		case useranswer := <-c:
			if useranswer == val.a {
				correct++
			}
		}
	}
}
func parse(data [][]string) []answer {
	ans := make([]answer, len(data))
	for index, value := range data {
		ans[index] = answer{
			value[0],
			strings.TrimSpace(value[1]),
		}
	}
	return ans
}

package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var MAX_TIME_FOR_QUESTION = 5 * time.Second // in seconds

func main(){
	var (
		numOfQuestions int
		numOfCorrectAnswers int
		userAnswer string

	)
	// Read CSV file
	filenamePointer := flag.String("file", "problems", "input file name for quiz")
	flag.Parse()
	fmt.Println(*filenamePointer)
	csvfile, _ := os.Open(*filenamePointer+".csv")
	reader := csv.NewReader(bufio.NewReader(csvfile))

	for{
		line, err := reader.Read()
		if(err == io.EOF){
			break;
		}else if(err != nil){
			log.Fatal("Error: ", err)
		}

		// Print the question
		numOfQuestions = numOfQuestions + 1
		fmt.Println("Question: ", line[0])

		// Wait for answer in another go routine
		ch := make(chan int)
		go func() {
			fmt.Scanln(&userAnswer)
			ch <- 1
		}()

		select {
		case <-ch:
			if(userAnswer == strings.TrimSpace(line[1])){
				numOfCorrectAnswers = numOfCorrectAnswers + 1
			}
			break
		case <-time.After(MAX_TIME_FOR_QUESTION):
			fmt.Println("Time out")
		}


	}

	fmt.Println("total questions: ", numOfQuestions)
	fmt.Println("correct answers: ", numOfCorrectAnswers)
}

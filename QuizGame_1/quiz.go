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
)

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
		numOfQuestions = numOfQuestions + 1
		fmt.Println("Question: ", line[0])
		fmt.Scanln(&userAnswer)
		if(userAnswer == strings.TrimSpace(line[1])){
			numOfCorrectAnswers = numOfCorrectAnswers + 1
		}

	}

	fmt.Println("total questions: ", numOfQuestions)
	fmt.Println("correct answers: ", numOfCorrectAnswers)
}

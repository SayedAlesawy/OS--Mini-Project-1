package fileIO

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"log"
)

//ReadFile Reads a file specified by user
func ReadFile(fileName string) []int{
	inputFile, err := os.Open(fileName + ".txt")
	defer inputFile.Close()

	if err != nil {
		fmt.Println("[Read File] Error opening input file: " + fileName);
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(inputFile)
	var arr []int 
	
	for scanner.Scan() {
		num := scanner.Text()
		numi, _ := strconv.Atoi(num)
		arr = append(arr, numi)
	}

	return arr
}

//WriteFile a function to write a file specified by user
func WriteFile(arr []int, fileName string){
	outputFile, err := os.Create(fileName)
	defer outputFile.Close()

	if err != nil {
		fmt.Println("[Write File] Error opening output file: " + fileName);
		log.Fatal(err)
	}

	writer := bufio.NewWriter(outputFile)

	for i := 0; i < len(arr); i++{
		_, err := writer.WriteString(fmt.Sprintf("%d\n", arr[i]))
		if err != nil{
			fmt.Println("[Write File] Error while converting numbers")
			log.Fatal(err)
		}
	}

	writer.Flush()
}
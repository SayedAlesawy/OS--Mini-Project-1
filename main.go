package main

import (
	"OS--Mini-Project-1/fileIO"
	"OS--Mini-Project-1/quicksort"
	"fmt"
)

//main main function
func main() {
	var inputFileName, outputFileName string
	var testInputSize string
	fmt.Print("Input file name: ")
	fmt.Scanf("%s\n", &inputFileName)

	fmt.Print("Output file name: ")
	fmt.Scanf("%s\n", &outputFileName)

	fmt.Print("Test array size: ")
	fmt.Scanf("%d\n", &testInputSize)

	intputArr := fileIO.ReadFile(inputFileName)

	quicksort.Sort(intputArr, 0, len(intputArr))

	fileIO.WriteFile(intputArr, outputFileName)

	fmt.Println("Done sorting")
}

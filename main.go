package main

import (
	"fmt"
	"OS--Mini-Project-1/quicksort"
	"OS--Mini-Project-1/fileIO"
)

//main main function
func main(){
	var inputFileName, outputFileName string
	
	fmt.Print("Input file name: ")
	fmt.Scanf("%s\n", &inputFileName)
	
	fmt.Print("Output file name: ")
	fmt.Scanf("%s\n", &outputFileName)

	fmt.Print("Test array size: ")
	fmt.Scanf("%d\n", &testInputSize)

	intputArr := ioFiles.ReadFile(inputFileName)

	quicksort.Sort(intputArr, 0, len(intputArr))

	ioFiles.WriteFile(intputArr, outputFileName)

	fmt.Println("Done sorting")
}
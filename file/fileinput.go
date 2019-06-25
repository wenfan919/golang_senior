package file

import (
	"os"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"bufio"
)

func Tst() {
	inputFile, inputError := os.Open("/Users/yanyong/workspace/go/src/golang_senior/main.go")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	//inputReader := bufio.NewReader(inputFile)
	//for {
	//	inputString, readerError := inputReader.ReadString('\n')
	//	fmt.Printf("The input was: %s", inputString)
	//	if readerError == io.EOF {
	//		return
	//	}
	//}


	inputfile := "/Users/yanyong/workspace/go/src/golang_senior/file/fileinput.go"
	outputfile := "/Users/yanyong/workspace/go/src/golang_senior/file/fileinput.gogo"

	buf, err := ioutil.ReadFile(inputfile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "file error %s\n", err)

	}

	fmt.Printf("file is : %s", buf)

	err = ioutil.WriteFile(outputfile, buf, 0644)


	if err != nil{
		panic(err.Error())
	}
}




func Tst2() {
	file, err := os.Open("/Users/yanyong/workspace/go/src/golang_senior/file/fileinput.gogo")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		// scans until newline
		if err != nil {
			fmt.Println("err: ", err)
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	suffix := filepath.Base("/Users/yanyong/workspace/go/src/golang_senior/file/fileinput.gogo")
	fmt.Println(suffix)
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}


func Tst3 () {
	// var outputWriter *bufio.Writer
	// var outputFile *os.File
	// var outputError os.Error
	// var outputString string
	outputFile, outputError := os.OpenFile("/Users/yanyong/workspace/go/src/golang_senior/file/fileinput.gogo", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i:=0; i<10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
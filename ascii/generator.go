package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"io"
)

func main(){

	if len(os.Args) != 4{
		fmt.Println("Provide all required arguments")
		return
	}

	flag := os.Args[1]
	coloredWord := os.Args[2]
	wholeString := os.Args[3]

	if !strings.HasPrefix(flag, "--color="){
		fmt.Println("go run . --color=<color> <substring to be colored> something")
		return
	}
	flag1 := strings.TrimPrefix(flag, "--color=")

	wordSlice := strings.Split(wholeString, " ")


	const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

colors := map[string]string{
	"red":    Red,
	"green":  Green,
	"yellow": Yellow,
	"blue":   Blue,
}

file, err := os.Open("standard.txt")
if err != nil {
	fmt.Println("Error", err)
}
defer file.Close()

reader := bufio.NewReader(file)

var asciiRead []string
	for {
		line, err := reader.ReadString('\n')
		asciiRead = append(asciiRead, strings.TrimRight(line, "\n"))
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading:", err)
			return
		}
	}


	for row := 0; row < 9; row++ { 
    for _, word := range wordSlice {
		index := strings.Index(word, coloredWord)
		var begin string
		var middle string
		var end string
		if index != -1{
			begin = word[:index]
			middle = word[index:index+len(coloredWord)]
			end = word[index+len(coloredWord):]
		} else{
			begin = word
			middle = ""
			end = ""
		}

		for _, char := range begin{
			start := (int(char) - 32) * 9 + 1
			fmt.Print(asciiRead[start + row])
		}
		for _, char := range middle{
			start := (int(char) - 32) * 9 + 1
			fmt.Print(colors[flag1] + asciiRead[start + row] + Reset)
		}
		for _, char := range end{
			start := (int(char) - 32) * 9 + 1
			fmt.Print(asciiRead[start + row])
		}
        fmt.Print(" ")
		
    }
    fmt.Println()
   }
}

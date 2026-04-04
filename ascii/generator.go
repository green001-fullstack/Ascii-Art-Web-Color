package ascii

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// func buildMask(text string, substring string) []bool {
// 	textRunes := []rune(text)
// 	subRunes := []rune(substring)

// 	mask := make([]bool, len(textRunes))

// 	// Edge case: empty substring → nothing to color
// 	if len(subRunes) == 0 {
// 		return mask
// 	}

// 	for i := 0; i <= len(textRunes)-len(subRunes); {
// 		match := true

// 		// Check if substring matches starting at i
// 		for j := 0; j < len(subRunes); j++ {
// 			if textRunes[i+j] != subRunes[j] {
// 				match = false
// 				break
// 			}
// 		}

// 		if match {
// 			// Mark all positions of the substring
// 			for j := 0; j < len(subRunes); j++ {
// 				mask[i+j] = true
// 			}
// 			// Move forward (no overlap)
// 			i += len(subRunes)
// 		} else {
// 			i++
// 		}
// 	}

// 	return mask
// }

func buildMask(text string, substring string)[]bool{
	textSlice := []rune(text)
	subStringSlice := []rune(substring)

	mask := make([]bool, len(textSlice))

	if len(subStringSlice) == 0 {
		return mask
	}
	for i := 0; i <= len(textSlice) - len(subStringSlice);{
		if string(textSlice[i:i+len(subStringSlice)]) == substring {
			for j := i; j < i + len(subStringSlice); j++{
				mask[j] = true
			}
			i+= len(subStringSlice)
		}else{
			i++
		}
	}
	return mask
}

func GenerateColor(text string, banner string, substring string, color string) (string, error) {


	file, err := os.Open("ascii/banners/" + banner + ".txt")
	if err != nil {
		return "", fmt.Errorf("Error opening file")
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var asciiRead []string
	for {
		line, err := reader.ReadString('\n')
		asciiRead = append(asciiRead, strings.TrimRight(line, "\r\n"))
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", fmt.Errorf("Error reading file")
		}
	}

	var result strings.Builder
	mask := buildMask(text, substring)
	textRune := []rune(text)

	for row := 0; row < 8; row++ {
		var oneLineString strings.Builder
		for i, char := range textRune{
			if char < 32 || char > 126{
				continue
			}
			start := int(char - 32) * 9 + 1
			if mask[i] == true{
				oneLineString.WriteString(`<span style="color:` + color +`">` + asciiRead[start + row] + `</span>`)
			} else {
				oneLineString.WriteString(asciiRead[start + row])
			}
		}
		result.WriteString(oneLineString.String() + "\n")
	}
	return result.String(), nil
}

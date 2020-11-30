////Small program which consists on receiving a string as an argument and outputting the string in a graphic representation of ASCII with different fonts
package asciiart

import (
	"fmt"
	"os"
	"strings"
)

// ImportASCIITxt : this function is used to import the font that will be printed
func ImportASCIITxt(filename string) []byte {
	var emptytable []byte
	file, err := os.Open("./Ascii/" + filename)
	if err != nil {
		return emptytable
	}
	fileinfo, err := file.Stat()
	if err != nil {
		return emptytable
	}
	size := fileinfo.Size()
	arr := make([]byte, size)
	file.Read(arr)
	file.Close()
	return arr
}

// GetCharLine : this function retrieve the line of a character
func GetCharLine(line int, tab []byte) string {
	count := line
	var index int
	var lastpos int
	for index = 0; count >= 0; index++ {
		if tab[index] == 13 {
			if count != 0 {
				count--
				lastpos = index
			} else {
				count--
			}
		}
	}
	return string(tab[lastpos+2 : index-1])
}

func AsciiArt(Text, font string) string {

	font += ".txt"
	arr := ImportASCIITxt(font)
	table := strings.Split(Text, "\\n")
	output := ""
	for _, str := range table {
		for i := 0; i < 8; i++ {
			for _, letter := range str {
				LetterCode := int(letter - 32)
				if LetterCode < 0 || LetterCode > 96 {
					fmt.Println("The character \"" + string(letter) + "\" isn't available, please use characters which have their ascii code between 32 and 126")
					return ""
				}
				pos := 9*(LetterCode) + 1 + i
				output = output + GetCharLine(pos, arr)
			}
			output += "\n"
		}
	}
	return output
}

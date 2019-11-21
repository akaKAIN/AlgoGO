package task_5

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func Test() {
	text_line := "aawdADWAWD AWD AWd awdawd WD Dawdawdawd ad sd aW Dawdfb jhrabf hjbdwhejbdqwgdcqwydcabdhsabdc a"
	text := strings.Split(text_line, " ")
	var input string
	if _, err := fmt.Scan(&input); err != nil {
		log.Println("Error: ", err)
	}
	regex(input, text)

}

func regex(pattern string, text []string) {
	for _, word := range text {

		matched, err := regexp.Match(pattern, []byte(word))
		if err != nil {
			log.Printf("Error matching %s", err)
		}
		if matched {
			fmt.Println("√ - ", pattern, ":", word)
		} else {
			fmt.Println("X", pattern, ":", word)
		}
	}
}

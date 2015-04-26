package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"bytes"
)

// DailyProgrammer challenge #211 (Easy) from 4/20/2015
// - http://www.reddit.com/r/dailyprogrammer/comments/338p28/20150420_challenge_211_easy_the_name_game/
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("It's the name game! (Press ctrl-c to exit.)")
	for {
		fmt.Print("What's your name? ")
		name, err := reader.ReadString('\n')
		if (err != nil) {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(RhymeName(name))
	}
}

// Returns a Name Game rhyme of the given name string.
func RhymeName(name string) string {
	name = strings.Trim(name, "!?., \r\n")
	tokens := strings.Split(name, "")
	tokens[0] = strings.ToUpper(tokens[0])

	name = strings.Join(tokens, "");
	firstChar := tokens[0]
	if (!strings.ContainsAny(tokens[0], "AEIOUY")) {
		tokens = tokens[1:]
	}
	tokens[0] = strings.ToLower(tokens[0])

	length := len(tokens)+1
	rhymedName := make([]string, length, length)
	copy(rhymedName[1:], tokens)

	rhyme := new(bytes.Buffer)
	rhyme.WriteString(fmt.Sprintf("%v, %v, %v\n", name, name, makeRhymedName("B", firstChar, rhymedName)))
	rhyme.WriteString(fmt.Sprintf("Banana fana %v\n", makeRhymedName("F", firstChar, rhymedName)))
	rhyme.WriteString(fmt.Sprintf("Fee fy %v\n", makeRhymedName("M", firstChar, rhymedName)))
	rhyme.WriteString(fmt.Sprintf("%v!\n", name))

	return rhyme.String()
}

func makeRhymedName(prefix string, firstChar string, nameTokens []string) string {
	if (prefix == firstChar) {
		nameTokens[0] = fmt.Sprintf("%vo-", prefix)
	} else {
		nameTokens[0] = fmt.Sprintf("%vo %v", strings.ToLower(prefix), prefix)
	}
	return strings.Join(nameTokens, "")
}

func debugString(sample string) {
	fmt.Println("Byte loop:")
    for i := 0; i < len(sample); i++ {
        fmt.Printf("%+q ", sample[i])
    }
    fmt.Printf("\n")
}

////// References
// http://en.wikipedia.org/wiki/The_Name_Game#Rules
//// Strings
// https://blog.golang.org/strings
// http://golang.org/pkg/strings/
// http://golang.org/pkg/unicode/
// http://golang.org/pkg/unicode/utf8/
//// Input
// http://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line
// http://golang.org/pkg/os/#Stdin
// http://golang.org/pkg/bufio/#Reader
//// Output
// http://golang.org/pkg/bytes/#Buffer
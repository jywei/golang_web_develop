package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "A cross-site scripting vulnerability may be used by attackers to bypass access controls such as the: you have a script on one site that makes a request to another site. For example: you come to my cool website about kittens, and a script runs to transfer money from UnionBank to my foreign account. If it wasn't for the 'same-origin policy' implemented in browsers, and if you had a cookie on your machine that said you were logged into Union Bank, then the money would transfer."

	scanner := bufio.NewScanner(strings.NewReader(s))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

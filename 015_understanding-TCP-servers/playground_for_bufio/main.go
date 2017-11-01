package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "Cross-site scripting carried out on websites accounted for roughly 84% of all security vulnerabilities documented by Symantec as of 2007. Their effect may range from a petty nuisance to a significant security risk, depending on the sensitivity of the data handled by the vulnerable site and the nature of any security mitigation implemented by the site's owner."

	scanner := bufio.NewScanner(strings.NewReader(s))

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

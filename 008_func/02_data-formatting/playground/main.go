package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	tf := t.Format(time.Kitchen)
	fmt.Println(tf)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	var S int
	var T time.Time
	fmt.Scan(&S)
	T = T.Add(time.Duration(S) * time.Second)
	duration := T.Sub(time.Time{})
	fmt.Printf("%v:%s\n", int(duration.Hours()), T.Format("4:5"))
}

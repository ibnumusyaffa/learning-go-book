package main

import "fmt"

type Score int
type HighScore Score

func main() {
	var i int = 100
	var s Score = 200
	var hs HighScore = 300

	//	hs = s            // compilation error!
	//	s = i             // compilation error!
	hs = HighScore(s) // ok
	s = Score(i)      // ok

	fmt.Println(i, s, hs)
}

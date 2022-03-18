package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Score int
type Converter func(string) Score
type TeamScores map[string]Score

type HighScore Score
type Employee Person

func (s Score) say() {
	fmt.Println("inside say", s)
}

func main() {
	var i int = 300
	var s Score = 100
	var hs HighScore = 200

	//hs = s                  // compilation error!
	//s = i                   // compilation error!
	s = Score(i)      // ok
	hs = HighScore(s) // ok

	fmt.Println(hs)

	s.say()
	//hs.say() highScore does not have say method
}

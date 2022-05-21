package main

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

func main() {

}

package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Score int

type Converter func(name string, age int) Score
type Converter2 func(string, int) Score

type TeamScores map[string]Score

type HighScore Score
type Employee Person

func main() {

}

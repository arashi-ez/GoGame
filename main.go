/*
main idea of the game is to make a different probabilities and possibilities
by the changing in that way the game always will be different
*/

//next change is to add functions which will applied randomly
//eg. first time player should get 1 way of situation but in the next circle of game
//player could take different story

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var name string
var age int

// this func will ask a namer of the hero and initial data
func namer() (int, string) {

	fmt.Println("What is your name?")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return 0, ""
	}
	fmt.Println("How old are you?")
	_, err2 := fmt.Scanln(&age)
	if err2 != nil {
		return 0, ""
	}
	return age, name
}

// this func will display all jobs and set it randomly.
// it could be or switch statement or an array
func jobs() string {
	var cur int
	var jobName [10]string

	return cur
}

// this func will make a random number generator
func randomizer() int {

	fmt.Println("Start generate your life!")

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	randee := r1.Intn(100)
	fmt.Println(randee)
	return randee
}

// this func will save all possible variant of actions
func possibilities(a int, b int, s string) {
	if a < 10 && a >= 0 {
		fmt.Printf("%s is a %d y.o king of UK", s, b)

	} else if a < 20 && a >= 10 {
		fmt.Printf("%s is a %d y.o streamer with 1b subs in YouTube", s, b)

	} else if a < 30 && a >= 20 {
		fmt.Printf("%s is a %d y.o average teenager who secretly loves his stepmom", s, b)

	} else if a < 40 && a >= 30 {
		fmt.Printf("%s is a %d y.o stepmom of an average teenager", s, b)

	} else if a < 50 && a >= 40 {
		fmt.Printf("%s is a %d y.o second-time-merried porn-producer ", s, b)

	} else {
		fmt.Printf("%s is a %d y.o GOAT of every sport all over the world", s, b)

	}
}

// func _empty() this func will show some extra data

// Here main code where we will apply some actions in the story game
func main() {
	namer()
	randee := randomizer()
	possibilities(randee, age, name)
}

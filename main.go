package main

import (
	"fmt"
	"math/rand"
	"time"
)

var name string
var age int

// func _empty() this func will show some extra data

// this func will save all possible variant of actions
func possibilities(int) {
	
}

//this func will make a random number generator
func randomizer() int {

	var seeding int

	fmt.Println("Start generate your life!")
	fmt.Println("Here you need to set the borders of your random number!")

	_, err := fmt.Scanln(&seeding)
	if err != nil {
		return 0
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	randee := r1.Intn(100)
	fmt.Println(randee)
	return randee
}

func namer() { //this func will ask a namer of the hero and initial data

	fmt.Println("What is your name?")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}
	fmt.Println("How old are you?")
	_, err2 := fmt.Scanln(&age)
	if err2 != nil {
		return
	}
	fmt.Printf("Hi, %s, you are %d!", name, age)
}

func main() {
	//Here main code where we will apply some actions in the story game
	namer()
	randee := randomizer()
	possibilities(randee)
}

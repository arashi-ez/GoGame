package main

import "fmt"

var name string
var age int

// func _empty() this func will show some extra data
// func _empty() this func will make a random number generator
// func _empty() this func will save all possible variant of actions

func namer() { //this func will ask a namer of the hero and initial data
	fmt.Println("What is your namer?")
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
}

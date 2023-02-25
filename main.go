/*
package main

import (

	"fmt"
	"math/rand"

)

var name string
var age int

// func _empty() this func will show some extra data
// func _empty() this func will save all possible variant of actions

func randomizer() int { //this func will make a random number generator

		var seeding int

		fmt.Println("Start generate your life!")
		fmt.Println("Here you need to set the borders of your random number!")

		_, err := fmt.Scanln(&seeding)
		if err != nil {
			return 0
		}
		random := rand.Intn(seeding)
		return random
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
		see := randomizer()
		ran := rand.Intn(see)
		fmt.Printf("so rand is %d", ran)
	}
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var word int
	fmt.Println("smth: ")
	_, err := fmt.Scanln(&word)
	if err != nil {
		return
	}
	random := rand.Intn(word)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 10; i++ {
		word++
		fmt.Printf("word is: %d, %d\n", word, random)

	}
	for i := 0; i < 10; i++ {
		word++
		fmt.Printf("word is: %d, %d\n", word, r1)

	}
	//best randomizer
	for i := 0; i < 10; i++ {
		word++
		fmt.Print("\n", word, " - ", r1.Intn(100), ",")
		fmt.Print(r1.Intn(100))

	}

	fmt.Println(word)
}

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

// this func will ask a namer of the hero and initial data
func namer() (int, string) {
	var name string
	var age int
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
	fmt.Println("Let's generate your life?")
	return age, name
}

// this func will display all jobs and set it randomly.
func jobs() string {

	randee := randomizer()

	jobName := [100]string{"Porn actor", "King", "Slave", "Bot", "Feminist", "Gay",
		"Gigachad", "Policeman", "Step-mom", "Data Scientist", "Dentist", "Nurse",
		"Pharmacist", "Computer Systems Analyst", "Physician", "Database Administrator", "Software Developer", "Physical Therapist", "Web Developer",
		"Dental Hygienist", "Occupational Therapist", "Veterinarian", "Computer Programmer", "School Psychologist", "Physical Therapist Assistant", "Interpreter & Translator",
		"Mechanical Engineer", "Technician", "Epidemiologist", "IT Manager", "Research Analyst", "Sonographer", "Computer Systems",
		"Civil Engineer", "Abuse Counselor", "Pathologist", "Landscape", "Groundskeeper", "Cost Estimator", "Financial Advisor",
		"Medical Assistant", "Lawyer", "Accountant", "Compliance Officer", "Teacher", "Repair Worker", "Fitness Worker",
		"Health Aide", "Dental Assistant", "Pharmacy Technician", "Relations Specialist", "Paramedic", "Hairdresser", "Preschool Teacher",
		"Marketing Manager", "Patrol Officer", "School Counselor", "Assistant", "Analyst", "Social Worker", "Event Planner",
		"Counselor", "Sales Representative", "Nursing Aide", "Sales Manager", "Architect", "Sales Manager", "HR Specialist",
		"Plumber", "Real Estate Agent", "Glazier", "Art Director", "Logistician", "Auto Mechanic", "Bus Driver",
		"Restaurant Cook", "Social Worker", "Receptionist", "Paralegal", "Cement Mason", "Concrete Finisher", "Painter",
		"Coach", "Concrete Finisher", "Brick-mason", "Block-mason", "Cashier", "Janitor", "Electrician",
		"Truck Driver", "Housekeeper", "Maid", "Carpenter", "Security Guard", "Worker", "Fabricator",
		"Fireman", "Whore", "Slut", "[ ]"}
	//fmt.Println(jobName)
	var cur = jobName[randee]

	//fmt.Println(randee)
	return cur
}

// func _empty() this func will  make the situations randomly
// this func will generate all possible variant of actions
func possibilities(r int, a int, s string, j string) {
	if r < 10 && r >= 0 {
		fmt.Printf("%s is r %d y.o %s of UK", s, a, j)

	} else if r < 20 && r >= 10 {

		fmt.Printf("%s is r %d y.o %s with 1b subs in YouTube", s, a, j)

	} else if r < 30 && r >= 20 {
		fmt.Printf("%s is r %d y.o %s who secretly loves his stepmom", s, a, j)

	} else if r < 40 && r >= 30 {
		fmt.Printf("%s is r %d y.o %s of an average teenager", s, a, j)

	} else if r < 50 && r >= 40 {
		fmt.Printf("%s is r %d y.o second-time-merried %s ", s, a, j)

	} else {
		fmt.Printf("%s is r %d y.o %s famous all over the world", s, a, j)

	}
}

/*
	func _empty() this func will apply different random numbers for every part of code
	ex. now if random is 49 all content will use 49 index

*/
// func _empty() this func will show some extra data
// Here main code where we will apply some actions in the story game
func main() {
	age, name := namer()
	randee := randomizer()
	job := jobs()
	possibilities(randee, age, name, job)
}

// this func will make a random number generator
func randomizer() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	randee := r1.Intn(100)
	return randee
}

package main

import "fmt"

func main() {
	printWelcome()
	printGreeting(getName())
	fmt.Println("Let's go on an adventure!")
	travel()
}

func printWelcome() {
	fmt.Println("Welcome to the Solar System!")
	fmt.Println("There are 8 planets to explore.")
}

func getName() string{
	var name string
	fmt.Println("What is your name?")
    fmt.Scan(&name)
    return name
}

func printGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

func travel() {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = handleUserResponse("Shall I randomly choose a planet for you to visit? (Y or N)")
		if choice == "Y" {
			handlePlanetChoice("Traveling to Jupiter...","","Arrived at Jupiter. The large red spot appears ominous.")
		} else if choice == "N" {
			planetName := handleUserResponse("Name the planet you would like to vist!")
			handlePlanetChoice("Traveling to", planetName,"Arrived at Neptune. A very cold planet, furthest from the sun.")
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
}

func handleUserResponse(prompt string) string {
	var response string
	fmt.Println(prompt)
	fmt.Scan(&response)
	return response
}

func handlePlanetChoice(response string, planetName string, planet string) string {
	fmt.Println(response, planetName)
	fmt.Println(planet)
	return response
}
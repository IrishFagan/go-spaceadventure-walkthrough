package main

import "fmt"

func main() {
	printIntroduction("Welcome to the Solar System!\n","There are 8 planets to explore.")
	fmt.Println("Let's go on an adventure!")
	travel()
}

func printIntroduction(greeting string, intro string) {
	fmt.Println(greeting,intro)
	name := responseToPrompt("What is your name?")
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

func travel() {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = responseToPrompt("Shall I randomly choose a planet for you to visit? (Y or N)")
		if choice == "Y" {
			handlePlanetChoice("Traveling to Jupiter...","","Arrived at Jupiter. The large red spot appears ominous.")
		} else if choice == "N" {
			planetName := responseToPrompt("Name the planet you would like to vist!")
			handlePlanetChoice("Traveling to", planetName,"Arrived at Neptune. A very cold planet, furthest from the sun.")
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
}

func responseToPrompt(prompt string) string {
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
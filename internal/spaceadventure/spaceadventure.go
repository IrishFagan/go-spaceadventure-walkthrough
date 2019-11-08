package spaceadventure

import "fmt"

func Start() {
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
	if (randomDestination) {
		handlePlanetChoice()
	} else {
		travel
	}
}

func promptRandomOrSpecific(prompt string) bool {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = responseToPrompt(prompt)
		if choice == "Y" {	
			return true
		} else if choice == "N" {
			return false
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
	return false
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
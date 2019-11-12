package spaceadventure

import "fmt"

func Start(planetarySystem PlanetarySystem) {
	printIntroduction("Welcome to the",planetarySystem,"There are 8 planets to explore.")
	fmt.Println("Let's go on an adventure!")
	fmt.Println("Would you like me to show you a random planet?")
	travel()
}

func printIntroduction(greeting string, planetarySystem PlanetarySystem, intro string,) {
	fmt.Println(greeting, planetarySystem.Name,"\n")
	fmt.Println(intro)
	name := responseToPrompt("What is your name?")
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

func responseToPrompt(prompt string) string {
	var response string
	fmt.Println(prompt)
	fmt.Scan(&response)
	return response
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

func travel() {
	if (promptRandomOrSpecific("") == true) {
		printRandomPlanet()
	} else {
		var planetName string
		fmt.Println("What planet would you like to visit?")
		fmt.Scan(&planetName)
		printSpecificPlanet(planetName)
	}
}

func printRandomPlanet() {
	fmt.Println("Travleing to Jupiter!")
	fmt.Println("Arrived at Jupiter. The large red spot seems ominous")
}

func printSpecificPlanet(planetName string) {
	fmt.Println("Traveling to ", planetName)
	fmt.Println("Arrived at Neptune. A very cold planet, furthest from the sun.")
}
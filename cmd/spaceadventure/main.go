package main

import "github.com/IrishFagan/go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	spaceadventure.Start(
		spaceadventure.PlanetarySystem{
			Name: "Solar System", Planets: []spaceadventure.Planet{
				spaceadventure.Planet{"Earth", "Nice planet"}, 
			},
		},
	)
}
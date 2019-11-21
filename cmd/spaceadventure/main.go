package main

import "github.com/IrishFagan/go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	spaceadventure.Start(
		spaceadventure.PlanetarySystem{
			Name: "Solar System", Planets: mockPlanets(),
		},
	)
}

func mockPlanets() []spaceadventure.Planet {
	return []spaceadventure.Planet{
		spaceadventure.Planet{"Geonosis","Rocky n' stuff"},
		spaceadventure.Planet{"Kamnio","Clones n stuff r here"},
		spaceadventure.Planet{"Utapau","Sinkholes n stufffff and deserts"},
		spaceadventure.Planet{"Scarif","Blown up, dont visit"},
		spaceadventure.Planet{"nur","OCEAN PLANET W/ VOLCANOES"},
		spaceadventure.Planet{"Dog","Woof"},
		spaceadventure.Planet{"Kessel","spice"},
		spaceadventure.Planet{"Dathomire","Wet spice + witches"},
		spaceadventure.Planet{"Coruscant","tech n stuff"},
		spaceadventure.Planet{"Austin world","trash"},
	}
}
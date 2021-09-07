package facts

import "math/rand"

type Fact struct {
	Title  string
	Detail string
}

var facts = [...]Fact{
	{
		Title:  "Octopuses have blue blood",
		Detail: "To survive in the deep ocean, octopuses evolved a copper rather than iron-based blood called hemocyanin, which turns its blood blue. This copper base is more efficient at transporting oxygen than hemoglobin when water temperature is very low and not much oxygen is around. But this system also causes them to be extremely sensitive to changes in acidity. If the surrounding water’s pH dips too low, octopuses can’t circulate enough oxygen. Accordingly, researchers worry about what will happen to the animals as a result of climate change-induced ocean acidification.",
	},
	{
		Title:  "Octopuses have three hearts",
		Detail: "Two of the hearts work exclusively to move blood beyond the animal’s gills, while the third keeps circulation flowing for the organs. The organ heart actually stops beating when the octopus swims, explaining the species’ penchant for crawling rather than swimming, which exhausts them.",
	},
	{
		Title:  "Octopus arms have a mind of their own",
		Detail: "Two-thirds of an octopus’ neurons reside in its arms, not its head. As a result, the arms can problem solve how to open a shellfish while their owners are busy doing something else, like checking out a cave for more edible goodies. The arms can even react after they’ve been completely severed. In one experiment, severed arms jerked away in pain when researchers pinched them.",
	}}

func RandomFact() Fact {
	return facts[rand.Intn(len(facts))]
}

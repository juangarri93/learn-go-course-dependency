package puppy

func Bark() string {
	return "Woof!"
}

func SuperBark() string {
	return "Woooooooof!"
}

type Puppy struct {
	name string
	breed string 
	age int
}
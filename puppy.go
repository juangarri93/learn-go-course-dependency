package puppy

func Bark() string {
	return "Woof!"
}

func SuperBark() string {
	return "Woooooooof!"
}

type Puppy struct {
	Name string
	Breed string 
	Age int
}
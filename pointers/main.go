package main

import "fmt"

type Musician struct {
	Name *string
	Age  int
}

func main() {
	var name = "Cartola"

	musician1 := Musician{
		Name: &name,
		Age:  80,
	}

	musician2 := musician1
	fmt.Println("Musician 1 Name:", *musician1.Name)
	fmt.Println("Musician 2 Name:", *musician2.Name)
	fmt.Println("---------------------")

	musician2.Name = toPointer("João Nogueira")
	fmt.Println("Musician 1 Name:", *musician1.Name)
	fmt.Println("Musician 2 Name:", *musician2.Name)

	fmt.Println("---------------------")

	musician3 := Musician{
		Name: toPointer("Beth Carvalho"),
		Age:  60,
	}
	musician4 := deepCopy(musician3)
	musician4.Name = toPointer("Alcione")
	fmt.Println("Musician 3 Name:", *musician3.Name)
	fmt.Println("Musician 4 Name:", *musician4.Name)

}

func toPointer(s string) *string {
	return &s
}

func deepCopy(source Musician) Musician {
	var destination Musician
	destination.Age = source.Age
	destination.Name = toPointer(*source.Name)
	return destination
}

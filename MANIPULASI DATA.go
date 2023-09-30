package main

import "fmt"

type Person struct {
	Nama string
	Umur  int
}

func updatePerson(p *Person, nama string, Umur int) {
	p.Nama = nama
	p.Umur = Umur
}

func main() {
	person := Person{
		Nama: "John",
		Umur:  30,
	}

	fmt.Println("Data awal:")
	fmt.Println("Nama:", person.Nama)
	fmt.Println("Umur:", person.Umur)

	updatePerson(&person, "Alice", 25)

	fmt.Println("\nData setelah diubah:")
	fmt.Println("Nama:", person.Nama)
	fmt.Println("Umur:", person.Umur)
}

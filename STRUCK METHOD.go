package main

import "fmt"

type Person struct {
    Nama       string
    Jurusan    string
    AsalDaerah string
}

func (p Person) Greet() string {
    return fmt.Sprintf("Halo, nama saya %s. Saya adalah mahasiswa jurusan %s dari %s.", p.Nama, p.Jurusan, p.AsalDaerah)
}

func main() {
    mahasiswa1 := Person{
        Nama:       "Andi",
        Jurusan:    "Informatika",
        AsalDaerah: "Jakarta",
    }

    mahasiswa2 := Person{
        Nama:       "Budi",
        Jurusan:    "Sistem Informasi",
        AsalDaerah: "Bandung",
    }

    fmt.Println(mahasiswa1.Greet()) // Output: "Halo, nama saya Andi. Saya adalah mahasiswa jurusan Informatika dari Jakarta."
    fmt.Println(mahasiswa2.Greet()) // Output: "Halo, nama saya Budi. Saya adalah mahasiswa jurusan Sistem Informasi dari Bandung."
}

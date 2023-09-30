package main

import "fmt"

func main() {
    mahasiswa := make(map[int]map[string]string)

    mahasiswa[1] = map[string]string{
        "nama":   "Andi",
        "jurusan": "Informatika",
    }
    mahasiswa[2] = map[string]string{
        "nama":   "Budi",
        "jurusan": "Sistem Informasi",
    }
    mahasiswa[3] = map[string]string{
        "nama":   "Citra",
        "jurusan": "Teknik Elektro",
    }

    fmt.Println("Data Mahasiswa:")
    for id, data := range mahasiswa {
        fmt.Printf("ID: %d\n", id)
        fmt.Printf("Nama: %s\n", data["nama"])
        fmt.Printf("Jurusan: %s\n", data["jurusan"])
        fmt.Println("----------")
    }
}

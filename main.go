package main

import "fmt"

func main() {
    // Membuat objek mobil dengan ban default Ban Karet
    mobil := Mobil{
        Roda: [4]Ban{&BanKaret{}, &BanKaret{}, &BanKaret{}, &BanKaret{}},
        PintuKanan: Pintu{Nama: "Kanan", KetukBunyi: "Knock", BukaBunyi: "beep"},
        PintuKiri: Pintu{Nama: "Kiri", KetukBunyi: "beep", BukaBunyi: "Knock"},
    }

    // Menampilkan jenis ban saat ini
    for i, roda := range mobil.Roda {
        fmt.Printf("Roda %d: %s\n", i+1, roda.JenisBan())
    }

    // Mengganti roda ke-2 dan ke-4 dengan Ban Besi dan Ban Kayu
    mobil.GantiRoda(1, &BanBesi{})
    mobil.GantiRoda(3, &BanKayu{})

    // Menampilkan jenis ban setelah diganti
    fmt.Println("\nSetelah penggantian roda:")
    for i, roda := range mobil.Roda {
        fmt.Printf("Roda %d: %s\n", i+1, roda.JenisBan())
    }

    // Ketuk dan buka kedua pintu
    fmt.Println("\nPintu Kanan:")
    mobil.PintuKanan.Ketuk()
    mobil.PintuKanan.Buka()

    fmt.Println("\nPintu Kiri:")
    mobil.PintuKiri.Ketuk()
    mobil.PintuKiri.Buka()
}




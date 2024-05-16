package main

import "fmt"

type Pintu struct {
    Nama string
    KetukBunyi string
    BukaBunyi string
}

// Method ketuk pintu
func (p Pintu) Ketuk() {
    fmt.Println(p.KetukBunyi)
}

// Method buka pintu
func (p Pintu) Buka() {
    fmt.Println(p.BukaBunyi)
}
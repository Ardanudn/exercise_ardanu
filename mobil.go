package main

import "fmt"

type Mobil struct {
	Roda       [4]Ban
	PintuKanan Pintu
	PintuKiri  Pintu
}

func (m *Mobil) GantiRoda(index int, roda Ban) {
	if index >= 0 && index < 4 {
		m.Roda[index] = roda
	} else {
		fmt.Println("jumlah roda salah")
	}
}

type Ban interface {
    JenisBan() string
}
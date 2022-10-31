package main

import "testing"

const (
	sisi           float64 = 10
	luasSeharusnya float64 = 100
)

func TestHitungLuas(t *testing.T) {
	bentuk := persegi{sisi}
	if bentuk.luas() != luasSeharusnya {
		//t.Fail()
		t.Errorf("SALAH! harusnya %2f", luasSeharusnya)
	}
}

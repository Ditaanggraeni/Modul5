package main

import "fmt"

func main() {
	var menu, brs int
	fmt.Println("=========================")
	fmt.Println("Sub Menu : ")
	fmt.Println("1. Segitiga")
	fmt.Println("2. Segitiga Terbalik")
	fmt.Println("=========================")
	fmt.Print("Pilih menu angka diatas : ")
	fmt.Scan(&menu)
	fmt.Println("=========================")

	if menu == 1 {
		fmt.Print("Masukkan Jumlah Baris :  ")
		fmt.Scan(&brs)
		for i := 1; i <= brs; i++ {
			for j := brs; j >= i; j-- {
				fmt.Print(" ")
			}
			for k := 1; k <= i; k++ {
				fmt.Printf("*")
			}
			for k := (i - 1); k >= 1; k-- {
				fmt.Printf("*")
			}
			fmt.Printf("\n")
		}
	} else {
		fmt.Print("Masukkan Jumlah Baris :  ")
		fmt.Scan(&brs)
		for i := brs; i >= 1; i-- {
			for j := brs; j >= i; j-- {
				fmt.Print(" ")
			}
			for k := 1; k <= i; k++ {
				fmt.Printf("*")
			}
			for k := (i - 1); k >= 1; k-- {
				fmt.Printf("*")
			}
			fmt.Printf("\n")
		}
	}
}

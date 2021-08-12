package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ulang := 100
	pemain := 3
	nilai := [3]int{0,0,0}
	lemparan := [3]int{4,4,4}
	rand.Seed(time.Now().UnixNano())

	fmt.Println(nilai)
	fmt.Println(nilai[0])

	dadu_acak := make([][]int , pemain)
	fmt.Println("===================================")
	for i := 0; i < ulang; i++{
		fmt.Println(len(dadu_acak))
		fmt.Println("Giliran", i+1, "lempar dadu :")
		for i := 0; i < pemain; i++ {
			dadu_acak[i] = make([]int, lemparan[i])
				
			for j := 0; j < lemparan[i]; j++{
				random := rand.Intn(6) + 1
				dadu_acak[i][j] = random
			}
		}
		for i := 0; i < pemain; i++ {
			fmt.Println("Pemain", i+1, "Nilai:", nilai[i], dadu_acak[i])
		}

		fmt.Println("\nSetelah revisi :")
		for i:=0; i<len(dadu_acak); i++ {
			for j := 0; j < len(dadu_acak[i]); j++{
				if dadu_acak[i][j] == 1 {
					dadu_acak[i] = remove(dadu_acak[i], j)
					j=0
					if len(dadu_acak[i]) == 0 {
						continue
					}
				}
				if dadu_acak[i][j] == 6 {
					dadu_acak[i] = remove(dadu_acak[i], j)
					nilai[i] += 1 
					j=0
				}
				lemparan[i] = len(dadu_acak[i])
			}
		}
		for i := 0; i < pemain; i++ {
			fmt.Println("Pemain", i+1, "Nilai:", nilai[i], dadu_acak[i])
		}

		// Break Winner
		if len(dadu_acak[0]) == 0 || len(dadu_acak[1]) == 0 {
			if len(dadu_acak[0]) == 0 && len(dadu_acak[1]) == 0 {
				break
			}else if len(dadu_acak[2]) == 0 {
				break
			}
		}

		fmt.Println("===================================")
	}

	var largerNumber, temp, person int  

	for i, element := range nilai {  
		if element > temp {  
			person = i
		 temp = element
		 largerNumber = temp  
		}  
	 } 
		
	fmt.Println("Pemenangnya adalah pemain", person+1, "dengan nilai :", largerNumber)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
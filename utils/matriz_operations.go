package utils

import "fmt"

func Rellenar_matriz(array *[3][3]int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Printf("Ingresa el valor #%d de la fila #%d\n", j+1, i+1)
			fmt.Scan(&array[i][j])
		}
	}
}

func Calcular_suma_diagonal(array *[3][3]int) int {

	var result int

	for i := 0; i < len(array); i++ {
		for j := i; j < len(array[i]); j++ {
			if i == j {
				result += array[i][j]
			}
		}
	}
	return result
}

func Mostrar_matriz(array *[3][3]int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Printf(" %d ", array[i][j])
		}
		fmt.Print("\n")
	}
}

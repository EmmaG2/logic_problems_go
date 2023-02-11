package problems

import (
	"fmt"

	"github.com/logic_problems_go/utils"
)

func Initialize_matriz_diagonal_sum() {
	var matrix [3][3]int

	utils.Rellenar_matriz(&matrix)

	fmt.Println()

	utils.Mostrar_matriz(&matrix)

	fmt.Println()

	result := utils.Calcular_suma_diagonal(&matrix)

	fmt.Printf("Resultado de la suma de la diagonal: %d\n", result)
}

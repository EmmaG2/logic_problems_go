package problems

import (
	"fmt"

	"github.com/logic_problems_go/problems/matriz/logic"
)

func Initialize_matriz_diagonal_sum() {
	var matrix [3][3]int

	logic.Rellenar_matriz(&matrix)

	fmt.Println()

	logic.Mostrar_matriz(&matrix)

	fmt.Println()

	result := logic.Calcular_suma_diagonal(&matrix)

	fmt.Printf("Resultado de la suma de la diagonal: %d\n", result)
}

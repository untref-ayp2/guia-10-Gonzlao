package ejercicios

// Problema de los dados
// Se tienen n dados de k caras cada uno, se desea saber la cantidad de formas de obtener
// una suma de x puntos al lanzar los n dados.
// Ejemplo: n = 3, k = 6, x = 7
// 1 1 5
// 1 2 4
// 1 3 3
// 1 4 2
// 1 5 1
// 2 1 4
// 2 2 3
// 2 3 2
// 2 4 1
// 3 1 3
// 3 2 2
// 3 3 1
// 4 1 2
// 4 2 1
// 5 1 1
// Total: 15 formas
// Devuelve la cantidad de formas de obtener x puntos al lanzar n dados de k caras
func Dados(n, k, x int) int {
	cantidad := 0
	dado := make([]int, n)

	backtracking(dado, 0, 0, x, k, &cantidad)

	return cantidad
}

func backtracking(dado []int, index, currentSum, x, k int, cantidad *int) {
	if index == len(dado) {
		if currentSum == x {
			*cantidad++
		}
		return
	}

	for i := 1; i <= k; i++ {
		dado[index] = i

		backtracking(dado, index+1, currentSum+i, x, k, cantidad)
	}
}

// Demonstra slices: a visão dinâmica sobre sequências em Go.
package main

import "fmt"

func main() {
	// Literal de slice
	nums := []int{10, 20, 30}
	fmt.Println("literal:", nums, "len:", len(nums), "cap:", cap(nums))

	// make: tamanho (e opcionalmente capacidade)
	zeros := make([]int, 3)
	fmt.Println("make:", zeros)

	// append pode realocar se a capacidade estourar
	nums = append(nums, 40, 50)
	fmt.Println("após append:", nums, "len:", len(nums), "cap:", cap(nums))

	// Fatiamento: [início:fim) - fim não incluso
	meio := nums[1:4]
	fmt.Println("fatiamento [1:4]:", meio)

	// Array tem tamanho fixo no tipo; slice é o dia a dia
	var arr [3]string = [3]string{"a", "b", "c"}
	parte := arr[:] // slice que cobre o array inteiro
	fmt.Println("array:", arr)
	fmt.Println("slice do array:", parte)

	// Copiar para não compartilhar o backing array sem querer
	copia := make([]int, len(nums))
	copy(copia, nums)
	copia[0] = 999
	fmt.Println("original:", nums)
	fmt.Println("cópia alterada:", copia)
}

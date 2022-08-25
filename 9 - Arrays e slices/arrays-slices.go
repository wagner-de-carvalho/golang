package main

import(
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e slices")
	// ARRAYS
	var array1[5] int
	fmt.Println(array1)
	
	var array2[3] string
	array2[0] = "Golang"
	array2[1] = "Golang2"
	array2[2] = "Golang3"
	fmt.Println(array2)
	
	array3 := [3] string{"Go 1.1", "Go 2.0", "Go 3.0"}
	fmt.Println(array3)
	
	array4 := [...] int{1, 2, 5, 8, 13, 14, 15, 19}
	fmt.Println(array4)
	
	// SLICES
	slice := []int{10, 40, 21, 30, 70}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 22)
	fmt.Println(slice)
	
	slice2 := array4[1:3]
	fmt.Println(slice2)

	// ARRAYS INTERNOS
	fmt.Println("Arrays INTERNOS")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //TAMANHO
	fmt.Println(cap(slice3)) //CAPACIDADE

}

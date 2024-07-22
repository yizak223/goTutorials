package main

import "fmt"

func main() {
	var intNun int = 32767
	intNun = intNun + 1
	fmt.Println(intNun)
	//how to store correctlly numbers
	var floatNun1 float32 = 12345678.9
	fmt.Println(floatNun1) // 123456789.00

	var floatNun2 float64 = 12345678.9
	fmt.Println(floatNun2) // 12345678.900

	var floatNun32 float32 =10.1
	var intNun32 int32 =2 
	//var result float32 = floatNun32 + intNun32// this cant be done
	var result float32 = floatNun32 + float32(intNun32)// this can be done
	fmt.Println(result) 

	var intNum1 int =3
	var intNum2 int =2
	fmt.Println(intNum1/intNum2) //no remainder result
	fmt.Println(intNum1%intNum2) //remainder result


}

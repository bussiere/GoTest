package main

import (
	"fmt"
)



func sortArray(arrayInt *[]int){
    var index2 int
	index2 = 0
	for _, arrayE := range *arrayInt{
		fmt.Println("arrayE")
		fmt.Println(arrayE)
		if (arrayE > 5||  arrayE < 1) {
			fmt.Println("arrayE inside")
			fmt.Println(arrayE)
			// if not in delimiting time we remove the objects
			*arrayInt = append((*arrayInt)[:index2], (*arrayInt)[index2+1:]...)
		} else {
			fmt.Println("index2")
			fmt.Println(index2)
			// we keep the good object
			index2 = index2 + 1
		}
	}

}

func sortType(toto *TestArray){
	var index2 int
	index2 = 0
	for _, arrayE := range toto.IntArray{
		fmt.Println("arrayE")
		fmt.Println(arrayE)
		if (arrayE > 5 ||  arrayE < 1) {
			fmt.Println("arrayE inside")
			fmt.Println(arrayE)
			// if not in delimiting time we remove the objects
			toto.IntArray = append(toto.IntArray[:index2], toto.IntArray[index2+1:]...)
		} else {
			fmt.Println("index2")
			fmt.Println(index2)
			// we keep the good object
			index2 = index2 + 1
		}
	}

}

type TestArray struct {
	IntArray []int         `json:"IntArray"`
}

func main(){
	var arrayInt []int = []int{1,2,3,4,5,6}
	fmt.Println(arrayInt)
	sortArray(&arrayInt)
	fmt.Println(arrayInt)
	var toto TestArray
	toto.IntArray  = arrayInt
	fmt.Println(toto)
	sortType(&toto)
	fmt.Println(toto)


}

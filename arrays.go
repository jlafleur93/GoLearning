package main

import "fmt"

func main() {
	var arr [10]int
	arr[0] = 42
	arr[1] = 13

	//	fmt.Printf("The first element is %d\n", arr[0])
	//	fmt.Printf("The second element is %d\n", arr[1])
	shortAssign()
	twodArray()
	sliceArray()

}

func shortAssign() {
	a := [4]int{1, 2, 3, 4}
	b := [10]int{1, 2, 3, 4}
	c := [...]int{4, 5, 6}

	fmt.Printf("The first element list is %d\n", a)
	fmt.Printf("The first element list is %d\n", b)
	fmt.Printf("The first element list is %d\n", c)

}
func twodArray() {
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{4, 2, 1, 7}}
	easyArray := [2][4]int{[4]int{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Printf("Double Array List %d\n", doubleArray)
	fmt.Printf("Easy Array List %d\n", easyArray)
}
func sliceArray() {
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var a, b []byte
	a = ar[2:5]
	b = ar[3:5]
	fmt.Printf("Slice array value: %d \n", a)
	fmt.Printf("Slice array value: %d \n", b)
}

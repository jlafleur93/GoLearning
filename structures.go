package main

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	specialty string
}

func main() {
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	println("His name is", mark.name)
	println("His age is", mark.age)
	println("His weight is", mark.weight)
	println("His specialty is", mark.specialty)

	jim := Student{Human{"jim:", 25, 120}, "Quick Maffs"}
	println("His name is", jim.name)
	println("His age is", jim.age)
	println("His weight is", jim.weight)
	println("His specialty is", jim.specialty)

}

package main

func main() {
	//define var with var and then type afterwards
	var stringList string
	//You can definite Multiple Variables, and definie the type later
	var coolList, lameList, sickList string
	//You can definite a var with init value
	var tryVar string = value

	// ---- Defining Multiple typed vars ----
	//var vname1, vname2, vname3 type = v1, v2, v3
	//You can short hand this by removing the type
	// Like below
	//var vname1, vname2, vname3 = v1, v2, v3

	//Even shorter example
	vname1, vname2, vname3 := v1, v2, v3
	/*
		Now it looks much better. Use := to replace var and type, this is called a short assignment. It has one limitation: this form can only be used inside of a functions. You will get compile errors if you try to use it outside of function bodies. Therefore, we usually use var to define global variables.

	*/
}

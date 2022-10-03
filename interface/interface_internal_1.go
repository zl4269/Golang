package main

func printNilInterface() {
	var i interface{}
	var err error
	println(i)
	println(err)
	println("i=nil:", i == nil)
	println("err=nil:", err == nil)
	println("i=err:", i == err)
	println("")
}

func main() {
	printNilInterface()
}

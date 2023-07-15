package main

import "go-exercises/exportedUnexportedPackage"

func main() {

	demo.Print(demo.A)

	//demo.print() unexported fonksiyon olduğu için ulaşılamaz
	//demo.B unexported değişken olduğu için ulaşılamaz.
}

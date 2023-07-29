package main

//Birden fazla paketi import etmek için aşağıdaki syntax kullanılabilir.
import (
	"fmt"
	"reflect"
)

func main() {

	var a int
	a = 10
	//a = "abc"
	fmt.Println(a)

	//Aşağıdaki gibi çoklu tanımlama yapabiliriz.
	var b, c, d float32
	b = 1.0
	c = 2.1
	fmt.Println(b, c, d)

	var e = true
	e = false
	fmt.Println(e)

	//Aşağıdaki gibi çoklu tanımlama yapıp değer ataması yapabiliriz.
	var g, h, j = 10, "g", true
	fmt.Println(g, h, j)

	//Aşağıdaki gibi tip belirtmeden atama yapabiliriz.
	//Fakat bu atamadan sonra depişken tipi belirlenir ve başka tipte veri atanamaz
	f := 10.12
	f += 2
	//f = "abcde"
	fmt.Println(f, reflect.TypeOf(f))

}

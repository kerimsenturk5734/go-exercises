package demo

import "fmt"

// Büyük harfle başlayan değişken veya fonksiyon isimleri o değişkeni Exported(Public) yapar.
var A = "Exported"

func Print(a ...string) {
	printS(a...)
}

// Küçük harfle başlayan değişken veya fonksiyon isimleri o değişkeni Unexported(private) yapar.
var b = "unexported"

func printS(a ...string) {
	fmt.Println(a, b)
}

package main

import "fmt"

func main() {
	// Çalıştırılabilir bir .go dosyası oluşturmak için aşağıdaki şartların karşılanması gerekir.
	/*
		- Dosyanın başında "package main" olarak belirtilmelidir.
		- "package main" yerine "package blabla" ifadesinin kullanılması bu dosayının çalışmasını engeller.
		- İçerisinde "main()" isimli bir fonksiyon bulunması gerekmektedir.
		- Dosya adının önemi yoktur.
		- Bu dosyayı çalıştırmak için "go run filename.go" terminal komutu kullanılır.
		- Yukardaki şartların sağlanmaması durumunda compile hatası alınacaktır.
	*/
	fmt.Println("Hello World")
}

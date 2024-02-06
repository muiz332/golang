/*

-  interface

dimateri kali ini kita akan bahas mengenai interface

apa itu interface?
interface adalah sebuah type data yang dapat mendefinisikan type data sebuah method dan dia tidak
memiliki implementasi secara langsung

lalu setelah kita membuat interface, bagaimana caranya kalo kita ingin
mengimplementasikan interfacenya?

dari pengertian diatas terdapat kata method, dan method itu berhubungan dengan struct
jadi kalo mau mengimplementasikan interface, kita harus membuat struct yang sama
dengan definisi method yang ada didalam intefacenya

kalo didalam bahasa pemrograman yang lain
interface itu seperti ini


misalkan diphp

interface Animal {
  public function makeSound();
}

class Cat implements Animal {
  public function makeSound() {
    echo "Meow";
  }
}


jadi kita punya class yang dimana kita mengikuti type data dari interface
Animal

jadi class Cat itu harus berisi method makeSound dan kita tuliskan keyword implements Animal
jadi kalo diphp harus kita tuliskan dulu kalo ingin mengikuti type interface Animal

sedangkan digolang tidak
digolang untuk implementasi interface itu terjadi secara otomatis

jadi kalo misalkan kita punya struct, dimana struct itu memiliki method yang
sama dengan definisi interface yang ada

maka struct tersebut secara otomatis langsung mengimplementasikan interface
tersebut

oke langsung saja kita coba




*/

package main

import (
	"fmt"
)

type Animal interface {
	MakeSound() string
}

type Cat struct {
	Name string
}

func (cat Cat) MakeSound() string {
	return "miaw"
}

func (cat Cat) Intro() {
	fmt.Println("hello i have a cat, name is", cat.Name)
}

type Dog struct {
	Name string
}

func (dog Dog) MakeSound() string {
	return "guk guk"
}

func main() {

	var cat Animal = Cat{
		Name: "ketty",
	}

	fmt.Println(cat.MakeSound())
	cat.(Cat).Intro()

	/*

		jadi interface itu type data, kalo type data kita bisa gunakan untuk mendeklarasikan
		type data dari sebuah variable ataupun parameter

		jadi disini bacanya
		variable cat harus memiliki type Animal

		dimana Animal ini adalah sebuah interface yang didalamnya terdapat definisi method
		secara otomatis struct Cat itu akan mengimplementasikan interface Animal

		tidak seperti bahasa pemrograman yang lain harus dituliskan keyword implements
		setelah itu kita assigment dengan struct Cat

		kalau misalkan Catnya ini tidak memiliki method MakeSound yang mereturn string
		maka interfacenya tidak akan diimplementasikan ke structnya

		akan tetapi kalo kelian memiliki method yang tidak ada didalam interfacenya
		kalo kalian panggil itu akan error

		misalkan kalo saya tuliskan gini
		cat.Intro()

		itu akan terjadi error, kerena tidak ada didalam deklarasi interface Animalnya
		agar bisa dipanggil kita bisa tulis seperti ini

		cat.(Cat).Intro() dan ini dinamakan dengan type assertion yang nanti akan kita bahas

		nah sekarang kita coba
		kalo kita buat struct lagi dimana struct tersebut memiliki
		method MakeSound, apakah struct tersebut mengimplementasikan interface Animal ?

	*/

	var dog Animal = Dog{
		Name: "apa",
	}
	fmt.Println(dog.MakeSound())

	/*

		kalo kita jalankan tidak error berarti, 1 interface bisa diimplementasikan kebanyak
		struct

		dan satu lagi, kita bisa membuat interface embedded, maksutnya kita bisa buat interface
		seperti ini

		type Animal1 {
			MakeSound() string
		}

		type Animal2{
			run()
		}

		type Animal{
			Animal1
			Animal2
		}

		jadi kita bisa penggil interface yang lain didalam interface yang kita buat
		oke jadi seperti itu ya

		jadi interface itu digunakan untuk mendeklarasikan method method
		kalo struct itu memiliki method yang sama dengan interfacenya

		maka secara otomatis struct tersebut mengimplementasikan interface tersebut
		oke jadi seperti itu mudah mudahan kalian paham

	*/

}

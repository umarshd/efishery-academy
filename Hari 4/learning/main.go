package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

// Struct
type student struct {
	grade int
	person
}

func main() {

	// Mencetak hello world
	fmt.Println("Hello World")

	// Numeric Decimal
	var decimalNUmber = 3.45

	fmt.Printf("bilangan decimal: %f\n", decimalNUmber)
	fmt.Printf("bilangan decimal: %.3f\n", decimalNUmber)

	// Boolean
	var exists bool = true
	fmt.Printf("exists? %t\n", exists)

	// String
	var message = `Nama Saya "Umar Sahid".
	Salam kenal.
	Mari Belajar "Golang".`

	fmt.Println(message)

	// Variable
	// Variable Declaration
	var firstName string = "Umar"
	var lastName string
	lastName = "Sahid"

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	// Declaration Variables without type data
	var firstNameAgain string = "Umar"
	lastNameAgain := "Sahid"

	fmt.Printf("halo %s %s!\n", firstNameAgain, lastNameAgain)

	// Declaration Multi Variables
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	fmt.Println(first, second, third, seventh, eight, ninth)

	// Declaration Underscore Varibles
	_ = "Belajar Golang"
	_ = "Golagn itu mudah"
	myName, _ := "Umar", "Bebas"

	fmt.Print("Hello ", myName, "!\n")

	// Constants
	const myFirstName string = "Umar"
	fmt.Print("Halo ", myFirstName, "!\n")

	//Condition
	// if .. else .. then
	var point = 8
	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point >= 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	// Switch Case

	point = 6

	switch point {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Looping
	// For Range

	var fruits = [4]string{"Apple", "Grape", "Banana", "Melon"}

	for i, fruit := range fruits {
		fmt.Printf("Elemen %d : %s \n", i, fruit)
	}

	// For Loop
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// For Loop Break Continue

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	// Functions

	// Use function swap
	a, b := swap("Hello", "Dunia")
	fmt.Println(a, b)

	// Struct

	var s1 = student{}
	s1.name = "Umar"
	s1.age = 25
	s1.grade = 3

	fmt.Println("name : ", s1.name)
	fmt.Println("age : ", s1.age)
	fmt.Println("age : ", s1.person.age)
	fmt.Println("grade : ", s1.grade)

	// Struct With Slice
	var allStudent = []person{
		{name: "Bamabang", age: 24},
		{name: "Junedi", age: 40},
		{name: "Wijaya", age: 21},
		{name: "Indah", age: 22},
	}

	for _, student := range allStudent {
		fmt.Println(student.name, "age is", student.age)
	}

	// Structure - Map
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	// Structure - Array
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	// Slice
	var fruitsSlice = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruitsSlice[0]) // "apple"

	var aFruits = fruitsSlice[0:3]
	var bFruits = fruitsSlice[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruitsSlice) // [apple grape banana melon]
	fmt.Println(aFruits)     // [apple grape banana]
	fmt.Println(bFruits)     // [apple banana melon]
	fmt.Println(aaFruits)    // [grape]
	fmt.Println(baFruits)    // [grape]

	// Buah "grape" di ubah menjadi pinnaple
	baFruits[0] = "pinnaple"

	fmt.Println(fruitsSlice) // [apple pinnaple banana melon]
	fmt.Println(aFruits)     // [apple pinnaple banana]
	fmt.Println(bFruits)     // [apple banana melon]
	fmt.Println(aaFruits)    // [pinnaple]
	fmt.Println(baFruits)    // [pinnaple]

	// Defer
	defer fmt.Println("Halo")
	fmt.Println("selamat datang")

	// Pointer
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

}

// Function Swap
// Multiple Return Values
func swap(x, y string) (string, string) {
	return y, x
}

// Consecutive named function parameters
// From
// func add(x int, y int) int {
// 	return x + y
// }

// to
func add(x, y int) int {
	return x + y
}

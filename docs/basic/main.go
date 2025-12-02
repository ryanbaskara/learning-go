package main

import (
	"fmt"
	"sort"
	"strings"
)

// =======================
//
//	Fungsi Utama
//
// =======================
func main() {
	// ===== Variabel =====
	var name string = "Ryan"
	age := 25 // short declaration
	const PI = 3.14
	fmt.Println(name, age, PI)

	// ===== Input/Output =====
	// fmt.Scan(&name)

	// ===== Tipe Data =====
	var x int = 10
	var y float64 = 2.5
	var ok bool = true
	fmt.Printf("x=%d, y=%.2f, ok=%t\n", x, y, ok)

	// ===== Operator =====
	sum := x + 5
	fmt.Println("sum:", sum)

	// ===== If Else =====
	if age >= 18 {
		fmt.Println("Dewasa")
	} else {
		fmt.Println("Anak-anak")
	}

	// ===== Switch =====
	day := 3
	switch day {
	case 1:
		fmt.Println("Senin")
	case 2, 3:
		fmt.Println("Selasa/Rabu")
	default:
		fmt.Println("Hari lain")
	}

	// ===== Looping =====
	for i := 0; i < 5; i++ {
		fmt.Println("For i:", i)
	}
	i := 0
	for i < 3 {
		fmt.Println("While-like:", i)
		i++
	}
	// for range
	for idx, val := range []string{"a", "b", "c"} {
		fmt.Println(idx, val)
	}

	// ===== Array =====
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)

	// ===== Slice =====
	slc := []int{3, 1, 4}
	slc = append(slc, 5)
	fmt.Println(slc, slc[1:3])

	// ===== Map =====
	m := map[string]int{"a": 1, "b": 2}
	m["c"] = 3
	fmt.Println(m, m["b"])
	delete(m, "a")

	// ===== String =====
	str := "halo dunia"
	fmt.Println(len(str))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.Contains(str, "dunia"))

	// ===== Sort =====
	nums := []int{5, 2, 9, 1}
	sort.Ints(nums)
	fmt.Println("sorted:", nums)
	words := []string{"z", "a", "c"}
	sort.Strings(words)
	fmt.Println("sorted:", words)

	// ===== Struct =====
	p := Person{Name: "Ryan", Age: 25}
	fmt.Println(p.Name, p.Age)

	// ===== Method =====
	p.sayHello()

	// ===== Pointer =====
	num := 10
	ptr := &num
	fmt.Println("ptr value:", *ptr)
	*ptr = 20
	fmt.Println("num:", num)

	// ===== Function =====
	result := tambah(5, 7)
	fmt.Println("tambah:", result)

	// ===== Multiple Return =====
	a, b := swap("foo", "bar")
	fmt.Println(a, b)

	// ===== Interface =====
	var s Speaker = p
	s.Speak()

	// ===== Error Handling =====
	div, err := bagi(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("hasil:", div)
	}

	// ===== Defer =====
	defer fmt.Println("Ini dipanggil terakhir")
	fmt.Println("Selesai main()")
}

// =======================
//  Fungsi Tambahan
// =======================

func tambah(a, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}

func bagi(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("tidak bisa bagi 0")
	}
	return a / b, nil
}

// ===== Method di struct =====
type Person struct {
	Name string
	Age  int
}

func (p Person) sayHello() {
	fmt.Println("Halo, saya", p.Name)
}

// ===== Interface =====
type Speaker interface {
	Speak()
}

func (p Person) Speak() {
	fmt.Println("Nama saya", p.Name)
}

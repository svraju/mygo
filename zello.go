package main

import "fmt"
import "stringutil"

func main() {
//var j int = 1;
	fmt.Printf("siva\n")
fmt.Printf(stringutil.Reverse("!oG ,olleH"));
fmt.Printf("\nadd :%v",stringutil.Add(1,2));
fmt.Printf("\nsub :%v",stringutil.Sub(4,2));
fmt.Printf("\nmul :%v",stringutil.Mul(4,2));
fmt.Printf("\ndiv :%v\n",stringutil.Div(4,2));

primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)


names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)


str := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(str)
	fmt.Println("completed")
 

}


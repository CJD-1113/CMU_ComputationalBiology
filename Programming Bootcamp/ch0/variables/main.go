/*
NOTES:
Every variable must have a type
You must use every variable that you declare

{var} name {type}

type - description - default
int - integer - 0
float64 - decimal - 0.0
string - word - ""
uint - unsigned integer - 0
byte - symbol - 0
bool - true or false - false
*/
package main

import "fmt"

func main() {
	fmt.Println("Variables and types.")

	var j int = 14            //0
	var x float64 = -2.3      //0.0
	var yo string = "HI"      //""
	var u uint = 14           //0
	var symbol byte = 'H'     //0
	var statement bool = true //false

	//shorthand declarations
	i := -6              // var i int = -6 automatically infers i is an int
	hi := "YO"           // var hi string = "YO"
	k := 34              // GO defaults to int
	y := 3.7             //float
	secStatement := true //bool

	//print variables
	fmt.Println(j)
	fmt.Println(x)
	fmt.Println(yo)
	fmt.Println(u)
	fmt.Println(symbol)
	fmt.Println(statement)

	fmt.Println(i, hi, k, y, secStatement)

	fmt.Println(2 * (i + 5) * k)
	fmt.Println(2*y - 3.16)

	fmt.Println(float64(k) * y) //must cast instead of infer
	fmt.Println(k / 14)         //int division gets rid of remainder
	fmt.Println(float64(k) / 14.0)

	var p int = -187
	var s uint = uint(p)
	fmt.Println(s)

	//largest possible int, when adding one go prints smallest possible int
	m := 9223372036854775807
	fmt.Println(m + 1)
}

package main

import (
	"fmt"
	"math"
	// "sort"
)

func sayGreating(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n int) {
	fmt.Printf("Hello you scored %v out of 100 \n", n)
}

func chatArea (r float64) float64{
	return math.Pi* r*r*r
}

func cycleNames (n []int, f func(int)){

	for _, v := range n  {
		
		f(v)
	}
}

func main () {

	sayGreating("Kalimu")
	sayGreating("Kalinda")
	sayBye(54)

	cycleNames([]int{1,1,2,3,5,6,6,7}, sayBye)
	area := chatArea(345.89)

	fmt.Printf("Your float number is %v", area)
	
	// Functions

	// lens && append
	

	// age := 67
	// name := "Eric"
	// points := 255.55

	// // Print

	// fmt.Print("hello, " )
	// fmt.Print("people \n")
	// fmt.Print("Line one \n")

	// // Println
	// fmt.Println("Hello Ninjas")
	// fmt.Println("good bye Ninjas")
	// fmt.Println("My age is", age, "and My name is ", name)

	// // Printf(formated string) %v or %_ = format specifier

	// fmt.Printf("My age is %v and my name is %v \n", age, name)
	// fmt.Printf("My age is %q and my name is %q \n", age, name)
	// fmt.Printf("My age type is %T and my name type is %T \n", age, name)
	// fmt.Printf("You scored %0.1f ponts \n", points)



	// // Sprintf ( save formated string)

	// var str = fmt.Sprintf("My age type is %T and my name type is %T \n", age, name)
	// fmt.Println(str)


	// var employeesAge = [3]int{20,30,40}
	// names := [4]string{"Carl", "Olen", "Keisha", "kelon"}
	// names[1] = "keria"

	// fmt.Println(employeesAge, len(employeesAge))
	// fmt.Println(names, len(names))


	// // slices(array under the hood)

	// var scores = []int{100, 200, 300,400}

	// scores[3] = 600

	// scores = append(scores, 78)

	// fmt.Println(scores, len(scores))

	// // Slice ranges

	// rangeOne := names[1: 3]
	// var rangeTwo = names[2:]
	// var rangeThree = names[:3]

	// fmt.Println(rangeOne, rangeTwo, rangeThree)

	// rangeOne = append(rangeOne, "kooper")

	// fmt.Println(rangeOne)


	// undesternding Go Packages

	// greatings := "hello my friends"

	// fmt.Println(strings.Contains(greatings, "Hello?"))
	// fmt.Println(strings.ReplaceAll(greatings, "Hello", "Hi"))

	// fmt.Println("original string Value is", greatings )

	// fmt.Println(strings.ToUpper(greatings))

	// fmt.Println(strings.Index(greatings, "ll"))

	// fmt.Println(strings.Split(greatings, ""))

	// ages := []int{30,25,26,29,28,27,24}

	// sort.Ints(ages)

	// fmt.Println(ages)

	// index := sort.SearchInts(ages, 30)

	// fmt.Println(index)

	// names := []string{"bahat", "Karamuka", "kelon", "karabo", "aline"}

	// sort.Strings(names)
	// fmt.Println(names)

	// fmt.Println(sort.SearchStrings(names, "Karamuka"))



	// Loops

	// x := 0

	// for x <5 {
	// 	fmt.Println("the value of x is :", x)

	// 	x ++
	// }

//   for i := 0; i<5; i++{
// 	fmt.Println("the value of i is :", i)
//   }

// names := []string{"bahat", "Karamuka", "kelon", "karabo", "aline"}

// for i :=0; i<len(names); i++{
// 	fmt.Println("the length of i is :", i)
// }

// x := 0

// for x =0; x<5; x++{

// 	fmt.Println("the value of X is :", x)

// }

// for i:=0; i<5; i++{
// 	fmt.Println("the value of i is also :", i)
// }

//  names := []string{"bahati", "Karamuka", "kelon", "karabo", "aline"}

//  for i := 0; i < len(names); i++ {

// 	fmt.Println(names[i])
	
//  }
// for index, value := range names{

// 	fmt.Printf("the value at index %v is %v \n", index, value)
// }


// age := 45

// fmt.Println(age <=50)
// fmt.Println(age >=50)
// fmt.Println(age ==45)
// fmt.Println(age !=50)

// if age< 30 || age != 45 {

// 	fmt.Println("This condition is correct")

// }else if age <40 || age > 50 {

// 	fmt.Println("This conditions are all true")
	
// }else{
// 	fmt.Println("All the Above conditions are wrong")
// }

// if age <30  && age <40 || age <20 && age <40 {
// 	fmt.Println("The age is less than thirty")
// }else if age >60 {
// 	fmt.Println("the age is less that 60")
// }else{
// 	fmt.Println("All the a bove stetement all false")
// }


// names := []string{"bahati", "Karamuka", "kelon", "karabo", "aline"}

// for index, value := range names{
// 	if index == 1 {
// 		fmt.Println("the continuing at pos", index)

// 		continue
// 	}

// 	if index < 2 {
// 		fmt.Println("Breaking line at", index)

// 		break
// 	}

	

// 	fmt.Printf("The value at pos %v is %v \n", index, value)
// }

}
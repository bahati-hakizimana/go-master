package main


import(
	"fmt"
)

func main () {
	

	age := 67
	name := "Eric"
	points := 255.55

	// Print

	fmt.Print("hello, " )
	fmt.Print("people \n")
	fmt.Print("Line one \n")

	// Println
	fmt.Println("Hello Ninjas")
	fmt.Println("good bye Ninjas")
	fmt.Println("My age is", age, "and My name is ", name)

	// Printf(formated string) %v or %_ = format specifier

	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("My age is %q and my name is %q \n", age, name)
	fmt.Printf("My age type is %T and my name type is %T \n", age, name)
	fmt.Printf("You scored %0.1f ponts \n", points)



	// Sprintf ( save formated string)

	var str = fmt.Sprintf("My age type is %T and my name type is %T \n", age, name)
	fmt.Println(str)


	var employeesAge = [3]int{20,30,40}
	names := [4]string{"Carl", "Olen", "Keisha", "kelon"}
	names[1] = "keria"

	fmt.Println(employeesAge, len(employeesAge))
	fmt.Println(names, len(names))


	// slices(array under the hood)

	var scores = []int{100, 200, 300,400}

	scores[3] = 600

	scores = append(scores, 78)

	fmt.Println(scores, len(scores))

	// Slice ranges

	rangeOne := names[1: 3]
	var rangeTwo = names[2:]
	var rangeThree = names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "kooper")

	fmt.Println(rangeOne)


}
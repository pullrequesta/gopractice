package main

func main() {
	// names := []string{"mario", "luigio", "peach"}
	// for _, value := range names {
	// 	fmt.Printf("the value is %s\n", value)
	// }
	// fmt.Println(names)

	// age := 30
	// if age < 45 {
	// 	fmt.Println("age is less than 45")
	// } else if age > 45 {
	// 	fmt.Println("age is greater than 45")
	// } else {
	// 	fmt.Println("age is perfect")
	// }

	// names := []string{"mario", "luigio", "peach", "lisa"}
	// for index, value := range names {
	// 	if index == 1 {
	// 		fmt.Println("continuing at pos", index)
	// 		continue
	// 	}
	// 	if index > 2 {
	// 		fmt.Println("breaking at pos", index)
	// 		break
	// 	}
	// 	fmt.Printf("the value at position %v is %v \n", index, value)
	// }

	// sayGreeting("mario")
	// sayBye("ludia")

	// cycleNames([]string{"mario", "luigio", "peach"}, sayGreeting)
	// fmt.Println(circleArea(10.5))

	// menu := map[string]float64{
	// 	"soup": 65.2,
	// 	"milk": 23.3,
	// 	"pie":  6.4,
	// }
	// fmt.Println(menu)
	// for k, v := range menu {
	// 	fmt.Println(k, "-", v)
	// }
	// phonebook := map[int]string{
	// 	372983: "mario",
	// 	232423: "ludia",
	// 	213434: "axel",
	// }
	// fmt.Println(phonebook[372983])

	// myBill := newBill("mario's bill")
	// fmt.Println(myBill)

	// myBill.updateTip(10)

	// myBill.items["cake"] = 2

	// myBill.addItem("fish", 4.9)

	// fmt.Println(myBill.format())

	// name := "mario"
	// fmt.Println(name)

	// m := &name
	// updateName(m)
	// fmt.Println(name)

	s := square{length: 4, width: 3}
	c := circle{radius: 5}

	measure(s)
	measure(c)

}

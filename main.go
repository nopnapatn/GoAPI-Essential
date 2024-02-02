package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Weight float32
	Height float32
	Grade  string
}

func main() {
	var fullname string = "neuw"
	var age int = 20
	fmt.Printf("Hello %s, Age: %d\n", fullname, age)

	test := "test"
	// test := "sd" can't change value
	fmt.Printf("%s\n", test)

	// while loop
	i := 0
	for i < 10 {
		fmt.Printf("number: %d\n", i)
		i++
	}

	var appArray [3]string
	appArray[0] = "test1"
	appArray[1] = "test2"
	appArray[2] = "test3"
	// can't use with this
	// appArray = ["test1", "test2", "test3"]
	// and if appArray[3] value in array is equal only 3!
	fmt.Println(appArray)

	for i := 0; i < len(appArray); i++ {
		fmt.Println(appArray[i])
	}

	appSlice := []int{10, 20, 30, 40, 50}
	fmt.Println(appSlice)
	// length
	fmt.Println(len(appSlice))
	// capacity
	fmt.Println(cap(appSlice))

	subSlice := appSlice[1:3]
	fmt.Println(subSlice)
	fmt.Println(len(subSlice))
	fmt.Println(cap(subSlice))

	appMap := make(map[string]int)
	appMap["test1"] = 5
	appMap["test2"] = 10
	appMap["test3"] = 15
	fmt.Println("Output:", appMap)

	delete(appMap, "test2")

	for key, value := range appMap {
		fmt.Printf("%s -> %d\n", key, value)
	}

	val, ok := appMap["test1"]
	if ok {
		fmt.Println("pear's value:", val)
	} else {
		fmt.Println("pear not found inmap")
	}

	var student Student
	student.Name = "neuw"
	student.Weight = 50
	student.Height = 170
	student.Grade = "A"
	fmt.Println(student)
}

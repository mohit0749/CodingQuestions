package main

import "reflect"

//import "fmt"

type MyCalendarTwo struct {
	mp map[int]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{make(map[int]int)}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	this.mp[start]++
	this.mp[end]--
	bookings:=0
	for _,v:=range this.mp{
		bookings+=v
		if bookings==3{
			this.mp[start]--
			this.mp[end]++
			return false
		}
	}
	return true
}

/*
["MyCalendarTwo","book","book","book","book","book","book"]
[[],[10,20],[10,50],[20,40],[40,50],[50,60],[60,70]]

["MyCalendarTwo","book","book","book","book","book","book","book","book","book","book"]
[[],[24,40],[43,50],[27,43],[5,21],[30,40],[14,29],[3,19],[3,14],[25,39],[6,19]]
*/

func Invoke(any interface{}, name string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	obj := reflect.ValueOf(any)
	if !obj.IsValid() {
		panic("not a valid")
	}
	meth := obj.MethodByName(name)
	if !meth.IsValid() {
		panic("not a valid method")
	}
	for _, v := range meth.Call(inputs) {
		println(v.Bool())
	}
}

func main() {
	obj := Constructor()
	fn := []string{"Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book", "Book"}
	values := [][2]int{{47, 50}, {1, 10}, {27, 36}, {40, 47}, {20, 27}, {15, 23}, {10, 18}, {27, 36}, {17, 25}, {8, 17}, {24, 33}, {23, 28}, {21, 27}, {47, 50}, {14, 21}, {26, 32}, {16, 21}, {2, 7}, {24, 33}, {6, 13}, {44, 50}, {33, 39}, {30, 36}, {6, 15}, {21, 27}, {49, 50}, {38, 45}, {4, 12}, {46, 50}, {13, 21}}
	for i := 0; i < len(fn); i++ {
		Invoke(&obj, fn[i], values[i][0], values[i][1])
	}
}

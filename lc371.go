package main

func getSum(a int, b int) {

}

func addNumberBitwise(a, b int) int {
	for b!=0{
		carry:=a&b
		a=a^b
		b=carry<<1
	}
	return a
}


func main() {
	a := -6
	b := -3
	n := addNumberBitwise(a, b)
	println(n)
	//println(getTwoComplements(n))
}

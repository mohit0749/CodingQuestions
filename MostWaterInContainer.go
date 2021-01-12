package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	l:=0
	r:=len(height)-1
	mxWater:=(r-l)* min(height[l],height[r])
	for i,j:=l+1,r-1;i<r;i,j=i+1,j-1{
		l=i
		if height[j]>height[r]{
			r=j
		}
		water:=(r-l)* min(height[l],height[r])
		if mxWater<water{
			mxWater=water
		}
	}
	return mxWater
}

func main(){
	height := []int{1,2,4,3}
	fmt.Println(maxArea(height))
}
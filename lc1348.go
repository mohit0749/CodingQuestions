package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

const (
	Minute=60
	Hour=3600
	Day=86400
)
type TweetCounts struct {
	tweets map[string][]int
}


func Constructor() TweetCounts {
	return TweetCounts{make(map[string][]int)}
}


func (this *TweetCounts) RecordTweet(tweetName string, time int)  {
	if _,ok:=this.tweets[tweetName];!ok{
		this.tweets[tweetName]=make([]int,0)
	}
	i:=0
	for ;i<len(this.tweets[tweetName]);i++{
		if time<this.tweets[tweetName][i]{
			break
		}
	}
	this.insert(tweetName,i,time)
}

func (this *TweetCounts)insert(tweetName string,at,time int){
	temp:=make([]int,0)
	temp=append(temp,this.tweets[tweetName][:at]...)
	temp=append(temp,time)
	temp=append(temp,this.tweets[tweetName][at:]...)
	this.tweets[tweetName]=temp
}


func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	var incrBy int
	switch freq {
	case "minute":
		incrBy = 60
	case "hour":
		incrBy = 3600
	case "day":
		incrBy = 86400
	}
	timeChunks:=createChunks(startTime,endTime,incrBy)
	counts:=make([]int,len(timeChunks))
	idx1:=-2
	idx2:=-2
	// fmt.Println("tweets = ",this.tweets)
	// fmt.Println("start,end = ",freq,startTime,endTime)
	// fmt.Println("chunks = ",timeChunks)
	for i,ch:=range timeChunks{
		idx1=ub(this.tweets[tweetName],ch[0])
		idx2=lb(this.tweets[tweetName],ch[1])
		// fmt.Println("chunk = ",ch," idx1 =",idx1," idx2 = ",idx2)
		if idx1==-1 && idx2==-1{
			continue
		}
		counts[i]=idx2-idx1+1
	}
	// fmt.Println("============================================")

	return counts
}

func lb(tweets []int,target int)int{
	if len(tweets)==0{
		return -1
	}
	low:=0
	high:=len(tweets)-1
	for low<high{
		mid:=low+(high-low)/2
		if tweets[mid]==target{
			return mid
		}else if tweets[mid]<target{
			low=mid+1
		}else{
			high=mid-1
		}
	}
	if tweets[low]>target{
		return low-1
	}
	return low
}
func ub(tweets []int,target int)int{
	if len(tweets)==0{
		return -1
	}
	low:=0
	high:=len(tweets)-1
	for low<high{
		mid:=low+(high-low)/2
		if tweets[mid]==target{
			return mid
		}else if tweets[mid]<target{
			low=mid+1
		}else{
			high=mid-1
		}
	}
	if tweets[low]<target{
		return low+1
	}
	return low
}

func createChunks(start,end int,incrby int)[][2]int{
	chunks:=make([][2]int,0)
	i:=start
	for ;;i+=incrby{
		if i-1>=end{
			break
		}
		chunks=append(chunks,[2]int{i,min(i-1+incrby,end)})
	}
	return chunks
}

func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}


/**
 * Your TweetCounts object will be instantiated and called as such:
 * obj := Constructor();
 * obj.RecordTweet(tweetName,time);
 * param_2 := obj.GetTweetCountsPerFrequency(freq,tweetName,startTime,endTime);
 */
/*
Test case:


["TweetCounts","recordTweet","recordTweet","recordTweet","getTweetCountsPerFrequency","getTweetCountsPerFrequency","recordTweet","getTweetCountsPerFrequency"]
[[],["tweet3",0],["tweet3",59],["tweet3",120],["hour","tweet3",0,59],["minute","tweet3",0,60],["tweet3",120],["hour","tweet3",0,210]]
*/

func main() {
	obj := Constructor()
	fn := []string{"recordTweet", "recordTweet", "recordTweet", "getTweetCountsPerFrequency", "getTweetCountsPerFrequency", "recordTweet", "getTweetCountsPerFrequency"}
	vv := []byte(`[["tweet3",0],["tweet3",59],["tweet3",120],["hour","tweet3",0,3600],["minute","tweet3",0,60],["tweet3",120],["hour","tweet3",0,210]]`)
	val := make([]interface{}, 0)
	_ = json.Unmarshal(vv, &val)
	for i := 0; i < len(fn); i++ {
		args := val[i].([]interface{})
		if fn[i]=="getTweetCountsPerFrequency"{
			Invoke(&obj, strings.Title(fn[i]), args[0].(string),args[1].(string), int(args[2].(float64)),int(args[3].(float64)))
		}else{
			Invoke(&obj, strings.Title(fn[i]), args[0].(string), int(args[1].(float64)))
		}
	}
}

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

type DemoServerice struct{}

type Args struct {
	A, B int
}

/**
相除
*/
//使用
func (DemoServerice) div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("fuck out")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}
func switch1() {
	isSpace := func(char byte) bool {
		switch char {
		case ' ': // 空格符会直接 break，返回 false // 和其他语言不一样
			log.Println("1")
			//return false
			fallthrough // 返回 true
		case '\t':
			log.Println("2")
			return false
		}
		return true
	}
	fmt.Println(isSpace('\t')) // true
	fmt.Println(isSpace(' '))  // false
}

func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	////低保一秒之后结束所有进程
	//defer cancel()
	//go handle(ctx, 500*time.Millisecond)
	//
	//select {
	//case <-ctx.Done():
	//	fmt.Println("main", ctx.Err())
	//}
	nums := []int{1, 7, 3, 6, 5, 6}
	log.Println(pivotIndex(nums))
}
func pivotIndex(nums []int) int {
	leftSum := 0
	rightSum := 0
	for i := 0; i < len(nums); i++ {
		rightSum += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		if i != 0 {
			//避开 两边极端
			leftSum += nums[i-1]
		}
		rightSum -= nums[i]
		if rightSum == leftSum {
			return i
		}
	}
	return -1
}
func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		log.Println("handle", ctx.Err())
	case <-time.After(duration):
		log.Println("process request with", duration)
	}

}

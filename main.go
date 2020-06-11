package main

import (
	"encoding/json"
	"fmt"
	"leetcode/model"
)

func main() {
	//GetTwoSum() //1
	//testSum()   //2
	JsonDemo()
}

func JsonDemo() {
	t := model.JsonDemo{
		Name:    "lanlan",
		Age:     "18",
		Address: "wojia",
	}

	tjson, _ := json.Marshal(t)

	fmt.Println("tjson==", string(tjson))
}

//01.给定一个整数数组nums和一个目标值target，请你在数组中找出和为目标值的那两个整数，并返回他们的数组下标

func GetTwoSum() {
	nums := []int{2, 11, 11, 7}
	target := 9

	date := TwoSum(nums, target)
	fmt.Println("data==", date)
}

func TwoSum(nums []int, target int) []int {

	v := make(map[int]int)

	for i, _ := range nums {
		dif := target - nums[i]
		tt, ok := v[dif]
		if ok {
			return []int{i, tt}
		}

		v[nums[i]] = i
	}

	return []int{-1, -1}
}

//02.两个非空链表表示两个非负整数，他们的位数按照逆序来存储，每个节点只能存储一位数字。
//如果将着两个数相加，将会返回一个新的链表表示他们的和。并且连个链表都不会以0来开头。

type ListNode struct {
	Val   int
	Next1 *ListNode
}

func AddTwoList(l1 *ListNode, l2 *ListNode) *ListNode {

	var result *ListNode

	num := 0
	Next := &result

	for l1 != nil || l2 != nil || num > 0 {

		if l1 != nil {
			num += l1.Val
			l1 = l1.Next1
		}

		if l2 != nil {
			num += l2.Val
			l2 = l2.Next1
		}

		*Next = &ListNode{
			Val:   num % 10,
			Next1: nil,
		}
		num = num / 10

		Next = &(*Next).Next1

	}

	return result

}

//遍历链表的值
func Ergodic(l *ListNode) {

	for l != nil {
		fmt.Print(l.Val)
		l = l.Next1
	}
}

//初始化链表
func GetTwoList() (l1, l2 *ListNode) {

	s1 := &ListNode{
		Val:   2,
		Next1: nil,
	}
	s2 := &ListNode{
		Val:   5,
		Next1: s1,
	}
	s3 := &ListNode{
		Val:   8,
		Next1: s2,
	}

	r1 := &ListNode{
		Val:   2,
		Next1: nil,
	}
	r2 := &ListNode{
		Val:   5,
		Next1: r1,
	}
	r3 := &ListNode{
		Val:   8,
		Next1: r2,
	}

	return s3, r3

}

func testSum() {

	l1, l2 := GetTwoList()

	Ergodic(l1)
	fmt.Println()
	Ergodic(l2)
	fmt.Println()

	l3 := AddTwoList(l1, l2)
	Ergodic(l3)

}

package main

import (
	"bytes"
	"fmt"
)

type merge struct {
	sum      int
	strings  string
	sumFloat float64
}

func init() {
	fmt.Println("start")
}

func main() {
	fmt.Println(unique([]string{"a", "b", "b", "d", "f"}))
	fmt.Println(unique("sjffahfadsfhjsadsjad"))
	//fmt.Println(groupBy([]int{3, 12, 2, 1, 2, 3, 56}))
	//test1("123")
	//test2("Hello", 1, 9324, 12, ", ", "World")
}

func unique(v interface{}) []string {
	var newStr []string
	switch v.(type) {
	case string:
		str := v.(string)
		lens := len(str)
		for i := 0; i < lens; i++ {
			if i == lens-1 {
				newStr = append(newStr, string(str[i]))
				break
			}

			if str[i] != str[i+1] {
				newStr = append(newStr, string(str[i]))
			}
		}
		break
	case []string:
		str := v.([]string)
		lens := len(str)
		for i := 0; i < lens; i++ {
			if i == lens-1 {
				newStr = append(newStr, str[i])
				break
			}

			if str[i] != str[i+1] {
				newStr = append(newStr, str[i])
			}
		}
		break
	default:
		break
	}

	return newStr
}

func groupBy(number []int) [][]int {
	var newArray [][]int
	tmp := map[int]int{}
	for _, v := range number {
		var ok bool
		_, ok = tmp[v]
		if !ok {
			tmp[v] = 1
			newArray = append(newArray, []int{v})
		}
	}

	return newArray
}

// 实现传单个任何类型的参数
func test1(v interface{}) {
	fmt.Println(v)
}

// 实现传任意多个任意类型的参数
func test2(v ...interface{}) {
	m := merge{}
	for _, val := range v {
		m.interface2Type(val)
	}
	fmt.Println(m)
}

func (m *merge) interface2Type(inter interface{}) {

	switch inter.(type) {

	case string:
		bt := bytes.Buffer{}
		bt.WriteString(m.strings)
		bt.WriteString(inter.(string))
		m.strings = bt.String()
		break
	case int:
		m.sum += inter.(int)
		break
	case float64:
		m.sumFloat += inter.(float64)
		break
	}

}

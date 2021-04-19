package algorithm

import (
	"fmt"
)

//不使用任何中间变量，将两个数进行交换
func Change2Num()  {
	//例如，a = 10,b = 20
	//方法1，a = a + b; b = a - b; a = a - b.当数字很大时，可能会产生溢出。
	//方法2，^运算：相当于十进制不同位做加法；
	//&运算：相当于十进制 相同位做加法的1/2；ex: 0101 & 0011   结果：二进制0001           十进制 (2^0 +2^0)/2   这里的"^"代表次幂
	//^运算：相当于十进制 不同位做加法的1/2；ex: 0101 ^ 0011   结果：二进制0110           十进制（2^2 + 2^1 ）
	//|运算：相当于十进制 相同位做加法的1/2与不同位做加法求和。ex: 0101 | 0011   结果：二进制0111  十进制 (2^0 +2^0)/2 +(2^2 +2^1)
	//即a = a ^ b;b = a ^ b;a = a ^ b.因为 b = a^b^b = a^(b^b) = a^0 = a
	a := 1
	b := 2
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Printf("a=%d\n", a)
	fmt.Printf("b=%d\n", b)
}
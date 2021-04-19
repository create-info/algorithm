package algorithm

//fibonacci numbers 斐波那契数列
//ex: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610,...
//方法1：递归 确定 当n很大时，耗时太慢，大量重复计算
func Fibonacci1(n int64) int64 {
	var a int64 = 0
	var b int64 = 1
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return a + b
	}
	return Fibonacci1(n-1) + Fibonacci1(n-2)
}

//动态规则
func Fibonacci2(n int64) int64 {
	var a int64 = 0
	var b int64 = 1
	if n <= 0 {
		return 0
	}
	var k int64 = 0
	for i:=int64(1); i<n; i++ {
		k = a
		a = b
		b = k + b
	}
	return b
}

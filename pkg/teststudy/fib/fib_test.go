package fib

import "testing"

// 性能比较测试
// go test -bench=Fib2
// go test -bench=.
// 指定执行时间：go test -bench=Fib2 -benchtime=20s
// 指定cpu数:
// b.SetParallelism(n)
// go test -bench=Split -cpu=1
func benchmarkFib(b *testing.B,n int) {
	for i:=0;i<b.N;i++{
		Fib(n)
	}
}

func BenchmarkFib2(b *testing.B){
	benchmarkFib(b,2)
}
func BenchmarkFib3(b *testing.B){
	benchmarkFib(b,3)
}
func BenchmarkFib10(b *testing.B){
	benchmarkFib(b,10)
}
func BenchmarkFib20(b *testing.B){
	benchmarkFib(b,20)
}
func BenchmarkFib40(b *testing.B){
	benchmarkFib(b,40)
}


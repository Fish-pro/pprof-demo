package fib

func Fib(n int)int{
	if n<2{
		return n
	}else {
		return Fib(n-1)+ Fib(n-2)
	}
}

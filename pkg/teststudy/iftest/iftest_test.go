package iftest

import "testing"

func BenchmarkSumA(b *testing.B) {
	list := []bool{false, true, false, true, false}
	count := 0
	for n := 0; n < b.N; n++ {
		for _, c := range list {
			if c {
				count++
			}
		}
	}
}

func BenchmarkSumB(b *testing.B) {
	list := []bool{false, true, false, true, false}
	count := 0
	for n := 0; n < b.N; n++ {
		for _, c := range list {
			if c == true {
				count++
			}
		}
	}
}

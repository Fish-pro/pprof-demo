package split

import (
	"reflect"
	"testing"
)

func Test0Split(t *testing.T){
	type testCase struct{
		Str string
		Sep string
		Want []string
	}
	tests := []testCase{
		{Str: "abc",Sep: "b",Want: []string{"a","c"}},
		{Str: "abcdbe",Sep: "b",Want: []string{"a","cd","e"}},
		{Str: "abcd",Sep: "bc",Want: []string{"a","d"}},
	}
	for _,tc :=range tests{
		got := Split(tc.Str,tc.Sep)
		if !reflect.DeepEqual(got,tc.Want){
			t.Errorf("want:%v,but got:%v",tc.Want,got)
		}
	}
}


// go test -v
// 子测试：go test -run=Test0Split/case_3 -v
// 测试覆盖率：go test -cover
// html工具：
// go test -coverprofile=c.out
// go tool cover -html=c.out
func Test1Split(t *testing.T){
	type testCase struct{
		Str string
		Sep string
		Want []string
	}
	tests := map[string]testCase{
		"case 1":{Str: "abc",Sep: "b",Want: []string{"a","c"}},
		"case 2":{Str: "abcdbe",Sep: "b",Want: []string{"a","cd","e"}},
		"case 3":{Str: "abcd",Sep: "bc",Want: []string{"a","d"}},
	}
	for name,tc :=range tests{
		teardownCase := setupTestCase(t)
		defer teardownCase(t)
		t.Run(name,func(t *testing.T){
			got := Split(tc.Str,tc.Sep)
			if !reflect.DeepEqual(got,tc.Want){
				t.Errorf("name:%v,want:%v,but got:%v",name,tc.Want,got)
			}
		})
	}
}

// 基准测试
// go test -bench=Split
// 内存：go test -bench=Split -benchmem
func BenchmarkSplit(b *testing.B) {
	for i:=0;i<b.N;i++{
		Split("abcbcbcbcbcb","b")
	}
}

func setupTestCase(t *testing.T)func(t *testing.T){
	t.Log("测试之前的工作")
	return func(t *testing.T){
		t.Log("测试之后的工作")
	}
}
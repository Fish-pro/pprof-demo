# golang测试
## 单元测试
```golang
func TestSplit(t *testing.T){
	got := Split("a:b:cw",":")
	want := []string{"a","b","cw"}
	if !reflect.DeepEqual(got,want){
		t.Errorf("want:%v,but got:%v",want,got)
	}
}

func Test2Split(t *testing.T){
	got := Split("a:b:c",":")
	want := []string{"a","b","c"}
	if !reflect.DeepEqual(got,want){
		t.Errorf("want:%v,but got:%v",want,got)
	}
}

func Test3Split(t *testing.T){
	got := Split("abcd","bc")
	want := []string{"a","b"}
	if !reflect.DeepEqual(got,want){
		t.Fatalf("want:%v,but got:%v",want,got)
	}
}
```
## 子测试
```golang
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
```
```bash
 go test -v
// 子测试：go test -run=Test0Split/case_3 -v
// 测试覆盖率：go test -cover
// html工具：
go test -coverprofile=c.out
go tool cover -html=c.out
```
## 基准测试
```golang
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
```
```bash
go test -bench=Split
// 内存：
go test -bench=Split -benchmem
```
## 性能比较测试
```golang
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
```
```bash
go test -bench=Fib2
go test -bench=.
// 指定执行时间：
go test -bench=Fib2 -benchtime=20s
// 指定cpu数:
// b.SetParallelism(n)
go test -bench=Split -cpu=1
```
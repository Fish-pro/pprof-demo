package split

import (
	"reflect"
	"testing"
)

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
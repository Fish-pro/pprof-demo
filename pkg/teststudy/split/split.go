package split

import "strings"

// Split 切割字符串
// example:
// abc，b => ["a","b"]
func Split(str string,sep string) []string{
	//var ret []string
	ret := make([]string,0,strings.Count(str,sep)+1)
	index := strings.Index(str,sep)
	for index >= 0{
		if str[:index] != ""{
			ret = append(ret,str[:index])
		}
		str = str[index+len(sep):]
		index = strings.Index(str,sep)
	}
	if str != ""{
		ret = append(ret,str)
	}
	return ret
}

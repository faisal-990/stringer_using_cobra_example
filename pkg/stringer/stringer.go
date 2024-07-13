package stringer

import (
	"strconv"
)
func Reverse(input string) (result string) {
	for _,c:=range(input) {
		result=string(c)+result
	}
	return result

}

func Inspect(input string ,digit bool) (count int,kind string) {
	//return the length of the string and inspect if it has any digit
	if !digit {
		return len(input), "char"
	}
	return inspectNumbers(input),"digit"
}
func inspectNumbers(input string) (count int) {
	//number in the form of string to be converted to int and then counted
	for _,c:=range input {
		_,err:=strconv.Atoi(string(c));
		if err == nil {
			count++
	}
	}
	return count;
	//now i want these reverse and inspect function to a part of my command 
	//i will create files in the root.go part and init the commands along with the rootCmd 
	//which is the root command (starting index) of my cli code:
}

package main

import (
	"fmt";
	"strconv";
	"strings";
	"os"
)

// adds/subtracts/multiplies all values that are in the *values array.
// nValues is the number of values you're reading from the array
// operator will indicate if it's an addition (1), subtraction (2) or
// multiplication (3)

func calc(operator int, values []int) int {
	res := values[0]
	if(operator == 1){
		for i := 1; i < len(values); i++ {
			res = res + values[i]
		}
	}else {
		if(operator == 2){
			for i := 1; i < len(values); i++ {
				res = res - values[i]
			}
		}else {
			for i := 1; i < len(values); i++ {
				res = res * values[i]
			}
		}
	}

	fmt.Println("Resultado: ", res)

	return 0
}

func main() {
	var arr []int;
	for i := 2; i < len(os.Args); i++ {
		num, e := strconv.Atoi(os.Args[i]);
		if (e != nil){
			fmt.Println(e)
		}
		arr = append(arr, num)
	}

	op :=0 ;

	if(strings.Compare(os.Args[1], "add") == 0){
		op = 1
	} else {
		if (strings.Compare(os.Args[1], "sub") == 0){
			op = 2
		} else {
			if (strings.Compare(os.Args[1], "mult") == 0){
				op = 3
			}else {
				fmt.Println("ERROR")
				return
			}
		}
	}

	calc(op, arr)

}
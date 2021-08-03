package main

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {
	fmt.Println([]byte(s))
	var res []string
	var getVal string
	for i, x :=  range s{

		if x >= 48 && x <= 57 {
			getVal += string(x)
		}

		if x == 91{

			for j:=i+1 ; j < len(s) -1; j++{
				if s[j] == 93{
					getInt, _ := strconv.Atoi(getVal)
					getVal = ""
					for k:=0; k < getInt ; k++{
						getVal += s[i+1:j]
					}

					res = append(res, getVal)

					getVal = ""

				}
				break
			}
		}

	}
	fmt.Println(res)
	return ""
}

func main()  {
	decodeString("3[a]2[bc]")
	//decodeString("09") // 48 57
	//decodeString("[]") // 91 - 93
	//decodeString("az") // 97 122
	//decodeString("AZ") // 65 90
}
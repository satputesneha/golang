package main

import "fmt"

func main() {
	fmt.Println("***")
	fmt.Println(checke("{[()]}"))
	fmt.Println(checke("{[[)]}"))

}

func checke(photo string) bool {
	braces := make(map[string]string)
	stack := make([]string, 0)
	braces["{"] = "}"
	braces["["] = "]"
	braces["("] = ")"

	//"{[()]}"

	for i := 0; i < len(photo); i++ {

		person := string(photo[i])

		_, exists := braces[person]
		if exists {
			stack = append(stack, person)

		} else {
			if len(stack) == 0 {
				return false

			}
			top := stack[len(stack)-1]
			if person == braces[top] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		}

		fmt.Println(person, stack)

	}
	return true
}

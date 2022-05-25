package users

import "fmt"

var Name string //name in uppercase this mean is public i can use it in this package or others package lowercase mean it is private i can use it in this package only

func Reply() {

	fmt.Println("u r welcome ", Name)
}

package test

import (
	"backend-go/utils/pwd"
	"fmt"
)

func main() {
	hash := pwd.HashPwd("1234")
	hash1 := pwd.HashPwd("1234")
	fmt.Println(hash, hash1)

	ok := pwd.CheckPwd(hash, "1234")
	fmt.Println(ok)
	ok1 := pwd.CheckPwd(hash1, "1234")
	fmt.Println(ok1)

}

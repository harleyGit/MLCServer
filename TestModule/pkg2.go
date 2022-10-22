/*
 * @title:
 * @Author: gang.huang
 * @Date: 2021-10-06 13:11:09
 * @LastEditTime: 2021-10-06 14:47:20
 * @FilePath: /GoDemo/TestModule/pkg2.go
 */
package pkg2

import "fmt"

func ExecPkg2() {
	fmt.Println("Exec Pkg2")
}

func init() {
	fmt.Println("pkg2 init")
}

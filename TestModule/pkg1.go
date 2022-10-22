/*
 * @title:
 * @Author: gang.huang
 * @Date: 2021-10-06 13:09:33
 * @LastEditTime: 2021-10-06 14:21:18
 * @FilePath: /GoDemo/TestModule/pkg1.go
 */
package pkg1

import (
	"fmt"
	// "pkg2"
)

func ExecPkg1() {
	fmt.Println("Exec Pkg1")

	// pkg2.ExecPkg2()

}

func init() {
	fmt.Println("pkg1 init")
}

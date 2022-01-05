package main

import (
	"fmt"
	"os"
)

func createFile(fileFullPath string) *os.File {
	//该方法最后的perm指定的是linux unix 等操作系统层面的权限，在windows系统中该参数无效，所有者权限_当前用户组权限_其他用户组的权限 r-4 w-2 x-1
	file, err := os.OpenFile(fileFullPath, os.O_WRONLY|os.O_CREATE|os.O_RDONLY, 666)
	if err != nil {
		fmt.Printf("openFile failed, err:%v\n", err)

	}
	return file
}

func main() {

}

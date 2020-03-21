package helper

import (
	"os"
	"fmt"
	"path/filepath"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dir string) {
	exist, err := PathExists(dir)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
	}

	if exist {
		fmt.Printf("has dir![%v]\n", dir)
	} else {
		fmt.Printf("no dir![%v]\n", dir)
		// 创建文件夹
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
}

func isExist(path string) (bool) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

func CreateFile(files string) {
	var path, fileName  = filepath.Split(files)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	if isExist(fileName) {
		return
	} else {
		file,err:=os.Create(files)
		if err!=nil{
			fmt.Println(err)
		}
		defer file.Close()
		return
	}
}

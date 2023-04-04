package basefunc

import (
	"fmt"
	"log"
	"testing"
)

func Test_CheckDir(t *testing.T) {
	res, err := CheckDir("/Users/fuao/Desktop/code/github/auto-clean")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res)

}

func Test_FilterFileList(t *testing.T) {
	res, err := FilterFileList("/Users/fuao/Desktop/code/github/auto-clean/test/", "zzzzzyyyyy","-3600s")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res)

}

func Test_ReadFileList(t *testing.T) {
	res, err := ReadFileList("/Users/fuao/Desktop/code/github/auto-clean/test/")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res)

}

func Test_GetDateFile(t *testing.T) {
	res := GetDateFile("/Users/fuao/Desktop/code/github/auto-clean/test/asd")
	fmt.Println(res)

}

func Test_CalculateTimeDifference(t *testing.T) {
	res := GetDateFile("/Users/fuao/Desktop/code/github/auto-clean/test/asd")
	fmt.Println("文件时间：", res)
	CalculateTimeDifference(res, "-1200s")

}

func Test_RemoveFile(t *testing.T) {
	res, err := FilterFileList("/Users/fuao/Desktop/code/github/auto-clean/test/", "zzzzzyyyyyy","-3600s")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res)

	RemoveFile(res)

}

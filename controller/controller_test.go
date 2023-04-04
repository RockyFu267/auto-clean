package controller

import (
	"fmt"
	"log"
	"testing"
)

func Test_CheckDir(t *testing.T) {
	res, err := ReadConfig("/Users/fuao/Desktop/code/github/auto-clean")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res, res.Service[0].Cleanpath)

}

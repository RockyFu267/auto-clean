package basefunc

import (
	"auto-clean/basecmd"
	"log"
	"os"
	"strings"
)

//FilterFileList dir路径 filterName为名称过滤条件 filterTime为时间过滤条件
func FilterFileList(dir string, filterName string, filterTime string) ([]string, error) {
	var res []string
	resTmp, err := ReadFileList(dir)
	if err != nil {
		log.Println(err)
		return resTmp, err
	}
	for _, v := range resTmp {
		resBool := strings.HasPrefix(v, filterName)
		if resBool {
			vTime := GetDateFile(dir+"/"+v)
			boolTmp:=CalculateTimeDifference(vTime,filterTime)
			if boolTmp {
				res = append(res, dir+"/"+v)
			}else{
				continue
			}
		} else {
			continue
		}
	}
	return res, nil
}

//ReadFileList 获取所有文件列表 过滤文件夹
func ReadFileList(dir string) ([]string, error) {
	resTmp, err := basecmd.CmdAndChangeDirToResAllInOne(dir, "ls -l|grep -v '^d' | awk '(NR > 1){print $9}'")
	if err != nil {
		log.Println(err)
		return resTmp, err
	}

	return resTmp, nil
}

//RemoveFile
func RemoveFile(input []string) {
	for _, v := range input {
		err := os.Remove(v)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

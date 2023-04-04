package basefunc

import (
	"os"
	"syscall"
	"time"
)

//CalculateTimeDifference Cyclelength应该为负值 返回true 说明应该在目标时间以前的状态，返回false，说明是近期目标时间以内的
func CalculateTimeDifference(Datefile time.Time, Cyclelength string) bool {
	nowTime := time.Now()
	pTmp, _ := time.ParseDuration(Cyclelength)
	T1 := nowTime.Add(pTmp)
	// fmt.Println("目标时间：", T1)
	// fmt.Println(Datefile.Before(T1))
	return Datefile.Before(T1)
}

func GetDateFile(file string) time.Time {
	finfo, _ := os.Stat(file)
	// Sys()返回的是interface{}，所以需要类型断言，不同平台需要的类型不一样，linux上为*syscall.Stat_t
	stat_t := finfo.Sys().(*syscall.Stat_t)
	return timespecToTime(stat_t.Mtimespec)
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

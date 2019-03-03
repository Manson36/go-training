package main

import (
	"fmt"
	"time"
)

func main(){

	now := time.Now()
	fmt.Printf("now= %v, now type=%T\n",now,now)

	//获取年月日 时分秒
	fmt.Printf("年=%v\n",now.Year())
	fmt.Printf("月=%v\n",now.Month())
	fmt.Printf("月=%v\n",int(now.Month()))
	fmt.Printf("日=%v\n",now.Day())
	fmt.Printf("时=%v\n",now.Hour())
	fmt.Printf("分=%v\n",now.Minute())
	fmt.Printf("秒=%v\n",now.Second())

	//格式化输出
	dateStr := fmt.Sprintf("当前的年月日 %d-%d-%d %d:%d:%d \n",now.Year(),
		now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Printf("dateStr=%v ",dateStr)

	//格式化输出方式2
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

}
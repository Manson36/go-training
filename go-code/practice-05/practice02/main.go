package main

import (
	"fmt"
	"strings"
)

func main(){

	//返回子串在的字符串第一次出现的index值
	index := strings.Index("TNL_abcabcabc","abc")
	fmt.Printf("index =%v\n",index)

	//返回子串在的字符串最后一次出现的index值
	index = strings.LastIndex("TNL_abcabcabc","abc")
	fmt.Printf("index =%v\n",index)

	//将指定的子串替换为另一个子串
	str2 := "go go hello"
	str := strings.Replace(str2,"go","北京",-1)
	fmt.Print("str =%v, str2=%v",str,str2)

	//按照某个指定的字符，为分割标识，将字符串拆成字符串组
	strArr := strings.Split("hello,world,ok",",")
	for i:=0; i< len(strArr); i++{
		fmt.Printf("str[%v]=%v\n",i,strArr[i])
	}
	fmt.Printf("strArr =%v",strArr)

	//将字符串的字母进行大小写转换
	str = "Hello goLang"
	str = strings.ToLower(str)
	str2 = strings.ToUpper(str)
	fmt.Printf("str=%v str2=%v\n",str,str2)

	//将字符串左右两边的空格去掉
	str = strings.TrimSpace(" ! Hello   !  ")
	fmt.Printf("str=%q\n",str)

	//将字符串左右两边指定的符号去掉
	str = strings.Trim("! Hello   !","!")
	fmt.Printf("str=%q\n",str)

	//将字符左边的！去掉
	str = strings.TrimLeft("! Hello   !","!")
	fmt.Printf("str=%q\n",str)

	//将字符右边的！去掉
	str = strings.TrimRight("! Hello   !","!")
	fmt.Printf("str=%q\n",str)

	//判断字符串是否已指定的字符串开头
	b:=strings.HasPrefix("ftp.//192.168.1.11","ftp")
	fmt.Printf("b=%v\n",b)

	//判断字符串是否已指定的字符串结尾
	b=strings.HasSuffix("ftp.//192.168.1.11","ftp")
	fmt.Printf("b=%v",b)
}
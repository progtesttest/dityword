package main

import (
	"dityword/dityword"
	"os"
	"path"
	"fmt"
	"log"
	"unicode/utf8"
	"github.com/henrylee2cn/mahonia"
	"bytes"
	"regexp"
)

func check(src string) bool  {
	str := "(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)" //此处改为“
	re, err := regexp.Compile(str)
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	b := re.MatchString(src)
	fmt.Println("lllll",b)  //打印出false。
	return  b
}

func main()  {

//	aa := "存在select"
//	check(aa)

//	bb := "存在"
//	check(bb)

//	cc := "select"
//	check(cc)
	//bb := "存在"
	//fmt.Println("bbbbb",re.MatchString(bb))  //打印出false。

	//1读取配置文件连
	cfgpath, _ := os.Getwd()
	filename := path.Join(cfgpath, "ditylist.txt")
	if !dityword.LoadDirtyWordsFile(filename) {
		os.Exit(1)
	}

	for {
		var input string
		fmt.Scanln(&input)
		log.Printf("input=%v len=%v \n",input,len(input))
		if utf8.ValidString(input) {

			enc:=mahonia.NewEncoder("gbk")
			gbkstr := enc.ConvertString(input)
			log.Printf("gbkstr=%v \n",[]byte(gbkstr))
			b := dityword.HasDirtyWords(gbkstr)
			usrc := bytes.Runes([]byte(input))
			log.Printf("check b=%v uscr=%#v %v\n",b,usrc,len(usrc))

		//	2018/05/26 00:02:12 input=日 len=3
		//	2018/05/26 00:02:12 gbkstr=[200 213]
		//	2018/05/26 00:02:12 check b=true uscr=[]int32{26085} 1

			//r, size := utf8.DecodeRuneInString(input)
			//fmt.Printf("%c %v\n", r, size)

		//	newdata := string(([]byte(input))[size:])
		//	fmt.Printf("%c %v  data=%v \n", r, size,newdata)
			//str = str[size:]
		//	if data,num := utf8.DecodeRuneInString(input); ok {
		//		b := dityword.HasDirtyWords(input)
		//		fmt.Printf("check b=%v \n",b)
		//	}

		}

	}


}

/*
testBytes := []byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1}

    var testStr string

    utfStr := "你好，世界！"

    var dec mahonia.Decoder

    var enc mahonia.Encoder



    testStr = string(testBytes)



    dec = mahonia.NewDecoder("gbk")

    if ret, ok := dec.ConvertStringOK(testStr); ok {

        fmt.Println("GBK to UTF-8: ", ret, " bytes:", testBytes)

    }



    enc = mahonia.NewEncoder("gbk")

    if ret, ok := enc.ConvertStringOK(utfStr); ok {

        fmt.Println("UTF-8 to GBK: ", ret, " bytes: ", []byte(ret))

    }
 */
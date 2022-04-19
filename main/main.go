package main

import (
	"fmt"
	"CipherSpider/spider"
)

func main(){
	key := make(map[rune]rune)
	fmt.Println(spider.SubEncode("Hello",key,true))
	fmt.Println(spider.SubEncode("Hello",key,false))
	key,_ = spider.QuickMap("qponzyxwumlkjihtsrgfedcbva")
	fmt.Println(spider.SubEncode("Hello",key,false))
	fmt.Println(spider.SubDecode("wzkkh",key,false))
	badKey,_ := spider.QuickMap("qponzyxsumlkjihtsrgfedcbva")
	fmt.Println(spider.SubDecode("wzkkh",badKey,false))
	fmt.Println(spider.SubDecode("wzkkh",badKey,true))
}
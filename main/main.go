package main

import (
	"fmt"
	"CipherSpider/spider"
)

func main(){
	key := make(map[rune]rune)
	fmt.Println(spider.SubEncode("Hello, world!",key,true))
	fmt.Println(spider.SubEncode("Hello, world!",key,false))
	key,_ = spider.QuickMap("mlkjihgfedcbazyxwvutsrqpon")
	fmt.Println(spider.SubEncode("Hello, world!",key,false))
	fmt.Println(spider.SubDecode("Fibby, qyvbj!",key,false))
	badKey,_ := spider.QuickMap("qponzyxsumlkjihtsrgfedcbva")
	fmt.Println(spider.SubDecode("Wzkkh, chrkn!",badKey,false))
	fmt.Println(spider.SubDecode("Wzkkh, chrkn!",badKey,true))
	fmt.Println(spider.CaesarEncode("Hello, world!",0))
	fmt.Println(spider.CaesarEncode("Hello, world!",1))
	fmt.Println(spider.CaesarEncode(spider.CaesarDecode("Hello, world!",1),1))
	fmt.Println(spider.Rot13("Hello, world!"))
	fmt.Println(spider.Rot13(spider.Rot13("Hello, world!")))
	fmt.Println(spider.Atbash("Hello, world!"))
	fmt.Println(spider.Atbash(spider.Atbash("Hello, world!")))
}
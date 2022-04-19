package main

import "CipherSpider/spider"

func main(){
	key := make(map[rune]rune)
	spider.SubEncode("Hello",key,true)
	spider.SubEncode("Hello",key,false)
	key,_ = spider.QuickMap("zyxwvutsrqponmlkjihgfedcba")
	spider.SubEncode("Hello",key,true)
	spider.SubEncode("Hello",key,false)
}
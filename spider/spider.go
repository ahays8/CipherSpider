package spider

import (
	"errors"
	"log"
	"strings"
	"unicode"
	"io/ioutil"
	//"math"
	//"fmt"
)


func QuickMap(alph string) (map[rune]rune,error){
	key := make(map[rune]rune)
	if(len(alph)!=26){
		alErr:=errors.New("26 rune alphabet required")
		log.Fatalln(alErr)
		return key, alErr
	}else{
		plain := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		plslice := []rune(plain)
		upper := []rune(strings.ToUpper(alph))
		lower := []rune(strings.ToLower(alph))
		for ind := range alph{
			key[plslice[ind]]=lower[ind]
			key[plslice[ind+26]]=upper[ind]
		}
	}
	return key, nil
}
func SubEncode(ptxt string, key map[rune]rune,fill bool) string{
	var ctxt []rune
	for _,pchar := range ptxt{
		cchar,exists := key[pchar]
		if(exists){
			ctxt = append(ctxt,cchar)
		}else if(!fill&&unicode.IsLetter(pchar)){
			ctxt = append(ctxt,'?')
		}else{
			ctxt = append(ctxt,pchar)
		}
	}
	return string(ctxt)
}

func SubDecode(ctxt string, key map[rune]rune,fill bool) string{
	var ptxt []rune
	for _,char := range ctxt{
		var found = false
		for dec,enc :=range key{
			if enc==char{
				ptxt = append(ptxt, dec)
				found = true
				break
			}
		}
		if(!found){
			if(!fill&&unicode.IsLetter(char)){
				ptxt = append(ptxt, '?')
			}else{
				ptxt = append(ptxt, char)
			}
		}
	}
	return string(ptxt)
}

func CaesarEncode(msg string, shft int) string{
	shft = shft%26
	for shft<0 {
		shft +=26
	}
	alph := "abcdefghijklmnopqrstuvwxyz"
	shiftkey,_ := QuickMap(alph[shft:]+alph[0:shft])
	return SubEncode(msg,shiftkey,true)
}

func CaesarDecode(msg string, shft int) string{
	return CaesarEncode(msg, -shft)
}

func Rot13(msg string) string{
	return CaesarEncode(msg, 13)
}

func Atbash(msg string) string{
	key,_ := QuickMap("zyxwvutsrqponmlkjihgfedcba")
	return SubEncode(msg,key,true)
}
func LetterNumber(msg string) []int{
	alph := "abcdefghijklmnopqrstuvwxyz"
	ctxt := []int{}
	for _,char := range msg{
		if(unicode.IsLetter(char)){
			ctxt=append(ctxt,strings.IndexRune(alph,unicode.ToLower(char)))
		}
	}
	return ctxt
}
func NumberLetter(msg []int) string{
	alph := "abcdefghijklmnopqrstuvwxyz"
	ptxt := ""
	for _,num := range msg{
		ptxt=ptxt+alph[num:num+1]
	}
	return ptxt
}

func VigenereEncode(msg string, key string) (string,error){
	newkey:=""
	for _,char:=range(key){
		if(unicode.IsLetter(char)){
			newkey+=string(char)
		}
		if(len(newkey)>=len(msg)){break}
	}
	if(len(newkey)<1){
		alErr:=errors.New("key is too short")
		log.Fatalln(alErr)
		return "", alErr
	}
	rkey:=[]rune(newkey)
	ctxt:=""
	ind:=0
	for _,char := range msg{
		if(unicode.IsLetter(char)){
			shift:=LetterNumber(string(rkey[ind%len(newkey)]))
			ctxt+=CaesarEncode(string(char), shift[0])
			ind++
		}else{
			ctxt+=string(char)
		}
	}
	return ctxt,nil
}

func VigenereDecode(msg string, key string) (string,error){
	newkey:=""
	for _,char:=range(key){
		if(unicode.IsLetter(char)){
			newkey+=string(char)
		}
		if(len(newkey)>=len(msg)){break}
	}
	if(len(newkey)<1){
		alErr:=errors.New("key is too short")
		log.Fatalln(alErr)
		return "", alErr
	}
	rkey:=[]rune(newkey)
	ptxt:=""
	ind:=0
	for _,char := range msg{
		if(unicode.IsLetter(char)){
			shift:=LetterNumber(string(rkey[ind%len(newkey)]))
			ptxt+=CaesarDecode(string(char), shift[0])
			ind++
		}else{
			ptxt+=string(char)
		}
	}
	return ptxt,nil
}

func RunningEncode(msg string, file string) (string,error){
	text,err:=ioutil.ReadFile(file)
    if err != nil {
        log.Fatal(err)
    }
	key:=string(text)
	if(len(key)<len(msg)){
		alErr:=errors.New("key is too short")
		log.Fatalln(alErr)
		return "", alErr
	}
	return VigenereEncode(msg, key)
}

func RunningDecode(msg string, file string) (string,error){
	text,err:=ioutil.ReadFile(file)
    if err != nil {
        log.Fatal(err)
    }
	key:=string(text)
	if(len(key)<len(msg)){
		alErr:=errors.New("key is too short")
		log.Fatalln(alErr)
		return "", alErr
	}
	return VigenereDecode(msg, key)
}
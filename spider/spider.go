package spider

import (
	"errors"
	"log"
	"strings"
	"unicode"
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
	if(shft<0){
		shft = shft+26
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
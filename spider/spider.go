package spider

import(
	"log"
	"errors"
	"strings"
)
//var atbash map[rune]rune
const plain = "abcdefghijklmnopqrstuvwxyz"
func QuickMap(alph string) (map[rune]rune,error){
	key := make(map[rune]rune)
	if(len(alph)!=26){
		alErr:=errors.New("26 rune alphabet required")
		log.Fatalln(alErr)
		return key, alErr
	}else{
		plslice := []rune(plain)
		for ind,char := range alph{
			key[plslice[ind]]=char
		}
	}
	return key, nil
}
func SubEncode(msg string, key map[rune]rune,fill bool) string{
	msg = strings.ToLower(msg)
	var cmsg []rune
	for _,char := range msg{
		cchar,exists := key[char]
		if(exists){
			cmsg = append(cmsg,cchar)
		}else if(fill){
			cmsg = append(cmsg,char)
		}else{
			cmsg = append(cmsg,'?')
		}
	}
	return string(cmsg)
}

func SubDecode(msg string, key map[rune]rune,fill bool) string{
	msg = strings.ToLower(msg)
	var pmsg []rune
	for _,char := range msg{
		var found = false
		for dec,enc :=range key{
			if enc==char{
				pmsg = append(pmsg, dec)
				found = true
				break
			}
		}
		if(!found){
			if(fill){
				pmsg = append(pmsg, char)
			}else{
				pmsg = append(pmsg, '?')
			}
		}
	}
	return string(pmsg)
}
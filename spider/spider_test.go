package spider

import (
	"testing"
)

//Utility tests
func TestFill(t *testing.T) {
	emptykey := make(map[rune]rune)
	want:="The quick brown fox jumps over the lazy dog."
	got:=SubEncode("The quick brown fox jumps over the lazy dog.",
					emptykey,true)
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestNoFill(t *testing.T) {
	emptykey := make(map[rune]rune)
	want:="??? ????? ????? ??? ????? ???? ??? ???? ???."
	got:=SubEncode("The quick brown fox jumps over the lazy dog.",
					emptykey,false)
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

//Generic sub tests
func TestEncode(t *testing.T) {
	key,_ := QuickMap("mlkjictsrqpobazyxwhgfedvun")
	want:="Gsi xfrkp lwzda czv qfbyh zeiw gsi omnu jzt."
	got:=SubEncode("The quick brown fox jumps over the lazy dog.",
					key,false)
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestDecode(t *testing.T) {
	key,_ := QuickMap("mlkjictsrqpobazyxwhgfedvun")
	want:="The quick brown fox jumps over the lazy dog."
	got:=SubDecode("Gsi xfrkp lwzda czv qfbyh zeiw gsi omnu jzt.",
					key,false)
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestEncodeDecode(t *testing.T) {
	key,_ := QuickMap("mlkjictsrqpobazyxwhgfedvun")
	want:="The quick brown fox jumps over the lazy dog."
	got:=SubDecode(
		SubEncode("The quick brown fox jumps over the lazy dog.",
					key,false),
		key,false)
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

//Caesar tests
func TestCaesarEncode(t *testing.T) {
	want:="Ymj vznhp gwtbs ktc ozrux tajw ymj qfed itl."
	got:=CaesarEncode("The quick brown fox jumps over the lazy dog.",5)
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestCaesarDecode(t *testing.T) {
	want:="The quick brown fox jumps over the lazy dog."
	got:=CaesarDecode("Ymj vznhp gwtbs ktc ozrux tajw ymj qfed itl.",5)
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestCaesarEncodeDecode(t *testing.T) {
	want:="The quick brown fox jumps over the lazy dog."
	got:=CaesarDecode(CaesarEncode("The quick brown fox jumps over the lazy dog.",5),5)
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

//Rot13 tests
func TestRot13Encode(t *testing.T) {
	want:="Gur dhvpx oebja sbk whzcf bire gur ynml qbt."
	got:=Rot13("The quick brown fox jumps over the lazy dog.")
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestRot13EncodeDecode(t *testing.T) {
	want:="The quick brown fox jumps over the lazy dog."
	got:=Rot13(Rot13("The quick brown fox jumps over the lazy dog."))
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

//Atbash tests
func TestAtbashEncode(t *testing.T) {
	want:="Gsv jfrxp yildm ulc qfnkh levi gsv ozab wlt."
	got:=Atbash("The quick brown fox jumps over the lazy dog.")
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestAtbashEncodeDecode(t *testing.T) {
	want:="The quick brown fox jumps over the lazy dog."
	got:=Atbash(Atbash("The quick brown fox jumps over the lazy dog."))
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

//LetterNumber tests
func TestLetterNumber(t *testing.T) {
	want:=[]int{19,7,4,16,20,8,2,10,1,17,14,22,13,5,14,23,9,20,12,15,18,14,21,4,17,19,7,4,11,0,25,24,3,14,6}
	got:=LetterNumber("The quick brown fox jumps over the lazy dog.")
	for ind,num:=range want{
		if(num!=got[ind]){
			t.Errorf("got %d, wanted %d", got[ind], want[ind])
		}
	}
}
func TestNumberLetter(t *testing.T) {
	want:="thequickbrownfoxjumpsoverthelazydog"
	got:=NumberLetter([]int{19,7,4,16,20,8,2,10,1,17,14,22,13,5,14,23,9,20,12,15,18,14,21,4,17,19,7,4,11,0,25,24,3,14,6})
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestLetterNumberLetter(t *testing.T) {
	want:="thequickbrownfoxjumpsoverthelazydog"
	got:=NumberLetter(LetterNumber(want))
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestNumberLetterNumber(t *testing.T) {
	want:=[]int{19,7,4,16,20,8,2,10,1,17,14,22,13,5,14,23,9,20,12,15,18,14,21,4,17,19,7,4,11,0,25,24,3,14,6}
	got:=LetterNumber(NumberLetter(want))
	for ind,num:=range want{
		if(num!=got[ind]){
			t.Errorf("got %d, wanted %d", got[ind], want[ind])
		}
	}
}

//Vigenere tests
func TestVigenereEncode(t *testing.T) {
	want:="Vho uwimo dryap fyb luwtu ofit tri najc foq."
	got,_:=VigenereEncode("The quick brown fox jumps over the lazy dog.","c a  k e\n")
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestVigenereDecode(t *testing.T) {
	want:="The quick brown fox jumps over the lazy dog."
	got,_:=VigenereDecode("Vho uwimo dryap fyb luwtu ofit tri najc foq.","cake")
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestVigenereEncodeDecode(t *testing.T) {
	want:="The quick brown fox jumps over the lazy dog."
	got1,_:=VigenereEncode(want,"cake")
	got,_:=VigenereDecode(got1,"cake")
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
//Running key tests
func TestRunningEncode(t *testing.T) {
	want:="Lvq uvwfi peqag tza vyfww kjvc wpw roml dfu."
	got,_:=RunningEncode("The quick brown fox jumps over the lazy dog.","AllStar.txt")
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestRunningDecode(t *testing.T) {
	want:="The quick brown fox jumps over the lazy dog."
	got,_:=RunningDecode("Lvq uvwfi peqag tza vyfww kjvc wpw roml dfu.","AllStar.txt")
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
func TestRunningEncodeDecode(t *testing.T) {
	want:="The quick brown fox jumps over the lazy dog."
	got1,_:=RunningEncode(want,"AllStar.txt")
	got,_:=RunningDecode(got1,"AllStar.txt")
	if (got != want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
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

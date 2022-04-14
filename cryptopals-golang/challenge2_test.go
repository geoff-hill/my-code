package cryptopals
     
import (
	"testing"
)

func Test_Challenge2(t *testing.T) {
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"

	ret, err := XORHexStrings(hex1, hex2)
	if err != nil {
		t.Log("An error was returned", err)
		t.Fail()
	} else if ret != expected {
		t.Log("Expected ", expected)
		t.Log("Returned ", ret)
		t.Fail()
	}
}
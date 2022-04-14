package cryptopals
     
import (
	"testing"
)

func TestChallenge1(t *testing.T) {
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	ret, err := ConvertHexToBase64(hex)
	if err != nil {
		t.Log("An error was returned", err)
		t.Fail()
	} else if ret != expected {
		t.Log("Expected ", expected)
		t.Log("Returned ", ret)
		t.Fail()
	}
}
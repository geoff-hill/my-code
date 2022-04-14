package cryptopals
     
import (
	"testing"
)

func TestChallenge3(t *testing.T) {
	cypherTextHex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expected := "Cooking MC's like a pound of bacon"
	_, clearTextHex, _, err := CrackSingleByteXORStringByFrequency(cypherTextHex)
	if err != nil {
		t.Log("An error was returned", err)
		t.Fail()
	}
	if clearTextHex != expected {
		t.Log("Expected", expected)
		t.Log("Extracted", clearTextHex)
		t.Fail()
	}

}

func TestChallenge3_UsingMostLetters(t *testing.T) {
	cypherTextHex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expected := "Cooking MC's like a pound of bacon"
	_, clearTextHex, _, err := CrackSingleByteXORStringByMostLetters(cypherTextHex)
	if err != nil {
		t.Log("An error was returned", err)
		t.Fail()
	}
	if clearTextHex != expected {
		t.Log("Expected", expected)
		t.Log("Extracted", clearTextHex)
		t.Fail()
	}

}
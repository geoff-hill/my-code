package cryptopals

import (
	"strings"
	"testing"
)


func TestChallenge4(t *testing.T) {
	challenge4Url := "https://cryptopals.com/static/challenge-data/4.txt"
	expectedClearText := "Now that the party is jumping"
		line, key, cipherTextHex, clearText, score, err := IdentifySingleByteXOR(challenge4Url)
	if err != nil {
		t.Log("An error was returned", err)
		t.Fail()
	}
	t.Log("line ", line)
	t.Log("key ", key)
	t.Log("cipherTextHex ", cipherTextHex)
	t.Log("clearText ", clearText)
	t.Log("score ", score)
	if strings.TrimSpace(clearText) != expectedClearText {
		t.Log("The wrong line has been identified")
		t.Log("Identified", "**"+clearText+"**")
		t.Log("Expected", "**"+expectedClearText+"**")
		t.Fail()
	}
}
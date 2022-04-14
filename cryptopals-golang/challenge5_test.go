package cryptopals

import (
	"encoding/hex"
	"fmt"
	"testing"
)


func TestChallenge5_RepeatBug(t *testing.T) {
	key := []byte("123");
	expected := "12312312312312312312"

	result := string(RepeatBuffer(key, 20))

	fmt.Println(result)
	if expected != result {
		t.Log("Returned", result)
		t.Log("expected", expected)
		t.Fail()
	}
}

func TestChallenge5(t *testing.T) {
	clearText := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	key := "ICE"
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	result, err := MultiByteXOR([]byte(key), []byte(clearText))
	if err != nil {
		t.Log("An error was returned", err)
		t.FailNow()
	}
	hexResult := hex.EncodeToString(result)
	if expected != hexResult {
		t.Log("Returned", hexResult)
		t.Log("expected", expected)
		t.Fail()
	}
}
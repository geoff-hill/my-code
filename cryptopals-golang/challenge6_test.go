package cryptopals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChallenge6_HammingDistanceStrings(t *testing.T) {
	thisIsATest := "this is a test"
	wokkaWokka := "wokka wokka!!!"
	expectedHammingDistance := 37

	returnedHammingDistance := HammingDistanceStrings(thisIsATest, wokkaWokka)
	t.Log("Expected:", expectedHammingDistance, ", Returned:", returnedHammingDistance)
	if expectedHammingDistance != returnedHammingDistance {
		t.Log("Expected:", expectedHammingDistance, ", Returned:", returnedHammingDistance)
		t.FailNow()
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func arraysEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestChallenge6_SortedKeys(t *testing.T) {
	m := make(map[int]float64)
	m[1] = 3.14159
	m[2] = 2
	m[20] = 0.02
	m[3] = 1.9999

	sorted := SortedKeys(m)
	expected := []int{20, 3, 2, 1}

	if !arraysEqual(expected, sorted) {
		t.Log("Expected:", expected, ", Returned:", sorted)
		t.FailNow()
	}
}

func TestChallenge6_HammingDistanceByteArrays(t *testing.T) {
	thisIsATest := []byte("this is a test")
	wokkaWokka := []byte("wokka wokka!!!")
	expectedEditDistance := 37

	returnedEditDistance := HammingDistanceByteArrays(thisIsATest, wokkaWokka)
	t.Log("Expected:", expectedEditDistance, ", Returned:", returnedEditDistance)
	if expectedEditDistance != returnedEditDistance {
		t.Log("Expected:", expectedEditDistance, ", Returned:", returnedEditDistance)
		t.FailNow()
	}
}

func TestChallenge6_NormalisedEditDistance1(t *testing.T) {
	thisIsATest := []byte("this is a test")
	wokkaWokka := []byte("wokka wokka!!!")
	expectedEditDistance := 37.00

	returnedEditDistance := NormalisedEditDistance(1, thisIsATest, wokkaWokka)
	t.Log("Expected:", expectedEditDistance, ", Returned:", returnedEditDistance)
	if expectedEditDistance != returnedEditDistance {
		t.Log("Expected:", expectedEditDistance, ", Returned:", returnedEditDistance)
		t.FailNow()
	}
}

func TestChallenge6_NormalisedEditDistance10(t *testing.T) {
	thisIsATest := []byte("this is a test")
	wokkaWokka := []byte("wokka wokka!!!")
	expectedEditDistance := 37.0 / 10.0

	returnedEditDistance := NormalisedEditDistance(10, thisIsATest, wokkaWokka)
	t.Log("Expected:", expectedEditDistance, ", Returned:", returnedEditDistance)
	if expectedEditDistance != returnedEditDistance {
		t.Log("Expected:", expectedEditDistance, ", Returned:", returnedEditDistance)
		t.FailNow()
	}
}

func assertResult(t *testing.T, expected, result string) {
	if expected != result {
		t.Log("FAIL: Expected", expected, "got", result)
		t.Fail()
	} else {
		t.Log("Expected", expected, "got", result, "good")
	}
}

func TestChallenge6_Transpose(t *testing.T) {
    s1 := []byte("145")
	s2 := []byte("BCF")
	s3 := []byte("678")
	s4 := []byte("1")

	testData := [][]byte { s1, s2, s3, s4 }
	result := TransposeBytes(testData)
	expectedLen := 3
	if len(result) != expectedLen {
		t.Log("Expected", expectedLen, "results", len(result))
		t.FailNow()
	}
	assert.EqualValues(t, "1B61", result[0])
	assert.EqualValues(t, "4C7", result[1])
	assert.EqualValues(t, "5F8", result[2])
}

func TestChallenge6_BreakIntoBlocks(t *testing.T) {
	thisIsATest := []byte("123456789ABCDEF123456789ABCDEF1")
    testBlockSize := 10
	expectedBlocks := 4
	result := BreakIntoBlocks(thisIsATest, testBlockSize)

	t.Log("Expected", expectedBlocks, "results", len(result))
	if len(result) != expectedBlocks {
		t.Log("Expected", expectedBlocks, "results", len(result))
		t.Fail()
	} 
	assertResult(t, "123456789A", string(result[0]))
	assertResult(t, "BCDEF12345", string(result[1]))
	assertResult(t, "6789ABCDEF", string(result[2]))
	assertResult(t, "1", string(result[3]))

}

func TestChallenge6_BreakIntoBlocksEndsWithPartialBuffer(t *testing.T) {
	thisIsATest := []byte("123456789ABCDEF123456789ABCDEF1")
    testBlockSize := 2
	expectedBlocks := 16
	result := BreakIntoBlocks(thisIsATest, testBlockSize)

	t.Log("Expected", expectedBlocks, "results", len(result))
	if len(result) != expectedBlocks {
		t.Log("Expected", expectedBlocks, "results", len(result))
		t.Fail()
	} 
	assertResult(t, "12", string(result[0]))
	assertResult(t, "34", string(result[1]))
	assertResult(t, "56", string(result[2]))
	assertResult(t, "78", string(result[3]))
	assertResult(t, "9A", string(result[4]))
	assertResult(t, "BC", string(result[5]))
	assertResult(t, "DE", string(result[6]))
	assertResult(t, "F1", string(result[7]))
	assertResult(t, "23", string(result[8]))
	assertResult(t, "45", string(result[9]))
	assertResult(t, "67", string(result[10]))
	assertResult(t, "89", string(result[11]))
	assertResult(t, "AB", string(result[12]))
	assertResult(t, "CD", string(result[13]))
	assertResult(t, "EF", string(result[14]))
	assertResult(t, "1", string(result[15]))
}

func TestChallenge6_BreakIntoBlocksExactNumberOfBuffers(t *testing.T) {
	thisIsATest := []byte("123456789ABCDEF123456789ABCDEF")
    testBlockSize := 2
	expectedBlocks := 15
	result := BreakIntoBlocks(thisIsATest, testBlockSize)

	t.Log("Expected", expectedBlocks, "results", len(result))
	if len(result) != expectedBlocks {
		t.Log("Expected", expectedBlocks, "results", len(result))
		t.Fail()
	} 
	assertResult(t, "12", string(result[0]))
	assertResult(t, "34", string(result[1]))
	assertResult(t, "56", string(result[2]))
	assertResult(t, "78", string(result[3]))
	assertResult(t, "9A", string(result[4]))
	assertResult(t, "BC", string(result[5]))
	assertResult(t, "DE", string(result[6]))
	assertResult(t, "F1", string(result[7]))
	assertResult(t, "23", string(result[8]))
	assertResult(t, "45", string(result[9]))
	assertResult(t, "67", string(result[10]))
	assertResult(t, "89", string(result[11]))
	assertResult(t, "AB", string(result[12]))
	assertResult(t, "CD", string(result[13]))
	assertResult(t, "EF", string(result[14]))
}

func TestChallenge6_BreakRepeatingKeyXORGivenKeySize(t *testing.T) {
	clearText := "Hello John have you got a new motor, to be or not to be that is the question?"
	testKey := "golang"
 	var err error
	testData, err := MultiByteXOR([]byte(testKey), []byte(clearText))
	if err != nil {
		t.Fatal("MultiByteXOR returns an error", err)
	}
	resultKey, err := BreakRepeatingKeyXORGivenKeySize(testData, len(testKey))
	if err != nil {
		t.Fatal("BreakRepeatingKeyXORGivenKeySize returns an error", err)
	}

	if testKey != string(resultKey) {
		t.Fatal("Expected", testKey, "results", string(resultKey))
	} 

	decoded, err := MultiByteXOR(resultKey, testData)

	if err != nil {
		t.Fatal("Decode with MultiByteXOR returns an error", err)
	}
	if clearText == string(decoded) {
		t.Fatal("Expected", clearText, "results", string(decoded))
	}
}

func TestChallenge6_BreakRepeatingKeyXOR(t *testing.T) {
	clearText := "Hello John have you got a new motor, to be or not to be that is the question?"
	testKey := "outcast"
 	var err error
	testData, err := MultiByteXOR([]byte(testKey), []byte(clearText))
	if err != nil {
		t.Fatal("MultiByteXOR returns an error", err)
	}
	resultKey, _, err := BreakRepeatingKeyXOR(testData)
	if err != nil {
		t.Fatal("BreakRepeatingKeyXOR returns an error", err)
	}
	if testKey != string(resultKey) {
		t.Fatal("Expected", testKey, "results", string(resultKey))
	}
	// if clearText != string(resultClearText) {
	// 	t.Fatal("Expected", clearText, "results", string(resultClearText))
	// } 
}


func TestChallenge6(t *testing.T) {
	url := "https://cryptopals.com/static/challenge-data/6.txt"
	expectedKey := "12345"
	expectedClearText := "Whatever"

	key, clearText, err := BreakRepeatingKeyXORBase64FromURL(url)
	if err != nil {
		t.Fatal("An error was returned", err)
	}
	if expectedKey != key {

		t.Log("Expected:", expectedKey)
		t.Log("Returned:", key)

		t.Error("Incorrect Key Selected")
	}
	if expectedClearText != clearText {
		t.Log("Expected:", expectedClearText)
		// t.Log("Returned:", clearText)

		t.Error("Incorrect decoding")
	}
}

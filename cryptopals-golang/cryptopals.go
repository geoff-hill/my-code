package cryptopals

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"

	"github.com/steakknife/hamming"
)

func ConvertHexToBase64(fromHex string) (result string, err error) {
	decoded, err := hex.DecodeString(fromHex)
	if err != nil {
		return
	}

	return base64.RawStdEncoding.EncodeToString(decoded), err
}

func XOR(buf1 []byte, buf2 []byte) (ret []byte, err error) {
	len1 := len(buf1)
	if len1 != len(buf2) {
		err = fmt.Errorf("XOR buffers have different sizes: %d vs %d", len1, len(buf2))
		return
	}

	ret = make([]byte, len1)
	for i := range buf1 {
		ret[i] = buf1[i] ^ buf2[i]
	}
	return
}

func RepeatBuffer(buf []byte, length int) []byte {
	repeats := int(math.Ceil(float64(length) / float64(len(buf))))
	newBuf := bytes.Repeat(buf, repeats)
	return newBuf[:length]
}

func MultiByteXOR(key []byte, buf []byte) ([]byte, error) {
	keyBuf := RepeatBuffer(key, len(buf))
	return XOR(keyBuf, buf)
}

func SingleByteXOR(key byte, buf []byte) (ret []byte) {
	ret = make([]byte, len(buf))
	for i := range buf {
		ret[i] = buf[i] ^ key
	}
	return
}

func XORHexStrings(hex1, hex2 string) (result string, err error) {
	decoded1, err := hex.DecodeString(hex1)
	if err != nil {
		return
	}
	decoded2, err := hex.DecodeString(hex2)
	if err != nil {
		return
	}

	ret, err := XOR(decoded1, decoded2)
	if err != nil {
		return
	}

	result = hex.EncodeToString(ret)
	return
}

func UrlToLines(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return LinesFromReader(resp.Body)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func isLetter(b byte) bool {
	return (b > 64 && b < 91) || (b > 96 && b < 123)
}

func isSpace(b byte) bool {
	return b == 32
}

func countLettersAndSpaces(clearText []byte) (count int) {
	for _, b := range clearText {
		if isLetter(b) || isSpace(b) {
			count++
		}
	}
	return
}

func scoreByMostLetters(clearText []byte) float64 {
	return float64 (countLettersAndSpaces(clearText))
}

var scoresByLetter = map[byte]float64 {
	'a': 0.0651738,
	'b': 0.0124248,
	'c': 0.0217339,
	'd': 0.0349835,
	'e': 0.1041442,
	'f': 0.0197881,
	'g': 0.0158610,
	'h': 0.0492888,
	'i': 0.0558094,
	'j': 0.0009033,
	'k': 0.0050529,
	'l': 0.0331490,
	'm': 0.0202124,
	'n': 0.0564513,
	'o': 0.0596302,
	'p': 0.0137645,
	'q': 0.0008606,
	'r': 0.0497563,
	's': 0.0515760,
	't': 0.0729357,
	'u': 0.0225134,
	'v': 0.0082903,
	'w': 0.0171272,
	'x': 0.0013692,
	'y': 0.0145984,
	'z': 0.0007836,
	' ': 0.1918182,
}

func scoreByFrequency(clearText []byte) (float64) {
	var score float64
	for _, b := range []byte(strings.ToLower(string(clearText))) {
		score += scoresByLetter[b]
	}
	return score / float64(len(clearText))
}

func createMapOfCounts(clearText []byte) map[byte] int {
	ret := make(map[byte]int)
	for _, b := range []byte(strings.ToLower(string(clearText))) {
		if b == ' ' || (b >= 'a' && b <= 'z') {
			ret[b] = ret[b] + 1
		}
	}
	return ret
}

func sumOfCounts(mapOfCounts map[byte]int) int { 
	tot := 0
	for _, count := range mapOfCounts {
		tot += count
	}
	return tot 
}

func scaleCountsToFrequencies(mapOfCounts map[byte]int) map[byte]float64 {
	tot := float64(sumOfCounts(mapOfCounts))
	ret := make(map[byte]float64)
	for key, val := range mapOfCounts {
		ret[key] = float64(val) / tot
	}
	return ret
}

func scoreByChiSquaredGOF(clearText []byte) float64 {
    clearTextFrequency := scaleCountsToFrequencies(createMapOfCounts(clearText))
	var sum float64	
	for key, val := range scoresByLetter {
		sum += (((clearTextFrequency[key] - val) * (clearTextFrequency[key] - val)) / val)
	}
	return -1.000 * sum
}

func defaultScorer(clearText []byte) (score float64) {
	return scoreByFrequency(clearText)
}


func crackSingleByteXOR(cipherText []byte, scorer func([]byte) float64) (key byte, clearText []byte, score float64, err error) {

	first := true

	for thisKey := byte(0); thisKey <= 127; thisKey++ {
		thisClearText := SingleByteXOR(thisKey, cipherText)
		thisScore := scorer(thisClearText)
		if first || thisScore > score {
			// fmt.Println("A new leader", thisKey, thisScore, string(thisClearText))
			first = false
			score = thisScore
			key = thisKey
			clearText = thisClearText
		}

	}

	return
}

func crackSingleByteXORString(cipherTextHex string, scorer func([]byte) float64) (key byte, clearText string, score float64, err error) {
	cipherText, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return
	}

	key, clearTextBytes, score, err := crackSingleByteXOR(cipherText, scorer)
	if err != nil {
		return
	}

	clearText = string(clearTextBytes)
	return
}

func CrackSingleByteXORStringByMostLetters(cipherTextHex string) (key byte, clearText string, score float64, err error) {
	return crackSingleByteXORString(cipherTextHex, scoreByMostLetters)
}

func CrackSingleByteXORStringByFrequency(cipherTextHex string) (key byte, clearText string, score float64, err error) {
	return crackSingleByteXORString(cipherTextHex, scoreByFrequency)
}


func identifySingleByteXORInternal(linesHex []string) (line int, key byte, cipherTextHex, clearText string, score float64, err error) {

	for ix, lineHex := range linesHex {
		var thisKey byte
		var thisScore float64
		var thisClearText string

		thisKey, thisClearText, thisScore, err = CrackSingleByteXORStringByFrequency(lineHex)
		if err != nil {
			return
		}
		if thisScore > score {
			score = thisScore
			line = ix
			key = thisKey
			cipherTextHex = lineHex
			clearText = thisClearText
		}
	}
	return
}

func IdentifySingleByteXOR(url string) (line int, key byte, cipherTextHex, clearText string, score float64, err error) {
	lines, err := UrlToLines(url)
	if err != nil {
		return
	}

	return identifySingleByteXORInternal(lines)
}

func HammingDistanceStrings(s1, s2 string) int {
	return hamming.Strings(s1, s2)
}

func HammingDistanceByteArrays(b1, b2 []byte) int {
	return hamming.Bytes(b1, b2)
} 

func getCandidateKeySizes(min, max, candidates int, cipherText[]byte) []int {

	scoreMap := make(map[int]float64)
	for i := min; i <= max && i*2 < len(cipherText); i++ {
		score := NormalisedEditDistance(i, cipherText[0:i], cipherText[i:2*i])
		scoreMap[i] = score
	}
	return SortedKeys(scoreMap)[0:candidates]
}

func NormalisedEditDistance(keySize int, buf1 []byte, buf2 []byte) float64 {
	return float64(HammingDistanceByteArrays(buf1, buf2)) / float64(keySize)
}

func BreakIntoBlocks(buf []byte, blockSize int) [][]byte {
	numBlocks := int(math.Ceil(float64(len(buf)) / float64(blockSize)))
	bufLen := len(buf)
	blocks := make([][]byte, numBlocks)
	for i := 0; i < numBlocks; i++ {
		for j := 0; j < blockSize; j++ {
			bufIx := blockSize*i + j
			if bufIx < bufLen {
				blocks[i] = append(blocks[i], buf[bufIx])
			}
		}
	}
	return blocks;																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																							
}


func BreakRepeatingKeyXORGivenKeySize(cipherText []byte, keySize int) (possibleKey []byte, err error) {
	blocks := BreakIntoBlocks(cipherText, keySize)
	transposed := TransposeBytes(blocks)
	possibleKey = make([]byte, keySize)

	var keyByte byte
	for i, buf := range transposed {
		keyByte, _, _, err = crackSingleByteXOR(buf, defaultScorer)
		if err != nil {
			return
		}
		possibleKey[i] = keyByte 
	}

    return
}

func BreakRepeatingKeyXOR(cipherText []byte) (key []byte, result []byte, err error) {
	numberOfCandidates := 5
	candidateKeySizes := getCandidateKeySizes(2, 40, numberOfCandidates, cipherText)
	fmt.Println("CKS", candidateKeySizes)
	candidateKeys := make([][]byte, numberOfCandidates)
	for i, keySize := range candidateKeySizes {
		candidateKeys[i], err = BreakRepeatingKeyXORGivenKeySize(cipherText, keySize)
		if err != nil {
			return
		}
	}
	var resultScore = float64(0.0)
    
	fmt.Println("Candidate Keys", candidateKeys)

	for _, candidateKey := range candidateKeys {
		
		var clear []byte
		clear, err = MultiByteXOR(candidateKey, cipherText)
		if err != nil {
			return
		}
		score := defaultScorer(clear)
		if score > resultScore {
			fmt.Println("Adding", candidateKey, "score", score)
			resultScore = score
			result = clear
			key = candidateKey
		}
	}
	return
}

func BreakRepeatingKeyXORBase64FromURL(url string) (key string, clearText string, err error) {
	lines, _ := UrlToLines(url)
	joined := strings.Join(lines, "")
	cipherText, err := base64.StdEncoding.DecodeString(joined)
	if err != nil {
		return 
	}
	keyBytes, clearBytes, err := BreakRepeatingKeyXOR(cipherText)
	if err != nil {
		return 
	}
	return string(keyBytes), string(clearBytes), nil
}

func TransposeBytes(matrix [][]byte) [][]byte {
		m := len(matrix)
		n := len(matrix[0])
		// transposed matrix
		T := make([][]byte, n)
		for i := 0; i < n; i++ {
			T[i] = make([]byte, m)
		}
	
		for j := 0; j < m; j++ {
			for i := 0; i < n; i++ {
				if i < len(matrix[j]) {
					T[i][j] = matrix[j][i]
				}
			}
		}

		for k, buf := range T {
			T[k] = bytes.Trim(buf, "\x00")
		}
		
		return T
}

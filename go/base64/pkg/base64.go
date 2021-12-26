package pkg

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"regexp"
)

var m = map[string]string{
	"000000": "A", "000001": "B", "000010": "C", "000011": "D",
	"000100": "E", "000101": "F", "000110": "G", "000111": "H",
	"001000": "I", "001001": "J", "001010": "K", "001011": "L",
	"001100": "M", "001101": "N", "001110": "O", "001111": "P",
	"010000": "Q", "010001": "R", "010010": "S", "010011": "T",
	"010100": "U", "010101": "V", "010110": "W", "010111": "X",
	"011000": "Y", "011001": "Z",

	"011010": "a", "011011": "b", "011100": "c", "011101": "d",
	"011110": "e", "011111": "f", "100000": "g", "100001": "h",
	"100010": "i", "100011": "j", "100100": "k", "100101": "l",
	"100110": "m", "100111": "n", "101000": "o", "101001": "p",
	"101010": "q", "101011": "r", "101100": "s", "101101": "t",
	"101110": "u", "101111": "v", "110000": "w", "110001": "x",
	"110010": "y", "110011": "z",

	"110100": "0", "110101": "1", "110110": "2", "110111": "3",
	"111000": "4", "111001": "5", "111010": "6", "111011": "7",
	"111100": "8", "111101": "9", "111110": "+", "111111": "/",
}

func strToBin(s string) []uint8 {
	var buffer bytes.Buffer

	for _, c := range s {
		fmt.Fprintf(&buffer, "%08b", c)
	}

	return buffer.Bytes()
}

func padding(s string) string {
	n := len(s)

	p := 6 - n

	for p > 0 {
		s = s + "0"
		p--
	}

	return s
}

// Encode string to base64
func Encode(s string) string {
	b8 := strToBin(s)

	b64 := ""

	for len(b8) > 6 {
		c, _ := m[string(b8[:6])]
		b64 = b64 + c

		b8 = b8[6:]
	}

	if len(b8) > 0 {
		t, _ := m[padding(string(b8))]
		b64 = b64 + t
	}

	p := 4 - (len(b64) % 4)
	for p > 0 {
		b64 = b64 + "="
		p--
	}

	return string(b64)
}

func base64ToBin(s string) string {
	b := ""

	for _, c := range s {
		// TODO: build reverse map will be quik
		for k, v := range m {
			if string(c) == v {
				b = b + k
				break
			}
		}
	}

	return b
}

// Decode base64 to string
func Decode(s string) (string, error) {
	if len(s)%4 != 0 {
		return "", errors.New("Wrong Length")
	}

    match, _ := regexp.MatchString("^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{4}|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)$", s)
	if !match {
		return "", errors.New("Wrong Regex")
	}

	b := base64ToBin(s)
	str := ""
	for i := 8; i < len(b); i += 8 {
		c, _ := strconv.ParseInt(b[i-8:i], 2, 8)
		str = str + string(c)
	}

	return str, nil
}

package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode/utf8"
)

var DATASTR = "010101110110000101110011001" +
	"0000001101000011000010111" +
	"0100001000000111011001101" +
	"0010110010101110010001000" +
	"0001000010011001010110100" +
	"10110111001100101001000000" +
	"111010101101110011001000010" +
	"00000110101101100001011011" +
	"10011011100010000001100110" +
	"011011000110100101100101011" +
	"0011101100101011011100011111" +
	"100100000010110100111011101" +
	"1001010110100100100000010" +
	"1011011000011101101100110011" +
	"1011001010110110000100001"

func toByte(s string) byte {
	i, _ := strconv.ParseUint(s, 2, 8)
	return byte(i)
}

func main() {
	if len(DATASTR)%8 != 0 {
		log.Fatal(fmt.Sprintf("Fehler beim Abtippen: Text hat nicht die erwartete Länge (teilbar durch 8):\n\tRest: %v\n\tErwartet %v\n", len(DATASTR)%8, 0))
	}

	var strBytes []byte
	for i := 0; i < len(DATASTR)/8; i++ {
		strBytes = append(strBytes, toByte(DATASTR[i*8:(i+1)*8]))
		if utf8.Valid(strBytes) {
			fmt.Printf("%s\n", string(strBytes))
		}
	}

	fmt.Println(";-) Haha")
}

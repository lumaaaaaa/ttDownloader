package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	HexKey = []int64{0xDF, 0x77, 0xB9, 0x40, 0xb9, 0x9b, 0x84, 0x83, 0xd1, 0xb9, 0xcb, 0xd1, 0xf7, 0xc2, 0xb9, 0x85, 0xc3, 0xd0, 0xfb, 0xc3}
)

func generateMD5(str string) string {
	var hash [16]byte = md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func hexString(n int64) string {
	var HEX = strconv.FormatInt(n, 16)
	if len(HEX) < 2 {
		HEX = "0" + HEX
	}
	return HEX
}

func RBIT(num int64) int64 {
	var binData string
	var bin = fmt.Sprintf("%b", num)

	for len(bin) < 8 {
		bin = "0" + bin
	}
	for i := 0; i < 8; i++ {
		binData += string(bin[7-i])
	}
	var data, _ = strconv.ParseInt(binData, 2, 64)
	return data
}

func reverseString(n int64) int64 {
	var hexStr string = hexString(n)
	var hexInt, _ = strconv.ParseInt(hexStr[1:]+hexStr[:1], 16, 64)
	return hexInt
}

func generateSignature(params string) (string, int64) {
	var MD5 string = generateMD5(params) + strings.Repeat("00000000000000000000000000000000", 3)
	var HEX []int64
	var EOR []int64

	var xGorgon string
	var xKhronos int64 = time.Now().Unix()

	for i := 0; i < 12; i += 4 {
		for x := 0; x < 4; x++ {
			var H, _ = strconv.ParseInt(MD5[8*i : 8*(i+1)][x*2:(x+1)*2], 16, 64)
			HEX = append(HEX, H)
		}
	}
	HEX = append(HEX, 0x0, 0x6, 0xB, 0x1C, (xKhronos&0xFF000000)>>24, (xKhronos&0x00FF0000)>>16, (xKhronos&0x0000FF00)>>8, (xKhronos&0x000000FF)>>0)

	for i := range HexKey {
		EOR = append(EOR, HEX[i]^HexKey[i])
	}
	for i := 0; i < 0x14; i++ {
		EOR[i] = ((RBIT(reverseString(EOR[i])^EOR[(i+1)%0x14]) ^ 0xFFFFFFFF) ^ 0x14) & 0xFF
	}
	for _, value := range EOR {
		xGorgon += hexString(value)
	}
	return "0408b0d30000" + xGorgon, xKhronos
}

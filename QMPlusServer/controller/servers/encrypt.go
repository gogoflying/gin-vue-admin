package servers

import (
	"strconv"
)

var EnKey []byte = []byte{0x0B, 0x0A, 0x8C, 0x0C, 0x93, 0x6D, 0x46, 0x47}

func DataEncryPt(eData string) string {
	var result string
	j := 0
	s := ""
	bt := []rune(eData)
	for i := 0; i < len(bt); i++ {
		s = strconv.FormatInt(int64(byte(bt[i])^EnKey[j]), 16)
		if len(s) == 1 {
			s = "0" + s
		}
		result = result + (s)
		j = (j + 1) % 8
	}
	return result
}

func DataDecode(eData string) string {
	var result string
	var s int64
	j := 0
	bt := []rune(eData)
	//fmt.Println(bt)
	for i := 0; i < len(eData)/2; i++ {
		s, _ = strconv.ParseInt(string(bt[i*2:i*2+2]), 16, 0)
		result = result + string(byte(s)^EnKey[j])
		j = (j + 1) % 8
	}
	return result
}

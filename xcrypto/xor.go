package xcrypto

import (
	"bytes"
	"encoding/base64"
)

func xor(key []byte, src []byte) []byte {
	l := len(src)
	if len(key) > l {
		l = len(key)
	}
	var ret []byte
	for i := 0; i < l; i++ {
		srcByte := byte(0)
		keyByte := byte(0)
		if i < len(src) {
			srcByte = src[i]
		}
		if i < len(key) {
			keyByte = key[i]
		}
		ret = append(ret, srcByte^keyByte)
	}
	return ret
}

func XorEncode(key string, src string) string {
	dst := xor([]byte(key), []byte(src))
	but := bytes.NewBuffer(nil)
	en := base64.NewEncoder(base64.StdEncoding, but)
	en.Write(dst)
	en.Close()
	return but.String()
}

func XorDecode(key string, src string) string {
	dbuf := make([]byte, base64.StdEncoding.DecodedLen(len(src)))
	de := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(src))
	de.Read(dbuf)
	return string(xor([]byte(key), dbuf))
}
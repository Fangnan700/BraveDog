package utils

// XorEncrypt XOR 加密函数
func XorEncrypt(data, key []byte) []byte {
	encrypted := make([]byte, len(data))
	for i := range data {
		encrypted[i] = data[i] ^ key[i%len(key)]
	}
	return encrypted
}

// XorDecrypt XOR 解密函数
func XorDecrypt(encrypted, key []byte) []byte {
	decrypted := make([]byte, len(encrypted))
	for i := range encrypted {
		decrypted[i] = encrypted[i] ^ key[i%len(key)]
	}
	return decrypted
}

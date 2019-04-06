package golangexamples
func Encrypt(sliceToEncrypt[] byte, ceaserCount int) string {
	for i:=0; i<len(sliceToEncrypt); i++ {
		sliceToEncrypt[i] = byte(int(sliceToEncrypt[i]) + ceaserCount)
	}
	
	return string(sliceToEncrypt)
}


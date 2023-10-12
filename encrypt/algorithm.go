package encrypt

func Nimbus(str string) string {
	decryptedstr := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode + 3)
		decryptedstr += character
	}
	return decryptedstr
}

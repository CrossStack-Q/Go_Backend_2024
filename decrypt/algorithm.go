package decrypt

func Nimbus(str string) string {
	encryptedstr := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode - 3)
		encryptedstr += character
	}
	return encryptedstr
}

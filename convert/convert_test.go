package convert

import "testing"

func TestConvertFileToDocx(t *testing.T) {
	ok, err := ConvertFileToDocx(
		"D:\\ChopperBot\\chopperbot-test\\src\\test\\java\\org\\example\\ws\\WebSocketTest.java",
		"D:\\ChopperBot_docx\\chopperbot-test\\src\\test\\java\\org\\example\\ws\\WebSocketTest.docx",
	)
	t.Log(ok, err)
}

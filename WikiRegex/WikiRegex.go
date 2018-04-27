package VCWikiRegex

import (
	"strings"
	"regexp"
	"log"
)

//각 타입 별 regex 태그가 전역변수로 있어야 한다.

func Save(s string){
	//client에서 받은 text를 html 파일로 바꿔서 저장하는 함수.

}

func ParseToWiki(s []string) []string{
	//큰 하나의 string으로 입력된 내용을 string 배열로 변경하는 함수.
	//이 안에서 파싱을 하고 안 닫혔으면 다 닫아줘야 한다.
	return []string{"dummy, data"}
}

func ToHTML(s string, listLen int) (string, int){
	//wiki text를 html로 변경해서 반환하는 함수
	stringtag := []string{"~~", "//", "\\*\\*", "__"}
	wikitag := []string{"~~", "//", "**", "__"}
	htmltag := []string{"del", "i", "b", "u"}
	s, listLen = TransferListTag(s, "*", "ul", listLen)
	for i := 0; i < len(wikitag); i++ {
		s = TransferPropertiesTag(s, stringtag[i], wikitag[i], htmltag[i])
	}
	// 태그에 따라 인풋값을 다르게 하여 TransferlistTag 호출
	return s, listLen
}

func TransferPropertiesTag(input string, stringtag string, wikitag string, htmltag string) string{
	//실질적으로 태그를 변경해서 반환하는 함수.
	var index int = 0
	realTag := []string{"<" + htmltag + ">", "</" + htmltag + ">"}
	matched, err := regexp.MatchString(stringtag + ".*" + stringtag, input)
	if err != nil {
		log.Fatal(nil)
	} else if matched {
		for strings.Contains(input, wikitag){
			input = strings.Replace(input, wikitag, realTag[index], 1)
			index = (index + 1) % 2
		}
	}
	return input
}

func TransferListTag(input string, tag string, htmltag string, indexLen int) (string, int){
	//실질적으로 태그를 변경해서 반환하는 함수
	len := len(strings.Split(input, tag + " ")[0])
	input = strings.Replace(input, "* ", "", 1)
	if indexLen == 0 {
		input = "<" + htmltag + ">\n<li>" + input + "</li>"
	} else{
		if indexLen == len{
			input = "<li>" + input + "</li>"
		} else if indexLen < len{
			input = "<" + htmltag + ">\n<li>" + input + "</li>"
		} else if indexLen > len{
			input = "</" + htmltag + ">\n<li>" + input + "</li>"
		}
	}
	//=====와 같은 태그 (dl 등)을 만나도 마무리 하고 다 닫아줘야 한다.
	//또한 =====태그(dl 등)에 대한 처리 필요
	return input, len
}
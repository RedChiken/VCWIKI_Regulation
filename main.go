package main

import (
	"fmt"
	""
)

func main() {
	fmt.Println(VCWikiRegex.ToHTML(" * ~~deleted~~", 0))
	fmt.Println(VCWikiRegex.TransferPropertiesTag("~~deleted~~", "~~", "~~", "del"))
	fmt.Println(VCWikiRegex.TransferPropertiesTag("//itelic//", "//", "//", "i"))
	fmt.Println(VCWikiRegex.TransferPropertiesTag("**bald**", "\\*\\*", "**", "b"))
	fmt.Println(VCWikiRegex.TransferPropertiesTag("__underline__", "__", "__", "u"))
	//fmt.Println(WikiRegex.ToHTML("real~~text"))
	//fmt.Println(WikiRegex.ToHTML("~~**corrupted text**~~"))
	//fmt.Println(WikiRegex.ToHTML("I hate **//network Homework//** because it's ~~Big Fucking**~~often"))
}

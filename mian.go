package main

import (
	"log"
	"regexp"
)

func main() {
	DotMatch()

	CharacterSetsMatch()

	NotCharacterSetsMatch()

	RepeatMatch()
}

// .是元字符中最简单的例子。
// .匹配任意单个字符，但不匹配换行符。
// 例如，表达式.ar匹配一个任意字符后面跟着是a和r的字符串。
func DotMatch() {
	regex, err := regexp.Compile(".ar")
	if err != nil {
		log.Fatalf("Compile err: %s", err.Error())
	}

	matched := regex.FindAllString("The car parked in the garage.", -1)
	log.Printf("matched: %v\n", matched) //  [car par gar]
}

// 字符集也叫做字符类。
// 方括号用来指定一个字符集。
// 在方括号中使用连字符来指定字符集的范围。
// 在方括号中的字符集不关心顺序。 例如，表达式[Tt]he 匹配 the 和 The。
func CharacterSetsMatch() {
	regex, err := regexp.Compile("[Tt]he")
	if err != nil {
		log.Fatalf("Compile err: %s", err.Error())
	}

	matched := regex.FindAllString("The car parked in the garage.", -1)
	log.Printf("matched: %v\n", matched) // [The the]
}

func NotCharacterSetsMatch() {
	regex, err := regexp.Compile("[^cp]ar")
	if err != nil {
		log.Fatalf("Compile err: %s", err.Error())
	}

	matched := regex.FindAllString("The car parked in the garage.", -1)
	log.Printf("matched: %v\n", matched) // [gar]
}

// 后面跟着元字符 +，* or ? 的，用来指定匹配子模式的次数。
// 这些元字符在不同的情况下有着不同的意思。
func RepeatMatch() {
	// *号匹配 在*之前的字符出现大于等于0次。
	// 例如，表达式 a* 匹配0或更多个以a开头的字符。
	// 表达式[a-z]* 匹配一个行中所有以小写字母开头的字符串。
	regex, err := regexp.Compile("[a-z]*")
	if err != nil {
		log.Fatalf("Compile err: %s", err.Error())
	}

	matched := regex.FindAllString("The car parked in the garage.", -1)
	log.Printf("matched: %v\n", matched) // [ he car parked in the garage ]

	// *字符和.字符搭配可以匹配所有的字符.*。
	// *和表示匹配空格的符号\s连起来用，
	// 如表达式\s*cat\s*匹配0或更多个空格开头和0或更多个空格结尾的cat字符串。
	regex, err = regexp.Compile(`\s*cat\s*`)
	if err != nil {
		log.Fatalf("Compile err: %s", err.Error())
	}

	matched = regex.FindAllString("The fat cat sat on the concatenation.", -1)
	log.Printf("matched: %v\n", matched) // [ cat  cat]

	// +号匹配+号之前的字符出现 >=1 次。
	// 例如表达式c.+t 匹配以首字母c开头以t结尾，中间跟着至少一个字符的字符串。
	regex, err = regexp.Compile("c.+t")
	if err != nil {
		log.Fatalf("Compile err: %s", err.Error())
	}

	matched = regex.FindAllString("The fat cat sat on the concatenation.", -1)
	log.Printf("len: %d, matched: %v\n", len(matched), matched) // [cat sat on the concatenat]

	// 在正则表达式中元字符 ? 标记在符号前面的字符为可选，
	// 即出现 0 或 1 次。 例如，表达式 [T]?he 匹配字符串 he 和 The。
	regex, err = regexp.Compile("T?he")
	if err != nil {
		log.Fatalf("Compile err: %s", err.Error())
	}

	matched = regex.FindAllString("The car is parked in the garage.", -1)
	log.Printf("len: %d, matched: %v\n", len(matched), matched) // [The he]
}

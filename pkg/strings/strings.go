package strings

// AbbreviateUsingEllipses abbreviates a string using ellipses
// 使用"..."对指定字符串进行省略替换
func AbbreviateUsingEllipses(str string, offset, maxWidth int) string {
	return Abbreviate(str, "...", offset, maxWidth)
}

// Abbreviate a string using a given repleacement marker
// 使用给定的字符串进行省略替换
func Abbreviate(str, abbrevMarker string, offset, maxWidth int) string {
	if IsEmpty(str) {
		return str
	}
	var abbrevMarkerLength = Length(abbrevMarker)
	var minAbbrevWidth = abbrevMarkerLength + 1
	var minAbbrevWidthOffset = abbrevMarkerLength + abbrevMarkerLength + 1

	if maxWidth < minAbbrevWidth {
		return str
	}

	if maxWidth < minAbbrevWidthOffset {
		return str
	}

	var strLength = Length(str)

	if strLength <= maxWidth {
		return str
	}

	if offset > strLength {
		offset = strLength
	}

	if (strLength - offset) < maxWidth-abbrevMarkerLength {
		offset = strLength - (maxWidth - abbrevMarkerLength)
	}

	var runes = []rune(str)
	if offset <= minAbbrevWidth {
		return string(runes[:maxWidth-abbrevMarkerLength]) + abbrevMarker
	}

	if (offset + maxWidth - abbrevMarkerLength) < strLength {
		return abbrevMarker + Abbreviate(string(runes[offset:]), abbrevMarker, 0, maxWidth-abbrevMarkerLength)
	}

	return abbrevMarker + string(runes[strLength-(maxWidth-abbrevMarkerLength):])
}

func Abbreviate2(str string, offset, maxWidth int) string {
	if maxWidth < 4 {
		return str
	}

	if Length(str)-offset <= maxWidth {
		return str
	}

	var runes = []rune(str)
	var nRunes = runes[:maxWidth]

	var endIndex = len(nRunes) - 1
	for i := endIndex; i > endIndex-3; i-- {
		nRunes[i] = '.'
	}

	return string(nRunes)
}

// IsEmpty Empty Checker
// 检查字符串是否为空
func IsEmpty(str string) bool {
	return Length(str) == 0
}

// IsNotEmpty Check string is not empty
// 检查字符串是否不为空
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

// Length Return string length
// 返回字符串长度
func Length(str string) int {
	return len(str)
}

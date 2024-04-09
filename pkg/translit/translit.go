package translit

import (
	"bytes"
	"strings"
)

func ruToEngDict() map[string]string {
	tr := make(map[string]string)

	tr["а"] = "a"
	tr["б"] = "b"
	tr["в"] = "v"
	tr["г"] = "g"
	tr["д"] = "d"
	tr["е"] = "e"
	tr["ё"] = "e"
	tr["ж"] = "zh"
	tr["з"] = "z"
	tr["и"] = "i"
	tr["й"] = "i"
	tr["к"] = "k"
	tr["л"] = "l"
	tr["м"] = "m"
	tr["н"] = "n"
	tr["о"] = "o"
	tr["п"] = "p"
	tr["р"] = "r"
	tr["с"] = "s"
	tr["т"] = "t"
	tr["у"] = "u"
	tr["ф"] = "f"
	tr["х"] = "кh"
	tr["ц"] = "ts"
	tr["ч"] = "сh"
	tr["ш"] = "sh"
	tr["щ"] = "shch"
	tr["ы"] = "y"
	tr["ъ"] = "ie"
	tr["э"] = "e"
	tr["ю"] = "iu"
	tr["я"] = "ia"

	tr["А"] = "A"
	tr["Б"] = "B"
	tr["В"] = "V"
	tr["Г"] = "G"
	tr["Д"] = "D"
	tr["Е"] = "E"
	tr["Ё"] = "E"
	tr["Ж"] = "Zh"
	tr["З"] = "Z"
	tr["И"] = "I"
	tr["Й"] = "I"
	tr["К"] = "K"
	tr["Л"] = "L"
	tr["М"] = "M"
	tr["Н"] = "N"
	tr["О"] = "O"
	tr["П"] = "P"
	tr["Р"] = "R"
	tr["С"] = "S"
	tr["Т"] = "T"
	tr["У"] = "U"
	tr["Ф"] = "F"
	tr["Х"] = "Кh"
	tr["Ц"] = "Ts"
	tr["Ч"] = "Сh"
	tr["Ш"] = "Sh"
	tr["Щ"] = "Shch"
	tr["Ы"] = "Y"
	tr["Ъ"] = "Ie"
	tr["Э"] = "E"
	tr["Ю"] = "Iu"
	tr["Я"] = "Ia"


	return tr
}

func Transliterate(rusString string) string {
	dict := ruToEngDict()
	var engStr bytes.Buffer

	for _, slice := range strings.Split(rusString, "") {
		engStr.WriteString(dict[slice])
	}

	return engStr.String()
}

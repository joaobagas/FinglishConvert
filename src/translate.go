package src

func TranslateEngFar(toTranslate string) string {
	pair := ReadCsv("res/convert.csv")
	ret := ""

	for i := 0; i < len(toTranslate); i++ {
		letter := ""
		if i+1 < len(toTranslate) {
			letter = pair[string(toTranslate[i])+string(toTranslate[i+1])]
		}
		if letter == "" {
			letter = pair[string(toTranslate[i])]
		}
		ret += letter
	}

	return ret
}

package lib

import (
	"strings"
)

func CleanText(text string) (string, error) {
	for _, punc := range punctuation {
		text = strings.ReplaceAll(text, punc, "")
	}
	return text, nil
}

func SplitIntoIndexWords(text string) ([]string, error) {
	var result []string
	text, err := CleanText(text)
	if err != nil {
		return result, nil
	}
	for _, word := range strings.Split(text, " ") {
		word = strings.TrimSpace(word)
		word = strings.ToLower(word)
		if word == "" {
			continue
		}
		_, inStopwords := stopwords[word]
		if inStopwords {
			continue
		}
		if len(word) == 1 {
			continue
		}
		result = append(result, word)
	}
	return result, nil
}

func InterfaceToString(data interface{}) (string, error) {
	var result string
	// fmt.Println(reflect.TypeOf(data), data)
	if data == nil {
		return "", nil
	}
	switch data.(type) {
	case string:
		result = data.(string)
	case map[string]interface{}:
		mapdata := data.(map[string]interface{})
		for _, v := range mapdata {
			result = v.([]interface{})[0].(string)
			break
		}
	case map[string]string:
		mapdata := data.(map[string]string)
		for _, v := range mapdata {
			result = v
			break
		}
	case []string:
		result = data.([]string)[0]
	case []interface{}:
		intdata := data.([]interface{})[0]
		mapdata := intdata.(map[string]interface{})
		for _, v := range mapdata {
			result = v.(string)
			break
		}
	default:
		panic("bad interface conversion")
	}

	return result, nil
}

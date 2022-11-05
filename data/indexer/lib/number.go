package lib

func ToHex(number int) (string, error) {
	table := map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5",
		6: "6", 7: "7", 8: "8", 9: "9", 10: "A", 11: "B", 12: "C", 13: "D",
		14: "E", 15: "F"}
	return ToBase(number, 16, table)
}

// https://datatracker.ietf.org/doc/html/rfc1924
func ToAscii85(number int) (string, error) {
	table := map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5",
		6: "6", 7: "7", 8: "8", 9: "9", 10: "A", 11: "B", 12: "C", 13: "D",
		14: "E", 15: "F", 16: "G", 17: "H", 18: "I", 19: "J", 20: "K", 21: "L",
		22: "M", 23: "N", 24: "O", 25: "P", 26: "Q", 27: "R", 28: "S", 29: "T",
		30: "U", 31: "V", 32: "W", 33: "X", 34: "Y", 35: "Z", 36: "a", 37: "b",
		38: "c", 39: "d", 40: "e", 41: "f", 42: "g", 43: "h", 44: "i", 45: "j",
		46: "k", 47: "l", 48: "m", 49: "n", 50: "o", 51: "p", 52: "q", 53: "r",
		54: "s", 55: "t", 56: "u", 57: "v", 58: "w", 59: "x", 60: "y", 61: "z",
		62: "!", 63: "#", 64: "$", 65: "%", 66: "&", 67: "(", 68: ")", 69: "*",
		70: "+", 71: "-", 72: ";", 73: "<", 74: "=", 75: ">", 76: "?", 77: "@",
		78: "^", 79: "_", 80: "`", 81: "{", 82: "|", 83: "}", 84: "~"}
	return ToBase(number, 85, table)
}

func ToBase(number int, base int, table map[int]string) (string, error) {
	if number < base {
		return table[number], nil
	}
	result := ""
	for {
		division := number / base
		remainder := number % base
		number = division
		result = table[remainder] + result
		if division == 0 {
			break
		}
	}
	return result, nil
}

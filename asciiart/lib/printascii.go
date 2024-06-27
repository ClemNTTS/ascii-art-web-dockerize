package lib

func PrintAscii(opt_justify, opt_output, opt_color bool, table_asset [][]string, sentence string) string {
	if sentence == "" {
		return sentence
	}
	var l1, l2, l3, l4, l5, l6, l7, l8 string
	result := ""
	//iteration above the user string
	for i := 0; i < len(sentence); i++ {
		//add on variables each line of characters
		if sentence[i] == '\\' && i < len(sentence) && len(sentence) > 1 {
			if i < len(sentence)-1 && sentence[i+1] == 'n' {
				result += l1 + "\n"
				result += l2 + "\n"
				result += l3 + "\n"
				result += l4 + "\n"
				result += l5 + "\n"
				result += l6 + "\n"
				result += l7 + "\n"
				result += l8 + "\n"
				l1, l2, l3, l4, l5, l6, l7, l8 = "", "", "", "", "", "", "", ""
				i++

			} else {
				l1 += table_asset[sentence[i]-32][0]
				l2 += table_asset[sentence[i]-32][1]
				l3 += table_asset[sentence[i]-32][2]
				l4 += table_asset[sentence[i]-32][3]
				l5 += table_asset[sentence[i]-32][4]
				l6 += table_asset[sentence[i]-32][5]
				l7 += table_asset[sentence[i]-32][6]
				l8 += table_asset[sentence[i]-32][7]
			}
		} else if sentence[i] >= 32 && sentence[i] <= 126 {
			l1 += table_asset[sentence[i]-32][0]
			l2 += table_asset[sentence[i]-32][1]
			l3 += table_asset[sentence[i]-32][2]
			l4 += table_asset[sentence[i]-32][3]
			l5 += table_asset[sentence[i]-32][4]
			l6 += table_asset[sentence[i]-32][5]
			l7 += table_asset[sentence[i]-32][6]
			l8 += table_asset[sentence[i]-32][7]

			//if \n print all if they arn't empty, else \n

			//case for invalid inpput
		} else if sentence[i] == '\r' {
			result += l1 + "\n"
			result += l2 + "\n"
			result += l3 + "\n"
			result += l4 + "\n"
			result += l5 + "\n"
			result += l6 + "\n"
			result += l7 + "\n"
			result += l8 + "\n"
			l1, l2, l3, l4, l5, l6, l7, l8 = "", "", "", "", "", "", "", ""
			i++

		} else {
			return "ERROR"
		}
	}
	result += l1 + "\n"
	result += l2 + "\n"
	result += l3 + "\n"
	result += l4 + "\n"
	result += l5 + "\n"
	result += l6 + "\n"
	result += l7 + "\n"
	result += l8 + "\n"
	return result
}

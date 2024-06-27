package lib

import (
	"os"
)

func CheckOptions() (wrong_input, opt_template, opt_justify, opt_output, opt_color bool, sentence, asset_file_name string) {
	wrong_input, opt_template, opt_justify, opt_output, opt_color = false, false, false, false, false

	asset_file_name = "standard.txt"
	//wrong number of parameters
	if len(os.Args) > 4 {
		wrong_input = true

		//just a string to print
	} else if len(os.Args) == 2 {
		sentence = os.Args[1]
		return wrong_input, opt_template, opt_justify, opt_output, opt_color, sentence, asset_file_name

		//For 2 parameters
	} else if len(os.Args) == 3 {
		//if it's a choice template
		if !StartWith("--align=", os.Args[1]) && !StartWith("--color=", os.Args[1]) && !StartWith("--output=", os.Args[1]) {
			opt_template = true
			path := "assets/" + os.Args[2] + ".txt"
			asset_file_name = os.Args[2] + ".txt"
			if !FileExist(path) {
				wrong_input = true
			}
			sentence = os.Args[1]
			//align case
		} else if StartWith("--align=", os.Args[1]) {
			sentence = os.Args[2]
			switch os.Args[1][7:] {
			case "=right", "=center", "=left", "=justify":
				opt_justify = true
			default:
				wrong_input = true
			}
			//case color
		} else if StartWith("--color=", os.Args[1]) {
			sentence = os.Args[2]
			opt_color = true
		} else if StartWith("--output=", os.Args[1]) {
			sentence = os.Args[2]
			opt_output = true
		} else {
			sentence = os.Args[2]
			wrong_input = true
		}
		//case for 3 parameters
	} else if len(os.Args) == 4 {
		sentence = os.Args[2]
		if !FileExist(string("assets/" + os.Args[3] + ".txt")) {
			wrong_input = true
		} else {
			opt_template = true
			asset_file_name = os.Args[3] + ".txt"
		}
		if StartWith("--align=", os.Args[1]) {
			switch os.Args[1][7:] {
			case "=right", "=center", "=left", "=justify":
				opt_justify = true
			default:
				wrong_input = true
			}
			//case color
		} else if StartWith("--color=", os.Args[1]) {
			opt_color = true
		} else if StartWith("--output=", os.Args[1]) {
			opt_output = true
		} else {
			wrong_input = true
		}

	}

	return wrong_input, opt_template, opt_justify, opt_output, opt_color, sentence, asset_file_name
}

func StartWith(s, sample string) bool {
	if len(s) <= 0 || len(sample) <= 0 || len(s) >= len(sample) {
		return false
	} else if sample[:len(s)] == s {
		return true
	} else {
		return false
	}
}

func FileExist(file_path string) bool {
	_, err := os.ReadFile(file_path)
	if err != nil {
		return false
	} else {
		return true
	}
}

func In(s, sample string) bool {
	if len(s) <= 0 || len(sample) <= 0 || len(s) > len(sample) {
		return false
	}
	for i := 0; i < len(sample)-len(s); i++ {
		for j := i + len(s); j < len(sample); j++ {
			if sample[i:j] == s {
				return true
			}
		}
	}
	return false
}

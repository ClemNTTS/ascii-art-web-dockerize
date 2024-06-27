package ascii

import (
	"errors"
	"lib"
)

func AsciiArt(s, template string) (string, error) {
	//check arguments
	wrong_input, _, opt_justify, opt_output, opt_color, _, _ := lib.CheckOptions()

	if wrong_input || template == "" {
		Error := errors.New("Wrong input!")
		return "", Error
	}
	asset, ThereIsThinker := lib.GetFileContent("assets/" + template + ".txt")
	table_asset := lib.CreateTable(asset, ThereIsThinker)
	result := lib.PrintAscii(opt_justify, opt_output, opt_color, table_asset, s)
	if result == "ERROR" {
		return "", errors.New("Wrong input!")
	}
	return result, nil

}

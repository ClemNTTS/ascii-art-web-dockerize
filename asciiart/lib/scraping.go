package lib

import (
	"strings"
)

func CreateTable(asset string, thinkertoy bool) [][]string {
	assetTable := [][]string{}
	assetChar := []string{}
	cpt := 0
	var workingString []string
	if thinkertoy {
		workingString = strings.Split(asset, "\r\n")
	} else {
		workingString = strings.Split(asset, "\n")
	}
	for i := 1; i < len(workingString); i++ {
		if cpt == 8 {
			assetTable = append(assetTable, assetChar)
			assetChar = []string{}
			cpt = 0
		} else {
			assetChar = append(assetChar, workingString[i])
			cpt++
		}
	}
	return assetTable
}

func GetLetter(asset string) [][]string {

	return nil
}

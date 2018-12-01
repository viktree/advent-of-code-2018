package adventOfCode

import "utils"

func ReadInputFile(day string, fileName string) []string {
	path := "./src/day" + day + "/" + fileName
	return utils.StringSliceFromFile(path)
}

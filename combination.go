package energytypes

import (
	"strconv"
	"strings"
)

const (
	matchAny = ""
)

var permissibleCombinationMap = map[int][]string{
	1: {"F01040100"},
	2: {"F010501xx"},
	3: {"F010502xx"},
	4: {"F010502xx"},
	5: {"F00000000", "F0100xxxx", "F0101xxxx", "F0102xxxx", "F0103xxxx", "F0104xxxx", "F010500xx", "F02xxxxxx"},
	6: {"F03xxxxxx"},
	7: {matchAny},
}

func extractLevelCodePrefix(fullCode string) string {
	return strings.Split(fullCode, "x")[0]
}

func extractLevelCodeSegment(fullCode string) (int, int) {
	codeStr := extractLevelCodePrefix(fullCode)
	codeNum, _ := strconv.Atoi(codeStr[1:])
	levels := (len(codeStr) - 1) / 2
	return codeNum, levels
}

// PermissibleCombination determines whether a proposed combination of a fuel source and technology code is
// permissible for electricity production.
func PermissibleCombination(fuelCode string, technologyCode string) (bool, error) {
	if fuelCode[0] != 'F' || technologyCode[0] != 'T' {
		return false, ErrTypeCodeInvalid
	}

	energyType, err := NewEnergyType(technologyCode)
	if err != nil {
		return false, err
	}

	fuelType, err := NewEnergyType(fuelCode)
	if err != nil {
		return false, err
	}

	typeL1, _ := energyType.Level1Code()

	if fuelCodes, ok := permissibleCombinationMap[typeL1]; ok {
		for _, fullCode := range fuelCodes {
			if fullCode == matchAny {
				return true, nil
			}

			match, level := extractLevelCodeSegment(fullCode)
			if level == 1 {
				fuelL1, _ := fuelType.Level1Code()
				if fuelL1 == match {
					return true, nil
				}
			} else {
				fuelLevelMatch, _ := fuelType.extractLevelCodeRange(1, level)
				if fuelLevelMatch == match {
					return true, nil
				}
			}
		}
	}

	return false, nil
}

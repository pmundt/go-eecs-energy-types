package energytypes

import "fmt"

type OtherGases struct {
	EnergyType
}

var otherGasesLevelMap = map[int]string{
	1: "Anaerobic digestion",
	2: "Gasification",

	100: unspecified,
	101: "Wet fermentation to biogas",
	102: "Dry fermentation to biogas",
	200: unspecified,
	201: "Thermal gasification",

	10000: unspecified,
	10100: unspecified,
	10200: unspecified,
	20000: unspecified,
	20100: unspecified,
}

func (o OtherGases) Level1Description() string {
	code, err := o.Level1Code()
	if err != nil {
		return unspecified
	}
	return otherGasesLevelMap[code]
}

func (o OtherGases) Level2Description() string {
	code, err := o.Level1to2Codes()
	if err != nil {
		return unspecified
	}
	return otherGasesLevelMap[code]
}

func (o OtherGases) Level3Description() string {
	code, err := o.Level1to3Codes()
	if err != nil {
		return unspecified
	}
	return otherGasesLevelMap[code]
}

func (o OtherGases) String() string {
	return fmt.Sprintf("%s > %s > %s", o.Level1Description(), o.Level2Description(), o.Level3Description())
}

func (o OtherGases) ContainsType() bool {
	code, _ := o.Level1to3Codes()
	_, contained := otherGasesLevelMap[code]
	return contained
}

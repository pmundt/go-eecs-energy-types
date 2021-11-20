package energytypes

import "fmt"

type Methane struct {
	EnergyType
}

var methaneLevelMap = map[int]string{
	1: "Anaerobic digestion",
	2: "Gasification",
	3: "Chemical synthesis",

	100: unspecified,
	101: "Fermentation to biogas",
	200: unspecified,
	201: "Thermal gasification",
	300: unspecified,
	301: "Methanation with CO2 from renewable origin",

	10000: unspecified,
	10100: unspecified,
	20000: unspecified,
	20100: unspecified,
	20101: "Gasification of wood",
	30000: unspecified,
	30100: unspecified,
	30101: "Catalytic methanation",
	30102: "Biological methanation",
}

func (m Methane) Level1Description() string {
	code, err := m.Level1Code()
	if err != nil {
		return unspecified
	}
	return methaneLevelMap[code]
}

func (m Methane) Level2Description() string {
	code, err := m.Level1to2Codes()
	if err != nil {
		return unspecified
	}
	return methaneLevelMap[code]
}

func (m Methane) Level3Description() string {
	code, err := m.Level1to3Codes()
	if err != nil {
		return unspecified
	}
	return methaneLevelMap[code]
}

func (m Methane) String() string {
	return fmt.Sprintf("%s > %s > %s", m.Level1Description(), m.Level2Description(), m.Level3Description())
}

func (m Methane) ContainsType() bool {
	code, _ := m.Level1to3Codes()
	_, contained := methaneLevelMap[code]
	return contained
}

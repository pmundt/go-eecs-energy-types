package energytypes

import "fmt"

type Hydrogen struct {
	EnergyType
}

var hydrogenLevelMap = map[int]string{
	1: "Water Electrolysis",
	2: "Chlor-alkali electrolysis",
	3: "Steam Methane Reforming",
	4: "Partial Oxidation",
	5: "Autothermal reforming",
	6: "Methanol reforming",
	7: "Ammonia reforming",
	8: "Gasification",

	100: unspecified,
	101: "Low temperature",
	102: "High temperature",
	200: unspecified,
	300: unspecified,
	301: "Without CCS/CCU",
	302: "With CCS/CCU",
	400: unspecified,
	500: unspecified,
	600: unspecified,
	700: unspecified,
	800: unspecified,

	10000: unspecified,
	10100: unspecified,
	10101: "Main product",
	10200: unspecified,
	10201: "Main product",
	20000: unspecified,
	20001: "By-product",
	30000: unspecified,
	30100: unspecified,
	30101: "Main product",
	30200: unspecified,
	30201: "Main product",
	40000: unspecified,
	50000: unspecified,
	60000: unspecified,
	70000: unspecified,
	80000: unspecified,
}

func (h Hydrogen) Level1Description() string {
	code, err := h.Level1Code()
	if err != nil {
		return unspecified
	}
	return hydrogenLevelMap[code]
}

func (h Hydrogen) Level2Description() string {
	code, err := h.Level1to2Codes()
	if err != nil {
		return unspecified
	}
	return hydrogenLevelMap[code]
}

func (h Hydrogen) Level3Description() string {
	code, err := h.Level1to3Codes()
	if err != nil {
		return unspecified
	}
	return hydrogenLevelMap[code]
}

func (h Hydrogen) String() string {
	return fmt.Sprintf("%s > %s > %s", h.Level1Description(), h.Level2Description(), h.Level3Description())
}

func (h Hydrogen) ContainsType() bool {
	code, _ := h.Level1to3Codes()
	_, contained := hydrogenLevelMap[code]
	return contained
}

package energytypes

import "fmt"

type HeatAndCold struct {
	EnergyType
}

var heatAndColdLevelMap = map[int]string{
	1: "Combustion for heating purpose",
	2: "Heat pump",
	3: "Heating or cooling recovery",
	4: "Geothermal pumping installation",
	5: "Solar thermal collector",
	6: "Electrical resistance heating",

	100: unspecified,
	101: "Flue gas Condensing",
	102: "Non-condensing",
	200: unspecified,
	201: "Electrical",
	202: "Absorption",
	300: unspecified,
	301: "Water-water heat exchange",
	302: "Water-air heat exchange",
	303: "Air-air heat exchange",
	304: "Air-water heat exchange",
	305: "Refrigerant cooling",
	306: "Steam production",
	400: unspecified,
	500: unspecified,
	501: "Non-concentrating",
	502: "Concentrating",
	600: unspecified,
	601: "Electrical boiler",

	10000: unspecified,
	10100: unspecified,
	10101: "CHP",
	10102: "Non-CHP",
	10200: unspecified,
	10201: "CHP",
	10202: "Non-CHP",
	20000: unspecified,
	20100: unspecified,
	20200: unspecified,
	30000: unspecified,
	30100: unspecified,
	30101: "CHP",
	30102: "Non-CHP",
	30200: unspecified,
	30201: "CHP",
	30202: "Non-CHP",
	30300: unspecified,
	30301: "CHP",
	30302: "Non-CHP",
	30400: unspecified,
	30401: "CHP",
	30402: "Non-CHP",
	30500: unspecified,
	30501: "CHP",
	30502: "Non-CHP",
	30600: unspecified,
	30601: "CHP",
	30602: "Non-CHP",
	40000: unspecified,
	50000: unspecified,
	50100: unspecified,
	50101: "Flat plate collector",
	50102: "Evacuated tube collector",
	50200: unspecified,
	50201: "Parabolic trough",
	50202: "Solar power tower",
	50203: "Linear Fresnel reflector",
	50204: "Dish reflector",
	60000: unspecified,
	60100: unspecified,
}

func (h HeatAndCold) Level1Description() string {
	code, err := h.Level1Code()
	if err != nil {
		return unspecified
	}
	return heatAndColdLevelMap[code]
}

func (h HeatAndCold) Level2Description() string {
	code, err := h.Level1to2Codes()
	if err != nil {
		return unspecified
	}
	return heatAndColdLevelMap[code]
}

func (h HeatAndCold) Level3Description() string {
	code, err := h.Level1to3Codes()
	if err != nil {
		return unspecified
	}
	return heatAndColdLevelMap[code]
}

func (h HeatAndCold) String() string {
	return fmt.Sprintf("%s > %s > %s", h.Level1Description(), h.Level2Description(), h.Level3Description())
}

func (h HeatAndCold) ContainsType() bool {
	code, _ := h.Level1to3Codes()
	_, contained := heatAndColdLevelMap[code]
	return contained
}

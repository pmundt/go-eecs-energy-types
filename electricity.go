package energytypes

import "fmt"

type Electricity struct {
	EnergyType
}

var electricityLevelMap = map[int]string{
	1: "Solar",
	2: "Wind",
	3: "Hydro-electric head installations",
	4: "Marine",
	5: "Thermal",
	6: "Nuclear",
	7: "Other",

	100: unspecified,
	101: "Photovoltaic",
	102: "Concentration",
	200: unspecified,
	300: unspecified,
	301: "Run-of-river head installation",
	302: "Storage head installation",
	303: "Pure pumped storage head installation",
	304: "Mixed pumped storage head",
	400: unspecified,
	401: "Tidal",
	402: "Wave",
	403: "Currents",
	404: "Pressure",
	500: unspecified,
	501: "Combined cycle gas turbine with heat recovery",
	502: "Steam turbine with back-pressure turbine (open cycle)",
	503: "Steam turbine with condensation turbine (closed cycle)",
	504: "Gas turbine with heat recovery",
	505: "Internal combustion engine",
	506: "Micro-turbine",
	507: "Stirling engine",
	508: "Fuel cell",
	509: "Steam engine",
	510: "Organic Rankine cycle",
	511: "Gas turbine without heat recovery",
	600: unspecified,
	601: "Heavy-water reactor",
	602: "Light water reactor",
	603: "Breeder",
	604: "Graphite reactor",
	700: unspecified,

	10000: unspecified,
	10100: unspecified,
	10101: "Classic silicon",
	10102: "Thin film",
	10200: unspecified,
	20000: unspecified,
	20001: "Onshore",
	20002: "Offshore",
	30000: unspecified,
	30100: unspecified,
	30200: unspecified,
	30300: unspecified,
	30400: unspecified,
	40000: unspecified,
	40100: unspecified,
	40101: "Onshore",
	40102: "Offshore",
	40200: unspecified,
	40201: "Onshore",
	40202: "Offshore",
	40300: unspecified,
	40400: unspecified,
	50000: unspecified,
	50100: unspecified,
	50101: "Non-CHP",
	50102: "CHP",
	50200: unspecified,
	50201: "Non-CHP",
	50202: "CHP",
	50300: unspecified,
	50301: "Non-CHP",
	50302: "CHP",
	50400: unspecified,
	50401: "Non-CHP",
	50402: "CHP",
	50500: unspecified,
	50501: "Non-CHP",
	50502: "CHP",
	50600: unspecified,
	50601: "Non-CHP",
	50602: "CHP",
	50700: unspecified,
	50701: "Non-CHP",
	50702: "CHP",
	50800: unspecified,
	50801: "Non-CHP",
	50802: "CHP",
	50900: unspecified,
	50901: "Non-CHP",
	50902: "CHP",
	51000: unspecified,
	51001: "Non-CHP",
	51002: "CHP",
	51100: unspecified,
	60000: unspecified,
	60100: unspecified,
	60200: unspecified,
	60300: unspecified,
	60400: unspecified,
	70000: unspecified,
}

func (e Electricity) Level1Description() string {
	code, err := e.Level1Code()
	if err != nil {
		return unspecified
	}
	return electricityLevelMap[code]
}

func (e Electricity) Level2Description() string {
	code, err := e.Level1to2Codes()
	if err != nil {
		return unspecified
	}
	return electricityLevelMap[code]
}

func (e Electricity) Level3Description() string {
	code, err := e.Level1to3Codes()
	if err != nil {
		return unspecified
	}
	return electricityLevelMap[code]
}

func (e Electricity) String() string {
	return fmt.Sprintf("%s > %s > %s", e.Level1Description(), e.Level2Description(), e.Level3Description())
}

func (e Electricity) ContainsType() bool {
	code, _ := e.Level1to3Codes()
	_, contained := electricityLevelMap[code]
	return contained
}

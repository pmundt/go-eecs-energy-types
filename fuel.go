package energytypes

import "fmt"

type Fuel struct {
	EnergyType
}

var carbonContentMap = map[int]float32{
	2000000: 247.4,
	2010000: 111.9,
	2010100: 111.9,
	2010101: 98.3,
	2010102: 94.7,
	2010103: 94.0,
	2010104: 111.9,
	2010105: 111.9,
	2010200: 106.0,
	2010201: 96.1,
	2010202: 101.2,
	2010203: 94.6,
	2010204: 106.0,
	2010300: 106.0,
	2010400: 73.6,
	2010500: 73.6,
	2010501: 73.6,
	2020000: 100.8,
	2020100: 73.3,
	2020101: 73.3,
	2020200: 63.1,
	2020300: 77.4,
	2020301: 61.6,
	2020302: 73.3,
	2020303: 72.0,
	2020304: 72.0,
	2020305: 71.5,
	2020306: 71.9,
	2020307: 74.3,
	2020308: 74.3,
	2020309: 77.4,
	2020310: 66.7,
	2020311: 80.7,
	2020312: 80.7,
	2020313: 73.3,
	2020314: 100.8,
	2020315: 73.3,
	2030000: 247.4,
	2030100: 56.1,
	2030200: 247.4,
	2030201: 247.4,
	2030202: 41.2,
	2030300: 73.7,
	2030301: 73.7,
	2030302: 73.7,
	2030303: 66.7,
	2030304: 66.7,
	2030500: 191.9,
	2030501: 155.2,
	2030502: 54.9,
	2030503: 56.1,
	2030504: 149.5,
	2030505: 191.9,
}

var fuelLevelMap = map[int]string{
	0: unspecified,
	1: "Renewable",
	2: "Fossil",
	3: "Nuclear",
	4: "Gas synthesis",
	5: "Waste heat and cold",

	100: unspecified,
	101: "Solid",
	102: "Liquid",
	103: "Gaseous",
	104: "Heating and cooling",
	105: "Mechanical source or other",
	200: unspecified,
	201: "Solid",
	202: "Liquid",
	203: "Gaseous",
	204: "Heat",
	301: "Solid",
	400: unspecified,
	401: "Furnace Gas",
	500: unspecified,
	501: "By-product in industrial installation",
	502: "By-product in power generation",
	503: "By-product in tertiary sector",

	10000: unspecified,
	10100: unspecified,
	10101: "Municipal waste",
	10102: "Industrial and commercial waste",
	10103: "Wood",
	10104: "Animal fats",
	10105: "Biomass from agriculture",
	10200: unspecified,
	10201: "Municipal biodegradable waste",
	10202: "Black liquor",
	10203: "Pure plant oil",
	10204: "Waste plant oil",
	10205: "Refined vegetable oil",
	10300: unspecified,
	10301: "Landfill gas",
	10302: "Sewage gas",
	10303: "Agricultural gas",
	10304: "Gas from organic waste digestion",
	10305: "Process gas",
	10306: "Other biogenic sources",
	10401: "Solar",
	10402: "Geothermal",
	10403: "Aerothermal",
	10404: "Hydrothermal",
	10405: "Process heat",
	10500: unspecified,
	10501: "Wind",
	10502: "Hydro & marine",
	20000: unspecified,
	20100: unspecified,
	20101: "Hard coal",
	20102: "Brown coal",
	20103: "Peat",
	20104: "Municipal waste",
	20105: "Industrial and commercial waste",
	20200: unspecified,
	20201: "Crude oil",
	20202: "Natural gas liquids (NGL)",
	20203: "Petroleum products",
	20300: unspecified,
	20301: "Natural gas",
	20302: "Coal-derived gas",
	20303: "Petroleum products",
	20304: "Municipal gas plant",
	20305: "Process gas",
	20400: unspecified,
	20401: "Process heat",
	30101: "Radioactive fuel",
	40000: unspecified,
	40100: unspecified,
	50000: unspecified,
	50100: unspecified,
	50200: unspecified,
	50300: unspecified,

	1000000: unspecified,
	1010000: unspecified,
	1010101: "Biogenic",
	1010201: "Biogenic",
	1010300: unspecified,
	1010301: "Forestry products",
	1010302: "Forestry by-products & waste",
	1010303: "Saw products, by-products and waste",
	1010400: unspecified,
	1010500: unspecified,
	1010501: "Agricultural products",
	1010502: "Agricultural by-products & waste",
	1020000: unspecified,
	1020100: unspecified,
	1020200: unspecified,
	1020300: unspecified,
	1020301: "Rapeseed (Brassica napus L.)",
	1020302: "Sunflower (Helianthus anuus L.)",
	1020303: "Oil palm (Elaeis guineensis Jacq.)",
	1020304: "Coconut (Cocos nucifera L.)",
	1020305: "Jatropha",
	1020400: unspecified,
	1020500: unspecified,
	1020501: "Biodiesel (mono-alkyl ester)",
	1020502: "Biogasoline (C6-C12 hydrocarbon)",
	1030000: unspecified,
	1030100: unspecified,
	1030200: unspecified,
	1030300: unspecified,
	1030301: "Pig manure",
	1030302: "Cow manure",
	1030303: "Chicken manure",
	1030304: "Unspecified manure",
	1030305: "Energy crops",
	1030306: "Digestion of pure manure",
	1030307: "Digestion of manure with energy crops",
	1030400: unspecified,
	1030401: "Organic waste unspecified",
	1030402: "Agricultural waste unspecified",
	1030403: "Agricultural waste from farm fertiliser",
	1030404: "Agricultural waste from straw",
	1030405: "Waste from food industry",
	1030406: "Manure with organic waste",
	1030407: "Manure with organic waste and energy crops",
	1030408: "Other biogenic waste",
	1030501: "Biogenic",
	1030601: unspecified,
	1040100: unspecified,
	1040200: unspecified,
	1040201: "Conventional geothermal heat",
	1040202: "Enhanced dry bed geothermal heat",

	// FIXME: The specification has a typo here, and defines F02040203 as the full code, while the level breakdown and
	// placement means it should be F01040203. Create an alias for both until the typo can be corrected in the next
	// version of the guidelines.
	1040203: "Shallow geothermal heat/cold",
	2040203: "Shallow geothermal heat/cold",

	1040300: unspecified,
	1040400: unspecified,
	1040401: "River",
	1040402: "Lake",
	1040501: "Biogenic",
	1050000: unspecified,
	1050100: unspecified,
	1050200: unspecified,
	2000000: unspecified,
	2010000: unspecified,
	2010100: unspecified,
	2010101: "Anthracite",
	2010102: "Bituminous coal",
	2010103: "Coking coal",
	2010104: "Coke-oven coke",
	2010105: "Lignite coke",
	2010200: unspecified,
	2010201: "Sub-bituminous coal",
	2010202: "Lignite",
	2010203: "Brown coal briquette",
	2010204: "Peat briquette",
	2010300: unspecified,
	2010400: unspecified,
	2010500: unspecified,
	2010501: "Non-renewable",
	2020000: unspecified,
	2020100: unspecified,
	2020101: "Shale oil",
	2020200: unspecified,
	2020300: unspecified,
	2020301: "Ethane",
	2020302: "Naphtha",
	2020303: "Aviation gasoline",
	2020304: "Motor gasoline",
	2020305: "Aviation turbine fuel",
	2020306: "Other kerosene",
	2020307: "Gas/diesel oil",
	2020308: "Fuel oil, low sulphur",
	2020309: "Fuel oil, high sulphur",
	2020310: "Liquid Petroleum Gas",
	2020311: "Orimulsion",
	2020312: "Bitumen",
	2020313: "Lubricants",
	2020314: "Petroleum coke",
	2020315: "Refinery Feedstock",
	2030000: unspecified,
	2030100: unspecified,
	2030200: unspecified,
	2030201: "Blast furnace gas",
	2030202: "Coke-oven gas",
	2030300: unspecified,
	2030301: "Propane",
	2030302: "Butane",
	2030303: "Refinery gas",
	2030304: "Chemical waste gas",
	2030400: unspecified,
	2030500: unspecified,
	2030501: "Carbon monoxide",
	2030502: "Methane",
	2030503: "Hydrogen (fossil sourced)",
	2030504: "Phosphor gas",
	2030505: "Oxy gas",
	2040000: unspecified,
	2040001: "Non-renewable",
	2040100: unspecified,
	2040101: "Non-renewable",
	3010100: unspecified,
	3010101: "UOX",
	3010102: "AGR",
	3010103: "MOX",
	4000000: unspecified,
	4010000: unspecified,
	5000000: unspecified,
	5010000: unspecified,
	5020000: unspecified,
	5030000: unspecified,
}

func (f Fuel) Level1Description() string {
	code, err := f.Level1Code()
	if err != nil {
		return unspecified
	}
	return fuelLevelMap[code]
}

func (f Fuel) Level2Description() string {
	code, err := f.Level1to2Codes()
	if err != nil {
		return unspecified
	}
	return fuelLevelMap[code]
}

func (f Fuel) Level3Description() string {
	code, err := f.Level1to3Codes()
	if err != nil {
		return unspecified
	}
	return fuelLevelMap[code]
}

func (f Fuel) Level4Description() string {
	code, err := f.Level1to4Codes()
	if err != nil {
		return unspecified
	}
	return fuelLevelMap[code]
}

func (f Fuel) CarbonContent() float32 {
	code, err := f.Level1to4Codes()
	if err != nil {
		return 0.00
	}
	return carbonContentMap[code]
}

func (f Fuel) String() string {
	return fmt.Sprintf("%s > %s > %s > %s", f.Level1Description(), f.Level2Description(), f.Level3Description(), f.Level4Description())
}

func (f Fuel) ContainsType() bool {
	code, _ := f.Level1to4Codes()
	_, contained := fuelLevelMap[code]
	return contained
}

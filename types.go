package energytypes

import (
	"errors"
	"strconv"
)

var (
	ErrLevelOutOfBounds      = errors.New("level out of bounds")
	ErrLevelRangeInvalid     = errors.New("invalid level range specified")
	ErrTypeCodeInvalid       = errors.New("invalid type code specified")
	ErrDeviceTypeNotFound    = errors.New("no matching device type found")
	ErrDeviceTypeCodeInvalid = errors.New("invalid code for device type")
	ErrEnergySourceNotFound  = errors.New("no matching energy source found")
)

const (
	unspecified = "Unspecified"
)

type ProductionDevice interface {
	ContainsType() bool
	Level1Description() string
	Level2Description() string
	Level3Description() string
	String() string
}

type EnergySource interface {
	ProductionDevice
	Level4Description() string
	CarbonContent() float32
}

type EnergyType struct {
	FullCode string
	levels   int
}

var typeCodeLevelMap = map[byte]int{
	'F': 4,
	'G': 3,
	'H': 3,
	'M': 3,
	'T': 3,
	'Q': 3,
}

func (e EnergyType) typeCode() byte {
	return e.FullCode[0]
}

func (e EnergyType) valid() bool {
	// Make sure that the numeric part of the code is valid
	_, err := strconv.Atoi(e.FullCode[1:])
	if err != nil {
		return false
	}

	// Ensure the type code and number of levels matches the spec
	if level, ok := typeCodeLevelMap[e.typeCode()]; ok {
		return e.levels == level
	}

	return false
}

func (e EnergyType) extractLevelCode(level int) (int, error) {
	if level <= 0 || level > e.levels {
		return -1, ErrLevelOutOfBounds
	}

	startIndex := level + (level - 1)
	endIndex := startIndex + 2

	return strconv.Atoi(e.FullCode[startIndex:endIndex])
}

func (e EnergyType) extractLevelCodeRange(startLevel int, endLevel int) (int, error) {
	if startLevel <= 0 || startLevel >= endLevel {
		return -1, ErrLevelRangeInvalid
	}
	if endLevel > e.levels {
		return -1, ErrLevelOutOfBounds
	}

	startIndex := startLevel + (startLevel - 1)
	endIndex := (endLevel + (endLevel - 1)) + 2

	return strconv.Atoi(e.FullCode[startIndex:endIndex])
}

func (e EnergyType) Level1Code() (int, error) {
	return e.extractLevelCode(1)
}

func (e EnergyType) Level2Code() (int, error) {
	return e.extractLevelCode(2)
}

func (e EnergyType) Level3Code() (int, error) {
	return e.extractLevelCode(3)
}

func (e EnergyType) Level4Code() (int, error) {
	return e.extractLevelCode(4)
}

func (e EnergyType) Level1to2Codes() (int, error) {
	return e.extractLevelCodeRange(1, 2)
}

func (e EnergyType) Level1to3Codes() (int, error) {
	return e.extractLevelCodeRange(1, 3)
}

func (e EnergyType) Level1to4Codes() (int, error) {
	return e.extractLevelCodeRange(1, 4)
}

func NewEnergyType(fullCode string) (EnergyType, error) {
	energyType := EnergyType{
		FullCode: fullCode,
		// Calculated as number of digit pairs less the single character type prefix
		levels: (len(fullCode) - 1) / 2,
	}

	if !energyType.valid() {
		return EnergyType{}, ErrTypeCodeInvalid
	}

	return energyType, nil
}

func NewEnergySource(fullCode string) (EnergySource, error) {
	var es EnergySource

	et, err := NewEnergyType(fullCode)
	if err != nil {
		return nil, err
	}

	switch et.typeCode() {
	case 'F':
		es = Fuel{et}
		break
	default:
		return nil, ErrEnergySourceNotFound
	}

	if !es.ContainsType() {
		return nil, ErrTypeCodeInvalid
	}

	return es, nil
}

func NewProductionDevice(fullCode string) (ProductionDevice, error) {
	var pd ProductionDevice

	et, err := NewEnergyType(fullCode)
	if err != nil {
		return nil, err
	}

	switch et.typeCode() {
	case 'G':
		pd = OtherGases{et}
		break
	case 'H':
		pd = Hydrogen{et}
		break
	case 'M':
		pd = Methane{et}
		break
	case 'Q':
		pd = HeatAndCold{et}
		break
	case 'T':
		pd = Electricity{et}
		break
	default:
		return nil, ErrDeviceTypeNotFound
	}

	if !pd.ContainsType() {
		return nil, ErrDeviceTypeCodeInvalid
	}

	return pd, nil
}

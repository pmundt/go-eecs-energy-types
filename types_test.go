package energytypes

import "testing"

func TestEnergyType_typeCode(t *testing.T) {
	fullCode := "T010000"
	energyType, _ := NewEnergyType(fullCode)
	typeCode := energyType.typeCode()
	if typeCode != 'T' {
		t.Errorf("Incorrect type code, got: %c, expected: T\n", typeCode)
	}
}

func TestEnergyType_Level1Code(t *testing.T) {
	fullCode := "T012345"
	energyType, _ := NewEnergyType(fullCode)
	l1, _ := energyType.Level1Code()
	if l1 != 1 {
		t.Errorf("Incorrect type code, got: %d, expected: 1\n", l1)
	}
}

func TestEnergyType_Level2Code(t *testing.T) {
	fullCode := "T012345"
	energyType, _ := NewEnergyType(fullCode)
	l2, _ := energyType.Level2Code()
	if l2 != 23 {
		t.Errorf("Incorrect type code, got: %d, expected: 23", l2)
	}
}

func TestEnergyType_Level3Code(t *testing.T) {
	fullCode := "T012345"
	energyType, _ := NewEnergyType(fullCode)
	l3, _ := energyType.Level3Code()
	if l3 != 45 {
		t.Errorf("Incorrect type code, got: %d, expected: 45", l3)
	}
}

func TestEnergyType_Level4Code(t *testing.T) {
	fullCode := "F01234567"
	energyType, _ := NewEnergyType(fullCode)
	l4, _ := energyType.Level4Code()
	if l4 != 67 {
		t.Errorf("Incorrect type code, got: %d, expected: 67", l4)
	}
}

func TestEnergyType_Level1to2Codes(t *testing.T) {
	fullCode := "T012345"
	energyType, _ := NewEnergyType(fullCode)
	level, _ := energyType.Level1to2Codes()
	if level != 123 {
		t.Errorf("Incorrect type code, got: %d, expected: 123\n", level)
	}
}

func TestEnergyType_Level1to3Codes(t *testing.T) {
	fullCode := "F01234567"
	energyType, _ := NewEnergyType(fullCode)
	level, _ := energyType.Level1to3Codes()
	if level != 12345 {
		t.Errorf("Incorrect type code, got: %d, expected: 12345\n", level)
	}
}

func TestEnergyType_Level1to4Codes(t *testing.T) {
	fullCode := "F01234567"
	energyType, _ := NewEnergyType(fullCode)
	level, _ := energyType.Level1to4Codes()
	if level != 1234567 {
		t.Errorf("Incorrect type code, got: %d, expected: 1234567\n", level)
	}

	level3Code := "T012345"
	energyType, _ = NewEnergyType(level3Code)
	level, err := energyType.Level1to4Codes()
	if err != ErrLevelOutOfBounds {
		t.Errorf("Expected out of bounds error\n")
	}
}

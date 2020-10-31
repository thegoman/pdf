package pdfunit

import (
	"fmt"
	"github.com/shopspring/decimal"
)

// UnitType for UnitType constants
type UnitType int

const (
	//Point unit type
	Point UnitType = iota

	//Pica unit type
	Pica

	//Inch unit type
	Inch

	//Millimeter unit type
	Millimeter
)

// Unit struct set/get and calculate units point, pica, inch and mm
//
// if one value is set it calculates the other units automatically
type Unit struct {
	point      decimal.Decimal
	pica       decimal.Decimal
	inch       decimal.Decimal
	millimeter decimal.Decimal
}

//NewUnit Creates a new Unit struct
func NewUnit() *Unit {
	return &Unit{}
}

//SetValue sets the value of the given unit and calculate the other units
func (u *Unit) SetValue(value float64, unit UnitType) error {
	switch unit {
	case Point:
		u.SetPoint(decimal.NewFromFloat(value))
		return nil
	case Pica:
		u.SetPica(decimal.NewFromFloat(value))
		return nil
	case Inch:
		u.SetInch(decimal.NewFromFloat(value))
		return nil
	case Millimeter:
		u.SetMillimeter(decimal.NewFromFloat(value))
		return nil
	default:
		return fmt.Errorf("invalid unit")
	}
}

//GetValue gets the unit value, the value rounded on 3 decimal places
func (u *Unit) GetValue(unit UnitType) (*float64, error) {
	switch unit {
	case Point:
		value, _ := u.GetPoint().Round(3).Float64()
		return &value, nil
	case Pica:
		value, _ := u.GetPica().Round(3).Float64()
		return &value, nil
	case Inch:
		value, _ := u.GetInch().Round(3).Float64()
		return &value, nil
	case Millimeter:
		value, _ := u.GetMillimeter().Round(3).Float64()
		return &value, nil
	default:
		return nil, fmt.Errorf("invalid unit")
	}
}

//GetPoint gets the point value as decimal.Decimal
func (u Unit) GetPoint() decimal.Decimal {
	return u.point
}

//GetPica gets the pica value as decimal.Decimal
func (u Unit) GetPica() decimal.Decimal {
	return u.pica
}

//GetInch gets the inch value as decimal.Decimal
func (u Unit) GetInch() decimal.Decimal {
	return u.inch
}

//GetMillimeter gets the millimeter as decimal.Decimal
func (u Unit) GetMillimeter() decimal.Decimal {
	return u.millimeter
}

//SetPoint sets the point value and calculate the other values
func (u *Unit) SetPoint(point decimal.Decimal) {
	u.point = point
	u.pica = point.Div(decimal.NewFromInt(12))
	u.inch = point.Div(decimal.NewFromInt(72))
	u.millimeter = u.inch.Mul(decimal.NewFromFloat(25.4))
}

//SetPica sets the pica value and calculate the other values
func (u *Unit) SetPica(pica decimal.Decimal) {
	u.point = pica.Mul(decimal.NewFromInt(12))
	u.pica = pica
	u.inch = pica.Div(decimal.NewFromInt(6))
	u.millimeter = u.inch.Mul(decimal.NewFromFloat(25.4))
}

//SetInch sets the inch value and calculate the other values
func (u *Unit) SetInch(inch decimal.Decimal) {
	u.point = inch.Mul(decimal.NewFromInt(72))
	u.pica = inch.Mul(decimal.NewFromInt(6))
	u.inch = inch
	u.millimeter = inch.Mul(decimal.NewFromFloat(25.4))
}

//SetMillimeter sets the millimeter value and calculate the other values
func (u *Unit) SetMillimeter(millimeter decimal.Decimal) {
	u.inch = millimeter.Div(decimal.NewFromFloat(25.4))
	u.point = u.inch.Mul(decimal.NewFromInt(72))
	u.pica = u.inch.Mul(decimal.NewFromInt(6))
	u.millimeter = millimeter
}

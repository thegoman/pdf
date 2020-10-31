package pdfunit

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

type data struct {
	Point      decimal.Decimal
	Pica       decimal.Decimal
	Inch       decimal.Decimal
	Millimeter decimal.Decimal
}

func getData() []data {

	pica := decimal.NewFromFloat(1.0).Div(decimal.NewFromFloat(12.0))

	inch := decimal.NewFromFloat(1.0).Div(decimal.NewFromFloat(72.0))

	millimeter := inch.Mul(decimal.NewFromFloat(25.4))

	var dataset = []data{
		{
			decimal.NewFromFloat(72.0),
			decimal.NewFromFloat(6.0),
			decimal.NewFromFloat(1.0),
			decimal.NewFromFloat(25.4),
		},
		{
			decimal.NewFromFloat(1.0),
			pica,
			inch,
			millimeter,
		},
	}
	return dataset
}

func TestUnit_GetPoint(t *testing.T) {
	u := NewUnit()
	assert.NotNil(t, u)
	assert.EqualValues(t, 0, u.GetPoint().Cmp(decimal.NewFromFloat(0.0)))
}

func TestUnit_GetPica(t *testing.T) {
	u := NewUnit()
	assert.NotNil(t, u)
	assert.EqualValues(t, 0, u.GetPica().Cmp(decimal.NewFromFloat(0.0)))
}

func TestUnit_GetInch(t *testing.T) {
	u := NewUnit()
	assert.NotNil(t, u)
	assert.EqualValues(t, 0, u.GetInch().Cmp(decimal.NewFromFloat(0.0)))
}

func TestUnit_GetMillimeter(t *testing.T) {
	u := NewUnit()
	assert.NotNil(t, u)
	assert.EqualValues(t, 0, u.GetMillimeter().Cmp(decimal.NewFromFloat(0.0)))
}

func TestUnit_SetPoint(t *testing.T) {
	u := NewUnit()
	assert.NotNil(t, u)
	for _, unitData := range getData() {
		u.SetPoint(unitData.Point)

		assert.EqualValues(t, 0, unitData.Pica.Cmp(u.GetPica()))
		assert.EqualValues(t, 0, unitData.Inch.Cmp(u.GetInch()))
		assert.EqualValues(t, 0, unitData.Millimeter.Cmp(u.GetMillimeter()))
	}
}

func TestUnit_SetPica(t *testing.T) {
	u := NewUnit()
	assert.NotNil(t, u)
	for _, unitData := range getData() {
		u.SetPica(unitData.Pica)

		assert.EqualValues(t, 0, unitData.Point.Round(0).Cmp(u.GetPoint().Round(0)))
		assert.EqualValues(t, 0, unitData.Inch.Cmp(u.GetInch()))
		assert.EqualValues(t, 0, unitData.Millimeter.Cmp(u.GetMillimeter()))
	}
}

func TestUnit_SetInch(t *testing.T) {
	u := NewUnit()
	assert.NotNil(t, u)
	for _, unitData := range getData() {
		u.SetInch(unitData.Inch)

		assert.EqualValues(t, 0, unitData.Point.Round(0).Cmp(u.GetPoint().Round(0)))
		assert.EqualValues(t, 0, unitData.Pica.Round(3).Cmp(u.GetPica().Round(3)))
		assert.EqualValues(t, 0, unitData.Millimeter.Cmp(u.GetMillimeter()))
	}
}

func TestUnit_SetMillimeter(t *testing.T) {
	u := NewUnit()
	assert.NotNil(t, u)
	for _, unitData := range getData() {
		u.SetMillimeter(unitData.Millimeter)

		assert.EqualValues(t, 0, unitData.Point.Round(0).Cmp(u.GetPoint().Round(0)))
		assert.EqualValues(t, 0, unitData.Pica.Round(3).Cmp(u.GetPica().Round(3)))
		assert.EqualValues(t, 0, unitData.Inch.Cmp(u.GetInch()))
	}
}

func TestUnit_SetValue(t *testing.T) {
	t.Run("error on invalid unit", func(t *testing.T) {
		u := NewUnit()
		err := u.SetValue(32.0, UnitType(99))

		assert.NotNil(t, err)
		if err != nil {
			assert.EqualValues(t, "invalid unit", err.Error())
		}
	})

	t.Run("unit type point", func(t *testing.T) {
		u := NewUnit()
		err := u.SetValue(32.2, Point)

		assert.Nil(t, err)
		val, _ := u.GetPoint().Float64()
		assert.EqualValues(t, 32.2, val)
	})

	t.Run("unit type pica", func(t *testing.T) {
		u := NewUnit()
		err := u.SetValue(32.2, Pica)

		assert.Nil(t, err)
		val, _ := u.GetPica().Float64()
		assert.EqualValues(t, 32.2, val)
	})

	t.Run("unit type inch", func(t *testing.T) {
		u := NewUnit()
		err := u.SetValue(32.2, Inch)

		assert.Nil(t, err)
		val, _ := u.GetInch().Float64()
		assert.EqualValues(t, 32.2, val)
	})

	t.Run("unit type millimeter", func(t *testing.T) {
		u := NewUnit()
		err := u.SetValue(32.2, Millimeter)

		assert.Nil(t, err)
		val, _ := u.GetMillimeter().Float64()
		assert.EqualValues(t, 32.2, val)
	})
}

func TestUnit_GetValue(t *testing.T) {
	t.Run("error on invalid unit", func(t *testing.T) {
		u := NewUnit()
		_ = u.SetValue(32.0, Millimeter)
		_, err := u.GetValue(UnitType(99))

		assert.NotNil(t, err)
		if err != nil {
			assert.EqualValues(t, "invalid unit", err.Error())
		}
	})

	t.Run("unit type point", func(t *testing.T) {
		u := NewUnit()
		_ = u.SetValue(32.2266, Point)
		val, err := u.GetValue(Point)

		assert.Nil(t, err)
		assert.EqualValues(t, 32.227, *val)
	})

	t.Run("unit type pica", func(t *testing.T) {
		u := NewUnit()
		_ = u.SetValue(32.2266, Pica)
		val, err := u.GetValue(Pica)

		assert.Nil(t, err)
		assert.EqualValues(t, 32.227, *val)
	})

	t.Run("unit type inch", func(t *testing.T) {
		u := NewUnit()
		_ = u.SetValue(32.2266, Inch)
		val, err := u.GetValue(Inch)

		assert.Nil(t, err)
		assert.EqualValues(t, 32.227, *val)
	})

	t.Run("unit type millimeter", func(t *testing.T) {
		u := NewUnit()
		_ = u.SetValue(32.2266, Millimeter)
		val, err := u.GetValue(Millimeter)

		assert.Nil(t, err)
		assert.EqualValues(t, 32.227, *val)
	})
}

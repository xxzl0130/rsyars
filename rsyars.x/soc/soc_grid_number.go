package soc

import (
	"github.com/pkg/errors"
)

const (
	GridNumberTypeNone GridNumberType = iota
	GridNumberType6
	GridNumberType5
	GridNumberType4
	GridNumberType3
)

type GridNumberType int

func GetGridNumber(c *SoC) (GridNumberType, error) {
	if len(c.ChipID) != 4 {
		return GridNumberTypeNone, errors.Errorf("illegal chip_id <%s>", c.ChipID)
	}
	switch c.ChipID[2] {
	case '6':
		return GridNumberType6, nil
	case '5':
		return GridNumberType5, nil
	case '4':
		return GridNumberType4, nil
	case '3':
		return GridNumberType3, nil
	}
	return GridNumberTypeNone, errors.Errorf("illegal chip_id <%s>", c.ChipID)
}

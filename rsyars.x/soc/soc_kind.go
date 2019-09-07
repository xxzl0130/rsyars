package soc

import "github.com/pkg/errors"

const (
	KindTypeNone KindType = iota
	KindType1
	KindType2
)

type KindType int

func GetKind(c *SoC) (KindType, error) {
	if len(c.ChipID) != 4 {
		return KindTypeNone, errors.Errorf("illegal chip_id <%s>", c.ChipID)
	}
	switch c.ChipID[3] {
	case '1':
		return KindType2, nil
	case '2':
		return KindType1, nil
	}
	return KindTypeNone, errors.Errorf("illegal chip_id <%s>", c.ChipID)
}

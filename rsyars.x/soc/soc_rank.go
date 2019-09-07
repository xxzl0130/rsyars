package soc

import "github.com/pkg/errors"

const (
	RankTypeNone RankType = iota
	RankType2
	RankType3
	RankType4
	RankType5
)

type RankType int

func GetRank(c *SoC) (RankType, error) {
	if len(c.ChipID) != 4 {
		return RankTypeNone, errors.Errorf("illegal chip_id <%s>", c.ChipID)
	}
	switch c.ChipID[0] {
	case '2':
		return RankType2, nil
	case '3':
		return RankType3, nil
	case '4':
		return RankType4, nil
	case '5':
		return RankType5, nil
	}
	return RankTypeNone, errors.Errorf("illegal chip_id <%s>", c.ChipID)
}

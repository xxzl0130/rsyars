package soc

import "github.com/pkg/errors"

const (
	ColorTypeNone ColorType = iota
	ColorTypeBlue
	ColorTypeOrange
)

type ColorType int

func GetColor(c *SoC) (ColorType, error) {
	switch c.ColorID {
	case "1":
		return ColorTypeOrange, nil
	case "2":
		return ColorTypeBlue, nil
	}
	return ColorTypeNone, errors.Errorf("illegal color_id <%s>", c.ColorID)
}

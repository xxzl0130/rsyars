package hycdes

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	"github.com/xxzl0130/rsyars/rsyars.x/soc"
)

var (
	rules = []string{"InfinityFrost", "FatalChapters"}
)

type SoC struct {
	ID      int
	Color   string
	Class   string
	Type    string
	Level   string
	Hit     int
	Reload  int
	Damage  int
	Destroy int
}

func NewSoC(c *soc.SoC) (*SoC, error) {
	ic := new(SoC)

	id, err := strconv.Atoi(c.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	color, err := soc.GetColor(c)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	gridNumber, err := soc.GetGridNumber(c)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	_, err = soc.GetKind(c)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	hit, reload, damage, destroy, err := soc.GetProperty(c)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	rank, err := soc.GetRank(c)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	shape, err := soc.GetShape(c)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ic.ID = id

	switch color {
	case soc.ColorTypeBlue:
		ic.Color = "1"
	case soc.ColorTypeOrange:
		ic.Color = "2"
	default:
		return nil, errors.Errorf("unknown color<%d>", color)
	}

	switch rank {
	case soc.RankType5:
		switch gridNumber {
		case soc.GridNumberType6:
			ic.Class = "56"
		case soc.GridNumberType5:
			ic.Class = "551"
		default:
			return nil, errors.Errorf("unknown grid_number<%d>", gridNumber)
		}
	default:
		return nil, errors.Errorf("unknown rank<%d>", rank)
	}

	switch shape {
	case soc.ShapeType61:
		ic.Type = "1"
	case soc.ShapeType62:
		ic.Type = "2"
	case soc.ShapeType63:
		ic.Type = "3"
	case soc.ShapeType641:
		ic.Type = "41"
	case soc.ShapeType642:
		ic.Type = "42"
	case soc.ShapeType65:
		ic.Type = "5"
	case soc.ShapeType66:
		ic.Type = "6"
	case soc.ShapeType67:
		ic.Type = "7"
	case soc.ShapeType68:
		ic.Type = "8"
	case soc.ShapeType69:
		ic.Type = "9"
	case soc.ShapeType511:
		ic.Type = "5"
	case soc.ShapeType512:
		ic.Type = "22"
	case soc.ShapeType513:
		ic.Type = "21"
	case soc.ShapeType514:
		ic.Type = "32"
	case soc.ShapeType515:
		ic.Type = "31"
	case soc.ShapeType516:
		ic.Type = "6"
	case soc.ShapeType517:
		ic.Type = "4"
	case soc.ShapeType518:
		ic.Type = "11"
	case soc.ShapeType519:
		ic.Type = "12"
	case soc.ShapeType521:
		ic.Type = "132"
	case soc.ShapeType522:
		ic.Type = "131"
	case soc.ShapeType523:
		ic.Type = "120"
	case soc.ShapeType524:
		ic.Type = "112"
	case soc.ShapeType525:
		ic.Type = "111"
	case soc.ShapeType526:
		ic.Type = "10"
	case soc.ShapeType527:
		ic.Type = "9"
	case soc.ShapeType528:
		ic.Type = "82"
	case soc.ShapeType529:
		ic.Type = "81"
	default:
		return nil, errors.Errorf("unknown shape<%d>", shape)
	}

	ic.Level = c.ChipLevel

	ic.Hit, ic.Reload, ic.Damage, ic.Destroy = hit, reload, damage, destroy

	return ic, nil
}

func Build(targets []*SoC) (string, error) {
	w := bytes.NewBufferString(fmt.Sprintf("[%c!", rules[0][getPos(targets)]))
	for i, target := range targets {
		w.WriteString(fmt.Sprintf("%d,", i))
		w.WriteString(fmt.Sprintf("%s,", target.Color))
		w.WriteString(fmt.Sprintf("%s,", target.Class))
		w.WriteString(fmt.Sprintf("%s,", target.Type))
		w.WriteString(fmt.Sprintf("%s,", target.Level))
		w.WriteString(fmt.Sprintf("%d,", target.Hit))
		w.WriteString(fmt.Sprintf("%d,", target.Reload))
		w.WriteString(fmt.Sprintf("%d,", target.Damage))
		w.WriteString(fmt.Sprintf("%d,", target.Destroy))
		w.WriteString(fmt.Sprintf("%d&", 1))
	}
	w.WriteString(fmt.Sprintf("?%c]", rules[1][getPos(targets)]))

	return w.String(), nil
}

func getPos(targets []*SoC) int {
	return len(targets) - 13*len(targets)/13
}

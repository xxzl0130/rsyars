package soc

import (
	"github.com/pkg/errors"
)

/**

**
**
**
ShapeType61

***
**
*
ShapeType62

****
 **
ShapeType63

***
 ***
ShapeType641

 ***
***
ShapeType642

**
 **
**
ShapeType65

 *
***
 *
 *
ShapeType66

 *
 *
 *
 *
 *
 *
ShapeType67

****
*  *
ShapeType68

**
***
 *
ShapeType69

 **
**
*
ShapeType511

 ***
**
ShapeType512

***
  **
ShapeType513

*
**
*
*
ShapeType514

 *
**
 *
 *
ShapeType515

 *
***
 *
ShapeType516

***
 *
 *
ShapeType517

 **
**
 *
ShapeType518

**
 **
 *
ShapeType519

 *
 *
 *
**
ShapeType521

*
*
*
**
ShapeType522

***
*
*
ShapeType523

 **
 *
**
ShapeType524

**
 *
 **
ShapeType525

**
*
**
ShapeType526

*
*
*
*
*
ShapeType527

 *
**
**
ShapeType528

*
**
**
ShapeType529
*/
const (
	ShapeTypeNone ShapeType = iota
	ShapeType61
	ShapeType62
	ShapeType63
	ShapeType641
	ShapeType642
	ShapeType65
	ShapeType66
	ShapeType67
	ShapeType68
	ShapeType69
	ShapeType511
	ShapeType512
	ShapeType513
	ShapeType514
	ShapeType515
	ShapeType516
	ShapeType517
	ShapeType518
	ShapeType519
	ShapeType521
	ShapeType522
	ShapeType523
	ShapeType524
	ShapeType525
	ShapeType526
	ShapeType527
	ShapeType528
	ShapeType529
)

type ShapeType int

func GetShape(c *SoC) (ShapeType, error) {
	switch c.GridID {
	case "30":
		return ShapeType61, nil
	case "31":
		return ShapeType62, nil
	case "32":
		return ShapeType63, nil
	case "33":
		return ShapeType641, nil
	case "34":
		return ShapeType642, nil
	case "35":
		return ShapeType65, nil
	case "36":
		return ShapeType66, nil
	case "37":
		return ShapeType67, nil
	case "38":
		return ShapeType68, nil
	case "39":
		return ShapeType69, nil
	case "21":
		return ShapeType511, nil
	case "22":
		return ShapeType512, nil
	case "23":
		return ShapeType513, nil
	case "24":
		return ShapeType514, nil
	case "25":
		return ShapeType515, nil
	case "26":
		return ShapeType516, nil
	case "27":
		return ShapeType517, nil
	case "28":
		return ShapeType518, nil
	case "29":
		return ShapeType519, nil
	case "20":
		return ShapeType521, nil
	case "19":
		return ShapeType522, nil
	case "18":
		return ShapeType523, nil
	case "17":
		return ShapeType524, nil
	case "16":
		return ShapeType525, nil
	case "15":
		return ShapeType526, nil
	case "14":
		return ShapeType527, nil
	case "13":
		return ShapeType528, nil
	case "12":
		return ShapeType529, nil
	case "11", "10", "9", "8", "7", "6", "5", "4", "3", "2", "1":
		return ShapeTypeNone, errors.Errorf("unknown grid_id <%s>", c.GridID)
	}
	return ShapeTypeNone, errors.Errorf("illegal grid_id <%s>", c.GridID)
}

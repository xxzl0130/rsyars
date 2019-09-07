package soc

import (
	"strconv"

	"github.com/pkg/errors"
)

func GetProperty(c *SoC) (hit, reload, damage, destroy int, err error) {
	hit, err = strconv.Atoi(c.AssistHit)
	if err != err {
		return 0, 0, 0, 0, errors.WithMessagef(err, "illegal assist_hit<%s>", c.AssistHit)
	}
	reload, err = strconv.Atoi(c.AssistReload)
	if err != err {
		return 0, 0, 0, 0, errors.WithMessagef(err, "illegal assist_reload<%s>", c.AssistHit)
	}
	damage, err = strconv.Atoi(c.AssistDamage)
	if err != err {
		return 0, 0, 0, 0, errors.WithMessagef(err, "illegal assist_damage<%s>", c.AssistHit)
	}
	destroy, err = strconv.Atoi(c.AssistDefBreak)
	if err != err {
		return 0, 0, 0, 0, errors.WithMessagef(err, "illegal assist_def_break<%s>", c.AssistHit)
	}
	return hit, reload, damage, destroy, nil
}

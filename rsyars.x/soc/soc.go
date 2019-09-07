package soc

type SoC struct {
	ID              string `json:"id"`
	ChipID          string `json:"chip_id"`
	ChipLevel       string `json:"chip_level"`
	ColorID         string `json:"color_id"`
	GridID          string `json:"grid_id"`
	SquadWithUserID string `json:"squad_with_user_id"`
	ShapeInfo       string `json:"shape_info"`
	AssistDamage    string `json:"assist_damage"`
	AssistReload    string `json:"assist_reload"`
	AssistHit       string `json:"assist_hit"`
	AssistDefBreak  string `json:"assist_def_break"`
	Locked          string `json:"is_locked"`
}

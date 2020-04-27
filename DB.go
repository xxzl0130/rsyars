package main

import (
	"encoding/json"
	"strconv"
	"fmt"

	"github.com/gobuffalo/packr"
)

type enemy_team_infoT struct {
	Id		int		`json:"id"`
	Leader  int		`json:"enemy_leader"`
	Members []int	`json:"member_ids"`
	Night	bool	`json:"is_night"`
}

type enemy_in_team_infoT struct {
	Id		int		`json:"id"`
	TeamId  int		`json:"enemy_team_id"`
	TypeId	int		`json:"enemy_character_type_id"`
}

type enemy_character_typeT struct {
	Id		int		`json:"id"`
	Name	string	`json:"name"`
}

type tranlsationT struct{
	EnemyName	map[string]string	`json:"enemy_char_name"`
}

func (rs *rsyars) InitDB() error{
	db := packr.NewBox("./GFDB/src/db/json")
	locale := packr.NewBox("./GFDB/src/locales/zh")
	data,e := db.FindString("enemy_team_info.json")
	if e != nil{
		rs.log.Errorf("读取JSON数据失败 -> %+v", e)
		return e
	}
	if err := json.Unmarshal([]byte(data), &rs.teamInfo); err != nil {
		rs.log.Errorf("解析JSON数据失败 -> %+v", err)
		return err
	}
	
	data,e = db.FindString("enemy_in_team_info.json")
	if e != nil{
		rs.log.Errorf("读取JSON数据失败 -> %+v", e)
		return e
	}
	if err := json.Unmarshal([]byte(data), &rs.enemyInfo); err != nil {
		rs.log.Errorf("解析JSON数据失败 -> %+v", err)
		return err
	}
	
	data,e = db.FindString("enemy_character_type_info.json")
	if e != nil{
		rs.log.Errorf("读取JSON数据失败 -> %+v", e)
		return e
	}
	if err := json.Unmarshal([]byte(data), &rs.characterInfo); err != nil {
		rs.log.Errorf("解析JSON数据失败 -> %+v", err)
		return err
	}

	data,e = locale.FindString("translation.json")
	if e != nil{
		rs.log.Errorf("读取JSON数据失败 -> %+v", e)
		return e
	}
	if err := json.Unmarshal([]byte(data), &rs.translation); err != nil {
		rs.log.Errorf("解析JSON数据失败 -> %+v", err)
		return err
	}

	data,e = locale.FindString("table.json")
	if e != nil{
		rs.log.Errorf("读取JSON数据失败 -> %+v", e)
		return e
	}
	if err := json.Unmarshal([]byte(data), &rs.table); err != nil {
		rs.log.Errorf("解析JSON数据失败 -> %+v", err)
		return err
	}

	return nil
}

func (rs *rsyars) areaNo2EnemyTeamId(area int, no int) string{
	id := 90000000 + area * 100 + (no + 1)
	return strconv.Itoa(id)
}

func (rs *rsyars) getEnemyTeamInfo(id string) string{
	team,ok := rs.teamInfo[id]
	if !ok{
		return ""
	}
	var count map[int]int
	count = make(map[int]int)
	for _,member := range team.Members{
		enemy,ok := rs.enemyInfo[strconv.Itoa(member)]
		if !ok{
			return ""
		}
		count[enemy.TypeId]++
	}
	str := ""
	for k,v := range count{
		name,ok := rs.translation.EnemyName[strconv.Itoa(k)]
		if !ok{
			character,ok := rs.characterInfo[strconv.Itoa(k)];
			if ok{
				name,ok = rs.table[character.Name];
				if !ok{
					name = ""
				}
			}
		}
		str += fmt.Sprintf("%s*%d ",name,v)
	}
	if team.Night{
		str += "夜战"
	}

	return str
}
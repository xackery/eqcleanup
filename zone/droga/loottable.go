package droga

import (
	"database/sql"
	"github.com/xackery/goeq/loot"
)

var loottables []loot.LootTable = []loot.LootTable{
	{Name: sql.NullString{String: "A_Goblin_Traitor", Valid: true}, Mincash: 858, Maxcash: 5685, Avgcoin: 0, Done: 1, Id: sql.NullInt64{Int64: 4502, Valid: true}},
	{Maxcash: 0, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 13492, Valid: true}, Name: sql.NullString{String: "Warlord_Skargus", Valid: true}, Mincash: 0},
	{Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 90833, Valid: true}, Name: sql.NullString{String: "90833_a_goblin_warder_MAGELO-GEN", Valid: true}, Mincash: 32, Maxcash: 3200},
	{Id: sql.NullInt64{Int64: 90837, Valid: true}, Name: sql.NullString{String: "90837_#Overseer_Dlubish_MAGELO-GEN", Valid: true}, Mincash: 30, Maxcash: 3000, Avgcoin: 0, Done: 0},
	{Id: sql.NullInt64{Int64: 90838, Valid: true}, Name: sql.NullString{String: "90838_a_goblin_stonechanter_MAGELO-GEN", Valid: true}, Mincash: 31, Maxcash: 3100, Avgcoin: 0, Done: 0},
	{Id: sql.NullInt64{Int64: 92258, Valid: true}, Name: sql.NullString{String: "92258_a_goblin_taskmaster_MAGELO-GEN", Valid: true}, Mincash: 33, Maxcash: 3300, Avgcoin: 0, Done: 0},
	{Id: sql.NullInt64{Int64: 92262, Valid: true}, Name: sql.NullString{String: "92262_a_goblin_legate_MAGELO-GEN", Valid: true}, Mincash: 33, Maxcash: 3300, Avgcoin: 0, Done: 0},
	{Name: sql.NullString{String: "92263_a_goblin_dirtscriber_MAGELO-GEN", Valid: true}, Mincash: 33, Maxcash: 3300, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 92263, Valid: true}},
	{Mincash: 33, Maxcash: 3300, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 92264, Valid: true}, Name: sql.NullString{String: "92264_a_goblin_flamedrowser_MAGELO-GEN", Valid: true}},
	{Id: sql.NullInt64{Int64: 92265, Valid: true}, Name: sql.NullString{String: "92265_a_goblin_interrogator_MAGELO-GEN", Valid: true}, Mincash: 35, Maxcash: 3500, Avgcoin: 0, Done: 0},
	{Maxcash: 3500, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 92266, Valid: true}, Name: sql.NullString{String: "92266_Jeren_Manri_MAGELO-GEN", Valid: true}, Mincash: 35},
	{Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 92268, Valid: true}, Name: sql.NullString{String: "92268_a_goblin_slavedriver_MAGELO-GEN", Valid: true}, Mincash: 31, Maxcash: 3100},
	{Id: sql.NullInt64{Int64: 92273, Valid: true}, Name: sql.NullString{String: "92273_a_goblin_plaguebringer_MAGELO-GEN", Valid: true}, Mincash: 29, Maxcash: 2900, Avgcoin: 0, Done: 0},
	{Id: sql.NullInt64{Int64: 92274, Valid: true}, Name: sql.NullString{String: "92274_a_goblin_slinker_MAGELO-GEN", Valid: true}, Mincash: 29, Maxcash: 2900, Avgcoin: 0, Done: 0},
	{Id: sql.NullInt64{Int64: 92276, Valid: true}, Name: sql.NullString{String: "92276_a_goblin_dirtcaller_MAGELO-GEN", Valid: true}, Mincash: 29, Maxcash: 2900, Avgcoin: 0, Done: 0},
	{Maxcash: 2900, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 92277, Valid: true}, Name: sql.NullString{String: "92277_a_goblin_mosstrooper_MAGELO-GEN", Valid: true}, Mincash: 29},
	{Maxcash: 2900, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 92278, Valid: true}, Name: sql.NullString{String: "92278_a_goblin_rockchanter_MAGELO-GEN", Valid: true}, Mincash: 29},
	{Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 92279, Valid: true}, Name: sql.NullString{String: "92279_a_goblin_firedrowser_MAGELO-GEN", Valid: true}, Mincash: 29, Maxcash: 2900},
	{Id: sql.NullInt64{Int64: 94575, Valid: true}, Name: sql.NullString{String: "94575_a_goblin_boneseer_MAGELO-GEN", Valid: true}, Mincash: 31, Maxcash: 3100, Avgcoin: 0, Done: 0},
	{Maxcash: 3100, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 94576, Valid: true}, Name: sql.NullString{String: "94576_a_goblin_blightcaller_MAGELO-GEN", Valid: true}, Mincash: 31},
	{Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 94587, Valid: true}, Name: sql.NullString{String: "94587_a_goblin_stalagknight_MAGELO-GEN", Valid: true}, Mincash: 34, Maxcash: 3400},
	{Id: sql.NullInt64{Int64: 94588, Valid: true}, Name: sql.NullString{String: "94588_a_goblin_backbiter_MAGELO-GEN", Valid: true}, Mincash: 34, Maxcash: 3400, Avgcoin: 0, Done: 0},
	{Done: 0, Id: sql.NullInt64{Int64: 94589, Valid: true}, Name: sql.NullString{String: "94589_a_goblin_dirttracer_MAGELO-GEN", Valid: true}, Mincash: 34, Maxcash: 3400, Avgcoin: 0},
	{Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 94591, Valid: true}, Name: sql.NullString{String: "94591_Soothsayer_Dregzak_MAGELO-GEN", Valid: true}, Mincash: 40, Maxcash: 4000},
	{Name: sql.NullString{String: "94592_Chief_RokGus_MAGELO-GEN", Valid: true}, Mincash: 40, Maxcash: 4000, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 94592, Valid: true}},
	{Maxcash: 4000, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 94593, Valid: true}, Name: sql.NullString{String: "94593_a_maddened_Burynai_MAGELO-GEN", Valid: true}, Mincash: 40},
	{Done: 0, Id: sql.NullInt64{Int64: 94594, Valid: true}, Name: sql.NullString{String: "94594_a_goblin_bodyguard_MAGELO-GEN", Valid: true}, Mincash: 40, Maxcash: 4000, Avgcoin: 0},
	{Done: 0, Id: sql.NullInt64{Int64: 94595, Valid: true}, Name: sql.NullString{String: "94595_a_goblin_whipcracker_MAGELO-GEN", Valid: true}, Mincash: 29, Maxcash: 2900, Avgcoin: 0},
	{Id: sql.NullInt64{Int64: 94596, Valid: true}, Name: sql.NullString{String: "94596_a_goblin_flashdrowser_MAGELO-GEN", Valid: true}, Mincash: 40, Maxcash: 4000, Avgcoin: 0, Done: 0},
	{Id: sql.NullInt64{Int64: 94597, Valid: true}, Name: sql.NullString{String: "94597_a_goblin_depredator_MAGELO-GEN", Valid: true}, Mincash: 40, Maxcash: 4000, Avgcoin: 0, Done: 0},
	{Done: 0, Id: sql.NullInt64{Int64: 94598, Valid: true}, Name: sql.NullString{String: "94598_a_goblin_cavehunter_MAGELO-GEN", Valid: true}, Mincash: 40, Maxcash: 4000, Avgcoin: 0},
	{Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 94599, Valid: true}, Name: sql.NullString{String: "94599_a_goblin_boneslaver_MAGELO-GEN", Valid: true}, Mincash: 40, Maxcash: 4000},
	{Id: sql.NullInt64{Int64: 94600, Valid: true}, Name: sql.NullString{String: "94600_a_goblin_bonecharmer_MAGELO-GEN", Valid: true}, Mincash: 40, Maxcash: 4000, Avgcoin: 0, Done: 0},
	{Name: sql.NullString{String: "94601_a_goblin_adept_MAGELO-GEN", Valid: true}, Mincash: 40, Maxcash: 4000, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 94601, Valid: true}},
	{Maxcash: 3400, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 94602, Valid: true}, Name: sql.NullString{String: "94602_a_goblin_penkeeper_MAGELO-GEN", Valid: true}, Mincash: 34},
	{Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 94603, Valid: true}, Name: sql.NullString{String: "94603_a_goblin_penmaster_MAGELO-GEN", Valid: true}, Mincash: 34, Maxcash: 3400},
	{Id: sql.NullInt64{Int64: 94873, Valid: true}, Name: sql.NullString{String: "an_Iksar_slave", Valid: true}, Mincash: 5, Maxcash: 500, Avgcoin: 0, Done: 0},
	{Id: sql.NullInt64{Int64: 94880, Valid: true}, Name: sql.NullString{String: "a_goblin_fanatic", Valid: true}, Mincash: 50, Maxcash: 500, Avgcoin: 0, Done: 0},
}

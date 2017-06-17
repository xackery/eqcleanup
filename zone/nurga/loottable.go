package nurga

import (
	"database/sql"
	"github.com/xackery/goeq/loot"
)

var loottables []loot.LootTable = []loot.LootTable{
	{Mincash: 968, Maxcash: 6411, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 11501, Valid: true}, Name: sql.NullString{String: "a_Burynai_slave", Valid: true}},
	{Name: sql.NullString{String: "90833_a_goblin_warder_MAGELO-GEN", Valid: true}, Mincash: 32, Maxcash: 3200, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 90833, Valid: true}},
	{Done: 0, Id: sql.NullInt64{Int64: 90837, Valid: true}, Name: sql.NullString{String: "90837_#Overseer_Dlubish_MAGELO-GEN", Valid: true}, Mincash: 30, Maxcash: 3000, Avgcoin: 0},
	{Done: 0, Id: sql.NullInt64{Int64: 90838, Valid: true}, Name: sql.NullString{String: "90838_a_goblin_stonechanter_MAGELO-GEN", Valid: true}, Mincash: 31, Maxcash: 3100, Avgcoin: 0},
	{Done: 0, Id: sql.NullInt64{Int64: 90840, Valid: true}, Name: sql.NullString{String: "90840_a_sleeping_ogre_MAGELO-GEN", Valid: true}, Mincash: 30, Maxcash: 3000, Avgcoin: 0},
	{Mincash: 31, Maxcash: 3100, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 94568, Valid: true}, Name: sql.NullString{String: "94568_a_goblin_slavedriver_MAGELO-GEN", Valid: true}},
	{Name: sql.NullString{String: "94575_a_goblin_boneseer_MAGELO-GEN", Valid: true}, Mincash: 31, Maxcash: 3100, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 94575, Valid: true}},
	{Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 94576, Valid: true}, Name: sql.NullString{String: "94576_a_goblin_blightcaller_MAGELO-GEN", Valid: true}, Mincash: 31, Maxcash: 3100},
	{Id: sql.NullInt64{Int64: 94585, Valid: true}, Name: sql.NullString{String: "Trunt", Valid: true}, Mincash: 0, Maxcash: 0, Avgcoin: 0, Done: 0},
	{Mincash: 31, Maxcash: 3100, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 94586, Valid: true}, Name: sql.NullString{String: "94586_a_goblin_flashdrowser_MAGELO-GEN", Valid: true}},
	{Id: sql.NullInt64{Int64: 94587, Valid: true}, Name: sql.NullString{String: "94587_a_goblin_stalagknight_MAGELO-GEN", Valid: true}, Mincash: 34, Maxcash: 3400, Avgcoin: 0, Done: 0},
	{Id: sql.NullInt64{Int64: 94588, Valid: true}, Name: sql.NullString{String: "94588_a_goblin_backbiter_MAGELO-GEN", Valid: true}, Mincash: 34, Maxcash: 3400, Avgcoin: 0, Done: 0},
	{Id: sql.NullInt64{Int64: 94589, Valid: true}, Name: sql.NullString{String: "94589_a_goblin_dirttracer_MAGELO-GEN", Valid: true}, Mincash: 34, Maxcash: 3400, Avgcoin: 0, Done: 0},
	{Maxcash: 0, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 94849, Valid: true}, Name: sql.NullString{String: "a_tortured_slave", Valid: true}, Mincash: 0},
}

package veeshan

import (
	"database/sql"
	"github.com/xackery/goeq/loot"
)

var loottables []loot.LootTable = []loot.LootTable{
	{Id: sql.NullInt64{Int64: 4642, Valid: true}, Name: sql.NullString{String: "veeshan A_Guardian_Wurm", Valid: true}, Mincash: 0, Maxcash: 0, Avgcoin: 0, Done: 1},
	{Id: sql.NullInt64{Int64: 87652, Valid: true}, Name: sql.NullString{String: "Silverwing_1.0", Valid: true}, Mincash: 1408, Maxcash: 9315, Avgcoin: 0, Done: 1},
	{Id: sql.NullInt64{Int64: 87653, Valid: true}, Name: sql.NullString{String: "Phara_Dar_1.0", Valid: true}, Mincash: 1408, Maxcash: 9315, Avgcoin: 0, Done: 1},
	{Id: sql.NullInt64{Int64: 87654, Valid: true}, Name: sql.NullString{String: "Xygoz_1.0", Valid: true}, Mincash: 1408, Maxcash: 9315, Avgcoin: 0, Done: 1},
	{Id: sql.NullInt64{Int64: 87655, Valid: true}, Name: sql.NullString{String: "Druushk_1.0", Valid: true}, Mincash: 1408, Maxcash: 9315, Avgcoin: 0, Done: 1},
	{Maxcash: 9315, Avgcoin: 0, Done: 1, Id: sql.NullInt64{Int64: 87656, Valid: true}, Name: sql.NullString{String: "Nexona_1.0", Valid: true}, Mincash: 1408},
	{Avgcoin: 0, Done: 1, Id: sql.NullInt64{Int64: 87657, Valid: true}, Name: sql.NullString{String: "Hoshkar_1.0", Valid: true}, Mincash: 1408, Maxcash: 9315},
	{Id: sql.NullInt64{Int64: 90845, Valid: true}, Name: sql.NullString{String: "90845_a_racnar_MAGELO-GEN", Valid: true}, Mincash: 0, Maxcash: 0, Avgcoin: 0, Done: 0},
	{Id: sql.NullInt64{Int64: 90846, Valid: true}, Name: sql.NullString{String: "90846_lava_drake_MAGELO-GEN", Valid: true}, Mincash: 60, Maxcash: 6000, Avgcoin: 0, Done: 0},
}

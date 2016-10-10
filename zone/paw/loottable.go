package paw

import (
	"database/sql"
	"github.com/xackery/goeq/loot"
)

var loottables []loot.LootTable = []loot.LootTable{
	{Maxcash: 4233, Avgcoin: 0, Done: 1, Id: sql.NullInt64{Int64: 2971, Valid: true}, Name: sql.NullString{String: "A_Lteth_Val_Scribe", Valid: true}, Mincash: 638},
	{Done: 1, Id: sql.NullInt64{Int64: 11627, Valid: true}, Name: sql.NullString{String: "a_Tesch_Val_Guard", Valid: true}, Mincash: 814, Maxcash: 5394, Avgcoin: 0},
	{Mincash: 0, Maxcash: 0, Avgcoin: 0, Done: 1, Id: sql.NullInt64{Int64: 11843, Valid: true}, Name: sql.NullString{String: "A_Gaduladian_Widemouth", Valid: true}},
	{Id: sql.NullInt64{Int64: 11889, Valid: true}, Name: sql.NullString{String: "A_Hiding_Gnoll", Valid: true}, Mincash: 418, Maxcash: 2781, Avgcoin: 0, Done: 1},
	{Id: sql.NullInt64{Int64: 12017, Valid: true}, Name: sql.NullString{String: "A_Gnoll_Prisoner", Valid: true}, Mincash: 396, Maxcash: 2636, Avgcoin: 0, Done: 1},
	{Maxcash: 3652, Avgcoin: 0, Done: 1, Id: sql.NullInt64{Int64: 90208, Valid: true}, Name: sql.NullString{String: "Tesch_Val_Deval`Nmak", Valid: true}, Mincash: 0},
	{Maxcash: 3300, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 91623, Valid: true}, Name: sql.NullString{String: "91623_Verishe_Mal_Executioner_MAGELO-GEN", Valid: true}, Mincash: 33},
	{Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 91631, Valid: true}, Name: sql.NullString{String: "91631_Verishe_Mal_Judge_MAGELO-GEN", Valid: true}, Mincash: 32, Maxcash: 3200},
	{Id: sql.NullInt64{Int64: 92466, Valid: true}, Name: sql.NullString{String: "92466_The_Ishva_Mal_MAGELO-GEN", Valid: true}, Mincash: 40, Maxcash: 4000, Avgcoin: 0, Done: 0},
	{Name: sql.NullString{String: "92467_a_Ishva_Lteth_gnoll_MAGELO-GEN", Valid: true}, Mincash: 35, Maxcash: 3500, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 92467, Valid: true}},
	{Mincash: 35, Maxcash: 3500, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 92468, Valid: true}, Name: sql.NullString{String: "92468_a_one_eyed_gnoll_MAGELO-GEN", Valid: true}},
	{Id: sql.NullInt64{Int64: 92469, Valid: true}, Name: sql.NullString{String: "92469_Nisch_Val_Torash_Mashk_MAGELO-GEN", Valid: true}, Mincash: 35, Maxcash: 3500, Avgcoin: 0, Done: 0},
	{Maxcash: 2500, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 92470, Valid: true}, Name: sql.NullString{String: "92470_Kurrpok_Splitpaw_MAGELO-GEN", Valid: true}, Mincash: 25},
	{Mincash: 5, Maxcash: 500, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 92471, Valid: true}, Name: sql.NullString{String: "92471_Rosch_Val_L`Vlor_MAGELO-GEN", Valid: true}},
	{Id: sql.NullInt64{Int64: 92472, Valid: true}, Name: sql.NullString{String: "92472_Tesch_Val_Kadvem_MAGELO-GEN", Valid: true}, Mincash: 5, Maxcash: 500, Avgcoin: 0, Done: 0},
	{Done: 0, Id: sql.NullInt64{Int64: 93724, Valid: true}, Name: sql.NullString{String: "93724_a_Tesch_Val_Gnoll_MAGELO-GEN", Valid: true}, Mincash: 32, Maxcash: 3200, Avgcoin: 0},
	{Mincash: 25, Maxcash: 2500, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 93725, Valid: true}, Name: sql.NullString{String: "93725_a_Tesch_Mas_Gnoll_MAGELO-GEN", Valid: true}},
	{Name: sql.NullString{String: "93726_a_Tesch_Mal_Gnoll_MAGELO-GEN", Valid: true}, Mincash: 29, Maxcash: 2900, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 93726, Valid: true}},
	{Mincash: 33, Maxcash: 3300, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 93727, Valid: true}, Name: sql.NullString{String: "93727_a_Rosch_Val_Gnoll_MAGELO-GEN", Valid: true}},
	{Id: sql.NullInt64{Int64: 93728, Valid: true}, Name: sql.NullString{String: "93728_a_Rosch_Mas_Gnoll_MAGELO-GEN", Valid: true}, Mincash: 25, Maxcash: 2500, Avgcoin: 0, Done: 0},
	{Id: sql.NullInt64{Int64: 93729, Valid: true}, Name: sql.NullString{String: "93729_a_Nisch_Mas_Gnoll_MAGELO-GEN", Valid: true}, Mincash: 27, Maxcash: 2700, Avgcoin: 0, Done: 0},
	{Mincash: 31, Maxcash: 3100, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 93730, Valid: true}, Name: sql.NullString{String: "93730_a_Nisch_Mal_Gnoll_MAGELO-GEN", Valid: true}},
	{Maxcash: 2400, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 93731, Valid: true}, Name: sql.NullString{String: "93731_a_Lteth_Val_Gnoll_MAGELO-GEN", Valid: true}, Mincash: 24},
	{Maxcash: 2000, Avgcoin: 0, Done: 0, Id: sql.NullInt64{Int64: 93732, Valid: true}, Name: sql.NullString{String: "93732_a_Lteth_Mas_Gnoll_MAGELO-GEN", Valid: true}, Mincash: 20},
}

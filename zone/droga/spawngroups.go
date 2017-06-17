package droga

import (
	"database/sql"
	"github.com/xackery/goeq/spawn"
)

var spawngroups []spawn.SpawnGroup = []spawn.SpawnGroup{
	{Id: sql.NullInt64{Int64: 81003, Valid: true}, Name: sql.NullString{String: "droga_an_Iksar_slave3", Valid: true}, Dist: 0.000000, Max_x: 0.000000, Delay: 0, Despawn_timer: 100, Spawn_limit: 0, Min_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn: 0},
	{Spawn_limit: 0, Max_x: 0.000000, Mindelay: 15000, Despawn: 0, Delay: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 81010, Valid: true}, Name: sql.NullString{String: "droga_a_goblin_taskmaster0", Valid: true}, Dist: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000},
	{Min_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Id: sql.NullInt64{Int64: 81074, Valid: true}, Name: sql.NullString{String: "droga_a_goblin_interrogator74", Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Delay: 0, Mindelay: 15000, Despawn: 0, Despawn_timer: 100},
	{Max_x: 0.000000, Min_y: 0.000000, Delay: 0, Mindelay: 15000, Name: sql.NullString{String: "droga_Jeren_Manri462", Valid: true}, Spawn_limit: 0, Dist: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 81116, Valid: true}},
	{Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 81117, Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Max_y: 0.000000, Delay: 0, Name: sql.NullString{String: "droga_a_goblin_traitor465", Valid: true}, Min_x: 0.000000, Min_y: 0.000000, Mindelay: 15000},
	{Min_x: 0.000000, Max_y: 0.000000, Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Min_y: 0.000000, Delay: 0, Mindelay: 15000, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 81135, Valid: true}, Name: sql.NullString{String: "droga_nothing463", Valid: true}},
	{Max_y: 0.000000, Min_y: 0.000000, Delay: 0, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 81449, Valid: true}, Name: sql.NullString{String: "droga_#Overseer_Dlubish449", Valid: true}, Max_x: 0.000000, Min_x: 0.000000, Mindelay: 15000, Spawn_limit: 0, Dist: 0.000000},
	{Max_x: 0.000000, Delay: 0, Mindelay: 15000, Despawn: 0, Id: sql.NullInt64{Int64: 81464, Valid: true}, Name: sql.NullString{String: "droga_an_angry_goblin464", Valid: true}, Spawn_limit: 0, Dist: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Despawn_timer: 100},
}

package nurga

import (
	"database/sql"
	"github.com/xackery/goeq/spawn"
)

var spawngroups []spawn.SpawnGroup = []spawn.SpawnGroup{
	{Dist: 0.000000, Min_x: 0.000000, Delay: 0, Despawn: 0, Despawn_timer: 100, Mindelay: 15000, Id: sql.NullInt64{Int64: 107003, Valid: true}, Name: sql.NullString{String: "nurga_a_goblin_slavedriver1", Valid: true}, Spawn_limit: 0, Max_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000},
	{Id: sql.NullInt64{Int64: 107006, Valid: true}, Max_y: 0.000000, Delay: 0, Despawn_timer: 100, Name: sql.NullString{String: "nurga_a_goblin_dirttracer6", Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn: 0},
	{Spawn_limit: 0, Max_x: 0.000000, Delay: 0, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 107009, Valid: true}, Name: sql.NullString{String: "nurga_a_goblin_stalagknight9", Valid: true}, Dist: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Mindelay: 15000},
	{Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Min_y: 0.000000, Despawn: 0, Id: sql.NullInt64{Int64: 107013, Valid: true}, Min_x: 0.000000, Max_y: 0.000000, Delay: 0, Mindelay: 15000, Despawn_timer: 100, Name: sql.NullString{String: "nurga_a_goblin_backbiter5", Valid: true}},
	{Despawn_timer: 100, Id: sql.NullInt64{Int64: 107030, Valid: true}, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Min_y: 0.000000, Despawn: 0, Name: sql.NullString{String: "nurga_a_goblin_warder0", Valid: true}, Spawn_limit: 0, Max_y: 0.000000, Delay: 0, Mindelay: 15000},
	{Dist: 0.000000, Min_x: 0.000000, Delay: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 107047, Valid: true}, Spawn_limit: 0, Max_y: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn: 0, Name: sql.NullString{String: "nurga_a_goblin_boneseer47", Valid: true}, Max_x: 0.000000},
	{Dist: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Delay: 0, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 107057, Valid: true}, Name: sql.NullString{String: "nurga_a_goblin_flashdrowser44", Valid: true}, Spawn_limit: 0, Max_x: 0.000000, Mindelay: 15000},
	{Max_x: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Delay: 0, Despawn: 0, Id: sql.NullInt64{Int64: 107058, Valid: true}, Name: sql.NullString{String: "nurga_a_goblin_blightcaller58", Valid: true}, Dist: 0.000000, Despawn_timer: 100, Spawn_limit: 0, Min_y: 0.000000, Mindelay: 15000},
	{Delay: 0, Mindelay: 15000, Name: sql.NullString{String: "nurga_a_goblin_stonechanter26", Valid: true}, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Despawn: 0, Id: sql.NullInt64{Int64: 107077, Valid: true}, Spawn_limit: 0, Despawn_timer: 100},
	{Delay: 0, Despawn: 0, Id: sql.NullInt64{Int64: 107079, Valid: true}, Name: sql.NullString{String: "nurga_an_Iksar_slave79", Valid: true}, Dist: 0.000000, Max_x: 0.000000, Min_y: 0.000000, Spawn_limit: 0, Min_x: 0.000000, Max_y: 0.000000, Mindelay: 15000, Despawn_timer: 100},
	{Id: sql.NullInt64{Int64: 107081, Valid: true}, Spawn_limit: 0, Delay: 0, Despawn: 0, Max_y: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn_timer: 100, Name: sql.NullString{String: "nurga_a_tortured_slave81", Valid: true}, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000},
	{Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Despawn: 0, Id: sql.NullInt64{Int64: 107101, Valid: true}, Name: sql.NullString{String: "nurga_a_sleeping_ogre91", Valid: true}, Min_y: 0.000000, Delay: 0, Mindelay: 15000, Despawn_timer: 100},
}

package cazicthule

import (
	"database/sql"
	"github.com/xackery/goeq/spawn"
)

var spawngroups []spawn.SpawnGroup = []spawn.SpawnGroup{
	{Delay: 0, Id: sql.NullInt64{Int64: 48001, Valid: true}, Name: sql.NullString{String: "cazicthule_a_Tae_Ew_Justicar1", Valid: true}, Max_x: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Spawn_limit: 0, Dist: 0.000000, Mindelay: 15000, Despawn: 0, Despawn_timer: 100},
	{Max_y: 0.000000, Min_y: 0.000000, Delay: 0, Despawn: 0, Id: sql.NullInt64{Int64: 48038, Valid: true}, Name: sql.NullString{String: "cazicthule_a_greenblood_piranha121", Valid: true}, Max_x: 0.000000, Min_x: 0.000000, Mindelay: 15000, Despawn_timer: 100, Spawn_limit: 0, Dist: 0.000000},
	{Despawn: 0, Despawn_timer: 100, Max_x: 0.000000, Max_y: 0.000000, Spawn_limit: 0, Dist: 0.000000, Min_x: 0.000000, Min_y: 0.000000, Delay: 0, Mindelay: 15000, Id: sql.NullInt64{Int64: 48040, Valid: true}, Name: sql.NullString{String: "cazicthule_Gimlik_Cogboggle40", Valid: true}},
	{Max_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Delay: 0, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48041, Valid: true}, Spawn_limit: 0, Min_x: 0.000000, Mindelay: 15000, Name: sql.NullString{String: "cazicthule_a_swirling_ooze0", Valid: true}, Dist: 0.000000},
	{Id: sql.NullInt64{Int64: 48049, Valid: true}, Delay: 0, Name: sql.NullString{String: "cazicthule_a_Tae_Ew_zealot94", Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn: 0, Despawn_timer: 100},
	{Min_x: 0.000000, Mindelay: 15000, Spawn_limit: 0, Name: sql.NullString{String: "cazicthule_a_Thul_Tae_Ew_judicator60", Valid: true}, Dist: 0.000000, Max_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Delay: 0, Despawn: 0, Id: sql.NullInt64{Int64: 48054, Valid: true}, Despawn_timer: 100},
	{Despawn_timer: 100, Id: sql.NullInt64{Int64: 48057, Valid: true}, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Mindelay: 15000, Name: sql.NullString{String: "cazicthule_Tahia_Felwah5", Valid: true}, Spawn_limit: 0, Min_y: 0.000000, Delay: 0, Despawn: 0},
	{Name: sql.NullString{String: "cazicthule_a_Tae_Ew_warrior81", Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Delay: 0, Despawn: 0, Id: sql.NullInt64{Int64: 48058, Valid: true}, Max_y: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn_timer: 100},
	{Spawn_limit: 0, Dist: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Mindelay: 15000, Id: sql.NullInt64{Int64: 48059, Valid: true}, Max_x: 0.000000, Min_x: 0.000000, Delay: 0, Despawn: 0, Despawn_timer: 100, Name: sql.NullString{String: "cazicthule_a_tiger_raptor54", Valid: true}},
	{Mindelay: 15000, Despawn: 0, Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Min_y: 0.000000, Delay: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48060, Valid: true}, Name: sql.NullString{String: "cazicthule_a_Tae_Ew_warder92", Valid: true}, Max_y: 0.000000},
	{Name: sql.NullString{String: "cazicthule_a_jungle_raptor79", Valid: true}, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48062, Valid: true}, Spawn_limit: 0, Max_y: 0.000000, Min_y: 0.000000, Delay: 0, Mindelay: 15000},
	{Delay: 0, Mindelay: 15000, Despawn: 0, Max_y: 0.000000, Min_y: 0.000000, Id: sql.NullInt64{Int64: 48064, Valid: true}, Name: sql.NullString{String: "cazicthule_a_pool_of_slime128", Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Despawn_timer: 100},
	{Max_x: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Delay: 0, Mindelay: 15000, Despawn: 0, Id: sql.NullInt64{Int64: 48065, Valid: true}, Dist: 0.000000, Despawn_timer: 100, Min_y: 0.000000, Name: sql.NullString{String: "cazicthule_a_lifestealer_mosquito136", Valid: true}, Spawn_limit: 0},
	{Max_y: 0.000000, Delay: 0, Mindelay: 15000, Dist: 0.000000, Max_x: 0.000000, Spawn_limit: 0, Min_x: 0.000000, Min_y: 0.000000, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48066, Valid: true}, Name: sql.NullString{String: "cazicthule_an_ooze42", Valid: true}},
	{Mindelay: 15000, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48067, Valid: true}, Name: sql.NullString{String: "cazicthule_a_Tae_Ew_lifedrinker95", Valid: true}, Dist: 0.000000, Max_y: 0.000000, Delay: 0, Spawn_limit: 0, Max_x: 0.000000, Min_x: 0.000000, Min_y: 0.000000},
	{Delay: 0, Mindelay: 15000, Dist: 0.000000, Max_x: 0.000000, Max_y: 0.000000, Min_x: 0.000000, Min_y: 0.000000, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48069, Valid: true}, Name: sql.NullString{String: "cazicthule_a_fierce_jungle_raptor51", Valid: true}, Spawn_limit: 0},
	{Max_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Mindelay: 15000, Id: sql.NullInt64{Int64: 48070, Valid: true}, Spawn_limit: 0, Dist: 0.000000, Despawn: 0, Despawn_timer: 100, Name: sql.NullString{String: "cazicthule_a_Tae_Ew_proselyte90", Valid: true}, Min_x: 0.000000, Delay: 0},
	{Spawn_limit: 0, Max_x: 0.000000, Min_x: 0.000000, Min_y: 0.000000, Delay: 0, Mindelay: 15000, Despawn_timer: 100, Name: sql.NullString{String: "cazicthule_a_Tae_Ew_disciple93", Valid: true}, Dist: 0.000000, Max_y: 0.000000, Despawn: 0, Id: sql.NullInt64{Int64: 48071, Valid: true}},
	{Delay: 0, Despawn: 0, Id: sql.NullInt64{Int64: 48073, Valid: true}, Name: sql.NullString{String: "cazicthule_a_jungle_stalker73", Valid: true}, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Min_y: 0.000000, Spawn_limit: 0, Max_y: 0.000000, Mindelay: 15000, Despawn_timer: 100},
	{Despawn_timer: 100, Id: sql.NullInt64{Int64: 48074, Valid: true}, Name: sql.NullString{String: "cazicthule_a_large_greenblood_piranha120", Valid: true}, Dist: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Delay: 0, Despawn: 0, Spawn_limit: 0, Max_x: 0.000000, Min_x: 0.000000, Mindelay: 15000},
	{Max_x: 0.000000, Delay: 0, Min_y: 0.000000, Mindelay: 15000, Id: sql.NullInt64{Int64: 48077, Valid: true}, Name: sql.NullString{String: "cazicthule_a_frenzied_tiger_raptor116", Valid: true}, Spawn_limit: 0, Dist: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Despawn: 0, Despawn_timer: 100},
	{Id: sql.NullInt64{Int64: 48079, Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Despawn: 0, Name: sql.NullString{String: "cazicthule_a_Tae_Ew_defender59", Valid: true}, Max_x: 0.000000, Min_x: 0.000000, Delay: 0, Mindelay: 15000, Despawn_timer: 100},
	{Min_y: 0.000000, Delay: 0, Despawn: 0, Id: sql.NullInt64{Int64: 48081, Valid: true}, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Despawn_timer: 100, Name: sql.NullString{String: "cazicthule_an_Amygdalan_protector34", Valid: true}, Spawn_limit: 0, Mindelay: 15000},
	{Min_x: 0.000000, Mindelay: 15000, Despawn_timer: 100, Spawn_limit: 0, Name: sql.NullString{String: "cazicthule_a_Thul_Tae_Ew_defender58", Valid: true}, Dist: 0.000000, Max_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Delay: 0, Despawn: 0, Id: sql.NullInt64{Int64: 48083, Valid: true}},
	{Name: sql.NullString{String: "cazicthule_a_Thul_Tae_Ew_justicar57", Valid: true}, Spawn_limit: 0, Max_x: 0.000000, Delay: 0, Despawn: 0, Id: sql.NullInt64{Int64: 48084, Valid: true}, Min_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn_timer: 100, Dist: 0.000000},
	{Mindelay: 15000, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48087, Valid: true}, Name: sql.NullString{String: "cazicthule_a_Tae_Ew_fanatic184", Valid: true}, Max_x: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Spawn_limit: 0, Dist: 0.000000, Delay: 0, Despawn: 0},
	{Name: sql.NullString{String: "cazicthule_a_Thul_Tae_Ew_zealot191", Valid: true}, Spawn_limit: 0, Delay: 0, Mindelay: 15000, Despawn: 0, Min_y: 0.000000, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48089, Valid: true}, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Max_y: 0.000000},
	{Max_y: 0.000000, Min_y: 0.000000, Despawn: 0, Id: sql.NullInt64{Int64: 48090, Valid: true}, Max_x: 0.000000, Min_x: 0.000000, Delay: 0, Mindelay: 15000, Despawn_timer: 100, Name: sql.NullString{String: "cazicthule_a_fierce_tiger_raptor52", Valid: true}, Spawn_limit: 0, Dist: 0.000000},
	{Despawn: 0, Spawn_limit: 0, Min_x: 0.000000, Min_y: 0.000000, Mindelay: 15000, Max_y: 0.000000, Delay: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48093, Valid: true}, Name: sql.NullString{String: "cazicthule_#a_Thul_Tae_Ew_High_Priest6", Valid: true}, Dist: 0.000000, Max_x: 0.000000},
	{Id: sql.NullInt64{Int64: 48094, Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Delay: 0, Despawn: 0, Despawn_timer: 100, Name: sql.NullString{String: "cazicthule_an_enraged_Tae_Ew_fanatic29", Valid: true}, Min_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Mindelay: 15000},
	{Max_y: 0.000000, Despawn: 0, Despawn_timer: 100, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Min_y: 0.000000, Delay: 0, Mindelay: 15000, Id: sql.NullInt64{Int64: 48095, Valid: true}, Name: sql.NullString{String: "cazicthule_a_living_void38", Valid: true}, Spawn_limit: 0},
	{Min_x: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48097, Valid: true}, Name: sql.NullString{String: "cazicthule_#an_envenomed_hunter97", Valid: true}, Dist: 0.000000, Max_x: 0.000000, Spawn_limit: 0, Max_y: 0.000000, Delay: 0, Despawn: 0},
	{Dist: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Delay: 0, Mindelay: 15000, Despawn: 0, Spawn_limit: 0, Despawn_timer: 100, Name: sql.NullString{String: "cazicthule_a_Tae_Ew_lifestealer61", Valid: true}, Max_x: 0.000000, Id: sql.NullInt64{Int64: 48098, Valid: true}},
	{Id: sql.NullInt64{Int64: 48099, Valid: true}, Max_y: 0.000000, Delay: 0, Despawn: 0, Min_x: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn_timer: 100, Name: sql.NullString{String: "cazicthule_a_Thul_Tae_Ew_Warder69", Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000},
	{Id: sql.NullInt64{Int64: 48100, Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Min_y: 0.000000, Delay: 0, Name: sql.NullString{String: "cazicthule_a_Tae_Ew_spiritcaller110", Valid: true}, Min_x: 0.000000, Max_y: 0.000000, Mindelay: 15000, Despawn: 0, Despawn_timer: 100},
	{Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Min_y: 0.000000, Delay: 0, Mindelay: 15000, Id: sql.NullInt64{Int64: 48102, Valid: true}, Name: sql.NullString{String: "cazicthule_a_Tae_Ew_protector171", Valid: true}, Min_x: 0.000000, Max_y: 0.000000, Despawn: 0, Despawn_timer: 100},
	{Dist: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48103, Valid: true}, Name: sql.NullString{String: "cazicthule_a_Thul_Tae_Ew_bloodcaller168", Valid: true}, Spawn_limit: 0, Max_x: 0.000000, Min_y: 0.000000, Delay: 0, Mindelay: 15000},
	{Min_x: 0.000000, Max_y: 0.000000, Delay: 0, Despawn: 0, Dist: 0.000000, Max_x: 0.000000, Spawn_limit: 0, Min_y: 0.000000, Mindelay: 15000, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48104, Valid: true}, Name: sql.NullString{String: "cazicthule__________203", Valid: true}},
	{Mindelay: 15000, Despawn_timer: 100, Spawn_limit: 0, Max_y: 0.000000, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Min_y: 0.000000, Delay: 0, Despawn: 0, Id: sql.NullInt64{Int64: 48106, Valid: true}, Name: sql.NullString{String: "cazicthule_a_Thul_Tae_Ew_lifestealer251", Valid: true}},
	{Name: sql.NullString{String: "cazicthule_a_frenzied_jungle_raptor53", Valid: true}, Dist: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48107, Valid: true}, Spawn_limit: 0, Max_x: 0.000000, Min_x: 0.000000, Delay: 0},
	{Delay: 0, Mindelay: 15000, Name: sql.NullString{String: "cazicthule_a_Thul_Tae_Ew_fanatic165", Valid: true}, Spawn_limit: 0, Min_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48109, Valid: true}, Dist: 0.000000, Max_x: 0.000000},
	{Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48110, Valid: true}, Dist: 0.000000, Max_x: 0.000000, Min_y: 0.000000, Delay: 0, Mindelay: 15000, Name: sql.NullString{String: "cazicthule_a_shiverback295", Valid: true}, Spawn_limit: 0, Min_x: 0.000000, Max_y: 0.000000},
	{Despawn_timer: 100, Id: sql.NullInt64{Int64: 48113, Valid: true}, Name: sql.NullString{String: "cazicthule_a_Tae_Ew_judicator174", Valid: true}, Dist: 0.000000, Min_x: 0.000000, Despawn: 0, Mindelay: 15000, Spawn_limit: 0, Max_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Delay: 0},
	{Spawn_limit: 0, Dist: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn_timer: 100, Name: sql.NullString{String: "cazicthule_a_jungle_hunter77", Valid: true}, Max_x: 0.000000, Min_x: 0.000000, Delay: 0, Despawn: 0, Id: sql.NullInt64{Int64: 48116, Valid: true}},
	{Id: sql.NullInt64{Int64: 48117, Valid: true}, Dist: 0.000000, Max_x: 0.000000, Max_y: 0.000000, Delay: 0, Mindelay: 15000, Despawn: 0, Name: sql.NullString{String: "cazicthule_a_large_hunter83", Valid: true}, Spawn_limit: 0, Min_x: 0.000000, Min_y: 0.000000, Despawn_timer: 100},
	{Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Delay: 0, Despawn: 0, Name: sql.NullString{String: "cazicthule_DlgFour24", Valid: true}, Min_y: 0.000000, Mindelay: 15000, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48120, Valid: true}},
	{Delay: 0, Despawn_timer: 100, Name: sql.NullString{String: "cazicthule_DlgTwo26", Valid: true}, Dist: 0.000000, Min_x: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn: 0, Id: sql.NullInt64{Int64: 48123, Valid: true}, Spawn_limit: 0, Max_x: 0.000000, Max_y: 0.000000},
	{Despawn_timer: 100, Id: sql.NullInt64{Int64: 48126, Valid: true}, Name: sql.NullString{String: "cazicthule_DlgSeven21", Valid: true}, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Delay: 0, Mindelay: 15000, Spawn_limit: 0, Max_y: 0.000000, Min_y: 0.000000, Despawn: 0},
	{Name: sql.NullString{String: "cazicthule_InterpreterZero8", Valid: true}, Spawn_limit: 0, Max_x: 0.000000, Min_x: 0.000000, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48129, Valid: true}, Dist: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Delay: 0, Mindelay: 15000, Despawn: 0},
	{Mindelay: 15000, Despawn_timer: 100, Min_y: 0.000000, Name: sql.NullString{String: "cazicthule_DlgEight20", Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Delay: 0, Id: sql.NullInt64{Int64: 48131, Valid: true}, Despawn: 0},
	{Max_x: 0.000000, Min_x: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48132, Valid: true}, Name: sql.NullString{String: "cazicthule_an_unseen_force65", Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_y: 0.000000, Delay: 0},
	{Max_x: 0.000000, Min_x: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48133, Valid: true}, Name: sql.NullString{String: "cazicthule_RaidStopper7", Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_y: 0.000000, Delay: 0},
	{Max_y: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48135, Valid: true}, Min_x: 0.000000, Dist: 0.000000, Max_x: 0.000000, Delay: 0, Despawn: 0, Name: sql.NullString{String: "cazicthule_DlgTwelve16", Valid: true}, Spawn_limit: 0},
	{Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Delay: 0, Despawn: 0, Id: sql.NullInt64{Int64: 48138, Valid: true}, Name: sql.NullString{String: "cazicthule_DlgNineteen9", Valid: true}, Max_y: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn_timer: 100},
	{Dist: 0.000000, Max_x: 0.000000, Max_y: 0.000000, Delay: 0, Mindelay: 15000, Id: sql.NullInt64{Int64: 48140, Valid: true}, Spawn_limit: 0, Min_x: 0.000000, Min_y: 0.000000, Despawn: 0, Despawn_timer: 100, Name: sql.NullString{String: "cazicthule_DlgSix22", Valid: true}},
	{Name: sql.NullString{String: "cazicthule_DlgEighteen10", Valid: true}, Spawn_limit: 0, Delay: 0, Despawn_timer: 100, Min_y: 0.000000, Mindelay: 15000, Despawn: 0, Id: sql.NullInt64{Int64: 48142, Valid: true}, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Max_y: 0.000000},
	{Id: sql.NullInt64{Int64: 48144, Valid: true}, Name: sql.NullString{String: "cazicthule_DlgEleven17", Valid: true}, Spawn_limit: 0, Max_x: 0.000000, Max_y: 0.000000, Delay: 0, Despawn: 0, Dist: 0.000000, Min_x: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn_timer: 100},
	{Id: sql.NullInt64{Int64: 48146, Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Delay: 0, Name: sql.NullString{String: "cazicthule_DlgFive23", Valid: true}, Min_y: 0.000000, Mindelay: 15000, Despawn: 0, Despawn_timer: 100},
	{Min_x: 0.000000, Mindelay: 15000, Despawn_timer: 100, Name: sql.NullString{String: "cazicthule_DlgSeventeen11", Valid: true}, Dist: 0.000000, Max_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000, Delay: 0, Despawn: 0, Id: sql.NullInt64{Int64: 48154, Valid: true}, Spawn_limit: 0},
	{Spawn_limit: 0, Max_y: 0.000000, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Min_y: 0.000000, Delay: 0, Mindelay: 15000, Id: sql.NullInt64{Int64: 48157, Valid: true}, Name: sql.NullString{String: "cazicthule_DlgFifteen13", Valid: true}, Despawn: 0, Despawn_timer: 100},
	{Min_y: 0.000000, Mindelay: 15000, Despawn_timer: 100, Min_x: 0.000000, Name: sql.NullString{String: "cazicthule_DlgThirteen15", Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Max_y: 0.000000, Delay: 0, Despawn: 0, Id: sql.NullInt64{Int64: 48163, Valid: true}},
	{Max_y: 0.000000, Min_y: 0.000000, Delay: 0, Despawn: 0, Name: sql.NullString{String: "cazicthule_Mazeman39", Valid: true}, Max_x: 0.000000, Min_x: 0.000000, Mindelay: 15000, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48164, Valid: true}, Spawn_limit: 0, Dist: 0.000000},
	{Min_y: 0.000000, Mindelay: 15000, Despawn: 0, Name: sql.NullString{String: "cazicthule_DlgSixteen12", Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Id: sql.NullInt64{Int64: 48173, Valid: true}, Max_y: 0.000000, Delay: 0, Despawn_timer: 100},
	{Name: sql.NullString{String: "cazicthule_#an_enraged_jungle_raptor75", Valid: true}, Spawn_limit: 0, Max_y: 0.000000, Min_y: 0.000000, Mindelay: 15000, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48177, Valid: true}, Max_x: 0.000000, Min_x: 0.000000, Delay: 0, Dist: 0.000000},
	{Id: sql.NullInt64{Int64: 48182, Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_y: 0.000000, Delay: 0, Mindelay: 15000, Despawn: 0, Despawn_timer: 100, Name: sql.NullString{String: "cazicthule_#an_enraged_tiger_raptor86", Valid: true}, Max_x: 0.000000, Min_x: 0.000000, Min_y: 0.000000},
	{Id: sql.NullInt64{Int64: 48187, Valid: true}, Max_x: 0.000000, Delay: 0, Mindelay: 15000, Despawn: 0, Despawn_timer: 100, Name: sql.NullString{String: "cazicthule_DlgTen18", Valid: true}, Spawn_limit: 0, Dist: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Min_y: 0.000000},
	{Min_y: 0.000000, Id: sql.NullInt64{Int64: 48188, Valid: true}, Name: sql.NullString{String: "cazicthule_DlgZero28", Valid: true}, Max_x: 0.000000, Max_y: 0.000000, Delay: 0, Mindelay: 15000, Despawn: 0, Despawn_timer: 100, Spawn_limit: 0, Dist: 0.000000, Min_x: 0.000000},
	{Name: sql.NullString{String: "cazicthule___194", Valid: true}, Spawn_limit: 0, Min_x: 0.000000, Max_y: 0.000000, Despawn_timer: 100, Despawn: 0, Id: sql.NullInt64{Int64: 48194, Valid: true}, Dist: 0.000000, Max_x: 0.000000, Min_y: 0.000000, Delay: 0, Mindelay: 15000},
	{Name: sql.NullString{String: "cazicthule_DlgNine19", Valid: true}, Dist: 0.000000, Min_y: 0.000000, Id: sql.NullInt64{Int64: 48196, Valid: true}, Spawn_limit: 0, Max_x: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Delay: 0, Mindelay: 15000, Despawn: 0, Despawn_timer: 100},
	{Delay: 0, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48197, Valid: true}, Spawn_limit: 0, Max_x: 0.000000, Max_y: 0.000000, Mindelay: 15000, Name: sql.NullString{String: "cazicthule_DlgOne27", Valid: true}, Dist: 0.000000, Min_x: 0.000000, Min_y: 0.000000},
	{Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48199, Valid: true}, Max_x: 0.000000, Min_x: 0.000000, Max_y: 0.000000, Mindelay: 15000, Name: sql.NullString{String: "cazicthule_DlgThree25", Valid: true}, Spawn_limit: 0, Dist: 0.000000, Min_y: 0.000000, Delay: 0},
	{Name: sql.NullString{String: "cazicthule_a_Tae_Ew_bloodcaller170", Valid: true}, Max_y: 0.000000, Mindelay: 15000, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48201, Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Min_y: 0.000000, Delay: 0},
	{Spawn_limit: 0, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Min_y: 0.000000, Mindelay: 15000, Name: sql.NullString{String: "cazicthule_DlgFourteen14", Valid: true}, Max_y: 0.000000, Delay: 0, Despawn: 0, Despawn_timer: 100, Id: sql.NullInt64{Int64: 48208, Valid: true}},
	{Id: sql.NullInt64{Int64: 48274, Valid: true}, Spawn_limit: 0, Dist: 0.000000, Max_y: 0.000000, Delay: 0, Mindelay: 15000, Name: sql.NullString{String: "cazicthule_#a_Tae_Ew_prophet274", Valid: true}, Max_x: 0.000000, Min_x: 0.000000, Min_y: 0.000000, Despawn: 0, Despawn_timer: 100},
	{Id: sql.NullInt64{Int64: 48298, Valid: true}, Name: sql.NullString{String: "cazicthule__196", Valid: true}, Dist: 0.000000, Max_x: 0.000000, Min_x: 0.000000, Despawn: 0, Spawn_limit: 0, Max_y: 0.000000, Min_y: 0.000000, Delay: 0, Mindelay: 15000, Despawn_timer: 100},
}

package veeshan

import (
	"database/sql"
	"github.com/xackery/goeq/spawn"
)

var spawns []spawn.Spawn2 = []spawn.Spawn2{
	{X: 614.000000, Variance: 0, Pathgrid: 0, Respawntime: 1200, Condition: 0, Cond_value: 1, Spawngroupid: 108067, Version: 0, Y: 1537.000000, Heading: 12.237960, Zone: sql.NullString{String: "veeshan", Valid: true}, Enabled: 1, Id: sql.NullInt64{Int64: 364656, Valid: true}, Z: 261.200012, Animation: 0},
	{Spawngroupid: 108067, Version: 0, X: 605.000000, Heading: 146.855499, Pathgrid: 0, Respawntime: 1200, Enabled: 1, Animation: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, Y: 1379.000000, Z: 261.000000, Condition: 0, Cond_value: 1, Id: sql.NullInt64{Int64: 364657, Valid: true}, Variance: 0},
	{Z: 260.899994, Respawntime: 1200, Variance: 0, Condition: 0, Animation: 0, Id: sql.NullInt64{Int64: 364658, Valid: true}, Zone: sql.NullString{String: "veeshan", Valid: true}, Y: 1380.000000, Cond_value: 1, Version: 0, X: 501.000000, Spawngroupid: 108067, Heading: 141.756393, Pathgrid: 0, Enabled: 1},
	{Condition: 0, Cond_value: 1, Animation: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, Version: 0, Z: 688.400024, Heading: 65.269119, Id: sql.NullInt64{Int64: 364596, Valid: true}, Spawngroupid: 108500, X: -479.000000, Variance: 0, Y: 1430.000000, Respawntime: 144, Pathgrid: 0, Enabled: 1},
	{Zone: sql.NullString{String: "veeshan", Valid: true}, Y: 1370.000000, Variance: 0, X: -479.000000, Enabled: 1, Id: sql.NullInt64{Int64: 364597, Valid: true}, Spawngroupid: 108500, Pathgrid: 0, Condition: 0, Animation: 0, Version: 0, Z: 688.400024, Heading: 65.269119, Respawntime: 144, Cond_value: 1},
	{Version: 0, Heading: 259.036804, Spawngroupid: 108500, Z: 288.600006, Variance: 0, Enabled: 1, X: -970.000000, Y: -647.000000, Respawntime: 144, Condition: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, Pathgrid: 0, Cond_value: 1, Animation: 0, Id: sql.NullInt64{Int64: 364627, Valid: true}},
	{Id: sql.NullInt64{Int64: 364628, Valid: true}, Version: 0, Condition: 0, Animation: 0, Y: -649.000000, Z: 288.600006, Heading: 259.036804, Respawntime: 144, Variance: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, X: -1030.000000, Spawngroupid: 108500, Pathgrid: 0, Cond_value: 1, Enabled: 1},
	{Spawngroupid: 108500, Variance: 0, Pathgrid: 0, Animation: 0, Id: sql.NullInt64{Int64: 364629, Valid: true}, Heading: 64.249290, Enabled: 1, Zone: sql.NullString{String: "veeshan", Valid: true}, Version: 0, Y: 342.000000, Condition: 0, X: -893.000000, Z: 468.500000, Respawntime: 144, Cond_value: 1},
	{Version: 0, X: -891.000000, Z: 468.500000, Respawntime: 144, Condition: 0, Enabled: 1, Id: sql.NullInt64{Int64: 364630, Valid: true}, Heading: 62.209629, Pathgrid: 0, Cond_value: 1, Animation: 0, Spawngroupid: 108500, Zone: sql.NullString{String: "veeshan", Valid: true}, Y: 284.000000, Variance: 0},
	{X: 574.000000, Y: 1779.000000, Pathgrid: 0, Enabled: 1, Id: sql.NullInt64{Int64: 364635, Valid: true}, Version: 0, Z: 609.500000, Respawntime: 144, Condition: 0, Animation: 0, Spawngroupid: 108500, Heading: 127.478798, Zone: sql.NullString{String: "veeshan", Valid: true}, Variance: 0, Cond_value: 1},
	{Z: 609.500000, Respawntime: 144, Enabled: 1, Animation: 0, Heading: 127.478798, Pathgrid: 0, Cond_value: 1, Zone: sql.NullString{String: "veeshan", Valid: true}, Version: 0, Y: 1782.000000, Variance: 0, Id: sql.NullInt64{Int64: 364636, Valid: true}, Spawngroupid: 108500, X: 614.000000, Condition: 0},
	{Z: 450.899994, Pathgrid: 0, Animation: 0, Cond_value: 1, Spawngroupid: 108500, Zone: sql.NullString{String: "veeshan", Valid: true}, X: -155.000000, Respawntime: 144, Variance: 0, Version: 0, Id: sql.NullInt64{Int64: 364637, Valid: true}, Y: 1336.000000, Heading: 4.079320, Condition: 0, Enabled: 1},
	{Id: sql.NullInt64{Int64: 364638, Valid: true}, Spawngroupid: 108500, Z: 448.500000, Heading: 1.019830, Zone: sql.NullString{String: "veeshan", Valid: true}, Version: 0, Variance: 0, Pathgrid: 0, Condition: 0, Cond_value: 1, Enabled: 1, X: -86.000000, Y: 1358.000000, Respawntime: 144, Animation: 0},
	{Heading: 65.269119, Respawntime: 144, Condition: 0, Enabled: 1, Y: 1467.000000, Id: sql.NullInt64{Int64: 364640, Valid: true}, Z: 448.500000, Variance: 0, Pathgrid: 0, Spawngroupid: 108500, Version: 0, X: -9.000000, Cond_value: 1, Animation: 0, Zone: sql.NullString{String: "veeshan", Valid: true}},
	{Version: 0, Pathgrid: 0, Cond_value: 1, Animation: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, Spawngroupid: 108500, Condition: 0, Id: sql.NullInt64{Int64: 364641, Valid: true}, Y: 1408.000000, Respawntime: 144, X: -10.000000, Heading: 65.269119, Variance: 0, Enabled: 1, Z: 448.500000},
	{X: 211.000000, Y: 1719.000000, Enabled: 1, Zone: sql.NullString{String: "veeshan", Valid: true}, Version: 0, Respawntime: 144, Pathgrid: 0, Z: 608.500000, Heading: 116.260597, Condition: 0, Cond_value: 1, Spawngroupid: 108500, Variance: 0, Id: sql.NullInt64{Int64: 364642, Valid: true}, Animation: 0},
	{Enabled: 1, Y: 1722.000000, Heading: 115.240799, Variance: 0, Pathgrid: 0, Cond_value: 1, Condition: 0, Id: sql.NullInt64{Int64: 364643, Valid: true}, Spawngroupid: 108500, Zone: sql.NullString{String: "veeshan", Valid: true}, Version: 0, X: 270.000000, Z: 608.500000, Respawntime: 144, Animation: 0},
	{Z: 608.400024, Respawntime: 144, Animation: 0, Id: sql.NullInt64{Int64: 364646, Valid: true}, Spawngroupid: 108500, Zone: sql.NullString{String: "veeshan", Valid: true}, Condition: 0, Enabled: 1, Version: 0, Cond_value: 1, X: 281.000000, Y: 1291.000000, Heading: 46.912182, Variance: 0, Pathgrid: 0},
	{X: 285.000000, Heading: 59.150139, Respawntime: 144, Condition: 0, Id: sql.NullInt64{Int64: 364647, Valid: true}, Zone: sql.NullString{String: "veeshan", Valid: true}, Animation: 0, Enabled: 1, Variance: 0, Pathgrid: 0, Cond_value: 1, Spawngroupid: 108500, Version: 0, Y: 1350.000000, Z: 607.799988},
	{Zone: sql.NullString{String: "veeshan", Valid: true}, Version: 0, Variance: 0, Animation: 0, Id: sql.NullInt64{Int64: 364648, Valid: true}, Spawngroupid: 108500, X: -65.000000, Heading: 194.787506, Cond_value: 1, Z: 448.500000, Respawntime: 144, Condition: 0, Enabled: 1, Y: 809.000000, Pathgrid: 0},
	{Condition: 0, Enabled: 1, Z: 448.500000, Variance: 0, Pathgrid: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, Version: 0, Y: 869.000000, Respawntime: 144, Id: sql.NullInt64{Int64: 364649, Valid: true}, Spawngroupid: 108500, X: -69.000000, Heading: 197.847000, Cond_value: 1, Animation: 0},
	{Y: -247.000000, Z: 428.500000, Respawntime: 144, Condition: 0, Version: 0, Heading: 128.498596, Zone: sql.NullString{String: "veeshan", Valid: true}, X: -152.000000, Animation: 0, Id: sql.NullInt64{Int64: 364650, Valid: true}, Variance: 0, Pathgrid: 0, Cond_value: 1, Enabled: 1, Spawngroupid: 108500},
	{Heading: 131.558105, Respawntime: 144, Variance: 0, Pathgrid: 0, Condition: 0, Enabled: 1, X: -89.000000, Y: -245.000000, Animation: 0, Id: sql.NullInt64{Int64: 364651, Valid: true}, Version: 0, Spawngroupid: 108500, Cond_value: 1, Zone: sql.NullString{String: "veeshan", Valid: true}, Z: 428.500000},
	{Z: 368.500000, Pathgrid: 0, Condition: 0, Cond_value: 1, Zone: sql.NullString{String: "veeshan", Valid: true}, Y: 447.000000, Version: 0, Animation: 0, Heading: 260.056702, Respawntime: 144, X: 93.000000, Variance: 0, Enabled: 1, Id: sql.NullInt64{Int64: 364652, Valid: true}, Spawngroupid: 108500},
	{Zone: sql.NullString{String: "veeshan", Valid: true}, Enabled: 1, Animation: 0, Id: sql.NullInt64{Int64: 364653, Valid: true}, Z: 368.500000, Respawntime: 144, Variance: 0, Spawngroupid: 108500, Y: 447.000000, Heading: 260.056702, Cond_value: 1, Version: 0, X: 154.000000, Pathgrid: 0, Condition: 0},
	{Pathgrid: 0, Spawngroupid: 108500, Y: 445.000000, Heading: 258.016998, Respawntime: 144, Condition: 0, Version: 0, X: 205.000000, Cond_value: 1, Enabled: 1, Id: sql.NullInt64{Int64: 364654, Valid: true}, Zone: sql.NullString{String: "veeshan", Valid: true}, Z: 368.500000, Variance: 0, Animation: 0},
	{X: 274.000000, Y: 442.000000, Respawntime: 144, Cond_value: 1, Animation: 0, Condition: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, Version: 0, Z: 368.500000, Variance: 0, Pathgrid: 0, Spawngroupid: 108500, Enabled: 1, Id: sql.NullInt64{Int64: 364655, Valid: true}, Heading: 256.997192},
	{Spawngroupid: 108500, Z: 207.600006, Respawntime: 144, Pathgrid: 0, X: 1217.000000, Y: 841.000000, Condition: 0, Cond_value: 1, Id: sql.NullInt64{Int64: 364660, Valid: true}, Zone: sql.NullString{String: "veeshan", Valid: true}, Heading: 78.526909, Variance: 0, Enabled: 1, Version: 0, Animation: 0},
	{Y: 755.000000, Pathgrid: 0, Cond_value: 1, Spawngroupid: 108500, Enabled: 1, Zone: sql.NullString{String: "veeshan", Valid: true}, Version: 0, X: 1213.000000, Z: 207.600006, Heading: 48.951839, Respawntime: 144, Variance: 0, Id: sql.NullInt64{Int64: 364661, Valid: true}, Animation: 0, Condition: 0},
	{Cond_value: 1, Id: sql.NullInt64{Int64: 364662, Valid: true}, Version: 0, Z: 48.700001, Respawntime: 144, Spawngroupid: 108500, Zone: sql.NullString{String: "veeshan", Valid: true}, Y: 893.000000, Pathgrid: 0, X: 1719.000000, Heading: 126.458900, Variance: 0, Condition: 0, Enabled: 1, Animation: 0},
	{Y: 801.000000, Respawntime: 144, Spawngroupid: 108500, Zone: sql.NullString{String: "veeshan", Valid: true}, Pathgrid: 0, Condition: 0, Cond_value: 1, Enabled: 1, Animation: 0, Version: 0, Heading: 198.866898, Id: sql.NullInt64{Int64: 364663, Valid: true}, Variance: 0, X: 1820.000000, Z: 48.700001},
	{Spawngroupid: 108500, Heading: 110.141602, Pathgrid: 0, Cond_value: 1, Version: 0, Enabled: 1, Y: 376.000000, Z: 28.750990, Respawntime: 144, Condition: 0, Animation: 0, Id: sql.NullInt64{Int64: 364664, Valid: true}, Zone: sql.NullString{String: "veeshan", Valid: true}, X: 1649.000000, Variance: 0},
	{Zone: sql.NullString{String: "veeshan", Valid: true}, Z: 28.750990, Heading: 149.914993, Pathgrid: 0, Version: 0, X: 1754.000000, Cond_value: 1, Id: sql.NullInt64{Int64: 364665, Valid: true}, Spawngroupid: 108500, Variance: 0, Condition: 0, Animation: 0, Y: 370.000000, Respawntime: 144, Enabled: 1},
	{Version: 0, Y: 1604.000000, Heading: 65.269119, Respawntime: 900, Variance: 0, Id: sql.NullInt64{Int64: 364592, Valid: true}, X: -412.000000, Pathgrid: 0, Enabled: 1, Spawngroupid: 108505, Cond_value: 1, Animation: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, Z: 692.200012, Condition: 0},
	{Heading: 130.538193, Variance: 0, Animation: 0, Cond_value: 1, Spawngroupid: 108505, Y: -218.000000, Pathgrid: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, X: -177.000000, Enabled: 1, Respawntime: 900, Condition: 0, Id: sql.NullInt64{Int64: 364621, Valid: true}, Version: 0, Z: 449.299988},
	{Y: -224.000000, Respawntime: 900, Zone: sql.NullString{String: "veeshan", Valid: true}, X: -59.000000, Pathgrid: 0, Spawngroupid: 108505, Variance: 0, Animation: 0, Heading: 129.518402, Enabled: 1, Z: 449.299988, Condition: 0, Cond_value: 1, Id: sql.NullInt64{Int64: 364622, Valid: true}, Version: 0},
	{Pathgrid: 0, Condition: 0, Z: 292.299988, Y: -999.000000, Enabled: 1, Zone: sql.NullString{String: "veeshan", Valid: true}, Version: 0, X: -1136.000000, Respawntime: 900, Animation: 0, Spawngroupid: 108505, Heading: 256.997192, Variance: 0, Cond_value: 1, Id: sql.NullInt64{Int64: 364624, Valid: true}},
	{Spawngroupid: 108505, Z: 292.299988, Pathgrid: 0, Condition: 0, Animation: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, Version: 0, Y: -1003.000000, Enabled: 1, Id: sql.NullInt64{Int64: 364625, Valid: true}, Heading: 5.099150, Variance: 0, Cond_value: 1, X: -1189.000000, Respawntime: 900},
	{Cond_value: 1, Enabled: 1, Id: sql.NullInt64{Int64: 364626, Valid: true}, Heading: 41.813030, Respawntime: 900, Pathgrid: 0, Spawngroupid: 108505, Zone: sql.NullString{String: "veeshan", Valid: true}, Z: 292.299988, Variance: 0, Condition: 0, Animation: 0, Version: 0, X: -1151.000000, Y: -789.000000},
	{Enabled: 1, Version: 0, Pathgrid: 0, Cond_value: 1, Y: 1174.000000, Heading: 63.229462, Respawntime: 900, Animation: 0, Spawngroupid: 108505, Zone: sql.NullString{String: "veeshan", Valid: true}, X: 1216.000000, Id: sql.NullInt64{Int64: 364632, Valid: true}, Z: 772.099976, Variance: 0, Condition: 0},
	{Version: 0, X: 1376.000000, Y: 1269.000000, Condition: 0, Spawngroupid: 108505, Zone: sql.NullString{String: "veeshan", Valid: true}, Heading: 49.971668, Cond_value: 1, Z: 772.099976, Variance: 0, Pathgrid: 0, Animation: 0, Id: sql.NullInt64{Int64: 364633, Valid: true}, Respawntime: 900, Enabled: 1},
	{Zone: sql.NullString{String: "veeshan", Valid: true}, Pathgrid: 0, Enabled: 1, Animation: 0, Id: sql.NullInt64{Int64: 364634, Valid: true}, Y: 1216.000000, Variance: 0, Cond_value: 1, Version: 0, X: 1369.000000, Z: 772.099976, Respawntime: 900, Spawngroupid: 108505, Heading: 26.515579, Condition: 0},
	{Id: sql.NullInt64{Int64: 364644, Valid: true}, X: -564.000000, Heading: 239.660004, Variance: 0, Cond_value: 1, Animation: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, Respawntime: 900, Pathgrid: 0, Enabled: 1, Spawngroupid: 108505, Version: 0, Y: 1274.000000, Z: 685.400024, Condition: 0},
	{Y: 1735.000000, Variance: 0, Cond_value: 1, Spawngroupid: 108506, Version: 0, X: -118.000000, Heading: 255.977295, Condition: 0, Enabled: 1, Animation: 0, Id: sql.NullInt64{Int64: 364591, Valid: true}, Z: 697.900024, Pathgrid: 90, Zone: sql.NullString{String: "veeshan", Valid: true}, Respawntime: 144},
	{Y: 840.000000, Respawntime: 144, Cond_value: 1, X: -62.000000, Heading: 64.249290, Variance: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, Version: 0, Z: 459.799988, Condition: 0, Animation: 0, Id: sql.NullInt64{Int64: 364593, Valid: true}, Spawngroupid: 108506, Pathgrid: 84, Enabled: 1},
	{X: -1099.000000, Variance: 0, Cond_value: 1, Version: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, Z: 282.000000, Respawntime: 144, Pathgrid: 97, Enabled: 1, Animation: 0, Id: sql.NullInt64{Int64: 364594, Valid: true}, Heading: 60.000000, Y: -760.000000, Condition: 0, Spawngroupid: 108506},
	{Cond_value: 1, Id: sql.NullInt64{Int64: 364595, Valid: true}, Z: 322.000000, Variance: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, Y: -405.000000, Respawntime: 144, Condition: 0, Enabled: 1, Spawngroupid: 108506, Version: 0, X: -1315.000000, Heading: 1.000000, Pathgrid: 96, Animation: 0},
	{Zone: sql.NullString{String: "veeshan", Valid: true}, Heading: 253.000000, Condition: 0, Cond_value: 1, Spawngroupid: 108506, Version: 0, Z: 761.799988, Variance: 0, Pathgrid: 95, Id: sql.NullInt64{Int64: 364599, Valid: true}, Enabled: 1, Animation: 0, X: 1441.000000, Y: 1304.000000, Respawntime: 144},
	{Zone: sql.NullString{String: "veeshan", Valid: true}, Respawntime: 144, Id: sql.NullInt64{Int64: 364600, Valid: true}, Z: 681.900024, Variance: 0, Condition: 0, Enabled: 1, Animation: 0, Version: 0, Y: 1734.000000, Pathgrid: 94, Cond_value: 1, Spawngroupid: 108506, X: 1060.000000, Heading: 251.000000},
	{Zone: sql.NullString{String: "veeshan", Valid: true}, Heading: 251.000000, Respawntime: 144, Condition: 0, Version: 0, Y: 1503.000000, Cond_value: 1, Enabled: 1, Animation: 0, Id: sql.NullInt64{Int64: 364601, Valid: true}, X: 81.000000, Z: 442.000000, Variance: 0, Pathgrid: 93, Spawngroupid: 108506},
	{X: -568.000000, Y: 1178.000000, Pathgrid: 92, Enabled: 1, Animation: 0, Id: sql.NullInt64{Int64: 364602, Valid: true}, Zone: sql.NullString{String: "veeshan", Valid: true}, Z: 688.299988, Variance: 0, Condition: 0, Spawngroupid: 108506, Cond_value: 1, Version: 0, Heading: 119.320099, Respawntime: 144},
	{Animation: 0, Id: sql.NullInt64{Int64: 364603, Valid: true}, Version: 0, X: 415.000000, Y: 1339.000000, Respawntime: 144, Condition: 0, Spawngroupid: 108506, Variance: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, Z: 609.700012, Heading: 61.189800, Pathgrid: 89, Cond_value: 1, Enabled: 1},
	{Version: 0, X: 325.000000, Variance: 0, Condition: 0, Enabled: 1, Spawngroupid: 108506, Z: 614.099976, Pathgrid: 91, Cond_value: 1, Animation: 0, Y: 1330.000000, Respawntime: 144, Zone: sql.NullString{String: "veeshan", Valid: true}, Heading: 61.189800, Id: sql.NullInt64{Int64: 364604, Valid: true}},
	{Zone: sql.NullString{String: "veeshan", Valid: true}, Respawntime: 144, Pathgrid: 88, Enabled: 1, Animation: 0, Spawngroupid: 108506, Y: 1691.000000, Id: sql.NullInt64{Int64: 364605, Valid: true}, Version: 0, Z: 591.900024, Condition: 0, X: 650.000000, Heading: 64.000000, Variance: 0, Cond_value: 1},
	{Pathgrid: 86, Condition: 0, Spawngroupid: 108506, Z: 521.900024, Respawntime: 144, Animation: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, X: 601.000000, Heading: 36.000000, Enabled: 1, Id: sql.NullInt64{Int64: 364606, Valid: true}, Y: 890.000000, Cond_value: 1, Version: 0, Variance: 0},
	{Heading: 255.000000, Variance: 0, Version: 0, X: 640.000000, Y: 1259.000000, Z: 591.900024, Zone: sql.NullString{String: "veeshan", Valid: true}, Enabled: 1, Spawngroupid: 108506, Respawntime: 144, Pathgrid: 87, Id: sql.NullInt64{Int64: 364607, Valid: true}, Condition: 0, Cond_value: 1, Animation: 0},
	{Y: 840.000000, Z: 442.000000, Enabled: 1, Cond_value: 1, Zone: sql.NullString{String: "veeshan", Valid: true}, Version: 0, Pathgrid: 85, Condition: 0, Id: sql.NullInt64{Int64: 364608, Valid: true}, X: 40.000000, Respawntime: 144, Variance: 0, Spawngroupid: 108506, Heading: 191.000000, Animation: 0},
	{Cond_value: 1, Heading: 65.000000, Spawngroupid: 108506, Version: 0, X: -768.000000, Y: 840.000000, Z: 442.000000, Condition: 0, Id: sql.NullInt64{Int64: 364610, Valid: true}, Variance: 0, Pathgrid: 83, Zone: sql.NullString{String: "veeshan", Valid: true}, Enabled: 1, Animation: 0, Respawntime: 144},
	{Id: sql.NullInt64{Int64: 364611, Valid: true}, Spawngroupid: 108506, Version: 0, Y: -256.000000, Condition: 0, Heading: 0.000000, Pathgrid: 82, Cond_value: 1, X: -717.000000, Variance: 0, Enabled: 1, Animation: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, Z: 459.799988, Respawntime: 144},
	{Pathgrid: 81, Id: sql.NullInt64{Int64: 364612, Valid: true}, Y: -482.000000, Z: 439.799988, Version: 0, X: -48.000000, Respawntime: 144, Variance: 0, Condition: 0, Cond_value: 1, Spawngroupid: 108506, Zone: sql.NullString{String: "veeshan", Valid: true}, Animation: 0, Heading: 193.767700, Enabled: 1},
	{Id: sql.NullInt64{Int64: 364613, Valid: true}, Version: 0, Heading: 194.787506, Variance: 0, Pathgrid: 79, Condition: 0, Cond_value: 1, Spawngroupid: 108506, Zone: sql.NullString{String: "veeshan", Valid: true}, Z: 378.799988, Respawntime: 144, Animation: 0, Enabled: 1, X: 52.000000, Y: 486.000000},
	{Zone: sql.NullString{String: "veeshan", Valid: true}, X: 1067.000000, Pathgrid: 78, Animation: 0, Id: sql.NullInt64{Int64: 364614, Valid: true}, Enabled: 1, Spawngroupid: 108506, Version: 0, Cond_value: 1, Y: 520.000000, Z: 219.899994, Heading: 42.832859, Respawntime: 144, Variance: 0, Condition: 0},
	{Variance: 0, Pathgrid: 74, Condition: 0, Version: 0, X: 466.000000, Y: 1188.000000, Id: sql.NullInt64{Int64: 364615, Valid: true}, Heading: 13.000000, Animation: 0, Spawngroupid: 108506, Z: 269.200012, Cond_value: 1, Zone: sql.NullString{String: "veeshan", Valid: true}, Respawntime: 144, Enabled: 1},
	{Id: sql.NullInt64{Int64: 364616, Valid: true}, Y: 1092.000000, Z: 242.000000, Respawntime: 144, Variance: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, X: 795.000000, Pathgrid: 75, Condition: 0, Cond_value: 1, Enabled: 1, Spawngroupid: 108506, Version: 0, Heading: 190.000000, Animation: 0},
	{Heading: 64.249290, Animation: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, X: 900.000000, Y: 798.000000, Z: 219.899994, Respawntime: 144, Pathgrid: 77, Condition: 0, Enabled: 1, Id: sql.NullInt64{Int64: 364617, Valid: true}, Spawngroupid: 108506, Version: 0, Variance: 0, Cond_value: 1},
	{Enabled: 1, Version: 0, Respawntime: 144, Cond_value: 1, Spawngroupid: 108506, Zone: sql.NullString{String: "veeshan", Valid: true}, X: 1497.000000, Variance: 0, Animation: 0, Id: sql.NullInt64{Int64: 364618, Valid: true}, Z: 91.699997, Heading: 65.269119, Y: 800.000000, Pathgrid: 76, Condition: 0},
	{Spawngroupid: 108506, Cond_value: 1, X: 365.000000, Variance: 0, Pathgrid: 80, Condition: 0, Id: sql.NullInt64{Int64: 364619, Valid: true}, Version: 0, Enabled: 1, Animation: 0, Zone: sql.NullString{String: "veeshan", Valid: true}, Y: -81.000000, Z: 378.799988, Heading: 9.178471, Respawntime: 144},
	{Spawngroupid: 108506, Enabled: 1, Animation: 0, Id: sql.NullInt64{Int64: 364620, Valid: true}, Zone: sql.NullString{String: "veeshan", Valid: true}, X: 1703.000000, Y: 707.000000, Z: 59.969749, Respawntime: 144, Pathgrid: 71, Condition: 0, Version: 0, Heading: 0.000000, Variance: 0, Cond_value: 1},
	{Zone: sql.NullString{String: "veeshan", Valid: true}, Version: 0, X: -194.000000, Condition: 0, Animation: 0, Id: sql.NullInt64{Int64: 364598, Valid: true}, Spawngroupid: 108509, Variance: 43200, Cond_value: 1, Heading: 192.000000, Respawntime: 604800, Y: 319.000000, Z: 442.000000, Pathgrid: 73, Enabled: 1},
	{Respawntime: 604800, Pathgrid: 0, Enabled: 1, Zone: sql.NullString{String: "veeshan", Valid: true}, Version: 0, Z: 299.799988, Y: -1154.000000, Variance: 43200, Cond_value: 1, Id: sql.NullInt64{Int64: 364623, Valid: true}, Spawngroupid: 108510, X: -1313.000000, Heading: 41.813030, Condition: 0, Animation: 0},
	{Id: sql.NullInt64{Int64: 364631, Valid: true}, Version: 0, Heading: 107.082199, Respawntime: 604800, Variance: 43200, Cond_value: 1, Enabled: 1, Spawngroupid: 108511, Zone: sql.NullString{String: "veeshan", Valid: true}, X: 1209.000000, Y: 1336.000000, Z: 779.599976, Pathgrid: 0, Condition: 0, Animation: 0},
	{Version: 0, Respawntime: 604800, Variance: 43200, Zone: sql.NullString{String: "veeshan", Valid: true}, Spawngroupid: 108512, Y: 1418.000000, Z: 459.799988, Heading: 54.050991, Pathgrid: 0, Cond_value: 1, Id: sql.NullInt64{Int64: 364639, Valid: true}, X: -205.000000, Condition: 0, Enabled: 1, Animation: 0},
	{Respawntime: 604800, Enabled: 1, Id: sql.NullInt64{Int64: 364645, Valid: true}, Heading: 244.759201, Animation: 0, Z: 689.400024, Version: 0, Y: 1171.000000, Condition: 0, Spawngroupid: 108513, X: -560.000000, Variance: 43200, Pathgrid: 0, Cond_value: 1, Zone: sql.NullString{String: "veeshan", Valid: true}},
	{Variance: 43200, Cond_value: 1, Z: 272.000000, Heading: 139.000000, Respawntime: 604800, Id: sql.NullInt64{Int64: 364659, Valid: true}, Y: 1463.000000, X: 562.000000, Pathgrid: 72, Condition: 0, Enabled: 1, Animation: 0, Spawngroupid: 108517, Zone: sql.NullString{String: "veeshan", Valid: true}, Version: 0},
}

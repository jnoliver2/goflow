// Generated from excellent/gen/Excellent.g4 by ANTLR 4.7.

package gen

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 27, 182,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 6, 20, 108, 10, 20, 13, 20,
	14, 20, 109, 3, 20, 3, 20, 6, 20, 114, 10, 20, 13, 20, 14, 20, 115, 5,
	20, 118, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 124, 10, 21, 12, 21,
	14, 21, 127, 11, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 6, 24, 143, 10, 24, 13, 24,
	14, 24, 144, 3, 24, 3, 24, 3, 24, 7, 24, 150, 10, 24, 12, 24, 14, 24, 153,
	11, 24, 3, 25, 6, 25, 156, 10, 25, 13, 25, 14, 25, 157, 3, 25, 3, 25, 3,
	26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 169, 10, 27, 3, 28,
	3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3,
	33, 2, 2, 34, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 2, 55,
	2, 57, 2, 59, 2, 61, 2, 63, 2, 65, 2, 3, 2, 19, 3, 2, 50, 59, 3, 2, 36,
	36, 4, 2, 86, 86, 118, 118, 4, 2, 84, 84, 116, 116, 4, 2, 87, 87, 119,
	119, 4, 2, 71, 71, 103, 103, 4, 2, 72, 72, 104, 104, 4, 2, 67, 67, 99,
	99, 4, 2, 78, 78, 110, 110, 4, 2, 85, 85, 117, 117, 5, 2, 11, 12, 15, 15,
	34, 34, 84, 2, 67, 92, 194, 216, 218, 224, 258, 312, 315, 329, 332, 383,
	387, 388, 390, 397, 400, 403, 405, 406, 408, 410, 414, 415, 417, 418, 420,
	427, 430, 437, 439, 446, 454, 463, 465, 477, 480, 496, 499, 502, 504, 506,
	508, 564, 572, 573, 575, 576, 579, 584, 586, 592, 882, 884, 888, 897, 904,
	908, 910, 931, 933, 941, 977, 982, 986, 1008, 1014, 1017, 1019, 1020, 1023,
	1073, 1122, 1154, 1164, 1231, 1234, 1328, 1331, 1368, 4258, 4295, 4297,
	4303, 7682, 7830, 7840, 7936, 7946, 7953, 7962, 7967, 7978, 7985, 7994,
	8001, 8010, 8015, 8027, 8033, 8042, 8049, 8122, 8125, 8138, 8141, 8154,
	8157, 8170, 8174, 8186, 8189, 8452, 8457, 8461, 8463, 8466, 8468, 8471,
	8479, 8486, 8495, 8498, 8501, 8512, 8513, 8519, 8581, 11266, 11312, 11362,
	11366, 11369, 11378, 11380, 11383, 11392, 11394, 11396, 11492, 11501, 11503,
	11508, 42562, 42564, 42606, 42626, 42652, 42788, 42800, 42804, 42864, 42875,
	42888, 42893, 42895, 42898, 42900, 42904, 42927, 42930, 42931, 65315, 65340,
	83, 2, 99, 124, 183, 248, 250, 257, 259, 377, 380, 386, 389, 391, 394,
	404, 407, 413, 416, 419, 421, 423, 426, 431, 434, 438, 440, 449, 456, 462,
	464, 501, 503, 507, 509, 571, 574, 580, 585, 661, 663, 689, 883, 885, 889,
	895, 914, 976, 978, 979, 983, 985, 987, 1013, 1015, 1121, 1123, 1155, 1165,
	1217, 1220, 1329, 1379, 1417, 7426, 7469, 7533, 7545, 7547, 7580, 7683,
	7839, 7841, 7945, 7954, 7959, 7970, 7977, 7986, 7993, 8002, 8007, 8018,
	8025, 8034, 8041, 8050, 8063, 8066, 8073, 8082, 8089, 8098, 8105, 8114,
	8118, 8120, 8121, 8128, 8134, 8136, 8137, 8146, 8149, 8152, 8153, 8162,
	8169, 8180, 8182, 8184, 8185, 8460, 8469, 8497, 8507, 8510, 8511, 8520,
	8523, 8528, 8582, 11314, 11360, 11363, 11374, 11379, 11389, 11395, 11502,
	11504, 11509, 11522, 11559, 11561, 11567, 42563, 42607, 42627, 42653, 42789,
	42803, 42805, 42874, 42876, 42878, 42881, 42889, 42894, 42896, 42899, 42903,
	42905, 42923, 43004, 43868, 43878, 43879, 64258, 64264, 64277, 64281, 65347,
	65372, 8, 2, 455, 461, 500, 8081, 8090, 8097, 8106, 8113, 8126, 8142, 8190,
	8190, 35, 2, 690, 707, 712, 723, 738, 742, 750, 752, 886, 892, 1371, 1602,
	1767, 1768, 2038, 2039, 2044, 2076, 2086, 2090, 2419, 3656, 3784, 4350,
	6105, 6213, 6825, 7295, 7470, 7532, 7546, 7617, 8307, 8321, 8338, 8350,
	11390, 11391, 11633, 11825, 12295, 12343, 12349, 12544, 40983, 42239, 42510,
	42625, 42654, 42655, 42777, 42785, 42866, 42890, 43002, 43003, 43473, 43496,
	43634, 43743, 43765, 43766, 43870, 43873, 65394, 65441, 236, 2, 172, 188,
	445, 453, 662, 1516, 1522, 1524, 1570, 1601, 1603, 1612, 1648, 1649, 1651,
	1749, 1751, 1790, 1793, 1810, 1812, 1841, 1871, 1959, 1971, 2028, 2050,
	2071, 2114, 2138, 2210, 2228, 2310, 2363, 2367, 2386, 2394, 2403, 2420,
	2434, 2439, 2446, 2449, 2450, 2453, 2474, 2476, 2482, 2484, 2491, 2495,
	2512, 2526, 2527, 2529, 2531, 2546, 2547, 2567, 2572, 2577, 2578, 2581,
	2602, 2604, 2610, 2612, 2613, 2615, 2616, 2618, 2619, 2651, 2654, 2656,
	2678, 2695, 2703, 2705, 2707, 2709, 2730, 2732, 2738, 2740, 2741, 2743,
	2747, 2751, 2770, 2786, 2787, 2823, 2830, 2833, 2834, 2837, 2858, 2860,
	2866, 2868, 2869, 2871, 2875, 2879, 2915, 2931, 2949, 2951, 2956, 2960,
	2962, 2964, 2967, 2971, 2972, 2974, 2988, 2992, 3003, 3026, 3086, 3088,
	3090, 3092, 3114, 3116, 3131, 3135, 3214, 3216, 3218, 3220, 3242, 3244,
	3253, 3255, 3259, 3263, 3296, 3298, 3299, 3315, 3316, 3335, 3342, 3344,
	3346, 3348, 3388, 3391, 3408, 3426, 3427, 3452, 3457, 3463, 3480, 3484,
	3507, 3509, 3517, 3519, 3528, 3587, 3634, 3636, 3637, 3650, 3655, 3715,
	3716, 3718, 3724, 3727, 3737, 3739, 3745, 3747, 3749, 3751, 3753, 3756,
	3757, 3759, 3762, 3764, 3765, 3775, 3782, 3806, 3809, 3842, 3913, 3915,
	3950, 3978, 3982, 4098, 4140, 4161, 4183, 4188, 4191, 4195, 4210, 4215,
	4227, 4240, 4348, 4351, 4682, 4684, 4687, 4690, 4696, 4698, 4703, 4706,
	4746, 4748, 4751, 4754, 4786, 4788, 4791, 4794, 4800, 4802, 4807, 4810,
	4824, 4826, 4882, 4884, 4887, 4890, 4956, 4994, 5009, 5026, 5110, 5123,
	5742, 5745, 5761, 5763, 5788, 5794, 5868, 5875, 5882, 5890, 5902, 5904,
	5907, 5922, 5939, 5954, 5971, 5986, 5998, 6000, 6002, 6018, 6069, 6110,
	6212, 6214, 6265, 6274, 6314, 6316, 6391, 6402, 6432, 6482, 6511, 6514,
	6518, 6530, 6573, 6595, 6601, 6658, 6680, 6690, 6742, 6919, 6965, 6983,
	6989, 7045, 7074, 7088, 7089, 7100, 7143, 7170, 7205, 7247, 7249, 7260,
	7289, 7403, 7406, 7408, 7411, 7415, 7416, 8503, 8506, 11570, 11625, 11650,
	11672, 11682, 11688, 11690, 11696, 11698, 11704, 11706, 11712, 11714, 11720,
	11722, 11728, 11730, 11736, 11738, 11744, 12296, 12350, 12355, 12440, 12449,
	12540, 12545, 12591, 12595, 12688, 12706, 12732, 12786, 12801, 13314, 19895,
	19970, 40910, 40962, 40982, 40984, 42126, 42194, 42233, 42242, 42509, 42514,
	42529, 42540, 42541, 42608, 42727, 43001, 43011, 43013, 43015, 43017, 43020,
	43022, 43044, 43074, 43125, 43140, 43189, 43252, 43257, 43261, 43303, 43314,
	43336, 43362, 43390, 43398, 43444, 43490, 43494, 43497, 43505, 43516, 43520,
	43522, 43562, 43586, 43588, 43590, 43597, 43618, 43633, 43635, 43640, 43644,
	43697, 43699, 43711, 43714, 43716, 43741, 43742, 43746, 43756, 43764, 43784,
	43787, 43792, 43795, 43800, 43810, 43816, 43818, 43824, 43970, 44004, 44034,
	55205, 55218, 55240, 55245, 55293, 63746, 64111, 64114, 64219, 64287, 64298,
	64300, 64312, 64314, 64318, 64320, 64435, 64469, 64831, 64850, 64913, 64916,
	64969, 65010, 65021, 65138, 65142, 65144, 65278, 65384, 65393, 65395, 65439,
	65442, 65472, 65476, 65481, 65484, 65489, 65492, 65497, 65500, 65502, 39,
	2, 50, 59, 1634, 1643, 1778, 1787, 1986, 1995, 2408, 2417, 2536, 2545,
	2664, 2673, 2792, 2801, 2920, 2929, 3048, 3057, 3176, 3185, 3304, 3313,
	3432, 3441, 3560, 3569, 3666, 3675, 3794, 3803, 3874, 3883, 4162, 4171,
	4242, 4251, 6114, 6123, 6162, 6171, 6472, 6481, 6610, 6619, 6786, 6795,
	6802, 6811, 6994, 7003, 7090, 7099, 7234, 7243, 7250, 7259, 42530, 42539,
	43218, 43227, 43266, 43275, 43474, 43483, 43506, 43515, 43602, 43611, 44018,
	44027, 65298, 65307, 2, 188, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7,
	3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2,
	15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2,
	2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2,
	2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2,
	2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3,
	2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 3, 67,
	3, 2, 2, 2, 5, 69, 3, 2, 2, 2, 7, 71, 3, 2, 2, 2, 9, 73, 3, 2, 2, 2, 11,
	75, 3, 2, 2, 2, 13, 77, 3, 2, 2, 2, 15, 79, 3, 2, 2, 2, 17, 81, 3, 2, 2,
	2, 19, 83, 3, 2, 2, 2, 21, 85, 3, 2, 2, 2, 23, 87, 3, 2, 2, 2, 25, 89,
	3, 2, 2, 2, 27, 91, 3, 2, 2, 2, 29, 94, 3, 2, 2, 2, 31, 97, 3, 2, 2, 2,
	33, 99, 3, 2, 2, 2, 35, 102, 3, 2, 2, 2, 37, 104, 3, 2, 2, 2, 39, 107,
	3, 2, 2, 2, 41, 119, 3, 2, 2, 2, 43, 130, 3, 2, 2, 2, 45, 135, 3, 2, 2,
	2, 47, 142, 3, 2, 2, 2, 49, 155, 3, 2, 2, 2, 51, 161, 3, 2, 2, 2, 53, 168,
	3, 2, 2, 2, 55, 170, 3, 2, 2, 2, 57, 172, 3, 2, 2, 2, 59, 174, 3, 2, 2,
	2, 61, 176, 3, 2, 2, 2, 63, 178, 3, 2, 2, 2, 65, 180, 3, 2, 2, 2, 67, 68,
	7, 46, 2, 2, 68, 4, 3, 2, 2, 2, 69, 70, 7, 42, 2, 2, 70, 6, 3, 2, 2, 2,
	71, 72, 7, 43, 2, 2, 72, 8, 3, 2, 2, 2, 73, 74, 7, 93, 2, 2, 74, 10, 3,
	2, 2, 2, 75, 76, 7, 95, 2, 2, 76, 12, 3, 2, 2, 2, 77, 78, 7, 48, 2, 2,
	78, 14, 3, 2, 2, 2, 79, 80, 7, 45, 2, 2, 80, 16, 3, 2, 2, 2, 81, 82, 7,
	47, 2, 2, 82, 18, 3, 2, 2, 2, 83, 84, 7, 44, 2, 2, 84, 20, 3, 2, 2, 2,
	85, 86, 7, 49, 2, 2, 86, 22, 3, 2, 2, 2, 87, 88, 7, 96, 2, 2, 88, 24, 3,
	2, 2, 2, 89, 90, 7, 63, 2, 2, 90, 26, 3, 2, 2, 2, 91, 92, 7, 35, 2, 2,
	92, 93, 7, 63, 2, 2, 93, 28, 3, 2, 2, 2, 94, 95, 7, 62, 2, 2, 95, 96, 7,
	63, 2, 2, 96, 30, 3, 2, 2, 2, 97, 98, 7, 62, 2, 2, 98, 32, 3, 2, 2, 2,
	99, 100, 7, 64, 2, 2, 100, 101, 7, 63, 2, 2, 101, 34, 3, 2, 2, 2, 102,
	103, 7, 64, 2, 2, 103, 36, 3, 2, 2, 2, 104, 105, 7, 40, 2, 2, 105, 38,
	3, 2, 2, 2, 106, 108, 9, 2, 2, 2, 107, 106, 3, 2, 2, 2, 108, 109, 3, 2,
	2, 2, 109, 107, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 117, 3, 2, 2, 2,
	111, 113, 7, 48, 2, 2, 112, 114, 9, 2, 2, 2, 113, 112, 3, 2, 2, 2, 114,
	115, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 118,
	3, 2, 2, 2, 117, 111, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 40, 3, 2,
	2, 2, 119, 125, 7, 36, 2, 2, 120, 124, 10, 3, 2, 2, 121, 122, 7, 36, 2,
	2, 122, 124, 7, 36, 2, 2, 123, 120, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2, 124,
	127, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 128,
	3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 128, 129, 7, 36, 2, 2, 129, 42, 3, 2,
	2, 2, 130, 131, 9, 4, 2, 2, 131, 132, 9, 5, 2, 2, 132, 133, 9, 6, 2, 2,
	133, 134, 9, 7, 2, 2, 134, 44, 3, 2, 2, 2, 135, 136, 9, 8, 2, 2, 136, 137,
	9, 9, 2, 2, 137, 138, 9, 10, 2, 2, 138, 139, 9, 11, 2, 2, 139, 140, 9,
	7, 2, 2, 140, 46, 3, 2, 2, 2, 141, 143, 5, 53, 27, 2, 142, 141, 3, 2, 2,
	2, 143, 144, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145,
	151, 3, 2, 2, 2, 146, 150, 5, 53, 27, 2, 147, 150, 5, 65, 33, 2, 148, 150,
	7, 97, 2, 2, 149, 146, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 149, 148, 3, 2,
	2, 2, 150, 153, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2,
	152, 48, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 154, 156, 9, 12, 2, 2, 155,
	154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 158,
	3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 160, 8, 25, 2, 2, 160, 50, 3, 2,
	2, 2, 161, 162, 11, 2, 2, 2, 162, 52, 3, 2, 2, 2, 163, 169, 5, 55, 28,
	2, 164, 169, 5, 57, 29, 2, 165, 169, 5, 59, 30, 2, 166, 169, 5, 61, 31,
	2, 167, 169, 5, 63, 32, 2, 168, 163, 3, 2, 2, 2, 168, 164, 3, 2, 2, 2,
	168, 165, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 168, 167, 3, 2, 2, 2, 169,
	54, 3, 2, 2, 2, 170, 171, 9, 13, 2, 2, 171, 56, 3, 2, 2, 2, 172, 173, 9,
	14, 2, 2, 173, 58, 3, 2, 2, 2, 174, 175, 9, 15, 2, 2, 175, 60, 3, 2, 2,
	2, 176, 177, 9, 16, 2, 2, 177, 62, 3, 2, 2, 2, 178, 179, 9, 17, 2, 2, 179,
	64, 3, 2, 2, 2, 180, 181, 9, 18, 2, 2, 181, 66, 3, 2, 2, 2, 13, 2, 109,
	115, 117, 123, 125, 144, 149, 151, 157, 168, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "'('", "')'", "'['", "']'", "'.'", "'+'", "'-'", "'*'", "'/'",
	"'^'", "'='", "'!='", "'<='", "'<'", "'>='", "'>'", "'&'",
}

var lexerSymbolicNames = []string{
	"", "COMMA", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "DOT", "PLUS", "MINUS",
	"TIMES", "DIVIDE", "EXPONENT", "EQ", "NEQ", "LTE", "LT", "GTE", "GT", "AMPERSAND",
	"DECIMAL", "STRING", "TRUE", "FALSE", "NAME", "WS", "ERROR",
}

var lexerRuleNames = []string{
	"COMMA", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "DOT", "PLUS", "MINUS",
	"TIMES", "DIVIDE", "EXPONENT", "EQ", "NEQ", "LTE", "LT", "GTE", "GT", "AMPERSAND",
	"DECIMAL", "STRING", "TRUE", "FALSE", "NAME", "WS", "ERROR", "UnicodeLetter",
	"UnicodeClass_LU", "UnicodeClass_LL", "UnicodeClass_LT", "UnicodeClass_LM",
	"UnicodeClass_LO", "UnicodeDigit",
}

type ExcellentLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewExcellentLexer(input antlr.CharStream) *ExcellentLexer {

	l := new(ExcellentLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Excellent.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExcellentLexer tokens.
const (
	ExcellentLexerCOMMA     = 1
	ExcellentLexerLPAREN    = 2
	ExcellentLexerRPAREN    = 3
	ExcellentLexerLBRACK    = 4
	ExcellentLexerRBRACK    = 5
	ExcellentLexerDOT       = 6
	ExcellentLexerPLUS      = 7
	ExcellentLexerMINUS     = 8
	ExcellentLexerTIMES     = 9
	ExcellentLexerDIVIDE    = 10
	ExcellentLexerEXPONENT  = 11
	ExcellentLexerEQ        = 12
	ExcellentLexerNEQ       = 13
	ExcellentLexerLTE       = 14
	ExcellentLexerLT        = 15
	ExcellentLexerGTE       = 16
	ExcellentLexerGT        = 17
	ExcellentLexerAMPERSAND = 18
	ExcellentLexerDECIMAL   = 19
	ExcellentLexerSTRING    = 20
	ExcellentLexerTRUE      = 21
	ExcellentLexerFALSE     = 22
	ExcellentLexerNAME      = 23
	ExcellentLexerWS        = 24
	ExcellentLexerERROR     = 25
)

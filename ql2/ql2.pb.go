// Code generated by protoc-gen-go.
// source: ql2.proto
// DO NOT EDIT!

package ql2

import proto "github.com/golang/protobuf"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type VersionDummy_Version int32

const (
	VersionDummy_V0_1 VersionDummy_Version = 1063369270
	VersionDummy_V0_2 VersionDummy_Version = 1915781601
	VersionDummy_V0_3 VersionDummy_Version = 1601562686
)

var VersionDummy_Version_name = map[int32]string{
	1063369270: "V0_1",
	1915781601: "V0_2",
	1601562686: "V0_3",
}
var VersionDummy_Version_value = map[string]int32{
	"V0_1": 1063369270,
	"V0_2": 1915781601,
	"V0_3": 1601562686,
}

func (x VersionDummy_Version) Enum() *VersionDummy_Version {
	p := new(VersionDummy_Version)
	*p = x
	return p
}
func (x VersionDummy_Version) String() string {
	return proto.EnumName(VersionDummy_Version_name, int32(x))
}
func (x VersionDummy_Version) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *VersionDummy_Version) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(VersionDummy_Version_value, data, "VersionDummy_Version")
	if err != nil {
		return err
	}
	*x = VersionDummy_Version(value)
	return nil
}

type VersionDummy_Protocol int32

const (
	VersionDummy_PROTOBUF VersionDummy_Protocol = 656407617
	VersionDummy_JSON     VersionDummy_Protocol = 2120839367
)

var VersionDummy_Protocol_name = map[int32]string{
	656407617:  "PROTOBUF",
	2120839367: "JSON",
}
var VersionDummy_Protocol_value = map[string]int32{
	"PROTOBUF": 656407617,
	"JSON":     2120839367,
}

func (x VersionDummy_Protocol) Enum() *VersionDummy_Protocol {
	p := new(VersionDummy_Protocol)
	*p = x
	return p
}
func (x VersionDummy_Protocol) String() string {
	return proto.EnumName(VersionDummy_Protocol_name, int32(x))
}
func (x VersionDummy_Protocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *VersionDummy_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(VersionDummy_Protocol_value, data, "VersionDummy_Protocol")
	if err != nil {
		return err
	}
	*x = VersionDummy_Protocol(value)
	return nil
}

type Query_QueryType int32

const (
	Query_START        Query_QueryType = 1
	Query_CONTINUE     Query_QueryType = 2
	Query_STOP         Query_QueryType = 3
	Query_NOREPLY_WAIT Query_QueryType = 4
)

var Query_QueryType_name = map[int32]string{
	1: "START",
	2: "CONTINUE",
	3: "STOP",
	4: "NOREPLY_WAIT",
}
var Query_QueryType_value = map[string]int32{
	"START":        1,
	"CONTINUE":     2,
	"STOP":         3,
	"NOREPLY_WAIT": 4,
}

func (x Query_QueryType) Enum() *Query_QueryType {
	p := new(Query_QueryType)
	*p = x
	return p
}
func (x Query_QueryType) String() string {
	return proto.EnumName(Query_QueryType_name, int32(x))
}
func (x Query_QueryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Query_QueryType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Query_QueryType_value, data, "Query_QueryType")
	if err != nil {
		return err
	}
	*x = Query_QueryType(value)
	return nil
}

type Frame_FrameType int32

const (
	Frame_POS Frame_FrameType = 1
	Frame_OPT Frame_FrameType = 2
)

var Frame_FrameType_name = map[int32]string{
	1: "POS",
	2: "OPT",
}
var Frame_FrameType_value = map[string]int32{
	"POS": 1,
	"OPT": 2,
}

func (x Frame_FrameType) Enum() *Frame_FrameType {
	p := new(Frame_FrameType)
	*p = x
	return p
}
func (x Frame_FrameType) String() string {
	return proto.EnumName(Frame_FrameType_name, int32(x))
}
func (x Frame_FrameType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Frame_FrameType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Frame_FrameType_value, data, "Frame_FrameType")
	if err != nil {
		return err
	}
	*x = Frame_FrameType(value)
	return nil
}

type Response_ResponseType int32

const (
	Response_SUCCESS_ATOM      Response_ResponseType = 1
	Response_SUCCESS_SEQUENCE  Response_ResponseType = 2
	Response_SUCCESS_PARTIAL   Response_ResponseType = 3
	Response_SUCCESS_FEED      Response_ResponseType = 5
	Response_WAIT_COMPLETE     Response_ResponseType = 4
	Response_SUCCESS_ATOM_FEED Response_ResponseType = 6
	Response_CLIENT_ERROR      Response_ResponseType = 16
	Response_COMPILE_ERROR     Response_ResponseType = 17
	Response_RUNTIME_ERROR     Response_ResponseType = 18
)

var Response_ResponseType_name = map[int32]string{
	1:  "SUCCESS_ATOM",
	2:  "SUCCESS_SEQUENCE",
	3:  "SUCCESS_PARTIAL",
	5:  "SUCCESS_FEED",
	4:  "WAIT_COMPLETE",
	6:  "SUCCESS_ATOM_FEED",
	16: "CLIENT_ERROR",
	17: "COMPILE_ERROR",
	18: "RUNTIME_ERROR",
}
var Response_ResponseType_value = map[string]int32{
	"SUCCESS_ATOM":      1,
	"SUCCESS_SEQUENCE":  2,
	"SUCCESS_PARTIAL":   3,
	"SUCCESS_FEED":      5,
	"WAIT_COMPLETE":     4,
	"SUCCESS_ATOM_FEED": 6,
	"CLIENT_ERROR":      16,
	"COMPILE_ERROR":     17,
	"RUNTIME_ERROR":     18,
}

func (x Response_ResponseType) Enum() *Response_ResponseType {
	p := new(Response_ResponseType)
	*p = x
	return p
}
func (x Response_ResponseType) String() string {
	return proto.EnumName(Response_ResponseType_name, int32(x))
}
func (x Response_ResponseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Response_ResponseType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Response_ResponseType_value, data, "Response_ResponseType")
	if err != nil {
		return err
	}
	*x = Response_ResponseType(value)
	return nil
}

type Datum_DatumType int32

const (
	Datum_R_NULL   Datum_DatumType = 1
	Datum_R_BOOL   Datum_DatumType = 2
	Datum_R_NUM    Datum_DatumType = 3
	Datum_R_STR    Datum_DatumType = 4
	Datum_R_ARRAY  Datum_DatumType = 5
	Datum_R_OBJECT Datum_DatumType = 6
	Datum_R_JSON   Datum_DatumType = 7
)

var Datum_DatumType_name = map[int32]string{
	1: "R_NULL",
	2: "R_BOOL",
	3: "R_NUM",
	4: "R_STR",
	5: "R_ARRAY",
	6: "R_OBJECT",
	7: "R_JSON",
}
var Datum_DatumType_value = map[string]int32{
	"R_NULL":   1,
	"R_BOOL":   2,
	"R_NUM":    3,
	"R_STR":    4,
	"R_ARRAY":  5,
	"R_OBJECT": 6,
	"R_JSON":   7,
}

func (x Datum_DatumType) Enum() *Datum_DatumType {
	p := new(Datum_DatumType)
	*p = x
	return p
}
func (x Datum_DatumType) String() string {
	return proto.EnumName(Datum_DatumType_name, int32(x))
}
func (x Datum_DatumType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Datum_DatumType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Datum_DatumType_value, data, "Datum_DatumType")
	if err != nil {
		return err
	}
	*x = Datum_DatumType(value)
	return nil
}

type Term_TermType int32

const (
	Term_DATUM            Term_TermType = 1
	Term_MAKE_ARRAY       Term_TermType = 2
	Term_MAKE_OBJ         Term_TermType = 3
	Term_VAR              Term_TermType = 10
	Term_JAVASCRIPT       Term_TermType = 11
	Term_UUID             Term_TermType = 169
	Term_HTTP             Term_TermType = 153
	Term_ERROR            Term_TermType = 12
	Term_IMPLICIT_VAR     Term_TermType = 13
	Term_DB               Term_TermType = 14
	Term_TABLE            Term_TermType = 15
	Term_GET              Term_TermType = 16
	Term_GET_ALL          Term_TermType = 78
	Term_EQ               Term_TermType = 17
	Term_NE               Term_TermType = 18
	Term_LT               Term_TermType = 19
	Term_LE               Term_TermType = 20
	Term_GT               Term_TermType = 21
	Term_GE               Term_TermType = 22
	Term_NOT              Term_TermType = 23
	Term_ADD              Term_TermType = 24
	Term_SUB              Term_TermType = 25
	Term_MUL              Term_TermType = 26
	Term_DIV              Term_TermType = 27
	Term_MOD              Term_TermType = 28
	Term_APPEND           Term_TermType = 29
	Term_PREPEND          Term_TermType = 80
	Term_DIFFERENCE       Term_TermType = 95
	Term_SET_INSERT       Term_TermType = 88
	Term_SET_INTERSECTION Term_TermType = 89
	Term_SET_UNION        Term_TermType = 90
	Term_SET_DIFFERENCE   Term_TermType = 91
	Term_SLICE            Term_TermType = 30
	Term_SKIP             Term_TermType = 70
	Term_LIMIT            Term_TermType = 71
	Term_INDEXES_OF       Term_TermType = 87
	Term_CONTAINS         Term_TermType = 93
	Term_GET_FIELD        Term_TermType = 31
	Term_KEYS             Term_TermType = 94
	Term_OBJECT           Term_TermType = 143
	Term_HAS_FIELDS       Term_TermType = 32
	Term_WITH_FIELDS      Term_TermType = 96
	Term_PLUCK            Term_TermType = 33
	Term_WITHOUT          Term_TermType = 34
	Term_MERGE            Term_TermType = 35
	Term_BETWEEN          Term_TermType = 36
	Term_REDUCE           Term_TermType = 37
	Term_MAP              Term_TermType = 38
	Term_FILTER           Term_TermType = 39
	Term_CONCAT_MAP       Term_TermType = 40
	Term_ORDER_BY         Term_TermType = 41
	Term_DISTINCT         Term_TermType = 42
	Term_COUNT            Term_TermType = 43
	Term_IS_EMPTY         Term_TermType = 86
	Term_UNION            Term_TermType = 44
	Term_NTH              Term_TermType = 45
	Term_BRACKET          Term_TermType = 170
	Term_INNER_JOIN       Term_TermType = 48
	Term_OUTER_JOIN       Term_TermType = 49
	Term_EQ_JOIN          Term_TermType = 50
	Term_ZIP              Term_TermType = 72
	Term_RANGE            Term_TermType = 173
	Term_INSERT_AT        Term_TermType = 82
	Term_DELETE_AT        Term_TermType = 83
	Term_CHANGE_AT        Term_TermType = 84
	Term_SPLICE_AT        Term_TermType = 85
	Term_COERCE_TO        Term_TermType = 51
	Term_TYPE_OF          Term_TermType = 52
	Term_UPDATE           Term_TermType = 53
	Term_DELETE           Term_TermType = 54
	Term_REPLACE          Term_TermType = 55
	Term_INSERT           Term_TermType = 56
	Term_DB_CREATE        Term_TermType = 57
	Term_DB_DROP          Term_TermType = 58
	Term_DB_LIST          Term_TermType = 59
	Term_TABLE_CREATE     Term_TermType = 60
	Term_TABLE_DROP       Term_TermType = 61
	Term_TABLE_LIST       Term_TermType = 62
	Term_CONFIG           Term_TermType = 174
	Term_STATUS           Term_TermType = 175
	Term_WAIT             Term_TermType = 177
	Term_RECONFIGURE      Term_TermType = 176
	Term_REBALANCE        Term_TermType = 179
	Term_SYNC             Term_TermType = 138
	Term_INDEX_CREATE     Term_TermType = 75
	Term_INDEX_DROP       Term_TermType = 76
	Term_INDEX_LIST       Term_TermType = 77
	Term_INDEX_STATUS     Term_TermType = 139
	Term_INDEX_WAIT       Term_TermType = 140
	Term_INDEX_RENAME     Term_TermType = 156
	Term_FUNCALL          Term_TermType = 64
	Term_BRANCH           Term_TermType = 65
	Term_ANY              Term_TermType = 66
	Term_ALL              Term_TermType = 67
	Term_FOR_EACH         Term_TermType = 68
	Term_FUNC             Term_TermType = 69
	Term_ASC              Term_TermType = 73
	Term_DESC             Term_TermType = 74
	Term_INFO             Term_TermType = 79
	Term_MATCH            Term_TermType = 97
	Term_UPCASE           Term_TermType = 141
	Term_DOWNCASE         Term_TermType = 142
	Term_SAMPLE           Term_TermType = 81
	Term_DEFAULT          Term_TermType = 92
	Term_JSON             Term_TermType = 98
	Term_TO_JSON_STRING   Term_TermType = 172
	Term_ISO8601          Term_TermType = 99
	Term_TO_ISO8601       Term_TermType = 100
	Term_EPOCH_TIME       Term_TermType = 101
	Term_TO_EPOCH_TIME    Term_TermType = 102
	Term_NOW              Term_TermType = 103
	Term_IN_TIMEZONE      Term_TermType = 104
	Term_DURING           Term_TermType = 105
	Term_DATE             Term_TermType = 106
	Term_TIME_OF_DAY      Term_TermType = 126
	Term_TIMEZONE         Term_TermType = 127
	Term_YEAR             Term_TermType = 128
	Term_MONTH            Term_TermType = 129
	Term_DAY              Term_TermType = 130
	Term_DAY_OF_WEEK      Term_TermType = 131
	Term_DAY_OF_YEAR      Term_TermType = 132
	Term_HOURS            Term_TermType = 133
	Term_MINUTES          Term_TermType = 134
	Term_SECONDS          Term_TermType = 135
	Term_TIME             Term_TermType = 136
	Term_MONDAY           Term_TermType = 107
	Term_TUESDAY          Term_TermType = 108
	Term_WEDNESDAY        Term_TermType = 109
	Term_THURSDAY         Term_TermType = 110
	Term_FRIDAY           Term_TermType = 111
	Term_SATURDAY         Term_TermType = 112
	Term_SUNDAY           Term_TermType = 113
	Term_JANUARY          Term_TermType = 114
	Term_FEBRUARY         Term_TermType = 115
	Term_MARCH            Term_TermType = 116
	Term_APRIL            Term_TermType = 117
	Term_MAY              Term_TermType = 118
	Term_JUNE             Term_TermType = 119
	Term_JULY             Term_TermType = 120
	Term_AUGUST           Term_TermType = 121
	Term_SEPTEMBER        Term_TermType = 122
	Term_OCTOBER          Term_TermType = 123
	Term_NOVEMBER         Term_TermType = 124
	Term_DECEMBER         Term_TermType = 125
	Term_LITERAL          Term_TermType = 137
	Term_GROUP            Term_TermType = 144
	Term_SUM              Term_TermType = 145
	Term_AVG              Term_TermType = 146
	Term_MIN              Term_TermType = 147
	Term_MAX              Term_TermType = 148
	Term_SPLIT            Term_TermType = 149
	Term_UNGROUP          Term_TermType = 150
	Term_RANDOM           Term_TermType = 151
	Term_CHANGES          Term_TermType = 152
	Term_ARGS             Term_TermType = 154
	Term_BINARY           Term_TermType = 155
	Term_GEOJSON          Term_TermType = 157
	Term_TO_GEOJSON       Term_TermType = 158
	Term_POINT            Term_TermType = 159
	Term_LINE             Term_TermType = 160
	Term_POLYGON          Term_TermType = 161
	Term_DISTANCE         Term_TermType = 162
	Term_INTERSECTS       Term_TermType = 163
	Term_INCLUDES         Term_TermType = 164
	Term_CIRCLE           Term_TermType = 165
	Term_GET_INTERSECTING Term_TermType = 166
	Term_FILL             Term_TermType = 167
	Term_GET_NEAREST      Term_TermType = 168
	Term_POLYGON_SUB      Term_TermType = 171
)

var Term_TermType_name = map[int32]string{
	1:   "DATUM",
	2:   "MAKE_ARRAY",
	3:   "MAKE_OBJ",
	10:  "VAR",
	11:  "JAVASCRIPT",
	169: "UUID",
	153: "HTTP",
	12:  "ERROR",
	13:  "IMPLICIT_VAR",
	14:  "DB",
	15:  "TABLE",
	16:  "GET",
	78:  "GET_ALL",
	17:  "EQ",
	18:  "NE",
	19:  "LT",
	20:  "LE",
	21:  "GT",
	22:  "GE",
	23:  "NOT",
	24:  "ADD",
	25:  "SUB",
	26:  "MUL",
	27:  "DIV",
	28:  "MOD",
	29:  "APPEND",
	80:  "PREPEND",
	95:  "DIFFERENCE",
	88:  "SET_INSERT",
	89:  "SET_INTERSECTION",
	90:  "SET_UNION",
	91:  "SET_DIFFERENCE",
	30:  "SLICE",
	70:  "SKIP",
	71:  "LIMIT",
	87:  "INDEXES_OF",
	93:  "CONTAINS",
	31:  "GET_FIELD",
	94:  "KEYS",
	143: "OBJECT",
	32:  "HAS_FIELDS",
	96:  "WITH_FIELDS",
	33:  "PLUCK",
	34:  "WITHOUT",
	35:  "MERGE",
	36:  "BETWEEN",
	37:  "REDUCE",
	38:  "MAP",
	39:  "FILTER",
	40:  "CONCAT_MAP",
	41:  "ORDER_BY",
	42:  "DISTINCT",
	43:  "COUNT",
	86:  "IS_EMPTY",
	44:  "UNION",
	45:  "NTH",
	170: "BRACKET",
	48:  "INNER_JOIN",
	49:  "OUTER_JOIN",
	50:  "EQ_JOIN",
	72:  "ZIP",
	173: "RANGE",
	82:  "INSERT_AT",
	83:  "DELETE_AT",
	84:  "CHANGE_AT",
	85:  "SPLICE_AT",
	51:  "COERCE_TO",
	52:  "TYPE_OF",
	53:  "UPDATE",
	54:  "DELETE",
	55:  "REPLACE",
	56:  "INSERT",
	57:  "DB_CREATE",
	58:  "DB_DROP",
	59:  "DB_LIST",
	60:  "TABLE_CREATE",
	61:  "TABLE_DROP",
	62:  "TABLE_LIST",
	174: "CONFIG",
	175: "STATUS",
	177: "WAIT",
	176: "RECONFIGURE",
	179: "REBALANCE",
	138: "SYNC",
	75:  "INDEX_CREATE",
	76:  "INDEX_DROP",
	77:  "INDEX_LIST",
	139: "INDEX_STATUS",
	140: "INDEX_WAIT",
	156: "INDEX_RENAME",
	64:  "FUNCALL",
	65:  "BRANCH",
	66:  "ANY",
	67:  "ALL",
	68:  "FOR_EACH",
	69:  "FUNC",
	73:  "ASC",
	74:  "DESC",
	79:  "INFO",
	97:  "MATCH",
	141: "UPCASE",
	142: "DOWNCASE",
	81:  "SAMPLE",
	92:  "DEFAULT",
	98:  "JSON",
	172: "TO_JSON_STRING",
	99:  "ISO8601",
	100: "TO_ISO8601",
	101: "EPOCH_TIME",
	102: "TO_EPOCH_TIME",
	103: "NOW",
	104: "IN_TIMEZONE",
	105: "DURING",
	106: "DATE",
	126: "TIME_OF_DAY",
	127: "TIMEZONE",
	128: "YEAR",
	129: "MONTH",
	130: "DAY",
	131: "DAY_OF_WEEK",
	132: "DAY_OF_YEAR",
	133: "HOURS",
	134: "MINUTES",
	135: "SECONDS",
	136: "TIME",
	107: "MONDAY",
	108: "TUESDAY",
	109: "WEDNESDAY",
	110: "THURSDAY",
	111: "FRIDAY",
	112: "SATURDAY",
	113: "SUNDAY",
	114: "JANUARY",
	115: "FEBRUARY",
	116: "MARCH",
	117: "APRIL",
	118: "MAY",
	119: "JUNE",
	120: "JULY",
	121: "AUGUST",
	122: "SEPTEMBER",
	123: "OCTOBER",
	124: "NOVEMBER",
	125: "DECEMBER",
	137: "LITERAL",
	144: "GROUP",
	145: "SUM",
	146: "AVG",
	147: "MIN",
	148: "MAX",
	149: "SPLIT",
	150: "UNGROUP",
	151: "RANDOM",
	152: "CHANGES",
	154: "ARGS",
	155: "BINARY",
	157: "GEOJSON",
	158: "TO_GEOJSON",
	159: "POINT",
	160: "LINE",
	161: "POLYGON",
	162: "DISTANCE",
	163: "INTERSECTS",
	164: "INCLUDES",
	165: "CIRCLE",
	166: "GET_INTERSECTING",
	167: "FILL",
	168: "GET_NEAREST",
	171: "POLYGON_SUB",
}
var Term_TermType_value = map[string]int32{
	"DATUM":            1,
	"MAKE_ARRAY":       2,
	"MAKE_OBJ":         3,
	"VAR":              10,
	"JAVASCRIPT":       11,
	"UUID":             169,
	"HTTP":             153,
	"ERROR":            12,
	"IMPLICIT_VAR":     13,
	"DB":               14,
	"TABLE":            15,
	"GET":              16,
	"GET_ALL":          78,
	"EQ":               17,
	"NE":               18,
	"LT":               19,
	"LE":               20,
	"GT":               21,
	"GE":               22,
	"NOT":              23,
	"ADD":              24,
	"SUB":              25,
	"MUL":              26,
	"DIV":              27,
	"MOD":              28,
	"APPEND":           29,
	"PREPEND":          80,
	"DIFFERENCE":       95,
	"SET_INSERT":       88,
	"SET_INTERSECTION": 89,
	"SET_UNION":        90,
	"SET_DIFFERENCE":   91,
	"SLICE":            30,
	"SKIP":             70,
	"LIMIT":            71,
	"INDEXES_OF":       87,
	"CONTAINS":         93,
	"GET_FIELD":        31,
	"KEYS":             94,
	"OBJECT":           143,
	"HAS_FIELDS":       32,
	"WITH_FIELDS":      96,
	"PLUCK":            33,
	"WITHOUT":          34,
	"MERGE":            35,
	"BETWEEN":          36,
	"REDUCE":           37,
	"MAP":              38,
	"FILTER":           39,
	"CONCAT_MAP":       40,
	"ORDER_BY":         41,
	"DISTINCT":         42,
	"COUNT":            43,
	"IS_EMPTY":         86,
	"UNION":            44,
	"NTH":              45,
	"BRACKET":          170,
	"INNER_JOIN":       48,
	"OUTER_JOIN":       49,
	"EQ_JOIN":          50,
	"ZIP":              72,
	"RANGE":            173,
	"INSERT_AT":        82,
	"DELETE_AT":        83,
	"CHANGE_AT":        84,
	"SPLICE_AT":        85,
	"COERCE_TO":        51,
	"TYPE_OF":          52,
	"UPDATE":           53,
	"DELETE":           54,
	"REPLACE":          55,
	"INSERT":           56,
	"DB_CREATE":        57,
	"DB_DROP":          58,
	"DB_LIST":          59,
	"TABLE_CREATE":     60,
	"TABLE_DROP":       61,
	"TABLE_LIST":       62,
	"CONFIG":           174,
	"STATUS":           175,
	"WAIT":             177,
	"RECONFIGURE":      176,
	"REBALANCE":        179,
	"SYNC":             138,
	"INDEX_CREATE":     75,
	"INDEX_DROP":       76,
	"INDEX_LIST":       77,
	"INDEX_STATUS":     139,
	"INDEX_WAIT":       140,
	"INDEX_RENAME":     156,
	"FUNCALL":          64,
	"BRANCH":           65,
	"ANY":              66,
	"ALL":              67,
	"FOR_EACH":         68,
	"FUNC":             69,
	"ASC":              73,
	"DESC":             74,
	"INFO":             79,
	"MATCH":            97,
	"UPCASE":           141,
	"DOWNCASE":         142,
	"SAMPLE":           81,
	"DEFAULT":          92,
	"JSON":             98,
	"TO_JSON_STRING":   172,
	"ISO8601":          99,
	"TO_ISO8601":       100,
	"EPOCH_TIME":       101,
	"TO_EPOCH_TIME":    102,
	"NOW":              103,
	"IN_TIMEZONE":      104,
	"DURING":           105,
	"DATE":             106,
	"TIME_OF_DAY":      126,
	"TIMEZONE":         127,
	"YEAR":             128,
	"MONTH":            129,
	"DAY":              130,
	"DAY_OF_WEEK":      131,
	"DAY_OF_YEAR":      132,
	"HOURS":            133,
	"MINUTES":          134,
	"SECONDS":          135,
	"TIME":             136,
	"MONDAY":           107,
	"TUESDAY":          108,
	"WEDNESDAY":        109,
	"THURSDAY":         110,
	"FRIDAY":           111,
	"SATURDAY":         112,
	"SUNDAY":           113,
	"JANUARY":          114,
	"FEBRUARY":         115,
	"MARCH":            116,
	"APRIL":            117,
	"MAY":              118,
	"JUNE":             119,
	"JULY":             120,
	"AUGUST":           121,
	"SEPTEMBER":        122,
	"OCTOBER":          123,
	"NOVEMBER":         124,
	"DECEMBER":         125,
	"LITERAL":          137,
	"GROUP":            144,
	"SUM":              145,
	"AVG":              146,
	"MIN":              147,
	"MAX":              148,
	"SPLIT":            149,
	"UNGROUP":          150,
	"RANDOM":           151,
	"CHANGES":          152,
	"ARGS":             154,
	"BINARY":           155,
	"GEOJSON":          157,
	"TO_GEOJSON":       158,
	"POINT":            159,
	"LINE":             160,
	"POLYGON":          161,
	"DISTANCE":         162,
	"INTERSECTS":       163,
	"INCLUDES":         164,
	"CIRCLE":           165,
	"GET_INTERSECTING": 166,
	"FILL":             167,
	"GET_NEAREST":      168,
	"POLYGON_SUB":      171,
}

func (x Term_TermType) Enum() *Term_TermType {
	p := new(Term_TermType)
	*p = x
	return p
}
func (x Term_TermType) String() string {
	return proto.EnumName(Term_TermType_name, int32(x))
}
func (x Term_TermType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Term_TermType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Term_TermType_value, data, "Term_TermType")
	if err != nil {
		return err
	}
	*x = Term_TermType(value)
	return nil
}

type VersionDummy struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *VersionDummy) Reset()         { *m = VersionDummy{} }
func (m *VersionDummy) String() string { return proto.CompactTextString(m) }
func (*VersionDummy) ProtoMessage()    {}

type Query struct {
	Type             *Query_QueryType   `protobuf:"varint,1,opt,name=type,enum=Query_QueryType" json:"type,omitempty"`
	Query            *Term              `protobuf:"bytes,2,opt,name=query" json:"query,omitempty"`
	Token            *int64             `protobuf:"varint,3,opt,name=token" json:"token,omitempty"`
	OBSOLETENoreply  *bool              `protobuf:"varint,4,opt,name=OBSOLETE_noreply,def=0" json:"OBSOLETE_noreply,omitempty"`
	AcceptsRJson     *bool              `protobuf:"varint,5,opt,name=accepts_r_json,def=0" json:"accepts_r_json,omitempty"`
	GlobalOptargs    []*Query_AssocPair `protobuf:"bytes,6,rep,name=global_optargs" json:"global_optargs,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}

const Default_Query_OBSOLETENoreply bool = false
const Default_Query_AcceptsRJson bool = false

func (m *Query) GetType() Query_QueryType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Query) GetQuery() *Term {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *Query) GetToken() int64 {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return 0
}

func (m *Query) GetOBSOLETENoreply() bool {
	if m != nil && m.OBSOLETENoreply != nil {
		return *m.OBSOLETENoreply
	}
	return Default_Query_OBSOLETENoreply
}

func (m *Query) GetAcceptsRJson() bool {
	if m != nil && m.AcceptsRJson != nil {
		return *m.AcceptsRJson
	}
	return Default_Query_AcceptsRJson
}

func (m *Query) GetGlobalOptargs() []*Query_AssocPair {
	if m != nil {
		return m.GlobalOptargs
	}
	return nil
}

type Query_AssocPair struct {
	Key              *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Val              *Term   `protobuf:"bytes,2,opt,name=val" json:"val,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Query_AssocPair) Reset()         { *m = Query_AssocPair{} }
func (m *Query_AssocPair) String() string { return proto.CompactTextString(m) }
func (*Query_AssocPair) ProtoMessage()    {}

func (m *Query_AssocPair) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Query_AssocPair) GetVal() *Term {
	if m != nil {
		return m.Val
	}
	return nil
}

type Frame struct {
	Type             *Frame_FrameType `protobuf:"varint,1,opt,name=type,enum=Frame_FrameType" json:"type,omitempty"`
	Pos              *int64           `protobuf:"varint,2,opt,name=pos" json:"pos,omitempty"`
	Opt              *string          `protobuf:"bytes,3,opt,name=opt" json:"opt,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Frame) Reset()         { *m = Frame{} }
func (m *Frame) String() string { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()    {}

func (m *Frame) GetType() Frame_FrameType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Frame) GetPos() int64 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func (m *Frame) GetOpt() string {
	if m != nil && m.Opt != nil {
		return *m.Opt
	}
	return ""
}

type Backtrace struct {
	Frames           []*Frame `protobuf:"bytes,1,rep,name=frames" json:"frames,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Backtrace) Reset()         { *m = Backtrace{} }
func (m *Backtrace) String() string { return proto.CompactTextString(m) }
func (*Backtrace) ProtoMessage()    {}

func (m *Backtrace) GetFrames() []*Frame {
	if m != nil {
		return m.Frames
	}
	return nil
}

type Response struct {
	Type             *Response_ResponseType `protobuf:"varint,1,opt,name=type,enum=Response_ResponseType" json:"type,omitempty"`
	Token            *int64                 `protobuf:"varint,2,opt,name=token" json:"token,omitempty"`
	Response         []*Datum               `protobuf:"bytes,3,rep,name=response" json:"response,omitempty"`
	Backtrace        *Backtrace             `protobuf:"bytes,4,opt,name=backtrace" json:"backtrace,omitempty"`
	Profile          *Datum                 `protobuf:"bytes,5,opt,name=profile" json:"profile,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetType() Response_ResponseType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Response) GetToken() int64 {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return 0
}

func (m *Response) GetResponse() []*Datum {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Response) GetBacktrace() *Backtrace {
	if m != nil {
		return m.Backtrace
	}
	return nil
}

func (m *Response) GetProfile() *Datum {
	if m != nil {
		return m.Profile
	}
	return nil
}

type Datum struct {
	Type             *Datum_DatumType          `protobuf:"varint,1,opt,name=type,enum=Datum_DatumType" json:"type,omitempty"`
	RBool            *bool                     `protobuf:"varint,2,opt,name=r_bool" json:"r_bool,omitempty"`
	RNum             *float64                  `protobuf:"fixed64,3,opt,name=r_num" json:"r_num,omitempty"`
	RStr             *string                   `protobuf:"bytes,4,opt,name=r_str" json:"r_str,omitempty"`
	RArray           []*Datum                  `protobuf:"bytes,5,rep,name=r_array" json:"r_array,omitempty"`
	RObject          []*Datum_AssocPair        `protobuf:"bytes,6,rep,name=r_object" json:"r_object,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Datum) Reset()         { *m = Datum{} }
func (m *Datum) String() string { return proto.CompactTextString(m) }
func (*Datum) ProtoMessage()    {}

var extRange_Datum = []proto.ExtensionRange{
	{10000, 20000},
}

func (*Datum) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Datum
}
func (m *Datum) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *Datum) GetType() Datum_DatumType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Datum) GetRBool() bool {
	if m != nil && m.RBool != nil {
		return *m.RBool
	}
	return false
}

func (m *Datum) GetRNum() float64 {
	if m != nil && m.RNum != nil {
		return *m.RNum
	}
	return 0
}

func (m *Datum) GetRStr() string {
	if m != nil && m.RStr != nil {
		return *m.RStr
	}
	return ""
}

func (m *Datum) GetRArray() []*Datum {
	if m != nil {
		return m.RArray
	}
	return nil
}

func (m *Datum) GetRObject() []*Datum_AssocPair {
	if m != nil {
		return m.RObject
	}
	return nil
}

type Datum_AssocPair struct {
	Key              *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Val              *Datum  `protobuf:"bytes,2,opt,name=val" json:"val,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Datum_AssocPair) Reset()         { *m = Datum_AssocPair{} }
func (m *Datum_AssocPair) String() string { return proto.CompactTextString(m) }
func (*Datum_AssocPair) ProtoMessage()    {}

func (m *Datum_AssocPair) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Datum_AssocPair) GetVal() *Datum {
	if m != nil {
		return m.Val
	}
	return nil
}

type Term struct {
	Type             *Term_TermType            `protobuf:"varint,1,opt,name=type,enum=Term_TermType" json:"type,omitempty"`
	Datum            *Datum                    `protobuf:"bytes,2,opt,name=datum" json:"datum,omitempty"`
	Args             []*Term                   `protobuf:"bytes,3,rep,name=args" json:"args,omitempty"`
	Optargs          []*Term_AssocPair         `protobuf:"bytes,4,rep,name=optargs" json:"optargs,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Term) Reset()         { *m = Term{} }
func (m *Term) String() string { return proto.CompactTextString(m) }
func (*Term) ProtoMessage()    {}

var extRange_Term = []proto.ExtensionRange{
	{10000, 20000},
}

func (*Term) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Term
}
func (m *Term) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *Term) GetType() Term_TermType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Term) GetDatum() *Datum {
	if m != nil {
		return m.Datum
	}
	return nil
}

func (m *Term) GetArgs() []*Term {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Term) GetOptargs() []*Term_AssocPair {
	if m != nil {
		return m.Optargs
	}
	return nil
}

type Term_AssocPair struct {
	Key              *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Val              *Term   `protobuf:"bytes,2,opt,name=val" json:"val,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Term_AssocPair) Reset()         { *m = Term_AssocPair{} }
func (m *Term_AssocPair) String() string { return proto.CompactTextString(m) }
func (*Term_AssocPair) ProtoMessage()    {}

func (m *Term_AssocPair) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Term_AssocPair) GetVal() *Term {
	if m != nil {
		return m.Val
	}
	return nil
}

func init() {
	proto.RegisterEnum("VersionDummy_Version", VersionDummy_Version_name, VersionDummy_Version_value)
	proto.RegisterEnum("VersionDummy_Protocol", VersionDummy_Protocol_name, VersionDummy_Protocol_value)
	proto.RegisterEnum("Query_QueryType", Query_QueryType_name, Query_QueryType_value)
	proto.RegisterEnum("Frame_FrameType", Frame_FrameType_name, Frame_FrameType_value)
	proto.RegisterEnum("Response_ResponseType", Response_ResponseType_name, Response_ResponseType_value)
	proto.RegisterEnum("Datum_DatumType", Datum_DatumType_name, Datum_DatumType_value)
	proto.RegisterEnum("Term_TermType", Term_TermType_name, Term_TermType_value)
}

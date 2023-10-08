package engine

var asciiMap = map[string]int{
	"left_mouse":   0x01,
	"right_mouse":  0x02,
	"center_mouse": 0x04,
	"left_win":     0x5B,
	"right_win":    0x5C,
	"right_0":      0x60,
	"right_1":      0x61,
	"right_2":      0x62,
	"right_3":      0x63,
	"right_4":      0x64,
	"right_5":      0x65,
	"right_6":      0x66,
	"right_7":      0x67,
	"right_8":      0x68,
	"right_9":      0x69,
	"left_0":         0x30,
	"left_1":         0x31,
	"left_2":         0x32,
	"left_3":         0x33,
	"left_4":         0x34,
	"left_5":         0x35,
	"left_6":         0x36,
	"left_7":         0x37,
	"left_8":         0x38,
	"left_9":         0x39,
	"left_shift":   0xA0,
	"right_shift":  0xA1,
	"left_ctrl":    0xA2,
	"right_ctrl":   0xA3,
	"left_alt":     0xA4,
	"right_alt":    0xA5,

	"backspace":      0x08,
	"tab":            0x09,
	"num_lock_close": 0x0C,
	"enter":          0x0D,
	"shift":          0x10,
	"ctrl":           0x11,
	"alt":            0x12,
	"pause":          0x13,
	"caps_lock":      0x14,
	"esc":            0x1B,
	"space":          0x20,
	"pgup":           0x21,
	"pgdn":           0x22,
	"end":            0x23,
	"home":           0x24,
	"left":           0x25,
	"up":             0x26,
	"right":          0x27,
	"down":           0x28,
	"select_":        0x29,
	"print_":         0x2A,
	"execute":        0x2B,
	"print_screen":   0x2C,
	"insert":         0x2D,
	"delete":         0x2E,
	"help":           0x2F,

	"A": 0x41,
	"B": 0x42,
	"C": 0x43,
	"D": 0x44,
	"E": 0x45,
	"F": 0x46,
	"G": 0x47,
	"H": 0x48,
	"I": 0x49,
	"J": 0x4A,
	"K": 0x4B,
	"L": 0x4C,
	"M": 0x4D,
	"N": 0x4E,
	"O": 0x4F,
	"P": 0x50,
	"Q": 0x51,
	"R": 0x52,
	"S": 0x53,
	"T": 0x54,
	"U": 0x55,
	"V": 0x56,
	"W": 0x57,
	"X": 0x58,
	"Y": 0x59,
	"Z": 0x5A,

	"a": 0x41,
	"b": 0x42,
	"c": 0x43,
	"d": 0x44,
	"e": 0x45,
	"f": 0x46,
	"g": 0x47,
	"h": 0x48,
	"i": 0x49,
	"j": 0x4A,
	"k": 0x4B,
	"l": 0x4C,
	"m": 0x4D,
	"n": 0x4E,
	"o": 0x4F,
	"p": 0x50,
	"q": 0x51,
	"r": 0x52,
	"s": 0x53,
	"t": 0x54,
	"u": 0x55,
	"v": 0x56,
	"w": 0x57,
	"x": 0x58,
	"y": 0x59,
	"z": 0x5A,

	"multiply":    0x6A,
	"add":         0x6B,
	"sepatator":   0x6C,
	"subtract":    0x6D,
	"decimal":     0x6E,
	"divide":      0x6F,
	"F1":          0x70,
	"F2":          0x71,
	"F3":          0x72,
	"F4":          0x73,
	"F5":          0x74,
	"F6":          0x75,
	"F7":          0x76,
	"F8":          0x77,
	"F9":          0x78,
	"F10":         0x79,
	"F11":         0x7A,
	"F12":         0x7B,
	"F13":         0x7C,
	"F14":         0x7D,
	"F15":         0x7E,
	"F16":         0x7F,
	"F17":         0x80,
	"F18":         0x81,
	"F19":         0x82,
	"F20":         0x83,
	"F21":         0x84,
	"F22":         0x85,
	"F23":         0x86,
	"F24":         0x87,
	"num_lock":    0x90,
	"scroll_lock": 0x91,

	"OEM_1":      0xBA,
	"OEM_PLUS":   0xBB,
	"OEM_COMMA":  0xBC,
	"OEM_MINUS":  0xBD,
	"OEM_PERIOD": 0xBE,
	"OEM_2":      0xBF,
	"OEM_3":      0xC0,
	"OEM_4":      0xDB,
	"OEM_5":      0xDC,
	"OEM_6":      0xDD,
	"OEM_7":      0xDE,
}

// 键盘事件集合
const mouse_left = 0x01   // 鼠标左键
const mouse_right = 0x02  // 鼠标右键
const mouse_center = 0x04 // 鼠标中键

const backspace = 0x08      // BackSpace
const tab = 0x09            // Tab
const num_lock_close = 0x0C // Num Lock关闭时的数字区 5
const enter = 0x0D          // Enter
const shift = 0x10          // Shift/左右shift都可以
const ctrl = 0x11           // Ctrl/左右Ctrl都可以
const alt = 0x12            // Alt/左右Alt都可以
const pause = 0x13          /* Pause*/
const caps_lock = 0x14      // Caps Lock
const esc = 0x1B            // Esc
const space = 0x20          // Space
const pgup = 0x21           // PgUp
const pgdn = 0x22           // PgDn
const end = 0x23            // End
const home = 0x24           // Home
const left = 0x25           // 方向左
const up = 0x26             // 方向上
const right = 0x27          // 方向右
const down = 0x28           // 方向下
const select_ = 0x29        /* Select*/
const print_ = 0x2A         /* Print*/
const execute = 0x2B        /* Execute*/
const print_screen = 0x2C   /* Print Screen键(抓屏)*/
const insert = 0x2D         // Insert
const delete = 0x2E         // Delete
const help = 0x2F           /* Help*/

const left_0 = 0x30 // '0'
const left_1 = 0x31 // '1'
const left_2 = 0x32 // '2'
const left_3 = 0x33 // '3'
const left_4 = 0x34 // '4'
const left_5 = 0x35 // '5'
const left_6 = 0x36 // '6'
const left_7 = 0x37 // '7'
const left_8 = 0x38 // '8'
const left_9 = 0x39 // '9'

const A = 0x41 // 'A'
const B = 0x42 // 'B'
const C = 0x43 // 'C'
const D = 0x44 // 'D'
const E = 0x45 // 'E'
const F = 0x46 // 'F'
const G = 0x47 // 'G'
const H = 0x48 // 'H'
const I = 0x49 // 'I'
const J = 0x4A // 'J'
const K = 0x4B // 'K'
const L = 0x4C // 'L'
const M = 0x4D // 'M'
const N = 0x4E // 'N'
const O = 0x4F // 'O'
const P = 0x50 // 'P'
const Q = 0x51 // 'Q'
const R = 0x52 // 'R'
const S = 0x53 // 'S'
const T = 0x54 // 'T'
const U = 0x55 // 'U'
const V = 0x56 // 'V'
const W = 0x57 // 'W'
const X = 0x58 // 'X'
const Y = 0x59 // 'Y'
const Z = 0x5A // 'Z'

const left_win = 0x5B  // 左win键
const right_win = 0x5C // 右win键

const right_0 = 0x60  // 数字区 0
const right_1 = 0x61  // 数字区 1
const right_2 = 0x62  // 数字区 2
const right_3 = 0x63  // 数字区 3
const right_4 = 0x64  // 数字区 4
const right_5 = 0x65  // 数字区 5
const right_6 = 0x66  // 数字区 6
const right_7 = 0x67  // 数字区 7
const right_8 = 0x68  // 数字区 8
const right_9 = 0x69  // 数字区 9
const multiply = 0x6A // 数字区 *
const add = 0x6B      // 数字区 +

const sepatator = 0x6C /* Sepatator*/
const subtract = 0x6D  // 数字区 -
const decimal = 0x6E   // 数字区 小数点
const divide = 0x6F    // 数字区 /

const F1 = 0x70  // F1
const F2 = 0x71  // F2
const F3 = 0x72  // F3
const F4 = 0x73  // F4
const F5 = 0x74  // F5
const F6 = 0x75  // F6
const F7 = 0x76  // F7
const F8 = 0x77  // F8
const F9 = 0x78  // F9
const F10 = 0x79 // F10
const F11 = 0x7A // F11
const F12 = 0x7B // F12
const F13 = 0x7C // F13
const F14 = 0x7D // F14
const F15 = 0x7E // F15
const F16 = 0x7F // F16
const F17 = 0x80 // F17
const F18 = 0x81 // F18
const F19 = 0x82 // F19
const F20 = 0x83 // F20
const F21 = 0x84 // F21
const F22 = 0x85 // F22
const F23 = 0x86 // F23
const F24 = 0x87 // F24

const num_lock = 0x90    // Num Lock
const scroll_lock = 0x91 /* Scroll Lock*/
const left_shift = 0xA0  // 左shift
const right_shift = 0xA1 // 右shift
const left_ctrl = 0xA2   // 左Ctrl
const right_ctrl = 0xA3  // 右Ctrl
const left_alt = 0xA4    // 左Alt
const right_alt = 0xA5   // 右Alt
const OEM_1 = 0xBA       // ';:'
const OEM_PLUS = 0xBB    // '=+'
const OEM_COMMA = 0xBC   // ',<'
const OEM_MINUS = 0xBD   // '-_'
const OEM_PERIOD = 0xBE  // '.>'
const OEM_2 = 0xBF       // '/?'
const OEM_3 = 0xC0       // '`~'
const OEM_4 = 0xDB       // '[{'
const OEM_5 = 0xDC       // '\|'
const OEM_6 = 0xDD       // ']}'
const OEM_7 = 0xDE       // ''"'

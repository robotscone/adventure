package input

import "github.com/veandco/go-sdl2/sdl"

var scancodes = map[string]int{
	"keyboard:unknown":            sdl.SCANCODE_UNKNOWN,
	"keyboard:a":                  sdl.SCANCODE_A,
	"keyboard:b":                  sdl.SCANCODE_B,
	"keyboard:c":                  sdl.SCANCODE_C,
	"keyboard:d":                  sdl.SCANCODE_D,
	"keyboard:e":                  sdl.SCANCODE_E,
	"keyboard:f":                  sdl.SCANCODE_F,
	"keyboard:g":                  sdl.SCANCODE_G,
	"keyboard:h":                  sdl.SCANCODE_H,
	"keyboard:i":                  sdl.SCANCODE_I,
	"keyboard:j":                  sdl.SCANCODE_J,
	"keyboard:k":                  sdl.SCANCODE_K,
	"keyboard:l":                  sdl.SCANCODE_L,
	"keyboard:m":                  sdl.SCANCODE_M,
	"keyboard:n":                  sdl.SCANCODE_N,
	"keyboard:o":                  sdl.SCANCODE_O,
	"keyboard:p":                  sdl.SCANCODE_P,
	"keyboard:q":                  sdl.SCANCODE_Q,
	"keyboard:r":                  sdl.SCANCODE_R,
	"keyboard:s":                  sdl.SCANCODE_S,
	"keyboard:t":                  sdl.SCANCODE_T,
	"keyboard:u":                  sdl.SCANCODE_U,
	"keyboard:v":                  sdl.SCANCODE_V,
	"keyboard:w":                  sdl.SCANCODE_W,
	"keyboard:x":                  sdl.SCANCODE_X,
	"keyboard:y":                  sdl.SCANCODE_Y,
	"keyboard:z":                  sdl.SCANCODE_Z,
	"keyboard:1":                  sdl.SCANCODE_1,
	"keyboard:2":                  sdl.SCANCODE_2,
	"keyboard:3":                  sdl.SCANCODE_3,
	"keyboard:4":                  sdl.SCANCODE_4,
	"keyboard:5":                  sdl.SCANCODE_5,
	"keyboard:6":                  sdl.SCANCODE_6,
	"keyboard:7":                  sdl.SCANCODE_7,
	"keyboard:8":                  sdl.SCANCODE_8,
	"keyboard:9":                  sdl.SCANCODE_9,
	"keyboard:0":                  sdl.SCANCODE_0,
	"keyboard:return":             sdl.SCANCODE_RETURN,
	"keyboard:escape":             sdl.SCANCODE_ESCAPE,
	"keyboard:backspace":          sdl.SCANCODE_BACKSPACE,
	"keyboard:tab":                sdl.SCANCODE_TAB,
	"keyboard:space":              sdl.SCANCODE_SPACE,
	"keyboard:minus":              sdl.SCANCODE_MINUS,
	"keyboard:equals":             sdl.SCANCODE_EQUALS,
	"keyboard:leftbracket":        sdl.SCANCODE_LEFTBRACKET,
	"keyboard:rightbracket":       sdl.SCANCODE_RIGHTBRACKET,
	"keyboard:backslash":          sdl.SCANCODE_BACKSLASH,
	"keyboard:nonushash":          sdl.SCANCODE_NONUSHASH,
	"keyboard:semicolon":          sdl.SCANCODE_SEMICOLON,
	"keyboard:apostrophe":         sdl.SCANCODE_APOSTROPHE,
	"keyboard:grave":              sdl.SCANCODE_GRAVE,
	"keyboard:comma":              sdl.SCANCODE_COMMA,
	"keyboard:period":             sdl.SCANCODE_PERIOD,
	"keyboard:slash":              sdl.SCANCODE_SLASH,
	"keyboard:capslock":           sdl.SCANCODE_CAPSLOCK,
	"keyboard:f1":                 sdl.SCANCODE_F1,
	"keyboard:f2":                 sdl.SCANCODE_F2,
	"keyboard:f3":                 sdl.SCANCODE_F3,
	"keyboard:f4":                 sdl.SCANCODE_F4,
	"keyboard:f5":                 sdl.SCANCODE_F5,
	"keyboard:f6":                 sdl.SCANCODE_F6,
	"keyboard:f7":                 sdl.SCANCODE_F7,
	"keyboard:f8":                 sdl.SCANCODE_F8,
	"keyboard:f9":                 sdl.SCANCODE_F9,
	"keyboard:f10":                sdl.SCANCODE_F10,
	"keyboard:f11":                sdl.SCANCODE_F11,
	"keyboard:f12":                sdl.SCANCODE_F12,
	"keyboard:printscreen":        sdl.SCANCODE_PRINTSCREEN,
	"keyboard:scrolllock":         sdl.SCANCODE_SCROLLLOCK,
	"keyboard:pause":              sdl.SCANCODE_PAUSE,
	"keyboard:insert":             sdl.SCANCODE_INSERT,
	"keyboard:home":               sdl.SCANCODE_HOME,
	"keyboard:pageup":             sdl.SCANCODE_PAGEUP,
	"keyboard:delete":             sdl.SCANCODE_DELETE,
	"keyboard:end":                sdl.SCANCODE_END,
	"keyboard:pagedown":           sdl.SCANCODE_PAGEDOWN,
	"keyboard:right":              sdl.SCANCODE_RIGHT,
	"keyboard:left":               sdl.SCANCODE_LEFT,
	"keyboard:down":               sdl.SCANCODE_DOWN,
	"keyboard:up":                 sdl.SCANCODE_UP,
	"keyboard:numlockclear":       sdl.SCANCODE_NUMLOCKCLEAR,
	"keyboard:kp:divide":          sdl.SCANCODE_KP_DIVIDE,
	"keyboard:kp:multiply":        sdl.SCANCODE_KP_MULTIPLY,
	"keyboard:kp:minus":           sdl.SCANCODE_KP_MINUS,
	"keyboard:kp:plus":            sdl.SCANCODE_KP_PLUS,
	"keyboard:kp:enter":           sdl.SCANCODE_KP_ENTER,
	"keyboard:kp:1":               sdl.SCANCODE_KP_1,
	"keyboard:kp:2":               sdl.SCANCODE_KP_2,
	"keyboard:kp:3":               sdl.SCANCODE_KP_3,
	"keyboard:kp:4":               sdl.SCANCODE_KP_4,
	"keyboard:kp:5":               sdl.SCANCODE_KP_5,
	"keyboard:kp:6":               sdl.SCANCODE_KP_6,
	"keyboard:kp:7":               sdl.SCANCODE_KP_7,
	"keyboard:kp:8":               sdl.SCANCODE_KP_8,
	"keyboard:kp:9":               sdl.SCANCODE_KP_9,
	"keyboard:kp:0":               sdl.SCANCODE_KP_0,
	"keyboard:kp:period":          sdl.SCANCODE_KP_PERIOD,
	"keyboard:nonusbackslash":     sdl.SCANCODE_NONUSBACKSLASH,
	"keyboard:application":        sdl.SCANCODE_APPLICATION,
	"keyboard:power":              sdl.SCANCODE_POWER,
	"keyboard:kp:equals":          sdl.SCANCODE_KP_EQUALS,
	"keyboard:f13":                sdl.SCANCODE_F13,
	"keyboard:f14":                sdl.SCANCODE_F14,
	"keyboard:f15":                sdl.SCANCODE_F15,
	"keyboard:f16":                sdl.SCANCODE_F16,
	"keyboard:f17":                sdl.SCANCODE_F17,
	"keyboard:f18":                sdl.SCANCODE_F18,
	"keyboard:f19":                sdl.SCANCODE_F19,
	"keyboard:f20":                sdl.SCANCODE_F20,
	"keyboard:f21":                sdl.SCANCODE_F21,
	"keyboard:f22":                sdl.SCANCODE_F22,
	"keyboard:f23":                sdl.SCANCODE_F23,
	"keyboard:f24":                sdl.SCANCODE_F24,
	"keyboard:execute":            sdl.SCANCODE_EXECUTE,
	"keyboard:help":               sdl.SCANCODE_HELP,
	"keyboard:menu":               sdl.SCANCODE_MENU,
	"keyboard:select":             sdl.SCANCODE_SELECT,
	"keyboard:stop":               sdl.SCANCODE_STOP,
	"keyboard:again":              sdl.SCANCODE_AGAIN,
	"keyboard:undo":               sdl.SCANCODE_UNDO,
	"keyboard:cut":                sdl.SCANCODE_CUT,
	"keyboard:copy":               sdl.SCANCODE_COPY,
	"keyboard:paste":              sdl.SCANCODE_PASTE,
	"keyboard:find":               sdl.SCANCODE_FIND,
	"keyboard:mute":               sdl.SCANCODE_MUTE,
	"keyboard:volumeup":           sdl.SCANCODE_VOLUMEUP,
	"keyboard:volumedown":         sdl.SCANCODE_VOLUMEDOWN,
	"keyboard:kp:comma":           sdl.SCANCODE_KP_COMMA,
	"keyboard:kp:equalsas400":     sdl.SCANCODE_KP_EQUALSAS400,
	"keyboard:international1":     sdl.SCANCODE_INTERNATIONAL1,
	"keyboard:international2":     sdl.SCANCODE_INTERNATIONAL2,
	"keyboard:international3":     sdl.SCANCODE_INTERNATIONAL3,
	"keyboard:international4":     sdl.SCANCODE_INTERNATIONAL4,
	"keyboard:international5":     sdl.SCANCODE_INTERNATIONAL5,
	"keyboard:international6":     sdl.SCANCODE_INTERNATIONAL6,
	"keyboard:international7":     sdl.SCANCODE_INTERNATIONAL7,
	"keyboard:international8":     sdl.SCANCODE_INTERNATIONAL8,
	"keyboard:international9":     sdl.SCANCODE_INTERNATIONAL9,
	"keyboard:lang1":              sdl.SCANCODE_LANG1,
	"keyboard:lang2":              sdl.SCANCODE_LANG2,
	"keyboard:lang3":              sdl.SCANCODE_LANG3,
	"keyboard:lang4":              sdl.SCANCODE_LANG4,
	"keyboard:lang5":              sdl.SCANCODE_LANG5,
	"keyboard:lang6":              sdl.SCANCODE_LANG6,
	"keyboard:lang7":              sdl.SCANCODE_LANG7,
	"keyboard:lang8":              sdl.SCANCODE_LANG8,
	"keyboard:lang9":              sdl.SCANCODE_LANG9,
	"keyboard:alterase":           sdl.SCANCODE_ALTERASE,
	"keyboard:sysreq":             sdl.SCANCODE_SYSREQ,
	"keyboard:cancel":             sdl.SCANCODE_CANCEL,
	"keyboard:clear":              sdl.SCANCODE_CLEAR,
	"keyboard:prior":              sdl.SCANCODE_PRIOR,
	"keyboard:return2":            sdl.SCANCODE_RETURN2,
	"keyboard:separator":          sdl.SCANCODE_SEPARATOR,
	"keyboard:out":                sdl.SCANCODE_OUT,
	"keyboard:oper":               sdl.SCANCODE_OPER,
	"keyboard:clearagain":         sdl.SCANCODE_CLEARAGAIN,
	"keyboard:crsel":              sdl.SCANCODE_CRSEL,
	"keyboard:exsel":              sdl.SCANCODE_EXSEL,
	"keyboard:kp:00":              sdl.SCANCODE_KP_00,
	"keyboard:kp:000":             sdl.SCANCODE_KP_000,
	"keyboard:thousandsseparator": sdl.SCANCODE_THOUSANDSSEPARATOR,
	"keyboard:decimalseparator":   sdl.SCANCODE_DECIMALSEPARATOR,
	"keyboard:currencyunit":       sdl.SCANCODE_CURRENCYUNIT,
	"keyboard:currencysubunit":    sdl.SCANCODE_CURRENCYSUBUNIT,
	"keyboard:kp:leftparen":       sdl.SCANCODE_KP_LEFTPAREN,
	"keyboard:kp:rightparen":      sdl.SCANCODE_KP_RIGHTPAREN,
	"keyboard:kp:leftbrace":       sdl.SCANCODE_KP_LEFTBRACE,
	"keyboard:kp:rightbrace":      sdl.SCANCODE_KP_RIGHTBRACE,
	"keyboard:kp:tab":             sdl.SCANCODE_KP_TAB,
	"keyboard:kp:backspace":       sdl.SCANCODE_KP_BACKSPACE,
	"keyboard:kp:a":               sdl.SCANCODE_KP_A,
	"keyboard:kp:b":               sdl.SCANCODE_KP_B,
	"keyboard:kp:c":               sdl.SCANCODE_KP_C,
	"keyboard:kp:d":               sdl.SCANCODE_KP_D,
	"keyboard:kp:e":               sdl.SCANCODE_KP_E,
	"keyboard:kp:f":               sdl.SCANCODE_KP_F,
	"keyboard:kp:xor":             sdl.SCANCODE_KP_XOR,
	"keyboard:kp:power":           sdl.SCANCODE_KP_POWER,
	"keyboard:kp:percent":         sdl.SCANCODE_KP_PERCENT,
	"keyboard:kp:less":            sdl.SCANCODE_KP_LESS,
	"keyboard:kp:greater":         sdl.SCANCODE_KP_GREATER,
	"keyboard:kp:ampersand":       sdl.SCANCODE_KP_AMPERSAND,
	"keyboard:kp:dblampersand":    sdl.SCANCODE_KP_DBLAMPERSAND,
	"keyboard:kp:verticalbar":     sdl.SCANCODE_KP_VERTICALBAR,
	"keyboard:kp:dblverticalbar":  sdl.SCANCODE_KP_DBLVERTICALBAR,
	"keyboard:kp:colon":           sdl.SCANCODE_KP_COLON,
	"keyboard:kp:hash":            sdl.SCANCODE_KP_HASH,
	"keyboard:kp:space":           sdl.SCANCODE_KP_SPACE,
	"keyboard:kp:at":              sdl.SCANCODE_KP_AT,
	"keyboard:kp:exclam":          sdl.SCANCODE_KP_EXCLAM,
	"keyboard:kp:memstore":        sdl.SCANCODE_KP_MEMSTORE,
	"keyboard:kp:memrecall":       sdl.SCANCODE_KP_MEMRECALL,
	"keyboard:kp:memclear":        sdl.SCANCODE_KP_MEMCLEAR,
	"keyboard:kp:memadd":          sdl.SCANCODE_KP_MEMADD,
	"keyboard:kp:memsubtract":     sdl.SCANCODE_KP_MEMSUBTRACT,
	"keyboard:kp:memmultiply":     sdl.SCANCODE_KP_MEMMULTIPLY,
	"keyboard:kp:memdivide":       sdl.SCANCODE_KP_MEMDIVIDE,
	"keyboard:kp:plusminus":       sdl.SCANCODE_KP_PLUSMINUS,
	"keyboard:kp:clear":           sdl.SCANCODE_KP_CLEAR,
	"keyboard:kp:clearentry":      sdl.SCANCODE_KP_CLEARENTRY,
	"keyboard:kp:binary":          sdl.SCANCODE_KP_BINARY,
	"keyboard:kp:octal":           sdl.SCANCODE_KP_OCTAL,
	"keyboard:kp:decimal":         sdl.SCANCODE_KP_DECIMAL,
	"keyboard:kp:hexadecimal":     sdl.SCANCODE_KP_HEXADECIMAL,
	"keyboard:lctrl":              sdl.SCANCODE_LCTRL,
	"keyboard:lshift":             sdl.SCANCODE_LSHIFT,
	"keyboard:lalt":               sdl.SCANCODE_LALT,
	"keyboard:lgui":               sdl.SCANCODE_LGUI,
	"keyboard:rctrl":              sdl.SCANCODE_RCTRL,
	"keyboard:rshift":             sdl.SCANCODE_RSHIFT,
	"keyboard:ralt":               sdl.SCANCODE_RALT,
	"keyboard:rgui":               sdl.SCANCODE_RGUI,
	"keyboard:mode":               sdl.SCANCODE_MODE,
	"keyboard:audionext":          sdl.SCANCODE_AUDIONEXT,
	"keyboard:audioprev":          sdl.SCANCODE_AUDIOPREV,
	"keyboard:audiostop":          sdl.SCANCODE_AUDIOSTOP,
	"keyboard:audioplay":          sdl.SCANCODE_AUDIOPLAY,
	"keyboard:audiomute":          sdl.SCANCODE_AUDIOMUTE,
	"keyboard:mediaselect":        sdl.SCANCODE_MEDIASELECT,
	"keyboard:www":                sdl.SCANCODE_WWW,
	"keyboard:mail":               sdl.SCANCODE_MAIL,
	"keyboard:calculator":         sdl.SCANCODE_CALCULATOR,
	"keyboard:computer":           sdl.SCANCODE_COMPUTER,
	"keyboard:ac:search":          sdl.SCANCODE_AC_SEARCH,
	"keyboard:ac:home":            sdl.SCANCODE_AC_HOME,
	"keyboard:ac:back":            sdl.SCANCODE_AC_BACK,
	"keyboard:ac:forward":         sdl.SCANCODE_AC_FORWARD,
	"keyboard:ac:stop":            sdl.SCANCODE_AC_STOP,
	"keyboard:ac:refresh":         sdl.SCANCODE_AC_REFRESH,
	"keyboard:ac:bookmarks":       sdl.SCANCODE_AC_BOOKMARKS,
	"keyboard:brightnessdown":     sdl.SCANCODE_BRIGHTNESSDOWN,
	"keyboard:brightnessup":       sdl.SCANCODE_BRIGHTNESSUP,
	"keyboard:displayswitch":      sdl.SCANCODE_DISPLAYSWITCH,
	"keyboard:kbdillumtoggle":     sdl.SCANCODE_KBDILLUMTOGGLE,
	"keyboard:kbdillumdown":       sdl.SCANCODE_KBDILLUMDOWN,
	"keyboard:kbdillumup":         sdl.SCANCODE_KBDILLUMUP,
	"keyboard:eject":              sdl.SCANCODE_EJECT,
	"keyboard:sleep":              sdl.SCANCODE_SLEEP,
	"keyboard:app1":               sdl.SCANCODE_APP1,
	"keyboard:app2":               sdl.SCANCODE_APP2,
}

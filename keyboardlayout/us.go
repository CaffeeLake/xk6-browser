/*
 *
 * xk6-browser - a browser automation extension for k6
 * Copyright (C) 2021 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package keyboardlayout

func initUS() {
	validKeys := map[KeyInput]bool{
		"0":                  true,
		"1":                  true,
		"2":                  true,
		"3":                  true,
		"4":                  true,
		"5":                  true,
		"6":                  true,
		"7":                  true,
		"8":                  true,
		"9":                  true,
		"Power":              true,
		"Eject":              true,
		"Abort":              true,
		"Help":               true,
		"Backspace":          true,
		"Tab":                true,
		"Numpad5":            true,
		"NumpadEnter":        true,
		"Enter":              true,
		"\r":                 true,
		"\n":                 true,
		"ShiftLeft":          true,
		"ShiftRight":         true,
		"ControlLeft":        true,
		"ControlRight":       true,
		"AltLeft":            true,
		"AltRight":           true,
		"Pause":              true,
		"CapsLock":           true,
		"Escape":             true,
		"Convert":            true,
		"NonConvert":         true,
		"Space":              true,
		"Numpad9":            true,
		"PageUp":             true,
		"Numpad3":            true,
		"PageDown":           true,
		"End":                true,
		"Numpad1":            true,
		"Home":               true,
		"Numpad7":            true,
		"ArrowLeft":          true,
		"Numpad4":            true,
		"Numpad8":            true,
		"ArrowUp":            true,
		"ArrowRight":         true,
		"Numpad6":            true,
		"Numpad2":            true,
		"ArrowDown":          true,
		"Select":             true,
		"Open":               true,
		"PrintScreen":        true,
		"Insert":             true,
		"Numpad0":            true,
		"Delete":             true,
		"NumpadDecimal":      true,
		"Digit0":             true,
		"Digit1":             true,
		"Digit2":             true,
		"Digit3":             true,
		"Digit4":             true,
		"Digit5":             true,
		"Digit6":             true,
		"Digit7":             true,
		"Digit8":             true,
		"Digit9":             true,
		"KeyA":               true,
		"KeyB":               true,
		"KeyC":               true,
		"KeyD":               true,
		"KeyE":               true,
		"KeyF":               true,
		"KeyG":               true,
		"KeyH":               true,
		"KeyI":               true,
		"KeyJ":               true,
		"KeyK":               true,
		"KeyL":               true,
		"KeyM":               true,
		"KeyN":               true,
		"KeyO":               true,
		"KeyP":               true,
		"KeyQ":               true,
		"KeyR":               true,
		"KeyS":               true,
		"KeyT":               true,
		"KeyU":               true,
		"KeyV":               true,
		"KeyW":               true,
		"KeyX":               true,
		"KeyY":               true,
		"KeyZ":               true,
		"MetaLeft":           true,
		"MetaRight":          true,
		"ConTextMenu":        true,
		"NumpadMultiply":     true,
		"NumpadAdd":          true,
		"NumpadSubtract":     true,
		"NumpadDivide":       true,
		"F1":                 true,
		"F2":                 true,
		"F3":                 true,
		"F4":                 true,
		"F5":                 true,
		"F6":                 true,
		"F7":                 true,
		"F8":                 true,
		"F9":                 true,
		"F10":                true,
		"F11":                true,
		"F12":                true,
		"F13":                true,
		"F14":                true,
		"F15":                true,
		"F16":                true,
		"F17":                true,
		"F18":                true,
		"F19":                true,
		"F20":                true,
		"F21":                true,
		"F22":                true,
		"F23":                true,
		"F24":                true,
		"NumLock":            true,
		"ScrollLock":         true,
		"AudioVolumeMute":    true,
		"AudioVolumeDown":    true,
		"AudioVolumeUp":      true,
		"MediaTrackNext":     true,
		"MediaTrackPrevious": true,
		"MediaStop":          true,
		"MediaPlayPause":     true,
		"Semicolon":          true,
		"Equal":              true,
		"NumpadEqual":        true,
		"Comma":              true,
		"Minus":              true,
		"Period":             true,
		"Slash":              true,
		"Backquote":          true,
		"BracketLeft":        true,
		"Backslash":          true,
		"BracketRight":       true,
		"Quote":              true,
		"AltGraph":           true,
		"Props":              true,
		"Cancel":             true,
		"Clear":              true,
		"Shift":              true,
		"Control":            true,
		"Alt":                true,
		"Accept":             true,
		"ModeChange":         true,
		" ":                  true,
		"Print":              true,
		"Execute":            true,
		"\u0000":             true,
		"a":                  true,
		"b":                  true,
		"c":                  true,
		"d":                  true,
		"e":                  true,
		"f":                  true,
		"g":                  true,
		"h":                  true,
		"i":                  true,
		"j":                  true,
		"k":                  true,
		"l":                  true,
		"m":                  true,
		"n":                  true,
		"o":                  true,
		"p":                  true,
		"q":                  true,
		"r":                  true,
		"s":                  true,
		"t":                  true,
		"u":                  true,
		"v":                  true,
		"w":                  true,
		"x":                  true,
		"y":                  true,
		"z":                  true,
		"Meta":               true,
		"*":                  true,
		"+":                  true,
		"-":                  true,
		"/":                  true,
		";":                  true,
		"=":                  true,
		",":                  true,
		".":                  true,
		"`":                  true,
		"[":                  true,
		"\\":                 true,
		"]":                  true,
		"'":                  true,
		"Attn":               true,
		"CrSel":              true,
		"ExSel":              true,
		"EraseEof":           true,
		"Play":               true,
		"ZoomOut":            true,
		")":                  true,
		"!":                  true,
		"@":                  true,
		"#":                  true,
		"$":                  true,
		"%":                  true,
		"^":                  true,
		"&":                  true,
		"(":                  true,
		"A":                  true,
		"B":                  true,
		"C":                  true,
		"D":                  true,
		"E":                  true,
		"F":                  true,
		"G":                  true,
		"H":                  true,
		"I":                  true,
		"J":                  true,
		"K":                  true,
		"L":                  true,
		"M":                  true,
		"N":                  true,
		"O":                  true,
		"P":                  true,
		"Q":                  true,
		"R":                  true,
		"S":                  true,
		"T":                  true,
		"U":                  true,
		"V":                  true,
		"W":                  true,
		"X":                  true,
		"Y":                  true,
		"Z":                  true,
		":":                  true,
		"<":                  true,
		"_":                  true,
		">":                  true,
		"?":                  true,
		"~":                  true,
		"{":                  true,
		"|":                  true,
		"}":                  true,
		"\"":                 true,
		"SoftLeft":           true,
		"SoftRight":          true,
		"Camera":             true,
		"Call":               true,
		"EndCall":            true,
		"VolumeDown":         true,
		"VolumeUp":           true,
	}
	Keys := map[KeyInput]KeyDefinition{
		// Functions row
		"Escape": {KeyCode: 27, Key: "Escape"},
		"F1":     {KeyCode: 112, Key: "F1"},
		"F2":     {KeyCode: 113, Key: "F2"},
		"F3":     {KeyCode: 114, Key: "F3"},
		"F4":     {KeyCode: 115, Key: "F4"},
		"F5":     {KeyCode: 116, Key: "F5"},
		"F6":     {KeyCode: 117, Key: "F6"},
		"F7":     {KeyCode: 118, Key: "F7"},
		"F8":     {KeyCode: 119, Key: "F8"},
		"F9":     {KeyCode: 120, Key: "F9"},
		"F10":    {KeyCode: 121, Key: "F10"},
		"F11":    {KeyCode: 122, Key: "F11"},
		"F12":    {KeyCode: 123, Key: "F12"},

		// Numbers row
		"Backquote": {KeyCode: 192, ShiftKey: "~", Key: "`"},
		"Digit1":    {KeyCode: 49, ShiftKey: "!", Key: "1"},
		"Digit2":    {KeyCode: 50, ShiftKey: "@", Key: "2"},
		"Digit3":    {KeyCode: 51, ShiftKey: "#", Key: "3"},
		"Digit4":    {KeyCode: 52, ShiftKey: "$", Key: "4"},
		"Digit5":    {KeyCode: 53, ShiftKey: "%", Key: "5"},
		"Digit6":    {KeyCode: 54, ShiftKey: "^", Key: "6"},
		"Digit7":    {KeyCode: 55, ShiftKey: "&", Key: "7"},
		"Digit8":    {KeyCode: 56, ShiftKey: "*", Key: "8"},
		"Digit9":    {KeyCode: 57, ShiftKey: "(", Key: "9"},
		"Digit0":    {KeyCode: 48, ShiftKey: ")", Key: "0"},
		"Minus":     {KeyCode: 189, ShiftKey: "_", Key: "-"},
		"Equal":     {KeyCode: 187, ShiftKey: "+", Key: "="},
		"Backslash": {KeyCode: 220, ShiftKey: "|", Key: "\\"},
		"Backspace": {KeyCode: 8, Key: "Backspace"},

		// First row
		"Tab":          {KeyCode: 9, Key: "Tab"},
		"KeyQ":         {KeyCode: 81, ShiftKey: "Q", Key: "q"},
		"KeyW":         {KeyCode: 87, ShiftKey: "W", Key: "w"},
		"KeyE":         {KeyCode: 69, ShiftKey: "E", Key: "e"},
		"KeyR":         {KeyCode: 82, ShiftKey: "R", Key: "r"},
		"KeyT":         {KeyCode: 84, ShiftKey: "T", Key: "t"},
		"KeyY":         {KeyCode: 89, ShiftKey: "Y", Key: "y"},
		"KeyU":         {KeyCode: 85, ShiftKey: "U", Key: "u"},
		"KeyI":         {KeyCode: 73, ShiftKey: "I", Key: "i"},
		"KeyO":         {KeyCode: 79, ShiftKey: "O", Key: "o"},
		"KeyP":         {KeyCode: 80, ShiftKey: "P", Key: "p"},
		"BracketLeft":  {KeyCode: 219, ShiftKey: "{", Key: "["},
		"BracketRight": {KeyCode: 221, ShiftKey: "}", Key: "]"},

		// Second row
		"CapsLock":  {KeyCode: 20, Key: "CapsLock"},
		"KeyA":      {KeyCode: 65, ShiftKey: "A", Key: "a"},
		"KeyS":      {KeyCode: 83, ShiftKey: "S", Key: "s"},
		"KeyD":      {KeyCode: 68, ShiftKey: "D", Key: "d"},
		"KeyF":      {KeyCode: 70, ShiftKey: "F", Key: "f"},
		"KeyG":      {KeyCode: 71, ShiftKey: "G", Key: "g"},
		"KeyH":      {KeyCode: 72, ShiftKey: "H", Key: "h"},
		"KeyJ":      {KeyCode: 74, ShiftKey: "J", Key: "j"},
		"KeyK":      {KeyCode: 75, ShiftKey: "K", Key: "k"},
		"KeyL":      {KeyCode: 76, ShiftKey: "L", Key: "l"},
		"Semicolon": {KeyCode: 186, ShiftKey: ":", Key: ";"},
		"Quote":     {KeyCode: 222, ShiftKey: "\"", Key: "'"},
		"Enter":     {KeyCode: 13, Key: "Enter", Text: "\r"},

		// Third row
		"ShiftLeft":  {KeyCode: 160, KeyCodeWithoutLocation: 16, Key: "Shift", Location: 1},
		"KeyZ":       {KeyCode: 90, ShiftKey: "Z", Key: "z"},
		"KeyX":       {KeyCode: 88, ShiftKey: "X", Key: "x"},
		"KeyC":       {KeyCode: 67, ShiftKey: "C", Key: "c"},
		"KeyV":       {KeyCode: 86, ShiftKey: "V", Key: "v"},
		"KeyB":       {KeyCode: 66, ShiftKey: "B", Key: "b"},
		"KeyN":       {KeyCode: 78, ShiftKey: "N", Key: "n"},
		"KeyM":       {KeyCode: 77, ShiftKey: "M", Key: "m"},
		"Comma":      {KeyCode: 188, ShiftKey: "<", Key: ","},
		"Period":     {KeyCode: 190, ShiftKey: ">", Key: "."},
		"Slash":      {KeyCode: 191, ShiftKey: "?", Key: "/"},
		"ShiftRight": {KeyCode: 161, KeyCodeWithoutLocation: 16, Key: "Shift", Location: 2},

		// Last row
		"ControlLeft":  {KeyCode: 162, KeyCodeWithoutLocation: 17, Key: "Control", Location: 1},
		"MetaLeft":     {KeyCode: 91, Key: "Meta", Location: 1},
		"AltLeft":      {KeyCode: 164, KeyCodeWithoutLocation: 18, Key: "Alt", Location: 1},
		"Space":        {KeyCode: 32, Key: " "},
		"AltRight":     {KeyCode: 165, KeyCodeWithoutLocation: 18, Key: "Alt", Location: 2},
		"AltGraph":     {KeyCode: 225, Key: "AltGraph"},
		"MetaRight":    {KeyCode: 92, Key: "Meta", Location: 2},
		"ConTextMenu":  {KeyCode: 93, Key: "ConTextMenu"},
		"ControlRight": {KeyCode: 163, KeyCodeWithoutLocation: 17, Key: "Control", Location: 2},

		// Center block
		"PrintScreen": {KeyCode: 44, Key: "PrintScreen"},
		"ScrollLock":  {KeyCode: 145, Key: "ScrollLock"},
		"Pause":       {KeyCode: 19, Key: "Pause"},

		"PageUp":   {KeyCode: 33, Key: "PageUp"},
		"PageDown": {KeyCode: 34, Key: "PageDown"},
		"Insert":   {KeyCode: 45, Key: "Insert"},
		"Delete":   {KeyCode: 46, Key: "Delete"},
		"Home":     {KeyCode: 36, Key: "Home"},
		"End":      {KeyCode: 35, Key: "End"},

		"ArrowLeft":  {KeyCode: 37, Key: "ArrowLeft"},
		"ArrowUp":    {KeyCode: 38, Key: "ArrowUp"},
		"ArrowRight": {KeyCode: 39, Key: "ArrowRight"},
		"ArrowDown":  {KeyCode: 40, Key: "ArrowDown"},

		// Numpad
		"NumLock":        {KeyCode: 144, Key: "NumLock"},
		"NumpadDivide":   {KeyCode: 111, Key: "/", Location: 3},
		"NumpadMultiply": {KeyCode: 106, Key: "*", Location: 3},
		"NumpadSubtract": {KeyCode: 109, Key: "-", Location: 3},
		"Numpad7":        {KeyCode: 36, ShiftKeyCode: 103, Key: "Home", ShiftKey: "7", Location: 3},
		"Numpad8":        {KeyCode: 38, ShiftKeyCode: 104, Key: "ArrowUp", ShiftKey: "8", Location: 3},
		"Numpad9":        {KeyCode: 33, ShiftKeyCode: 105, Key: "PageUp", ShiftKey: "9", Location: 3},
		"Numpad4":        {KeyCode: 37, ShiftKeyCode: 100, Key: "ArrowLeft", ShiftKey: "4", Location: 3},
		"Numpad5":        {KeyCode: 12, ShiftKeyCode: 101, Key: "Clear", ShiftKey: "5", Location: 3},
		"Numpad6":        {KeyCode: 39, ShiftKeyCode: 102, Key: "ArrowRight", ShiftKey: "6", Location: 3},
		"NumpadAdd":      {KeyCode: 107, Key: "+", Location: 3},
		"Numpad1":        {KeyCode: 35, ShiftKeyCode: 97, Key: "End", ShiftKey: "1", Location: 3},
		"Numpad2":        {KeyCode: 40, ShiftKeyCode: 98, Key: "ArrowDown", ShiftKey: "2", Location: 3},
		"Numpad3":        {KeyCode: 34, ShiftKeyCode: 99, Key: "PageDown", ShiftKey: "3", Location: 3},
		"Numpad0":        {KeyCode: 45, ShiftKeyCode: 96, Key: "Insert", ShiftKey: "0", Location: 3},
		"NumpadDecimal":  {KeyCode: 46, ShiftKeyCode: 110, Key: "\u0000", ShiftKey: ".", Location: 3},
		"NumpadEnter":    {KeyCode: 13, Key: "Enter", Text: "\r", Location: 3},
	}

	register("us", validKeys, Keys)
}

//
// Copyright (c) 2016 Konstantin Ivanov <kostyarin.ivanov@gmail.com>.
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Do What
// The Fuck You Want To Public License, Version 2, as published by
// Sam Hocevar. See LICENSE file for more details or see below.
//

//
//        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//                    Version 2, December 2004
//
// Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>
//
// Everyone is permitted to copy and distribute verbatim or modified
// copies of this license document, and changing it is allowed as long
// as the name is changed.
//
//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
//  0. You just DO WHAT THE FUCK YOU WANT TO.
//

package aurora

import (
	"fmt"
	"testing"
)

func TestValue_String(t *testing.T) {
	var v Value
	// colorized
	v = value{"x", 0, 0}
	if x := v.String(); x != "x" {
		t.Errorf("(value).String: want %q, got %q", "x", x)
	}
	v = value{"x", BlackFg, RedBg}
	want := "\033[30mx\033[0m\033[41m"
	if got := v.String(); want != got {
		t.Errorf("(value).String: want %q, got %q", want, got)
	}
	// clear
	v = valueClear{"x"}
	if x := v.String(); x != "x" {
		t.Errorf("(value).String: want %q, got %q", "x", x)
	}

}

func TestValue_Color(t *testing.T) {
	// colorized
	if (value{"", RedFg, 0}).Color() != RedFg {
		t.Error("wrong color")
	}
	// clear
	if (valueClear{0}).Color() != 0 {
		t.Error("wrong color")
	}
}

func TestValue_Value(t *testing.T) {
	// colorized
	if (value{"x", RedFg, BlueBg}).Value() != "x" {
		t.Error("wrong value")
	}
	// clear
	if (valueClear{"x"}).Value() != "x" {
		t.Error("wrong value")
	}
}

func TestValue_Bleach(t *testing.T) {
	// colorized
	if (value{"x", RedFg, BlueBg}).Bleach() != (value{value: "x"}) {
		t.Error("wrong bleached")
	}
	// clear
	if (valueClear{"x"}).Bleach() != (valueClear{"x"}) {
		t.Error("wrong bleached")
	}
}

func TestValue_Format(t *testing.T) {
	var v Value
	var want, got string
	//
	// colorized
	//
	v = value{3.14, RedFg, BlueBg}
	got = fmt.Sprintf("%+1.3g", v)
	want = "\033[31m" + fmt.Sprintf("%+1.3g", 3.14) +
		"\033[0m\033[44m"
	if want != got {
		t.Errorf("Format: want %q, got %q", want, got)
	}
	//
	var utf8Verb = "%+1.3世" // verb that fit more then 1 byte
	got = fmt.Sprintf(utf8Verb, v)
	want = "\033[31m" + fmt.Sprintf(utf8Verb, 3.14) +
		"\033[0m\033[44m"
	if want != got {
		t.Errorf("Format: want %q, got %q", want, got)
	}
	//
	// clear
	//
	v = valueClear{3.14}
	got = fmt.Sprintf("%+1.3g", v)
	want = fmt.Sprintf("%+1.3g", 3.14)
	if want != got {
		t.Errorf("Format: want %q, got %q", want, got)
	}
	//
	got = fmt.Sprintf(utf8Verb, v)
	want = fmt.Sprintf(utf8Verb, 3.14)
	if want != got {
		t.Errorf("Format: want %q, got %q", want, got)
	}
}

func Test_tail(t *testing.T) {
	// colorized
	if (value{"x", 0, BlueBg}).tail() != BlueBg {
		t.Error("wrong tail color")
	}
	// clear
	if (valueClear{"x"}).tail() != 0 {
		t.Error("wrong tail color")
	}
}

func Test_setTail(t *testing.T) {
	// colorized
	if (value{"x", 0, 0}).setTail(RedFg) != (value{"x", 0, RedFg}) {
		t.Error("wrong setTail behavior")
	}
	// clear
	if (valueClear{"x"}).setTail(RedFg) != (valueClear{"x"}) {
		t.Error("wrong setTail behavior")
	}
}

func TestValue_colors(t *testing.T) {
	test := func(name string, v Value, clr Color) {
		if c := v.Color(); c != clr {
			t.Errorf("wrong color for %s: want %d, got %d", name, clr, c)
		}
	}
	// colorized
	test("Black", Red("x").Black(), BlackFg)
	test("Red", Bold("x").Red(), RedFg|BoldFm)
	test("Green", Inverse("x").Green(), GreenFg|InverseFm)
	test("Inverse&Brown", Bold("x").Inverse().Brown(), BoldFm|InverseFm|BrownFg)
	test("Blue", BgBlue("x").Blue(), BlueFg|BlueBg)
	test("Magenta", Magenta("x").Magenta(), MagentaFg)
	test("Cyan", Red("x").Cyan(), CyanFg)
	test("LightGray", Green("x").LightGray(), GrayFg)
	test("Gray", Red("x").Gray(), BlackFg|FgBrightFm)
	test("BrightRed", LightRed("x"), RedFg|FgBrightFm)
	test("LightRed", Bold("x").LightRed(), RedFg|BoldFm|FgBrightFm)
	test("LightGreen", Inverse("x").LightGreen(), GreenFg|InverseFm|FgBrightFm)
	test("Inverse&Yellow", Bold("x").Inverse().Yellow(), BoldFm|InverseFm|BrownFg|FgBrightFm)
	test("LightBlue", BgBlue("x").LightBlue(), BlueFg|BlueBg|FgBrightFm)
	test("LightMagenta", Magenta("x").LightMagenta(), MagentaFg|FgBrightFm)
	test("LightCyan", Red("x").LightCyan(), CyanFg|FgBrightFm)
	test("White", Green("x").White(), GrayFg|FgBrightFm)
	test("BgBlack", Black("x").BgBlack(), BlackFg|BlackBg)
	test("BgRed", Red("x").BgRed(), RedFg|RedBg)
	test("BgGreen", Green("x").BgGreen(), GreenFg|GreenBg)
	test("BgBrown", Brown("x").BgBrown(), BrownFg|BrownBg)
	test("BgBlue", Blue("x").BgBlue(), BlueFg|BlueBg)
	test("BgMagenta", BgCyan("x").BgMagenta(), MagentaBg)
	test("BgCyan", Cyan("x").BgCyan(), CyanFg|CyanBg)
	test("BgLightGray", LightGray("x").BgLightGray(), GrayFg|GrayBg)
	test("BgGray", Black("x").BgGray(), BlackFg|BlackBg|BgBrightFm)
	test("BgLightRed", LightRed("x").BgLightRed(), RedFg|RedBg|BgBrightFm|FgBrightFm)
	test("BgLightGreen", LightGreen("x").BgLightGreen(), GreenFg|GreenBg|BgBrightFm|FgBrightFm)
	test("BgYellow", Yellow("x").BgYellow(), BrownFg|BrownBg|BgBrightFm|FgBrightFm)
	test("BgLightBlue", LightBlue("x").BgLightBlue(), BlueFg|BlueBg|BgBrightFm|FgBrightFm)
	test("BgLightMagenta", BgLightCyan("x").BgLightMagenta(), MagentaBg|BgBrightFm)
	test("BgLightCyan", LightCyan("x").BgLightCyan(), CyanFg|CyanBg|BgBrightFm|FgBrightFm)
	test("BgWhite", White("x").BgWhite(), GrayFg|GrayBg|BgBrightFm|FgBrightFm)
	test("Bold & BlueBg", Red("x").BgBlue().Bold(), RedFg|BoldFm|BlueBg)
	test("Inverse", Black("x").Inverse(), BlackFg|InverseFm)
	// clear
	test("Black", valueClear{"x"}.Black(), 0)
	test("Red", valueClear{"x"}.Red(), 0)
	test("Green", valueClear{"x"}.Green(), 0)
	test("Inverse&Brown", valueClear{"x"}.Inverse().Brown(), 0)
	test("Blue", valueClear{"x"}.Blue(), 0)
	test("Magenta", valueClear{"x"}.Magenta(), 0)
	test("Cyan", valueClear{"x"}.Cyan(), 0)
	test("LightGray", valueClear{"x"}.LightGray(), 0)
	test("Gray", valueClear{"x"}.Gray(), 0)
	test("LightRed", valueClear{"x"}.LightRed(), 0)
	test("LightGreen", valueClear{"x"}.LightGreen(), 0)
	test("Inverse&Yellow", valueClear{"x"}.Inverse().Yellow(), 0)
	test("LightBlue", valueClear{"x"}.LightBlue(), 0)
	test("LightMagenta", valueClear{"x"}.LightMagenta(), 0)
	test("LightCyan", valueClear{"x"}.LightCyan(), 0)
	test("White", valueClear{"x"}.White(), 0)
	test("BgBlack", valueClear{"x"}.BgBlack(), 0)
	test("BgRed", valueClear{"x"}.BgRed(), 0)
	test("BgGreen", valueClear{"x"}.BgGreen(), 0)
	test("BgBrown", valueClear{"x"}.BgBrown(), 0)
	test("BgBlue", valueClear{"x"}.BgBlue(), 0)
	test("BgMagenta", valueClear{"x"}.BgMagenta(), 0)
	test("BgCyan", valueClear{"x"}.BgCyan(), 0)
	test("BgLightGray", valueClear{"x"}.BgLightGray(), 0)
	test("BgGray", valueClear{"x"}.BgGray(), 0)
	test("BgLightRed", valueClear{"x"}.BgLightRed(), 0)
	test("BgLightGreen", valueClear{"x"}.BgLightGreen(), 0)
	test("BgYellow", valueClear{"x"}.BgYellow(), 0)
	test("BgLightBlue", valueClear{"x"}.BgLightBlue(), 0)
	test("BgLightMagenta", valueClear{"x"}.BgLightMagenta(), 0)
	test("BgLightCyan", valueClear{"x"}.BgLightCyan(), 0)
	test("BgWhite", valueClear{"x"}.BgWhite(), 0)
	test("Bold & BlueBg", valueClear{"x"}.BgBlue().Bold(), 0)
	test("Inverse", valueClear{"x"}.Inverse(), 0)
}

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
	"testing"
)

func TestColor_Nos(t *testing.T) {
	c := Color(0)
	if c.Nos() != "" {
		t.Error("some Nos for 0 color")
	}
	c = BoldFm | InverseFm | RedFg | MagentaBg
	want := "1;7;31;45"
	if nos := c.Nos(); nos != want {
		t.Errorf("wrong Nos: want %q, got %q", want, nos)
	}
	c = InverseFm | BlackBg
	want = "7;40"
	if nos := c.Nos(); nos != want {
		t.Errorf("wrong Nos: want %q, got %q", want, nos)
	}
}

func TestColor_Nos_Bright(t *testing.T) {
	c := Color(0)
	if c.Nos() != "" {
		t.Error("some Nos for 0 color")
	}
	c = BoldFm | InverseFm | RedFg | MagentaBg | FgBrightFm
	want := "1;7;91;45"
	if nos := c.Nos(); nos != want {
		t.Errorf("wrong Nos: want %q, got %q", want, nos)
	}
	c = InverseFm | BlackBg | BgBrightFm
	want = "7;100"
	if nos := c.Nos(); nos != want {
		t.Errorf("wrong Nos: want %q, got %q", want, nos)
	}
	c = InverseFm | CyanBg | RedFg | FgBrightFm | BgBrightFm
	want = "7;91;106"
	if nos := c.Nos(); nos != want {
		t.Errorf("wrong Nos: want %q, got %q", want, nos)
	}
}

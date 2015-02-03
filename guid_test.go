package guid

import "testing"

func TestNew(t *testing.T) {
	for i := 0; i < 1024; i++ {
		g := New()
		if !g.IsConformant() {
			t.Errorf("Guid '%v' is not RFC4122 compliant.\n", g)
		}
	}
}

func BenchmarkNewGuid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}

func TestParseString(t *testing.T) {
	var goodGuids = [...]string{
		"0e545c9c-6942-4988-fab0-645274cfaded",
		"22e2c08b-e2bd-449a-8fc7-6ff9558ba733",
		"3D7670ff-48CC-42D3-91E4-B09177487D0C",
		"33C69DB0-3895-4D6F-D128-1855D3995742",
	}
	var goodGuidBytes = [...][16]byte{
		{14, 84, 92, 156, 105, 66, 73, 136, 250, 176, 100, 82, 116, 207, 173, 237},
		{34, 226, 192, 139, 226, 189, 68, 154, 143, 199, 111, 249, 85, 139, 167, 51},
		{61, 118, 112, 255, 72, 204, 66, 211, 145, 228, 176, 145, 119, 72, 125, 12},
		{51, 198, 157, 176, 56, 149, 77, 111, 209, 40, 24, 85, 211, 153, 87, 66},
	}
	for i, s := range goodGuids {
		if !IsGuid(s) {
			t.Errorf("good guid '%v' failed IsGuid test\n", s)
		}
		g, err := ParseString(s)
		if err != nil {
			t.Errorf("good guid '%v' failed to parse [%v]\n", s, err)
		}
		if [16]byte(*g) != goodGuidBytes[i] {
			t.Error("guid does not match bytes")
		}
	}
	var badGuids = [...]string{
		"0g545c9c-f942-4988-4ab0-645274cfaded",
		"2e2c08b-82bd-449a-7fc7-6ff9558ba733",
		"3D76-709898CC-42D3-41E4-B09177487D0C",
		"33C69DB0D8954D6F71281855D3995742",
	}
	for _, s := range badGuids {
		if IsGuid(s) {
			t.Error("bad guid passed IsGuid test")
		}
		if _, err := ParseString(s); err != ErrInvalid {
			t.Error("bad guid parsed")
		}
	}
}

func TestParseString2(t *testing.T) {
	for i := 0; i < 1024; i++ {
		g := New()
		g2, _ := ParseString(g.String())
		if *g != *g2 {
			t.Error("guid changed in conversion")
		}
	}
}

func BenchmarkParseString(b *testing.B) {
	var guids [16]string
	for i := 0; i < 16; i++ {
		guids[i] = New().String()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParseString(guids[i%16])
	}
}

func BenchmarkIsGuid(b *testing.B) {
	var guids [16]string
	for i := 0; i < 16; i++ {
		guids[i] = New().String()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsGuid(guids[i%16])
	}
}

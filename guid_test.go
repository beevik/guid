package guid

import (
    "testing"
)

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
        "0e545c9c-f942-4988-4ab0-645274cfaded",
        "22e2c08b-82bd-449a-7fc7-6ff9558ba733",
        "3D7670ff-98CC-42D3-41E4-B09177487D0C",
        "33C69DB0-D895-4D6F-7128-1855D3995742",
    }
    for _, s := range goodGuids {
        if !IsGuid(s) {
            t.Errorf("good guid '%v' failed IsGuid test\n", s)
        }
        if _, err := ParseString(s); err != nil {
            t.Errorf("good guid '%v' failed to parse [%v]\n", s, err)
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

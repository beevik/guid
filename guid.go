// Package guid defines a type for globally unique identifiers.  It
// provides functions to generate RFC4122-compliant guids, to parse
// strings into guids, and to convert guids to strings.
package guid

import (
    "crypto/rand"
    "encoding/hex"
    "errors"
    "fmt"
)

// Guid is a globally unique 16 byte identifier
type Guid [16]byte

var (
    ErrInvalid error // Parsed value is not a valid Guid
)

func init() {
    ErrInvalid = errors.New("guid: invalid format")
}

// New generates a random RFC4122-conformant version 4 Guid.
func New() *Guid {
    g := new(Guid)
    if _, err := rand.Read(g[:]); err != nil {
        panic(err)
    }
    g[6] = (g[6] & 0x0f) | 0x40 // version = 4
    g[8] = (g[8] & 0x3f) | 0x80 // variant = RFC4122
    return g
}

// ParseString returns the Guid represented by the string s.
func ParseString(s string) (*Guid, error) {
    if !IsGuid(s) {
        return nil, ErrInvalid
    }
    g := new(Guid)
    b, _ := hex.DecodeString(s[0:8] + s[9:13] + s[14:18] + s[19:23] + s[24:36])
    copy(g[:], b)
    return g, nil
}

// IsGuid returns true if the string contains a properly formatted Guid.
func IsGuid(s string) bool {
    if len(s) != 36 {
        return false
    }
    if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
        return false
    }
    stripped := [...]string{s[0:8], s[9:13], s[14:18], s[19:23], s[24:36]}
    for _, sub := range stripped {
        for _, c := range sub {
            if (c < '0' || c > '9') && (c < 'a' || c > 'f') && (c < 'A' || c > 'F') {
                return false
            }
        }
    }
    return true
}

// String returns a standard hexadecimal string version of the Guid.
// Lowercase characters are used.
func (g *Guid) String() string {
    return fmt.Sprintf("%x-%x-%x-%x-%x", g[0:4], g[4:6], g[6:8], g[8:10], g[10:16])
}

// StringUpper returns a standard hexadecimal string version of the Guid.
// Uppercase characters are used.
func (g *Guid) StringUpper() string {
    return fmt.Sprintf("%X-%X-%X-%X-%X", g[0:4], g[4:6], g[6:8], g[8:10], g[10:16])
}

// IsConformant determines if the GUID is RFC-4122 conformant.  If the
// variant is "reserved for future definition" or the version is unknown,
// then it is non-conformant.
func (g *Guid) IsConformant() bool {
    version := (g[6] & 0xf0) >> 4
    if version < 1 || version > 5 {
        return false
    }
    return (g[8] & 0xe0) != 0xe0
}

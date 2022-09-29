package ids

import (
	"github.com/aofei/sandid"
)

// SandID returns a sandid.SandID value.
func SandID() sandid.SandID {
	return sandid.New()
}

// SandIDString returns the string representation of the return value of
// SandID
func SandIDString() string {
	return SandID().String()
}

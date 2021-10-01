package utils

import (
	"fmt"

	"github.com/galaxy126/icicb-base/hash"
	"github.com/galaxy126/icicb-base/inter/idx"
)

// NameOf returns human readable string representation.
func NameOf(p idx.ValidatorID) string {
	if name := hash.GetNodeName(p); len(name) > 0 {
		return name
	}

	return fmt.Sprintf("%d", p)
}

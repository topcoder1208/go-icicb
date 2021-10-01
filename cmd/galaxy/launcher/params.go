package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://701414b5b824984b166c8219759e9cca08c54eef855096a6acdda14f1d6ee890b39d1abd3202a5be754ea35c99913558aa9fb8c6126fef633de52716ab3d2279@192.168.115.160:5060",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}

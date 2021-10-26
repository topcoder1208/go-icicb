package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://f8dcdfa611e4116758f46290169ae908ae840a769571e20c8178bda7b35709c363691c2d56429072ac12468b9d0c963e68575a1fe5f1129548b016eb3fca00c7@18.117.255.252:5060",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}

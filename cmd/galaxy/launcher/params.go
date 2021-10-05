package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://335872da24ac5775d0e68fae5519463794f01f44d9f2234c2b2f57c2133900a196c76bd87ef5ad11394e0a66bf9c0b7159412b5333de0b6be425df0303504457@192.168.115.160:5060",
		"enode://374352c78999485c435d3858279658e3d34cf82afced8706bbda3a1cbcac4e18a59010bd8ec63d86d7402f5d8a597738b7f88b2821adf0e7ad051dbc832bf533@192.168.115.163:5060",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}

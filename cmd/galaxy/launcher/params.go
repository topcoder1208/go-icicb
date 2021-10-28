package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		// testnet bootnodes
		"enode://6033f037e7a60340ba47b46ecdba07a5a87c1c41ed1d9422f871f33ed6a88477016c0caa1571905e650bccc5c03565103c951d0d0aaf715afbc2eb31c8861d09@18.117.255.252:5060",
		// mainnet bootnodes
		// "enode://6033f037e7a60340ba47b46ecdba07a5a87c1c41ed1d9422f871f33ed6a88477016c0caa1571905e650bccc5c03565103c951d0d0aaf715afbc2eb31c8861d09@18.117.255.252:5060",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}

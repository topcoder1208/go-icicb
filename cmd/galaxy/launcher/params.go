package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		// testnet bootnodes
		"enode://c68a3c5efaf9a570507ac66f9e02b34418d3591b6ede2fe09e39b317e1cf22e7af87c974c047aee5c2be5e815c66180e1589cf832a837c0f7336fc716a87b7b5@192.168.115.160:5060",
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

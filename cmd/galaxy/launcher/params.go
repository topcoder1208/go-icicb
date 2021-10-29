package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		// testnet bootnodes
		"enode://178dc1fec4eb1ab4664960d3f7fa4ad623c98aef800a9fd50b6573709639b2a1b08575c88e1ffec21303b8d270161443d17223d136d04674146db66f25412659@192.168.115.160:5060",
		"enode://381bfe4a9240e494aa5ce93b04ed1aeec3fedde8d8bfc079b62ffacc012002ceb0b3785798f65f242fd0e7e4c936b5094ed183e2ce5bd42ba1cf7be22c234900@18.117.255.252:5060",
		"enode://c4665dfb0b6614733302aa4050620cad6e49bf40ac545c36b770e129d42e001ed36f270477e1c05cff2418049664c085e6d11b3a2f2f22856fe4a0c5016b84b7@13.59.118.124:5060",
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

package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		// testnet bootnodes
		"enode://178dc1fec4eb1ab4664960d3f7fa4ad623c98aef800a9fd50b6573709639b2a1b08575c88e1ffec21303b8d270161443d17223d136d04674146db66f25412659@192.168.115.160:5060",
		"enode://1494e08b969f46400da2d7445d5a829266001ae43862e11b34cb28c939e0de3240167d511ac1ac5403a3c20751b87deca4bcd34fbf8b1edfe7ef851703928180@18.117.255.252:5060",
		"enode://0c79fcb218d3e1014f99e48c6dab99a56ec2ad73bd31034bce303561effacadedf4c667749329807722940ed7b62581fcc66456762a7af2e72f27ba643181b2a@13.59.118.124:5060",
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

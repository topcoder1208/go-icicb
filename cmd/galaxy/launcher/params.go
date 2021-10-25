package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://2f0b9f8e655ec1b9ac126d012ec2f0ab4973d1b51b5e404cf865a2bb24591db74553841e6df86388b226a01822fe37229f3895565faf7dbd5842542bdc5e4ca6@192.168.115.160:5060",
		"enode://06a194fef6f143c14a64929db816db85000a49d93a682be43e43e9ca56fb33bf43a4b3d61a6046a71ff67c15d7222fdffa18330fb1affef1bb9e5c8390d6f79f@18.117.227.242:5060",
		"enode://6eb878a6c6ec4c5606c0e92d95dacc41cf504fbf9add6c1e7be54b84bf4ab1bb84fc23f046cffe9421facf92a04ec8dcd50078e3a79b6fd6d050060bba5b22f0@3.17.193.52:5060",

		/* "enode://8a90e8391075e6776e2647ad912ad6d7fa5204032870c3a97db9461ab75ea910ae5213ec49d6089ac47c22965062babb932fae2df6917d43894481b2b5fad906@192.168.115.160:5060",
		"enode://c7334eca89572566f2b9108afcdacb684a2427f5742a27f7d9fda76e91d2563fca617411acfa6b1693e9ddbc8534f891a68f925f1ee0453834cbe3c06e40eb24@192.168.115.163:5060", */
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}

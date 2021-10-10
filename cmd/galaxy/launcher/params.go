package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://043fb0771d85c82cc4f71130094508e250f84df64368ad427678a51acb36d50450458a92f395f038021e04806d74ec968cae3fb59535ed554cfb266aa1faf90a@13.59.118.124:5060",
		"enode://06a194fef6f143c14a64929db816db85000a49d93a682be43e43e9ca56fb33bf43a4b3d61a6046a71ff67c15d7222fdffa18330fb1affef1bb9e5c8390d6f79f@18.117.227.242:5060",
		"enode://6eb878a6c6ec4c5606c0e92d95dacc41cf504fbf9add6c1e7be54b84bf4ab1bb84fc23f046cffe9421facf92a04ec8dcd50078e3a79b6fd6d050060bba5b22f0@3.17.193.52:5060",

		"enode://5ef32ab2092024dbaae18edbe1528588c845684823d52c6d19ffc065247c57050d30d67238706af66ef3ab508f53aff74a4fda4a4016230a54dd1001b7985daf@192.168.115.160:5060",
		"enode://c7334eca89572566f2b9108afcdacb684a2427f5742a27f7d9fda76e91d2563fca617411acfa6b1693e9ddbc8534f891a68f925f1ee0453834cbe3c06e40eb24@192.168.115.163:5060",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}

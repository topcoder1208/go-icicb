package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://c6f16d1b65fe8a94fdf20037dea25597fdacc198cea61c16c2bbe5f421dbce74e674f057580f367cb18779c4e59e53189ddaeceba4da01c2ff00eb4b9e3782e9@13.59.118.124:5060",
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

package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		// local   bootnodes
		// "enode://178dc1fec4eb1ab4664960d3f7fa4ad623c98aef800a9fd50b6573709639b2a1b08575c88e1ffec21303b8d270161443d17223d136d04674146db66f25412659@192.168.115.160:5060",
		// testnet bootnodes
		"enode://f874cf80b187d7811a9655819af718a507863a3ad8a1d66a3cb0e313ed22f975aa6d0e40bd66dc499a773cf2e3db1275f4291d3da0b0e09070ed94911c33a760@18.219.43.180:5060",
		"enode://9a12b4a5f183e47597d83d4b4822e2f8b9cbec04f8b79649dd44110674a6898c8014923b2255e0bc89d268f9976efad1de47ec044aea81018f7624c096051ca2@18.117.255.252:5060",
		"enode://c150cc3794081526e405553acecac52ffe1d7e99ca1687a4c17fbbb49a8e70bc66a1b92b2d87644de831f8b91bc551c30199be7eb862220a1e577524f580b6b2@3.15.233.209:5060",
		"enode://71ad819f5d1bcc98712a7333595fd865b7aefb0cf56a2b9fd3f5622011628a1e817004c5dc4db26b9a88bb61df17e26ca91ce65d3ef266f8ec066115e65453ce@18.117.227.242:5060",
		"enode://05d0742456c4f9f5f68b411c59975bd553bca0da72993a73e3f2b9acc92b9388b25df1059724898cb4e236450f456ef25474251c6f156ecac511dcfdae9bb680@3.17.193.52:5060",
		"enode://81ba242ad6ed97affe26cf02d42274679415bde402f8dcb263cca11c1594903849a90e0753f7a086da26e50de177836604f35962cc8ae61e785cfa6f8bb06df7@13.59.118.124:5060",
		"enode://e6616a45894efa0bf2cd22f3b3aceaa0ef77c66f68a6d78b3bcf2b15300ecd5452184ed55b26b13fd491904f0feb448c7971c986a00b8751b046d05c81d9582e@3.144.206.44:5060",
		"enode://d7ca2adc1f0a8a4cefef53610d36058ce4b81fa8ff2f61412edd969e8c2b194717e49e4486bd5a219ab0ee9e69f1e0d6f1b3ff153d6d5efd3e519c513fd9654b@3.128.156.220:5060",

		// mainnet bootnodes
		// "enode://66f8bff83a520e1394b827d5f08a9731b95551d2a0b07b2065edff14b28b1dbd8286ea0c69855d110608b3034f5862a3eb0b9cbb092aab7c0e673b6c82a58b4f@3.144.185.169:5060",
		// "enode://653fe3aa421c9a7b687c4dbd3649ced0df45c88b1ef3538de87496f265215dce65edea3c5b93df16f04385af9ae5416014644695525d5ccb572c45632aebbc53@52.14.233.85:5060",
		// "enode://767c02c165dfec134ce14f6612dca7e805b13f34cb80a4fc5d1301b7c00296ae6fc4619065b4a5fdd76fa3c7f039fe22d397d4069f93ba3215cf1ea03f87741a@18.216.88.94:5060",
		// "enode://eae6034d342eda73aa470a7c6eecb3451d7256b95837b613453a26fae595f9d5d94b9015cb6e9709263f3bcef4a45e429a422ca86dafb39546707dbf20f1b966@18.220.249.133:5060",
		// "enode://1171cf1c04727f9d7fc5766547641445a75cf992ff3a14b34efb18c091c0326b56a60813c5e26d0c2cab5f21c24f981dfda5dd8ac281f2e3f777f462ae6361e5@18.219.78.202:5060",
		// "enode://0e9b30fa17804d1a5b205bd09a12d1437cdec4a1a187a6bf837a85fd293326bf8133bc3fd773e2126071343fd0d1c5c54b2c26901a68a6a23131a010a226cf67@18.117.169.16:5060",
		// "enode://52b75153c5d6b44198fa869b6f9cf6b98fab648a4bcf25bc47e85e9ea0ae9d85ba9fe579182e65501239ee719c8a0d296e8168c9c4982b0951d46d080b4fb87d@3.140.193.10:5060",
		// "enode://d4bb0d5e72af5cd41fed4f5487bb156a8156f933779c66c6b2017e4f5d35550101d57afc598a4b9f52d5006b720cd9a1d74bd31a518fe53c03161cf9767186f6@3.138.199.193:5060",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}

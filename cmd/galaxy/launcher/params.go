package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		// local   bootnodes
		// "enode://178dc1fec4eb1ab4664960d3f7fa4ad623c98aef800a9fd50b6573709639b2a1b08575c88e1ffec21303b8d270161443d17223d136d04674146db66f25412659@192.168.115.160:5060",
		// testnet bootnodes
		// "enode://94fd0818deb6e1a553c732db6b7ea083797075c9df01486c415348d97958ed269f4b849f01be2c6fb8358fbbde830de8727d420f29f87d59e0900f22a4afd60a@18.219.43.180:5060",
		// "enode://bbda5878e8404d29810b28919e6d104b5be0053e0de86c189aaa930f953bc94de7d6ee4ed02e8de2ee9637a4abcfac57db4a1b7b5659c1039700cc10873bf9e5@18.117.255.252:5060",
		// "enode://27147b57a6fd469fdaef737a11b5c19af3c2309540ce861f36690c56fcb5fcdd9627b22ef52fae3c4f000d09b82f6eda258448fea41e928582dc734cf51a4df5@3.15.233.209:5060",
		// "enode://b4ce26028e3f6334aaf8f7a1927453f6e39f7d9ec51e3d4face07d6a84d47603b1ad76d8945f5e8cc0ae6a98d03929c077ae3dd23b9618415e2203f4eb08d6e1@18.117.227.242:5060",
		// "enode://229aee78aff13cf064a7b812cce0c596285566aa93487ac638a5816211396ead92aceb4efb5ce61d79513241ce13b563608d499e429060d546cafb62aab796cc@3.17.193.52:5060",
		// "enode://7e601ddf52c276384979cf3271496942a71fe094abbc881e02f5a83254f2eb6361c457e3bd8fb078386d99cc522856703e9af73f1a85cca34c7cd1b00f93aa85@13.59.118.124:5060",
		// "enode://173b4137e5ac529e7f7755ffe38e415dd8dcfe910a4e99d5379e54784ae2f16bbe37764737e4cf4cac2294c78addca367bb404013ea74d4c8dd64d1710c4044b@3.144.206.44:5060",
		// "enode://8e92c4b4c6eb8c244bf392404fed1f74541211bdaf2c75b8c12c7275f8d26c5f7cf1c2a2a38c1ef8fd24bd27a73944d8972cdfc94bfa25a67832553fa7f51d87@3.128.156.220:5060",

		// mainnet bootnodes
		"enode://c004cbedf9bfdc8816ceaf32c25bc3c611788fcbf350ba4ee2225b404b84a2a75e6b3192558bb0545818fc31a2a0dd81db0d35735ccd4b1926b800831497b91a@3.144.185.169:5060",
		"enode://74522d11621b907211dd9a7ad04be4f19201547a00df11d892dd7831eb1c5f79aaea5bf413b81f150d4c52cd3a173db9fd77bce490a00bfc9f8e6c11e97be539@52.14.233.85:5060",
		"enode://aeef03857f3a76cf1f3ed7ee4e01fca73bef9fe5ce7e5ec8a5c213f94a1304caaa0bb2aa2c4a0ce241bd848959248551df14a69768d44ac8814423937013efb4@18.216.88.94:5060",
		"enode://f4a59e01d9791494e68bca72b25e25ba8dbc4c67d547a46d3b98597713e89fa7bfc3faf4c82e5b92f210f4a8691aa5fe5f73dfb0d4ee86c47cdf5dc7b2c309d0@18.220.249.133:5060",
		"enode://1ff3a72a132f5e0717218303f8df238239cf9fdf8f9e55c15aa80a7b5f220b17217cca2de5690d9908c5b2ec58608b8b4c6406fe62fbab3b7cb39d1a4b9fac31@18.219.78.202:5060",
		"enode://d434bd3b10e0ebf302e258eb7ebea3fe77a95d87938c40d23c49c10c5d12b018a63f7185ae45294495acb89e43345d84aa29895a2bece3f62043541544472a40@18.117.169.16:5060",
		"enode://a335272010f3168ed5fa5f7d56aa430caa71e5855339f9955f99af9d826143efab40b110db387883741611f8343ead2d481fa80bd9fcb480a0eb22b8ccb4ae8c@3.140.193.10:5060",
		"enode://b1fc560c1f241af5781aed2860ffeef0d6e325aeda39e3089664295cde5f8d7675926d8996996e047b761f6eb0053a588987c394f47cf73c02e0b014d9b0cdf6@3.138.199.193:5060",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}

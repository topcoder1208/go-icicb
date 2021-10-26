package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://cb5d36ea772ca02adee79086d592c336a9c6c86956f499687e763ddc374dc0bfde63ba93413a1f93bf486561f08cfef5fc82eac84ee8792e8a1ee3695e9fbd94@18.117.255.252:5060",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}

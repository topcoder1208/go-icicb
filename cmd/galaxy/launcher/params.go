package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://6870ccd7f2693cb422e0a4379b45756cc229aa26cfcf8a52d5c2846a87e955316725570f1fd5f8372e3f1aef7adec9341ea3e927eb9cbebee30cd6ac66287dbd@85.206.161.66:5060",
		"enode://c57a64fc9863a78d32099098093a119d292f4603cb62317af1df93f6ae84cb9436e8fac4d40ccb9351f99a5d35b2fe1ea9dc0b0946e865d4937ae080f464b4f6@85.206.161.65:5060",
		"enode://84cf45969b58f78dc5f630551ece1be21dcc4f8ac845080bebd003afb7cac07a612f3d13549ad0733ebb258478c0a9c9eb825f9194739b245c318a99555b561e@185.25.48.191:5060",
		"enode://841f2670d826454137f1d242962e3070f52578124b1694541d6b08ea103f9e03377e53da45dcd2d7b2ac9b3e69f0bc9847c4b9706c7cafbf543934cc62e3bab3@185.25.48.190:5060",
		"enode://c671de8656997a74ed2065506932e7379ed70b85f62be1e9b2eb753e422f626a8212ee4e98debf6887a35dd7983aba5a23ec692583f24976ed8b4ff615a23070@185.64.106.248:5060",
		"enode://4909baf7eb73417ad6aff115b881d46a142c055e2353c9878359750ad1e67300672242125bbf5d785bd969411597c1301c3522c2f15c53edc7573d0561f8cb3d@185.64.104.196:5060",
		"enode://2f08061eefa8459cec58e2004f82d2d26922369e36d794cbd6e20e622b9d2392e18739f4835e1cc81b114c3e3fdf92fe0836dc4b0bd3332f6635f44c3aa69970@185.64.104.193:5060",
		"enode://d40dbb9ed4872eeacab96160361d2d05fe8740503cc8641b88086ddab53ebf4e4348ed67c442c2dcd9adf9c681f7d50487de33cf893f1cad0076d13094590793@185.64.104.87:5060",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}

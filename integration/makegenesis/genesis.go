package makegenesis

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/galaxy-foundation/icicb-base/hash"
	"github.com/galaxy-foundation/icicb-base/inter/idx"

	galaxy "github.com/goicicb/galaxy"
	"github.com/goicicb/galaxy/genesis"
	"github.com/goicicb/galaxy/genesis/driver"
	"github.com/goicicb/galaxy/genesis/driverauth"
	"github.com/goicicb/galaxy/genesis/evmwriter"
	"github.com/goicicb/galaxy/genesis/gpos"
	"github.com/goicicb/galaxy/genesis/netinit"
	"github.com/goicicb/galaxy/genesis/sfc"
	"github.com/goicicb/galaxy/genesis/sicicb"
	"github.com/goicicb/galaxy/genesis/sti"
	"github.com/goicicb/galaxy/genesis/tokenizer"
	"github.com/goicicb/galaxy/genesisstore"
	"github.com/goicicb/inter"
	"github.com/goicicb/inter/validatorpk"
	futils "github.com/goicicb/utils"
)

var (
	FakeGenesisTime = inter.Timestamp(1608600000 * time.Second)
)

// FakeKey gets n-th fake private key.
func FakeKey(n int) *ecdsa.PrivateKey {
	reader := rand.New(rand.NewSource(int64(n)))

	key, err := ecdsa.GenerateKey(crypto.S256(), reader)

	fmt.Printf("\nYour new privatekey was generated %x\n", key.D)

	if err != nil {
		panic(err)
	}

	return key
}

type ValidatorAccount struct {
	address   string
	validator string
}

func MakeGenesisStore() *genesisstore.Store {
	genStore := genesisstore.NewMemStore()
	genStore.SetRules(galaxy.MainNetRules())

	var validatorAccounts = []ValidatorAccount{
		// for mainnet
		{
			address:   "0x5BF4C398ff75a84EbFaDB758c0bEc886685690Fb",
			validator: "0xc004b7440ac3295c78574330dadf2c15c2bbea118c8c8531a1455311f808232a3de23e8a8bf5ffdad3fd0d9a9289fbb709a21eae614840a28d28c8739772e30d2c4d",
		},
		{
			address:   "0x9F3ed13529c2C00fE11270d586c32E390FD9618b",
			validator: "0xc0040e677c769a6056998e104ffcfeb2ed854479899c08857325f60d0ec8b88286aa873ef9061bb8468179b5076b13627f050cd8ab62589a6127a028d5a739e1548f",
		},
		{
			address:   "0x7980AC08c410b2B84Edef88dbc21b50cD9fF23E7",
			validator: "0xc004e2908148dde384116cc321279965406d990192e3b142007c21a1a88c06ed1f5f61393f6f6bb6bb493dfed9968185dd4a69d1b64b67b14c58c7460b1b6f821273",
		},
		{
			address:   "0x1341980EDE9D847359f685247A8eE6749601baec",
			validator: "0xc004fd458d2eaf4faf810854e4b974ca19e9aa5fec8be85833b8cc70f2367a84a4ec604aec9eea83482910d9d40e0211815c368098f46c8aa819f15ff010a13f4ce8",
		},
		{
			address:   "0xA45086Fb5edDa7e6474E090ab13CB21Ca2168262",
			validator: "0xc004f246443f0237ac6ab16774751afffeddeabed478b8cbfe95345cc4067fcb8ac720b9a53aabd0ad7cf271effbcad8f0398aa31745d88c0e58ea9073b3357daa72",
		},
		{
			address:   "0x219B271a1fB54636cDD20905bba3ff8C37405952",
			validator: "0xc0046f961174d34cae1887a623e2c78ebe1aa4ba1f71e29387564e1d2c7c2094377477b53782524d310faab9da0358bddfa3c2bcd6d6dbe54dee579250231e6927d4",
		},
		{
			address:   "0x99C34f67BBb426661e7892965D7934f6898Cc660",
			validator: "0xc004757e63ae2def97b07a4dc7d46e63e9f44c3ffdcee0751da5398a8f8321f5c9227c0dc5bce7c55421bf2ab095f7eb5237ab2931c24e26bd6b92222d3ec8fd096d",
		},
		{
			address:   "0x59B4A45D9C36e7721ACB1eD45900fCc5F510a94c",
			validator: "0xc004aa4cf3ed2f01e9c70048d9af65e449802ed44d760d54a19d118e4c172597245b1479303dcee264b667146deab9ddd1c80972c25787b2039285ebaa3a9d16401d",
		},
	}

	var initialAccounts = []string{
		"0x32De1D547067d6f9B9e04B11180b696Cca792Bdb",
		"0x2B8241F7d108b4a796dB260B501DF2221d951262",
		"0x26577ad836B5FF57510da6516aa5e8b670AEE0A0",
		"0x2685F70B97BE5ad586ad567531200A209fcD7890",
		"0xac655dc57FA548ed3233ABd3d8735917f29504e0",
		"0x3a85Ba5005D1732f05995557853B828791cCf2d3",
		"0xB8c151FB77f6593806483121e9fB8a8aD28c4dbE",
		"0x019Fac6e645C8b2bC7B9A4bdF8aEB71d878170c0",
		"0xCdB1ee111A23168ad38E5227013ca979C270dE64",
		"0x585871Dd4dF23643BB4A45ccD6Fa05F35a21d5C2",
	}
	num := len(validatorAccounts)

	_total := 5000
	_validator := 10
	_staker := 100
	_initial := (5000 - (_validator+_staker)*num) / 10

	totalSupply := futils.ToIcicb(uint64(_total) * 1e6)
	balance := futils.ToIcicb(uint64(_validator) * 1e6)
	stake := futils.ToIcicb(uint64(_staker) * 1e6)
	initialBalance := futils.ToIcicb(uint64(_initial) * 1e6)

	validators := make(gpos.Validators, 0, num)

	now := time.Now() // current local time
	// sec := now.Unix()      // number of seconds since January 1, 1970 UTC
	nsec := now.UnixNano()
	time := inter.Timestamp(nsec)
	for i := 1; i <= num; i++ {
		addr := common.HexToAddress(validatorAccounts[i-1].address)
		pubkeyraw := common.Hex2Bytes(validatorAccounts[i-1].validator)
		fmt.Printf("\n# addr %x pubkeyraw %s len %d\n", addr, hex.EncodeToString(pubkeyraw), len(pubkeyraw))
		validatorID := idx.ValidatorID(i)
		pubKey := validatorpk.PubKey{
			Raw:  pubkeyraw,
			Type: validatorpk.Types.Secp256k1,
		}

		validators = append(validators, gpos.Validator{
			ID:               validatorID,
			Address:          addr,
			PubKey:           pubKey,
			CreationTime:     time,
			CreationEpoch:    0,
			DeactivatedTime:  0,
			DeactivatedEpoch: 0,
			Status:           0,
		})
	}
	for _, val := range initialAccounts {
		genStore.SetEvmAccount(common.HexToAddress(val), genesis.Account{
			Code:    []byte{},
			Balance: initialBalance,
			Nonce:   0,
		})
		// totalSupply.Add(totalSupply, initialBalance)
	}
	for _, val := range validators {
		genStore.SetEvmAccount(val.Address, genesis.Account{
			Code:    []byte{},
			Balance: balance,
			Nonce:   0,
		})
		genStore.SetDelegation(val.Address, val.ID, genesis.Delegation{
			Stake:              stake,
			Rewards:            new(big.Int),
			LockedStake:        new(big.Int),
			LockupFromEpoch:    0,
			LockupEndTime:      0,
			LockupDuration:     0,
			EarlyUnlockPenalty: new(big.Int),
		})
		// totalSupply.Add(totalSupply, stake)
		// totalSupply.Add(totalSupply, balance)
	}

	var owner common.Address
	if num != 0 {
		owner = validators[0].Address
	}

	genStore.SetMetadata(genesisstore.Metadata{
		Validators:    validators,
		FirstEpoch:    2,
		Time:          time,
		PrevEpochTime: time - inter.Timestamp(time.Time().Hour()),
		ExtraData:     []byte("galaxy"),
		DriverOwner:   owner,
		TotalSupply:   totalSupply,
	})
	genStore.SetBlock(0, genesis.Block{
		Time:        time - inter.Timestamp(time.Time().Minute()),
		Atropos:     hash.Event{},
		Txs:         types.Transactions{},
		InternalTxs: types.Transactions{},
		Root:        hash.Hash{},
		Receipts:    []*types.ReceiptForStorage{},
	})
	// pre deploy NetworkInitializer
	genStore.SetEvmAccount(netinit.ContractAddress, genesis.Account{
		Code:    netinit.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy NodeDriver
	genStore.SetEvmAccount(driver.ContractAddress, genesis.Account{
		Code:    driver.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy NodeDriverAuth
	genStore.SetEvmAccount(driverauth.ContractAddress, genesis.Account{
		Code:    driverauth.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy SFC
	genStore.SetEvmAccount(sfc.ContractAddress, genesis.Account{
		Code:    sfc.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy STI
	genStore.SetEvmAccount(sti.ContractAddress, genesis.Account{
		Code:    sti.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy sICICB
	genStore.SetEvmAccount(sicicb.ContractAddress, genesis.Account{
		Code:    sicicb.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy Tokenizer
	genStore.SetEvmAccount(tokenizer.ContractAddress, genesis.Account{
		Code:    tokenizer.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// set non-zero code for pre-compiled contracts
	genStore.SetEvmAccount(evmwriter.ContractAddress, genesis.Account{
		Code:    []byte{0},
		Balance: new(big.Int),
		Nonce:   0,
	})

	return genStore
}
func MakeTestnetGenesisStore() *genesisstore.Store {
	genStore := genesisstore.NewMemStore()
	genStore.SetRules(galaxy.TestNetRules())
	var validatorAccounts = []ValidatorAccount{
		{
			address:   "0x034c37E7850A0DB0298664f723b357ff25FF31E2",
			validator: "0439b9f3f5a56c6aa8c79e01094d496d4e5b0b2116f6e26790177fb7639ffdf473ed428b71eec45e9789e3210cd46e663b9852d2f58ce7070bf1c928ace37d904a",
		},
		{
			address:   "0xA3F3571734840d9A01279D2696bDD20342eBF302",
			validator: "04be3ddfc6d48ad5d0ab793968f30e412cfaf0a1e1bdf3af63f542c4082191349ef8d2d13a1a1ce2b1512526f1ff53bbb8365d7f2e953b64fcc8cc93ce6ab60d9d",
		},
		{
			address:   "0x4d16A5DCA915C2f2CE039f4204548A94f520EF2e",
			validator: "0454c530e6781b0c7bb378199d903651745c68239c150c793751d4fbb2bf923eb7e3fd155b5235140e8534fdafd726f0dd213ab5bc07d51bcc598aa67bb260901a",
		},
		{
			address:   "0x7ECB5240FB7237bE35ddd5E6B08994A8FC43E52D",
			validator: "04847343604e986ba2b2fdd64c905503b85918fa206cac3df141e8bb61afee7c07e77196b4372d0b9ff9331b1eeff0a989ca22ab5c95cbfec4f77dc584a96f7278",
		},
		{
			address:   "0xdD7792225BD36410F9deFD98878890eb6c8135ad",
			validator: "04303bbac1433b0d46feae674db33fd2b62794803f53f0335d8743ae7f6005608bee87c7f73a25fe296fd184536bd83a79e60e6c35d53ca208d2aa00f0b18d36d2",
		},
		{
			address:   "0xaA492E71d793C99D824efe945aE2091eb6e41977",
			validator: "043ecdc02c855c322be643aee9b8f735e82bc664746b09304a9883e553cf64ba2e91c1f3ab39ae24ae27ed3d41c93d947b0e64cb78c06dd408532f99ef2d207895",
		},
		{
			address:   "0xa73B9365479fBB5008E5222078F639b6039c2Aed",
			validator: "0xc00486fb9204c56ce4fc2e2de29db5c7df9917ea22cdd39b21234dd35a191d3e9677e199142799001101d79b5b8c4a2966c072b3a9f7c06c55151cd2ad27a3d3cd8b",
		},
		{
			address:   "0x0716C6Bb0573e3FD902Eb4A7311863f8a1E411b9",
			validator: "0xc004be2adee5c7b3d15cdb8ae0a099a759117cc6d5bbe45018fe6e0b05d645ca43069051819a3173c7e87b8947d1d6d3ae85c9dbf725f6778ef417caf186bd8fcac9",
		},
	}

	var initialAccounts = []string{
		"0x9cD60D0D9e4404Be3a1C890cAF477A08903Aca2b",
		"0x8E1c7C2960B5298Bc2580619224E56023a27996B",
		"0xEe8E84116F1903c1F0d723E9d1a92D20613a50d2",
		"0x5f4632ceD4D32B02c9d2217B19888b8eC9749114",
		"0x294cD6A64d63e9cbd92358C74cE751c43DE9F3dC",
		"0x6752bDd135D92025611c01ab0f16b532a046E863",
		"0xFaca4DAe41dcDD2618FfD083cf03Ef4C05078B79",
		"0xd68ccE056fe53c6C349AdE5De472597B8D2b576c",
		"0x4Af5d38b634C36d29F28Fe948383fA3be9fccda2",
		"0x9509eb170B5007e5Ac607944F800b8A475cc9bC7",
	}

	num := len(validatorAccounts)

	_total := 5000
	_validator := 10
	_staker := 100
	_initial := (5000 - (_validator+_staker)*num) / 10

	totalSupply := futils.ToIcicb(uint64(_total) * 1e6)
	balance := futils.ToIcicb(uint64(_validator) * 1e6)
	stake := futils.ToIcicb(uint64(_staker) * 1e6)
	initialBalance := futils.ToIcicb(uint64(_initial) * 1e6)

	validators := make(gpos.Validators, 0, num)

	now := time.Now() // current local time
	// sec := now.Unix()      // number of seconds since January 1, 1970 UTC
	nsec := now.UnixNano()
	time := inter.Timestamp(nsec)
	for i := 1; i <= num; i++ {
		addr := common.HexToAddress(validatorAccounts[i-1].address)
		pubkeyraw := common.Hex2Bytes(validatorAccounts[i-1].validator)
		fmt.Printf("\n# addr %x pubkeyraw %s len %d\n", addr, hex.EncodeToString(pubkeyraw), len(pubkeyraw))
		validatorID := idx.ValidatorID(i)
		pubKey := validatorpk.PubKey{
			Raw:  pubkeyraw,
			Type: validatorpk.Types.Secp256k1,
		}

		validators = append(validators, gpos.Validator{
			ID:               validatorID,
			Address:          addr,
			PubKey:           pubKey,
			CreationTime:     time,
			CreationEpoch:    0,
			DeactivatedTime:  0,
			DeactivatedEpoch: 0,
			Status:           0,
		})
	}

	for _, val := range initialAccounts {
		genStore.SetEvmAccount(common.HexToAddress(val), genesis.Account{
			Code:    []byte{},
			Balance: initialBalance,
			Nonce:   0,
		})
	}

	for _, val := range validators {
		genStore.SetEvmAccount(val.Address, genesis.Account{
			Code:    []byte{},
			Balance: balance,
			Nonce:   0,
		})
		genStore.SetDelegation(val.Address, val.ID, genesis.Delegation{
			Stake:              stake,
			Rewards:            new(big.Int),
			LockedStake:        new(big.Int),
			LockupFromEpoch:    0,
			LockupEndTime:      0,
			LockupDuration:     0,
			EarlyUnlockPenalty: new(big.Int),
		})
		// totalSupply.Add(totalSupply, stake)
		// totalSupply.Add(totalSupply, balance)
	}

	var owner common.Address
	if num != 0 {
		owner = validators[0].Address
	}

	genStore.SetMetadata(genesisstore.Metadata{
		Validators:    validators,
		FirstEpoch:    2,
		Time:          time,
		PrevEpochTime: time - inter.Timestamp(time.Time().Hour()),
		ExtraData:     []byte("fake"),
		DriverOwner:   owner,
		TotalSupply:   totalSupply,
	})
	genStore.SetBlock(0, genesis.Block{
		Time:        time - inter.Timestamp(time.Time().Minute()),
		Atropos:     hash.Event{},
		Txs:         types.Transactions{},
		InternalTxs: types.Transactions{},
		Root:        hash.Hash{},
		Receipts:    []*types.ReceiptForStorage{},
	})
	// pre deploy NetworkInitializer
	genStore.SetEvmAccount(netinit.ContractAddress, genesis.Account{
		Code:    netinit.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy NodeDriver
	genStore.SetEvmAccount(driver.ContractAddress, genesis.Account{
		Code:    driver.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy NodeDriverAuth
	genStore.SetEvmAccount(driverauth.ContractAddress, genesis.Account{
		Code:    driverauth.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy SFC
	genStore.SetEvmAccount(sfc.ContractAddress, genesis.Account{
		Code:    sfc.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy STI
	genStore.SetEvmAccount(sti.ContractAddress, genesis.Account{
		Code:    sti.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy sICICB
	genStore.SetEvmAccount(sicicb.ContractAddress, genesis.Account{
		Code:    sicicb.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy Tokenizer
	genStore.SetEvmAccount(tokenizer.ContractAddress, genesis.Account{
		Code:    tokenizer.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// set non-zero code for pre-compiled contracts
	genStore.SetEvmAccount(evmwriter.ContractAddress, genesis.Account{
		Code:    []byte{0},
		Balance: new(big.Int),
		Nonce:   0,
	})

	return genStore
}
func FakeGenesisStore(num int, balance, stake *big.Int) *genesisstore.Store {
	genStore := genesisstore.NewMemStore()
	genStore.SetRules(galaxy.FakeNetRules())

	validators := GetFakeValidators(num)

	totalSupply := new(big.Int)
	for _, val := range validators {
		genStore.SetEvmAccount(val.Address, genesis.Account{
			Code:    []byte{},
			Balance: balance,
			Nonce:   0,
		})
		genStore.SetDelegation(val.Address, val.ID, genesis.Delegation{
			Stake:              stake,
			Rewards:            new(big.Int),
			LockedStake:        new(big.Int),
			LockupFromEpoch:    0,
			LockupEndTime:      0,
			LockupDuration:     0,
			EarlyUnlockPenalty: new(big.Int),
		})
		totalSupply.Add(totalSupply, balance)
	}

	var owner common.Address
	if num != 0 {
		owner = validators[0].Address
	}

	genStore.SetMetadata(genesisstore.Metadata{
		Validators:    validators,
		FirstEpoch:    2,
		Time:          FakeGenesisTime,
		PrevEpochTime: FakeGenesisTime - inter.Timestamp(time.Hour),
		ExtraData:     []byte("fake"),
		DriverOwner:   owner,
		TotalSupply:   totalSupply,
	})
	genStore.SetBlock(0, genesis.Block{
		Time:        FakeGenesisTime - inter.Timestamp(time.Minute),
		Atropos:     hash.Event{},
		Txs:         types.Transactions{},
		InternalTxs: types.Transactions{},
		Root:        hash.Hash{},
		Receipts:    []*types.ReceiptForStorage{},
	})
	// pre deploy NetworkInitializer
	genStore.SetEvmAccount(netinit.ContractAddress, genesis.Account{
		Code:    netinit.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy NodeDriver
	genStore.SetEvmAccount(driver.ContractAddress, genesis.Account{
		Code:    driver.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy NodeDriverAuth
	genStore.SetEvmAccount(driverauth.ContractAddress, genesis.Account{
		Code:    driverauth.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// pre deploy SFC
	genStore.SetEvmAccount(sfc.ContractAddress, genesis.Account{
		Code:    sfc.GetContractBin(),
		Balance: new(big.Int),
		Nonce:   0,
	})
	// set non-zero code for pre-compiled contracts
	genStore.SetEvmAccount(evmwriter.ContractAddress, genesis.Account{
		Code:    []byte{0},
		Balance: new(big.Int),
		Nonce:   0,
	})

	return genStore
}

func GetFakeValidators(num int) gpos.Validators {
	validators := make(gpos.Validators, 0, num)

	for i := 1; i <= num; i++ {
		key := FakeKey(i)
		addr := crypto.PubkeyToAddress(key.PublicKey)
		pubkeyraw := crypto.FromECDSAPub(&key.PublicKey)

		validatorID := idx.ValidatorID(i)
		validators = append(validators, gpos.Validator{
			ID:      validatorID,
			Address: addr,
			PubKey: validatorpk.PubKey{
				Raw:  pubkeyraw,
				Type: validatorpk.Types.Secp256k1,
			},
			CreationTime:     FakeGenesisTime,
			CreationEpoch:    0,
			DeactivatedTime:  0,
			DeactivatedEpoch: 0,
			Status:           0,
		})
	}

	return validators
}

type Genesis struct {
	Nonce      uint64         `json:"nonce"`
	Timestamp  uint64         `json:"timestamp"`
	ExtraData  []byte         `json:"extraData"`
	GasLimit   uint64         `json:"gasLimit"   gencodec:"required"`
	Difficulty *big.Int       `json:"difficulty" gencodec:"required"`
	Mixhash    common.Hash    `json:"mixHash"`
	Coinbase   common.Address `json:"coinbase"`
	Alloc      GenesisAlloc   `json:"alloc"      gencodec:"required"`

	// These fields are used for consensus tests. Please don't use them
	// in actual genesis blocks.
	Number     uint64      `json:"number"`
	GasUsed    uint64      `json:"gasUsed"`
	ParentHash common.Hash `json:"parentHash"`
	BaseFee    *big.Int    `json:"baseFeePerGas"`
}

type GenesisAlloc map[common.Address]GenesisAccount

type GenesisAccount struct {
	Code       []byte                      `json:"code,omitempty"`
	Storage    map[common.Hash]common.Hash `json:"storage,omitempty"`
	Balance    *big.Int                    `json:"balance" gencodec:"required"`
	Nonce      uint64                      `json:"nonce,omitempty"`
	PrivateKey []byte                      `json:"secretKey,omitempty"` // for tests
}

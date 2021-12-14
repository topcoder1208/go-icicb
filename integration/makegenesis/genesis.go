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
			address:   "0x765CBa86f898De0210B5c110732E24de89D893a1",
			validator: "04ce49f7c6afdf7a457f31fc6e6216c8c0df3f46d0f7a80ccd80b35a4f91313184b345efd3724b4639736405a2c3bfb4b1aabea58a34fa2a58cdc3bbea9d80f234",
		},
		{
			address:   "0x8525e36dFdd0191326595e57C5f0e020c8CFf371",
			validator: "04653c606752353c3196193ca132466d137863c2966ac0479dbb989893e41bde818e7fcc6e1ec3d22f5262fb1c82814edf6dfb1368f34c6db400d7b44186449719",
		},
		{
			address:   "0x4a7A58a07cc1eaB994d0863673518D4121A9548f",
			validator: "04a4008e352eabeff85db3f1eea164f6970d6d0d7846061b061c666037e482f68e0b52f4df3bfdc309658f6b52665791e1da10b890f8f7715b218ee28ae5c09ed7",
		},
		{
			address:   "0x96d383E6E678B443C7B4bE832682d0A489f7FEb7",
			validator: "04939c39a101d5520339bff351a5cba585d85a72db013c187f8a329962c5ce0ffbc9745e680e9cbef2e448aec19809ad8e5fbdf318720d205e89897d4704849950",
		},
		{
			address:   "0xb32f9F1e98525F7cCaD6225162D2cDD8DBD971d9",
			validator: "0405fb835b662f599d4aeb1952a3e5ddd8002e6043fea0e0b1126adbd0ea8cebe5c6a5024fb8fc069c139b405ed261ea3f1a779c3a83b1041a37e1f9a8dfb163be",
		},
		{
			address:   "0xd7F6465f20530856864FDa690E95Cf29e0c9037A",
			validator: "0417b350522e3f8411e171ae3d4f7709ce0b10a02900e71581a7ac268c80472f97a6e6c87d5e2b1edbfa35357032d3627ccf3f90ee8f51ce49360fb1855e12f8ef",
		},
		{
			address:   "0xb61A862aD5D3385AB34B72ed77CcD6B830411dc6",
			validator: "04f704a31767b8f6d3606e166c35adf6ec725675196047b5012c0e05f849901d327bdaf44585a826ef7e986fbf060b8682341385b2db1bda7394edaca00752cd29",
		},
		{
			address:   "0x85a8b7bf26F26f5a828EB0ae5F497fa9c6baDB0a",
			validator: "04171eb754a53a173c500bf3c5b1e04a632924a04198b73f1881a19574c79e63e561b17f244dbbfcaec1e45a9310487d4c88bb78df187f17dfc1da98f619b184de",
		},
	}

	var initialAccounts = []string{
		"0xb8B5BE7122f317F86b47778422e277cD91C0B031",
	}
	num := len(validatorAccounts)

	_total := 5000
	_validator := 1
	_staker := 100
	_initial := (5000 - (_validator+_staker)*num) / len(initialAccounts)

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
	}
	/* var validatorAccounts = []ValidatorAccount{
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
			validator: "0486fb9204c56ce4fc2e2de29db5c7df9917ea22cdd39b21234dd35a191d3e9677e199142799001101d79b5b8c4a2966c072b3a9f7c06c55151cd2ad27a3d3cd8b",
		},
		{
			address:   "0x0716C6Bb0573e3FD902Eb4A7311863f8a1E411b9",
			validator: "04be2adee5c7b3d15cdb8ae0a099a759117cc6d5bbe45018fe6e0b05d645ca43069051819a3173c7e87b8947d1d6d3ae85c9dbf725f6778ef417caf186bd8fcac9",
		},
	}
	*/
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

# ICICB-Galaxy

EVM-compatible chain secured by the Lachesis consensus algorithm.

## Building the source

Building `galaxy` requires both a Go (version 1.14 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make galaxy
```
The build output is ```build/galaxy``` executable.

## Running `galaxy`

Going through all the possible command line flags is out of scope here,
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `galaxy` instance.

### Launching a network

Launching `galaxy` for a network:

```shell
$ galaxy --genesis /path/to/genesis.g
```

### Configuration

As an alternative to passing the numerous flags to the `galaxy` binary, you can also pass a
configuration file via:

```shell
$ galaxy --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to
export your existing configuration:

```shell
$ galaxy --your-favourite-flags dumpconfig
```

#### Validator

New validator private key may be created with `galaxy validator new` command.

To launch a validator, you have to use `--validator.id` and `--validator.pubkey` flags to enable events emitter.

```shell
$ galaxy --nousb --validator.id YOUR_ID --validator.pubkey 0xYOUR_PUBKEY
```



`galaxy` will prompt you for a password to decrypt your validator private key. Optionally, you can
specify password with a file using `--validator.password` flag.

#### Participation in discovery

Optionally you can specify your public IP to straighten connectivity of the network.
Ensure your TCP/UDP p2p port (5050 by default) isn't blocked by your firewall.

```shell
$ galaxy --nat extip:1.2.3.4
```

## Dev

### Running testnet

The network is specified only by its genesis file, so running a testnet node is equivalent to
using a testnet genesis file instead of a mainnet genesis file:
```shell
$ galaxy --genesis /path/to/testnet.g # launch node
```

It may be convenient to use a separate datadir for your testnet node to avoid collisions with other networks:
```shell
$ galaxy --genesis /path/to/testnet.g --datadir /path/to/datadir # launch node
$ galaxy --datadir /path/to/datadir account new # create new account
$ galaxy --datadir /path/to/datadir attach # attach to IPC
```

### Testing

Lachesis has extensive unit-testing. Use the Go tool to run tests:
```shell
go test ./...
```

If everything goes well, it should output something along these lines:
```
ok  	github.com/goicicb/app	0.033s
?   	github.com/goicicb/cmd/cmdtest	[no test files]
ok  	github.com/goicicb/cmd/galaxy	13.890s
?   	github.com/goicicb/cmd/galaxy/metrics	[no test files]
?   	github.com/goicicb/cmd/galaxy/tracing	[no test files]
?   	github.com/goicicb/crypto	[no test files]
?   	github.com/goicicb/debug	[no test files]
?   	github.com/goicicb/ethapi	[no test files]
?   	github.com/goicicb/eventcheck	[no test files]
?   	github.com/goicicb/eventcheck/basiccheck	[no test files]
?   	github.com/goicicb/eventcheck/gaspowercheck	[no test files]
?   	github.com/goicicb/eventcheck/heavycheck	[no test files]
?   	github.com/goicicb/eventcheck/parentscheck	[no test files]
ok  	github.com/goicicb/evmcore	6.322s
?   	github.com/goicicb/gossip	[no test files]
?   	github.com/goicicb/gossip/emitter	[no test files]
ok  	github.com/goicicb/gossip/filters	1.250s
?   	github.com/goicicb/gossip/gasprice	[no test files]
?   	github.com/goicicb/gossip/occuredtxs	[no test files]
?   	github.com/goicicb/gossip/piecefunc	[no test files]
ok  	github.com/goicicb/integration	21.640s
```

Also it is tested with [fuzzing](./FUZZING.md).


### Operating a private network (fakenet)

Fakenet is a private network optimized for your private testing.
It'll generate a genesis containing N validators with equal stakes.
To launch a validator in this network, all you need to do is specify a validator ID you're willing to launch.

Pay attention that validator's private keys are deterministically generated in this network, so you must use it only for private testing.

Maintaining your own private network is more involved as a lot of configurations taken for
granted in the official networks need to be manually set up.

To run the fakenet with just one validator (which will work practically as a PoA blockchain), use:
```shell
$ galaxy --fakenet 1/1
```

To run the fakenet with 5 validators, run the command for each validator:
```shell
$ galaxy --fakenet 1/5 # first node, use 2/5 for second node
```

If you have to launch a non-validator node in fakenet, use 0 as ID:
```shell
$ galaxy --fakenet 0/5
```

After that, you have to connect your nodes. Either connect them statically or specify a bootnode:
```shell
$ galaxy --fakenet 1/5 --bootnodes "enode://verylonghex@1.2.3.4:5050"
```

### Running the demo

For the testing purposes, the full demo may be launched using:
```shell
cd demo/
./start.sh # start the Galaxy processes
./stop.sh # stop the demo
./clean.sh # erase the chain data
```

### Make config file
```shell
./galaxy --genesis --rpc  -rpcaddr 0.0.0.0 --datadir=d:/icicb/galaxy --rpcapi "net,eth,txpool,web3" --maxpeers 128 --maxpendpeers 128 --txpool.globalqueue 4096  dumpconfig > d:/icicb/config.toml
```

### Running fakenet
```shell
./build/galaxy --fakenet 1/1 --http --http.addr="127.0.0.1" --http.port=5050 --http.corsdomain="*" --http.api="eth,debug,net,admin,web3,personal,txpool,icicb,dag" --datadir=d:/icicb/fake
```

### Running testnet
./build/opera --metrics  --cache 64000 --genesis
$HOME/fantom/mainnet.g --nousb --http --http.addr '0.0.0.0' --http.port 8545 --http.corsdomain "*" --http.vhosts "*" --ws --ws.addr '0.0.0.0' --ws.port 8546  --ws.origins '0.0.0.0' --graphql --graphql.corsdomain '*' --graphql.vhosts '*' --datadir "$HOME/fantom/node" --http.api "net,eth,web3" --ws.api "net,eth,web3"


### testing

generate genesis

```shell
./build/galaxy init d:/icicb/testgenesis.json --datadir d:/icicb/genesis.g
```

### generate validator key
```shell
./build/galaxy validator new
```

### run mainnet
```shell
./build/galaxy --genesis d:/icicb/genesis.g --datadir d:/icicb/mainnet --validator.id 0 --validator.pubkey 0xc0042ddbf4fda4bfefbb4368e3c0626faacfeb9319d08588cdb1a414eee7870ebc4c806aed6c5b01c2b2de5c60ba1cf98510dd4bcaecd56ad2dadc3c976279003797  --validator.password 123456789
```

```shell
./build/galaxy --genesis d:/icicb/genesis.g --datadir d:/icicb/mainnet --http --http.addr '0.0.0.0' --http.port 8545 --http.corsdomain "*" --http.vhosts "*" --http.api="eth,debug,net,admin,web3,personal,txpool,icicb,dag" --validator.id 1 --validator.pubkey 0xc004aa0713a0bc43226ab1f19e9547afc0d4b13fde317142436072d4d959621dc61a1d6ac185fb87ffe81d578c2a89a79e151a77ae71f31b8281e1c690c0f625dcf7 --validator.password D:\\icicb\\pass-1.txt

./build/galaxy --genesis d:/icicb/genesis.g --datadir d:/icicb/mainnet --nousb --validator.id 2 --validator.pubkey 0xc004398d3a0ee4514dfd111d22a4a54137b2ea79dc179f575f5b3dfe66ccc3348559a28af9ed7c075e3b93a5ff338dc657606104b3030ebcd95d7ad52fa62199553f --validator.password D:\\icicb\\pass-2.txt
```

`galaxy` will prompt you for a password to decrypt your validator private key. Optionally, you can
specify password with a file using `--validator.password` flag.

./build/galaxy --genesis d:/icicb/genesis.g --datadir d:/icicb/mainnet --bootnodes="enode://2cecf66045ee5f0defb2d0d88020a181504295882547b0442bd65246ab1a40e0164eb105558a22c4ce376bc01d25e3344521bc693915aaa5d19eb917c7acfc08@192.168.115.163:5060"
./build/galaxy init d:\\icicb\\genesis.json 

"init","d:\\icicb\\genesis.json"
"--genesis" 
"d:/icicb/genesis.g"
"--datadir"
"d:/icicb/mainnet"
"--http"
"--http.addr"
"'0.0.0.0'"
"--http.port"
"8545"
"--http.corsdomain"
"\"*\""
"--http.vhosts"
"\"*\""
"--http.api=\"eth,debug,net,admin,web3,personal,txpool,icicb,dag\""
"--nousb"
"--validator.id"
"1"
"--validator.pubkey"
"0xc004aa0713a0bc43226ab1f19e9547afc0d4b13fde317142436072d4d959621dc61a1d6ac185fb87ffe81d578c2a89a79e151a77ae71f31b8281e1c690c0f625dcf7"
"--validator.password"
"D:\\icicb\\pass-1.txt"

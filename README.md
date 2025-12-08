# How to use
ignite chain serve

The easiest way to build the blockchain is by running the command
`go build -o divined cmd/divined/main.go`

This will create the `divined` binary which is both the Node and the CLI to the Node

## Development Environment

Initialize the blockchain by running

`divined init divine --chain-id divine`

Set up test keypairs by running

`divined keys add alice --keyring-backend test`

`divined keys add bob --keyring-backend test`

Initialize them with tokens by running

`divined genesis add-genesis-account alice 100000000000000000div --keyring-backend test`

`divined genesis add-genesis-account bob 200000000000000000div --keyring-backend test`

Generate the Create Validator transaction by running

`divined genesis gentx alice 10000000000000000div --keyring-backend test`

Collect the transaction by running 
`divined genesis collect-gentxs`


Setup the no-fee requirements in the configuration by changing `min-gas-prices` in `~/.divined/config/app.toml` to `0div`

Change the staking token name to `div` in `~/.divined/config/genesis.json` by replacing all instances of `stake` with `div`

Start the node by running
`divined start`
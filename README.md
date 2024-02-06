# Procyon - A minimal Cosmos SDK chain

Procyon is the name of the star that represents maera (the hound) in the Virgo cluster. It is a copy of cosmos sdk `chain-minimal`, and is being used for researching module development.

`procyon` uses the **latest** version of the [Cosmos-SDK](https://github.com/cosmos/cosmos-sdk).

### Installation

Install and run `procyon`:

```sh
make install # install the procyon binary
make init # initialize the chain
procyon start # start the chain
```

# list test keys

```shell
procyon keys list --keyring-backend test

- address: mini16ajnus3hhpcsfqem55m5awf3mfwfvhpp36rc7d
  name: alice
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A0gUNtXpBqggTdnVICr04GHqIQOa3ZEpjAhn50889AQX"}'
  type: local
- address: mini1hv85y6h5rkqxgshcyzpn2zralmmcgnqwsjn3qg
  name: bob
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"ArXLlxUs2gEw8+clqPp6YoVNmy36PrJ7aYbV+W8GrcnQ"}'
  type: local
```

## Envoy module

FIXME: Requires (for now) a local version of `envoy` module which this copy of chain-minimal is being used to develop

Parallel to this project, checkout the envoy module

```shell
git clone git@github.com:christophercampbell/envoy.git
```

## create a lock

This will not be how locks are created, the system will create/configure them internally for named lockable (node exclusive) actions.

```shell
procyon tx envoy create lock1 mini16ajnus3hhpcsfqem55m5awf3mfwfvhpp36rc7d 666 12 --from alice --yes
```

## read a lock

```shell
procyon query envoy get-lock lock1
```
```
lock:
  at_block: 666
  envoy: mini16ajnus3hhpcsfqem55m5awf3mfwfvhpp36rc7d
  name: lock1
  num_blocks: 12
```

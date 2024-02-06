#!/bin/bash

rm -r ~/.procyon || true
PROCYON_BIN=$(which procyon)
# configure procyon
$PROCYON_BIN config set client chain-id demo
$PROCYON_BIN config set client keyring-backend test
$PROCYON_BIN keys add alice
$PROCYON_BIN keys add bob
$PROCYON_BIN init test --chain-id demo --default-denom mini
# update genesis
$PROCYON_BIN genesis add-genesis-account alice 10000000mini --keyring-backend test
$PROCYON_BIN genesis add-genesis-account bob 1000mini --keyring-backend test
# create default validator
$PROCYON_BIN genesis gentx alice 1000000mini --chain-id demo
$PROCYON_BIN genesis collect-gentxs

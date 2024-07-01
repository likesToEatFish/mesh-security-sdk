#!/bin/bash

set -ex

# initialize Hermes relayer configuration
mkdir -p /root/.hermes/
touch /root/.hermes/config.toml

# setup Hermes relayer configuration
tee /root/.hermes/config.toml <<EOF
[global]
log_level = 'info'
[mode]
[mode.clients]
enabled = true
refresh = true
misbehaviour = true
[mode.connections]
enabled = true
[mode.channels]
enabled = true
[mode.packets]
enabled = true
clear_interval = 100
clear_on_start = true
tx_confirmation = true
[rest]
enabled = true
host = '0.0.0.0'
port = 3031
[telemetry]
enabled = true
host = '127.0.0.1'
port = 3001
[[chains]]
id = '$MESH_A_E2E_CHAIN_ID'
rpc_addr = 'http://$MESH_A_E2E_VAL_HOST:26657'
grpc_addr = 'http://$MESH_A_E2E_VAL_HOST:9090'
websocket_addr = 'ws://$MESH_A_E2E_VAL_HOST:26657/websocket'
rpc_timeout = '10s'
account_prefix = 'mesh'
key_name = 'val01-mesh-a'
store_prefix = 'ibc'
max_gas = 6000000
gas_price = { price = 0.1, denom = 'umesh' }
gas_multiplier = 1.1
max_msg_num = 30
max_tx_size = 2097152
clock_drift = '5s' # to accomdate docker containers
max_block_time = '30s'
memo_prefix = ''
sequential_batch_tx = false
[chains.trust_threshold]
numerator = '1'
denominator = '3'
[chains.packet_filter]
policy = 'allow'
list = [[
    'transfer',
    'channel-*',
]]
[[chains]]
id = '$MESH_B_E2E_CHAIN_ID'
rpc_addr = 'http://$MESH_B_E2E_VAL_HOST:26657'
grpc_addr = 'http://$MESH_B_E2E_VAL_HOST:9090'
websocket_addr = 'ws://$MESH_B_E2E_VAL_HOST:26657/websocket'
rpc_timeout = '10s'
account_prefix = 'mesh'
key_name = 'val01-mesh-b'
store_prefix = 'ibc'
max_gas = 6000000
gas_price = { price = 0.1, denom = 'umesh' }
gas_multiplier = 1.1
max_msg_num = 30
max_tx_size = 2097152
clock_drift = '5s' # to accomdate docker containers
max_block_time = '30s'
memo_prefix = ''
sequential_batch_tx = false
[chains.trust_threshold]
numerator = '1'
denominator = '3'
[chains.packet_filter]
policy = 'allow'
list = [[
    'transfer',
    'channel-*',
]]
EOF

# import keys

hermes keys add --hd-path "m/44'/330'/0'/0/0" --chain ${MESH_A_E2E_CHAIN_ID} --key-name "val01-mesh-a" --mnemonic-file "${MESH_A_E2E_VAL_MNEMONIC}" --overwrite
hermes keys add --hd-path "m/44'/330'/0'/0/0" --chain ${MESH_B_E2E_CHAIN_ID} --key-name "val01-mesh-b" --mnemonic-file "${MESH_B_E2E_VAL_MNEMONIC}" --overwrite

# start Hermes relayer
hermes start

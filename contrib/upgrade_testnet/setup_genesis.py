import json
import os
import sys
import requests
import hashlib
import base64
import bech32
import cosmospy

args = sys.argv[1:]

# Get the args
build_dir = args[0]
genesis_file = f"{build_dir}/node0/desmos/config/genesis.json"

chain_state_url = args[2]
chain_state_file = f"{build_dir}/state.json"
output_file = f"{build_dir}/output_state.json"

# Get the chain state inside the build dir
with requests.get(chain_state_url) as r, open(chain_state_file, 'w') as f:
    f.write(json.dumps(r.json()))


def get_genesis_validators_delegations(gentxs: [dict]):
    __validators = []
    __delegations = []
    for gentx in gentxs:
        message = gentx['body']['messages'][0]
        __validators.append({
            'commission': {
                'commission_rates': message['commission'],
                'update_time': genesis['genesis_time']
            },
            'consensus_pubkey': message['pubkey'],
            'delegator_shares': message['value']['amount'],
            'description': message['description'],
            'jailed': False,
            'min_self_delegation': message['min_self_delegation'],
            'operator_address': message['validator_address'],
            'status': 'BOND_STATUS_BONDED',
            'tokens': message['value']['amount'],
            'unbonding_height': '0',
            'unbonding_time': '1970-01-01T00:00:00Z'
        })
        __delegations.append({
            'delegator_address': message['delegator_address'],
            'validator_address': message['validator_address'],
            'shares': message['value']['amount']
        })

    return __validators, __delegations


def get_tendermint_validators(gentxs: [dict]):
    __validators = []
    for gentx in gentxs:
        message = gentx['body']['messages'][0]
        pubkey = base64.b64decode(message['pubkey']['key'].encode('ascii'))
        alg = hashlib.sha256()
        alg.update(pubkey)

        __validators.append({
            'address': alg.digest()[:20].hex(),
            'name': message['description']['moniker'],
            'power': message['value']['amount'],
            'pub_key': {
                'type': 'tendermint/PubKeyEd25519',
                'value': message['pubkey']['key']
            }
        })
    return __validators


with open(chain_state_file, 'r') as chain_state_f, open(genesis_file, 'r') as genesis_f, open(output_file, 'w') as out:
    chain_state = json.load(chain_state_f)
    genesis = json.load(genesis_f)

    chain_state['genesis_time'] = genesis['genesis_time']
    chain_state['chain_id'] = genesis['chain_id']
    chain_state['initial_height'] = genesis['initial_height']
    chain_state['app_state']['auth']['accounts'] += genesis['app_state']['auth']['accounts']

    # Transform the gentxs into validators and delegations
    gentxs = genesis['app_state']['genutil']['gen_txs']
    genesis_validators, genesis_delegations = get_genesis_validators_delegations(gentxs)
    tendermint_validators = get_tendermint_validators(gentxs)

    # -------------------------------
    # --- Update the staking state

    # Change all the delegations so that they are delegated to the first validator
    delegations = chain_state['app_state']['staking']['delegations']
    for __delegation in delegations:
        __delegation['validator_address'] = genesis_validators[0]['operator_address']
    delegations += genesis_delegations

    # Get the maps of validator -> [delegation]
    validators_delegations = {}
    for __delegation in delegations:
        __validator_address = __delegation['validator_address']
        __delegations = validators_delegations.get(__validator_address, [])
        __delegations.append(__delegation)
        validators_delegations[__validator_address] = __delegations

    # Merge the delegations for the same delegator and validator together
    final_delegations = []
    for validator, delegations in validators_delegations.items():
        # Map delegator -> shares
        __delegator_shares = {}
        for __delegation in delegations:
            __delegator_address = __delegation.get('delegator_address')
            __shares = float(__delegator_shares.get(__delegator_address, '0'))
            __shares += float(__delegation.get('shares'))
            __delegator_shares[__delegator_address] = "{:.0f}".format(__shares)

        for delegator, shares in __delegator_shares.items():
            final_delegations.append({
                'delegator_address': delegator,
                'validator_address': validator,
                'shares': shares
            })

    # Remove all 0 shares delegations
    for __delegation in final_delegations:
        __delegation_shares = float(__delegation.get('shares'))
        if __delegation_shares < 1:
            final_delegations.remove(__delegation)

    validators_delegators_shares = {}
    for __delegation in final_delegations:
        __validator_address = __delegation.get('validator_address')
        __validator_shares = float(validators_delegators_shares.get(__validator_address, '0'))
        __validator_shares += __delegation_shares
        validators_delegators_shares[__validator_address] = "{:.0f}".format(__validator_shares)

    # Update the validator delegator shares
    for validator in genesis_validators:
        __shares = validators_delegators_shares.get(validator['operator_address'], '0')
        validator['delegator_shares'] = __shares

    # Update the validator voting powers
    final_voting_powers = []
    for validator in genesis_validators:
        final_voting_powers.append({
            'address': validator['operator_address'],
            'power': validator['tokens']
        })

    # Set the final staking state:
    # - validators to only be the genesis ones
    # - validator voting powers
    # - correct delegations
    # - delete the redelegations
    chain_state['app_state']['staking']['validators'] = genesis_validators
    chain_state['app_state']['staking']['last_validator_powers'] = final_voting_powers
    chain_state['app_state']['staking']['delegations'] = final_delegations
    del (chain_state['app_state']['staking']['redelegations'])

    # -----------------------------------------
    # --- Update the distribution state

    # Get the validator starting info
    validator_periods = {}
    genesis_starting_info = []
    for __delegation in final_delegations:
        __validator_address = __delegation['validator_address']

        # Increment of 1 the period
        __period = validator_periods.get(__validator_address, 0) + 1

        # Store the starting info
        genesis_starting_info.append({
            'delegator_address': __delegation['delegator_address'],
            'starting_info': {
                'height': genesis['initial_height'],
                'previous_period': str(__period),
                'stake': __delegation['shares']
            },
            'validator_address': __validator_address
        })

        # Store the updated period
        validator_periods[__validator_address] = __period

    # Get the historical rewards info
    genesis_current_rewards = []
    genesis_validator_historical_rewards = []
    for __validator, __period in validator_periods.items():
        genesis_current_rewards.append({
            'rewards': {
                'period': str(__period_number + 1),
                'rewards': []
            },
            'validator_address': __validator
        })

        for __period_number in range(0, __period):
            genesis_validator_historical_rewards.append({
                'period': str(__period_number + 1),
                'rewards': {
                    'cumulative_reward_ratio': [],
                    'reference_count': 1,
                },
                'validator_address': __validator
            })

    chain_state['app_state']['distribution']['delegator_starting_infos'] = genesis_starting_info
    chain_state['app_state']['distribution']['validator_historical_rewards'] = genesis_validator_historical_rewards
    chain_state['app_state']['distribution']['outstanding_rewards'] = []
    chain_state['app_state']['distribution']['validator_current_rewards'] = genesis_current_rewards
    chain_state['app_state']['distribution']['validator_accumulated_commissions'] = []
    chain_state['app_state']['distribution']['validator_slash_events'] = []

    # # Update the distribution starting info
    # for changed_validator in changed_validators:
    #     for distribution_info in chain_state['app_state']['distribution']['delegator_starting_infos']:
    #         if distribution_info['validator_address'] == changed_validator:
    #             distribution_info['validator_address'] = validator_address

    # -------------------------------
    # --- Update the bank state
    delegated_amount = 0
    for __delegation in genesis_delegations:
        delegated_amount += int(__delegation['shares'])

    chain_state['app_state']['bank']['balances'] += genesis['app_state']['bank']['balances']
    for balance in chain_state['app_state']['bank']['balances']:
        # Remove the distribution balance
        if balance['address'] == 'desmos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8n8fv78':
            balance['coins'][0]['amount'] = '1665300627184'

        # Remove the bonded tokens balance
        if balance['address'] == 'desmos1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3prylw0':
            balance['coins'][0]['amount'] = str(delegated_amount)

        # Remove the unbonded tokens balance
        elif balance['address'] == 'desmos1tygms3xhhs3yv487phx3dw4a95jn7t7l4rcwcm':
            balance['coins'][0]['amount'] = '630839623447'

    # -------------------------------
    # --- Update the slashing state

    signing_infos = []
    for idx, validator in enumerate(tendermint_validators):
        __signing_info = chain_state['app_state']['slashing']['signing_infos'][idx]
        __five_bits_r = bech32.convertbits(bytearray.fromhex(validator['address']), 8, 5)
        __cons_addr = bech32.bech32_encode('desmosvalcons', __five_bits_r)
        signing_infos.append({
            'address': __cons_addr,
            'validator_signing_info': {
                "address": __cons_addr,
                'index_offset': __signing_info['validator_signing_info']['index_offset'],
                'jailed_until': __signing_info['validator_signing_info']['jailed_until'],
                'missed_blocks_counter': __signing_info['validator_signing_info']['missed_blocks_counter'],
                'start_height': __signing_info['validator_signing_info']['start_height'],
                'tombstoned': __signing_info['validator_signing_info']['tombstoned']
            }
        })

    chain_state['app_state']['slashing'] = genesis['app_state']['slashing']
    chain_state['app_state']['slashing']['signing_infos'] = signing_infos

    # -------------------------------
    # --- Fix modules state

    chain_state['app_state']['bank']['supply'] = []

    chain_state['app_state']['gov']['deposit_params']['max_deposit_period'] = '120s'
    chain_state['app_state']['gov']['voting_params']['voting_period'] = '120s'
    chain_state['app_state']['gov']['deposit_params']['min_deposit'] = [{'amount': '10000000', 'denom': 'udaric'}]

    # -------------------------------
    # --- Clear the validators list

    chain_state['validators'] = tendermint_validators

    del (genesis['app_state']['genutil'])

    # -------------------------------
    # --- Write the file

    out.write(json.dumps(chain_state))
    os.system(f"sed -i 's/udsm/udaric/g' {output_file}")

nodes_amount = args[1]
for i in range(0, int(nodes_amount)):
    genesis_path = f"{build_dir}/node{i}/desmos/config/genesis.json"
    with open(genesis_path, 'w') as file:
        os.system(f"cp {output_file} {genesis_path}")

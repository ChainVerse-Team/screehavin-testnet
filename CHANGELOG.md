# CHANGELOG

## Current

### Screehavin-Testnet v1.1.1-beta

+ Add new contracts ( Grant Contract, MarketSaleContract )

+ Update new features: Add Foundation Name, Warning, Ban, Suspend in Staking Contract

+ Refactor flow POM in Screehavin_Core

+ Config new design of Metadata

+ Implement new design of covenant social mobility

+ Update transaction fee share rebalance

+ Implement unstake mechanism

### Screehavin-Testnet v1.1.0

Features:

+ Feat: Update New Metadata

+ Feat: Covenant Social Mobility

+ Feat: Update Staking Contract

+ Feat: Modify Txfee Share


### Screehavin-Testnet v1.1.0-beta

Features:

+ Update Subset Size Calc Function

+ Feat: Modify Purge Time

+ Feat/Add Gas Price

+ Feat: Add API Return Recommend Gas Price Recommend

### Screehavin-Testnet v1.0.10

Features:

+ Feat: Limiting Purge Results

+ Fix: Refactor Earning Stats

+ Store Receipts To Its Cache After Build Block Done

+ Fix: Mitigate Unmarshal RLP Errors 

+ Feat: Deploy k8s 

+ Revert To EVM Queries

+ Fix/Logging Fetch Metada

+ Logging: Log Time To Investigate

+ Fix: Add Map Indexes

+ Feat: Optimize Fetch Data 

+ Hot Fix/Nil Pointer 

+ Hotfix: Txn Snapshot

+ Fix: Increase Default Open File Cap

+ Fix: Update Go Mod

+ Debug Log Amount

+ Fix: Extend Epoch Timeout To 30s

+ Fix: Add Nil Pointer Validate

+ Log Address And Reward Amount

+ Fix Cache, Snapshot Delegators

+ Refactor, And Log Debug For Txn State Root

+ Fix Mounting Wrong Dir Path Inside Container


### Screehavin-Testnet v1.0.9

Features:

+ Feat/Pom Execution Detail

+ Refactor: Update Library

+ Feat: Add Export Port For Prometheus To Access After Running

+ Fix/Leverage Cache

+ Feat: Extend Epoch Block Time In Core

+ Update NewSyncPeerClient And Fix Nil Pointer

+ Panic -> Send On Closed Channel Inside Syncer

+ Feat: Some Events Might Not Be Received By Subscriber

+ Refactor big Int

+ Fix Fetch Covenants At Epoch Snapshot

+ Feat: Inactive Validators Listen To Blk Creation

+ Hotfix: Assign Validator Subset For Writing Seals Incorrectly

+ Fix: Move Global Total Coin Base To Local Scope

+ Feat: Global Counter 

+ Fix: Cache Metadata At Epoch Block & Fix Cache Test

+ Fix Set Node Storage Snapshot

+ Logging Time Execution

+ Feat: Fork From Go-Ibft

+ Change Blocktime To 3s 

### Screehavin-Testnet v1.0.8

Features:

+ Fix: Bulk Sync Nil Pointer

+ Feat: Add Postfix For Logs File, Move Locks To Executor

+ Fix: Thread Locking When Processing Tracking Rewards

+ Feat: Remove Panics In Code And Update Check Nil

+ Feat/group rpc

+ Feat: Add Whoami ci/cd

+ Hotfix: Move Size Check Variable Into Scope 

+ Fix Out Of Range In Sorting Covenant Reward Function

+ Hotfix: Validator Node Catching Up Incorrect Timestamp Data

+ Fix: Nil Pointer Accessing

+ Fix: Covenant Share Transaction Fee

+ Fix: Invalid State Root In Distribute Reward Covenant 

+ Feat: Save Purge History 

+ Feat: First Attempt To Fix Invalid Root Hash 

+ Feat: Config Validator Subset For Small Staker

+ Fix: Errors In Folder Scripts And Query_test File

+ Fix: Go-ibft Version

### Screehavin-Testnet v1.0.7

Features:

+ Fix: Predeploy Slot 

+ Hotfix: Fix Commiting Date

+ Feat: Proposer Delegation 

+ Refactor: Add Logger Location

+ Feat: Predeploy Withdrawal Address

+ Update Api Portal

+ Feat: refactor reward seed db

+ Feat: Run Docker in node server


### Screehavin-Testnet v1.0.6

Features:

+ Restore Seed Init

+ Remove Log

+ Check Errs, Optimize Writing Seeds

+ Change Geth Module


### ScreeHavin-Testnet v1.0.5

Features:

+ Revert Contract Calls


### ScreeHavin-Testnet v1.0.4

Features:

+ Add Track Total Supply Covenant

+ Refactor Tracking Block Rewards


### ScreeHavin-Testnet v1.0.3

Features:

+ Update Blockreward Validator Distribution in 20 years

+ Update Increase Covenant Share Of Bonus


### ScreeHavin-Testnet v1.0.2

Features:

+ Update Epoch Size = 150 blocks

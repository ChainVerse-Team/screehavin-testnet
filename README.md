# Screehavin-Testnet

Screehavin is a customized blockchain protocol based on polygon edge.

# Minimum System Requirements:

RAM: 2 GB

CPU: 2 cores CPU 

Disk: 

+ 10 GB root patition

+ 30 GB root partition with LVM for disk extension

# 

# Overview

### Screehavin Explorer:  [http://explorer.screehavin.io/](http://explorer.screehavin.io/)

![image](https://user-images.githubusercontent.com/55268800/208595093-77cfbbf7-7c8a-4adb-b21a-2d2508eaedf4.png)

### Demo: View nodes in blockchain

![image](https://user-images.githubusercontent.com/55268800/208595265-b7c106d2-7ffa-40a1-adf9-1ebade64152e.png)

***

# How to add screehavin on metamask

**Step 1**: Go to the network => Add Network
# # 
![image](https://user-images.githubusercontent.com/55268800/208595850-28c77b45-2fa3-4aed-af22-3cbfcc001017.png)

**Step 2**: Fill in these informations
<img width="940" alt="Screenshot 2023-01-10 at 10 42 58" src="https://user-images.githubusercontent.com/55268800/211457147-6b7e3a04-51de-402a-9dc5-3902bb2f0a28.png">


```ruby
Name : screehavin testnet
RPC URL : http://103.138.113.121:8545/ (IP is interchangeable with other server’s IP, default server’s RPC port is 8545)
ChainID : 100 (default)
Currency Symbol: SCREE
Block explorer URL: http://explorer.screehavin.io/
```

***

# BlockChain Development Procedure

## Deployment steps

### Clone a repository into a new directory

**Step 1**: 


(if you haven't downloaded `git`, please follow the link https://git-scm.com/downloads)

`git version`

`git clone https://github.com/ChainVerse-Team/screehavin-testnet.git`

`cd screehavin-testnet`

***

### Create a data directory

**Step 2**: `make data-dir-macOS` or `make data-dir-linux` or `make data-dir-windows`

(choose the operating system that is compatible with your computer)

***

### Setting environment

**Step 3**: Create `.env` file

```ruby
STAKING_CONTRACT=0x0000000000000000000000000000000000001001
JSONRPC_URL=http://103.138.113.121:8545/
```
***

### How to install the Make package on server

**Step 4**:
`sudo apt update`
`sudo apt install make`
`make -version`

(if you run Windows, please follow the instructions in the link https://www.technewstoday.com/install-and-use-make-in-windows/)

### Put in public and private key 

***

**Step 5**: There are two ways to implement:

1. you use the available information in `data-dir` folder.

Example:
```ruby
[SECRETS INIT]
Public key (address) = 0xfDA89fDE5917d88a7823925e810B0FB73cE87A65
BLS Public key       = 0xadb78c42d3b13c026e60750499c0dc0d5a618a8671af1fe49c52fe7101bf16e9147f125e46e7d1fe6b5bdee8e9fdb2d9
Node ID              = 16Uiu2HAkzDhKRhdGdmBwdUTLy5W5ee2DGKTy1JikEHm1PxesMZ24
```

2. you find `validator.key` file in `data-dir/consensus`  and then replace it with your private-key.

(Note: check your private-key in metamask)

<img width="326" alt="Screenshot 2023-01-05 at 10 11 06" src="https://user-images.githubusercontent.com/55268800/210697465-1cb8b817-9431-4c00-b66b-8cf1973bff77.png">

***

### Run server makefile (choose Linux, Windows, MacOS)

**Step 6**: `make run-macOS` or `make run-windows` or `make run-linux`

(choose the operating system that is compatible with your computer)

***

### Request for testnet staking tokens

**Step 7**:

=> Telegram: 'https://t.me/screehavintestnet'

=> Email: `hello@chainverse.org`

***

### Become a validator

**Step 8**:
you find `keys.text` in `scripts/keys.txt` and add your private-key in it.
Next, you run command `make batch-staking`

(VALIDATOR_THRESHOLD = 10 SCREE)


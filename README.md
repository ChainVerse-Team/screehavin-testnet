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
<img width="979" alt="Screenshot 2022-12-26 at 16 59 31" src="https://user-images.githubusercontent.com/55268800/209535494-88e788e4-c593-4de6-9d59-90bfeeae2bea.png">

```ruby
Name : screehavin testnet
RPC URL : http://103.138.113.121:8545/ (IP is interchangeable with other server’s IP, default server’s RPC port is 8545)
ChainID : 100 (default)
Currency Symbol: SCREE
Block explorer URL: http://explorer.screehavin.io/
```

***

# BlockChain Development Procedure

## Deployment Steps

**Step 1**: Add file screehavin in server ubuntu
`git clone https://github.com/ChainVerse-Team/screehavin-testnet.git`

## Setting Environment
**Step 2**: Create `.env` file

```ruby
STAKING_CONTRACT=0x0000000000000000000000000000000000001001
JSONRPC_URL=http://103.138.113.121:8545/
```

### How to install the Make package on Ubuntu
**Step 3**:
`sudo apt update`
`sudo apt install make`
`make -version`

### Put in public and private key 
**Step 4**: you find `validator.key` file in `data-dir` folder and then replace private-key in this file.

(Note: check your private-key in metamask)

<img width="326" alt="Screenshot 2023-01-05 at 10 11 06" src="https://user-images.githubusercontent.com/55268800/210697465-1cb8b817-9431-4c00-b66b-8cf1973bff77.png">

### Run server makefile (choose Linux, Windows, MacOS)
**Step 5**: `make run-macOS` or `make run-windows` or `make run-linux`

### Request for testnet staking tokens
**Step 6**:
### Telegram: 'https://t.me/screehavintestnet'
### Email: `hello@chainverse.org`

### Become A Validator
**Step 7**:
you find `key.text` in scripts/key.txt and add your private-key in it.
Next, you run command `make batch-staking`

(VALIDATOR_THRESHOLD = 10 SCREE)


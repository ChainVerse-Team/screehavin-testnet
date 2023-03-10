# Screehavin-Testnet

Screehavin is a customized blockchain protocol based on polygon edge.

# Minimum System Requirements:

RAM: 2 GB

CPU: 2 cores CPU 

Disk: 

+ 10 GB root patition

+ 30 GB root partition with LVM for disk extension

# 

# Overview 📝

### Screehavin Explorer:  [http://explorer.screehavin.com/](http://explorer.screehavin.com/)

<img width="1496" alt="Screenshot 2023-02-28 at 14 02 43" src="https://user-images.githubusercontent.com/55268800/221778698-26b83edf-fbfe-4e4a-b830-69769c2e3523.png">

### Demo: View nodes in blockchain

![image](https://user-images.githubusercontent.com/55268800/208595265-b7c106d2-7ffa-40a1-adf9-1ebade64152e.png)

***

# How to add screehavin on metamask

**Step 1**: Go to the network => Add Network
# # 
![image](https://user-images.githubusercontent.com/55268800/208595850-28c77b45-2fa3-4aed-af22-3cbfcc001017.png)

**Step 2**: Fill in these informations
<img width="962" alt="Screenshot 2023-03-10 at 16 29 47" src="https://user-images.githubusercontent.com/55268800/224278707-7ae3d254-a2c4-4163-a885-8b8acc2328a9.png">


```ruby
Name : screehavin testnet
RPC URL : http://103.138.113.121:8545/ (IP is interchangeable with other server’s IP, default server’s RPC port is 8545)
ChainID : 100 (default)
Currency Symbol: SCREE
Block explorer URL: http://explorer.screehavin.com/
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

### How to install the Make package on server

**Step 2**:
`sudo apt update`
`sudo apt install make`
`make -version`

(if you run Windows, please follow the instructions in the link https://www.technewstoday.com/install-and-use-make-in-windows/)

### Create a data directory

**Step 3**: `make data-dir-macOS` or `make data-dir-linux` or `make data-dir-windows`

(choose the operating system that is compatible with your computer)

***

### Setting environment

**Step 4**: Create `.env` file

```ruby
STAKING_CONTRACT=0x0000000000000000000000000000000000001001
JSONRPC_URL=http://103.138.113.121:8545/
```
***

### How to get private key in data-dir

**Step 5**:

<img width="326" alt="Screenshot 2023-01-05 at 10 11 06" src="https://user-images.githubusercontent.com/55268800/210697465-1cb8b817-9431-4c00-b66b-8cf1973bff77.png">

you find `validator.key` file in `data-dir/consensus`.

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
you find `keys.text` in `scripts/keys.txt` and add your private-key in it (you get private-key in `data-dir/consensus/validator.key`).
Next, you run command `make batch-staking`

(VALIDATOR_THRESHOLD = 10 SCREE)


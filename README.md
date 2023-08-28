![screehavin_logo](https://github.com/ChainVerse-Team/screehavin-testnet/assets/55268800/5239e348-16d2-434b-bff8-410b6182b7f4)

# Screehavin-Testnet

Screehavin is a customized blockchain protocol based on polygon edge.

# Minimum System Requirements:

| Component| Minimum Requirement |  Recommended  |
|----------|---------------------|---------------|
| Processor|     2-core CPU      |  4-core CPU   |
|  Memory  |      2 GB RAM       |  4 GB RAM     |
|  Storage |     25 GB SSD       |   50 GB SSD   |  
|  Network | High-speed internet connection | Dedicated server with gigabit connection |

# 

# Overview 📝

### Screehavin Explorer:  [http://explorer.screehavin.com/](http://explorer.screehavin.com/)

<img width="1300" alt="Screenshot 2023-04-03 at 15 33 02" src="https://user-images.githubusercontent.com/55268800/263613497-a98882db-e8ee-4120-bec8-4e640b153945.png">

### Demo: View nodes in blockchain

<img width="1300" src="https://user-images.githubusercontent.com/55268800/208595265-b7c106d2-7ffa-40a1-adf9-1ebade64152e.png">

***

# How to add screehavin on metamask

**Step 1**: Go to the network => Add Network

<img width="300" src="https://user-images.githubusercontent.com/55268800/208595850-28c77b45-2fa3-4aed-af22-3cbfcc001017.png"><br>

**Step 2**: Fill in these informations

<img width="700" alt="Screenshot 2023-03-10 at 16 29 47" src="https://user-images.githubusercontent.com/55268800/224278707-7ae3d254-a2c4-4163-a885-8b8acc2328a9.png">


```ruby
Name : screehavin testnet
RPC URL : http://103.138.113.121:8545
ChainID : 100 (default)
Currency Symbol: SCREE
Block explorer URL: http://explorer.screehavin.com/
```

***

# :hammer_and_wrench: BlockChain Development Procedure

## Deployment steps

### Clone a repository into a new directory

**Step 1**: Open terminal on your PC


(if you haven't downloaded `git`, please follow the link https://git-scm.com/downloads)

`git version`

![image](https://github.com/ChainVerse-Team/screehavin-testnet/assets/55268800/7b5dd9a3-d60d-4f57-9549-4f8d9a2ce941)


`git clone https://github.com/ChainVerse-Team/screehavin-testnet.git`

![telegram-cloud-photo-size-5-6161413019322661258-x](https://github.com/ChainVerse-Team/screehavin-testnet/assets/55268800/d27903b6-7a23-4ab4-b493-78c49b62c3d2)


`cd screehavin-testnet`

![telegram-cloud-photo-size-5-6161413019322661262-x](https://github.com/ChainVerse-Team/screehavin-testnet/assets/55268800/d4744b68-cb72-42a3-96d1-6b29999bd682)


***

### How to install the Make package on server

**Step 2**: For Linux and MacOS users, you can execute this command:

`sudo apt update`
`sudo apt install make`
`make -version`

![telegram-cloud-photo-size-5-6161413019322661263-x](https://github.com/ChainVerse-Team/screehavin-testnet/assets/55268800/1cda0530-5b52-4e31-ad2a-29bff7f8197e)


For Windows users, please follow the instructions in the link https://www.technewstoday.com/install-and-use-make-in-windows/ (we recommend using 
 the chocolatey method)

### Create a data directory

**Step 3**: Choose the operating system that is compatible with your computer

`make data-dir-macOS` or `make data-dir-linux` or `make data-dir-windows`

![telegram-cloud-photo-size-5-6161413019322661264-y](https://github.com/ChainVerse-Team/screehavin-testnet/assets/55268800/0fc41e2d-74eb-44cd-ab14-8cd0b1a89e8a)


***

### Setting environment

**Step 4**: Create `.env` file in screehavin-testnet folder

For example: In Windows

C :\users\ <PC_name> \screehavin-testnet

```ruby
STAKING_CONTRACT=0x0000000000000000000000000000000000001001
JSONRPC_URL=http://103.138.113.121:8545/
```
***

### How to get private key in data-dir

**Step 5**:

<img width="326" alt="Screenshot 2023-01-05 at 10 11 06" src="https://user-images.githubusercontent.com/55268800/210697465-1cb8b817-9431-4c00-b66b-8cf1973bff77.png">

you find `validator.key` file in `data-dir/consensus`.

(do not reveal your private key to anyone)

***

### Run server makefile (choose Linux, Windows, MacOS)

**Step 6**: `make run-macOS` or `make run-windows` or `make run-linux`

(choose the operating system that is compatible with your computer)

***

### Request for testnet staking tokens

**Step 7**: Contact us to request tokens

Email: `hello@chainverse.org`

<div>
  <a href="https://t.me/screehavintestnet">
   <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/8/82/Telegram_logo.svg/600px-Telegram_logo.svg.png" alt="Telegram" width="32">
  </a>
  
  <a href="https://twitter.com/screehavin">
   <img src="https://cdn.icon-icons.com/icons2/836/PNG/512/Twitter_icon-icons.com_66803.png" alt="Twitter" width="32">
  </a>
</div>

***

### Become a validator

**Step 8**: `make batch-staking`

(VALIDATOR_THRESHOLD = 10 SCREE)


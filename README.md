# screehavin-testnet

# Overview

### ScreeHavin Explorer:  [http://explorer.screehavin.io/](http://explorer.screehavin.io/)

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

**Step 2**: How to install go on Ubuntu 
https://linuxhint.com/install-go-ubuntu-2/


## Install Dependencies
### How to install the Make package on Ubuntu
`sudo apt update`
`sudo apt install make`
`make -version`

![image](https://user-images.githubusercontent.com/55268800/208600604-0d3c60ee-3ac9-412c-9cc8-9e212704aea3.png)

***

# Run Commands
## CLI Commands
| Command | Usage |
| --- | --- |
|data| Create data directory in nodes |
| make run1  | Run test server 1, used on testnet |
| make run2  | Run test server 2, used on testnet |
| make run3  | Run test server 3, used on testnet |
| make run4  | Run test server 4, used on testnet |

## Become A Validator
you can stake to Staking Contract `0x0000000000000000000000000000000000001001`

### Please contract with us if you want to request for testnet tokens
### Email : `hello@chainverse.org`






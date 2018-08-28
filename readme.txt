//注释掉hcashd里面dns的代码，启动如下就为私链


/root/gopath/src/github.com/HcashOrg/hcashd/hcashd --listen=127.0.0.1:11010 --rpclisten=127.0.0.1:12010 --rpcuser=bitcoinrpc --rpcpass=123456 --blockmaxsize=999000

/root/gopath/src/github.com/HcashOrg/hcashd/hcashctl --rpcserver=127.0.0.1:12010 --rpcuser=bitcoinrpc --rpcpass=123456 getinfo

//wallet启动
/root/gopath/src/github.com/HcashOrg/hcashwallet/hcashwallet -C ./hcashwallet.conf --create
/root/gopath/src/github.com/HcashOrg/hcashwallet/hcashwallet -C ./hcashwallet.conf

//连接到wallet执行
/root/gopath/src/github.com/HcashOrg/hcashd/hcashctl --rpcserver=127.0.0.1:14010 --rpcuser=bitcoinrpc --rpcpass=123456 --wallet createnewaccount "test"

/root/gopath/src/github.com/HcashOrg/hcashd/hcashctl --rpcserver=127.0.0.1:14010 --rpcuser=bitcoinrpc --rpcpass=123456 --wallet -l

/root/gopath/src/github.com/HcashOrg/hcashd/hcashctl --rpcserver=127.0.0.1:14010 --rpcuser=bitcoinrpc --rpcpass=123456 --wallet getnewaddress
HsEvRq9fxQ4pW5AJr1Cffia3zzepEELiQry

/root/gopath/src/github.com/HcashOrg/hcashd/hcashctl --rpcserver=127.0.0.1:14010 --rpcuser=bitcoinrpc --rpcpass=123456 generate 10

/root/gopath/src/github.com/HcashOrg/hcashd/hcashd --listen=127.0.0.1:11010 --rpclisten=127.0.0.1:12010 --rpcuser=bitcoinrpc --rpcpass=123456 --blockmaxsize=999000 --generate --miningaddr=HsEvRq9fxQ4pW5AJr1Cffia3zzepEELiQry

/root/gopath/src/github.com/HcashOrg/hcashd/hcashctl --rpcserver=127.0.0.1:14010 --rpcuser=bitcoinrpc --rpcpass=123456 --wallet getbalance

/root/gopath/src/github.com/HcashOrg/hcashd/hcashctl --rpcserver=127.0.0.1:14010 --rpcuser=bitcoinrpc --rpcpass=123456 --wallet setgenerate generate

/root/gopath/src/github.com/HcashOrg/hcashd/hcashctl --rpcserver=127.0.0.1:14010 --rpcuser=bitcoinrpc --rpcpass=123456 --wallet getreceivedbyaccount default


//hcashd 12010 hcashwallet 14010

目前问题：
1、hcashwallet是否是HD钱包，其中的命令getmasterpubkey需要加参数account，返回值的含义是
	master/purposecode/coincode/account 这个路径的公钥？
2、hcashwallet gui启动会启动hcashd、hcashwallet及hcashctl，其中hcashd没挖矿的情况下占用cpu高，但linux版的没问题

3、hcashwallet.conf 连接到hcashd端口地址不可配置

4、gui异常退出后hcashd有可能没有退出，且占cpu高

5、hcashd.conf里的配置  datadir=~/.hcashd必须默认，更改为其他地方会带来找不到数据库的问题，因为其他一些信息比如认证文件等还是在~/.hcashd并没有更改位置

Your wallet generation seed is:
skullcap yesteryear ragtime pedigree surmount typewriter 
Zulu infancy tumor direction egghead inferno 
woodlark paramount berserk caravan woodlark opulent 
stagnate voyager hockey borderline enlist Waterloo 
kiwi hemisphere steamship impetus scenic paperweight 
fracture vacancy accrue 

Hex: bdfea0aadae8ff7aec4a567bfea81e25fe9fd0f773185af97c6fd476b4a565f1



https://docs.decred.org/getting-started/user-guides/agenda-voting/#voting-preparation 投票过程

winningTickets：选出幸运的投票
calcNextRequiredDifficulty：计算下一个PoW的难度
calcNextRequiredStakeDifficulty：计算下一个PoS的难度
CreatePremineBlock：制作下个区块模板
NextBlock
交易类型：
Regular：正常的交易
SStx (fresh stake tickets)：
SSGen (votes)：
SSRtx (revocations for missed tickets)


//买票相关
./hcashctl -C ./testhcashctl.conf --wallet purchaseticket "default" 1000
[
  "d15b82f87befb51a56d2101cd040d6bd6ed6d27e4e52ba91978e7879a90b73ed"
]

./hcashctl -C ./testhcashctl.conf --wallet getvotechoices
{
  "version": 5,
  "choices": [
    {
      "agendaid": "sdiffalgorithm",
      "agendadescription": "Change stake difficulty algorithm as defined in DCP0001",
      "choiceid": "abstain",
      "choicedescription": "abstain voting for change"
    }
  ]
}

./hcashctl -C ./testhcashctl.conf --wallet setvotechoice sdiffalgorithm yes

./hcashctl -C ./testhcashctl.conf --wallet sendtossgen "default" "d15b82f87befb51a56d2101cd040d6bd6ed6d27e4e52ba91978e7879a90b73ed" "0000cf4919bf268129cac05abc539754285bfcdf5ebae8171d58bdcdf8c212b8" 142 votebits ("comment")
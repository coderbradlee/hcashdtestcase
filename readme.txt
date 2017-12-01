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

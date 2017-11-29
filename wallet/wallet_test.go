// Copyright (c) 2016 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wallet

import (
	// "bytes"
	// "encoding/hex"
	// "strings"
	"testing"
	// "bufio"
	
	// "fmt"
	// "io"
	// "os"
	
	// "github.com/HcashOrg/hcashrpcclient"
	"github.com/HcashOrg/hcashd/hcashjson"
)
//wallet test command list

// accountaddressindex "account" branch
// accountsyncaddressindex "account" branch index
// addmultisigaddress nrequired ["key",...] ("account")
// addticket "tickethex"
// calcpowsubsidy ("account" from=0 count=0 includewatchonly=false)
// consolidate inputs ("account" "address")
// createmultisig nrequired ["key",...]
// createnewaccount "account"
// createrawssgentx [{"txid":"value","vout":n,"tree":n},...] votebits
// createrawssrtx [{"txid":"value","vout":n,"tree":n},...] (fee)
// createrawsstx [{"txid":"value","vout":n,"tree":n,"amt":n},...] amount [{"add
// r":"value","commitamt":n,"changeaddr":"value","changeamt":n},...]
// dumpprivkey "address"
// estimatepriority numblocks
// generatevote "blockhash" height keyheight "tickethash" votebits "votebitsext
// "
// getaccount "address"
// getaccountaddress "account"
// getaddressesbyaccount "account"
// getbalance ("account" minconf=1)
// getmasterpubkey ("account")
// getmultisigoutinfo "hash" index
// getnewaddress ("account" "gappolicy")
// getrawchangeaddress ("account")
// getreceivedbyaccount "account" (minconf=1)
// getreceivedbyaddress "address" (minconf=1)
// getseed
// getstakeinfo (isliveticketdetails=false)
// getticketfee
// gettickets includeimmature
// gettransaction "txid" (includewatchonly=false)
// getvotechoices
// getwalletfee
// importaddress "address" (rescan=true)
// importprivkey "privkey" ("label" rescan=true scanfrom)
// importpubkey "pubkey" (rescan=true)
// importscript "hex" (rescan=true scanfrom)
// keypoolrefill (newsize=100)
// listaccounts (minconf=1)
// listlockunspent
// listreceivedbyaccount (minconf=1 includeempty=false includewatchonly=false)
// listreceivedbyaddress (minconf=1 includeempty=false includewatchonly=false)
// listscripts
// listsinceblock ("blockhash" targetconfirmations=1 includewatchonly=false)
// listtransactions ("account" count=10 from=0 includewatchonly=false)
// listtxs ("account" txtype=63 count=10 from=0 includewatchonly=false)
// listunspent (minconf=1 maxconf=9999999 ["address",...])
// lockunspent unlock [{"txid":"value","vout":n,"tree":n},...]
// purchaseticket "fromaccount" spendlimit (minconf=1 "ticketaddress" numticket
// s "pooladdress" poolfees expiry "comment")
// redeemmultisigout "hash" index tree ("address")
// redeemmultisigouts "fromscraddress" ("toaddress" number)
// renameaccount "oldaccount" "newaccount"
// rescanwallet (beginheight=0)
// revoketickets
// sendfrom "fromaccount" "toaddress" amount (notsend=0 minconf=1 "comment" "co
// mmentto")
// sendmany "fromaccount" {"address":amount,...} (notsend=0 minconf=1 "comment"
// )
// sendtoaddress "address" amount (notsend=0 "comment" "commentto")
// sendtomultisig "fromaccount" amount ["pubkey",...] (nrequired=1 minconf=1 "c
// omment")
// sendtossgen "fromaccount" "tickethash" "blockhash" height votebits ("comment
// ")
// sendtossrtx "fromaccount" "tickethash" ("comment")
// sendtosstx "fromaccount" amounts [{"txid":"value","vout":n,"tree":n,"amt":n}
// ,...] [{"addr":"value","commitamt":n,"changeaddr":"value","changeamt":n},...
// ] (minconf=1 "comment")
// setbalancetomaintain balance
// setticketfee fee
// setticketmaxprice max
// settxfee amount
// setvotechoice "agendaid" "choiceid"
// signmessage "address" "message"
// signrawtransaction "rawtx" ([{"txid":"value","vout":n,"tree":n,"scriptpubkey
// ":"value","redeemscript":"value"},...] ["privkey",...] flags="ALL")
// signrawtransactions ["rawtx",...] (send=true)
// stakepooluserinfo "user"
// walletinfo
// walletlock
// walletpassphrase "passphrase" timeout
// walletpassphrasechange "oldpassphrase" "newpassphrase"

// func TestDecodeHex(t *testing.T) {
// 	test2()
// }
func TestAll(t *testing.T) {
	t.Parallel()

	// testID := int(1)
	tests := []struct {
		name         string
		newCmd       func() (interface{}, error)
		staticCmd    func() interface{}
		marshalled   string
		unmarshalled interface{}
	}{
		{
			name: "addmultisigaddress",
			newCmd: func() (interface{}, error) {
				return hcashjson.NewCmd("addmultisigaddress", 2, []string{"031234", "035678"})
			},
			staticCmd: func() interface{} {
				keys := []string{"031234", "035678"}
				return hcashjson.NewAddMultisigAddressCmd(2, keys, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"addmultisigaddress","params":[2,["031234","035678"]],"id":1}`,
			unmarshalled: &hcashjson.AddMultisigAddressCmd{
				NRequired: 2,
				Keys:      []string{"031234", "035678"},
				Account:   nil,
			},
		},
		{
			name: "addmultisigaddress optional",
			newCmd: func() (interface{}, error) {
				return hcashjson.NewCmd("addmultisigaddress", 2, []string{"031234", "035678"}, "test")
			},
			staticCmd: func() interface{} {
				keys := []string{"031234", "035678"}
				return hcashjson.NewAddMultisigAddressCmd(2, keys, hcashjson.String("test"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"addmultisigaddress","params":[2,["031234","035678"],"test"],"id":1}`,
			unmarshalled: &hcashjson.AddMultisigAddressCmd{
				NRequired: 2,
				Keys:      []string{"031234", "035678"},
				Account:   hcashjson.String("test"),
			},
		},
		{
			name: "createmultisig",
			newCmd: func() (interface{}, error) {
				return hcashjson.NewCmd("createmultisig", 2, []string{"031234", "035678"})
			},
			staticCmd: func() interface{} {
				keys := []string{"031234", "035678"}
				return hcashjson.NewCreateMultisigCmd(2, keys)
			},
			marshalled: `{"jsonrpc":"1.0","method":"createmultisig","params":[2,["031234","035678"]],"id":1}`,
			unmarshalled: &hcashjson.CreateMultisigCmd{
				NRequired: 2,
				Keys:      []string{"031234", "035678"},
			},
		},
		{
			name: "dumpprivkey",
			newCmd: func() (interface{}, error) {
				return hcashjson.NewCmd("dumpprivkey", "HsEvRq9fxQ4pW5AJr1Cffia3zzepEELiQry")
			},
			staticCmd: func() interface{} {
				return hcashjson.NewDumpPrivKeyCmd("1Address")
			},
			marshalled: `{"jsonrpc":"1.0","method":"dumpprivkey","params":["1Address"],"id":1}`,
			unmarshalled: &hcashjson.DumpPrivKeyCmd{
				Address: "1Address",
			},
		},
		{
			name: "estimatefee",
			newCmd: func() (interface{}, error) {
				return hcashjson.NewCmd("estimatefee", 6)
			},
			staticCmd: func() interface{} {
				return hcashjson.NewEstimateFeeCmd(6)
			},
			marshalled: `{"jsonrpc":"1.0","method":"estimatefee","params":[6],"id":1}`,
			unmarshalled: &hcashjson.EstimateFeeCmd{
				NumBlocks: 6,
			},
		},
		// {
		// 	name: "estimatepriority",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("estimatepriority", 6)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewEstimatePriorityCmd(6)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"estimatepriority","params":[6],"id":1}`,
		// 	unmarshalled: &hcashjson.EstimatePriorityCmd{
		// 		NumBlocks: 6,
		// 	},
		// },
		// {
		// 	name: "getaccount",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("getaccount", "1Address")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewGetAccountCmd("1Address")
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"getaccount","params":["1Address"],"id":1}`,
		// 	unmarshalled: &hcashjson.GetAccountCmd{
		// 		Address: "1Address",
		// 	},
		// },
		// {
		// 	name: "getaccountaddress",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("getaccountaddress", "acct")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewGetAccountAddressCmd("acct")
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"getaccountaddress","params":["acct"],"id":1}`,
		// 	unmarshalled: &hcashjson.GetAccountAddressCmd{
		// 		Account: "acct",
		// 	},
		// },
		// {
		// 	name: "getaddressesbyaccount",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("getaddressesbyaccount", "acct")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewGetAddressesByAccountCmd("acct")
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"getaddressesbyaccount","params":["acct"],"id":1}`,
		// 	unmarshalled: &hcashjson.GetAddressesByAccountCmd{
		// 		Account: "acct",
		// 	},
		// },
		// {
		// 	name: "getbalance",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("getbalance")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewGetBalanceCmd(nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"getbalance","params":[],"id":1}`,
		// 	unmarshalled: &hcashjson.GetBalanceCmd{
		// 		Account: nil,
		// 		MinConf: hcashjson.Int(1),
		// 	},
		// },
		// {
		// 	name: "getbalance optional1",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("getbalance", "acct")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewGetBalanceCmd(hcashjson.String("acct"), nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"getbalance","params":["acct"],"id":1}`,
		// 	unmarshalled: &hcashjson.GetBalanceCmd{
		// 		Account: hcashjson.String("acct"),
		// 		MinConf: hcashjson.Int(1),
		// 	},
		// },
		// {
		// 	name: "getbalance optional2",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("getbalance", "acct", 6)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewGetBalanceCmd(hcashjson.String("acct"), hcashjson.Int(6))
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"getbalance","params":["acct",6],"id":1}`,
		// 	unmarshalled: &hcashjson.GetBalanceCmd{
		// 		Account: hcashjson.String("acct"),
		// 		MinConf: hcashjson.Int(6),
		// 	},
		// },
		// {
		// 	name: "getnewaddress",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("getnewaddress")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewGetNewAddressCmd(nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"getnewaddress","params":[],"id":1}`,
		// 	unmarshalled: &hcashjson.GetNewAddressCmd{
		// 		Account:   nil,
		// 		GapPolicy: nil,
		// 	},
		// },
		// {
		// 	name: "getnewaddress optional",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("getnewaddress", "acct", "ignore")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewGetNewAddressCmd(hcashjson.String("acct"), hcashjson.String("ignore"))
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"getnewaddress","params":["acct","ignore"],"id":1}`,
		// 	unmarshalled: &hcashjson.GetNewAddressCmd{
		// 		Account:   hcashjson.String("acct"),
		// 		GapPolicy: hcashjson.String("ignore"),
		// 	},
		// },
		// {
		// 	name: "getrawchangeaddress",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("getrawchangeaddress")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewGetRawChangeAddressCmd(nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"getrawchangeaddress","params":[],"id":1}`,
		// 	unmarshalled: &hcashjson.GetRawChangeAddressCmd{
		// 		Account: nil,
		// 	},
		// },
		// {
		// 	name: "getrawchangeaddress optional",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("getrawchangeaddress", "acct")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewGetRawChangeAddressCmd(hcashjson.String("acct"))
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"getrawchangeaddress","params":["acct"],"id":1}`,
		// 	unmarshalled: &hcashjson.GetRawChangeAddressCmd{
		// 		Account: hcashjson.String("acct"),
		// 	},
		// },
		// {
		// 	name: "getreceivedbyaccount",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("getreceivedbyaccount", "acct")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewGetReceivedByAccountCmd("acct", nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"getreceivedbyaccount","params":["acct"],"id":1}`,
		// 	unmarshalled: &hcashjson.GetReceivedByAccountCmd{
		// 		Account: "acct",
		// 		MinConf: hcashjson.Int(1),
		// 	},
		// },
		// {
		// 	name: "getreceivedbyaccount optional",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("getreceivedbyaccount", "acct", 6)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewGetReceivedByAccountCmd("acct", hcashjson.Int(6))
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"getreceivedbyaccount","params":["acct",6],"id":1}`,
		// 	unmarshalled: &hcashjson.GetReceivedByAccountCmd{
		// 		Account: "acct",
		// 		MinConf: hcashjson.Int(6),
		// 	},
		// },
		// {
		// 	name: "getreceivedbyaddress",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("getreceivedbyaddress", "1Address")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewGetReceivedByAddressCmd("1Address", nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"getreceivedbyaddress","params":["1Address"],"id":1}`,
		// 	unmarshalled: &hcashjson.GetReceivedByAddressCmd{
		// 		Address: "1Address",
		// 		MinConf: hcashjson.Int(1),
		// 	},
		// },
		// {
		// 	name: "getreceivedbyaddress optional",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("getreceivedbyaddress", "1Address", 6)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewGetReceivedByAddressCmd("1Address", hcashjson.Int(6))
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"getreceivedbyaddress","params":["1Address",6],"id":1}`,
		// 	unmarshalled: &hcashjson.GetReceivedByAddressCmd{
		// 		Address: "1Address",
		// 		MinConf: hcashjson.Int(6),
		// 	},
		// },
		// {
		// 	name: "gettransaction",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("gettransaction", "123")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewGetTransactionCmd("123", nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"gettransaction","params":["123"],"id":1}`,
		// 	unmarshalled: &hcashjson.GetTransactionCmd{
		// 		Txid:             "123",
		// 		IncludeWatchOnly: hcashjson.Bool(false),
		// 	},
		// },
		// {
		// 	name: "gettransaction optional",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("gettransaction", "123", true)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewGetTransactionCmd("123", hcashjson.Bool(true))
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"gettransaction","params":["123",true],"id":1}`,
		// 	unmarshalled: &hcashjson.GetTransactionCmd{
		// 		Txid:             "123",
		// 		IncludeWatchOnly: hcashjson.Bool(true),
		// 	},
		// },
		// {
		// 	name: "importprivkey",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("importprivkey", "abc")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewImportPrivKeyCmd("abc", nil, nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"importprivkey","params":["abc"],"id":1}`,
		// 	unmarshalled: &hcashjson.ImportPrivKeyCmd{
		// 		PrivKey: "abc",
		// 		Label:   nil,
		// 		Rescan:  hcashjson.Bool(true),
		// 	},
		// },
		// {
		// 	name: "importprivkey optional1",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("importprivkey", "abc", "label")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewImportPrivKeyCmd("abc", hcashjson.String("label"), nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"importprivkey","params":["abc","label"],"id":1}`,
		// 	unmarshalled: &hcashjson.ImportPrivKeyCmd{
		// 		PrivKey: "abc",
		// 		Label:   hcashjson.String("label"),
		// 		Rescan:  hcashjson.Bool(true),
		// 	},
		// },
		// {
		// 	name: "importprivkey optional2",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("importprivkey", "abc", "label", false)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewImportPrivKeyCmd("abc", hcashjson.String("label"), hcashjson.Bool(false), nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"importprivkey","params":["abc","label",false],"id":1}`,
		// 	unmarshalled: &hcashjson.ImportPrivKeyCmd{
		// 		PrivKey: "abc",
		// 		Label:   hcashjson.String("label"),
		// 		Rescan:  hcashjson.Bool(false),
		// 	},
		// },
		// {
		// 	name: "importprivkey optional3",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("importprivkey", "abc", "label", false, 12345)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewImportPrivKeyCmd("abc", hcashjson.String("label"), hcashjson.Bool(false), hcashjson.Int(12345))
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"importprivkey","params":["abc","label",false,12345],"id":1}`,
		// 	unmarshalled: &hcashjson.ImportPrivKeyCmd{
		// 		PrivKey:  "abc",
		// 		Label:    hcashjson.String("label"),
		// 		Rescan:   hcashjson.Bool(false),
		// 		ScanFrom: hcashjson.Int(12345),
		// 	},
		// },
		// {
		// 	name: "keypoolrefill",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("keypoolrefill")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewKeyPoolRefillCmd(nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"keypoolrefill","params":[],"id":1}`,
		// 	unmarshalled: &hcashjson.KeyPoolRefillCmd{
		// 		NewSize: hcashjson.Uint(100),
		// 	},
		// },
		// {
		// 	name: "keypoolrefill optional",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("keypoolrefill", 200)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewKeyPoolRefillCmd(hcashjson.Uint(200))
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"keypoolrefill","params":[200],"id":1}`,
		// 	unmarshalled: &hcashjson.KeyPoolRefillCmd{
		// 		NewSize: hcashjson.Uint(200),
		// 	},
		// },
		// {
		// 	name: "listaccounts",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listaccounts")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListAccountsCmd(nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listaccounts","params":[],"id":1}`,
		// 	unmarshalled: &hcashjson.ListAccountsCmd{
		// 		MinConf: hcashjson.Int(1),
		// 	},
		// },
		// {
		// 	name: "listaccounts optional",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listaccounts", 6)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListAccountsCmd(hcashjson.Int(6))
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listaccounts","params":[6],"id":1}`,
		// 	unmarshalled: &hcashjson.ListAccountsCmd{
		// 		MinConf: hcashjson.Int(6),
		// 	},
		// },
		// {
		// 	name: "listlockunspent",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listlockunspent")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListLockUnspentCmd()
		// 	},
		// 	marshalled:   `{"jsonrpc":"1.0","method":"listlockunspent","params":[],"id":1}`,
		// 	unmarshalled: &hcashjson.ListLockUnspentCmd{},
		// },
		// {
		// 	name: "listreceivedbyaccount",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listreceivedbyaccount")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListReceivedByAccountCmd(nil, nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaccount","params":[],"id":1}`,
		// 	unmarshalled: &hcashjson.ListReceivedByAccountCmd{
		// 		MinConf:          hcashjson.Int(1),
		// 		IncludeEmpty:     hcashjson.Bool(false),
		// 		IncludeWatchOnly: hcashjson.Bool(false),
		// 	},
		// },
		// {
		// 	name: "listreceivedbyaccount optional1",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listreceivedbyaccount", 6)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListReceivedByAccountCmd(hcashjson.Int(6), nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaccount","params":[6],"id":1}`,
		// 	unmarshalled: &hcashjson.ListReceivedByAccountCmd{
		// 		MinConf:          hcashjson.Int(6),
		// 		IncludeEmpty:     hcashjson.Bool(false),
		// 		IncludeWatchOnly: hcashjson.Bool(false),
		// 	},
		// },
		// {
		// 	name: "listreceivedbyaccount optional2",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listreceivedbyaccount", 6, true)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListReceivedByAccountCmd(hcashjson.Int(6), hcashjson.Bool(true), nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaccount","params":[6,true],"id":1}`,
		// 	unmarshalled: &hcashjson.ListReceivedByAccountCmd{
		// 		MinConf:          hcashjson.Int(6),
		// 		IncludeEmpty:     hcashjson.Bool(true),
		// 		IncludeWatchOnly: hcashjson.Bool(false),
		// 	},
		// },
		// {
		// 	name: "listreceivedbyaccount optional3",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listreceivedbyaccount", 6, true, false)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListReceivedByAccountCmd(hcashjson.Int(6), hcashjson.Bool(true), hcashjson.Bool(false))
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaccount","params":[6,true,false],"id":1}`,
		// 	unmarshalled: &hcashjson.ListReceivedByAccountCmd{
		// 		MinConf:          hcashjson.Int(6),
		// 		IncludeEmpty:     hcashjson.Bool(true),
		// 		IncludeWatchOnly: hcashjson.Bool(false),
		// 	},
		// },
		// {
		// 	name: "listreceivedbyaddress",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listreceivedbyaddress")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListReceivedByAddressCmd(nil, nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaddress","params":[],"id":1}`,
		// 	unmarshalled: &hcashjson.ListReceivedByAddressCmd{
		// 		MinConf:          hcashjson.Int(1),
		// 		IncludeEmpty:     hcashjson.Bool(false),
		// 		IncludeWatchOnly: hcashjson.Bool(false),
		// 	},
		// },
		// {
		// 	name: "listreceivedbyaddress optional1",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listreceivedbyaddress", 6)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListReceivedByAddressCmd(hcashjson.Int(6), nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaddress","params":[6],"id":1}`,
		// 	unmarshalled: &hcashjson.ListReceivedByAddressCmd{
		// 		MinConf:          hcashjson.Int(6),
		// 		IncludeEmpty:     hcashjson.Bool(false),
		// 		IncludeWatchOnly: hcashjson.Bool(false),
		// 	},
		// },
		// {
		// 	name: "listreceivedbyaddress optional2",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listreceivedbyaddress", 6, true)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListReceivedByAddressCmd(hcashjson.Int(6), hcashjson.Bool(true), nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaddress","params":[6,true],"id":1}`,
		// 	unmarshalled: &hcashjson.ListReceivedByAddressCmd{
		// 		MinConf:          hcashjson.Int(6),
		// 		IncludeEmpty:     hcashjson.Bool(true),
		// 		IncludeWatchOnly: hcashjson.Bool(false),
		// 	},
		// },
		// {
		// 	name: "listreceivedbyaddress optional3",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listreceivedbyaddress", 6, true, false)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListReceivedByAddressCmd(hcashjson.Int(6), hcashjson.Bool(true), hcashjson.Bool(false))
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaddress","params":[6,true,false],"id":1}`,
		// 	unmarshalled: &hcashjson.ListReceivedByAddressCmd{
		// 		MinConf:          hcashjson.Int(6),
		// 		IncludeEmpty:     hcashjson.Bool(true),
		// 		IncludeWatchOnly: hcashjson.Bool(false),
		// 	},
		// },
		// {
		// 	name: "listsinceblock",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listsinceblock")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListSinceBlockCmd(nil, nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listsinceblock","params":[],"id":1}`,
		// 	unmarshalled: &hcashjson.ListSinceBlockCmd{
		// 		BlockHash:           nil,
		// 		TargetConfirmations: hcashjson.Int(1),
		// 		IncludeWatchOnly:    hcashjson.Bool(false),
		// 	},
		// },
		// {
		// 	name: "listsinceblock optional1",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listsinceblock", "123")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListSinceBlockCmd(hcashjson.String("123"), nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listsinceblock","params":["123"],"id":1}`,
		// 	unmarshalled: &hcashjson.ListSinceBlockCmd{
		// 		BlockHash:           hcashjson.String("123"),
		// 		TargetConfirmations: hcashjson.Int(1),
		// 		IncludeWatchOnly:    hcashjson.Bool(false),
		// 	},
		// },
		// {
		// 	name: "listsinceblock optional2",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listsinceblock", "123", 6)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListSinceBlockCmd(hcashjson.String("123"), hcashjson.Int(6), nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listsinceblock","params":["123",6],"id":1}`,
		// 	unmarshalled: &hcashjson.ListSinceBlockCmd{
		// 		BlockHash:           hcashjson.String("123"),
		// 		TargetConfirmations: hcashjson.Int(6),
		// 		IncludeWatchOnly:    hcashjson.Bool(false),
		// 	},
		// },
		// {
		// 	name: "listsinceblock optional3",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listsinceblock", "123", 6, true)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListSinceBlockCmd(hcashjson.String("123"), hcashjson.Int(6), hcashjson.Bool(true))
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listsinceblock","params":["123",6,true],"id":1}`,
		// 	unmarshalled: &hcashjson.ListSinceBlockCmd{
		// 		BlockHash:           hcashjson.String("123"),
		// 		TargetConfirmations: hcashjson.Int(6),
		// 		IncludeWatchOnly:    hcashjson.Bool(true),
		// 	},
		// },
		// {
		// 	name: "listtransactions",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listtransactions")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListTransactionsCmd(nil, nil, nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":[],"id":1}`,
		// 	unmarshalled: &hcashjson.ListTransactionsCmd{
		// 		Account:          nil,
		// 		Count:            hcashjson.Int(10),
		// 		From:             hcashjson.Int(0),
		// 		IncludeWatchOnly: hcashjson.Bool(false),
		// 	},
		// },
		// {
		// 	name: "listtransactions optional1",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listtransactions", "acct")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListTransactionsCmd(hcashjson.String("acct"), nil, nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":["acct"],"id":1}`,
		// 	unmarshalled: &hcashjson.ListTransactionsCmd{
		// 		Account:          hcashjson.String("acct"),
		// 		Count:            hcashjson.Int(10),
		// 		From:             hcashjson.Int(0),
		// 		IncludeWatchOnly: hcashjson.Bool(false),
		// 	},
		// },
		// {
		// 	name: "listtransactions optional2",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listtransactions", "acct", 20)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListTransactionsCmd(hcashjson.String("acct"), hcashjson.Int(20), nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":["acct",20],"id":1}`,
		// 	unmarshalled: &hcashjson.ListTransactionsCmd{
		// 		Account:          hcashjson.String("acct"),
		// 		Count:            hcashjson.Int(20),
		// 		From:             hcashjson.Int(0),
		// 		IncludeWatchOnly: hcashjson.Bool(false),
		// 	},
		// },
		// {
		// 	name: "listtransactions optional3",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listtransactions", "acct", 20, 1)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListTransactionsCmd(hcashjson.String("acct"), hcashjson.Int(20),
		// 			hcashjson.Int(1), nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":["acct",20,1],"id":1}`,
		// 	unmarshalled: &hcashjson.ListTransactionsCmd{
		// 		Account:          hcashjson.String("acct"),
		// 		Count:            hcashjson.Int(20),
		// 		From:             hcashjson.Int(1),
		// 		IncludeWatchOnly: hcashjson.Bool(false),
		// 	},
		// },
		// {
		// 	name: "listtransactions optional4",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listtransactions", "acct", 20, 1, true)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListTransactionsCmd(hcashjson.String("acct"), hcashjson.Int(20),
		// 			hcashjson.Int(1), hcashjson.Bool(true))
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":["acct",20,1,true],"id":1}`,
		// 	unmarshalled: &hcashjson.ListTransactionsCmd{
		// 		Account:          hcashjson.String("acct"),
		// 		Count:            hcashjson.Int(20),
		// 		From:             hcashjson.Int(1),
		// 		IncludeWatchOnly: hcashjson.Bool(true),
		// 	},
		// },
		// {
		// 	name: "listunspent",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listunspent")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListUnspentCmd(nil, nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listunspent","params":[],"id":1}`,
		// 	unmarshalled: &hcashjson.ListUnspentCmd{
		// 		MinConf:   hcashjson.Int(1),
		// 		MaxConf:   hcashjson.Int(9999999),
		// 		Addresses: nil,
		// 	},
		// },
		// {
		// 	name: "listunspent optional1",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listunspent", 6)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListUnspentCmd(hcashjson.Int(6), nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listunspent","params":[6],"id":1}`,
		// 	unmarshalled: &hcashjson.ListUnspentCmd{
		// 		MinConf:   hcashjson.Int(6),
		// 		MaxConf:   hcashjson.Int(9999999),
		// 		Addresses: nil,
		// 	},
		// },
		// {
		// 	name: "listunspent optional2",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listunspent", 6, 100)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListUnspentCmd(hcashjson.Int(6), hcashjson.Int(100), nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listunspent","params":[6,100],"id":1}`,
		// 	unmarshalled: &hcashjson.ListUnspentCmd{
		// 		MinConf:   hcashjson.Int(6),
		// 		MaxConf:   hcashjson.Int(100),
		// 		Addresses: nil,
		// 	},
		// },
		// {
		// 	name: "listunspent optional3",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("listunspent", 6, 100, []string{"1Address", "1Address2"})
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewListUnspentCmd(hcashjson.Int(6), hcashjson.Int(100),
		// 			&[]string{"1Address", "1Address2"})
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"listunspent","params":[6,100,["1Address","1Address2"]],"id":1}`,
		// 	unmarshalled: &hcashjson.ListUnspentCmd{
		// 		MinConf:   hcashjson.Int(6),
		// 		MaxConf:   hcashjson.Int(100),
		// 		Addresses: &[]string{"1Address", "1Address2"},
		// 	},
		// },
		// {
		// 	name: "lockunspent",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("lockunspent", true, `[{"txid":"123","vout":1}]`)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		txInputs := []hcashjson.TransactionInput{
		// 			{Txid: "123", Vout: 1},
		// 		}
		// 		return hcashjson.NewLockUnspentCmd(true, txInputs)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"lockunspent","params":[true,[{"txid":"123","vout":1,"tree":0}]],"id":1}`,
		// 	unmarshalled: &hcashjson.LockUnspentCmd{
		// 		Unlock: true,
		// 		Transactions: []hcashjson.TransactionInput{
		// 			{Txid: "123", Vout: 1},
		// 		},
		// 	},
		// },
		// {
		// 	name: "sendfrom",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("sendfrom", "from", "1Address", 0.5)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		// revised by sammy at 2017-10-27
		// 		//return hcashjson.NewSendFromCmd("from", "1Address", 0.5, nil, nil, nil)
		// 		return hcashjson.NewSendFromCmd("from", "1Address", 0.5, nil, nil, nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"sendfrom","params":["from","1Address",0.5],"id":1}`,
		// 	unmarshalled: &hcashjson.SendFromCmd{
		// 		FromAccount: "from",
		// 		ToAddress:   "1Address",
		// 		Amount:      0.5,
		// 		MinConf:     hcashjson.Int(1),
		// 		Comment:     nil,
		// 		CommentTo:   nil,
		// 	},
		// },
		// {
		// 	name: "sendfrom optional1",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("sendfrom", "from", "1Address", 0.5, 6)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		// revised by sammy at 2017-10-27
		// 		//return hcashjson.NewSendFromCmd("from", "1Address", 0.5, hcashjson.Int(6), nil, nil)
		// 		return hcashjson.NewSendFromCmd("from", "1Address", 0.5, nil, hcashjson.Int(6), nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"sendfrom","params":["from","1Address",0.5,6],"id":1}`,
		// 	unmarshalled: &hcashjson.SendFromCmd{
		// 		FromAccount: "from",
		// 		ToAddress:   "1Address",
		// 		Amount:      0.5,
		// 		MinConf:     hcashjson.Int(6),
		// 		Comment:     nil,
		// 		CommentTo:   nil,
		// 	},
		// },
		// {
		// 	name: "sendfrom optional2",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("sendfrom", "from", "1Address", 0.5, 6, "comment")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		// revised by sammy at 2017-10-27
		// 		//return hcashjson.NewSendFromCmd("from", "1Address", 0.5, hcashjson.Int(6),
		// 		return hcashjson.NewSendFromCmd("from", "1Address", 0.5, nil, hcashjson.Int(6),
		// 			hcashjson.String("comment"), nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"sendfrom","params":["from","1Address",0.5,6,"comment"],"id":1}`,
		// 	unmarshalled: &hcashjson.SendFromCmd{
		// 		FromAccount: "from",
		// 		ToAddress:   "1Address",
		// 		Amount:      0.5,
		// 		MinConf:     hcashjson.Int(6),
		// 		Comment:     hcashjson.String("comment"),
		// 		CommentTo:   nil,
		// 	},
		// },
		// {
		// 	name: "sendfrom optional3",
		// 	newCmd: func() (interface{}, error) {
		// 		// revised by sammy at 2017-10-27
		// 		//return hcashjson.NewCmd("sendfrom", "from", "1Address", 0.5, 6, "comment", "commentto")
		// 		return hcashjson.NewCmd("sendfrom", "from", "1Address", 0.5, nil, 6, "comment", "commentto")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		// revised by sammy at 2017-10-27
		// 		return hcashjson.NewSendFromCmd("from", "1Address", 0.5, hcashjson.Int(6), hcashjson.Int(6),
		// 			hcashjson.String("comment"), hcashjson.String("commentto"))
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"sendfrom","params":["from","1Address",0.5,6,"comment","commentto"],"id":1}`,
		// 	unmarshalled: &hcashjson.SendFromCmd{
		// 		FromAccount: "from",
		// 		ToAddress:   "1Address",
		// 		Amount:      0.5,
		// 		MinConf:     hcashjson.Int(6),
		// 		Comment:     hcashjson.String("comment"),
		// 		CommentTo:   hcashjson.String("commentto"),
		// 	},
		// },
		// {
		// 	name: "sendmany",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("sendmany", "from", `{"1Address":0.5}`)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		amounts := map[string]float64{"1Address": 0.5}
		// 		return hcashjson.NewSendManyCmd("from", amounts, nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"sendmany","params":["from",{"1Address":0.5}],"id":1}`,
		// 	unmarshalled: &hcashjson.SendManyCmd{
		// 		FromAccount: "from",
		// 		Amounts:     map[string]float64{"1Address": 0.5},
		// 		MinConf:     hcashjson.Int(1),
		// 		Comment:     nil,
		// 	},
		// },
		// {
		// 	name: "sendmany optional1",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("sendmany", "from", `{"1Address":0.5}`, 6)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		amounts := map[string]float64{"1Address": 0.5}
		// 		return hcashjson.NewSendManyCmd("from", amounts, hcashjson.Int(6), nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"sendmany","params":["from",{"1Address":0.5},6],"id":1}`,
		// 	unmarshalled: &hcashjson.SendManyCmd{
		// 		FromAccount: "from",
		// 		Amounts:     map[string]float64{"1Address": 0.5},
		// 		MinConf:     hcashjson.Int(6),
		// 		Comment:     nil,
		// 	},
		// },
		// {
		// 	name: "sendmany optional2",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("sendmany", "from", `{"1Address":0.5}`, 6, "comment")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		amounts := map[string]float64{"1Address": 0.5}
		// 		return hcashjson.NewSendManyCmd("from", amounts, hcashjson.Int(6), hcashjson.String("comment"))
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"sendmany","params":["from",{"1Address":0.5},6,"comment"],"id":1}`,
		// 	unmarshalled: &hcashjson.SendManyCmd{
		// 		FromAccount: "from",
		// 		Amounts:     map[string]float64{"1Address": 0.5},
		// 		MinConf:     hcashjson.Int(6),
		// 		Comment:     hcashjson.String("comment"),
		// 	},
		// },
		// {
		// 	name: "sendtoaddress",
		// 	newCmd: func() (interface{}, error) {
		// 		// revised by sammy 2017-10-27
		// 		//return hcashjson.NewCmd("sendtoaddress", "1Address", 0.5)
		// 		return hcashjson.NewCmd("sendtoaddress", "1Address", 0.5, 0)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		// revised by sammy 2017-10-27
		// 		//return hcashjson.NewSendToAddressCmd("1Address", 0.5, nil, nil)
		// 		return hcashjson.NewSendToAddressCmd("1Address", 0.5, 0, nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"sendtoaddress","params":["1Address",0.5],"id":1}`,
		// 	unmarshalled: &hcashjson.SendToAddressCmd{
		// 		Address:   "1Address",
		// 		Amount:    0.5,
		// 		Comment:   nil,
		// 		CommentTo: nil,
		// 	},
		// },
		// {
		// 	name: "sendtoaddress optional1",
		// 	newCmd: func() (interface{}, error) {
		// 		// revised by sammy at 2017-10-27
		// 		//return hcashjson.NewCmd("sendtoaddress", "1Address", 0.5, "comment", "commentto")
		// 		return hcashjson.NewCmd("sendtoaddress", "1Address", 0.5, 0, "comment", "commentto")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		// revised by sammy at 2017-10-27
		// 		//return hcashjson.NewSendToAddressCmd("1Address", 0.5, hcashjson.String("comment"),
		// 		return hcashjson.NewSendToAddressCmd("1Address", 0.5, 0, hcashjson.String("comment"),
		// 			hcashjson.String("commentto"))
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"sendtoaddress","params":["1Address",0.5,"comment","commentto"],"id":1}`,
		// 	unmarshalled: &hcashjson.SendToAddressCmd{
		// 		Address:   "1Address",
		// 		Amount:    0.5,
		// 		Comment:   hcashjson.String("comment"),
		// 		CommentTo: hcashjson.String("commentto"),
		// 	},
		// },
		// {
		// 	name: "settxfee",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("settxfee", 0.0001)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewSetTxFeeCmd(0.0001)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"settxfee","params":[0.0001],"id":1}`,
		// 	unmarshalled: &hcashjson.SetTxFeeCmd{
		// 		Amount: 0.0001,
		// 	},
		// },
		// {
		// 	name: "signmessage",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("signmessage", "1Address", "message")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewSignMessageCmd("1Address", "message")
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"signmessage","params":["1Address","message"],"id":1}`,
		// 	unmarshalled: &hcashjson.SignMessageCmd{
		// 		Address: "1Address",
		// 		Message: "message",
		// 	},
		// },
		// {
		// 	name: "signrawtransaction",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("signrawtransaction", "001122")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewSignRawTransactionCmd("001122", nil, nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"signrawtransaction","params":["001122"],"id":1}`,
		// 	unmarshalled: &hcashjson.SignRawTransactionCmd{
		// 		RawTx:    "001122",
		// 		Inputs:   nil,
		// 		PrivKeys: nil,
		// 		Flags:    hcashjson.String("ALL"),
		// 	},
		// },
		// {
		// 	name: "signrawtransaction optional1",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("signrawtransaction", "001122", `[{"txid":"123","vout":1,"tree":0,"scriptPubKey":"00","redeemScript":"01"}]`)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		txInputs := []hcashjson.RawTxInput{
		// 			{
		// 				Txid:         "123",
		// 				Vout:         1,
		// 				ScriptPubKey: "00",
		// 				RedeemScript: "01",
		// 			},
		// 		}

		// 		return hcashjson.NewSignRawTransactionCmd("001122", &txInputs, nil, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"signrawtransaction","params":["001122",[{"txid":"123","vout":1,"tree":0,"scriptPubKey":"00","redeemScript":"01"}]],"id":1}`,
		// 	unmarshalled: &hcashjson.SignRawTransactionCmd{
		// 		RawTx: "001122",
		// 		Inputs: &[]hcashjson.RawTxInput{
		// 			{
		// 				Txid:         "123",
		// 				Vout:         1,
		// 				ScriptPubKey: "00",
		// 				RedeemScript: "01",
		// 			},
		// 		},
		// 		PrivKeys: nil,
		// 		Flags:    hcashjson.String("ALL"),
		// 	},
		// },
		// {
		// 	name: "signrawtransaction optional2",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("signrawtransaction", "001122", `[]`, `["abc"]`)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		txInputs := []hcashjson.RawTxInput{}
		// 		privKeys := []string{"abc"}
		// 		return hcashjson.NewSignRawTransactionCmd("001122", &txInputs, &privKeys, nil)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"signrawtransaction","params":["001122",[],["abc"]],"id":1}`,
		// 	unmarshalled: &hcashjson.SignRawTransactionCmd{
		// 		RawTx:    "001122",
		// 		Inputs:   &[]hcashjson.RawTxInput{},
		// 		PrivKeys: &[]string{"abc"},
		// 		Flags:    hcashjson.String("ALL"),
		// 	},
		// },
		// {
		// 	name: "signrawtransaction optional3",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("signrawtransaction", "001122", `[]`, `[]`, "ALL")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		txInputs := []hcashjson.RawTxInput{}
		// 		privKeys := []string{}
		// 		return hcashjson.NewSignRawTransactionCmd("001122", &txInputs, &privKeys,
		// 			hcashjson.String("ALL"))
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"signrawtransaction","params":["001122",[],[],"ALL"],"id":1}`,
		// 	unmarshalled: &hcashjson.SignRawTransactionCmd{
		// 		RawTx:    "001122",
		// 		Inputs:   &[]hcashjson.RawTxInput{},
		// 		PrivKeys: &[]string{},
		// 		Flags:    hcashjson.String("ALL"),
		// 	},
		// },
		// {
		// 	name: "walletlock",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("walletlock")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewWalletLockCmd()
		// 	},
		// 	marshalled:   `{"jsonrpc":"1.0","method":"walletlock","params":[],"id":1}`,
		// 	unmarshalled: &hcashjson.WalletLockCmd{},
		// },
		// {
		// 	name: "walletpassphrase",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("walletpassphrase", "pass", 60)
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewWalletPassphraseCmd("pass", 60)
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"walletpassphrase","params":["pass",60],"id":1}`,
		// 	unmarshalled: &hcashjson.WalletPassphraseCmd{
		// 		Passphrase: "pass",
		// 		Timeout:    60,
		// 	},
		// },
		// {
		// 	name: "walletpassphrasechange",
		// 	newCmd: func() (interface{}, error) {
		// 		return hcashjson.NewCmd("walletpassphrasechange", "old", "new")
		// 	},
		// 	staticCmd: func() interface{} {
		// 		return hcashjson.NewWalletPassphraseChangeCmd("old", "new")
		// 	},
		// 	marshalled: `{"jsonrpc":"1.0","method":"walletpassphrasechange","params":["old","new"],"id":1}`,
		// 	unmarshalled: &hcashjson.WalletPassphraseChangeCmd{
		// 		OldPassphrase: "old",
		// 		NewPassphrase: "new",
		// 	},
		// },
	}
	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		cmd, err := test.newCmd()
		if err != nil {
			t.Errorf("Test #%d (%s) unexpected NewCmd error: %v ",
				i, test.name, err)
		}
		handleCmd(cmd)
	}

}

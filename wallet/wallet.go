// Copyright (c) 2016 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wallet

import (
	"github.com/HcashOrg/hcashrpcclient"
	"github.com/HcashOrg/hcashutil"
	// "github.com/HcashOrg/hcashd/wire"
	"io/ioutil"
	"log"
	"path/filepath"
	// "time"
	// "github.com/HcashOrg/hcashd/chaincfg/chainhash"
)
func test2() {
	hcashdHomeDir := hcashutil.AppDataDir(".hcashd", false)
	certs, err := ioutil.ReadFile(filepath.Join(hcashdHomeDir, "rpc.cert"))
	if err != nil {
		log.Fatal(err)
	}

	// Create a new RPC client using websockets.  Since this example is
	// not long-lived, the connection will be closed as soon as the program
	// exits.
	connCfg := &hcashrpcclient.ConnConfig{
		Host:         "localhost:12010",
		Endpoint:     "ws",
		User:         "bitcoinrpc",
		Pass:         "123456",
		Certificates: certs,
	}
	client, err := hcashrpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Query the RPC server for the current block count and display it.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)
}

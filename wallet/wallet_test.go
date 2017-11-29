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
	
	"fmt"
	// "io"
	"os"
	
	// "github.com/HcashOrg/hcashrpcclient"
	"github.com/HcashOrg/hcashd/hcashjson"
)

func TestDecodeHex(t *testing.T) {
	test2()
}
func Testtest() {
	// hcashdHomeDir := hcashutil.AppDataDir(".hcashd", false)
	// certs, err := ioutil.ReadFile(filepath.Join(hcashdHomeDir, "rpc.cert"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Create a new RPC client using websockets.  Since this example is
	// not long-lived, the connection will be closed as soon as the program
	// exits.
	// connCfg := &hcashrpcclient.ConnConfig{
	// 	Host:         "localhost:12010",
	// 	Endpoint:     "ws",
	// 	User:         "bitcoinrpc",
	// 	Pass:         "123456",
	// 	Certificates: certs,
	// }

	cmd,_:=hcashjson.NewCmd("getbestblockhash")
	marshalledJSON, err := hcashjson.MarshalCmd(1, cmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}

	// Send the JSON-RPC request to the server using the user-specified
	// connection configuration.
	connCfg,_,_:=loadConfig()
	result, err := sendPostRequest(marshalledJSON, connCfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}
	// strResult := string(result)
	showRet(result)
}
// Copyright (c) 2016 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wallet

import (
	// "bytes"
	// "encoding/hex"
	// "strings"
	"testing"
	// "net"
	// "net/http"
	"log"
	"io/ioutil"
	"github.com/HcashOrg/hcashutil"
	"github.com/HcashOrg/hcashrpcclient"
	"github.com/HcashOrg/hcashd/hcashjson"
)

func TestDecodeHex(t *testing.T) {
	test2()
}
func Testtest() {
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

	cmd:=hcashjson.NewCmd("getbestblockhash")
	marshalledJSON, err := hcashjson.MarshalCmd(1, cmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}

	// Send the JSON-RPC request to the server using the user-specified
	// connection configuration.
	result, err := sendPostRequest(marshalledJSON, connCfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}

	// Choose how to display the result based on its type.
	strResult := string(result)
	if strings.HasPrefix(strResult, "{") || strings.HasPrefix(strResult, "[") {
		var dst bytes.Buffer
		if err := json.Indent(&dst, result, "", "  "); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to format result: %v",
				err)
			os.Exit(1)
		}
		fmt.Println(dst.String())

	} else if strings.HasPrefix(strResult, `"`) {
		var str string
		if err := json.Unmarshal(result, &str); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to unmarshal result: %v",
				err)
			os.Exit(1)
		}
		fmt.Println(str)

	} else if strResult != "null" {
		fmt.Println(strResult)
	}
}
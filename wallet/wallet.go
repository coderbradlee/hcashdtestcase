// Copyright (c) 2016 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wallet

import (	
	// "github.com/HcashOrg/hcashd/wire"
	"io/ioutil"
	"log"
	"path/filepath"
	"fmt"
	// "github.com/HcashOrg/hcashd/chaincfg/chainhash"
	"net"
	"net/http"
	"os"
	"encoding/json"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"strings"
	"github.com/HcashOrg/hcashutil"
	"github.com/HcashOrg/hcashd/hcashjson"
	"github.com/HcashOrg/hcashrpcclient"
	"github.com/btcsuite/go-socks/socks"
)
func handleCmd(cmd interface{}) {
	
	// cmd,_:=hcashjson.NewCmd("getbestblockhash")
	marshalledJSON, err := hcashjson.MarshalCmd(1, cmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}

	// Send the JSON-RPC request to the server using the user-specified
	// connection configuration.
	connCfg,_,err:=loadConfig(true)
	if err!=nil{
		fmt.Println(err)
	}
	// fmt.Println(str)
	// connCfg.Wallet=true
	result, err := sendPostRequest(marshalledJSON, connCfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		// os.Exit(1)
	}
	// strResult := string(result)
	showRet(result)
}
func showRet(result []byte) {
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
	fmt.Println("*************************************")
}
func newHTTPClient(cfg *config) (*http.Client, error) {
	// Configure proxy if needed.
	var dial func(network, addr string) (net.Conn, error)
	if cfg.Proxy != "" {
		proxy := &socks.Proxy{
			Addr:     cfg.Proxy,
			Username: cfg.ProxyUser,
			Password: cfg.ProxyPass,
		}
		dial = func(network, addr string) (net.Conn, error) {
			c, err := proxy.Dial(network, addr)
			if err != nil {
				return nil, err
			}
			return c, nil
		}
	}

	// Configure TLS if needed.
	var tlsConfig *tls.Config
	if !cfg.NoTLS && cfg.RPCCert != "" {
		pem, err := ioutil.ReadFile(cfg.RPCCert)
		if err != nil {
			return nil, err
		}

		pool := x509.NewCertPool()
		if ok := pool.AppendCertsFromPEM(pem); !ok {
			return nil, fmt.Errorf("invalid certificate file: %v",
				cfg.RPCCert)
		}
		tlsConfig = &tls.Config{
			RootCAs:            pool,
			InsecureSkipVerify: cfg.TLSSkipVerify,
		}
	}

	// Create and return the new HTTP client potentially configured with a
	// proxy and TLS.
	client := http.Client{
		Transport: &http.Transport{
			Dial:            dial,
			TLSClientConfig: tlsConfig,
		},
	}
	return &client, nil
}
func sendPostRequest(marshalledJSON []byte, cfg *config) ([]byte, error) {
	// Generate a request to the configured RPC server.
	protocol := "http"
	if !cfg.NoTLS {
		protocol = "https"
	}
	url := protocol + "://" + cfg.RPCServer
	if cfg.PrintJSON {
		fmt.Println(string(marshalledJSON))
	}
	bodyReader := bytes.NewReader(marshalledJSON)
	httpRequest, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		return nil, err
	}
	httpRequest.Close = true
	httpRequest.Header.Set("Content-Type", "application/json")

	// Configure basic access authorization.
	httpRequest.SetBasicAuth(cfg.RPCUser, cfg.RPCPassword)

	// Create the new HTTP client that is configured according to the user-
	// specified options and submit the request.
	httpClient, err := newHTTPClient(cfg)
	if err != nil {
		return nil, err
	}
	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	// Read the raw bytes and close the response.
	respBytes, err := ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		err = fmt.Errorf("error reading json reply: %v", err)
		return nil, err
	}

	// Handle unsuccessful HTTP responses
	if httpResponse.StatusCode < 200 || httpResponse.StatusCode >= 300 {
		// Generate a standard error to return if the server body is
		// empty.  This should not happen very often, but it's better
		// than showing nothing in case the target server has a poor
		// implementation.
		if len(respBytes) == 0 {
			return nil, fmt.Errorf("%d %s", httpResponse.StatusCode,
				http.StatusText(httpResponse.StatusCode))
		}
		return nil, fmt.Errorf("%s", respBytes)
	}

	// If requested, print raw json response.
	if cfg.PrintJSON {
		fmt.Println(string(respBytes))
	}

	// Unmarshal the response.
	var resp hcashjson.Response
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, resp.Error
	}
	return resp.Result, nil
}

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

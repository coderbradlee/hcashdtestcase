package main

import (
	"testing"
// 	"github.com/HcashOrg/hcashrpcclient"
// 	"github.com/HcashOrg/hcashutil"
// 	// "github.com/HcashOrg/hcashd/wire"
// 	"io/ioutil"
// 	// "log"
// 	"path/filepath"
// 	"time"
// 	// "bufio"
// 	"fmt"
// 	"github.com/HcashOrg/hcashd/chaincfg/chainhash"
// 	// "github.com/HcashOrg/hcashwallet/loader"
// 	// "github.com/HcashOrg/hcashd/chaincfg"
)
func BenchmarkCreateNewWallet(b *testing.B) {
        // loader.CreateNewWallet
        for n := 0; n < b.N; n++ {
            testCreateNewWallet()
        }
}
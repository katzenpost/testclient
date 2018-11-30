// main.go - mixnet client
// Copyright (C) 2018  David Anthony Stainton, Masala
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package main provides a mixnet client daemon
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/katzenpost/client"
	"github.com/katzenpost/client/config"
	cclient "github.com/katzenpost/registration_client/client"
	rclient "github.com/katzenpost/registration_client"
	"github.com/katzenpost/core/utils"
)

func main() {
	cfgFile := flag.String("f", "katzenpost.toml", "Path to the config file.")
	genOnly := flag.Bool("g", false, "Generate the keys and exit immediately.")
	register := flag.Bool("r", false, "Register the account.")
	accountName := flag.String("account", "", "Account name to register.")
	dataDir := flag.String("dataDir", "", "Testclient data directory.")

	flag.Parse()

	if *register {
		if len(*accountName) == 0 {
			flag.Usage()
			return
		}

		if _, err := os.Stat(*dataDir); !os.IsNotExist(err) {
			panic(fmt.Sprintf("aborting registration, %s already exists", *dataDir))
		}
		if err := utils.MkDataDir(*dataDir); err != nil {
			panic(err)
		}

		// 2. generate testclient key material and configuration
		linkKey, identityKey, err := cclient.GenerateConfig(*accountName, ProviderName, ProviderKeyPin, "", "", "", *dataDir, "", "", false, Authorities)
		if err != nil {
			panic(err)
		}

		options := &rclient.Options{Scheme: "http"}
		// 3. perform registration with the mixnet Provider
		c, err := rclient.New(RegistrationAddr, options)
		if err != nil {
			panic(err)
		}
		err = c.RegisterAccountWithIdentityAndLinkKey(*accountName, linkKey, identityKey)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Successfully registered %s@%s\n", *accountName, ProviderName)
		fmt.Printf("testclient -f %s\n", *dataDir+"/testclient.toml")
		return
	}

	flag.Parse()

	// Set the umask to something "paranoid".
	syscall.Umask(0077)

	cfg, err := config.LoadFile(*cfgFile, *genOnly)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config file '%v': %v\n", *cfgFile, err)
		os.Exit(-1)
	}

	// Generate keys and exit
	if *genOnly {
		if err := config.GenerateKeys(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to generate client keys")
			os.Exit(-1)
		}
		os.Exit(0)
	}

	// Setup the signal handling.
	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	// Start up the client.
	client, err := client.New(cfg)
	if err != nil {
		panic(err)
	}
	defer client.Shutdown()

	<-signalCh
	client.Shutdown()
}

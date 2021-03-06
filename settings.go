// settings.go - Mixnet Playground settings.
// Copyright (C) 2018  David Stainton, Masala.
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

package main

import (
	"github.com/katzenpost/core/crypto/ecdh"
	"github.com/katzenpost/core/crypto/eddsa"
	"github.com/katzenpost/authority/voting/server/config"
)

const (
	RegistrationAddr      = "103.104.244.156:30004"
	OnionRegistrationAddr = ""
	ProviderName          = "😼"
	ProviderKeyPin        = "CNMZltUR8Bq810e3tZILkftwH4KAeo+MMjSXpaZQd9I="
)

func b64eddsa(a string) *eddsa.PublicKey {
	k := new(eddsa.PublicKey)
	if err := k.FromString(a); err != nil {
		return nil
	}
	return k
}

func b64ecdh(a string) *ecdh.PublicKey {
	k := new(ecdh.PublicKey)
	if err := k.FromString(a); err != nil {
		return nil
	}
	return k
}

var Authorities []*config.AuthorityPeer = []*config.AuthorityPeer{
	&config.AuthorityPeer{
		Addresses: []string{"103.104.244.148:30000",},
		IdentityPublicKey: b64eddsa("EmUWxb6ocBBXhxlrAKgxVd/6tyIDVK/8pIY/nZrqSDQ="),
		LinkPublicKey: b64ecdh("LGw+ZqpN6KmGErMmliRHRUFwWwhr6d8WZWNMkjHerAQ="),
	},
	&config.AuthorityPeer{
		Addresses: []string{"103.104.244.146:30000",},
		IdentityPublicKey: b64eddsa("vdOAeoRtWKFDw+W4k3sNN1EMT9ZsaHHmuCHOEKSg1aA="),
		LinkPublicKey: b64ecdh("zLzr8HIFHrZzRxIMsjhxhRDCu+UVnlxo8KbtLsOkrH8="),
	},
	&config.AuthorityPeer{
		Addresses: []string{"103.104.244.147:30000",},
		IdentityPublicKey: b64eddsa("bFgvws69dJrc3ACKXN5aCJKLHjkN7D8DA2HDKkhSNIk="),
		LinkPublicKey: b64ecdh("CZQadnvwlhpzJI+foNBru3odteb+tCn1n+sSfJtRc34="),
	},
}

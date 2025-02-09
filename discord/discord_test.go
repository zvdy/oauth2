// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discord

import (
	"testing"

	"golang.org/x/oauth2"
)

func TestDiscordEndpoint(t *testing.T) {
	expected := oauth2.Endpoint{
		AuthURL:  "https://discord.com/oauth2/authorize",
		TokenURL: "https://discord.com/api/oauth2/token",
	}
	if Endpoint.AuthURL != expected.AuthURL {
		t.Errorf("expected AuthURL %q; got %q", expected.AuthURL, Endpoint.AuthURL)
	}
	if Endpoint.TokenURL != expected.TokenURL {
		t.Errorf("expected TokenURL %q; got %q", expected.TokenURL, Endpoint.TokenURL)
	}
}

// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package discord provides constants for using OAuth2 to access Discord.
package discord

import "golang.org/x/oauth2"

// Endpoint is Discord's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://discord.com/oauth2/authorize",
	TokenURL: "https://discord.com/api/oauth2/token",
	// Note: oauth2.Endpoint does not directly support a token revocation URL.
	// The Token Revocation URL (https://discord.com/api/oauth2/token/revoke) can
	// be used with a custom implementation if needed.
}

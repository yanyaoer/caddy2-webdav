package main

import (
	caddycmd "github.com/caddyserver/caddy/cmd"
	_ "github.com/caddyserver/caddy/modules/standard"

	_ "github.com/mholt/caddy-webdav"
)

func main() { caddycmd.Main() }

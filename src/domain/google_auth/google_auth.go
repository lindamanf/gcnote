package google_auth

import "net"

type LocalServerConfig struct {
	Port    int
	Timeout int
	OS      string
}

type RedirectResult struct {
	Code string
	Err  error
}

type Redirect struct {
	Result      chan RedirectResult
	ServerStart chan bool
	ServerStop  chan bool
	Listener    net.Listener
}

type OpenBrowser struct {
	EscapeAnd string
	arg       []string
}

package aozora

import "net/http"

const (
	DefaultHost = "www.aozorahack.net"
)

//Server is informations of PA-API
type Server struct {
	scheme string
	name   string //Aozora API server name
}

//ServerOptFunc is self-referential function for functional options pattern
type ServerOptFunc func(*Server)

//New returns new Server instance
func New(opts ...ServerOptFunc) *Server {
	server := &Server{scheme: "http", name: DefaultHost}
	for _, opt := range opts {
		opt(server)
	}
	return server
}

//WithScheme returns function for setting Marketplace
func WithScheme(scheme string) ServerOptFunc {
	return func(s *Server) {
		if s != nil {
			s.scheme = scheme
		}
	}
}

//WithServerName returns function for setting Marketplace
func WithServerName(host string) ServerOptFunc {
	return func(s *Server) {
		if s != nil {
			s.name = host
		}
	}
}

//CreateClient returns new Client instance
func (s *Server) CreateClient(client *http.Client) *Client {
	svr := s
	if svr == nil {
		svr = New()
	}
	if client != nil {
		return &Client{server: svr, client: client}
	}
	return &Client{server: svr, client: http.DefaultClient}
}

//DefaultClient returns new Client instance with default setting
func DefaultClient() *Client {
	return New().CreateClient(nil)
}

/* Copyright 2019 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package proxy

import (
	"io"
	"log"
	"net"
	"net/http"
	//"time"
    //"strings"
)
// TODO
// 1) Establish a HTTP handler on a selected port
// 2) Parse the HTTP request
// 3) Take care of proxy headers
// 4) Forward request to server
// 5) Receive response from server
// 6) Forward response to client

// type struct for Proxy
type Proxy struct {
	// TODO(greatwhite): what does a proxy need?
	Port uint32
	Intercept bool
	log *log.Logger
}

type Header struct {
  Connection string
  Host string
  User-Agent string
  Accept string
  Accept-Encoding string
  Accept-Charset string
  Cache-Control string
  Keep-Alive string
  Proxy-Authenticate string
  Proxy-Authorization string
  Transfer-Encoding string
  Age int
  Expires int
}

// New returns a new proxy
func New(port uint32) (Proxy, err) {
  return Proxy{
    Port: port,
    Intercept: true,
    log: log.New() // this needs a buffer
  }
}

// Run starts the proxy
func (p Proxy) Run() error {
  listener := p.setup_listener() // Returns a TCP listener // on which port?
  for 1==1 {
    conn, err = listener.Accept()
    if err != nil {
      // do something
    }
    // concurrency made easy
    go handleConnection(conn)
  }
}

// returns an error if the handle connection fails
func handleConnection(conn net.Conn()) (error) {
  req := readRequest(conn)
  logRequest(log)
  req := cleanRequest(conn)
  res := forwardRequest(req)
  logResponse(res)
}

// read the request from a connection
func readRequest(conn net.Conn) (net.Request) {
  buffer := bufio.ReadAll(conn) // InfosecGuruji: ReadAll expects io.Reader ?
  // how do we do something with the buffer to make it into a 
  // net.Request ?
  buffer := makeRequest(buffer)
  return req
}

// log the request
func logRequest(r net.Request) {
  log, err := ioutil.ReadAll(r)
  logString := string(log)
  logger.Print(logString) 
} 

// take care of any proxy headers here
func cleanRequest(r net.Request) net.Request {
  
}

// forward the request: return the response
func forwardRequest(r net.Request) net.Response {
  
}

// log the response
func logResponse(r net.Response) {
  logger.Print(r.toString())
}
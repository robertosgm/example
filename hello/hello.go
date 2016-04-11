/*
Copyright 2014 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/robertosgm/example/stringutil"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, stringutil.Reverse("!selpmaxe oG ,olleH")+"\n")
	fmt.Fprintf(w, "%q"+":"+"%q"+"\n", r.Method, r.Proto)
	fmt.Fprintf(w, "%q"+":"+"%q"+"\n", html.EscapeString(r.URL.Path), r.RemoteAddr)
}

func main() {
	caCert, err := ioutil.ReadFile("./rootCA.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cert, err := tls.LoadX509KeyPair("./server.crt", "./server.key")
	if err != nil {
		log.Fatal(err)
	}
	config := tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caCertPool,
	}

	s := &http.Server{
		Addr:      ":4443",
		TLSConfig: &config,
	}
	http.Handle("/", http.HandlerFunc(barHandler))
	log.Fatal(s.ListenAndServeTLS("", ""))
}

package server

import "golang.org/x/crypto/acme/autocert"

// getCertificateManager allows for the creation of LetsEncrypt TLS Certificates
func getCertificateManager() {
	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		// Set cache directory
		Cache: autocert.DirCache("certs"),
	}
	conf.CertManager = &certManager
}


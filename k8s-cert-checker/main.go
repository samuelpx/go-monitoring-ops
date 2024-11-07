package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"
)

func checkCertExpiry(certPath string) (string, error) {
	certPEM, err := os.ReadFile(certPath)
	if err != nil {
		return "", fmt.Errorf("failed to read certificate: %v", err)
	}

	block, _ := pem.Decode(certPEM)
	if block == nil {
		return "", fmt.Errorf("failed to parse certificate PEM")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse certificate: %v", err)
	}

	expiry := cert.NotAfter
	now := time.Now()
	daysUntilExpiry := int(expiry.Sub(now).Hours() / 24)

	var status string
	if daysUntilExpiry <= 7 {
		status = "major"
	} else if daysUntilExpiry <= 31 {
		status = "minor"
	} else {
		status = "valid"
	}

	return fmt.Sprintf("%s: %s", certPath, status), nil
}

func main() {
	certPaths := []string{
		// Add more certificate paths as needed
        "/home/samuelpx/Documents/Projects/go/go-monitoring-ops/k8s-cert-checker/test-certs/test-cert-5days.pem",
        "/home/samuelpx/Documents/Projects/go/go-monitoring-ops/k8s-cert-checker/test-certs/test-cert-10days.pem",
        "/home/samuelpx/Documents/Projects/go/go-monitoring-ops/k8s-cert-checker/test-certs/test-cert-40days.pem",
        "/home/samuelpx/Documents/Projects/go/go-monitoring-ops/k8s-cert-checker/test-certs/test-cert-expired.pem",
	}

	for _, certPath := range certPaths {
		result, err := checkCertExpiry(certPath)
		if err != nil {
			fmt.Printf("Error checking certificate %s: %v\n", certPath, err)
		} else {
			fmt.Println(result)
		}
	}
}

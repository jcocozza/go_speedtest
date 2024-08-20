package speedtest

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

// create a list of bytes of size Mb
func CreateData(size int) []byte {
	return bytes.Repeat([]byte("A"), size*1024*1024)
}

func Download(url string) error {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	nBytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		return err
	}

	totalTime := time.Since(start).Seconds()
	totalBits := float64(nBytes * 8)         // Convert bytes to bits
	speedMbps := totalBits / (totalTime * 1e6) // Convert to Mbps

	fmt.Printf("Data Transferred: %.2f MB\n", float64(nBytes)/(1024*1024))
	fmt.Printf("Time Taken: %.2f seconds\n", totalTime)
	fmt.Printf("Network Speed: %.2f Mbps\n", speedMbps)
	return nil
}

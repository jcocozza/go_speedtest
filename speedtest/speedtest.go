package speedtest

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os/exec"
	"time"
)

type Ping struct {
	Jitter  float64 `json:"jitter"`
	Latency float64 `json:"latency"`
	Low     float64 `json:"low"`
	High    float64 `json:"high"`
}

type Latency struct {
	Iqm    float64 `json:"iqm"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Jitter float64 `json:"jitter"`
}

type PacketInfo struct {
	Bandwidth int     `json:"bandwidth"`
	Bytes     int     `json:"bytes"`
	Elapsed   int     `json:"elapsed"`
	Latency   Latency `json:"latency"`
}

type Interface struct {
	InternalIP string `json:"internalIp"`
	Name       string `json:"name"`
	MacAddr    string `json:"macAddr"`
	IsVpn      bool   `json:"isVpn"`
	ExternalIP string `json:"externalIp"`
}

type Server struct {
	ID       int    `json:"id"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Country  string `json:"country"`
	IP       string `json:"ip"`
}

type Result struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	Persisted bool   `json:"persisted"`
}

type SpeedTestResult struct {
	Type       string     `json:"type"`
	Timestamp  time.Time  `json:"timestamp"`
	Ping       Ping 	  `json:"ping"`
	Download   PacketInfo `json:"download"`
	Upload 	   PacketInfo `json:"upload"`
	PacketLoss int        `json:"packetLoss"`
	Isp        string     `json:"isp"`
	Interface  Interface  `json:"interface"`
	Server     Server     `json:"server"`
	Result 	   Result     `json:"result"`
}

func SpeedTestJSON() SpeedTestResult {
	slog.Info("Running speedtest...")
	cmd := exec.Command("speedtest", "-f", "json")

	output, err := cmd.CombinedOutput()

	if err != nil {
		slog.Error("Error: " + fmt.Sprint(err))
		panic(err)
	}

	var speedtestResult SpeedTestResult
	err = json.Unmarshal(output, &speedtestResult)
    if err != nil {
        slog.Error("Error unmarshaling JSON: " + fmt.Sprint(err))
        panic(err)
    }

	slog.Info("Speedtest output: " + string(output))
	return speedtestResult
}
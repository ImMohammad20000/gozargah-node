package rest

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"

	"github.com/m03ed/marzban-node-go/common"
	"github.com/m03ed/marzban-node-go/config"
	nodeLogger "github.com/m03ed/marzban-node-go/logger"
	"github.com/m03ed/marzban-node-go/tools"
)

var (
	servicePort         = 8002
	nodeHost            = "127.0.0.1"
	xrayExecutablePath  = "/usr/local/bin/xray"
	xrayAssetsPath      = "/usr/local/share/xray"
	sslCertFile         = "../../certs/ssl_cert.pem"
	sslKeyFile          = "../../certs/ssl_key.pem"
	sslClientCertFile   = "../../certs/ssl_client_cert.pem"
	sslClientKeyFile    = "../../certs/ssl_client_key.pem"
	generatedConfigPath = "../../generated/"
	addr                = fmt.Sprintf("%s:%d", nodeHost, servicePort)
	configPath          = "../../backend/xray/config.json"
)

// httpClient creates a custom HTTP client with TLS configuration
func createHTTPClient(tlsConfig *tls.Config) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
		Timeout: 10 * time.Second,
	}
}

func TestRESTConnection(t *testing.T) {
	config.SetEnv(servicePort, 1000, nodeHost, xrayExecutablePath, xrayAssetsPath,
		sslCertFile, sslKeyFile, sslClientCertFile, "rest", generatedConfigPath, true)

	nodeLogger.SetOutputMode(true)

	certFileExists := tools.FileExists(sslCertFile)
	keyFileExists := tools.FileExists(sslKeyFile)
	if !certFileExists || !keyFileExists {
		if err := tools.RewriteSslFile(sslCertFile, sslKeyFile); err != nil {
			t.Fatal(err)
		}
	}

	clientFileExists := tools.FileExists(sslClientCertFile)
	if !clientFileExists {
		t.Fatal("SSL_CLIENT_CERT_FILE is required.")
	}

	tlsConfig, err := tools.LoadTLSCredentials(sslCertFile, sslKeyFile, sslClientCertFile, false)
	if err != nil {
		t.Fatal(err)
	}

	shutdownFunc, s, err := StartHttpListener(tlsConfig, addr)
	if err != nil {
		t.Fatalf("Failed to start HTTP listener: %v", err)
	}
	defer s.StopService()

	creds, err := tools.LoadTLSCredentials(sslClientCertFile, sslClientKeyFile, sslCertFile, true)
	if err != nil {
		t.Fatal(err)
	}

	client := createHTTPClient(creds)

	configFile, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatal(err)
	}

	user := &common.User{
		Email: "test_user1@example.com",
		Inbounds: []string{
			"VMESS TCP NOTLS",
			"VLESS TCP REALITY",
			"TROJAN TCP NOTLS",
			"Shadowsocks TCP",
			"Shadowsocks UDP",
		},
		Proxies: &common.Proxy{
			Vmess: &common.Vmess{
				Id: uuid.New().String(),
			},
			Vless: &common.Vless{
				Id: uuid.New().String(),
			},
			Trojan: &common.Trojan{
				Password: "try a random string",
			},
			Shadowsocks: &common.Shadowsocks{
				Password: "try a random string",
				Method:   "aes-256-gcm",
			},
		},
	}

	user2 := &common.User{
		Email: "test_user2@example.com",
		Inbounds: []string{
			"VMESS TCP NOTLS",
			"VLESS TCP REALITY",
			"TROJAN TCP NOTLS",
			"Shadowsocks TCP",
			"Shadowsocks UDP",
		},
		Proxies: &common.Proxy{
			Vmess: &common.Vmess{
				Id: uuid.New().String(),
			},
			Vless: &common.Vless{
				Id: uuid.New().String(),
			},
			Trojan: &common.Trojan{
				Password: "try a random string",
			},
			Shadowsocks: &common.Shadowsocks{
				Password: "try a random string",
				Method:   "aes-256-gcm",
			},
		},
	}

	backendStartReq := &common.Backend{
		Type:   common.BackendType_XRAY,
		Config: string(configFile),
		Users:  []*common.User{user, user2},
	}

	jsonBody, _ := proto.Marshal(backendStartReq)

	url := fmt.Sprintf("https://%s", addr)

	resp, err := client.Post(url+"/start", "application/x-protobuf", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed to start backend: %v", err)
	}
	defer resp.Body.Close()

	var baseInfoResp common.BaseInfoResponse
	body, _ := io.ReadAll(resp.Body)
	err = proto.Unmarshal(body, &baseInfoResp)
	if err != nil {
		t.Fatalf("Failed to parse start response: %v", err)
	}

	sessionID := baseInfoResp.SessionId
	if sessionID == "" {
		t.Fatal("No session ID received")
	}

	createAuthenticatedRequest := func(method, endpoint string, body io.Reader) (*http.Request, error) {
		req, err := http.NewRequest(method, url+endpoint, body)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+sessionID)
		if body != nil {
			req.Header.Set("Content-Type", "application/x-protobuf")
		}
		return req, nil
	}

	// Try To Get Outbounds Stats
	outboundsStatsReq, _ := createAuthenticatedRequest("GET", "/stats/outbounds", nil)
	outboundsStatsResp, err := client.Do(outboundsStatsReq)
	if err != nil {
		t.Fatalf("Outbounds stats request failed: %v", err)
	}
	defer outboundsStatsResp.Body.Close()

	var outboundsStats common.StatResponse
	outboundsStatsBody, _ := io.ReadAll(outboundsStatsResp.Body)
	err = proto.Unmarshal(outboundsStatsBody, &outboundsStats)
	if err != nil {
		t.Fatalf("Failed to parse outbounds stats: %v", err)
	}

	for _, stat := range outboundsStats.Stats {
		log.Printf("Outbound Stat - Name: %s, Traffic: %d, Type: %s, Link: %s",
			stat.Name, stat.Value, stat.Type, stat.Link)
	}

	// Try To Get Inbounds Stats
	inboundsStatsReq, _ := createAuthenticatedRequest("GET", "/stats/inbounds", nil)
	inboundsStatsResp, err := client.Do(inboundsStatsReq)
	if err != nil {
		t.Fatalf("Inbounds stats request failed: %v", err)
	}
	defer inboundsStatsResp.Body.Close()

	var inboundsStats common.StatResponse
	inboundsStatsBody, _ := io.ReadAll(inboundsStatsResp.Body)
	err = proto.Unmarshal(inboundsStatsBody, &inboundsStats)
	if err != nil {
		t.Fatalf("Failed to parse inbounds stats: %v", err)
	}

	for _, stat := range inboundsStats.Stats {
		log.Printf("Inbound Stat - Name: %s, Traffic: %d, Type: %s, Link: %s",
			stat.Name, stat.Value, stat.Type, stat.Link)
	}

	// Try To Get Users Stats
	usersStatsReq, _ := createAuthenticatedRequest("GET", "/stats/users", nil)
	usersStatsResp, err := client.Do(usersStatsReq)
	if err != nil {
		t.Fatalf("Users stats request failed: %v", err)
	}
	defer usersStatsResp.Body.Close()

	var usersStats common.StatResponse
	usersStatsBody, _ := io.ReadAll(usersStatsResp.Body)
	err = proto.Unmarshal(usersStatsBody, &usersStats)
	if err != nil {
		t.Fatalf("Failed to parse Users stats: %v", err)
	}

	for _, stat := range inboundsStats.Stats {
		log.Printf("User Stat - Name: %s, Traffic: %d, Type: %s, Link: %s",
			stat.Name, stat.Value, stat.Type, stat.Link)
	}

	// Try To Get Backend Stats
	backendStatsReq, _ := createAuthenticatedRequest("GET", "/stats/backend", nil)
	backendStatsResp, err := client.Do(backendStatsReq)
	if err != nil {
		t.Fatalf("Backend stats request failed: %v", err)
	}
	defer backendStatsResp.Body.Close()

	var backendStats common.BackendStatsResponse
	backendStatsBody, _ := io.ReadAll(backendStatsResp.Body)
	err = proto.Unmarshal(backendStatsBody, &backendStats)
	if err != nil {
		t.Fatalf("Failed to parse backend stats: %v", err)
	}

	for _, stat := range inboundsStats.Stats {
		log.Printf("Users Stat - Name: %s, Traffic: %d, Type: %s, Link: %s",
			stat.Name, stat.Value, stat.Type, stat.Link)
	}

	jsonBody, _ = proto.Marshal(user)

	// Try To Add User
	addUserReq, _ := createAuthenticatedRequest("PUT", "/user/sync", bytes.NewBuffer(jsonBody))
	addUserResp, err := client.Do(addUserReq)
	if err != nil {
		t.Fatalf("Add user request failed: %v", err)
	}
	defer addUserResp.Body.Close()

	logsReq, _ := createAuthenticatedRequest("GET", "/logs", nil)
	logsResp, err := client.Do(logsReq)
	if err != nil {
		t.Fatalf("Logs request failed: %v", err)
	}
	defer logsResp.Body.Close()

	logsBody, _ := io.ReadAll(logsResp.Body)
	var logs common.LogList
	err = proto.Unmarshal(logsBody, &logs)
	if err != nil {
		t.Fatalf("Failed to parse logs: %v", err)
	}

	for _, newLog := range logs.GetLogs() {
		fmt.Println("Log detail:", newLog)
	}

	time.Sleep(2 * time.Second)

	// Try To Get Node Stats
	nodeStatsReq, _ := createAuthenticatedRequest("GET", "/stats/system", nil)
	nodeStatsResp, err := client.Do(nodeStatsReq)
	if err != nil {
		t.Fatalf("Node stats request failed: %v", err)
	}
	defer nodeStatsResp.Body.Close()

	var systemStats common.SystemStatsResponse
	nodeStatsBody, _ := io.ReadAll(nodeStatsResp.Body)
	err = proto.Unmarshal(nodeStatsBody, &systemStats)
	if err != nil {
		t.Fatalf("Failed to parse node stats: %v", err)
	}
	fmt.Printf("System Stats: \nMem Total: %d \nMem Used: %d \nCpu Number: %d \nCpu Usage: %f \nIncoming: %d \nOutgoing: %d \n",
		systemStats.MemTotal, systemStats.MemUsed, systemStats.CpuCores, systemStats.CpuUsage, systemStats.IncomingBandwidthSpeed, systemStats.OutgoingBandwidthSpeed)

	stopReq, _ := createAuthenticatedRequest("PUT", "/stop", nil)
	stopResp, err := client.Do(stopReq)
	if err != nil {
		t.Fatalf("Stop request failed: %v", err)
	}
	defer stopResp.Body.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = shutdownFunc(ctx); err != nil {
		t.Fatalf("Failed to shutdown server: %v", err)
	}
}

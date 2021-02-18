package nutanixkpsclient

import (
	"crypto/tls"
	"errors"
	"net/http"
	"sync"

	cloudmgmt_client "terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client"

	httptransport "github.com/go-openapi/runtime/client"
)

// HostURL - Default Karbon platform services URL
const HostURL string = "karbon.nutanix.com"

// Client -
type Client struct {
	HostURL            string
	CloudmgmtAPIClient *cloudmgmt_client.Sherlock
	mutex              sync.Mutex
	Token              string
}

var _ KPSClient = (*Client)(nil)

// NewClient -
func NewClient(host, apiToken, username, password *string) (*Client, error) {
	var hostURL string
	if host != nil {
		hostURL = *host
	} else {
		hostURL = HostURL
	}

	clientConfig := &cloudmgmt_client.TransportConfig{
		Host:     hostURL,
		BasePath: "",
		Schemes:  []string{"https"},
	}
	cloudMgmtClient := cloudmgmt_client.NewHTTPClientWithConfig(nil, clientConfig)
	transport := httptransport.New(clientConfig.Host, clientConfig.BasePath, clientConfig.Schemes)
	transport.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cloudMgmtClient.SetTransport(transport)

	c := Client{
		CloudmgmtAPIClient: cloudMgmtClient,
		HostURL:            hostURL,
	}

	if (apiToken != nil) && (*apiToken != "") {
		c.Token = *apiToken
	} else if (username != nil) && (password != nil) {
		token, err := c.Login(username, password)
		if err != nil {
			return nil, errors.New(*err.Message)
		}
		c.Token = token
	}

	return &c, nil
}

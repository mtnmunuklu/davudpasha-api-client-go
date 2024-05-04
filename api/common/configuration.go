package common

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

// contextKeys are used to identify the type of value in the context.
type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes an oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKeys takes a string apikey as authentication for the request.
	ContextAPIKeys = contextKey("apiKeys")

	// ContextServerIndex uses a server configuration from the index.
	ContextServerIndex = contextKey("serverIndex")

	// ContexOperationServerIndices uses a server configuration from the index mapping.
	ContexOperationServerIndices = contextKey("serverOperationIndices")

	// ContextServerVariables overrides a server configuration varibles.
	ContextServerVariables = contextKey("serverVariables")

	// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextOperationServerVariables = contextKey("serverOperationVaribles")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth.
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey.
type APIKey struct {
	Key    string
	Prefix string
}

// ServerVariable stores the information about a server variable.
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about server.
type ServerConfiguration struct {
	URL         string
	Description string
	Variables   map[string]ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items.
type ServerConfigurations []ServerConfiguration

// URL formats template on a index using given variables.
func (sc ServerConfigurations) URL(index int, variables map[string]string) (string, error) {
	if index < 0 || len(sc) <= index {
		return "", fmt.Errorf("index %v out of range %v", index, len(sc)-1)
	}

	server := sc[index]
	url := server.URL

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("the variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			url = strings.Replace(url, "{"+name+"}", value, -1)
		} else {
			url = strings.Replace(url, "{"+name+"}", variable.DefaultValue, -1)
		}
	}
	return url, nil
}

// RetryConfiguration stores the configuration of the retry behavior of the api client.
type RetryConfiguration struct {
	EnableRetry       bool
	BackOffMultiplier float64
	BackOffBase       float64
	HTTPRetryTimeout  time.Duration
	MaxRetries        int
}

// Configuration stores the configuration of the API client.
type Configuration struct {
	Host               string            `json:"host,omitempty"`
	Scheme             string            `json:"scheme,omitempty"`
	DefaultHeader      map[string]string `json:"defaultHeader,omitempty"`
	UserAgent          string            `json:"userAgent,omitempty"`
	Debug              bool              `json:"debug,omitempty"`
	Compress           bool              `json:"compress,omitempty"`
	Servers            ServerConfigurations
	OperationServers   map[string]ServerConfigurations
	HTTPClient         *http.Client
	unstableOperations map[string]bool
	RetryConfiguration RetryConfiguration
}

// AddDefaultHeader adds a new HTTP header to the default header in request.
func (c *Configuration) AddDefaultHeader(key, value string) {
	c.DefaultHeader[key] = value
}

// ServerURL returns URL based on server settings.
func (c *Configuration) ServerURL(index int, variables map[string]string) (string, error) {
	return c.Servers.URL(index, variables)
}

// ServerURLWithContext returns a new server URL given an endpoint
func (c *Configuration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	sc, ok := c.OperationServers[endpoint]
	if !ok {
		sc = c.Servers
	}

	if ctx == nil {
		return sc.URL(0, nil)
	}

	index, err := getServerOperationIndex(ctx, endpoint)
	if err != nil {
		return "", err
	}

	variables, err := getServerOperationVariables(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return sc.URL(index, variables)
}

// GetUnstableOperations returns a slice with all unstable operation Ids.
func (c *Configuration) GetUnstableOperations() []string {
	ids := make([]string, len(c.unstableOperations))
	for id := range c.unstableOperations {
		ids = append(ids, id)
	}
	return ids
}

// SetUnstableOperationEnabled sets an unstable operation as enabled (true) or disabled (false).
// This function accepts operation ID as an argument - this is the name of the method on the API class.
// Returns true if the operation is marked as unstable and thus was enabled/disabled, false otherwise.
func (c *Configuration) SetUnstableOperationEnabled(operation string, enabled bool) bool {
	if _, ok := c.unstableOperations[operation]; ok {
		c.unstableOperations[operation] = enabled
		return true
	}
	log.Printf("WARNING: '%s' is not unstable operation, can't enable/disable", operation)
	return false
}

// IsUnstableOperation determines whether an operation is an unstable operation.
// This function accepts opereation ID as an argument - this is the name of the method on the API class.
func (c *Configuration) IsUnstableOperation(operation string) bool {
	_, present := c.unstableOperations[operation]
	return present
}

func (c *Configuration) IsUnstableOperationEnabled(operation string) bool {
	if enabled, present := c.unstableOperations[operation]; present {
		return enabled
	}
	log.Printf("WARNING: '%s' is not an unstable operation, is always enabled", operation)
	return false
}

// SetHTTPClientWithInsecureSkipVerify is a custom function for Configuration structure, it sets the HTTP client with TLS certificate verification disabled.
func (c *Configuration) SetHTTPClientWithInsecureSkipVerify() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c.HTTPClient = &http.Client{Transport: tr}
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     getUserAgent(),
		Debug:         false,
		Compress:      true,
		Servers: ServerConfigurations{
			{
				URL:         "https://{site}/{api_path}",
				Description: "No description provided",
				Variables: map[string]ServerVariable{
					"site": {
						Description:  "The web site for Davudpasha customers.",
						DefaultValue: "davudpasha.com",
					},
					"api_path": {
						Description:  "The path where the api is deploeyed",
						DefaultValue: "api/DpConnection/CallByInterfaceApi",
					},
				},
			},
			{
				URL:         "https://{site}/{api_path}",
				Description: "No description provided",
				Variables: map[string]ServerVariable{
					"site": {
						Description:  "The web site for Davudpasha customers.",
						DefaultValue: "davudpasha.com",
						EnumValues: []string{
							"davudpasha.com",
							"csiem.davudpasha.com",
							"soair.davudpasha.com",
							"vatos.davudpasha.com",
						},
					},
					"api_path": {
						Description:  "The path where the api is deploeyed",
						DefaultValue: "api/DpConnection/CallByInterfaceApi",
					},
				},
			},
			{
				URL:         "{protocol}://{name}",
				Description: "No description provided",
				Variables: map[string]ServerVariable{
					"name": {
						Description:  "Full site DNS name",
						DefaultValue: "davudpasha.com",
					},
					"protocol": {
						Description:  "The protocol  for accessing the API",
						DefaultValue: "https",
					},
				},
			},
		},
		OperationServers:   map[string]ServerConfigurations{},
		unstableOperations: map[string]bool{},
		RetryConfiguration: RetryConfiguration{
			EnableRetry:       false,
			BackOffMultiplier: 2,
			BackOffBase:       2,
			HTTPRetryTimeout:  60 * time.Second,
			MaxRetries:        3,
		},
	}
	return cfg
}

func getServerIndex(ctx context.Context) (int, error) {
	si := ctx.Value(ContextServerIndex)
	if si != nil {
		if index, ok := si.(int); ok {
			return index, nil
		}
		return 0, ReportError("invalid type %T should be int", si)
	}
	return 0, nil
}

func getServerOperationIndex(ctx context.Context, endpoint string) (int, error) {
	osi := ctx.Value(ContexOperationServerIndices)
	if osi != nil {
		operationIndices, ok := osi.(map[string]int)
		if !ok {
			return 0, ReportError("invalid type %T should be map[string]int", osi)
		}
		index, ok := operationIndices[endpoint]
		if ok {
			return index, nil
		}
	}
	return getServerIndex(ctx)
}

func getServerVariables(ctx context.Context) (map[string]string, error) {
	sv := ctx.Value(ContextServerVariables)
	if sv != nil {
		if variables, ok := sv.(map[string]string); ok {
			return variables, nil
		}
		return nil, ReportError("ctx value of ContextServerVariables has invalid type %T should be map[string]string", sv)
	}
	return nil, nil
}

func getServerOperationVariables(ctx context.Context, endpoint string) (map[string]string, error) {
	osv := ctx.Value(ContextOperationServerVariables)
	if osv != nil {
		operationVariables, ok := osv.(map[string]map[string]string)
		if !ok {
			return nil, ReportError("ctx value of ContextOperationServerVariables has invalid type %T should be map[string]map[string]string", osv)
		}
		variables, ok := operationVariables[endpoint]
		if ok {
			return variables, nil
		}
	}
	return getServerVariables(ctx)
}

func getUserAgent() string {
	return fmt.Sprintf(
		"davudpasha-api-client-go (go %s; os %s; arch %s)",
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
	)
}

func NewDefaultContext(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	if site, ok := os.LookupEnv("DP_SITE"); ok {
		ctx = context.WithValue(
			ctx,
			ContextServerVariables,
			map[string]string{"site": site},
		)
	}

	keys := make(map[string]APIKey)
	if apiKey, ok := os.LookupEnv("DP_API_KEY"); ok {
		keys["apiKeyAuth"] = APIKey{Key: apiKey}
	}
	ctx = context.WithValue(
		ctx,
		ContextAPIKeys,
		keys,
	)

	return ctx
}

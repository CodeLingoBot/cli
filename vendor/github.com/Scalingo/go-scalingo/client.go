package scalingo

import (
	"crypto/tls"
	"time"

	"github.com/Scalingo/go-scalingo/http"
)

type API interface {
	AddonsService
	AddonProvidersService
	AppsService
	AlertsService
	AutoscalersService
	BackupsService
	CollaboratorsService
	DeploymentsService
	DomainsService
	VariablesService
	EventsService
	KeysService
	LoginService
	LogsArchivesService
	LogsService
	NotificationPlatformsService
	NotificationsService
	NotifiersService
	OperationsService
	RunsService
	SignUpService
	SourcesService
	TokensService
	UsersService
	http.TokenGenerator

	ScalingoAPI() http.Client
	AuthAPI() http.Client
	DBAPI(app, addon string) http.Client
}

var _ API = (*Client)(nil)

type Client struct {
	config     ClientConfig
	apiClient  http.Client
	dbClient   http.Client
	authClient http.Client
}

type ClientConfig struct {
	Timeout   time.Duration
	TLSConfig *tls.Config
	APIToken  string
}

func NewClient(cfg ClientConfig) *Client {
	client := &Client{
		config: cfg,
	}
	return client
}

func (c *Client) ScalingoAPI() http.Client {
	if c.apiClient != nil {
		return c.apiClient
	}
	var tokenGenerator http.TokenGenerator
	if len(c.config.APIToken) != 0 {
		tokenGenerator = http.NewAPITokenGenerator(c, c.config.APIToken)
	}

	return http.NewClient(http.ScalingoAPI, http.ClientConfig{
		Timeout:        c.config.Timeout,
		TLSConfig:      c.config.TLSConfig,
		APIVersion:     "1",
		TokenGenerator: tokenGenerator,
	})
}

func (c *Client) DBAPI(app, addon string) http.Client {
	if c.dbClient != nil {
		return c.dbClient
	}
	return http.NewClient(http.DBAPI, http.ClientConfig{
		Timeout:        c.config.Timeout,
		TLSConfig:      c.config.TLSConfig,
		TokenGenerator: http.NewAddonTokenGenerator(app, addon, c),
	})
}

func (c *Client) AuthAPI() http.Client {
	if c.authClient != nil {
		return c.authClient
	}
	var tokenGenerator http.TokenGenerator
	if len(c.config.APIToken) != 0 {
		tokenGenerator = http.NewAPITokenGenerator(c, c.config.APIToken)
	}

	return http.NewClient(http.AuthAPI, http.ClientConfig{
		Timeout:        c.config.Timeout,
		TLSConfig:      c.config.TLSConfig,
		APIVersion:     "1",
		TokenGenerator: tokenGenerator,
	})
}

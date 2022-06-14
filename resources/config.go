package resources

import "time"

// AppConfig is main struct for configuration
type AppConfig struct {
	Gateway        string             `json:",omitempty" mapstructure:"gateway"`
	ProjectID      string             `json:",omitempty"`
	Context        SectionContext     `json:",omitempty"`
	HTTPConfig     SectionHTTP        `json:",omitempty" mapstructure:"httpconfig"`
	ConsumerConfig SectionConsumer    `json:",omitempty"`
	Core           SectionCore        `json:",omitempty" mapstructure:"core"`
	Token          SectionToken       `json:",omitempty"`
	ExternalAPI    SectionExternalAPI `json:",omitempty"`
	GINMode        string             `json:",omitempty"`
	Log            SectionLogger      `json:",omitempty"`
	Resilience     Resilience         `json:",omitempty" mapstructure:"resilience"`
	LogLevel       string             `json:",omitempty"`
}

type SectionCore struct {
	Kvothe SectionService `json:",omitempty"`
}

// SectionService is a struct for Service Configuration
type SectionService struct {
	Name           string           `json:",omitempty" mapstructure:"name"`
	Environment    string           `json:",omitempty" mapstructure:"environment"`
	Version        string           `json:",omitempty" mapstructure:"version"`
	Port           string           `json:",omitempty" mapstructure:"port"`
	Protocol       string           `json:",omitempty" mapstructure:"protocol"`
	Hostname       string           `json:",omitempty" mapstructure:"hostname"`
	Gateway        string           `json:",omitempty" mapstructure:"gateway"`
	LogLevel       int              `json:",omitempty" mapstructure:"loglevel"`
	IsEnable       bool             `json:",omitempty"`
	FirstDB        string           `json:",omitempty"`
	GRPC           SectionGRPC      `json:",omitempty"`
	DBPostgres     SectionDB        `json:",omitempty" mapstructure:"dbpostgres"`
	Redis          RedisAccount     `json:",omitempty" mapstructure:"redis"`
	Timeout        SectionTimeout   `json:",omitempty"`
	AWS            SectionAWS       `json:",omitempty"`
	AppKey         string           `json:",omitempty"`
	DomainFrontend string           `json:",omitempty" mapstructure:"domainfrontend"`
	JWT            SectionJWT       `json:",omitempty"`
	ExternalJWT    SectionJWT       `json:",omitempty"`
	BasicAuth      SectionBasicAuth `json:",omitempty"`
	Slack          SectionSlack     `json:",omitempty"`
	MaxProccess    int              `json:",omitempty"`
}

type SectionGRPC struct {
	Port string `json:",omitempty" mapstructure:"port"`
}

type SectionSlack struct {
	WebhookURL        string `json:"webhook_url,omitempty" mapstructure:"webhookurl"`
	WebhookChannel    string `json:"webhook_channel,omitempty" mapstructure:"webhookchannel"`
	IsEnableSlack     bool   `json:"is_enable_slack,omitempty" mapstructure:"isenableslack"`
	VerificationToken string `json:"verification_token,omitempty" mapstructure:"verificationtoken"`
	BotToken          string `json:"bot_token,omitempty" mapstructure:"bottoken"`
	ChannelID         string `json:"channel_id,omitempty" mapstructure:"channelid"`
	Icon              string `json:"icon,omitempty" mapstructure:"icon"`
}

type Resilience struct {
	Retrier Retrier `json:",omitempty" mapstructure:"retrier"`
}

type Retrier struct {
	WaitBase       time.Duration `json:",omitempty" mapstructure:"waitbase"`
	DisableBackoff bool          `json:",omitempty" mapstructure:"disablebackoff"`
	Times          int           `json:",omitempty" mapstructure:"times"`
}

type SectionLogger struct {
	FileName string `json:",omitempty"`
	Mode     string `json:",omitempty"`
}

// SectionJWT is a config for JWT
type SectionJWT struct {
	JWTSecretKey         string `json:",omitempty"`
	JWTRefreshSecretKey  string `json:",omitempty"`
	JWTExpiration        int    `json:",omitempty"`
	JWTRefreshExpiration int    `json:",omitempty"`
}

// SectionHTTP is a config for http request
type SectionHTTP struct {
	Timeout          int  `json:",omitempty" mapstructure:"timeout"`
	DisableKeepAlive bool `json:",omitempty" mapstructure:"disablekeepalive"`
}

type SectionContext struct {
	Timeout int `json:",omitempty"`
}

// SectionConsumer for config consumer
type SectionConsumer struct {
	Timeout int `json:",omitempty"`
}

// SectionToken is a config for Token
type SectionToken struct {
	TokenKey         string           `json:",omitempty"`
	RefreshTokenKey  string           `json:",omitempty"`
	TokenDaySecond   int              `json:",omitempty"`
	TokenMonthSecond int              `json:",omitempty"`
	BasicAuth        SectionBasicAuth `json:",omitempty"`
}

//SectionTimeout for service timeout in second
type SectionTimeout struct {
	Read  int `json:",omitempty"`
	Write int `json:",omitempty"`
}

type SectionBasicAuth struct {
	Username string
	Password string
}

type SectionAWS struct {
	SES SectionAWSItem `json:",omitempty"`
	S3  SectionAWSItem `json:",omitempty"`
}

type SectionAWSItem struct {
	AccessKeyID     string `json:",omitempty"`
	SecretAccessKey string `json:",omitempty"`
	DefaultBucket   string `json:",omitempty"`
	Region          string `json:",omitempty"`
	ScoutPrefix     string `json:",omitempty"`
	URL             string `json:",omitempty"`
	Storage         string `json:",omitempty"`
}

type SectionExternalAPI struct {
	AWS      SectionAWS      `json:",omitempty"`
	Google   SectionGoogle   `json:",omitempty"`
	Xendit   SectionXendit   `json:",omitempty"`
	Midtrans SectionMidtrans `json:",omitempty"`
}

type SectionGoogle struct {
	APIKey string `json:",omitempty"`
}

type SectionXendit struct {
	PublicKey string `json:",omitempty"`
	APIKey    string `json:",omitempty"`
}

type SectionMidtrans struct {
	ClientKey    string `json:",omitempty"`
	APIKey       string `json:",omitempty"`
	ServerKey    string `json:",omitempty"`
	MerchantCode string `json:",omitempty"`
}

type RedisAccount struct {
	URL                  string   `json:",omitempty" mapstructure:"url"`
	Port                 int      `json:",omitempty" mapstructure:"port"`
	DB                   int      `json:",omitempty" mapstructure:"db"`
	Password             string   `json:",omitempty" mapstructure:"password"`
	PoolSize             int      `json:",omitempty" mapstructure:"poolsize"`
	MinIdleConns         int      `json:",omitempty" mapstructure:"minidleconns"`
	MaxIdle              int      `json:",omitempty" mapstructure:"maxidle"`
	MaxActive            int      `json:",omitempty" mapstructure:"maxactive"`
	MaxConnLifetime      int      `json:",omitempty" mapstructure:"maxconnlifetime"`
	RedisearchIndex      []string `json:",omitempty"`
	RedisMutexExpiryTime int      `json:",omitempty"`
	RedisMutexLockTries  int      `json:",omitempty"`
}

type SectionDB struct {
	READ  DBAccount `json:",omitempty" mapstructure:"read"`
	WRITE DBAccount `json:",omitempty" mapstructure:"write"`
}

// DBAccount is struct for database configuration or database account
type DBAccount struct {
	Username     string `json:",omitempty" mapstructure:"username"`
	Password     string `json:",omitempty" mapstructure:"password"`
	URL          string `json:",omitempty" mapstructure:"url"`
	Port         string `json:",omitempty" mapstructure:"port"`
	DBName       string `json:",omitempty" mapstructure:"dbname"`
	Flavor       string `json:",omitempty" mapstructure:"flavor"`
	MaxIdleConns int    `json:",omitempty" mapstructure:"maxidleconns"`
	MaxOpenConns int    `json:",omitempty" mapstructure:"maxopenconns"`
	MaxLifeTime  int    `json:",omitempty" mapstructure:"maxlifetime"`
	Location     string `json:",omitempty" mapstructure:"location"`
	Timeout      string `json:",omitempty" mapstructure:"timeout"`
}

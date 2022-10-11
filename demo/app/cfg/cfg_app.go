package cfg

type DBCfg struct {
	Addr         string
	Auth         string
	Name         string
	ShowLog      bool
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

type LogCfg struct {
	AppName    string
	Level      string
	IsStdOut   bool
	IsStdErr   bool
	TimeFormat string
	Encoding   string

	IsFileOut   bool
	FileDir     string
	FileName    string
	FileMaxSize int
	FileMaxAge  int
}

type OssCfg struct {
	AK          string
	SK          string
	EpOuter     string
	EpInner     string
	BucketPub   string
	BucketPri   string
	Internal    bool
	TTL         int64
	CallbackUrl string
	FastUpload  bool
}

type CacheCfg struct {
	Addr     string
	Password string
	Database int
}

type SMSConfig struct {
	AccessKey       string
	SecretKey       string
	URL             string
	LimitIPCount    int
	LimitIPTTL      int
	LimitPhoneCount int
	LimitPhoneTTL   int
	TTL             int
}

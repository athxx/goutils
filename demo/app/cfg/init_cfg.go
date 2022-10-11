package cfg

import (
	"path/filepath"

	"github.com/joho/godotenv"

	"app/lib/alioss"
	"app/util"
	"app/util/logx"
)

var (
	Debug     bool
	DebugCurl bool
	PortAPI   string
	Cache     *CacheCfg

	DB                          *DBCfg
	Log                         *LogCfg
	Oss                         *OssCfg
	SMS                         *SMSConfig
	EmailAK, EmailSK, EmailSmtp string
)

func InitCache() {
	Cache = &CacheCfg{}
	Cache.Addr = util.EnvStr("redis_addr", "127.0.0.1:6379")
	Cache.Database = util.EnvInt("redis_db_number", 0)
	Cache.Password = util.EnvStr("redis_password", "")
}

func InitBase(file string) {
	if err := godotenv.Load(file); err != nil {
		logx.Warn(err.Error())
	}

	PortAPI = util.EnvStr("port_api", `:8000`)

	// debug
	Debug = util.EnvBool("debug", false)
	DebugCurl = util.EnvBool("debug_curl", false)

	//init cache
	InitCache()
}

func InitDB() {
	// Database
	DB = &DBCfg{
		Addr:         util.EnvStr("db_addr"),
		Auth:         util.EnvStr("db_auth"),
		Name:         util.EnvStr("db_name"),
		ShowLog:      util.EnvBool("db_show_log", false),
		MaxLifetime:  util.EnvInt("db_max_lifetime", 30),
		MaxOpenConns: util.EnvInt("db_max_open_conns", 100),
		MaxIdleConns: util.EnvInt("db_max_idle_conns", 100),
	}
	//db.InitDB(DB)
}

// InitCache 单例缓存
// func InitCache()error{
// 	// RedisAddr = util.EnvStr(SlmRedisAddr)
// 	InitRedis()
// 	rdx.InitRdb(Redis.Addr, Redis.Password, Redis.Database)
// }

func InitOss() {
	//OSS
	Oss = &OssCfg{
		AK:          util.EnvStr("oss_ak"),
		SK:          util.EnvStr("oss_sk"),
		EpOuter:     util.EnvStr("oss_endpoint_pub"),
		EpInner:     util.EnvStr("oss_endpoint_pri"),
		BucketPub:   util.EnvStr("oss_bucket_pub"),
		BucketPri:   util.EnvStr("oss_bucket_pri"),
		Internal:    util.EnvBool("oss_internal"),
		TTL:         util.EnvInt64("oss_ttl", 1800),
		CallbackUrl: util.EnvStr("oss_callback"),
		FastUpload:  util.EnvBool("oss_fast_upload", false),
	}
	alioss.Init(Oss.AK, Oss.SK, Oss.EpOuter, Oss.EpInner, Oss.BucketPub, Oss.BucketPri, Oss.TTL, Oss.Internal)
}

func InitLog(appName string) {
	// logger
	Log = &LogCfg{
		AppName:    appName,
		Level:      util.EnvStr("log_level", `warn`),
		IsStdOut:   util.EnvBool("log_is_stdout", true),
		IsStdErr:   util.EnvBool("log_is_stderr", false),
		TimeFormat: util.EnvStr("log_time_format", `stand`),
		Encoding:   util.EnvStr("log_encoding", `json`),

		IsFileOut:   util.EnvBool("log_is_file_out", false),
		FileDir:     util.EnvStr("log_file_dir", `./log/`),
		FileName:    util.EnvStr("log_file_name", `app.log`),
		FileMaxAge:  util.EnvInt("log_file_max_age", 1),
		FileMaxSize: util.EnvInt("log_file_max_size", 256),
	}
	logCfg := logx.LogConfig{
		AppName:         Log.AppName,
		Level:           Log.Level,
		StacktraceLevel: "error",
		IsStdOut:        Log.IsStdOut,
		IsStdErr:        Log.IsStdErr,
		TimeFormat:      Log.TimeFormat,
		Encoding:        Log.Encoding,
		IsFileOut:       Log.IsFileOut,
		FileDir:         filepath.Join(Log.FileDir, appName),
		FileName:        appName + ".log", //app使用自己单独的日志文件，便于排错
		FileMaxSize:     Log.FileMaxSize,
		FileMaxAge:      Log.FileMaxAge,
	}
	logx.InitDefaultLogger(&logCfg)
}

func InitSMS() {
	// sms
	SMS = &SMSConfig{}
	SMS.AccessKey = util.EnvStr("sms_ak")
	SMS.SecretKey = util.EnvStr("sms_sk")
	SMS.URL = util.EnvStr("sms_url")
}

func InitEmail() {
	// email
	EmailAK = util.EnvStr("email_ak")
	EmailSK = util.EnvStr("email_sk")
	EmailSmtp = util.EnvStr("email_smtp")
}

package config

import (
	"flag"
	"github.com/spf13/viper"
)

const (
	EnvProd = "prod"
)

var Environment string
var ConfigPath string
var EnvSlice []string

var App app

type app struct {
	Debug          bool      `json:"debug"`
	HttpPort       int       `json:"httpPort"`  //端口
	HttpsPort      int       `json:"httpsPort"` //端口
	RpcPort        int       `json:"rpcPort"`
	Name           string    `json:"name"`
	IP             string    `json:"ip"`
	Databases      databases `json:"databases"`
	Logger         logger    `json:"logger"`
	Redis          redis     `json:"redis"`
	LiveChatImgURL string    `json:"liveChatImgUrl"` //文件上传地址
	Consul         Consul    `json:"consul"`
}

type Consul struct {
	Host  string `json:"host"`
	Tag   string `json:"tag"`
	Token string `json:"token"`
}

type databases struct {
	Driver string     `json:"driver"`
	List   []database `json:"list"`
}

type database struct {
	Name               string `json:"name"`
	Type               string `json:"type"`
	Host               string `json:"host"`
	Port               int    `json:"port"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	MaxIdleConnections int    `json:"max_idle_connections"`
	MaxOpenConnections int    `json:"max_open_connections"`
}

type logger struct {
	Path      string `json:"path"`
	InfoFile  string `json:"info_file"`
	ErrorFile string `json:"error_file"`
	Format    string `json:"format"`
}

type redis struct {
	Cluster  []string `json:"cluster"`
	Password string   `json:"password"`
	MinIdle  int      `json:"minIdle"`
	PoolSize int      `json:"poolSize"`
}

func init() {
	flag.StringVar(&Environment, "env", `dev`, "运行模式")
	flag.StringVar(&ConfigPath, "config", "./config", "配置目录")
	//环境列表
	EnvSlice = []string{`dev`, `13`, `14`, `im`, `prod`}
}

//加载配置
func LoadConfig() {
	var configFileName = "app"
	var notMatch int
	//根据环境选择配置
	for _, v := range EnvSlice {
		if v == Environment {
			configFileName += "_" + v
			notMatch = 1
			break
		}
	}
	//未匹配上的走dev
	if notMatch == 0 {
		configFileName += "_dev"
	}
	config := viper.New()
	config.SetConfigName(configFileName)
	config.SetConfigType("yaml")
	config.AddConfigPath(ConfigPath)
	readErr := config.ReadInConfig()
	if readErr != nil {
		panic("配置文件读取失败，原因：" + readErr.Error())
	}
	unmarshalErr := config.Unmarshal(&App)
	if unmarshalErr != nil {
		panic("配置文件初始化结构体失败，原因：" + unmarshalErr.Error())
	}
}

func Mode() string {
	switch Environment {
	case EnvProd:
		return "release"
	default:
		return "debug"
	}
}

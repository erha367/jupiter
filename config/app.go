package config

import (
	"flag"
	"github.com/spf13/viper"
	"os"
)

const (
	EnvProd = "prod"
	EnvTest = "test"
	EnvDev  = "dev"
	EnvIm   = "im"
)

//运行环境 prod、test、dev
var Environment string

var App app

type app struct {
	Debug     bool      `json:"debug"`
	HttpPort  int       `json:"httpPort"`  //端口
	HttpsPort int       `json:"httpsPort"` //端口
	Sn        int       `json:"sn"`        //机器编号
	Machines  int       `json:"machines"`  //部署机器数量
	Name      string    `json:"name"`
	Domain    string    `json:"domain"`
	Databases databases `json:"databases"`
	Logger    logger    `json:"logger"`
	Redis     redis     `json:"redis"`
	Kafka     kafka     `json:"kafka"`
	Urls      urls      `json:"urls"`
	CertFile  string    `json:"certFile"`
	KeyFile   string    `json:"keyFile"`
}

type urls struct {
	EoOs   eoOs `json:"eoOs"`
	Course eoOs `json:"course"`
}

type eoOs struct {
	Host          string `json:"host"`
	GetCourseInfo string `json:"getCourseInfo"`
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
}

type redis struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

type kafka struct {
	Brokers string `json:"brokers"`
	Topic   string `json:"topic"`
	GroupId string `json:"groupId"`
}

func init() {
	flag.StringVar(&Environment, "env", EnvDev, "运行模式")
}

//加载配置
func LoadConfig() {
	var configFileName = "app"
	switch Environment {
	case EnvProd:
		configFileName += "_" + EnvProd
	case EnvTest:
		configFileName += "_" + EnvTest
	case EnvIm:
		configFileName += "_" + EnvIm
	default:
		configFileName += "_" + EnvDev
	}
	config := viper.New()
	config.SetConfigName(configFileName)
	config.SetConfigType("yaml")
	/* -- 源代码编译后，必须放在相同的项目目录才可以运行，这里修改一下。
	_, filename, _, _ := runtime.Caller(0)
	config.AddConfigPath(path.Dir(filename))
	*/
	path2, _ := os.Getwd()
	config.AddConfigPath(path2 + `/config`)
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
	case EnvTest:
		return "test"
	case EnvIm:
		return "test"
	default:
		return "debug"
	}
}

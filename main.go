package main

import (
	"content-grpc/core"
	"content-grpc/global"
	"content-grpc/initialize"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"reflect"
)

func main() {

	global.VIPER = core.Viper() 		//初始化viper
	global.LOG = core.Zap()				//初始化zap日志库
	global.DB = initialize.Gorm()   	//初始化GORM连接数据库
	initialize.Redis() 					//初始化redis连接

	global.GRPC = core.InitGrpcServer(global.CONFIG.GRPC.Port) //初始化GRPC连接

	defer global.GRPC.Stop()
}

func test() {
	p := message.NewPrinter(language.Chinese)
	p.Println("hello")
}


func test2() {
	var str  = `{"code":2,"data":{"age":1},"msg":"record not found"}`

	var jsonStr map[string]interface{}


	if err := json.Unmarshal([]byte(string(str)), &jsonStr); err != nil {
		fmt.Println(err)
	}

	fmt.Println(jsonStr["data"].(map[string]interface{})["age"])
}


func test3() {
	viper.AddRemoteProvider("consul", "127.0.0.1:8500", "name")
	viper.SetConfigType("json") // Need to explicitly set this to json
	if err := viper.ReadRemoteConfig(); err != nil {
		fmt.Println(err, 8888)
	}

	fmt.Println(viper.Get("port")) // 8080
	fmt.Println(viper.Get("hostname")) // myhostname.com


	var name struct{
		Port string `json:"port"`
		Hostname string `json:"hostname"`
	}


	if err := viper.Unmarshal(&name); err != nil {
		fmt.Println(err, 99999)
	}

	fmt.Println(name.Port)

}


func test4() {
	var str  = `{"code":2,"data":{"age":1},"msg":"record not found","dbname":"test"}`

	var config struct{
		Dbname string `json:"dbname" yaml:"db-name"`
		Code int `mapstructure:"code" json:"code" yaml:"code"`
	}

	if err := json.Unmarshal([]byte(string(str)), &config); err != nil {
		fmt.Println(err)
	}

	fmt.Println(config.Dbname, config.Code)
}

func test5() {

	type People struct {
		Name string
	}

	lei := &People{
		Name: "章磊",
	}

	ref := reflect.TypeOf(lei)

	fmt.Println(ref.Name(), ref.String(), ref.Kind(), reflect.Ptr, reflect.TypeOf(People{}).Kind().String())

	if a, ok := interface{}(lei).(*People); ok {
		fmt.Println("是他的实类", a, ok)
	}

}

func test6() {
	v := viper.New()

	v.SetConfigFile("test.yaml")

	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var config struct{
		Dbname string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
		Name string `mapstructure:"name" json:"name" yaml:"name"`
	}

	if err := v.Unmarshal(&config); err != nil {
		fmt.Println(err)
	}

	str, err := json.Marshal(v.AllSettings())
	fmt.Println(config, config.Name, config.Dbname)
	fmt.Println(string(str))


	str, err = json.Marshal(config)

	fmt.Println(string(str))

	err = json.Unmarshal([]byte(string(str)), &config)

	fmt.Println(config, config.Dbname, config.Name)

	str, err = json.Marshal(config)

	fmt.Println(string(str))
}

func test7() {

	var config struct{
		Dbname string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
		Name string `mapstructure:"name" json:"name" yaml:"name"`
	}

	v := viper.New()

	v.AddRemoteProvider("consul", "127.0.0.1:8500", "test")
	v.SetConfigType("json") // Need to explicitly set this to json
	if err := v.ReadRemoteConfig(); err == nil {
		fmt.Println(v.AllSettings())

		if err := v.Unmarshal(&config); err == nil {
			fmt.Println(config)
		}

		if str, err := json.Marshal(v.AllSettings()); err == nil {
			if err = json.Unmarshal([]byte(string(str)), &config); err == nil {
				fmt.Println(config)
			}
		}
	}

}
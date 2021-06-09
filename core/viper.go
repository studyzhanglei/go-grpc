package core

import (
	"content-grpc/global"
	"content-grpc/utils"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func Viper(path ...string) *viper.Viper {
	var config string
	v := viper.New()

	//先走consul 没有再走本地
	v.AddRemoteProvider("consul", "127.0.0.1:8500", "go")
	v.SetConfigType("json") // Need to explicitly set this to json
	if err := v.ReadRemoteConfig(); err == nil {
		if str, err := json.Marshal(v.AllSettings()); err != nil {
			if err = json.Unmarshal([]byte(string(str)), &global.CONFIG); err != nil {
				return v
			}
		}
	}


	v = viper.New()

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
				config = utils.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", utils.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}


	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	//
	//fmt.Println(global.CONFIG.Mysql.Dbname, 666666)
	//str, err := json.Marshal(v.AllSettings())
	//fmt.Println(err, string(str))

	return v
}

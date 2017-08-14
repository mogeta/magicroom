package firebase

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/zabawaba99/firego"
)

var firebase firego.Firebase

type Irdata struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Data   string `json:"data"`
}

var v map[string]Irdata

func init() {
	viper.SetConfigName("config") // 設定ファイル名を拡張子抜きで指定する
	viper.AddConfigPath(".")      // 現在のワーキングディレクトリを探索することもできる
	err := viper.ReadInConfig()   // 設定ファイルを探索して読み取る
	if err != nil {               // 設定ファイルの読み取りエラー対応
		panic(fmt.Errorf("設定ファイル読み込みエラー: %s \n", err))
	}

	firebase := firego.New(viper.GetString("url"), nil)
	firebase.Auth(viper.GetString("token"))
	if err := firebase.Value(&v); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", v)
	fmt.Printf("%v\n", v["tv"])
}

func GetList() map[string]Irdata {
	return v
}

func Get(key string) Irdata {
	return v[key]
}

package openaiezgo

import (
	"fmt"
	"strconv"
	"testing"

	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

var baseurl string
var tokenLimiter int

func TestGPT(t *testing.T) {
	viper.SetDefault("gpttokenmax", 512)
	viper.SetDefault("gpttoken", "0")
	viper.SetDefault("baseurl", openai.DefaultConfig("").BaseURL)
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	gpttoken := viper.Get("gpttoken").(string)
	fmt.Println("gpttoken=" + gpttoken)
	tokenLimiter = viper.GetInt("gpttokenmax")
	fmt.Println("gpttokenmax=" + strconv.Itoa(tokenLimiter))
	baseurl = viper.Get("baseurl").(string)
	fmt.Println("baseurl=" + baseurl)

	cfg := DefaultConfig(gpttoken)
	cfg.BaseURL = baseurl
	NewClientWithConfig(cfg)

	fmt.Println(NewCharacterSet("12345", "你是一个精神病患者，对于我的任何问题都会以一种非常奇怪的超出常理的角度回答。"))
	fmt.Println(NewSpeech("12345", "不用番茄和蛋如何做出番茄炒蛋这道菜呢？"))
}

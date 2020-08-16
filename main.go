package main

import (
	"editimage/comment"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	configfile = "./config/image.yaml"
	textfile   string
	imagefile  string
)

//读取yaml中字体的配置
func readyaml() (conf comment.Config, err error) {
	yamlfile, err := ioutil.ReadFile(configfile)
	if err != nil {
		log.Println("open yaml file err:", err)
		return conf, err
	}
	err = yaml.Unmarshal(yamlfile, &conf)
	if err != nil {
		log.Println("read yaml file err:", err)
		return conf, err
	}
	return conf, nil
}

//获取需要编辑的文字
func readtext(textfile string) (content []string, err error) {
	textbytes, err := ioutil.ReadFile(textfile)
	if err != nil {
		log.Println("read text file error:", err)
		return []string{}, err
	}
	return strings.Split(string(textbytes), "\n"), nil
}

func init() {
	flag.StringVar(&textfile, "txt", "./data/text.txt", "text file")
	flag.StringVar(&imagefile, "img", "./data/image.jpg", "image file")
}

func main() {
	flag.Parse()
	config, err := readyaml()
	if err != nil {
		os.Exit(1)
	}
	content, err := readtext(textfile)
	if err != nil {
		os.Exit(1)
	}
	comment.StartAddText(imagefile, content, config)
}

package main

import (
	"github.com/Unknwon/goconfig"
	"log"
)

func main() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}

	//基本读写操作
	//v, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "key_default")
	v, err := cfg.GetValue("", "key_default") //这是简便写法

	if err != nil {
		log.Fatalf("Couldn't get the key value(%s):%s ", "key_default", v)
	}
	log.Printf("%s > %s:%s", goconfig.DEFAULT_SECTION, "key_default", v)

	isInsert := cfg.SetValue(goconfig.DEFAULT_SECTION, "key_default", "这是新值")
	if err != nil {
		log.Fatalf("Couldn't set the key value(%s):%s ", "key_default", v)
	}
	log.Printf("设置键值%s 为 %v", "key_default", isInsert) //这里inInsert是布尔值  所以我们用%v打印

	//注释读写操作
	comment := cfg.GetSectionComments("super") //获取分区注释    //这个是特色
	log.Printf("分区 %s 的注释是：%s", "super", comment)

	val := cfg.SetSectionComments("super", "# 这是新的super分区注释")
	//如果返回是true  则代表“插入”  如果返回的是false 则代表“重写”
	log.Printf("分区 %s 的注释被修改为：%v", "super", val)

	vInt, err := cfg.Int("must", "int")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "int", err)
	}
	log.Printf("%s > %s:%v", "must", "int", vInt)
}

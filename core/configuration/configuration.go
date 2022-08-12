package configuration

import (
	"flag"
	"log"
	"seed-admin/common"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

var readyTask = new(sync.WaitGroup)

func AddConfig(path ...string) *config.Config {
	filePath := ""
	if len(path) == 0 {
		flag.StringVar(&filePath, "c", "", "您的配置文件.")
		flag.Parse()
		if filePath == "" {
			filePath = "config.toml"
		}
	} else {
		filePath = path[0]
	}
	hookFn := func(event string, c *config.Config) {
		// // 使用set和setData时重写文件(!!!!!!!!!导出的文件没有注释会自动格式化 这里只建议做备份切莫覆盖原配置文件)
		// if event == "set.value" || event == "set.data" {
		// 	buf := new(buffer.Buffer)
		// // 第二个参数是导出格式,有config.JSON | config.Toml | config.Yaml等等等 根据你到导出的格式选用
		// 	_, err := c.DumpTo(buf, config.Toml)
		// 	if err != nil {
		// 		common.LOG.Error(err.Error())
		// 		return
		// 	}
		// 	if err = ioutil.WriteFile("需要导出的文件地址", buf.Bytes(), 0755); err != nil {
		// 		common.LOG.Error(err.Error())
		// 		return
		// 	}
		// }
	}
	c := config.NewWithOptions("appConfig", config.ParseEnv, config.WithHookFunc(hookFn))
	// 添加驱动程序以支持toml内容解析（除了JSON是默认支持，其他的则是按需使用）
	c.AddDriver(toml.Driver)
	if err := c.LoadFiles(filePath); err != nil {
		panic(err.Error())
	}
	// 文件热重载
	watchConfig(filePath)
	return c
}

// 监听配置变化
func watchConfig(path string) {
	readyTask.Add(1)
	go func() {
		watch(path)
	}()
	readyTask.Wait()
}

// 创建监听器
func watch(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer watcher.Close()
	// 开个协程处理出错或通道关闭时的退出问题
	eventsTask := new(sync.WaitGroup)
	eventsTask.Add(1)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					eventsTask.Done()
					return
				}
				// 只有写入时才重新创建数据
				switch event.Op.String() {
				case "WRITE":
					// 重载数据
					if err := common.CONFIG.LoadFiles(path); err != nil {
						eventsTask.Done()
						log.Panicf("重载%s数据出错,err:%s", event.Name, err.Error())
						return
					}
					log.Printf("监听到%s变动", event.Name)
				case "REMOVE":
					eventsTask.Done()
					log.Panicf("重载%s数据出错,err:文件被删除,请不要删除配置文件", event.Name)
					return
				default:
					log.Printf("监听到%s变动 Op->%s\n", event.Name, event.Op.String())
				}
			case err, ok := <-watcher.Errors:
				if ok {
					log.Println(err.Error())
				}
				eventsTask.Done()
				return
			}
		}
	}()
	// 加载文件的监听
	if err = watcher.Add(path); err != nil {
		log.Fatal(err.Error())
	}
	readyTask.Done()
	eventsTask.Wait()
}

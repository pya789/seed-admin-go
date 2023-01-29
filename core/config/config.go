package config

import (
	"flag"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
	"github.com/gookit/goutil/cliutil"
)

// 添加配置文件
func Add(path ...string) *config.Config {
	filePath := ""
	// 判断是否使用函数入参设置配置文件
	if len(path) == 0 {
		// 判断启动时是否使用-c设置配置文件
		flag.StringVar(&filePath, "c", "", "您的配置文件.")
		flag.Parse()
		if filePath == "" {
			// 没有的话使用默认
			filePath = "config.toml"
		}
	} else {
		filePath = path[0]
	}
	// 创建新的配置实例
	c := config.NewWithOptions("appConfig", func(opts *config.Options) {
		opts.ParseEnv = true
		opts.HookFunc = hookFunc
	})
	// 加载驱动
	c.AddDriver(toml.Driver)
	// 加载配置文件
	if err := c.LoadFiles(filePath); err != nil {
		panic(err.Error())
	}
	// 监听配置文件热修改
	watchConfigFiles(c)
	return c
}

// 监听配置文件热修改
func watchConfigFiles(cfg *config.Config) {
	// 开一个新线程防止主线程被卡住
	readyTask := new(sync.WaitGroup)
	readyTask.Add(1)
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			cliutil.Errorln(err.Error())
			return
		}
		defer watcher.Close()
		// 获取加载的配置文件
		files := cfg.LoadedFiles()
		if len(files) == 0 {
			cliutil.Errorln("未读取到配置文件")
			return
		}
		// 处理出错或通道关闭时的退出问题
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
						if err := cfg.ReloadFiles(); err != nil {
							eventsTask.Done()
							cliutil.Errorf("重载%s数据出错,err:%s\n", event.Name, err.Error())
							return
						}
						cliutil.Infof("监听到%s变动\n", event.Name)
					case "REMOVE":
						eventsTask.Done()
						cliutil.Errorf("重载%s数据出错,err:文件被删除,请不要删除配置文件\n", event.Name)
						return
					default:
						cliutil.Infof("监听到%s变动 Op->%s\n", event.Name, event.Op.String())
					}
				case err, ok := <-watcher.Errors:
					if ok {
						cliutil.Errorln(err.Error())
					}
					if err != nil {
						cliutil.Errorln(err.Error())
					}
					eventsTask.Done()
					return
				}
			}
		}()
		// 加载文件的监听
		for _, path := range files {
			if err := watcher.Add(path); err != nil {
				cliutil.Errorln(err.Error())
			}
		}
		// 加载文件监听成功后释放创建监听的线程
		readyTask.Done()
		// 等待事件释放
		eventsTask.Wait()
	}()
	// 等待监听成功
	readyTask.Wait()
}

// 监听配置修改钩子
func hookFunc(event string, c *config.Config) {
	// if event == "set.value" || event == "set.data" {
	// buf := new(buffer.Buffer)
	// // 第二个参数是导出格式,有config.JSON | config.Toml | config.Yaml等等等 根据你到导出的格式选用
	// _, err := c.DumpTo(buf, config.Toml)
	// if err != nil {
	// 	common.LOG.Error(err.Error())
	// 	return
	// }
	// if err = ioutil.WriteFile("需要导出的文件地址", buf.Bytes(), 0755); err != nil {
	// 	common.LOG.Error(err.Error())
	// 	return
	// }
	// }
}

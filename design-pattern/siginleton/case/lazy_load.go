package _case

// 懒汉式加载，线程不安全
var lazyConf *Config

func GetLazyConf() *Config {
	if lazyConf == nil {
		lazyConf = &Config{
			name: "lazy",
		}
	}
	return lazyConf
}

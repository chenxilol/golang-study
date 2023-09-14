package _case

var hungryConf *Config

func init() {
	hungryConf = &Config{
		name: "hungry",
	}
}
func GetHungryConf() *Config {
	return hungryConf
}

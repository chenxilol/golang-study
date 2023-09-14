package _case

type Config struct {
	name string
}

func (c *Config) GetName() string {
	return c.name
}

package services

import "fmt"

type Configuration struct {
	Host string
	Port string
}

func (c *Configuration) URL() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

package queue

import "time"

type Config struct {
	Host      string        `yaml:"host"`
	Port      string        `yaml:"port"`
	User      string        `yaml:"user"`
	Password  string        `yaml:"password"`
	QueueName string        `yaml:"queueName"`
	Timeout   time.Duration `yaml:"timeout"`
}

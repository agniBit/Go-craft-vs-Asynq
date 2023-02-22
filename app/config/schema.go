package config

type Config struct {
	Redis  RedisConfig  `mapstructure:"redis"`
	Worker WorkerConfig `mapstructure:"worker"`
	Job    JobConfig    `mapstructure:"job"`
}

type WorkerConfig struct {
	Concurrency int `mapstructure:"concurrency"`
	RandomDelay int `mapstructure:"random_delay"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	PoolSize int    `mapstructure:"pool_size"`
}

type JobConfig struct {
	QueueCount      int `mapstructure:"queue_count"`
	TaskPerQueue    int `mapstructure:"task_per_queue"`
	PublisherCount  int `mapstructure:"publishers_count"`
	DefaultDeadline int `mapstructure:"default_deadline"`
	ErrorRate       int `mapstructure:"error_rate"`
}

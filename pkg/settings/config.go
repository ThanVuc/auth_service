package settings

/*
	@Author: Sinh
	@Date: 2025/6/1
	@Description: This file create a configuration struct for the application.
	@Note: The configuration struct is used to store the configuration of the application.
	All the configuration is loaded from a YAML file using Viper.
	And it is stored in the global variable `Config` in the `global` package.
*/

type Config struct {
	Server   Server   `mapstructure:"server" json:"server" yaml:"server"`
	Postgres Progres  `mapstructure:"postgres" json:"postgres" yaml:"postgres"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	Log      Log      `mapstructure:"log" json:"log" yaml:"log"`
	RabbitMQ RabbitMQ `mapstructure:"rabbitmq" json:"rabbitmq" yaml:"rabbitmq"`
}

type Progres struct {
	Host            string `mapstructure:"host" json:"host" yaml:"host"`
	Port            int    `mapstructure:"port" json:"port" yaml:"port"`
	User            string `mapstructure:"user" json:"user" yaml:"user"`
	Password        string `mapstructure:"password" json:"password" yaml:"password"`
	Database        string `mapstructure:"database" json:"database" yaml:"database"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	MaxLifetime     int    `mapstructure:"conn_max_lifetime" json:"conn_max_lifetime" yaml:"conn_max_lifetime"`    // in seconds
	ConnMaxIdleTime int    `mapstructure:"conn_max_idle_time" json:"conn_max_idle_time" yaml:"conn_max_idle_time"` // in seconds
}

type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	PoolSize int    `mapstructure:"pool_size" json:"pool_size" yaml:"pool_size"`
	MinIdle  int    `mapstructure:"min_idle" json:"min_idle" yaml:"min_idle"`
}

type Log struct {
	Level string `mapstructure:"level" json:"level" yaml:"level"`
}

type RabbitMQ struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	User     string `mapstructure:"user" json:"user" yaml:"user"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

type Server struct {
	Host             string `mapstructure:"host" json:"host" yaml:"host"`
	AuthPort         int    `mapstructure:"auth_port" json:"auth_port" yaml:"auth_port"`
	PermissionPort   int    `mapstructure:"permission_port" json:"permission_port" yaml:"permission_port"`
	RolePort         int    `mapstructure:"role_port" json:"role_port" yaml:"role_port"`
	TokenPort        int    `mapstructure:"token_port" json:"token_port" yaml:"token_port"`
	MaxRecvMsgSize   int    `mapstructure:"max_recv_msg_size" json:"max_recv_msg_size" yaml:"max_recv_msg_size"`
	MaxSendMsgSize   int    `mapstructure:"max_send_msg_size" json:"max_send_msg_size" yaml:"max_send_msg_size"`
	KeepaliveTime    int    `mapstructure:"keepalive_time" json:"keepalive_time" yaml:"keepalive_time"`          // in seconds
	KeepaliveTimeout int    `mapstructure:"keepalive_timeout" json:"keepalive_timeout" yaml:"keepalive_timeout"` // in seconds
}

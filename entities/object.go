package entities

// IObject 对象基类
type IObject struct{}

// IModel 数据模型基类
type IModel struct {
	IObject
}

type IConfig interface {
	GetConfig()
}

// BaseConfig 配置对象基类
type BaseConfig struct {
	IObject `json:"-"`
	Name    string `json:",omitempty"`
	Enable  bool   `json:",default=true"`
	IConfig `json:"-"`
}

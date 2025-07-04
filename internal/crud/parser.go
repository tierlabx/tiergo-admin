package crud

import (
	"reflect"
	"strings"
)

type RouteConfig struct {
	Prefix string
	Create bool
	Update bool
	Delete bool
	Page   bool
}

func ParseModelConfig[T any]() RouteConfig {
	var entity T
	var config RouteConfig
	// 反射解析结构体体tag
	t := reflect.TypeOf(entity)
	// 指针类型取值
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("crud")
		if tag == "" {
			continue
		}
		parts := strings.Split(tag, ",")
		for _, part := range parts {
			part := strings.TrimSpace(part)
			switch {
			case strings.HasPrefix(part, "prefix:"):
				config.Prefix = strings.TrimPrefix(part, "prefix:")
			case part == "create":
				config.Create = true
			case part == "update":
				config.Update = true
			case part == "delete":
				config.Delete = true
			case part == "page":
				config.Page = true
			}
		}
	}
	return config
}

package niuhe

// niuhe 生成的路由结构
type RouteItem struct {
	Method string         `json:"method"` // API 请求方法
	Path   string         `json:"path"`   // API 请求路径
	Name   string         `json:"name"`   // API 名称
	Codes  []IntConstItem `json:"codes"`  // API 限定错误码
}

var methodCodes = map[string]int{
	"GET":      GET,
	"POST":     POST,
	"PUT":      PUT,
	"DELETE":   DELETE,
	"PATCH":    PATCH,
	"HEAD":     HEAD,
	"OPTIONS":  OPTIONS,
	"GET_POST": GET_POST,
}

func (route RouteItem) IntMethod() int {
	if method, ok := methodCodes[route.Method]; ok {
		return method
	}
	return GET_POST
}

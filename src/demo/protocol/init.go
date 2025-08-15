package protocol

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/ma-guo/niuhe"
)

type StringEnum struct {
	Value string `json:"vallue"` // 枚举值 - 数字也是 string
	Name  string `json:"name"`   // 字段名
	Desc  string `json:"desc"`   // 描述
}

// 枚举类型
type Enum struct {
	Mode   string       `json:"mode"`
	Etype  string       `json:"type"` // 'string' | 'integer'
	Name   string       `json:"name"`
	Desc   string       `json:"desc"`
	Values []StringEnum `json:"values"`
}
type Field struct {
	Label string // required, optional, repeated, mapping
	Type  string `json:"type"` // 'integer', 'decimal', 'float', 'long', 'string', 'boolean', 'message', 'enum', 'stringenum','file' 'dict' 'any'
	Name  string `json:"name"` // 字段名
	Desc  string `json:"desc"` //  描述
	Ref   string `json:"ref"`  // 引用的结构
}
type Message struct {
	Mode   string   `json:"mode"` //api
	Name   string   `json:"name"` // 名称
	Desc   string   `json:"desc"` // 描述
	Fields []*Field `json:"fields"`
}
type Route struct {
	Method string   `json:"method"` // post, get[, put, patch, delete, head, options]
	Url    string   `json:"url"`    // /api/sys/ts/
	Mode   string   `json:"mode"`   // api
	View   string   `json:"view"`
	Name   string   `json:"name"`
	Desc   string   `json:"desc"` // 描述
	Req    *Message `json:"req"`
	Rsp    *Message `json:"rsp"`
}
type Docs struct {
	App string `json:"app"`
	/** api 路由信息 */
	Routes []Route `json:"routes"`
	/** 数组 */
	Enums map[string]*Enum `json:"enums"`
	/** 消息类型 */
	Messages map[string]*Message `json:"messages"`
}

// 解析结果
type ParseResult struct {
	Interfaces string
	Beans      string
}

func (doc *Docs) getKtType(meta *Field) string {
	switch meta.Type {
	case "integer":
		return "Int"
	case "decimal":
		return "Float"
	case "float":
		return "Float"
	case "long":
		return "Long"
	case "stringenum":
	case "string":
		return "String"
	case "boolean":
		return "Boolean"
	case "message":
		if meta.Ref != "" {
			return doc.fmtName(meta.Ref)
		}
		return doc.fmtName(meta.Ref)
	case "enum":
		return "Int"
	case "file":
		return "Any?"
	case "any":
		return "Any?"
	}
	return "Any?"
}
func (doc *Docs) getKtLabel(meta *Field) string {
	switch meta.Label {
	case "required":
		return doc.getKtType(meta)
	case "optional":
		return doc.getKtType(meta) + "?"
	case "repeated":
		return "List<" + doc.getKtType(meta) + ">?"
	case "mapping":
		return "Map<String, " + doc.getKtType(meta) + ">?"
	}
	return doc.getKtType(meta)
}
func getKtDefault(ntype string) string {
	switch ntype {
	case "integer":
		return "0"
	case "decimal":
		return "0f"
	case "float":
		return "0f"
	case "long":
		return "0"
	case "string":
		return "\"\""
	case "boolean":
		return "false"
	case "message":
		return "null"
	case "enum":
		return "0"
	case "file":
		return "null"
	case "any":
		return "null"
	}
	return "null"
}
func (doc *Docs) fmtName(name string) string {
	// 从 Hrhelper.StringEnum 返回 StringEnum
	if strings.Contains(name, ".") {
		return name[strings.Index(name, ".")+1:]
	}
	return name
}

// toCamelCase 函数用于将下划线分隔的字符串转换为驼峰命名法
func (doc *Docs) toCamelCase(s string) string {
	// 按下划线分割字符串
	parts := strings.Split(s, "_")
	var result strings.Builder

	for i, part := range parts {
		if i == 0 {
			// 第一个部分保持原样
			result.WriteString(part)
		} else {
			// 后续部分首字母大写
			for j, r := range part {
				if j == 0 {
					result.WriteRune(unicode.ToUpper(r))
				} else {
					result.WriteRune(r)
				}
			}
		}
	}

	return result.String()
}

func (doc *Docs) buildBean(item *Message) string {
	var lines strings.Builder
	lines.WriteString("@Serializable\n")
	lines.WriteString("data class " + doc.fmtName(item.Name) + "(\n")

	for _, field := range item.Fields {
		if field.Desc != "" {
			lines.WriteString("\n\t/** " + field.Desc + " */\n")
		}
		lines.WriteString("\t@SerialName(\"" + field.Name + "\")\n")
		lines.WriteString("\tval " + doc.toCamelCase(field.Name) + ": " + doc.getKtLabel(field) + " = " + getKtDefault(field.Type) + ",\n")
	}
	lines.WriteString(") {}\n")
	return lines.String()
}

func (doc *Docs) buildApi(route Route) string {
	var lines strings.Builder
	lines.WriteString("\t/**\n")
	if route.Desc != "" {
		lines.WriteString("\t * " + route.Desc + "\n\t *\n")
	}
	req := route.Req

	if len(req.Fields) > 0 {
		lines.WriteString("\t * 请求参数:\n")
		for _, value := range req.Fields {
			lines.WriteString("\t * @param " + doc.toCamelCase(value.Name) + " " + value.Desc + "\n")
		}
	}
	lines.WriteString("\t */\n")
	lines.WriteString("\t@FormUrlEncoded\n")
	lines.WriteString(fmt.Sprintf("\t@%v(\"%v\")\n", strings.ToUpper(route.Method), route.Url))
	rspRef := route.Rsp
	if len(req.Fields) > 0 {
		lines.WriteString("\tsuspend fun " + doc.fmtMethodName(route) + "(\n")
		for _, value := range req.Fields {
			lines.WriteString("\t\t@Field(\"" + value.Name + "\") " + doc.toCamelCase(value.Name) + ": " + doc.getKtLabel(value) + ",\n")
		}
		lines.WriteString("\t): CommRsp<" + doc.fmtName(rspRef.Name) + ">\n")
	} else {
		lines.WriteString("\tsuspend fun " + doc.fmtMethodName(route) + "(): CommRsp<" + doc.fmtName(rspRef.Name) + ">\n")
	}

	return lines.String()
}

// 将 /api/user/info/, post 转换为 postUserInfo
func (doc *Docs) fmtMethodName(route Route) string {
	method := strings.ToLower(route.Method)
	tmps := []string{method, route.View, route.Name}
	return doc.toCamelCase(strings.Join(tmps, "_"))
}

// 解析 protocol 文档
// docs 为 protocol.json 文件内容
func Parse(appName string, docs []byte) (*ParseResult, error) {
	info := &Docs{}
	err := json.Unmarshal(docs, info)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal protocol docs: %w", err)
	}
	var interfaces strings.Builder
	interfaces.WriteString("\nimport kotlinx.serialization.SerialName\n")
	interfaces.WriteString("interface " + appName + "{\n")

	for _, route := range info.Routes {
		interfaces.WriteString(info.buildApi(route))
	}
	interfaces.WriteString("}\n")
	result := &ParseResult{
		Interfaces: interfaces.String(),
	}

	var beans strings.Builder
	beans.WriteString("\nimport kotlinx.serialization.SerialName\n")
	beans.WriteString("import kotlinx.serialization.Serializable\n\n")
	for _, msg := range info.Messages {
		beans.WriteString(info.buildBean(msg))
	}
	result.Beans = beans.String()
	return result, nil

}

func ShowKtProtocol(docPath string) {
	// docPath := "you protocol.json path"
	docs, err := os.ReadFile(docPath)
	if err != nil {
		niuhe.LogInfo("read docs error: %v", err)
		return
	}
	result, err := Parse("HrHelper", docs)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return
	}
	// 保存到对应的文件
	niuhe.LogInfo("%v", result.Interfaces)
	niuhe.LogInfo("%v", result.Beans)

}

package views

// Generated by niuhe.idl
// 此文件由 niuhe.idl 自动生成, 请勿手动修改

import (
	"demo/app/api/protos"

	"github.com/ma-guo/niuhe"
)

type _Gen_Hellow struct{}

// 示例接口
func (v *_Gen_Hellow) World_GET(c *niuhe.Context, req *protos.HelloReq, rsp *protos.HelloRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}

// 协议文档
func (v *_Gen_Hellow) Docs_GET(c *niuhe.Context, req *protos.NoneReq, rsp *protos.NoneRsp) error {
	return niuhe.NewCommError(-1, "Not Implemented")
}

// 网页示例
func (v *_Gen_Hellow) Web_GET(c *niuhe.Context) {
	// 在此实现具体方法, 需要自行处理请求参数和返回数据
}

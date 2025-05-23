package views

// Generated by niuhe.idl

import (
	"demo/app/api/protos"
	"demo/xorm/models"
	"demo/xorm/services"

	"github.com/ma-guo/niuhe"
)

type Config struct {
	_Gen_Config
}

// 分页查询获取 Config 信息
func (v *Config) Page_GET(c *niuhe.Context, req *protos.ConfigPageReq, rsp *protos.ConfigPageRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	rows, total, err := svc.Config().GetPage(req.Page, req.Size, req.Value)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	rsp.Total = total
	rsp.Items = make([]*protos.ConfigItem, 0, len(rows))
	for _, row := range rows {
		rsp.Items = append(rsp.Items, row.ToProto(nil))
	}
	return nil
}

// 查询获取 Config 信息
func (v *Config) Form_GET(c *niuhe.Context, req *protos.ConfigFormReq, rsp *protos.ConfigItem) error {
	svc := services.NewSvc()
	defer svc.Close()
	row, has, err := svc.Config().GetById(req.Id)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	if !has {
		return niuhe.NewCommError(-1, "Not Found")
	}
	row.ToProto(rsp)
	return nil
}

// 添加 Config 信息
func (v *Config) Add_POST(c *niuhe.Context, req *protos.ConfigItem, rsp *protos.ConfigItem) error {
	svc := services.NewSvc()
	defer svc.Close()
	row := &models.Config{
		Name:  req.Name,
		Value: req.Value,
	}
	err := svc.Config().Insert(row)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	row.ToProto(rsp)
	return nil
}

// 更新 Config 信息
func (v *Config) Update_POST(c *niuhe.Context, req *protos.ConfigItem, rsp *protos.ConfigItem) error {
	svc := services.NewSvc()
	defer svc.Close()
	row, has, err := svc.Config().GetById(req.Id)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	if !has {
		return niuhe.NewCommError(-1, "Not Found")
	}
	row.Name = req.Name
	row.Value = req.Value
	_, err = svc.Config().Update(row)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	row.ToProto(rsp)
	return nil
}

// 删除 Config 信息
func (v *Config) Delete_DELETE(c *niuhe.Context, req *protos.ConfigDeleteReq, rsp *protos.ConfigNoneRsp) error {
	svc := services.NewSvc()
	defer svc.Close()
	rowsMap, err := svc.Config().GetByIds(req.Ids...)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	rows := make([]*models.Config, 0)
	for _, row := range rowsMap {
		rows = append(rows, row)
	}
	err = svc.Config().Delete(rows...)
	if err != nil {
		niuhe.LogInfo("%v", err)
		return err
	}
	return nil
}
func init() {
	GetModule().Register(&Config{})
}

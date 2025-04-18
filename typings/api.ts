
// Generated by niuhe.idl, search "niuhe" plugin for vscode
// 重新生成时本文件将被复写, 请不要手动修改. 配置 "tsapi" 可指定新路径

import { ajax_get, ajax_post, ajax_any } from "./request";


/**
 * 分页查询获取 Config 信息
 * @path GET /api/config/page/
 * @param page number 页码
 * @param size number 每页数量
 * @param value number 配置值
 * @return Demo.ConfigPageRsp
 */
export const getConfigPage = (data: Demo.ConfigPageReq): Promise<Rsp<Demo.ConfigPageRsp>> => {
	return ajax_get("/api/config/page/", data);
};

/**
 * 查询获取 Config 信息
 * @path GET /api/config/form/
 * @param id number 
 * @return Demo.ConfigItem
 */
export const getConfigForm = (data: Demo.ConfigFormReq): Promise<Rsp<Demo.ConfigItem>> => {
	return ajax_get("/api/config/form/", data);
};

/**
 * 添加 Config 信息
 * @path POST /api/config/add/
 * @param id number id
 * @param name string 配置名称
 * @param value number 配置值
 * @param create_at string 创建时间
 * @param update_at string 更新时间
 * @return Demo.ConfigItem
 */
export const setConfigAdd = (data: Demo.ConfigItem): Promise<Rsp<Demo.ConfigItem>> => {
	return ajax_post("/api/config/add/", data);
};

/**
 * 更新 Config 信息
 * @path POST /api/config/update/
 * @param id number id
 * @param name string 配置名称
 * @param value number 配置值
 * @param create_at string 创建时间
 * @param update_at string 更新时间
 * @return Demo.ConfigItem
 */
export const setConfigUpdate = (data: Demo.ConfigItem): Promise<Rsp<Demo.ConfigItem>> => {
	return ajax_post("/api/config/update/", data);
};

/**
 * 删除 Config 信息
 * @path DELETE /api/config/delete/
 * @param ids number 记录id列表
 * @return Demo.ConfigNoneRsp
 */
export const deleteConfigDelete = (data: Demo.ConfigDeleteReq): Promise<Rsp<Demo.ConfigNoneRsp>> => {
	return ajax_any("delete", "/api/config/delete/", data);
};

/**
 * 示例接口
 * @path GET /api/hellow/world/
 * @param name string 用户名
 * @return Demo.HelloRsp
 */
export const getHellowWorld = (data: Demo.HelloReq): Promise<Rsp<Demo.HelloRsp>> => {
	return ajax_get("/api/hellow/world/", data);
};

/**
 * 协议文档
 * @path GET /api/hellow/docs/
 * @return Demo.NoneRsp
 */
export const getHellowDocs = (data: Demo.NoneReq): Promise<Rsp<Demo.NoneRsp>> => {
	return ajax_get("/api/hellow/docs/", data);
};

/**
 * 网页示例
 * @path GET /api/hellow/web/
 * @return Object 任意结构
 */
export const getHellowWeb = (data: Object): Promise<Object> => {
	return ajax_get("/api/hellow/web/", data);
};

/**
 * RPC测试用例
 * @path GET /api/xxx/yyy/
 * @param name string 用户名
 * @param password string 密码
 * @return Demo.XxxYyyRspMsg
  * @codes 错误码列表
 * 1 (ZH_CN) 简体中文
 * 404 (NOT_FOUND) 查找的数据不存在
*/
export const getXxxYyy = (data: Demo.XxxYyyReqMsg): Promise<Rsp<Demo.XxxYyyRspMsg>> => {
	return ajax_get("/api/xxx/yyy/", data);
};

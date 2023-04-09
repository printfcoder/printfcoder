package common

var (
	/** 通用基础业务 */

	ErrorDBQueryTotal    = &Error{No: "100001", Msg: "查询数据库总数失败"}
	ErrorDBQuery         = &Error{No: "100002", Msg: "查询数据库失败"}
	ErrorDBQueryScan     = &Error{No: "100003", Msg: "查询Scan失败"}
	ErrorInvalidJSONBody = &Error{No: "100004", Msg: "JSON 参数不正确"}
	ErrorDBBeginTx       = &Error{No: "100005", Msg: "事务启动失败"}
	ErrorDBInsert        = &Error{No: "100006", Msg: "DB插入失败"}
	ErrorDBCommit        = &Error{No: "100007", Msg: "事务提交失败"}
	ErrorDBNoSuchRecord  = &Error{No: "100008", Msg: "数据库记录不存在"}
	ErrorDBDelete        = &Error{No: "100009", Msg: "数据库删除失败"}
	ErrorParseBody       = &Error{No: "100010", Msg: "参数解析失败"}
	ErrorDBNil           = &Error{No: "100011", Msg: "DB不存在"}

	/** People业务区 */

	ErrorPeopleInvalidName             = &Error{No: "110001", Msg: "名字非法"}
	ErrorPeopleInvalidBirthDay         = &Error{No: "110002", Msg: "生日非法"}
	ErrorPeopleInvalidDeathDay         = &Error{No: "110003", Msg: "忌日非法"}
	ErrorPeopleInvalidRelationPeopleID = &Error{No: "110004", Msg: "关联人物id不正确"}
	ErrorPeopleInvalidRelation         = &Error{No: "110005", Msg: "关联人物的关系"}
	ErrorPeopleInvalidNation           = &Error{No: "110006", Msg: "国家或地区非法"}
	ErrorPeopleIsNobody                = &Error{No: "110007", Msg: "人物不存在"}

	/** 股票区 **/

	ErrorStockUnimplementedMethod      = &Error{No: "120001", Msg: "未实现的方法"}
	ErrorStockSyncStockBaseInfo        = &Error{No: "120002", Msg: "同步股票基本信息异常"}
	ErrorStockInvalidCode              = &Error{No: "120003", Msg: "证券代码不存在"}
	ErrorStockVendorGuBenInvalidStruct = &Error{No: "120004", Msg: "供应商股本结构异常"}
	ErrorStockSyncGuBenToDB            = &Error{No: "120005", Msg: "同步股票股本信息异常"}
	ErrorStockVendorGuBenIsNil         = &Error{No: "120006", Msg: "供应商股本信息不存在"}
	ErrorStockSyncAllGuBenToDB         = &Error{No: "120007", Msg: "同步所有股票股本信息异常"}
	ErrorStockVendorGuBenReadError     = &Error{No: "120008", Msg: "供应商股本结构接口读取异常"}
)

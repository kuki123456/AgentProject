package EventBusiness

import (
	"AgentProject/Utils"
	"github.com/tidwall/gjson"
	"io"
)

func JSErrorEventBusiness(JSErrorEventbody string,requestbody string,file io.Writer){
	//sin对应的索引json结构校验与事件发生时间戳
	Utils.AssertEventSin(JSErrorEventbody, file, requestbody, "[JSErrorEvent]")
	//page id H5页面id [默认值0,非空字段]
	Utils.AssertType(gjson.Get(JSErrorEventbody,"v.pvid").Type,3,file,gjson.Get(JSErrorEventbody,"v.pvid"),"v.pvid",JSErrorEventbody,"JSErrorEvent")
	//url 所属H5页面地址
	Utils.AssertType(gjson.Get(JSErrorEventbody,"v.url").Type,3,file,gjson.Get(JSErrorEventbody,"v.url"),"v.url",JSErrorEventbody,"JSErrorEvent")
    //name js文件名称
	Utils.AssertType(gjson.Get(JSErrorEventbody,"v.n").Type,3,file,gjson.Get(JSErrorEventbody,"v.n"),"v.n",JSErrorEventbody,"JSErrorEvent")
//et,js错误类型
	Utils.AssertType(gjson.Get(JSErrorEventbody,"v.et").Type,3,file,gjson.Get(JSErrorEventbody,"v.et"),"v.et",JSErrorEventbody,"JSErrorEvent")
//m,错误信息
	Utils.AssertType(gjson.Get(JSErrorEventbody,"v.m").Type,3,file,gjson.Get(JSErrorEventbody,"v.m"),"v.m",JSErrorEventbody,"JSErrorEvent")
//l,错误行
	Utils.AssertType(gjson.Get(JSErrorEventbody,"v.l").Type,2,file,gjson.Get(JSErrorEventbody,"v.l"),"v.l",JSErrorEventbody,"JSErrorEvent")
//col,错误列
	Utils.AssertType(gjson.Get(JSErrorEventbody,"v.col").Type,2,file,gjson.Get(JSErrorEventbody,"v.col"),"v.col",JSErrorEventbody,"JSErrorEvent")
//sta,错误堆栈
	Utils.AssertType(gjson.Get(JSErrorEventbody,"v.sta").Type,3,file,gjson.Get(JSErrorEventbody,"v.sta"),"v.sta",JSErrorEventbody,"JSErrorEvent")
//pct,所属H5页面创建时间
	Utils.AssertType(gjson.Get(JSErrorEventbody,"v.pct").Type,2,file,gjson.Get(JSErrorEventbody,"v.pct"),"v.pct",JSErrorEventbody,"JSErrorEvent")

}
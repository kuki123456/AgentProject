package EventBusiness

import (
	"AgentProject/Utils"
	"github.com/tidwall/gjson"
	"io"
)

func CustomEventBusiness(CustomEventEventbody string,requestbody string,file io.Writer) {
	//sin对应的索引json结构校验
	Utils.AssertEventSin(CustomEventEventbody, file, requestbody, "[CustomEventEvent]")
	//事件ID
	Utils.AssertType(gjson.Get(CustomEventEventbody,"v.i").Type,3,file,gjson.Get(CustomEventEventbody,"v.i"),"v.i",CustomEventEventbody,"CustomEventEvent")
    //"n":"",//name 事件名称 [非必要字段]
    if gjson.Get(CustomEventEventbody,"v.n").Exists(){
		Utils.AssertType(gjson.Get(CustomEventEventbody,"v.i").Type,3,file,gjson.Get(CustomEventEventbody,"v.n"),"v.n",CustomEventEventbody,"CustomEventEvent")
	}
	//"p":"",//param 附加信息 [非必要字段]
	if gjson.Get(CustomEventEventbody,"v.p").Exists(){
		Utils.AssertType(gjson.Get(CustomEventEventbody,"v.p").Type,3,file,gjson.Get(CustomEventEventbody,"v.p"),"v.p",CustomEventEventbody,"CustomEventEvent")
	}
}
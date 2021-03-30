package EventBusiness

import (
	"AgentProject/Utils"
	"github.com/tidwall/gjson"
	"io"
)

func CustomLogEventBusiness(CustomLogEventbody string,requestbody string,file io.Writer){
	//sin对应的索引json结构校验
	Utils.AssertEventSin(CustomLogEventbody, file, requestbody, "[CustomLogEvent]")
	//“i”:"",//info 信息 [非空字段]
	Utils.AssertType(gjson.Get(CustomLogEventbody,"v.i").Type,3,file,gjson.Get(CustomLogEventbody,"v.i"),"v.i",CustomLogEventbody,"CustomLogEvent")
    //"p":"",//param 附加信息 [非必要字段]
    if gjson.Get(CustomLogEventbody,"v.p").Exists(){
    	Utils.AssertType(gjson.Get(CustomLogEventbody,"v.p").Type,3,file,gjson.Get(CustomLogEventbody,"v.p"),"v.p",CustomLogEventbody,"CustomLogEvent")
	}

}
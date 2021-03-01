package EventBusiness

import (
	"AgentProject/Utils"
	"github.com/tidwall/gjson"
	"io"
)

func CustomMetricEventBusiness(CustomMetricEventbody string,requestbody string,file io.Writer){
	//sin对应的索引json结构校验
	Utils.AssertEventSin(CustomMetricEventbody, file, requestbody, "CustomMetricEvent")
	//"n":"",//name 名称 [非空字段]
	Utils.AssertType(gjson.Get(CustomMetricEventbody,"v.n").Type,3,file,gjson.Get(CustomMetricEventbody,"v.n"),"v.n",CustomMetricEventbody)
	//“v”:1,//value 值 [非空字段]
	Utils.AssertType(gjson.Get(CustomMetricEventbody,"v.v").Type,2,file,gjson.Get(CustomMetricEventbody,"v.v"),"v.v",CustomMetricEventbody)
	if gjson.Get(CustomMetricEventbody,"v.p").Exists(){
		Utils.AssertType(gjson.Get(CustomMetricEventbody,"v.p").Type,3,file,gjson.Get(CustomMetricEventbody,"v.p"),"v.p",CustomMetricEventbody)
	}
}

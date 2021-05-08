package EventBusiness

import (
	"AgentProject/Utils"
	"github.com/tidwall/gjson"
	"io"
)

func CrashEventBusiness(CrashEventbody string,requestbody string,file io.Writer){
	//sin对应的索引json结构校验
	Utils.AssertEventSin(CrashEventbody, file, requestbody, "[CrashEvent]")
	//"ic":true,//iscustom 是否是自定义 [默认值false,非空字段]
	//先判断是否是自定义异常
	Utils.AssertBool(gjson.Get(CrashEventbody,"v.ic").Type,file,gjson.Get(CrashEventbody,"v.ic").Value(),"v.ic",CrashEventbody,"CrashEvent")
	if gjson.Get(CrashEventbody,"v.ic").Bool(){
		Utils.AssertType(gjson.Get(CrashEventbody,"v.t").Type,3,file,gjson.Get(CrashEventbody,"v.t").Value(),"v.t",CrashEventbody,"CrashEvent")
		Utils.AssertType(gjson.Get(CrashEventbody,"v.p").Type,3,file,gjson.Get(CrashEventbody,"v.p").Value(),"v.p",CrashEventbody,"CrashEvent")
		Utils.AssertType(gjson.Get(CrashEventbody,"v.cab").Type,3,file,gjson.Get(CrashEventbody,"v.cab").Value(),"v.cab",CrashEventbody,"CrashEvent")
	}else {
		//cab
		Utils.AssertType(gjson.Get(CrashEventbody,"v.cab").Type,3,file,gjson.Get(CrashEventbody,"v.cab").Value(),"v.cab",CrashEventbody,"CrashEvent")
		//t
		Utils.AssertType(gjson.Get(CrashEventbody,"v.t").Type,3,file,gjson.Get(CrashEventbody,"v.t").Value(),"v.t",CrashEventbody,"CrashEvent")
		//cti
		Utils.AssertType(gjson.Get(CrashEventbody,"v.cti").Type,3,file,gjson.Get(CrashEventbody,"v.cti").Value(),"v.cti",CrashEventbody,"CrashEvent")
		//mti
		Utils.AssertType(gjson.Get(CrashEventbody,"v.mti").Type,3,file,gjson.Get(CrashEventbody,"v.mti").Value(),"v.mti",CrashEventbody,"CrashEvent")
		//bi
		Utils.AssertType(gjson.Get(CrashEventbody,"v.bi").Type,3,file,gjson.Get(CrashEventbody,"v.bi").Value(),"v.bi",CrashEventbody,"CrashEvent")
		//tdi
		Utils.AssertType(gjson.Get(CrashEventbody,"v.tdi").Type,5,file,gjson.Get(CrashEventbody,"v.tdi").Value(),"v.tdi",CrashEventbody,"CrashEvent")
		for _,value:=range gjson.Get(CrashEventbody,"v.tdi").Array(){
			//"ti":"",//tid 线程tid [非空字段]
			Utils.AssertType(gjson.Get(value.String(),"ti").Type,3,file,gjson.Get(value.String(),"ti").Value(),"ti",value.String(),"CrashEvent")
			//"n":"",//name 线程名称 [非空字段]
			Utils.AssertType(gjson.Get(value.String(),"n").Type,3,file,gjson.Get(value.String(),"n").Value(),"n",value.String(),"CrashEvent")
			//"di":"",//dump info 堆栈信息 [非空字段]
			Utils.AssertType(gjson.Get(value.String(),"di").Type,3,file,gjson.Get(value.String(),"di").Value(),"di",value.String(),"CrashEvent")
		}
	}
}
package EventBusiness

import (
	"AgentProject/Utils"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
)

func LagEventBusiness(LagEventbody string,requestbody string,file io.Writer){
	//sin对应的索引json结构校验
	Utils.AssertEventSin(LagEventbody, file, requestbody, "LagEvent")
	//"bi":"",//binary info iOS二进制信息 [非必要字段]
	Utils.AssertType(gjson.Get(LagEventbody,"v.bi").Type,3,file,gjson.Get(LagEventbody,"v.bi").Value(),"v.bi",LagEventbody)
	//"vn":"",//view name 发生视图名称 [非必要字段]
	Utils.AssertType(gjson.Get(LagEventbody,"v.vn").Type,3,file,gjson.Get(LagEventbody,"v.vn").Value(),"v.vn",LagEventbody)
	//"tdi":{},//ThreadDumpInfo 线程堆栈信息 [只要主线程]
	Utils.AssertType(gjson.Get(LagEventbody,"v.tdi").Type,5,file,gjson.Get(LagEventbody,"v.tdi").Value(),"v.tdi",LagEventbody)
	Utils.AssertType(gjson.Get(LagEventbody,"v.tdi.ti").Type,3,file,gjson.Get(LagEventbody,"v.tdi.ti").Value(),"v.tdi.ti",LagEventbody)
	Utils.AssertType(gjson.Get(LagEventbody,"v.tdi.n").Type,3,file,gjson.Get(LagEventbody,"v.tdi.n").Value(),"v.tdi.n",LagEventbody)
	Utils.AssertType(gjson.Get(LagEventbody,"v.tdi.di").Type,3,file,gjson.Get(LagEventbody,"v.tdi.di").Value(),"v.tdi.di",LagEventbody)
	fmt.Println("LAG字段校验完成!")
}
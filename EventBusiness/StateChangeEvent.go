package EventBusiness

import (
	"AgentProject/Utils"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"time"
)

func StateChangeEventBusiness(StateChangeEventbody string,requestbody string,file io.Writer){
	//sin对应的索引json结构校验
	Utils.AssertEventSin(StateChangeEventbody, file, requestbody, "[StateChangeEvent]")
	//1,//type 状态切换类型「1:网络切换,2:置前台,3:置后台」[非空字段]
	Utils.AssertType(gjson.Get(StateChangeEventbody,"v.t").Type,2,file,gjson.Get(StateChangeEventbody,"v.t").Value(),"v.t",StateChangeEventbody,"StateChangeEvent")
	if gjson.Get(StateChangeEventbody,"v.t").Type==1 || gjson.Get(StateChangeEventbody,"v.t").Type==2 || gjson.Get(StateChangeEventbody,"v.t").Type==3{
		fmt.Println("StateChangeEvent字段t数值正确!")
		if gjson.Get(StateChangeEventbody,"v.t").Type==1{
			Utils.AssertType(gjson.Get(StateChangeEventbody,"v.tn.ns").Type,3,file,gjson.Get(StateChangeEventbody,"v.tn.ns").Value(),"v.tn.ns",StateChangeEventbody,"StateChangeEvent")
			if gjson.Get(StateChangeEventbody,"v.tn.dsi").Exists(){
				Utils.AssertType(gjson.Get(StateChangeEventbody,"v.tn.dsi").Type,3,file,gjson.Get(StateChangeEventbody,"v.tn.dsi").Value(),"v.tn.dsi",StateChangeEventbody,"StateChangeEvent")

			}
			if gjson.Get(StateChangeEventbody,"v.tn.dip").Exists(){
				Utils.AssertType(gjson.Get(StateChangeEventbody,"v.tn.dip").Type,3,file,gjson.Get(StateChangeEventbody,"v.tn.dip").Value(),"v.tn.dip",StateChangeEventbody,"StateChangeEvent")

			}
		}
	}else {
		_, _ = fmt.Fprintf(file, "%v[ERROR]:[StateChangeEvent]--%s的值错误：[当前值:]%v\n", time.Now().Format("2006/01/02 15:04:05"),"StateChangeEvent.tn.ns",gjson.Get(StateChangeEventbody,"v.t").Num)
		//Utils.EmailTo(fmt.Sprintf("%s捕捉到异常(Statechange.t值异常)\n,%s的值错误：[当前值:]%v请求体如下，请排查日志:\n\n%s",time.Now().Format("2006/01/02 15:04:05"),"statechange.t",gjson.Get(StateChangeEventbody,"v.t").Num,requestbody))
	}
	fmt.Println("StateChangeEvent参数校验完毕!")
}
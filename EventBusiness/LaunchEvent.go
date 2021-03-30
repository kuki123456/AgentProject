package EventBusiness

import (
	"AgentProject/Utils"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"time"
)

func LaunchEventBusiness(LaunchEventbody string,requestbody string,file io.Writer){
	//sin对应的索引json结构校验与事件发生事件戳校验
	Utils.AssertEventSin(LaunchEventbody, file, requestbody, "[LaunchEvent]")
	//t启动类型
	Utils.AssertType(gjson.Get(LaunchEventbody,"v.t").Type,2,file,gjson.Get(LaunchEventbody,"v.t"),"v.t",LaunchEventbody,"LaunchEvent")
	if gjson.Get(LaunchEventbody,"v.t").Num==1 || gjson.Get(LaunchEventbody,"v.t").Num==2{
		fmt.Println("LaunchEvent字段t值正确!")
	}else {
		_, _ = fmt.Fprintf(file, "%v[ERROR]:[LaunchEvent]--%s的值错误：[当前:]%v\n", time.Now().Format("2006/01/02 15:04:05"),"LaunchEvent.t",gjson.Get(LaunchEventbody," v.t").Num)
		//Utils.EmailTo(fmt.Sprintf("%s捕捉到异常(LaunchEvent.t异常)\n,%s的值错误：[当前值:]%v请求体如下，请排查日志:\n\n%s",time.Now().Format("2006/01/02 15:04:05"),"LaunchEvent.t",gjson.Get(LaunchEventbody,"v.t").Value(),requestbody))

	}
	//lt
	Utils.AssertType(gjson.Get(LaunchEventbody,"v.lt").Type,2,file,gjson.Get(LaunchEventbody,"v.lt"),"v.lt",LaunchEventbody,"LaunchEvent")
	//判断是否发生慢启动
	//冷启动判断
	if gjson.Get(LaunchEventbody,"v.t").Num==1 && gjson.Get(LaunchEventbody,"v.lt").Num>1000000{
		//is是否为true
		if gjson.Get(LaunchEventbody,"v.is").Bool(){
			fmt.Println("is与阈值判断符合!")
		}else {
			//Utils.EmailTo(fmt.Sprintf("%v冷启动事件中is字段与阈值判断不符,期望值:TURE,当前值:%v\n请求体:%s\n",time.Now().Format("2006/01/02 15:04:05"),gjson.Get(LaunchEventbody,"v.is").Bool(),LaunchEventbody))
			_,_=fmt.Fprintf(file,"%v[ERROR]:[LaunchEvent]--冷启动事件中is字段与阈值判断不符,期望值:TURE,当前值:%v\n请求体:%v\n",time.Now().Format("2006/01/02 15:04:05"),gjson.Get(LaunchEventbody,"v.is").Bool(),LaunchEventbody)
		}
		//tmi是否存在
		Utils.AssertType(gjson.Get(LaunchEventbody,"v.tmi").Type,5,file,gjson.Get(LaunchEventbody,"v.tmi"),"v.tmi",LaunchEventbody,"LaunchEvent")
		//tmi结构体字段校验
		for _,value:=range gjson.Get(LaunchEventbody,"v.tmi").Array(){
			Utils.AssertType(gjson.Get(value.String(),"ti").Type,3,file,gjson.Get(value.String(),"ti"),"ti",value.String(),"LaunchEvent")
			Utils.AssertBool(gjson.Get(value.String(),"im").Type,file,gjson.Get(value.String(),"im"),"LaunchEvent.tmi.im","LaunchEvent")
			Utils.AssertType(gjson.Get(value.String(),"n").Type,3,file,gjson.Get(value.String(),"n"),"n",value.String(),"LaunchEvent")
			Utils.AssertType(gjson.Get(value.String(),"mi").Type,5,file,gjson.Get(value.String(),"mi"),"mi",value.String(),"LaunchEvent")
			for _,element:=range gjson.Get(value.String(),"mi").Array(){
				Utils.AssertType(gjson.Get(element.String(),"st").Type,2,file,gjson.Get(element.String(),"st"),"st",element.String(),"LaunchEvent")
				Utils.AssertType(gjson.Get(element.String(),"et").Type,2,file,gjson.Get(element.String(),"et"),"et",element.String(),"LaunchEvent")
				Utils.AssertType(gjson.Get(element.String(),"n").Type,3,file,gjson.Get(element.String(),"n"),"n",element.String(),"LaunchEvent")
				if gjson.Get(element.String(),"p").Exists(){
					Utils.AssertType(gjson.Get(element.String(),"p").Type,3,file,gjson.Get(element.String(),"p"),"p",element.String(),"LaunchEvent")
				}
				//5.13新增
				//Utils.AssertBool(gjson.Get(element.String(),"ic").Type,file,gjson.Get(element.String(),"ic"),"ic",element.String())
			}
		}
		fmt.Println("COOLLaunch的tmi校验完毕!")
	}
			//热启动判断
	if gjson.Get(LaunchEventbody,"v.t").Num==2 && gjson.Get(LaunchEventbody,"v.lt").Num>200000{
		//is是否为true
		if gjson.Get(LaunchEventbody,"v.is").Bool(){
			fmt.Println("is与阈值判断符合!")
		}else {
			//Utils.EmailTo(fmt.Sprintf("冷启动事件中is字段与阈值判断不符,期望值:TURE,当前值:%v",gjson.Get(LaunchEventbody,"v.is").Bool()))
			_,_=fmt.Fprintf(file,"%v[ERROR]:[LaunchEvent]--冷启动事件中is字段与阈值判断不符,期望值:TURE,当前值:%v\n",gjson.Get(LaunchEventbody,"v.is").Bool())
		}
		//tmi是否存在(存在监控方法还没调用启动已完成情况)
		if gjson.Get(LaunchEventbody,"v.tmi").Exists(){
			Utils.AssertType(gjson.Get(LaunchEventbody,"v.tmi").Type,5,file,gjson.Get(LaunchEventbody,"v.tmi"),"v.tmi",LaunchEventbody,"LaunchEvent")
		}
		//tmi结构体字段校验
		for _,value:=range gjson.Get(LaunchEventbody,"v.tmi").Array(){
			Utils.AssertType(gjson.Get(value.String(),"ti").Type,3,file,gjson.Get(value.String(),"ti"),"ti",value.String(),"LaunchEvent")
			Utils.AssertBool(gjson.Get(value.String(),"im").Type,file,gjson.Get(value.String(),"im"),"LaunchEvent.tmi.im","LaunchEvent")
			Utils.AssertType(gjson.Get(value.String(),"n").Type,3,file,gjson.Get(value.String(),"n"),"n",value.String(),"LaunchEvent")
			Utils.AssertType(gjson.Get(value.String(),"mi").Type,5,file,gjson.Get(value.String(),"mi"),"mi",value.String(),"LaunchEvent")
			for _,element:=range gjson.Get(value.String(),"mi").Array(){
				Utils.AssertType(gjson.Get(element.String(),"st").Type,2,file,gjson.Get(element.String(),"st"),"st",element.String(),"LaunchEvent")
				Utils.AssertType(gjson.Get(element.String(),"et").Type,2,file,gjson.Get(element.String(),"et"),"et",element.String(),"LaunchEvent")
				Utils.AssertType(gjson.Get(element.String(),"n").Type,3,file,gjson.Get(element.String(),"n"),"n",element.String(),"LaunchEvent")
				if gjson.Get(element.String(),"p").Exists(){
					Utils.AssertType(gjson.Get(element.String(),"p").Type,3,file,gjson.Get(element.String(),"p"),"p",element.String(),"LaunchEvent")
				}
				//5.13新增
				//Utils.AssertBool(gjson.Get(element.String(),"ic").Type,file,gjson.Get(element.String(),"ic"),"ic",element.String())
			}
		}
		fmt.Println("HotLaunch的tmi校验完毕!")
	}
}
package EventBusiness

import (
	"AgentProject/Utils"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"time"
)

func ViewEventBusiness(ViewEventbody string,requestbody string,file io.Writer){
	//sin对应的索引json结构校验
	Utils.AssertEventSin(ViewEventbody, file, requestbody, "[ViewEvent]")
	//"ci":"",//correlation id 视图关联ID [非空字段]
	Utils.AssertType(gjson.Get(ViewEventbody,"v.ci").Type,3,file,gjson.Get(ViewEventbody,"v.ci").Value(),"v.ci",ViewEventbody,"ViewEvent")
	// "lt":1,//load time 加载耗时 [单位us,非空字段]
	Utils.AssertType(gjson.Get(ViewEventbody,"v.lt").Type,2,file,gjson.Get(ViewEventbody,"v.lt").Value(),"v.lt",ViewEventbody,"ViewEvent")
	//"m":1,//model 方式 「1:进入，2:退出」[非空字段]
	Utils.AssertType(gjson.Get(ViewEventbody,"v.m").Type,2,file,gjson.Get(ViewEventbody,"v.m").Value(),"v.m",ViewEventbody,"ViewEvent")
	if gjson.Get(ViewEventbody,"v.m").Num ==1 || gjson.Get(ViewEventbody,"v.m").Num==2{
		fmt.Println("m字段在范围内!")
	}else {
		_, _ = fmt.Fprintf(file, "%v[ERROR]:[ViewEvent]--%s的值错误：[当前:]%v,[期望:]%v\n", time.Now().Format("2006/01/02 15:04:05"),"v.m",gjson.Get(ViewEventbody,"v.m").Num,"1 or 2")
		//Utils.EmailTo(fmt.Sprintf("%s捕捉到异常(字段m值不对)\n,%s的值错误：[当前:]%v,[期望:]%v\n请求体如下，请排查日志:\n\n%s",time.Now().Format("2006/01/02 15:04:05"),"v.m",gjson.Get(ViewEventbody,"v.m").Num,"1 OR 2",ViewEventbody))
	}
	// "st":1,//stay time 停留耗时 [单位us,非必要字段]
	if gjson.Get(ViewEventbody,"v.m").Num==2{
		Utils.AssertType(gjson.Get(ViewEventbody,"v.st").Type,2,file,gjson.Get(ViewEventbody,"v.st").Value(),"v.st",ViewEventbody,"ViewEvent")
	}
	//"t":1,//type 类型 「1:h5,2:activity,3:fragment,4:window,5:controler,6:rn」[非空字段]
	typeslice:=[]float64{1,2,3,4,5,6}
	Utils.AssertType(gjson.Get(ViewEventbody,"v.t").Type,2,file,gjson.Get(ViewEventbody,"v.t").Value(),"v.t",ViewEventbody,"ViewEvent")
	if Utils.Assertin(gjson.Get(ViewEventbody,"v.t").Num,typeslice){
		fmt.Println("t字段在范围内!")
	}else {
		_, _ = fmt.Fprintf(file, "%v[ERROR]:%s的值错误：[当前:]%v,[期望:]%v\n", time.Now().Format("2006/01/02 15:04:05"),"v.t",gjson.Get(ViewEventbody,"v.t").Num,typeslice)
		//Utils.EmailTo(fmt.Sprintf("%s捕捉到异常(字段m值不对)\n,%s的值错误：[当前:]%v,[期望:]%v\n请求体如下，请排查日志:\n\n%s",time.Now().Format("2006/01/02 15:04:05"),"v.t",gjson.Get(ViewEventbody,"v.t").Num,typeslice,ViewEventbody))
	}
	//"pv":"",//parent view 父视图 [非必要字段]
	if gjson.Get(ViewEventbody,"v.pv").Exists(){
		Utils.AssertType(gjson.Get(ViewEventbody,"v.pv").Type,3,file,gjson.Get(ViewEventbody,"v.pv").Value(),"v.pv",ViewEventbody,"ViewEvent")
	}
	//"n":"", //name 视图名称 [非空字段]
	Utils.AssertType(gjson.Get(ViewEventbody,"v.n").Type,3,file,gjson.Get(ViewEventbody,"v.n").Value(),"v.n",ViewEventbody,"ViewEvent")
	//"ic":true,//is custom 是否是自定义 [默认值false,非空字段]
	Utils.AssertBool(gjson.Get(ViewEventbody,"v.ic").Type,file,gjson.Get(ViewEventbody,"v.ic").Value(),"v.ic",ViewEventbody,"ViewEvent")
	if gjson.Get(ViewEventbody,"v.ic").Bool(){
		//自定义视图就判断p字段
		Utils.AssertType(gjson.Get(ViewEventbody,"v.p").Type,3,file,gjson.Get(ViewEventbody,"v.p").Value(),"v.p",ViewEventbody,"ViewEvent")

	}else {
		//非自定义就判断"is":false,//is slow 视图是否发生慢加载 [非自定义数据非空, 自定义数据为空]
		Utils.AssertBool(gjson.Get(ViewEventbody,"v.is").Type,file,gjson.Get(ViewEventbody,"v.is").Value(),"v.is",ViewEventbody,"ViewEvent")
		if gjson.Get(ViewEventbody,"v.is").Bool()&&gjson.Get(ViewEventbody,"v.m").Num==1{
			//"tmi":[{}],//ThreadMethodInfo 线程方法信息 [慢加载线程方法信息,进入方式有,非必要字段]
			Utils.AssertType(gjson.Get(ViewEventbody,"v.tmi").Type,5,file,gjson.Get(ViewEventbody,"v.tmi").Value(),"v.tmi",ViewEventbody,"ViewEvent")
			for _,value:=range gjson.Get(ViewEventbody,"v.tmi").Array(){
				Utils.AssertType(gjson.Get(value.String(),"ti").Type,3,file,gjson.Get(value.String(),"ti"),"ti",value.String(),"ViewEvent")
				Utils.AssertBool(gjson.Get(value.String(),"im").Type,file,gjson.Get(value.String(),"im").Value(),"ViewEvent.tmi.im",ViewEventbody,"ViewEvent")
				Utils.AssertType(gjson.Get(value.String(),"n").Type,3,file,gjson.Get(value.String(),"n"),"n",value.String(),"ViewEvent")
				Utils.AssertType(gjson.Get(value.String(),"mi").Type,5,file,gjson.Get(value.String(),"mi"),"mi",value.String(),"ViewEvent")
				for _,element:=range gjson.Get(value.String(),"mi").Array(){
					Utils.AssertType(gjson.Get(element.String(),"st").Type,2,file,gjson.Get(element.String(),"st"),"st",element.String(),"ViewEvent")
					Utils.AssertType(gjson.Get(element.String(),"et").Type,2,file,gjson.Get(element.String(),"et"),"et",element.String(),"ViewEvent")
					Utils.AssertType(gjson.Get(element.String(),"n").Type,3,file,gjson.Get(element.String(),"n"),"n",element.String(),"ViewEvent")
					Utils.AssertBool(gjson.Get(element.String(),"ic").Type,file,gjson.Get(element.String(),"ic").Value(),"ic",element.String(),"ViewEvent")
					if gjson.Get(element.String(),"p").Exists(){
						Utils.AssertType(gjson.Get(element.String(),"p").Type,3,file,gjson.Get(element.String(),"p"),"p",element.String(),"ViewEvent")
					}
					//5.13新增
					//Utils.AssertBool(gjson.Get(element.String(),"ic").Type,file,gjson.Get(element.String(),"ic"),"ic",element.String())
				}
			}
		}else {
			fmt.Println("未发生慢视图!")
		}
	}
fmt.Println("ViewEvent 参数校验完成!")
}
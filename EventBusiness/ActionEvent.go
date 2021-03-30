package EventBusiness

import (
	"AgentProject/Utils"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"time"
)

func ActionEventBusiness(ActionEventbody string,requestbody string,file io.Writer) {
	typeslice := []float64{0, 1, 2, 3}
	//sin对应的索引json结构校验
	Utils.AssertEventSin(ActionEventbody, file, requestbody, "	[ActionEvent]")
	//"t":1,//type 操作类型「0:其他 1:点击,2:手势，3:键盘」[非空字段]
	Utils.AssertType(gjson.Get(ActionEventbody, "v.t").Type, 2, file, gjson.Get(ActionEventbody, "v.t"), "v.t", ActionEventbody,"ActionEvent")
	if Utils.Assertin(gjson.Get(ActionEventbody, "v.t").Num, typeslice) {
		fmt.Println("t值在范围中!")
	} else {
		_, _ = fmt.Fprintf(file, "%v[ERROR]:[ActionEvent]--%s的值错误：[当前:]%v,[期望:]%v\n", time.Now().Format("2006/01/02 15:04:05"), "art", gjson.Get(ActionEventbody, "v.t").Num, typeslice)
		//Utils.EmailTo(fmt.Sprintf("%s捕捉到异常(art值不在范围内)\n,%s的值类型错误：[当前:]%v,[期望:]%v请求体如下，请排查日志:\n%s", time.Now().Format("2006/01/02 15:04:05"), "art", gjson.Get(ActionEventbody, "v.t").Num, typeslice, ActionEventbody))
	}
	//"n":"",//name 名称 「method」[非空字段]
	Utils.AssertType(gjson.Get(ActionEventbody, "v.n").Type, 3, file, gjson.Get(ActionEventbody, "v.n"), "v.n", ActionEventbody,"ActionEvent")
	//"i":"",//info 信息  [非空字段]
	Utils.AssertType(gjson.Get(ActionEventbody, "v.i").Type, 3, file, gjson.Get(ActionEventbody, "v.i"), "v.i", ActionEventbody,"ActionEvent")
	//"vn":"",//view name 发生视图名称 [非空字段]
	Utils.AssertType(gjson.Get(ActionEventbody, "v.vn").Type, 3, file, gjson.Get(ActionEventbody, "v.vn"), "v.vn", ActionEventbody,"ActionEvent")
	//"ic":true,//is custom 是否是自定义 [默认值false,非空字段]
	Utils.AssertBool(gjson.Get(ActionEventbody, "v.ic").Type, file, gjson.Get(ActionEventbody, "v.ic").Value(), "v.ic","ActionEvent")
	//"p":"",//param 参数 [非必要字段]
	if gjson.Get(ActionEventbody, "v.p").Exists() {
		Utils.AssertType(gjson.Get(ActionEventbody, "v.p").Type, 3, file, gjson.Get(ActionEventbody, "v.p"), "v.p", ActionEventbody,"ActionEvent")
	}
	//"lt":1,//load time 加载耗时 [单位us,非空字段]
	Utils.AssertType(gjson.Get(ActionEventbody, "v.lt").Type, 2, file, gjson.Get(ActionEventbody, "v.lt"), "v.lt", ActionEventbody,"ActionEvent")
	//"is":false,//is slow 是否发生慢操作 [非自定义数据非空, 自定义数据为空]
	if gjson.Get(ActionEventbody, "v.ic").Bool() {
		fmt.Println("当前为自定义操作数据!")
	} else {
		Utils.AssertBool(gjson.Get(ActionEventbody, "v.is").Type, file, gjson.Get(ActionEventbody, "v.is").Value(), "v.is","ActionEvent")
	}
	//"tmi":[{}],//ThreadMethodInfo 线程方法信息 [慢操作线程方法信息,非必要字段]
	if gjson.Get(ActionEventbody, "v.is").Bool() {
		for _, value := range gjson.Get(ActionEventbody, "v.tmi").Array() {
			Utils.AssertType(gjson.Get(value.String(), "ti").Type, 3, file, gjson.Get(value.String(), "ti"), "ti", value.String(),"ActionEvent")
			Utils.AssertBool(gjson.Get(value.String(), "im").Type, file, gjson.Get(value.String(), "im"), "LaunchEvent.tmi.im","ActionEvent")
			Utils.AssertType(gjson.Get(value.String(), "n").Type, 3, file, gjson.Get(value.String(), "n"), "n", value.String(),"ActionEvent")
			Utils.AssertType(gjson.Get(value.String(), "mi").Type, 5, file, gjson.Get(value.String(), "mi"), "mi", value.String(),"ActionEvent")
			for _, element := range gjson.Get(value.String(), "mi").Array() {
				Utils.AssertType(gjson.Get(element.String(), "st").Type, 2, file, gjson.Get(element.String(), "st"), "st", element.String(),"ActionEvent")
				Utils.AssertType(gjson.Get(element.String(), "et").Type, 2, file, gjson.Get(element.String(), "et"), "et", element.String(),"ActionEvent")
				Utils.AssertType(gjson.Get(element.String(), "n").Type, 3, file, gjson.Get(element.String(), "n"), "n", element.String(),"ActionEvent")
				if gjson.Get(element.String(), "p").Exists() {
					Utils.AssertType(gjson.Get(element.String(), "p").Type, 3, file, gjson.Get(element.String(), "p"), "p", element.String(),"ActionEvent")
				}
				//5.13新增
				//Utils.AssertBool(gjson.Get(element.String(),"ic").Type,file,gjson.Get(element.String(),"ic"),"ic",element.String())
			}
		}
	}
}
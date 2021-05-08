package EventBusiness

import (
	"AgentProject/Utils"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"time"
)

func RouteChangeEventBusiness(RouteChangeEventtbody string,requestbody string,file io.Writer){
	//sin对应的索引json结构校验及事件发生时刻时间戳
	Utils.AssertEventSin(RouteChangeEventtbody, file, requestbody, "[RouteChangeEvent]")
	//"tu":"",   //toUrl 目标路由 [非空字段]
	Utils.AssertType(gjson.Get(RouteChangeEventtbody,"v.tu").Type,3,file,gjson.Get(RouteChangeEventtbody,"v.tu").Value(),"v.tu",RouteChangeEventtbody,"RouteChangeEvent")
	//"fu":"",   //fromUrl 来源路由 [非空字段]
	Utils.AssertType(gjson.Get(RouteChangeEventtbody,"v.fu").Type,3,file,gjson.Get(RouteChangeEventtbody,"v.fu").Value(),"v.fu",RouteChangeEventtbody,"RouteChangeEvent")
	//"d":0,     //duration 路由切换耗时 [单位us，非空字段]
	Utils.AssertType(gjson.Get(RouteChangeEventtbody,"v.d").Type,2,file,gjson.Get(RouteChangeEventtbody,"v.d").Value(),"v.d",RouteChangeEventtbody,"RouteChangeEvent")
	//"sta":0,   //status 路由切换状态「0:正常 2:异常」 [非空字段]
	Utils.AssertType(gjson.Get(RouteChangeEventtbody,"v.sta").Type,2,file,gjson.Get(RouteChangeEventtbody,"v.d").Value(),"v.d",RouteChangeEventtbody,"RouteChangeEvent")
	sta_eum:=[]float64{0,2}
	if Utils.Assertin(gjson.Get(RouteChangeEventtbody,"v.sta").Num,sta_eum){
		fmt.Println("sta在枚举范围内!")
	}else {
		_, _ = fmt.Fprintf(file, "%v[ERROR]:[RouteChangeEvent]--%s的值错误：[当前:]%v,[期望:]%v\n", time.Now().Format("2006/01/02 15:04:05"),"sta",gjson.Get(RouteChangeEventtbody,"v.sta").Num,sta_eum)
	}
	//"al":"",   //alias 路由地址名称(别名) [非必要字段]
	if gjson.Get(RouteChangeEventtbody,"v.al").Exists(){
		Utils.AssertType(gjson.Get(RouteChangeEventtbody,"v.al").Type,3,file,gjson.Get(RouteChangeEventtbody,"v.al").Value(),"v.al",RouteChangeEventtbody,"RouteChangeEvent")
	}
	//"pt":"",   //path 当前路由的子路由地址 [非空字段]
	Utils.AssertType(gjson.Get(RouteChangeEventtbody,"v.pt").Type,3,file,gjson.Get(RouteChangeEventtbody,"v.pt").Value(),"v.pt",RouteChangeEventtbody,"RouteChangeEvent")
	//"rt":"",   //root 路由全量地址(去锚点) [非空字段]
	Utils.AssertType(gjson.Get(RouteChangeEventtbody,"v.rt").Type,3,file,gjson.Get(RouteChangeEventtbody,"v.rt").Value(),"v.rt",RouteChangeEventtbody,"RouteChangeEvent")
	//"pu":"",   //page url 主页面地址 [非空字段]
	Utils.AssertType(gjson.Get(RouteChangeEventtbody,"v.pu").Type,3,file,gjson.Get(RouteChangeEventtbody,"v.pu").Value(),"v.pu",RouteChangeEventtbody,"RouteChangeEvent")
	//"fw":"",   //framework 框架名称「Vue,Angular,React」 [非空字段]
	Utils.AssertType(gjson.Get(RouteChangeEventtbody,"v.fw").Type,3,file,gjson.Get(RouteChangeEventtbody,"v.fw").Value(),"v.fw",RouteChangeEventtbody,"RouteChangeEvent")
	//"ic":false,//isCustom 是否为自定义  [非空字段]
	Utils.AssertBool(gjson.Get(RouteChangeEventtbody,"v.ic").Type,file,gjson.Get(RouteChangeEventtbody,"v.ic").Value(),"v.ic",RouteChangeEventtbody,"RouteChangeEvent")
	//"ctp":1,   //client type 客户端类型「1:PC 2:移动」[非空字段]
	Utils.AssertType(gjson.Get(RouteChangeEventtbody,"v.ctp").Type,2,file,gjson.Get(RouteChangeEventtbody,"v.ctp").Value(),"v.ctp",RouteChangeEventtbody,"RouteChangeEvent")
    ctp_eum:=[]float64{1,2}
	if Utils.Assertin(gjson.Get(RouteChangeEventtbody,"v.ctp").Num,ctp_eum){
		fmt.Println("ctp在枚举范围内!")
	}else {
		_, _ = fmt.Fprintf(file, "%v[ERROR]:[RouteChangeEvent]--%s的值错误：[当前:]%v,[期望:]%v\n", time.Now().Format("2006/01/02 15:04:05"),"ctp",gjson.Get(RouteChangeEventtbody,"v.ctp").Num,sta_eum)
	}
}

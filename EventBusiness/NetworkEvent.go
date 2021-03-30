package EventBusiness

import (
	"AgentProject/Utils"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"time"
)
func NetworkBusiness(networkbody string,requestbody string,file io.Writer){
	Utils.AssertEventSin(networkbody,file,requestbody,"[NetworkEvent]")
	//网络事件实体非空字段类型判断,针对必要字段校验
	//ru请求地址字段类型判断
	Utils.AssertType(gjson.Get(networkbody,"v.ru").Type,3,file,gjson.Get(networkbody,"v.ru"),"v.ru",networkbody,"NetworkEvent")
	//method 请求方式,m
	Utils.AssertType(gjson.Get(networkbody,"v.m").Type,3,file,gjson.Get(networkbody,"v.m"),"v.m",networkbody,"NetworkEvent")
	//target ip 目标IP,ti
	Utils.AssertType(gjson.Get(networkbody,"v.ti").Type,3,file,gjson.Get(networkbody,"v.ti"),"v.ti",networkbody,"NetworkEvent")
	//target port 目标端口tp
	Utils.AssertType(gjson.Get(networkbody,"v.tp").Type,2,file,gjson.Get(networkbody,"v.tp"),"v.tp",networkbody,"NetworkEvent")
	//dns time dns解析时间
	Utils.AssertType(gjson.Get(networkbody,"v.dt").Type,2,file,gjson.Get(networkbody,"v.dt"),"v.dt",networkbody,"NetworkEvent")
	if gjson.Get(networkbody,"v.dt").Num ==0 || 999<=gjson.Get(networkbody,"v.dt").Num{
		fmt.Println("dt值范围通过")
	}else{
		_, _ = fmt.Fprintf(file, "%v[ERROR]:[NetworkEvent]--%s的值错误：[当前值:]%v\n", time.Now().Format("2006/01/02 15:04:05"),gjson.Get(networkbody,"v.dt").Num)
		//Utils.EmailTo(fmt.Sprintf("%s捕捉到异常(network.dt值异常)\n,%s的值错误：[当前值:]%v请求体如下，请排查日志:\n\n%s",time.Now().Format("2006/01/02 15:04:05"),"network.dt",gjson.Get(networkbody,"v.dt").Num,requestbody))
	}
	//connect time tcp建连时间ct
	Utils.AssertType(gjson.Get(networkbody,"v.ct").Type,2,file,gjson.Get(networkbody,"v.ct"),"v.ct",networkbody,"NetworkEvent")
	if gjson.Get(networkbody,"v.ct").Num ==0 || 999<=gjson.Get(networkbody,"v.ct").Num{
		fmt.Println("ct值范围通过")
	}else{
		_, _ = fmt.Fprintf(file, "%v[ERROR]:[NetworkEvent]--%s的值错误：[当前值:]%v\n", time.Now().Format("2006/01/02 15:04:05"),gjson.Get(networkbody,"v.ct").Num)
		//Utils.EmailTo(fmt.Sprintf("%s捕捉到异常(network.ct值异常)\n,%s的值错误：[当前值:]%v请求体如下，请排查日志:\n\n%s",time.Now().Format("2006/01/02 15:04:05"),"network.ct",gjson.Get(networkbody,"v.ct").Num,requestbody))
	}
	//ssl time ssl解析时间sslt
	Utils.AssertType(gjson.Get(networkbody,"v.sslt").Type,2,file,gjson.Get(networkbody,"v.sslt"),"v.sslt",networkbody,"NetworkEvent")
	if gjson.Get(networkbody,"v.sslt").Num ==0 || 999<=gjson.Get(networkbody,"v.sslt").Num{
		fmt.Println("sslt值范围通过")
	}else{
		_, _ = fmt.Fprintf(file, "%v[ERROR]:[NetworkEvent]--%s的值错误：[当前值:]%v\n", time.Now().Format("2006/01/02 15:04:05"),gjson.Get(networkbody,"v.sslt").Num)
		//Utils.EmailTo(fmt.Sprintf("%s捕捉到异常(network.sslt值异常)\n,%s的值错误：[当前值:]%v请求体如下，请排查日志:\n\n%s",time.Now().Format("2006/01/02 15:04:05"),"network.sslt",gjson.Get(networkbody,"v.sslt").Num,requestbody))
	}
	//request time 请求时间 rt
	Utils.AssertType(gjson.Get(networkbody,"v.rt").Type,2,file,gjson.Get(networkbody,"v.rt"),"v.rt",networkbody,"NetworkEvent")
	if gjson.Get(networkbody,"v.rt").Num ==0 || 999<=gjson.Get(networkbody,"v.rt").Num{
		fmt.Println("rt值范围通过")
	} else{
		_, _ = fmt.Fprintf(file, "%v[ERROR]:[NetworkEvent]--%s的值错误：[当前值:]%v\n", time.Now().Format("2006/01/02 15:04:05"),gjson.Get(networkbody,"v.rt").Num)
		//Utils.EmailTo(fmt.Sprintf("%s捕捉到异常(network.rt值异常)\n,%s的值错误：[当前值:]%v请求体如下，请排查日志:\n\n%s",time.Now().Format("2006/01/02 15:04:05"),"network.rt",gjson.Get(networkbody,"v.rt").Num,requestbody))
	}
	//response time 响应时间 「有过程>=999，无过程0」[单位us，非空字段]
	Utils.AssertType(gjson.Get(networkbody,"v.rti").Type,2,file,gjson.Get(networkbody,"v.rti"),"v.rti",networkbody,"NetworkEvent")
	if gjson.Get(networkbody,"v.rti").Num ==0 || 999<=gjson.Get(networkbody,"v.rti").Num{
		fmt.Println("rti值范围通过")
	}else{
		_, _ = fmt.Fprintf(file, "%v[ERROR]:[NetworkEvent]--%s的值错误：[当前值:]%v\n", time.Now().Format("2006/01/02 15:04:05"),gjson.Get(networkbody,"v.rti").Num)
		//Utils.EmailTo(fmt.Sprintf("%s捕捉到异常(network.rti值异常)\n,%s的值错误：[当前值:]%v请求体如下，请排查日志:\n\n%s",time.Now().Format("2006/01/02 15:04:05"),"network.rti",gjson.Get(networkbody,"v.rti").Num,requestbody))
	}
	//download time 下载用时 「有过程>=999，无过程0」[单位us，非空字段]
	Utils.AssertType(gjson.Get(networkbody,"v.dti").Type,2,file,gjson.Get(networkbody,"v.dti"),"v.dti",networkbody,"NetworkEvent")
	if gjson.Get(networkbody,"v.dti").Num ==0 || 999<=gjson.Get(networkbody,"v.dti").Num{
		fmt.Println("dti值范围通过")
	}else{
		_, _ = fmt.Fprintf(file, "%v[ERROR]:[NetworkEvent]--%s的值错误：[当前值:]%v\n", time.Now().Format("2006/01/02 15:04:05"),gjson.Get(networkbody,"v.dti").Num)
		//Utils.EmailTo(fmt.Sprintf("%s捕捉到异常(network.dti值异常)\n,%s的值错误：[当前值:]%v请求体如下，请排查日志:\n\n%s",time.Now().Format("2006/01/02 15:04:05"),"network.dti",gjson.Get(networkbody,"v.dti").Num,requestbody))
	}
	//download size 下载大小 [单位Byte，非空字段]
	Utils.AssertType(gjson.Get(networkbody,"v.ds").Type,2,file,gjson.Get(networkbody,"v.ds"),"v.ds",networkbody,"NetworkEvent")
	//protocol type 协议类型 「1:h1,2:h1s,3:h2,5:ws,6:wss,7:tcp,10:udp」[非空字段]
	Utils.AssertType(gjson.Get(networkbody,"v.pt").Type,2,file,gjson.Get(networkbody,"v.pt"),"v.pt",networkbody,"NetworkEvent")
	//art app request type App请求类型枚举
	ArtArrary :=[]float64{0,1,2,3,10}
	if Utils.Assertin(gjson.Get(networkbody,"v.art").Num,ArtArrary){
		fmt.Println("art值在范围中!")
	}else {
		_, _ = fmt.Fprintf(file, "%v[ERROR]:[NetworkEvent]--%s的值错误：[当前:]%v,[期望:]%v\n", time.Now().Format("2006/01/02 15:04:05"),"art",gjson.Get(networkbody,"v.art").Num,ArtArrary)
		//Utils.EmailTo(fmt.Sprintf("%s捕捉到异常(art值不在范围内)\n,%s的值类型错误：[当前:]%v,[期望:]%v请求体如下，请排查日志:\n%s",time.Now().Format("2006/01/02 15:04:05"),"art",gjson.Get(networkbody,"v.art").Num,ArtArrary,networkbody))

	}
	Utils.AssertType(gjson.Get(networkbody,"v.art").Type,2,file,gjson.Get(networkbody,"v.art"),"v.art",networkbody,"NetworkEvent")
	//“ic”:true,//is custom 是否是自定义 [默认值false,非空字段]
	Utils.AssertBool(gjson.Get(networkbody,"v.ic").Type,file,gjson.Get(networkbody,"v.ic"),"v.ic","NetworkEvent")
	//非必要字段校验
	//cna
	if gjson.Get(networkbody,"v.cna").Exists(){
		Utils.AssertType(gjson.Get(networkbody,"v.cna").Type,5,file,gjson.Get(networkbody,"v.cna"),"v.cna",networkbody,"NetworkEvent")
	}else {
		fmt.Println("该网络请求没有cna字段!")
	}
	//"rh":"1",//request header 请求header [非必要字段]
	if gjson.Get(networkbody,"v.rh").Exists() {
		Utils.AssertType(gjson.Get(networkbody,"v.rh").Type,3,file,gjson.Get(networkbody,"v.rh"),"v.rh",networkbody,"NetworkEvent")
	}else {
		fmt.Println("该网络请求没有rh字段!")
	}
	  //"rhe":"",//response header 响应header [非必要字段]
	if gjson.Get(networkbody,"v.rhe").Exists() {
		Utils.AssertType(gjson.Get(networkbody,"v.rhe").Type,3,file,gjson.Get(networkbody,"v.rhe"),"v.rhe",networkbody,"NetworkEvent")
	}else {
		fmt.Println("该网络请求没有rhe字段!")
	}
	//"tid":"", //trace id 端到端打通功能 [非必要字段]
	if gjson.Get(networkbody,"v.tid").Exists() {
		Utils.AssertType(gjson.Get(networkbody,"v.tid").Type,3,file,gjson.Get(networkbody,"v.tid"),"v.tid",networkbody,"NetworkEvent")
	}else {
		fmt.Println("该网络请求没有tid字段!")
	}
	//“ep”:"",//error platform 错误码分区(http,js,mpaas及ios特有域(一段字符串)) [非必要字段]
	if gjson.Get(networkbody,"v.ep").Exists() {
		Utils.AssertType(gjson.Get(networkbody,"v.ep").Type,3,file,gjson.Get(networkbody,"v.ep"),"v.ep",networkbody,"NetworkEvent")
	}else {
		fmt.Println("该网络请求没有ep字段!")
	}
	//"eop":1,//error occurrent process 错误发生的阶段 「1:SSL过程，2:DNS过程，3:TCP过程，4:其他过程」[非必要字段]
	if gjson.Get(networkbody,"v.eop").Exists() {
		Utils.AssertType(gjson.Get(networkbody,"v.eop").Type,2,file,gjson.Get(networkbody,"v.eop"),"v.eop",networkbody,"NetworkEvent")
	}else {
		fmt.Println("该网络请求没有eop字段!")
	}
	// "ec":1,//error code 错误码 [非必要字段]
	if gjson.Get(networkbody,"v.ec").Exists() {
		Utils.AssertType(gjson.Get(networkbody,"v.ec").Type,2,file,gjson.Get(networkbody,"v.ec"),"v.ec",networkbody,"NetworkEvent")
	}else {
		fmt.Println("该网络请求没有ec字段!")
	}
	//"em":"",//error msg 错误信息 [非必要字段]
	if gjson.Get(networkbody,"v.em").Exists() {
		Utils.AssertType(gjson.Get(networkbody,"v.em").Type,3,file,gjson.Get(networkbody,"v.em"),"v.em",networkbody,"NetworkEvent")
	}else {
		fmt.Println("该网络请求没有em字段!")
	}
	//“ret”:“”,//resource type 资源类型 [统一响应头字段,非必要字段]
	if gjson.Get(networkbody,"v.ret").Exists() {
		Utils.AssertType(gjson.Get(networkbody,"v.ret").Type,3,file,gjson.Get(networkbody,"v.ret"),"v.ret",networkbody,"NetworkEvent")
	}else {
		fmt.Println("该网络请求没有ret字段!")
	}
	////page id 页面ID 网络请求对应的页面ID（目前原生网络为空，WebView网络非JS探针数据为空） [非必要字段]
	if gjson.Get(networkbody,"v.pvid").Exists() {
		Utils.AssertType(gjson.Get(networkbody,"v.pvid").Type,3,file,gjson.Get(networkbody,"v.pvid"),"v.pvid",networkbody,"NetworkEvent")
	}else {
		fmt.Println("该网络请求没有pvid字段!")
	}
	//"cbh":"",//custom business header 自定义业务头 [非必要字段]
	if gjson.Get(networkbody,"v.cbh").Exists() {
		Utils.AssertType(gjson.Get(networkbody,"v.cbh").Type,3,file,gjson.Get(networkbody,"v.cbh"),"v.cbh",networkbody,"NetworkEvent")
	}else {
		fmt.Println("该网络请求没有cbh字段!")
	}
}

package main

import (
	"AgentProject/EventBusiness"
	"AgentProject/Utils"
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)
func max(vals...float64) float64 {
	var max float64
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}


func min(vals...float64) float64 {
	var min float64
	for _, val := range vals {
		if min == 0 || val <= min {
			min = val
		}
	}
	return min
}
var lastmaxent float64
var temmin float64
var count int
type param struct {
	Url string `json:"url"`
	Data string `json:"data"`
}
type ToolStruct struct {
	Ai  ToolStruct_sub1 `json:"ai"`
	Di  ToolStruct_sub2 `json:"di"`
	Nsi ToolStruct_sub3 `json:"nsi"`
	UI  ToolStruct_sub4 `json:"ui"`
	V   string          `json:"v"`
}

type ToolStruct_sub2 struct {
	A   string  `json:"a"`
	Bn  string  `json:"bn"`
	Ch  string  `json:"ch"`
	Ci  string  `json:"ci"`
	Cm  string  `json:"cm"`
	Di  string  `json:"di"`
	Ds  string  `json:"ds"`
	L   string  `json:"l"`
	M   string  `json:"m"`
	Obv string  `json:"obv"`
	Omv string  `json:"omv"`
	Ot  int64   `json:"ot"`
	RAM float64   `json:"ram"`
	Rom float64 `json:"rom"`
}

type ToolStruct_sub1 struct {
	Ai string `json:"ai"`
	An string `json:"an"`
	Av string `json:"av"`
	Ci string `json:"ci"`
}

type ToolStruct_sub4 struct {
	Ei string `json:"ei"`
	UI string `json:"ui"`
}

type ToolStruct_sub3 struct {
	Ns string `json:"ns"`
}
func ParseGzip(data []byte) ([]byte, error) {
	b := new(bytes.Buffer)
	_ = binary.Write(b, binary.LittleEndian, data)
	r, err := gzip.NewReader(b)
	if err != nil {
		log.Printf("[ParseGzip] NewReader error: %v, maybe data is ungzip", err)
		return nil, err
	} else {
		defer r.Close()
		undatas, err := ioutil.ReadAll(r)
		if err != nil {
			log.Printf("[ParseGzip]  ioutil.ReadAll error: %v", err)
			return nil, err
		}
		return undatas, nil
	}
}
func CompressGzip(data []byte)*bytes.Buffer{
	var result bytes.Buffer
	newwriter :=gzip.NewWriter(&result)
	defer newwriter.Close()
	_, _ = newwriter.Write(data)
	_ = newwriter.Flush()
	return &result
}
/*字段类型断言：
0：null
1：bool-false
2：number
3：string
4：bool-true
5：json结构块
 */
func ReturninputMap(input int)[]map[string]interface{} {
	switch input {
	case 0:
		return []map[string]interface{}{{"c":100,"n":"network"}}
	case 1:
		return []map[string]interface{}{{"c":100,"n":"h5"}}
	case 2:
		return []map[string]interface{}{{"c":100,"n":"crash"}}
	case 3:
		return []map[string]interface{}{{"c":100,"n":"view","tv":200}}
	case 4:
		return []map[string]interface{}{{"c":100,"n":"coollaunch","tv":500},{"c":100,"n":"hotlaunch","tv":200}}
	case 5:
		return []map[string]interface{}{{"c":100,"n":"coollaunch","tv":500},{"c":100,"n":"hotlaunch","tv":200}}
	case 6:
		return []map[string]interface{}{{"c":100,"n":"action","tv":3000}}
	case 7:
		return []map[string]interface{}{{"c":100,"n":"lag","tv":5}}

	case 8:
		return []map[string]interface{}{{"c":100,"n":"statechange"}}
	case 9:
		return []map[string]interface{}{{"c":100,"n":"anr"}}
	case 10:
		return []map[string]interface{}{{"c":100,"n":"customlog"}}
	case 11:
		return []map[string]interface{}{{"c":100,"n":"customevent"}}
	case 12:
		return []map[string]interface{}{{"c":100,"n":"custommetric"}}
	case 13:
		return []map[string]interface{}{{"c":100,"n":"routechange"}}
	case 14:
		//return []map[string]interface{}{{"c":100,"n":"network"},{"c":100,"n":"h5"},{"c":100,"n":"view","tv":2000},{"c":100,"n":"crash"},{"c":100,"n":"coollaunch","tv":500},{"c":100,"n":"hotlaunch","tv":200},{"c":100,"n":"action","tv":3000},{"c":100,"n":"statechange"},{"c":100,"n":"anr"},{"c":100,"n":"customlog"},{"c":100,"n":"customevent"},{"c":100,"n":"custommetric"},{"c":100,"n":"lag","tv":6}}
		return []map[string]interface{}{{"n":"routechange","c":100},{"c":100,"n":"network"},{"c":100,"n":"h5"},{"c":100,"n":"view","tv":2000},{"c":100,"n":"crash"},{"c":100,"n":"coollaunch","tv":500},{"c":100,"n":"hotlaunch","tv":200},{"c":100,"n":"action","tv":3000},{"c":100,"n":"statechange"},{"c":100,"n":"anr"},{"c":100,"n":"customlog"},{"c":100,"n":"customevent"},{"c":100,"n":"custommetric"},{"c":100,"n":"lagfps","tv":59},{"c":100,"n":"lagstuck","tv":6}}
	default:
		log.Println("输入有误！")
	}
return nil
}
var logtask ="Documents/BRSDK/Log/log-2021-01-04-11-13-25.txt"
func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	r.LoadHTMLFiles("./custom-service.html")
	r.GET("set", func(context *gin.Context) {
	    task:=context.Query("task")
	    logtask=task
	    context.JSON(http.StatusOK,gin.H{"code":200,"message":fmt.Sprintf("%s:更新成功！",task)})
	})
	//const (
	//	network int=iota
	//	h5
	//	crash
	//	view
	//	coollaunch
	//	hotlaunch
	//	action
	//	lag
	//	statechange
	//	anr
	//	customlog
	//	customevent
	//	custommetric
	//	all
	//)
	//log.Printf("开启network模块请输入对应数字：%v\n",network)
	//log.Printf("开启h5模块请输入对应数字：%v\n",h5)
	//log.Printf("开启crash模块请输入对应数字：%v\n",crash)
	//log.Printf("开启view模块请输入对应数字：%v\n",view)
	//log.Printf("开启coollaunch模块请输入对应数字：%v\n",coollaunch)
	//log.Printf("开启hotlaunch模块请输入对应数字：%v\n",hotlaunch)
	//log.Printf("开启action模块请输入对应数字：%v\n",action)
	//log.Printf("开启lag模块请输入对应数字：%v\n",lag)
	//log.Printf("开启statechange模块请输入对应数字：%v\n",statechange)
	//log.Printf("开启anr模块请输入对应数字：%v\n",anr)
	//log.Printf("开启customlog模块请输入对应数字：%v\n",customlog)
	//log.Printf("开启customevent模块请输入对应数字：%v\n",customevent)
	//log.Printf("开启custommetric模块请输入对应数字：%v\n",custommetric)
	//log.Printf("开启所有模块请输入对应数字：%v\n",all)
	//var input string
	//fmt.Println("请输入需要开启的模块！")
	//fmt.Scanln(&input)
	//var ip string
	//fmt.Println("请输入upload的IP地址！")
	//fmt.Scanln(&ip)
	r.GET("/test", func(context *gin.Context) {
		context.HTML(http.StatusOK,"custom-service.html",gin.H{"hello":"world"})
	})
	r.POST("/config",func(c *gin.Context){
		version :=c.Query("v")
		appid :=c.Query("a")
		device := c.Query("d")
		//请求体解析分析
		var request_body ToolStruct
		body,_ := ioutil.ReadAll(c.Request.Body)
		Address :="http://devtest.ibr.cc:20107/config?"+"v="+version+"&a="+appid+"&d="+device
		//转发到对应测试环境
		fmt.Println("ADDRESS",Address)
		result,_:=ParseGzip(body)
		err:=json.Unmarshal(result,&request_body)
		RESPONSE:=Utils.CONFIGData(Address,strings.NewReader(string(result)))
		log.Printf("CONFIG请求body是：%v",string(result))
		log.Println(c.ClientIP())
		if err!=nil{
			log.Panicf("请求体参数与定义字段类型不符：%v",err)
		}
		//var mc int
		//mc,_=strconv.Atoi(input)
		var cp=map[string]interface{}{"log":map[string]interface{}{"upload_user_info":"http://192.168.1.152:8080/update/userinfo","upload_logfile":"http://192.168.31.124:8080/upload/logfile","get_logtask":"http://192.168.31.124:8080/get/logtask"},"speed":""}
		cp_byte,_:=json.Marshal(cp)
		cp_json:=string(cp_byte)
		c.JSON(http.StatusOK, gin.H{
			"brss" : true,
			"di" : c.ClientIP(),
			"gdia" : "http://devtest.ibr.cc/grip",
			"gpa" : "http://www.baidu.com:80",
			"rc" : 10000,
			"rct" : 15,
			//"s" : fmt.Sprintf("2726b602-07c5-447a-abeb-26a121106754%v",int32(time.Now().Unix())),
			"s":gjson.Get(RESPONSE,"s").String(),
			"sat" : 24,
			"sp" : 100,
			"spt" : 1440,
			"st" : int64(time.Now().UnixNano())/1000,
			"ua" : fmt.Sprintf("http://%v:8080/upload","192.168.31.124"),
			//"ua":"http://devtest.ibr.cc:20107/upload",
			"mc":ReturninputMap(14),
           "cp":cp_json,
           "cbhk":[]string{"a","b"},
			"cbbk":[]string{"a","b"},
			"cbqk":[]string{"a","b"},
			"usc":1,

		})
		})
	r.POST("/upload",func(c *gin.Context){
		file,_:=os.OpenFile(fmt.Sprintf("./log/%v%v%v_ERR.log",time.Now().Year(),int(time.Now().Month()),time.Now().Day()),os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_RDONLY,os.ModeAppend|os.ModePerm)
		defer file.Close()
		version :=c.Query("v")
		appid :=c.Query("a")
		device :=c.Query("d")
		var request_body map[string]interface{}
		body,_ := ioutil.ReadAll(c.Request.Body)
		result,_:=ParseGzip(body)
		err:=json.Unmarshal(result,&request_body)
		if err!=nil{
			log.Panicf("请求体解析错误：%v",err)
		}else {
			log.Printf("UPLOAD请求body是：%v\n",request_body)
			log.Printf("UPLOAD请求body是：%v\n",string(result))
		}
		Address :="http://devtest.ibr.cc:20107/upload?"+"v="+version+"&a="+appid+"&d="+device
		fmt.Println("ADDRESS\n",Address)
		////转发到对应测试环境
		paramer:=param{
			Url: Address,
			Data: string(result),
		}
		param_byte,_:=json.Marshal(paramer)
		Utils.UPLOADData("http://127.0.0.1:8000/upload",strings.NewReader(string(param_byte)))
		//command := exec.Command("/Users/wangxiaoxiao/AgentProject/Utils/venv/bin/python", "/Users/wangxiaoxiao/AgentProject/Utils/upload.py ",Address,string(result))
		//////创建获取命令输出管道
		//out, err := command.CombinedOutput()
		//if err != nil {
		//	fmt.Printf("combined out:\n%s\n", string(out))
		//	log.Fatalf("cmd.Run() failed with %s\n", err)
		//}
		//fmt.Printf("combined out:\n%s\n", string(out))
		//执行命令
		//if err := cmd.Start(); err != nil {
		//	fmt.Println("Error:The command is err,", err)
		//	return
		//}
		fmt.Println("******************************sdk版本*****************************")
		v:=gjson.Get(string(result),"v")
		Utils.AssertType(v.Type,3,file,v,"v",string(result),"UPLOAD")
		fmt.Println("******************************会话ID*****************************")
		s:=gjson.Get(string(result),"s")
		Utils.AssertType(s.Type,3,file,s,"s",string(result),"UPLOAD")
		fmt.Println("******************************upload监控时间*****************************")
		mt:=gjson.Get(string(result),"mt").Value()
		fmt.Println(time.Unix(int64(mt.(float64)/1000/1000),0).Format("2006-01-02" ),time.Now().Format("2006-01-02"))
		if time.Unix(int64(mt.(float64)/1000/1000),0).Format("2006-01-02" )==time.Now().Format("2006-01-02"){
			fmt.Println("mt校验正确!")
		}else {
			_, _ = fmt.Fprintf(file, "%v[ERROR]:[UPLOAD]--mt时间戳上报错误,上报值:%v\n",time.Now().Format("2006/01/02 15:04:05"),mt)
			//Utils.EmailTo(fmt.Sprintf("%s捕捉到异常(mt上报异常):\n,mt的值为:%v\n,请求体为:%s",time.Now().Format("2006/01/02 15:04:05"),mt,string(result)))
		}
		fmt.Println("******************************Config下发的服务器时间*****************************")
		cmt:=gjson.Get(string(result),"cmt").Value()
		fmt.Println(time.Unix(int64(cmt.(float64)/1000/1000),0).Format("2006-01-02" ),time.Now().Format("2006-01-02"))
		//if time.Unix(int64(cmt.(float64)/1000/1000),0).Format("2006-01-02" )==time.Now().Format("2006-01-02"){
		//	fmt.Println("cmt校验正确!")
		//}else {
		//	_, _ = fmt.Fprintf(file, "%v[ERROR]:cmt时间戳上报错误,上报值:%v\n", time.Now().Format("2006/01/02 15:04:05"),cmt)
		//	Utils.EmailTo(fmt.Sprintf("%s捕捉到异常(cmt上报异常):\n,cmt的值为:%v\n,请求体为:%s",time.Now().Format("2006/01/02 15:04:05"),cmt,string(result)))
		//}
		Utils.AssertType(gjson.Get(string(result),"cmt").Type,2,file,gjson.Get(string(result),"cmt").Value(),"cmt",string(result),"UPLOAD")
//******************设备信息字段校验************************
		//设备信息的设备ID字段类型断言
		fmt.Println("****************************设备信息*************************")
		di:=gjson.Get(string(result), "di.di")
		Utils.AssertType(di.Type,3,file,di,"di.di",string(result),"UPLOAD")
		//设备信息的cpu指令集校验
		ci:=gjson.Get(string(result), "di.ci")
		Utils.AssertType(ci.Type,3,file,ci,"di.ci",string(result),"UPLOAD")
		//设备信息的品牌校验
		bn:=gjson.Get(string(result), "di.bn")
		if bn.Raw!="apple"&& bn.Type.String()!="String" {
			_, _ = fmt.Fprintf(file, "%v[ERROR]:[UPLOAD]--di.bn的值类型错误：%v\n", time.Now().Format("2006/01/02 15:04:05"),bn.Type)
		}else {
			log.Printf("di.bn类型是string，值：%v\n",bn)
		}
		//设备信息的cpu厂商校验
		ch:=gjson.Get(string(result), "di.ch")
		if ch.Raw!="apple" && ch.Type.String()!="String" {
			_, _ = fmt.Fprintf(file, "%v[ERROR]:[UPLOAD]--di.ch的值类型错误：%v\n",time.Now().Format("2006/01/02 15:04:05"), ch.Type)
		}else {
			log.Printf("di.ch类型是string，值：%v\n",ch)
		}
		//设备信息的CPU型号
		cm:=gjson.Get(string(result), "di.cm")
		//if cm.Type.String()!="String" {
		//	_, _ = fmt.Fprintf(file, "%v[ERROR]:di.cm的值类型错误：%v\n",time.Now().Format("2006/01/02 15:04:05"), cm.Type)
		//}else {
		//	log.Printf("di.cm类型是string，值：%v\n",cm)
		//}
		Utils.AssertType(cm.Type,3,file,cm,"di.cm",string(result),"UPLOAD")
		//设备信息的屏幕分辨率字段校验
		ds:=gjson.Get(string(result), "di.ds")
		//if ds.Type.String()!="String" {
		//	_, _ = fmt.Fprintf(file, "%v[ERROR]:di.ds的值类型错误：%v\n",time.Now().Format("2006/01/02 15:04:05"), ds.Type)
		//}else {
		//	log.Printf("di.ds类型是string，值：%v\n",ds)
		//}
		Utils.AssertType(ds.Type,3,file,ds,"di.ds",string(result),"UPLOAD")
		//设备信息的语言字段校验
		l:=gjson.Get(string(result), "di.l")
		//if l.Type.String()!="String" {
		//	_, _ = fmt.Fprintf(file, "%v[ERROR]:di.l的值类型错误：%v\n",time.Now().Format("2006/01/02 15:04:05"), l.Type)
		//}else {
		//	log.Printf("di.l类型是string，值：%v\n",l)
		//}
		Utils.AssertType(l.Type,3,file,l,"di.l",string(result),"UPLOAD")
		//设备信息的手机型号校验
		m:=gjson.Get(string(result), "di.m")
		//if m.Type.String()!="String" {
		//	_, _ = fmt.Fprintf(file, "%v[ERROR]:di.m的值类型错误：%v\n", time.Now().Format("2006/01/02 15:04:05"),m.Type)
		//}else {
		//	log.Printf("di.m类型是string，值：%v\n",m)
		//}
		Utils.AssertType(m.Type,3,file,m,"di.m",string(result),"UPLOAD")
		//设备信息的编译版本校验
		obv:=gjson.Get(string(result), "di.obv")
		//if obv.Type.String()!="String" {
		//	_, _ = fmt.Fprintf(file, "%v[ERROR]:di.obv的值类型错误：%v\n", time.Now().Format("2006/01/02 15:04:05"),obv.Type)
		//}else {
		//	log.Printf("di.obv类型是string，值：%v\n",obv)
		//}
		Utils.AssertType(obv.Type,3,file,obv,"di.obv",string(result),"UPLOAD")
		//设备信息的os版本校验
		omv:=gjson.Get(string(result), "di.omv")
		//if omv.Type.String()!="String" {
		//	_, _ = fmt.Fprintf(file, "%v[ERROR]:di.omv的值类型错误：%v\n", time.Now().Format("2006/01/02 15:04:05"),omv.Type)
		//}else {
		//	log.Printf("di.omv类型是string，值：%v\n",omv)
		//}
		Utils.AssertType(omv.Type,3,file,omv,"di.omv",string(result),"UPLOAD")
		//设备信息的系统标志校验
		ot:=gjson.Get(string(result), "di.ot")
		//if ot.Type!=2 {
		//	_, _ = fmt.Fprintf(file, "%v[ERROR]:di.ot的值类型错误：%v\n",time.Now().Format("2006/01/02 15:04:05"), ot.Type)
		//}else {
		//	log.Printf("di.ot类型是number，值：%v\n",ot)
		//}
		Utils.AssertType(ot.Type,2,file,ot,"di.ot",string(result),"UPLOAD")
		Utils.AssertNumValue(ot.Num,0,file,"di.ot",string(result),"UPLOAD")
		//设备信息剩余内存校验
		ram:=gjson.Get(string(result), "di.ram")
		Utils.AssertType(ram.Type,2,file,ram,"di.ram",string(result),"UPLOAD")
		//设备信息的存储空间校验
		rom:=gjson.Get(string(result), "di.rom")
		Utils.AssertType(rom.Type,2,file,rom,"di.rom",string(result),"UPLOAD")
/****************************************device end***********************/
/****************************************appinfo start***********************/
		fmt.Println("****************************应用信息*************************")
		//应用ID校验
		ai:=gjson.Get(string(result), "ai.ai")
		Utils.AssertType(ai.Type,3,file,ai,"ai.ai",string(result),"UPLOAD")
		//app名称
		an:=gjson.Get(string(result), "ai.an")
		Utils.AssertType(an.Type,3,file,an,"ai.an",string(result),"UPLOAD")
		//渠道商
		if gjson.Get(string(result), "ai.ci").Exists() {
			Utils.AssertType(gjson.Get(string(result), "ai.ci").Type, 3, file, gjson.Get(string(result), "ai.ci"), "ai.ci", string(result), "UPLOAD")
		}//app版本
		av:=gjson.Get(string(result), "ai.av")
		Utils.AssertType(av.Type,3,file,av,"ai.av",string(result),"UPLOAD")
/****************************************appinfo end************************************/
/****************************************fui************************************/
		if gjson.Get(string(result),"fui").Exists(){
			Utils.AssertType(gjson.Get(string(result),"fui").Type,3,file,gjson.Get(string(result),"fui").Value(),"fui",string(result),"UPLOAD")
		}
/****************************************ti************************************/
			if gjson.Get(string(result),"ti").Exists(){
				Utils.AssertType(gjson.Get(string(result),"ti").Type,5,file,gjson.Get(string(result),"ti").Value(),"ti",string(result),"UPLOAD")
			}
			for _,val:=range gjson.Get(string(result),"ti").Array(){
				Utils.AssertType(gjson.Get(val.String(),"tu").Type,2,file,gjson.Get(val.String(),"tu").Value(),"tu",val.String(),"UPLOAD")
				Utils.AssertType(gjson.Get(val.String(),"sin").Type,3,file,gjson.Get(val.String(),"sin").Value(),"sin",val.String(),"UPLOAD")
			}
			if gjson.Get(string(result),"ui").Exists(){
				log.Print("ui 存在！")
			}
		/****************************************设备状态信息*************************************/
fmt.Println("*******************************各事件业务****************************************")
		var testarrary []float64
		eventbody:=gjson.Get(string(result),"e").Array()
		for _,value:=range eventbody{
			testarrary=append(testarrary,gjson.Get(value.String(),"ent").Num)
			switch gjson.Get(value.String(),"k").String() {
			case "network":
				fmt.Println("**************************检测network****************************")
				EventBusiness.NetworkBusiness(value.String(),string(result),file)
			case "jserror":
				fmt.Println("*******************************检测jserror********************")
				EventBusiness.JSErrorEventBusiness(value.String(),string(result),file)
			case "view":
				fmt.Println("*******************************检测View********************")
				EventBusiness.ViewEventBusiness(value.String(),string(result),file)
			case "crash":
				fmt.Println("***************************检测crash************************")
				EventBusiness.CrashEventBusiness(value.String(),string(result),file)
			case "action":
				fmt.Println("***************************检测action************************")
				EventBusiness.ActionEventBusiness(value.String(),string(result),file)
			case "statechange":
				fmt.Println("***************************检测statechange************************")
				EventBusiness.StateChangeEventBusiness(value.String(),string(result),file)
			case "lag":
				fmt.Println("***************************检测lag************************")
				EventBusiness.LagEventBusiness(value.String(),string(result),file)
			case "launch":
				fmt.Println("***************************检测launch************************")
				EventBusiness.LaunchEventBusiness(value.String(),string(result),file)
			case "customevent":
				fmt.Println("***************************检测customevent************************")
				EventBusiness.CustomEventBusiness(value.String(),string(result),file)
			case "customlog":
				fmt.Println("***************************检测customlog************************")
				EventBusiness.CustomLogEventBusiness(value.String(),string(result),file)
			case "custommetric":
				fmt.Println("***************************检测custommetric************************")
				EventBusiness.CustomMetricEventBusiness(value.String(),string(result),file)
			case "h5":
				fmt.Println("***************************检测H5************************")
				EventBusiness.H5EventBusiness(value.String(),string(result),file)
			case "routechange":
				fmt.Println("***************************检测routechange************************")
				EventBusiness.RouteChangeEventBusiness(value.String(),string(result),file)
			}
		//}
		//if count==0{
		//	lastmaxent=max(testarrary...)
		//	temmin=min(testarrary...)
		//	if lastmaxent-temmin==gjson.Get(string(result),"usd").Num{
		//		fmt.Println("首包usd正常!")
		//	}
		//	count=count+1
		//}else {
		//	fmt.Println("上一包最大的ent:",lastmaxent)
		//	fmt.Println("当前包最大的ent:",max(testarrary...))
		//	fmt.Println("当前包最小的ent:",min(testarrary...))
		//	fmt.Println("当前temmin",temmin)
			fmt.Println("usd",gjson.Get(string(result),"usd").Num)
			if gjson.Get(string(result),"usd").Num >3600000000 {
				     fmt.Fprintf(file,"%v[ERROR]:[UPLOAD]的usd异常！\n", time.Now().Format("2006/01/02 15:04:05"))
					fmt.Fprintf(file,"%s\n",string(result))
			}
			//if max(testarrary...)-lastmaxent==gjson.Get(string(result),"usd").Num{
			//	fmt.Println("usd正常!")
			//}else {
			//	fmt.Println("usd异常!")
			//}
			if gjson.Get(string(result),"usd").Num<0{
				Utils.EmailTo("usd出现负值")
				fmt.Fprintf(file,"%v[ERROR]:[UPLOAD]的usd异常！\n", time.Now().Format("2006/01/02 15:04:05"))
				fmt.Fprintf(file,"%s\n",string(result))
			}
			//if temmin>min(testarrary...){
			//	fmt.Println("temmin yichang!")
			//	temmin=min(testarrary...)
			//}
			//lastmaxent=max(testarrary...)
		}
		c.JSON(http.StatusOK, gin.H{
			"rc":10000,
		})
	})
	r.POST("/get/logtask", func(c *gin.Context) {
		var result =map[string]interface{}{"fp":logtask,"fs":1024,"ns":2}
		c.JSON(http.StatusOK,gin.H{"code":200,"data":result,"message":"test"})
	})
	r.POST("/update/userinfo",func(c *gin.Context){
		body,_ := ioutil.ReadAll(c.Request.Body)
		fmt.Printf("用户信息更新：%v\n",string(body))
		c.JSON(http.StatusOK,gin.H{"code":200,"message":"successful!"})
	})
	r.POST("/upload/logfile", func(c *gin.Context) {
		_=c.Query("di")
		body,_ := ioutil.ReadAll(c.Request.Body)
		fmt.Printf("上报文件长度：%v\n",len(body))
		c.JSON(http.StatusOK,gin.H{"code":200,"message":"successful!"})
	})
	c := cron.New()
	spec := "00 30 23 * *  "
	c.AddFunc(spec, func() {
		err:=Utils.ReadLine(fmt.Sprintf("./log/%v%v%v_ERR.log",time.Now().Year(),int(time.Now().Month()),time.Now().Day()))
		if err!=nil{
			fmt.Println(err)
		}
		Utils.EmailTo(fmt.Sprintf("upload次数:%v,network次数:%v,ViewEvent次数:%v,StateChangeEvent次数:%v,LaunchEvent次数:%v\n,LagEvent次数:%v,JSErrorEvent次数:%v,H5Event次数:%v,CustomMetricEvent次数:%v,CustomLogEvent次数:%v,CustomEventEvent次数:%v,CrashEvent次数:%v,ActionEvent次数:%v",Utils.UPLOAD,Utils.NetworkEvent,Utils.ViewEvent,Utils.StateChangeEvent,
		Utils.LaunchEvent,Utils.LagEvent,Utils.JSErrorEvent,Utils.H5Event,Utils.CustomMetricEvent,Utils.CustomLogEvent,Utils.CustomEventEvent,Utils.CrashEvent,Utils.ActionEvent))
		//初始化统计
		Utils.InitCount()
	})
	c.Start()
	r.Run("192.168.31.124:8080")
	}
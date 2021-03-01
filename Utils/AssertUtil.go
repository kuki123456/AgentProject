package Utils

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"time"
)
/*字段类型断言：
0：null
1：bool-false
2：number
3：string
4：bool-true
5：json结构块
*/
//值是否存在数组中判断，返回bool
//Target:判断是否存在列表的元素；EnsureArraary:要判断的列表
func Assertin(Target float64,EnsureArrary []float64)bool  {
   for _,value:=range EnsureArrary{
   	if value==Target{
   		return true
	}
   }
   return false
}
//类型判断
func AssertType(Actualtype gjson.Type,Expecttype gjson.Type,writer io.Writer,Actualvalue interface{},EnsureStr string,curbody string){
	if Actualtype==Expecttype && gjson.Get(curbody,EnsureStr).Exists() {
		log.Printf("%s类型是%v，值：%v\n",EnsureStr,Expecttype,Actualvalue)
	}else {
		_, _ = fmt.Fprintf(writer, "%v[ERROR]:%s的值类型错误：[当前:]%v,[期望:]%v,[当前值:]%v\n", time.Now().Format("2006/01/02 15:04:05"),EnsureStr,Actualtype,Expecttype,Actualvalue)
		EmailTo(fmt.Sprintf("%s捕捉到异常(字段类型不对或字段不存在)\n,%s的值类型错误：[当前:]%v,[期望:]%v,[当前值:]%v请求体如下，请排查日志:\n\n%s",time.Now().Format("2006/01/02 15:04:05"),EnsureStr,Actualtype,Expecttype,Actualvalue,curbody))
		}
}
//数值断言
func AssertNumValue(Actualvalue float64,Expectvalue float64,writer io.Writer,EnsureStr string,curbody string){
	if Actualvalue!=Expectvalue && gjson.Get(curbody,EnsureStr).Exists() {
		_, _ = fmt.Fprintf(writer, "%v[ERROR]:%s的值错误：[当前:]%v,[期望:]%v\n",time.Now().Format("2006/01/02 15:04:05"),EnsureStr,Actualvalue,Expectvalue)
		EmailTo(fmt.Sprintf("%s捕捉到异常(字段值不对)\n,%s的值错误：[当前:]%v,[期望:]%v\n请求体如下，请排查日志:\n\n%s",time.Now().Format("2006/01/02 15:04:05"),EnsureStr,Actualvalue,Expectvalue,curbody))
	}else {
		log.Printf("%s值是%v，值：%v\n",EnsureStr,Expectvalue,Actualvalue)
	}

}
//断言是否是bool值
func AssertBool(Actualtype gjson.Type,writer io.Writer,Actualvalue interface{},EnsureStr string,curbody string){
	if Actualtype==1 || Actualtype==4  {
		log.Printf("%s类型是%v，值：%v\n",EnsureStr,Actualtype,Actualvalue)
		}else {
		_, _ = fmt.Fprintf(writer, "%v[ERROR]:%s的值类型错误：[当前:]%v,[期望:]%v,[当前值:]%v\n", time.Now().Format("2006/01/02 15:04:05"),EnsureStr,Actualtype,"False or True",Actualvalue)
		EmailTo(fmt.Sprintf("%s捕捉到异常(字段类型不对或字段不存在)\n,%s的值类型错误：[当前:]%v,[期望:]%v,[当前值:]%v请求体如下，请排查日志:\n\n%s",time.Now().Format("2006/01/02 15:04:05"),EnsureStr,Actualtype,Actualtype,Actualvalue,curbody))
	}
}
//事件业务所对应的事件发生事件戳与用户信息，设备状态信息，网络信息
func AssertEventSin(BusinessBody string,file io.Writer,requestbody string,BusinessName string){
	//事件发生时刻ent校验
	ent:=gjson.Get(BusinessBody,"ent").Value()
	//fmt.Println(time.Unix(int64(ent.(float64)/1000/1000),0).Format("2006-01-02" ),time.Now().Format("2006-01-02"))
	if time.Unix(int64(ent.(float64)/1000/1000),0).Format("2006-01-02" )==time.Now().Format("2006-01-02"){
		fmt.Println("ent校验正确!")
	}else {
		_, _ = fmt.Fprintf(file, "%v[ERROR]:%s的ent上报异常，值为:%v\n", time.Now().Format("2006/01/02 15:04:05"),BusinessName,ent)
		EmailTo(fmt.Sprintf("%s--%s捕捉到异常(ent上报异常):\n,ent的值为:%v\n,请求体为:%s",time.Now().Format("2006/01/02 15:04:05"),BusinessName,ent,BusinessBody))

	}
	//sin索引值
	sin:=gjson.Get(BusinessBody,"sin").Value().([]interface{})
	for index,value:=range sin{
		//用户信息校验
		if index==0{
			if value==""{
				fmt.Println("user index is null string!")
			}else{
				if gjson.Get(requestbody,"ui."+value.(string)+".ui").Exists()&&gjson.Get(requestbody,"ui."+value.(string)+".ui").Type==3&&gjson.Get(requestbody,"ui."+value.(string)+".ui").Value()!="" || gjson.Get(requestbody,"ui."+value.(string)+".ei").Exists()&&gjson.Get(requestbody,"ui."+value.(string)+".ei").Type==3&&gjson.Get(requestbody,"ui."+value.(string)+".ei").Value()!=""{
					fmt.Printf("ui与ei校验成功,ui:%v,ei:%v!\n",gjson.Get(requestbody,"ui."+value.(string)+".ui"),gjson.Get(requestbody,"ui."+value.(string)+".ei"))
				}else {
					_, _ = fmt.Fprintf(file, "%v[ERROR]:ui上报异常,ui值为%v", time.Now().Format("2006/01/02 15:04:05"),gjson.Get(requestbody,"ui."+value.(string)))
					EmailTo(fmt.Sprintf("%s--%s对应的用户信息校验异常,请求体为:\n%s",time.Now().Format("2006/01/02 15:04:05"),BusinessName))
				}
			}
		}
		//设备状态信息校验
		if index==1{
			//suc
			AssertType(gjson.Get(requestbody,"ds."+value.(string)+".suc").Type,2,file,gjson.Get(requestbody,"ds."+value.(string)+".suc"),"ds."+value.(string)+".suc",requestbody)
			//auc
			AssertType(gjson.Get(requestbody,"ds."+value.(string)+".auc").Type,2,file,gjson.Get(requestbody,"ds."+value.(string)+".auc"),"ds."+value.(string)+".auc",requestbody)
			//aura
			AssertType(gjson.Get(requestbody,"ds."+value.(string)+".aura").Type,2,file,gjson.Get(requestbody,"ds."+value.(string)+".aura"),"ds."+value.(string)+".aura",requestbody)
			//sab
			AssertType(gjson.Get(requestbody,"ds."+value.(string)+".sab").Type,2,file,gjson.Get(requestbody,"ds."+value.(string)+".sab"),"ds."+value.(string)+".sab",requestbody)
			//saro
			AssertType(gjson.Get(requestbody,"ds."+value.(string)+".saro").Type,2,file,gjson.Get(requestbody,"ds."+value.(string)+".saro"),"ds."+value.(string)+".saro",requestbody)
			//sara
			AssertType(gjson.Get(requestbody,"ds."+value.(string)+".sara").Type,2,file,gjson.Get(requestbody,"ds."+value.(string)+".sara"),"ds."+value.(string)+".sara",requestbody)
			//ot
			AssertNumValue(gjson.Get(requestbody,"ds."+value.(string)+".ot").Value().(float64),1,file,"ds."+value.(string)+".ot",requestbody)
			//s
			AssertType(gjson.Get(requestbody,"ds."+value.(string)+".s").Type,2,file,gjson.Get(requestbody,"ds."+value.(string)+".s"),"ds."+value.(string)+".s",requestbody)
		}
		//网络状态信息校验
		if index==2{
			if gjson.Get(requestbody,"nsi."+value.(string)+".ns").Raw!="NaN" {
				AssertType(gjson.Get(requestbody, "nsi."+value.(string)+".ns").Type, 3, file, gjson.Get(requestbody, "nsi."+value.(string)+".ns"), "nsi."+value.(string)+".ns", requestbody)
				AssertType(gjson.Get(requestbody, "nsi."+value.(string)+".dip").Type, 3, file, gjson.Get(requestbody, "nsi."+value.(string)+".dip"), "nsi."+value.(string)+".dip", requestbody)
				AssertType(gjson.Get(requestbody, "nsi."+value.(string)+".dsi").Type, 3, file, gjson.Get(requestbody, "nsi."+value.(string)+".dsi"), "nsi."+value.(string)+".dsi", requestbody)
			}else {
				fmt.Println("网络状态信息为无网！")
			}
		}

	}
}
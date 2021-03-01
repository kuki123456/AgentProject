package EventBusiness

import (
	"AgentProject/Utils"
	"github.com/tidwall/gjson"
	"io"
)

func H5EventBusiness(H5Eventbody string,requestbody string,file io.Writer){
	//sin对应的索引json结构校验及事件发生时刻时间戳
	Utils.AssertEventSin(H5Eventbody, file, requestbody, "H5Event")
	//pvid页面id
	Utils.AssertType(gjson.Get(H5Eventbody,"v.pvid").Type,3,file,gjson.Get(H5Eventbody,"v.pvid").Value(),"v.pvid",H5Eventbody)
   //url,request请求地址url
	Utils.AssertType(gjson.Get(H5Eventbody,"v.url").Type,3,file,gjson.Get(H5Eventbody,"v.url").Value(),"v.url",H5Eventbody)
	//"wpi":{},//WebviewPerformanceInfo 页面性能数据 [非空字段]
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi").Type,5,file,gjson.Get(H5Eventbody,"v.wpi").Value(),"v.wpi",H5Eventbody)
	//{
	//	"ns":1,//navigation start [时间戳] (表征了从同一个浏览器上下文的上一个文档卸载(unload)结束时的UNIX时间戳。如果没有上一个文档，这个值会和PerformanceTiming.fetchStart相同)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.ns").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.ns").Value(),"v.wpi.ns",H5Eventbody)
	//    "ues":1,//unload event start [相对值] (表征了unload事件抛出时的UNIX时间戳。如果没有上一个文档，or if the previous document, or one of the needed redirects, is not of the same origin, 这个值会返回0)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.ues").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.ues").Value(),"v.wpi.ues",H5Eventbody)
	//    "uee":1,//unload event end [相对值] (表征了unload事件处理完成时的UNIX时间戳。如果没有上一个文档，or if the previous document, or one of the needed redirects, is not of the same origin, 这个值会返回0)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.uee").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.uee").Value(),"v.wpi.uee",H5Eventbody)
	//    "rds":1,//redirect start [相对值] (表征了第一个HTTP重定向开始时的UNIX时间戳。如果没有重定向，或者重定向中的一个不同源，这个值会返回0)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.rds").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.rds").Value(),"v.wpi.rds",H5Eventbody)
	//    "rde":1,//redirec end [相对值] (表征了最后一个HTTP重定向完成时（也就是说是HTTP响应的最后一个比特直接被收到的时间）的UNIX时间戳。如果没有重定向，或者重定向中的一个不同源，这个值会返回0)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.rde").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.rde").Value(),"v.wpi.rde",H5Eventbody)
	//    "fs":1,//fetch start [相对值] (表征了浏览器准备好使用HTTP请求来获取(fetch)文档的UNIX时间戳。这个时间点会在检查任何应用缓存之前)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.fs").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.fs").Value(),"v.wpi.fs",H5Eventbody)
	//    "dls":1,//domain lookup start [相对值] (表征了域名查询开始的UNIX时间戳。如果使用了持续连接(persistent connection)，或者这个信息存储到了缓存或者本地资源上，这个值将和 PerformanceTiming.fetchStart一致)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.dls").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.dls").Value(),"v.wpi.dls",H5Eventbody)
	//    "dle":1,//domain lookup end [相对值] (表征了域名查询结束的UNIX时间戳。如果使用了持续连接(persistent connection)，或者这个信息存储到了缓存或者本地资源上，这个值将和 PerformanceTiming.fetchStart一致)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.dle").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.dle").Value(),"v.wpi.dle",H5Eventbody)
	//    "cs":1,//connect start [相对值] (返回HTTP请求开始向服务器发送时的Unix毫秒时间戳。如果使用持久连接（persistent connection），则返回值等同于fetchStart属性的值)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.cs").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.cs").Value(),"v.wpi.cs",H5Eventbody)
	//    "scs":1,//secure connection start [相对值] (返回浏览器与服务器开始安全链接的握手时的Unix毫秒时间戳。如果当前网页不要求安全连接，则返回0)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.scs").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.scs").Value(),"v.wpi.scs",H5Eventbody)
	//    "ce":1,//connect end [相对值] (返回浏览器与服务器之间的连接建立时的Unix毫秒时间戳。如果建立的是持久连接，则返回值等同于fetchStart属性的值。连接建立指的是所有握手和认证过程全部结束)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.ce").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.ce").Value(),"v.wpi.ce",H5Eventbody)
	//    "reqs":1,//request start [相对值] (返回浏览器向服务器发出HTTP请求时（或开始读取本地缓存时）的Unix毫秒时间戳)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.reqs").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.reqs").Value(),"v.wpi.reqs",H5Eventbody)
	//    "rsps":1,//response start [相对值] (返回浏览器从服务器收到（或从本地缓存读取）第一个字节时的Unix毫秒时间戳。如果传输层在开始请求之后失败并且连接被重开，该属性将会被数制成新的请求的相对应的发起时间)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.rsps").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.rsps").Value(),"v.wpi.rsps",H5Eventbody)
	//    "rspe":1,//response end [相对值] (返回浏览器从服务器收到（或从本地缓存读取，或从本地资源读取）最后一个字节时（如果在此之前HTTP连接已经关闭，则返回关闭时）的Unix毫秒时间戳)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.rspe").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.rspe").Value(),"v.wpi.rspe",H5Eventbody)
	//    "dl":1,//dom loading [相对值] (返回当前网页DOM结构开始解析时（即Document.readyState属性变为"loading"、相应的 readystatechange事件触发时）的Unix毫秒时间戳)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.dl").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.dl").Value(),"v.wpi.dl",H5Eventbody)
	//    "di":1,//dom interactive [相对值] (返回当前网页DOM结构结束解析、开始加载内嵌资源时（即Document.readyState属性变为"interactive"、相应的readystatechange事件触发时）的Unix毫秒时间戳)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.di").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.di").Value(),"v.wpi.di",H5Eventbody)
	//    "dcles":1,//dom content loaded event start [相对值] (返回当解析器发送DOMContentLoaded 事件，即所有需要被执行的脚本已经被解析时的Unix毫秒时间戳)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.dcles").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.dcles").Value(),"v.wpi.dcles",H5Eventbody)
	//    "dclee":1,//dom content loaded event end [相对值] (返回当所有需要立即执行的脚本已经被执行（不论执行顺序）时的Unix毫秒时间戳)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.dclee").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.dclee").Value(),"v.wpi.dclee",H5Eventbody)
	//    "dc":1,//dom complete [相对值] (返回当前文档解析完成，即Document.readyState 变为 'complete'且相对应的readystatechange 被触发时的Unix毫秒时间戳)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.dc").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.dc").Value(),"v.wpi.dc",H5Eventbody)
	//    "les":1,//load event start [相对值] (返回该文档下，load事件被发送时的Unix毫秒时间戳。如果这个事件还未被发送，它的值将会是0)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.les").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.les").Value(),"v.wpi.les",H5Eventbody)
	//    "lee":1,//load event end [相对值] (返回当load事件结束，即加载事件完成时的Unix毫秒时间戳。如果这个事件还未被发送，或者尚未完成，它的值将会是0)
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wpi.lee").Type,2,file,gjson.Get(H5Eventbody,"v.wpi.lee").Value(),"v.wpi.lee",H5Eventbody)
	//}
	// "wri":[{}],//WebviewResourceInfo 页面资源数据 [非必要字段]
	Utils.AssertType(gjson.Get(H5Eventbody,"v.wri").Type,5,file,gjson.Get(H5Eventbody,"v.wri").Value(),"v.wri",H5Eventbody)
	//"WebviewResourceInfo": //该结构内时间单位统一为ms
	//{
	//	"st":1,//start time 发生时间 [时间戳，非空字段]
	//   "rt":"",//resource type 资源类型 [非空字段]
	//   "name":"",//name 资源名称 [非空字段]
	//   "dura":1,//duration 加载时间[相对时间，非空字段]
	//   "fs":1,//fetch start [相对时间，非空字段] 为浏览器已经准备好去使用HTTP请求抓取文档之时的 Unix毫秒时间戳。这一时刻在检查应用的缓存之前。
	//   "dls":1,//domain lookup start [相对时间，非空字段] 为域名开始解析之时的 Unix毫秒时间戳
	//   "dle":1,//domain lookup end [相对时间，非空字段] 为解析域名结束时的 Unix毫秒时间戳
	//   "cs":1,//connect start [相对时间，非空字段] 请求连接被发送到网络之时的Unix毫秒时间戳。如果传输层报告错误并且连接的建立重新开始，则把最后建立连接的开始时间作为该值。
	//   "ce":1,//connect end [相对时间，非空字段] 它以毫秒为单位，代表了网络链接建立的时间节点。如果传输层报告了错误或者链接又被重新建立，则采用最后一次链接建立的时间。如果链接是长久的，那么这个值等同于PerformanceTiming.fetchStart。
	//   "scs":1,//secure connection start [相对时间，非空字段] 为安全连接握手开始的时刻的 Unix毫秒时间戳
	//   "reqs":1,//request start [相对时间，非空字段] 为浏览器发送从服务器或者缓存获取实际文档的请求之时的 Unix毫秒时间戳
	//   "rsps":1,//response start [相对时间，非空字段] 为浏览器从服务器、缓存或者本地资源接收到响应的第一个字节之时的 Unix毫秒时间戳。
	//   "rspe":1,//response end [相对时间，非空字段] 为浏览器从服务器、缓存或者本地资源接收响应的最后一个字节或者连接被关闭之时的 Unix毫秒时间戳。
	//   "ts":1,//transfer size [相对时间，非空字段] the size (in octets) of the fetched resource. The size includes the response header fields plus the response payload body
	//   "ebs":1,//encoded body size [相对时间，非空字段] the size (in octets) received from the fetch (HTTP or cache), of the payload body, before removing any applied content-codings.
	//   "dbs":1,//decoded body size [相对时间，非空字段] the size (in octets) received from the fetch (HTTP or cache) of the message body, after removing any applied content-codings.
	//}


}
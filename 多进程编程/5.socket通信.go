package main

import (
	"sync"
	"strings"
	"fmt"
	"strconv"
	"math"
	"net"
	"bytes"
	"time"
	"io"
)

const (
	SEVER_NETWORK="tcp"
	SEVER_ADDRESS="127.0.0.1:8848"
	DELIMITER='\t'
)
var wg sync.WaitGroup//批量等待
//i
//i*i


//打印日志
func printlog(role string,sn int ,format string ,args...interface{}){
	if !strings.HasSuffix(format,"\n"){
		format+="\n"
	}
	fmt.Printf("%s[%d]:%s",role,sn,fmt.Sprintf(format,args...))
}

func PrintCLientLog(id int,format string ,args...interface{}){
	printlog("client",id ,format,args)
}
func PrintSeverLog(format string ,args...interface{}){
	printlog("server",0,format,args)
}

func strToInt32(str string)(int32,error) {
	num, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0, err
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0,err
	}
	return int32(num),nil
}
//处理m³
func CBrt(data int32)float64{
	return math.Cbrt(float64(data))
}

func Read(conn net.Conn)(string,error){
	readbytes:=make([]byte,1)
	var buffer bytes.Buffer
	for{
		_,err:=conn.Read(readbytes)
		if err!=nil{
			return "",err
		}
		readbyte:=readbytes[0]
		if readbyte==DELIMITER{
			break
		}
		buffer.WriteByte(readbyte)
	}
	return buffer.String(),nil
}
func Write(conn net.Conn,content string)(int,error){
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(DELIMITER)//分隔符
	return conn.Write(buffer.Bytes())
}
//Server处理链接，s收到一个整数，返回立方
func HandLeConn(conn net.Conn){
	defer func() {
		conn.Close()
		wg.Done()
	}()
	for {
		conn.SetReadDeadline(time.Now().Add(time.Second*10))//最多等待10秒
		strreq,err:=Read(conn)
		if err!=nil{
			if err==io.EOF{
				PrintSeverLog("客户端关闭")
			}else{
				PrintSeverLog("读取错误",err)
			}
			break
		}
		PrintSeverLog("收到%s\n",conn.RemoteAddr(),strreq)
		intreq,err:=strToInt32(strreq)
		if err!=nil{
			n,err:=Write(conn,err.Error())//返回错误信息
			PrintSeverLog("服务器错误",n,err)
			continue
		}
	//	floatdata:=CBrt(intreq)//处理立方
		floatdata:=float64(intreq*intreq*intreq)
		respMsg:=fmt.Sprintf("%d->%f",intreq,	floatdata)
		n,err:=Write(conn,respMsg)//返回错误信息
		if err!=nil{
			PrintSeverLog("发送",n,err)
			continue
		}
		fmt.Println("整数",intreq,	"数据",floatdata)
	}
}

func SeverGo(){
	var listener net.Listener
	listener,err:=net.Listen(SEVER_NETWORK,SEVER_ADDRESS)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer listener.Close()//延迟关闭
	fmt.Println("服务器开启",listener.Addr())
	for{
		conn,err:=listener.Accept()//接收链接
		if err!=nil{
			PrintSeverLog("%s",err)
		}
		PrintSeverLog("有客户加入",conn.RemoteAddr())
		go HandLeConn(conn ) //为每一个客户端开启一个线程
	}



}
func ClientGo(id int){
	defer wg.Done()
	conn,err:=net.DialTimeout(SEVER_NETWORK,SEVER_ADDRESS,time.Second*3)
	if err!=nil{
		PrintCLientLog(id,"%s",err)
		return
	}
	defer conn.Close()

	fmt.Printf("%d, client%s -> server %s \n ",id,conn.LocalAddr(),conn.RemoteAddr())
	time.Sleep(3)
	conn.SetDeadline(time.Now().Add(5*time.Second))

	//压入5个数据
	for i:=0;i<5;i++{
		n,err:=Write(conn,fmt.Sprintf("%d",i))
		if err!=nil{
			PrintCLientLog(id,"%s",err)
			continue
		}
		//PrintCLientLog(id,"%d,%d",n,i)
		fmt.Println(n,"客户端压入",id,"数据",i)
	}

	for i:=0;i<5;i++{
		strresp,err:=Read(conn)
		if err!=nil{
			if err==io.EOF{
				//PrintCLientLog(id,"%d  close",id)
			}else{
				//PrintCLientLog(id,"%d  %s",id,err)
			}
			break
		}
		//PrintCLientLog(id,"%d  %s",id,strresp)
		fmt.Println("客户端读取",id,"数据",strresp)
	}




}
func main(){
	wg.Add(3)
	go SeverGo()
	time.Sleep(time.Second*2)
	go ClientGo(1)
	go ClientGo(2)
	wg.Wait()
}
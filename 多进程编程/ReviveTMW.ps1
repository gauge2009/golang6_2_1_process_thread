#
#
#  描述：持续部署组件——水晶台监控脚本-reviver复活者
#  命名：CheckAndReviveTaskScheduler_x86
#  上次更新：2021-11-21
#
#
#
#备注： 1）请以管理员身份运行此脚本 
#       2）dev环境可采用 set-executionpolicy remotesigned 实现脚本运行权限，生产环境务必采用签名脚本！！
#       3）推荐安装路径：D:\HRLinkMain\HRLinkCal\CrystalBeacon\Reviver\


#1  脚本形参
Param(
    [string]$servicename, #新增无*，修改要*
   
    [string]$bit_type #x86 x64 
)
#2.1  输入参数校验
if($servicename   -eq "") {
    $servicename  = "HRLink.Worker.Node1" 
} 
 

set-executionpolicy remotesigned
#cd  $Source
#$servicename = "HRLink.TaskScheduler"
Write-Host  "正在重启$servicename......"   -ForegroundColor yellow 
Stop-Service -Name $servicename
Start-Service  -Name $servicename


 
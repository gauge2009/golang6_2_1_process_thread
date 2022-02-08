Param(
    [string]$servicename,  
   
    [string]$bit_type  
)
if($servicename   -eq "") {
    $servicename  = "HRLink.Worker.Node1" 
} 
set-executionpolicy remotesigned
Write-Host  "restart $servicename......"   -ForegroundColor yellow 
Stop-Service -Name $servicename
Start-Service  -Name $servicename
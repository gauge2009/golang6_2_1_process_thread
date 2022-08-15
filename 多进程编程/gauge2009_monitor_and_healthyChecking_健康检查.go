package main

import "gmonitor"

func main() {

	gmonitor.MoniorMemory()
	//gmonitor.MoniorCPU()
	gmonitor.MoniorDisk()
	gmonitor.MoniorHost()
	gmonitor.MoniorWindowsService()
	gmonitor.MoniorProcess()
	//gmonitor.MoniorMemory()
	//gmonitor.MoniorMemory()

}

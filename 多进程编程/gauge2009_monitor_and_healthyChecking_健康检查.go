package main

import "gmonitor"

func main() {

	gmonitor.MoniorMemory()
	//gmonitor.MoniorCPU()
	gmonitor.MoniorDisk()
	gmonitor.MoniorHost()
	gmonitor.MoniorProcess()
	gmonitor.MoniorWindowsService()
	//gmonitor.MoniorMemory()
	//gmonitor.MoniorMemory()

}

package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"golang.org/x/sys/windows"
	"log"

	//"errors"
	"unsafe"
)

func main() {
	CheckMemory()
	CheckMDiskUsage("C:")
	CheckMDiskUsage("D:")
	CheckMDiskUsage("R:")
	//DiskUsage("D:")
}

func CheckMemory() {
	v, _ := mem.VirtualMemory()

	fmt.Printf("Total: %v, Available: %v, UsedPercent:%f%%\n", v.Total, v.Available, v.UsedPercent)

	fmt.Println(v)
}
func CheckMDiskUsage(path string) {
	v, _ := disk.Usage(path)

	fmt.Printf("Total: %v, Available: %v, UsedPercent:%f%%\n", v.Total, v.Used, v.UsedPercent)

	fmt.Println(v)
}

type DiskStatus struct {
	All  uint64
	Used uint64
	Free uint64
}

func DiskUsage(path string) (disk DiskStatus) {
	h := windows.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")
	lpFreeBytesAvailable := uint64(0)
	lpTotalNumberOfBytes := uint64(0)
	lpTotalNumberOfFreeBytes := uint64(0)
	r1, r2, err := c.Call(uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("C:"))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)))
	//failOnError(err,"Failed to check FreeBytesAvailable")
	if err != nil {
		log.Fatalf("channel.open: %s", err)
	}
	fmt.Println(r1)
	fmt.Println(r2)
	disk.All = lpTotalNumberOfBytes
	disk.Free = lpTotalNumberOfFreeBytes
	disk.Used = lpFreeBytesAvailable
	return
}

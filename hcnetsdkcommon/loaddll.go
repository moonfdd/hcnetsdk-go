package hcnetsdkcommon

import (
	"sync"

	"github.com/ying32/dylib"
)

var hcnetsdkDll *dylib.LazyDLL
var hcnetsdkDllOnce sync.Once

func GetHcnetsdkDll() (ans *dylib.LazyDLL) {
	hcnetsdkDllOnce.Do(func() {
		hcnetsdkDll = dylib.NewLazyDLL(hcnetsdkPath)
	})
	ans = hcnetsdkDll
	return
}

var hcnetsdkPath = "hcnetsdk.dll"

func SetHcnetsdkPath(path0 string) {
	hcnetsdkPath = path0
}

package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"strings"
)

const authHeader = "x-pal-auth"

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

type vmContext struct {
	// Embed the default VM context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultVMContext
}

// Override types.DefaultVMContext.
func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginDemo{}
}

type pluginDemo struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultPluginContext
	contextID uint32
}

// Override types.DefaultPluginContext.
func (ctx *pluginDemo) OnPluginStart(pluginConfigurationSize int) types.OnPluginStartStatus {
	proxywasm.LogInfo("OnPluginStart from Go!")
	return types.OnPluginStartStatusOK
}

// Override types.DefaultPluginContext.
func (*pluginDemo) NewHttpContext(contextID uint32) types.HttpContext {
	return &httpAuthRandom{contextID: contextID}
}

type httpAuthRandom struct {
	// Embed the default http context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultHttpContext
	contextID uint32
}

// Override types.DefaultHttpContext.
func (ctx *httpAuthRandom) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	hs, err := proxywasm.GetHttpRequestHeaders()
	if err != nil {
		proxywasm.LogCriticalf("failed to get request headers: %v", err)
		return types.ActionPause
	}

	for _, h := range hs {
		proxywasm.LogInfof("request header <-- %s: %s ", h[0], h[1])
		if strings.ToLower(h[0]) == authHeader && !isPalindrome(h[1]) {
			proxywasm.LogErrorf("auth failure: '%s' is not a palindrome", h[1])

			body := "access forbidden"
			if err := proxywasm.SendHttpResponse(403, [][2]string{
				{"powered-by", "proxy-wasm-go-sdk!!"},
			}, []byte(body), -1); err != nil {
				proxywasm.LogCriticalf("failed to send local response: %v", err)
				return types.ActionPause
			}
		}
	}

	proxywasm.LogInfof("http call allowed")
	return types.ActionContinue
}

func isPalindrome(input string) bool {
	i, j := 0, len(input)-1
	for i < j {
		// Convert both characters to lowercase before comparison
		if input[i]|0x20 != input[j]|0x20 {
			return false
		}
		i++
		j--
	}
	return true
}

package proxy

import (
	"log"
	"net/http/httputil"
	"net/url"

	"github.com/pdxjohnny/discovery/random"
)

type ProxyManager interface {
	Add(name string)
	Random() *httputil.ReverseProxy
}

type BaseProxyManager struct {
	ProxyMap  map[string]*httputil.ReverseProxy
	ProxyList []string
}

func NewBaseProxyManager() *BaseProxyManager {
	return &BaseProxyManager{
		ProxyMap:  make(map[string]*httputil.ReverseProxy, 0),
		ProxyList: make([]string, 0),
	}
}

func (proxy *BaseProxyManager) Random() *httputil.ReverseProxy {
	index := random.Range(0, len(proxy.ProxyList) - 1)
  indexUrl := proxy.ProxyList[index]
  return proxy.ProxyMap[indexUrl]
}

func (proxy *BaseProxyManager) Add(addUrl string) {
	remote, err := url.Parse(addUrl)
	if err != nil {
		log.Println("BaseProxyManager.Add:\tERROR parsing url:\t", err)
	}
	proxy.ProxyMap[addUrl] = httputil.NewSingleHostReverseProxy(remote)
	proxy.ProxyList = append(proxy.ProxyList, addUrl)
}

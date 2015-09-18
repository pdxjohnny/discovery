package proxy

import (
	"log"
	"net/http/httputil"
	"net/url"

	"github.com/pdxjohnny/discovery/discovery"
	"github.com/pdxjohnny/discovery/random"
)

type Manager interface {
	Add(name string)
	Random() *httputil.ReverseProxy
	Discover(service discovery.Service, addr, port string)
}

type BaseManager struct {
	ProxyMap  map[string]*httputil.ReverseProxy
	ProxyList []string
}

func NewBaseManager() *BaseManager {
	return &BaseManager{
		ProxyMap:  make(map[string]*httputil.ReverseProxy, 0),
		ProxyList: make([]string, 0),
	}
}

func (proxy *BaseManager) Random() *httputil.ReverseProxy {
	index := random.Range(0, len(proxy.ProxyList)-1)
	indexUrl := proxy.ProxyList[index]
	return proxy.ProxyMap[indexUrl]
}

func (proxy *BaseManager) Add(addUrl string) {
	remote, err := url.Parse(addUrl)
	if err != nil {
		log.Println("ERROR: BaseManager.Add parsing url: ", err)
	}
	// Check if url is already in proxy map
	_, ok := proxy.ProxyMap[addUrl]
	if ok {
		return
	}
	// If it is not then add it
	log.Println("LOG: BaseManager.Add: Added web host", addUrl)
	proxy.ProxyMap[addUrl] = httputil.NewSingleHostReverseProxy(remote)
	proxy.ProxyList = append(proxy.ProxyList, addUrl)
}

func (proxy *BaseManager) Discover(service discovery.Service, addr, port string) {
	discovery.Listen(service, addr, port)
}

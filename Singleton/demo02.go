/**
@auther:liuzhijun
@data:2023/3/12
**/
package Singleton

import "sync"

//懒汉式单例
var (
	lazySingleton *Singleton
	once          sync.Once
)

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return singletion
}

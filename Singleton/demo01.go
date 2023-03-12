/**
@auther:liuzhijun
@data:2023/3/12
**/
package Singleton

//Singleton 饿汉式单例
type Singleton struct {
}

var singletion *Singleton

func init() {
	singletion = &Singleton{}
}

//GetSingletion 获取实例
func GetSingletion() *Singleton {
	return singletion
}

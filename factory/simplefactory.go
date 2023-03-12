/**
@auther:liuzhijun
@data:2023/3/12
**/
package factory

import "fmt"

//工厂-->创建相应的对象

type IRuleConfigParser interface {
	Parse(data []byte)
}
type jsonRuleConfigParser struct {
}

func (j *jsonRuleConfigParser) Parse(data []byte) {
	fmt.Println("json rule config parser")
}

type yamlRuleConfigParser struct {
}

func (j *yamlRuleConfigParser) Parse(data []byte) {
	fmt.Println("yaml rule config parser")
}

type xmlRuleConfigParser struct {
}

func (j *xmlRuleConfigParser) Parse(data []byte) {
	fmt.Println("xml rule config parser")
}

func newIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "xml":
		return &xmlRuleConfigParser{}
	case "jsom":
		return &jsonRuleConfigParser{}
	case "yaml":
		return &yamlRuleConfigParser{}
	default:
		return nil
	}
}

/**
@auther:liuzhijun
@data:2023/3/12
**/
package factory

//工厂-->创造相应的对象工厂-->创造相应的对象

type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

type yamlRuleConfigParseFactory struct {
}

func (y *yamlRuleConfigParseFactory) CreateParser() IRuleConfigParser {
	return &yamlRuleConfigParser{}
}

type jsonRuleConfigParseFactory struct {
}

func (y *jsonRuleConfigParseFactory) CreateParser() IRuleConfigParser {
	return &jsonRuleConfigParser{}
}

type xmlRuleConfigParseFactory struct {
}

func (y *xmlRuleConfigParseFactory) CreateParser() IRuleConfigParser {
	return &xmlRuleConfigParser{}
}

func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "xml":
		return &xmlRuleConfigParseFactory{}
	case "jsom":
		return &jsonRuleConfigParseFactory{}
	case "yaml":
		return &yamlRuleConfigParseFactory{}
	default:
		return nil
	}
}

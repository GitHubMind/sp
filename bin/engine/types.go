package engine

//请求借口以及 过滤的方法
type Request struct {
	Url       string
	ParserFun func([]byte, string) (ParseResult, error)
	Pre       string //记录上一个返回的id
}

const HostUrl = "http://www.quanben5.com"
const HostUrlPen = "https://www.qbiqu.com"

type ParseResult struct {
	//继续返回请求
	Requests []Request
	//interface{}可以做任何类型
	Items []interface{}
}

func NilParser([]byte, string) ParseResult {
	return ParseResult{}
}

func NilParseResult(r Request) ParseResult {
	RequestsArr := []Request{r}
	//item :=nil
	return ParseResult{RequestsArr, nil}
}

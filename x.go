package main;
 
import (
    "regexp"
    "fmt"
)
 
func main() {
 
    //2、解析正则表达式
    //正则表达式如果合法，Compile会返回一个Regexp对象指针，通过该指针可以在任意字符串上进行操作
    re, _ := regexp.Compile("a(.*?)m");
 
    //3、查找正则匹配的字符串
    data := "I am a good man";
 
    //Find函数返回匹配的第一个字符串
    one := re.Find([]byte(data));
    fmt.Println(string(one));
 
    //FindAll函数返回匹配的所有字符串，n小于0返回全部字符串，否则返回指定长度
    all := re.FindAll([]byte(data), 2);
    //all为长度为2的slice
    fmt.Println(string(all[0]));
	fmt.Println(string(all[1]));
	
	flysnowRegexp := regexp.MustCompile(`^http://www.flysnow.org/(.*?).html$`)
	params := flysnowRegexp.FindStringSubmatch("http://www.flysnow.org/2018/01/20/golang-goquery-examples-selector.html")

	for _,param :=range params {fmt.Println(param)}

 
    
}
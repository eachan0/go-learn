package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	//fmtPackage()
	//timePackage()
	//flagPackage()
	//logPackage()
	strconvPackage()
}

// bytes,字节操作,相应的还有 strings 等
func bytesPackage() {
	// bytes 包
	arrA := [...]byte{97, 98, 99}
	arrB := [...]byte{65, 66, 67}
	sA := arrA[:]
	sB := arrB[:]
	fmt.Println("EqualFold", bytes.EqualFold(sA, sB))
	prefixA := [1]byte{65}
	prefix := prefixA[:]
	fmt.Println("HasPrefix", bytes.HasPrefix(sA, prefix))
	fmt.Println("ContainsAny", bytes.ContainsAny(sA, "cde"))
}

// 基本数据类型与其字符串表示的转换
func strconvPackage() {

	// Atoi()函数用于将字符串类型的整数转换为int类型
	// A/a的典故,这是C语言遗留下的典故。C语言中没有string类型而是用字符数组(array)表示字符串
	i, err := strconv.Atoi("s")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("type:%T value:%#v\n", i, i) //type:int value:100
	}

	// Itoa()函数用于将int类型数据转换为对应的字符串表示
	s := strconv.Itoa(200)
	fmt.Printf("type:%T value:%#v\n", s, s) //type:string value:"200"

	// Parse系列函数,用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。
	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	// base指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；
	// bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
	i64, err := strconv.ParseInt("-2", 10, 64)
	u, err := strconv.ParseUint("2", 10, 64)
	fmt.Println(b, f, i64, u)

	// Format系列函数实现了将给定类型数据格式化为string类型数据的功能。
	s1 := strconv.FormatBool(true)
	// bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
	// fmt表示格式：’f’（-ddd.dddd）、’b’（-ddddp±ddd，指数为二进制）、’e’（-d.dddde±dd，十进制指数）、’E’（-d.ddddE±dd，十进制指数）、’g’（指数很大时用’e’格式，否则’f’格式）、’G’（指数很大时用’E’格式，否则’f’格式）。
	// prec控制精度（排除指数部分）：对’f’、’e’、’E’，它表示小数点后的数字个数；对’g’、’G’，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-2, 16)
	s4 := strconv.FormatUint(2, 16)
	fmt.Println(s1, s2, s3, s4)
}

// fmt
func fmtPackage() {
	// 向外输出
	// Print函数会将内容输出到系统的标准输出
	fmt.Print("在终端打印该信息。")
	// Printf函数支持格式化输出字符串
	name := "枯藤"
	fmt.Printf("我是：%s\n", name)
	// Println函数会在输出内容的结尾添加一个换行符。
	fmt.Println("在终端打印单独一行显示")
	// Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
	fileObj, err := os.Create("xx.txt")
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	} else {
		// 向打开的文件句柄中写入内容
		_, err := fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
		if err != nil {
			fmt.Println(err)
		}
	}
	// Sprint系列函数会把传入的数据生成并返回一个字符串。
	s1 := fmt.Sprint(name)
	s2 := fmt.Sprintf("name:%s,age:%d", name, 18)
	s3 := fmt.Sprintln(name)
	fmt.Println(s1, s2, s3)
	// Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。
	err = fmt.Errorf("这是一个错误")
	fmt.Println(err)

	/**
	通用占位符
	%v	值的默认格式表示
	%+v	类似%v，但输出结构体时会添加字段名
	%#v	值的Go语法表示
	%T	打印值的类型
	%%	百分号
	------bool占位符
	%t	true或false
	------整型
	%b	表示为二进制
	%c	该值对应的unicode码值
	%d	表示为十进制
	%o	表示为八进制
	%x	表示为十六进制，使用a-f
	%X	表示为十六进制，使用A-F
	%U	表示为Unicode格式：U+1234，等价于”U+%04X”
	%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	------浮点数与复数
	%b	无小数部分、二进制指数的科学计数法，如-123456p-78
	%e	科学计数法，如-1234.456e+78
	%E	科学计数法，如-1234.456E+78
	%f	有小数部分但无指数部分，如123.456
	%F	等价于%f
	%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
	--------字符串和[]byte
	%s	直接输出字符串或者[]byte
	%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
	%x	每个字节用两字符十六进制数表示（使用a-f
	%X	每个字节用两字符十六进制数表示（使用A-F）
	--------指针
	%p	表示为十六进制，并加上前导的0x
	--------宽度标识符
	%f	默认宽度，默认精度
	%9f	宽度9，默认精度
	%.2f	默认宽度，精度2
	%9.2f	宽度9，精度2
	%9.f	宽度9，精度0
	--------其他
	’+’	总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
	’ ‘	对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格
	’-’	在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
	’#’	八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的go字面值；
	‘0’	使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；
	*/

	// 获取输入
	var (
		title   string
		age     int
		married bool
	)
	//fmt.Scan(&title, &age, &married)
	//fmt.Printf("扫描结果 name:%s age:%d married:%t \n", title, age, married)
	//fmt.Scanf("1:%s 2:%d 3:%t", &title, &age, &married)
	//fmt.Printf("扫描结果 name:%s age:%d married:%t \n", title, age, married)
	fmt.Scanln(&title, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", title, age, married)
}

// crypto包,加密
func cryptoPackage() {
	// MD5
	str := "123456" // e10adc3949ba59abbe56e057f20f883e
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	fmt.Println(cipherStr)
	fmt.Println(hex.EncodeToString(cipherStr))
}

// 日期时间
func timePackage() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)

	// 使用time.Unix()函数可以将时间戳转为时间格式。
	timeObj := time.Unix(1631694640, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year = timeObj.Year()     //年
	month = timeObj.Month()   //月
	day = timeObj.Day()       //日
	hour = timeObj.Hour()     //小时
	minute = timeObj.Minute() //分钟
	second = timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	// 时间格式化
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

	// 解析字符串格式的时间
	// time.Parse("2006-01-02 15:04:05", "2021-09-15 16:30:40")
	// 按照指定时区和指定格式解析字符串时间
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	timeObj, err = time.ParseInLocation("2006-01-02 15:04:05", "2021-09-15 16:30:40", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))

	// 时间间隔 time.Duration, 时间操作
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println(later)

	/**
	Sub: 求两个时间之间的差值： func (t Time) Sub(u Time) Duration
	返回一个时间段t-u。如果结果超出了Duration可以表示的最大值/最小值，将返回最大值/最小值。要获取时间点t-d（d为Duration），可以使用t.Add(-d)。
	Equal:判断两个时间是否相同 func (t Time) Equal(u Time) bool
	会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息。
	Before: func (t Time) Before(u Time) bool
	如果t代表的时间点在u之前，返回真；否则返回假。
	After: func (t Time) After(u Time) bool
	如果t代表的时间点在u之后，返回真；否则返回假。
	*/

	// 定时器,使用time.Tick(时间间隔)来设置定时器，定时器的本质上是一个通道（channel）。
	/*ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i)//每秒都会执行的任务
	}*/
}

// 命令行参数解析
func flagPackage() {
	/**
	字符串flag	 合法字符串
	整数flag	     1234、0664、0x1234等类型，也可以是负数。
	浮点数flag	 合法浮点数
	bool类型flag	 1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False。
	时间段flag	 任何合法的时间段字符串。如”300ms”、”-1.5h”、”2h45m”。合法的单位有”ns”、”us” /“µs”、”ms”、”s”、”m”、”h”。
	*/

	// 定义命令行参数, 自带 --help
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}

// 日志
func logPackage() {
	/**
	logger会打印每条日志信息的日期、时间，默认输出到系统的标准错误。Fatal系列函数会在写入日志信息后调用os.Exit(1)。
	Panic系列函数会在写入日志信息后panic。
	*/
	log.Println("这是一条很普通的日志。")
	log.Printf("这是一条%s日志。\n", "很普通的")
	// log.Fatalln("这是一条会触发fatal的日志。")
	// log.Panicln("这是一条会触发panic的日志。")

	// 配置logger
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	// 配置日志前缀
	log.SetPrefix("[info]")
	log.Println("这是一条很普通的日志。")

	// 配置日志输出位置,通常会把上面的配置操作写到init函数中
	/*logFile, err := os.OpenFile("xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		log.SetOutput(logFile)
		log.Println("这是一条很普通的日志。但会写到文件中")
	}*/

	// 创建logger,log标准库中还提供了一个创建新logger对象的构造函数–New，支持我们创建自己的logger示例。
	logger := log.New(os.Stdout, "<error>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志。")
}

// init 函数备份,使用时去掉 _bak
func init_bak() {
	logFile, err := os.OpenFile("xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}

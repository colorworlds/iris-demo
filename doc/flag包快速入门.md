



### flag包快速入门  
flag包和os.Args功能差不多，相当于提供了更方便的命令行参数的解析工具。

使用`flag.Type()`定义一个命令行参数，比如下面定义了3个命令行参数：
```
name := flag.String("name", "小明", "姓名")
age := flag.Int("age", 18, "年龄")
married := flag.Bool("married", false, "婚否")  
```

可以看到 `flag.Typ()` 类型的方法有三个参数，这三个参数分别表示：  
1.命令行参数名。比如之后可以在运行程序时通过命令行中输入参数 `-name 小红` 来将变量 `name` 赋值为字符串“小红”。  
2.参数的默认值。如果在运行程序时如果没有为参数指定值，那么就将使用该默认值。比如程序运行后没有使用参数`-name 小红`，那么此时`name`的值就是“小明”。  
3.参数的描述。在运行程序时使用内置的`-help`能看到每个参数的描述。  


使用 `flag.Parse()` 来读取命令行参数。  
上面 `flag.Type()` 只是指定了每个参数如何读取，并把读取的值赋予哪个参数。而没有真正读到值。  
只有程序执行 `flag.Parse()` 后，上面的 `name`、`age`、`married` 三个变量才会有值。 

最终的程序可能是下面这样，比如叫flag_demo.go

```
func main() {
    //指定变量
    var name string
    var age int
    var married bool

    //为变量指定对应的命令行参数
    flag.StringVar(&name, "name", "小明", "姓名")
    flag.IntVar(&age, "age", 18, "年龄")
    flag.BoolVar(&married, "married", false, "婚否")


    //解析命令行参数，使得上面的变量获取值
    flag.Parse()
    
    //打印每个变量最终获得的值
    fmt.Println(name, age, married)
}
```

在命令行运行参数并使用-help指令：
```
$ ./flag_demo -help
```

能够看到：
```
Usage of ./flag_demo:
  -age int
        年龄 (default 18)
  -name string
        姓名 (default "小明")
  -married
        婚否
```

也可以像这样赋值：
```
$ ./flag_demo -name 小红 --age 16 -married=true
```
最终的打印结果：
```
小红 16 true
```


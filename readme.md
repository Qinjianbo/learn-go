1. 简单运行go文件

    go run helloworld.go

2. 生成可复用程序

    go build helloworld.go

3. 格式输出相关包

    fmt

4. 系统操作相关包

    os

5. 包引入

```
import (
    "fmt",
    "os"
)
```

6. 包声明

    package main

7. for循环 
```
for ini; condition; post {

}

// 死循环

for {

}
```

8. 变量声明及初始化

```
s := ""
var s string
var s = ""
var s string = ""
```

9. 字符串处理包

    strings

10. 类型转换包

    strconv

    ```
    # string => int
    int, err := strconv.Atoi(string)
    # string => int64
    int64, err := strconv.ParseInt(string, 10, 64)
    # int => string
    string := strconv.Itoa(int)
    # int64 => string
    string := strconv.FormatInt(int64, 10)
    ```

11. make 函数

```
    make 可以用来新建map
    counts := make(map[string]int)
```

12. 高效处理输入和输出的包 bufio

```
    input := bufio.NewScanner(os.Stdin)
    line := input.Text()
```

13. verb(转义字符)

```
%d                    十进制整数
%x, %o, %b            十六进制, 八进制, 二进制整数
%f, %g, %e            浮点数：如 3.1415926, 3.141592653589793, 3.141593e+00
%t                    布尔类型: true 或 false
%c                    字符(Unicode 码点)
%s                    字符串
%q                    带引号字符串(如 "abc") 或者字符(如 'c')
%v                    内置格式的任何值
%T                    任何值类型
%%                    百分号本身(无操作数)
```

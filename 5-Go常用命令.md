# 5、Go常用命令

## 1、go env

## 2、go build

## 3、go run

## 4、go install

## 5、go get

## 6、go test

运行当前目录下的tests
常用命令 
   go test:一般输出最终结果
   go test -v:输入详细信息

 必须引入testing包
import testing

 普通test文件规则：
   1、文件名规则：xxx_test.go  xxx不一定是要测试的文件名，默认规则
   2、要测试的函数(test case)，函数名规则：TestXxx,否则go test会跳过不执行
   3、函数参数必须含有 t *testing.T或者 b *testing.B最少一种类型
        testing.T ：测试，跑一遍
           A：t.Errorf()打印错误信息，并且当前函数会跳过执行  
           B：t.SkipNow() 跳过当前test case，并且直接按照PASS输出，必须放在函数第一行
           C：t.Run("测试函数标识",test case) 如果多个test case要按照某种顺序执行时使用t.Run：（go test不保证顺序执行）


    4、TestMain函数：初始化test，并且在函数内部调用m.Run()调用其他tests，完成一些初始化操作，比如初始化数据链接，如果有TestMain函数但没用调用m.Run()则本文件其他tests不执行
       函数原型
         func TestMain(m *testing.M){
           m.Run()
         }

  性能测试test case

     1、文件名规则：xxx_test.go  xxx不一定是要测试的文件名，默认规则
     2、函数名规则： BenchmarkXxx  一般以Benchmark开头
     3、Benchmark函数一般会跑b.N次，每次go test都会跑这么多次
     4、在执行过程中会根据test case函数每次执行时间是否稳定来增加b.N的次数，以达到稳态
     5、执行命令 go test -bench=. ，不带-bench=.是不会跑benchmark



## 7、go doc
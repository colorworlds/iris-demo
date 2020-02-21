# iris-demo
这个项目的目的是方便入门iris，所以会尽可能提供详细的解释。


这个项目的代码拷贝自：
https://github.com/wyanlord/golang-iris-web
最近刚入门go，最近一个月一直在用go写项目。
原本想要把自己的实践总结下，写个demo，不过刚好看到已经有大佬写了一个更加详细的例子程序（比我写的demo可好多了）。
所以干脆用大佬的项目做示例。顺便可以再学习下。




### 基于iris框架的demo
+ 基于IRIS框架,最主要的是合理的分配目录结构
+ 集成了validator参数验证工具
+ 集成了cache、mysql、redis、redis-locker等小工具
+ 集成了log按天分割，定期清理工具
+ 使用jwt进行权限认证
+ 使用etcd进行动态配置
+ 实现了一些helper方法

# golang的web开发之后端架构篇

&emsp;&emsp;闲来无事，玩玩golang，玩玩web了。计林一直盛传ruby的ruby on rails乃web开发之王，
未能有幸见之，python的小巧异步高性能web框架tornado到是有所了解，动态脚本语言开发的高效，
五花八门的三方库，让胶水之王的python在各个领域都占有一席之地。玩了一下golang之后，仅对比
web开发的后端，golang比python自我感觉确实有过之，而无不及呢，呵呵，各位python粉别喷我了。
当然语言孰优孰劣，本是一个伪命题，只有最强的人，而无最强的语言。技术的选型是需要考虑很多
东西的，也非仅仅开发语言所能代表的。废话少说了，言归正传，开始讨论讨论golang的web开发了。
</br   >传统的web开发分三层架构，用户界面层（ui层），业务逻辑层（business层）和数据层（db层），
可别问我为什么这么分，问了我也不知道。ui层归为前端，不在本文讨论范围之内，因为我也不会前端，
business层和db层归为后端，是本文讨论的重点。

&emsp;&emsp;既然是golang的web开发，当然得选个web开发框架，无须自己再造轮子了。选哪个好呢？这个我
也不知道呢，就选个gin呢！为什么选gin呢？网上都说gin的效率最高呢？于是乎我也就信了，那就拿
gin来练练手了。

&emsp;  ui层和business层采用http(s)通讯，进行请求，即business层采用gin框架。business层和
db层采用rpc通讯，进行请求，即db层采用rpc的服务端以及business层采用rpc的客户端。
&emsp;  golang的基础知识及如何搭建golang开发环境及web开发知识，请自行脑补，自行百度了。首
先开始business层的介绍。主函数main启动，读取配置文件，启动web服务，处理ui客户端的http
(s)请求，将所有的http(s)请求归为get，post，put，delete等，即对应httprequest包中的
getroute，postroute，putroute，putroute。每一种route管理着若干个handler，每一个
handler对应着一个业务处理函数，即若添加一个新功能，先判断该功能对应的是何种http(s)请求，
再在对应的route中增加一个对应的handler处理函数即可。源码见business中的httprequest包，
照着抄即可，无须太多的知识。如果涉及数据库的相关操作，就作为客户端向db层发请求即可。一句
话概括business层就是接受ui层客户端的请求，如果business自己能处理就处理，如果不能处理，
需要db层的相关数据，就将ui层客户端的请求在business层作为rpc的客户端向db层进行转发，
获取返回数据进行组装成预定的数据格式返回给ui层。详细的business的代码组织结构见business
目录，无须再多言了。
&emsp;  db层的介绍，db层主要是和数据库有关的相关操作，增删改查。将db层作为rpc的服务端运行，
为business层提供调用，熟悉各个数据库的操作api即可，没有什么可以说的了。
&emsp;  以上只是golang的web后端开发的一个简单原型的框架架构介绍，还有很多很多东西需要完善。
既然是玩golang，当然也应该体现golang的特色，后面可能会考虑采用grpc代替原生的rpc作为
business层和db层之间的通讯，采用微服务go-micro提供服务。当然在程序的具体写法中，也还应该
考虑golang的协程goroutine多么牛逼简洁的写法来支持大并发。
&emsp;  写在最后的话，本人这些年一直做的是小数据，低并发，低性能的开发，也不懂啥大数据，高
并发，高性能。嘴炮的架构师们在你们搭建好固若金汤的架构后，请你们也完善完善代码的组织结构
不要让新手有太多的入口去写一个功能，有且只有一处让其修补代码，保持代码风格的一直性何不乐
哉？

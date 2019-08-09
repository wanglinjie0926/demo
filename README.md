# demo
 2019年华为云鲲鹏开发者

1-C++ demo

这是xl调通的C++demo，然后只有初步功能，详情见里面得readme，有问题随时可以在群里交流或者这个帖子下回复我。

https://bbs.huaweicloud.com/forum/thread-22065-1-1.html

2019/8/4

分享一份跨windows和linux平台的C++demo一份出来

因为我虚拟机安装后网络有问题，然后云服务器配置逻辑局域网也失败。于是干脆写了一份跨平台的demo，方便C++的调试

只对网络通信进行了跨平台的封装，目前win下测试通过了，然后上传后跑分1分，详情见readme

2- golang-demo

由hadrianl大佬提供

https://bbs.huaweicloud.com/forum/thread-22099-1-1.html


3-python-demo 
由rysander大佬提供

python3.6+ 调试demo + 混分demo，下载上传直接拿分

=====以下是调试demo （本地调试时使用）=====

比赛没有给出Python3的demo，只能自己改Python2的代码，参考了泽胖胖QaQ的帖子
在他基础上我改了byte数据的判定格式和相对路径，目前运行正常正在大战AI……
下载附件Ballclient.zip后，覆盖掉ballclient文件夹即可开始调试。

要注意测试的时候可能需要修改run.bat文件的倒数第二段改为测试文件夹的目录
pushd %CD%
cd /d "demo_proj/demo_for_python"
start gameclient.bat 1112 127.0.0.1 6001 
popd

然后直接运行run.bat就可以，游戏界面会出来，只要蓝方队伍名是Team_name就成功了。

Window下就能调试！Windows下就能调试！Windows下就能调试！

=====以下是混分demo （上传混分时使用）=====

简单说就是提交的时候需要改动很多东西，比如改.sh文件改成python3运行，另外我pip3一直安不对，导致我把所有的numpy代码全删了…
不过好在给的demo程序不用怎么改就能用，反正随机算法也没啥可改…我传了五次，最高刷到了23分，而我自己辛辛苦苦写的算法分数才16分！

这比赛每天有一百次机会刷分，最后只取最高分，而这个算法又是随机算法。。。
那么。。。激(jian)动(yan)人(xue)心(tong)的时候到了！！！
请下载ython_demo.zip，改成自己名字直接上传，就能直接拿分！得多少分全凭天意

P S：程序里还埋了一处不容易发现的bug，debug出来的和de不出来的程序平均分数会相差5分左右，不过非常好找，用过Python3的人都能找到。
===============

编者再
PS：bug 就是有个print 函数里多了个小东西一眼就可以看到。

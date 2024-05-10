# 简介
本项目实现了 ZstuWlan 的多设备共享  

# 教程
## 路由器方案
### OpenWrt 路由器
1. 购买一个支持 OpenWrt 的路由器：[支持列表](http://www.ttcoder.cn/index.php/2024/04/14/support-openwrt-routers/)  
2. 为其刷上 OpenWrt 系统
3. 设置当中创建 Wifi  

相关教程有很多，不多做赘述，参考资料：[恩山论坛](https://www.right.com.cn/forum/forum.php)  

### 自动登录
上一步刷完 OpenWrt 并开启 Wifi 后，已可实现共享，连接路由器 Wifi，登录一次后，全部设备无需登录即可上网  

但是 12 点断网后路由器会失去验证，需要再次手动登录，比较麻烦，可以使用程序实现自动重连
#### 登录程序
ZSTU 使用的是锐捷 ePortal Web ，MenToHust 不可用，这里参考了 [RuijiePortalLoginTool](https://github.com/callmeliwen/RuijiePortalLoginTool) 项目

本项目使用了 Go 重写了 [RuijiePortalLoginTool](https://github.com/callmeliwen/RuijiePortalLoginTool) 的脚本，已验证可用，可在 release 根据路由器架构进行下载    
使用方法：  
```shell
./login 学号 密码
```

> 自行编译需要注意路由器的编译参数


#### 定时登录
OpenWrt 自带了 Cron 定时执行程序，将登录程序加入 Cron 定时任务，美团早上7点执行一次，[Cron 表达式工具](https://cron.ciding.cc/)

## 手机方案
实际上直接使用手机连接 ZstuWlan 再开热点就能实现共享，只是这样不方便和舍友共享网络


# 问题和建议
有问题或建议可以在 [Issues ](https://github.com/YeSZ1520/ZstuWlan/issues)提出

# 常见路由器处理器架构

|           处理器            |架构|
|:------------------------:|:---:|
|    MT7921<br>MT7620A     |MIPS|
| MT7981B<br>MT7986A<br>MT7622B<br>IPQ8071A<br>IPQ6018<br>IPQ8072A<br>IPQ9570<br>|ARM|

# 鸣谢
[callmeliwen/RuijiePortalLoginTool](https://github.com/callmeliwen/RuijiePortalLoginTool)
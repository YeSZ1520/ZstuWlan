# 简介
本项目实现了 ZstuWlan 的多设备共享  

# 教程
## 路由器方案
### OpenWrt 路由器
1. 购买一个支持 OpenWrt 的路由器  
2. 为其刷上 OpenWrt 系统
3. 设置当中创建 Wifi  

相关教程有很多，不多做赘述，参考资料：[恩山论坛](https://www.right.com.cn/forum/forum.php)  

### 自动登录
上一步刷完 OpenWrt 并开启 Wifi 后，已可实现共享，连接路由器 Wifi，登录一次后，全部设备无需登录即可上网  

但是 12 点断网后路由器会失去验证，需要再次手动登录，比较麻烦，可以使用程序实现自动重连
#### 登录程序
ZSTU 使用的是锐捷 ePortal Web ，MenToHust 不可用，这里参考了 [RuijiePortalLoginTool](https://github.com/callmeliwen/RuijiePortalLoginTool) 项目

我使用了 Go 重写了 RuijiePortalLoginTool 的脚本，结果如 main.go，可用性未知，感兴趣可以自己测试一下 ( 之前的代码找不到了，这是重写的，我不在学校暂时测试不了 )  
使用方法：  
```shell
./login 学号 密码
```

> MTK7621 路由器编译参数：`GOOS=linux;GOARCH=mipsle;GOMIPS=softfloat;CGO_ENABLED=0

#### 定时登录
OpenWrt 应该自带了 Cron 定时执行程序，将登录程序加入 Cron 定时任务，早上7点执行一次即可

## 手机方案
实际上直接使用手机连接 ZstuWlan 再开热点就能实现共享，只是这样不方便和舍友共享网络

# 改进
1. 教程不是很详细
2. main.go 可用性未知
3. 可添加自动填写 Cron 定时任务的功能

# 问题和建议
有问题或建议可以在 [Issues ](https://github.com/YeSZ1520/ZstuWlan/issues)提出

# 鸣谢
[callmeliwen/RuijiePortalLoginTool](https://github.com/callmeliwen/RuijiePortalLoginTool)
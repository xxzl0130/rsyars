# Rsyars

Rsyars is a tool for GIRLS' FRONTLINE.

## 台服請在手機/模擬器安裝ca.crt
> 台服玩家，在模拟器上安装了证书设置了代理，模拟器上能够正常联网，控制台里也能看到开各种网页的请求。  
> 但在进游戏点开始游戏的时候会卡死在“请求版本号”上转圈圈，然后就连接逾时。如果在这一步关掉代理，可以正常进入游戏，且在游戏中开着代理好像也没有问题。但这样就读不到芯片数据了。  
> 经验证确认是因为安卓7以上app可以主动选择不认用户安装的证书，人工把装好的证书丢进系统证书目录里才能用  
> 从/data/misc/user/0/cacerts-added/把证书挪到/system/etc/security/cacerts/  
> By.[林凌a神林](https://bbs.nga.cn/read.php?pid=388760319&opt=128)

## 开源说明

除数据加解密模块以外均已开源

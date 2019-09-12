# Containers from scratch

大多數人都應該聽說過LXC——LinuX Containers，它是一個加強版的Chroot。簡單的說，LXC就是將不同的應用隔離開來，這有點類似於chroot，chroot是將應用隔離到一個虛擬的私有root下，而LXC在這之上更進了一步。LXC內部依賴Linux內核的3種隔離機制（isolation infrastructure）：

1. Chroot
2. Cgroups
3. Namespaces


Linux的3.12內核支持6種Namespace：

- UTS: hostname（本文介紹）
- IPC: 進程間通信 （之後的文章會講到）
- PID: "chroot"進程樹（之後的文章會講到）
- NS: 掛載點，首次登陸Linux（之後的文章會講到）
- NET: 網絡訪問，包括接口（之後的文章會講到）
- USER: 將本地的虛擬user-id映射到真實的user-id（之後的文章會講到）


cgroup

```
/sys/fs/cgroup/

$ cat /sys/fs/cgroup/memory/memory.limit_in_bytes 
9223372036854771712
```


## ref

- [Liz Rice - Containers from scratch](https://www.youtube.com/watch?v=oSlheqvaRso)
- [Linux Namespace系列（01）：Namespace概述](https://segmentfault.com/a/1190000006908272)
- [Linux命名空間學習教程（一） UTS](http://dockerone.com/article/76)

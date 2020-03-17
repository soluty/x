## 标准库time包的文档翻译
time包提供了测量和显示时间的功能

日历的计算都是基于格里历，就是我们使用的公历

Monotonic Clocks

以绝对时间为准，获取的时间为系统重启到现在的时间，更改系统时间对它没有影响。

wall lock
相对时间，从1970.1.1到目前的时间。更改系统时间会更改获取的值。它以系统时间为坐标。

通用的规则的wall clock用来告知时间，monotonic clock用来度量时间。

为了不分裂api，time.Now返回的时间包含了两个时间，下一个time-telling的操作使用wall clock，
下一个time-measuring的操作如比较和加减则使用monotonic clock

例如，下面的代码总是计算出大约20ms，即使wall clock在操作之前被改变。
```
start := time.Now()
... operation that takes 20 milliseconds ...
t := time.Now()
elapsed := t.Sub(start)
```

类似还有比如`time.Since(start), time.Until(deadline), time.Now().Before(deadline)`
等都对更改系统时间不敏感

下面小节给出精确的使用monotonic clocks的细节，但是使用这个包而言这个不是必须要理解的。

简单来说，两个时间如果都有monotonic clock读，那么操作t.After(u), t.Before(u), t.Equal(u), and t.Sub(u)就使用monotonic clock

在某些系统上，monotonic clock可能会停止，如果操作系统进入休眠。这时t.Sub(u)就可能不准了。

因为monotonic clock对其它进程毫无意义，所以序列化和反序列化不包含它们。

==操作比较时间，Location和monotonic clock reading。

为了debug用，t.String方法包含了monotonic clock reading

## 标准库time包在游戏使用上的缺陷

1. 时间不好自由调节，需要调节系统时间，这样不方便测试。
2. 时间精度过高，大部分游戏中精确到秒已经很足够了，并且游戏一般有大量的定时器，所以自己实现个时间轮很有必要。



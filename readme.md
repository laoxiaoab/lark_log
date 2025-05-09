1. 项目名称
百灵鸟日志
lark意为百灵鸟

2. 日志效率
每秒写入日志为 25000条, 1毫秒写入的数量为25条, 平均写入一条的时长为 0.04毫秒
0.04毫秒, 也就是40微秒
效率评估: 好像还行, 因为加了写锁, 1毫秒能写25条日志, 写入1条日志延迟0.04毫秒, 
在实际中应该不会导致程序的性能问题
【使用 os.File的Write写入】

3. 效率是否可以继续提升？
第一版用 os.File的 Write方法
经过查找, 使用bufIo.Writer可以显著提高写入的效率
因为bufIo内部使用了缓冲区
不过bufIo在写入后需要调用flush来确保所有缓冲的数据都被写入文件

4. 使用bufio写入文件的效率
改为bufIo的方式后, 写入效率测试如下:
total log count: 1000000, useTime:[2812]ms, write one log time: 0.002812[ms]
这样平均每秒写入的日志条数为: 100W/2.812s = 355618条, 也就是35W条
每条日志的延迟为 0.002812毫秒, 也就是2.8微秒, 这个效率是很不错的了
因为同样是加了互斥锁的

5. bufIo方式同时开启3个goroutine, 每个goroutine写入100w条日志时的效率
total log count: 1000000, useTime:[12642]ms, write one log time: 0.012642[ms]
goroutine数量3, 使用sync.WaitGroup写入, 每个goroutine写入100w条, 总耗时约12.6秒

6. bufIo方式单个goroutine写入300W条数据效率
单个goroutine写入300w条耗时:
total log count: 3000000, useTime:[8789]ms, write one log time: 0.002929666666666667[ms]

7. 单个goroutine写入300W条和3个goroutine每个写入100W条的效率对比
单个goroutine: [8789]ms
3个goroutine: [12642]ms
因为多个goroutine之间有mutex加锁解锁开销, 因此总时长会比单个goroutine高很多


package xtime

import (
	"math"
	"time"
)

// 两个实现，一个时间轮，一个标准库wrapper, 在wrap中可以随意修改时间测试。

func _()  {
	time.Now()
}

type Job = func()

const gDEFAULT_TIMES          = math.MaxInt32
const STATUS_READY            = 0
const STATUS_RUNNING          = 1
const STATUS_STOPPED          = 2
const STATUS_CLOSED           = -1

type wheel struct {
	timer      *Timer        // 所属定时器
	level      int           // 所属分层索引号
	slots      []*glist.List // 所有的循环任务项, 按照Slot Number进行分组
	number     int64         // Slot Number=len(slots)
	ticks      *gtype.Int64  // 当前时间轮已转动的刻度数量
	totalMs    int64         // 整个时间轮的时间长度(毫秒)=number*interval
	createMs   int64         // 创建时间(毫秒)
	intervalMs int64         // 时间间隔(slot时间长度, 毫秒)
}

type Timer struct {
	status     int // 定时器状态
	wheels     []*wheel   // 分层时间轮对象
	length     int        // 分层层数
	number     int        // 每一层Slot Number
	intervalMs int64      // 最小时间刻度(毫秒)
}

func New(slot int, interval time.Duration, level ...int) *Timer {
	length := 6
	if len(level) > 0 {
		length = level[0]
	}
	t := &Timer{
		status:     STATUS_RUNNING,
		wheels:     make([]*wheel, length),
		length:     length,
		number:     slot,
		intervalMs: interval.Nanoseconds() / 1e6,
	}
	for i := 0; i < length; i++ {
		if i > 0 {
			n := time.Duration(t.wheels[i-1].totalMs) * time.Millisecond
			w := t.newWheel(i, slot, n)
			t.wheels[i] = w
			t.wheels[i-1].addEntry(n, w.proceed, false, gDEFAULT_TIMES, STATUS_READY)
		} else {
			t.wheels[i] = t.newWheel(i, slot, interval)
		}
	}
	t.wheels[0].start()
	return t
}

//
func (t *Timer) newWheel(level int, slot int, interval time.Duration) *wheel {
	w := &wheel{
		timer:      t,
		level:      level,
		slots:      make([]*glist.List, slot),
		number:     int64(slot),
		ticks:      0,  // 已经转动的刻度
		totalMs:    int64(slot) * interval.Nanoseconds() / 1e6,
		createMs:   time.Now().UnixNano() / 1e6,
		intervalMs: interval.Nanoseconds() / 1e6,
	}
	for i := int64(0); i < w.number; i++ {
		w.slots[i] = glist.New(true)
	}
	return w
}

func (t *Timer) Start() {
	t.status.Set(STATUS_RUNNING)
}

type Entry struct {
	wheel         *wheel      // 所属时间轮
	job           func()     // 注册循环任务方法
	singleton     *gtype.Bool // 任务是否单例运行
	status        *gtype.Int  // 任务状态(0: ready;  1: running; 2: stopped; -1: closed), 层级entry共享状态
	times         *gtype.Int  // 还需运行次数
	create        int64       // 注册时的时间轮ticks
	interval      int64       // 设置的运行间隔(时间轮刻度数量)
	createMs      int64       // 创建时间(毫秒)
	intervalMs    int64       // 间隔时间(毫秒)
	rawIntervalMs int64       // 原始间隔
}


func (w *wheel) addEntry(interval time.Duration, job func(), singleton bool, times int, status int) *Entry {
	if times <= 0 {
		times = gDEFAULT_TIMES
	}
	ms := interval.Nanoseconds() / 1e6
	num := ms / w.intervalMs
	if num == 0 {
		// 如果安装的任务间隔小于时间轮刻度，
		// 那么将会在下一刻度被执行
		num = 1
	}
	nowMs := time.Now().UnixNano() / 1e6
	ticks := w.ticks.Val()
	entry := &Entry{
		wheel:         w,
		job:           job,
		times:         gtype.NewInt(times),
		status:        gtype.NewInt(status),
		create:        ticks,
		interval:      num,
		singleton:     gtype.NewBool(singleton),
		createMs:      nowMs,
		intervalMs:    ms,
		rawIntervalMs: ms,
	}
	// 安装任务
	w.slots[(ticks+num)%w.number].PushBack(entry)
	return entry
}


func (w *wheel) start() {
	go func() {
		ticker := time.NewTicker(time.Duration(w.intervalMs) * time.Millisecond)
		for {
			select {
			case <-ticker.C:
				switch w.timer.status.Val() {
				case STATUS_RUNNING:
					w.proceed()

				case STATUS_STOPPED:
				case STATUS_CLOSED:
					ticker.Stop()
					return
				}

			}
		}
	}()
}

func (w *wheel) proceed() {
	n := w.ticks.Add(1)
	l := w.slots[int(n%w.number)]
	length := l.Len()
	if length > 0 {
		go func(l *glist.List, nowTicks int64) {
			entry := (*Entry)(nil)
			nowMs := time.Now().UnixNano() / 1e6
			for i := length; i > 0; i-- {
				if v := l.PopFront(); v == nil {
					break
				} else {
					entry = v.(*Entry)
				}
				// 是否满足运行条件
				runnable, addable := entry.check(nowTicks, nowMs)
				if runnable {
					// 异步执行运行
					go func(entry *Entry) {
						defer func() {
							if err := recover(); err != nil {
								if err != gPANIC_EXIT {
									panic(err)
								} else {
									entry.Close()
								}
							}
							if entry.Status() == STATUS_RUNNING {
								entry.SetStatus(STATUS_READY)
							}
						}()
						entry.job()
					}(entry)
				}
				// 是否继续添运行, 滚动任务
				if addable {
					entry.wheel.timer.doAddEntryByParent(entry.rawIntervalMs, entry)
				}
			}
		}(l, n)
	}
}

// 添加定时任务
func (t *Timer) AddEntry(interval time.Duration, job Job, singleton bool, times int, status int) *Entry {
	return t.doAddEntry(interval, job, singleton, times, status)
}
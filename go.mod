module github.com/soluty/x

go 1.13

require (
	github.com/AsynkronIT/protoactor-go v0.0.0-20191119210846-07df21a705b8
	github.com/alexeyco/simpletable v0.0.0-20191023080658-fe3ac9971811
	github.com/cloudflare/tableflip v1.0.0 // indirect
	github.com/codegangsta/inject v0.0.0-20150114235600-33e0aa1cb7c0
	github.com/coreos/etcd v3.3.17+incompatible
	github.com/emirpasic/gods v1.12.0
	github.com/faiface/glhf v0.0.0-20181018222622-82a6317ac380 // indirect
	github.com/faiface/mainthread v0.0.0-20171120011319-8b78f0a41ae3 // indirect
	github.com/faiface/pixel v0.8.0
	github.com/fsnotify/fsnotify v1.4.7
	github.com/go-gl/gl v0.0.0-20190320180904-bf2b1f2f34d7 // indirect
	github.com/go-gl/mathgl v0.0.0-20190713194549-592312d8590a // indirect
	github.com/gogf/gf v1.11.4 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/golang/protobuf v1.3.2
	github.com/hajimehoshi/ebiten v1.10.3
	github.com/magicsea/behavior3go v0.0.0-20190816070510-c4fcf60da748
	github.com/rivo/tview v0.0.0-20200219210816-cd38d7432498 // indirect
	github.com/rs/xid v1.2.1
	github.com/sirupsen/logrus v1.4.2
	github.com/smallnest/libkv-etcdv3-store v0.0.0-20191101045330-f92940446965
	github.com/smallnest/rpcx v0.0.0-20191228024106-2e3195bbbddb
	github.com/soluty/x/cc v1.0.0
	github.com/soluty/x/conv v0.0.0-20200110100703-54d0bab3b953
	github.com/soluty/x/ecs v1.0.0
	github.com/soluty/x/xcrypto v0.0.0-20191111054007-09768a6afb59
	github.com/soluty/x/xnet v1.0.0
	github.com/soluty/x/xrand v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.1
	github.com/tidwall/gjson v1.5.0
	github.com/tidwall/sjson v1.0.4
	github.com/twitchtv/twirp v5.10.1+incompatible
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8
	golang.org/x/net v0.0.0-20191116160921-f9c825593386
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/text v0.3.2
	google.golang.org/grpc v1.25.1
	gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637
	logur.dev/adapter/logrus v0.2.0
	logur.dev/logur v0.15.1
	upper.io/db.v3 v3.6.3+incompatible // indirect
)

replace github.com/soluty/x/xrand v1.0.0 => ./xrand

replace github.com/soluty/x/ecs v1.0.0 => ./ecs

replace github.com/soluty/x/cc v1.0.0 => ./cc

replace github.com/soluty/x/emitter v1.0.0 => ./emitter

replace github.com/soluty/x/xnet v1.0.0 => ./xnet

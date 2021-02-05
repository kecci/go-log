# go-log

Comparison between logger in go

Log a message and 10 fields:
```sh
Library	Time	Bytes Allocated	Objects Allocated
zerolog	767 ns/op	552 B/op	6 allocs/op
⚡ zap	848 ns/op	704 B/op	2 allocs/op
⚡ zap (sugared)	1363 ns/op	1610 B/op	20 allocs/op
go-kit	3614 ns/op	2895 B/op	66 allocs/op
lion	5392 ns/op	5807 B/op	63 allocs/op
logrus	5661 ns/op	6092 B/op	78 allocs/op
apex/log	15332 ns/op	3832 B/op	65 allocs/op
log15	20657 ns/op	5632 B/op	93 allocs/op
```

Log a message with a logger that already has 10 fields of context:
```sh
Library	Time	Bytes Allocated	Objects Allocated
zerolog	52 ns/op	0 B/op	0 allocs/op
⚡ zap	283 ns/op	0 B/op	0 allocs/op
⚡ zap (sugared)	337 ns/op	80 B/op	2 allocs/op
lion	2702 ns/op	4074 B/op	38 allocs/op
go-kit	3378 ns/op	3046 B/op	52 allocs/op
logrus	4309 ns/op	4564 B/op	63 allocs/op
apex/log	13456 ns/op	2898 B/op	51 allocs/op
log15	14179 ns/op	2642 B/op	44 allocs/op
```

Log a static string, without any context or printf-style templating:
```sh
Library	Time	Bytes Allocated	Objects Allocated
zerolog	50 ns/op	0 B/op	0 allocs/op
⚡ zap	236 ns/op	0 B/op	0 allocs/op
standard library	453 ns/op	80 B/op	2 allocs/op
⚡ zap (sugared)	337 ns/op	80 B/op	2 allocs/op
go-kit	508 ns/op	656 B/op	13 allocs/op
lion	771 ns/op	1224 B/op	10 allocs/op
logrus	1244 ns/op	1505 B/op	27 allocs/op
apex/log	2751 ns/op	584 B/op	11 allocs/op
log15	5181 ns/op	1592 B/op	26 allocs/op
```
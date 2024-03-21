## Running the Program

Build the Program:

```
go build -o main
```

Run the program:
```
./main
```
To print statistics about the resource usage of the command, including the elapsed real time, user CPU time, and system CPU time:
```
time ./main
```
Output would look like:
```
"ego.txt": 160578

real    0m2.102s
user    0m0.005s
sys     0m0.027s
```
> `real`: Elapsed real (wall-clock) time. This is the actual time taken from the start to the end of the command's execution.
>
> `user`: CPU time spent in user-mode code (outside the kernel). This includes the time spent executing your program's code.
>
> `sys`: CPU time spent by the kernel on behalf of the command. This includes time spent on system calls and other kernel operations.


If you want to compare it with the inbuilf `wc` command:
```
time wc -w fileName.txt
```

> Note: The `time` command is useful for profiling and understanding the performance characteristics of our programs.


## Benchmarking

```
go test -bench=BenchmarkWC
```

### Baseline:

<u>Benchmarking Result:</u>
```
"ego.txt": 160578
goos: windows
goarch: amd64
pkg: wc
cpu: AMD Ryzen 5 5500U with Radeon Graphics
BenchmarkWC-12                 1        1903381200 ns/op
PASS
ok      wc      2.476s
```
<u>Resource Stats:</u>
```
"ego.txt": 160578

real    0m2.223s
user    0m0.004s
sys     0m0.030s
```
<u>Resource Stats for standard `wc` command:</u>
> Note: There's a difference between word count here, that is because, there are some forms of words that I'm not considering.
```
161916 ego.txt

real    0m0.101s
user    0m0.008s
sys     0m0.059s
```

<u>Profiling:</u>

CPU Profile:
```
go test -cpuprofile cpu.prof -bench .
```

Visualize the profile in web:
```
go tool pprof -http=:8080 cpu.prof
```
![Profile Image](profile.png)

Want a PNG image of the profile?
```
go tool pprof -output="profile.png" wc.test.exe cpu.prof
```
This will open `pprof` interactive shell. Type in `png` there and press `Enter`
> Note: Enter your own binary file name in place of `wc.test.exe`

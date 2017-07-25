execve("./cmd", ["./cmd"], [/* 60 vars */]) = 0
arch_prctl(ARCH_SET_FS, 0x584f70)       = 0
sched_getaffinity(0, 128, {f, 0, 0, 0}) = 32
mmap(0xc000000000, 65536, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0xc000000000
munmap(0xc000000000, 65536)             = 0
mmap(NULL, 262144, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f814232e000
mmap(0xc208000000, 1048576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0xc208000000
mmap(0xc207ff0000, 65536, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0xc207ff0000
mmap(0xc000000000, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0xc000000000
mmap(NULL, 65536, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f814231e000
mmap(NULL, 1439992, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f81421be000
mmap(NULL, 131072, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f814219e000
sigaltstack({ss_sp=0xc208006000, ss_flags=0, ss_size=32768}, NULL) = 0
rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
rt_sigaction(SIGHUP, NULL, {SIG_DFL, [], 0}, 8) = 0
rt_sigaction(SIGHUP, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGINT, NULL, {SIG_DFL, [], 0}, 8) = 0
rt_sigaction(SIGINT, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGQUIT, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGILL, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGTRAP, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGABRT, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGBUS, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGFPE, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGUSR1, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGSEGV, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGUSR2, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGPIPE, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGALRM, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGTERM, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGSTKFLT, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGCHLD, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGURG, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGXCPU, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGXFSZ, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGVTALRM, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGPROF, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGWINCH, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGIO, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGPWR, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGSYS, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_2, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_3, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_4, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_5, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_6, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_7, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_8, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_9, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_10, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_11, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_12, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_13, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_14, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_15, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_16, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_17, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_18, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_19, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_20, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_21, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_22, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_23, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_24, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_25, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_26, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_27, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_28, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_29, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_30, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_31, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigaction(SIGRT_32, {0x429100, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429170}, NULL, 8) = 0
rt_sigprocmask(SIG_SETMASK, ~[], [], 8) = 0
clone(Process 20017 attached (waiting for parent)
Process 20017 resumed (parent 20016 ready)
child_stack=0x7f81421bdfa8, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD) = 20017
[pid 20017] gettid( <unfinished ...>
[pid 20016] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20017] <... gettid resumed> )      = 20017
[pid 20016] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20017] arch_prctl(ARCH_SET_FS, 0xc208016070) = 0
[pid 20016] mmap(NULL, 1048576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0 <unfinished ...>
[pid 20017] sigaltstack({ss_sp=0xc20801a000, ss_flags=0, ss_size=32768} <unfinished ...>
[pid 20016] <... mmap resumed> )        = 0x7f814209e000
[pid 20017] <... sigaltstack resumed> , NULL) = 0
[pid 20017] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20016] stat("/home/qboxtest/.rvm/gems/ruby-2.1.3/bin/go",  <unfinished ...>
[pid 20017] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20016] <... stat resumed> 0xc208040000) = -1 ENOENT (No such file or directory)
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] stat("/home/qboxtest/.rvm/gems/ruby-2.1.3@global/bin/go",  <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... stat resumed> 0xc208040090) = -1 ENOENT (No such file or directory)
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] stat("/home/qboxtest/.rvm/rubies/ruby-2.1.3/bin/go", 0xc208040120) = -1 ENOENT (No such file or directory)
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] stat("/home/qboxtest/bin/go",  <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] <... stat resumed> 0xc2080401b0) = -1 ENOENT (No such file or directory)
[pid 20016] stat("/usr/lib/lightdm/lightdm/go",  <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... stat resumed> 0xc208040240) = -1 ENOENT (No such file or directory)
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] stat("/usr/local/sbin/go",  <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... stat resumed> 0xc2080402d0) = -1 ENOENT (No such file or directory)
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] stat("/usr/local/bin/go", 0xc208040360) = -1 ENOENT (No such file or directory)
[pid 20016] stat("/usr/sbin/go",  <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... stat resumed> 0xc2080403f0) = -1 ENOENT (No such file or directory)
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] stat("/usr/bin/go", 0xc208040480) = -1 ENOENT (No such file or directory)
[pid 20016] stat("/sbin/go",  <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... stat resumed> 0xc208040510) = -1 ENOENT (No such file or directory)
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] stat("/bin/go", 0xc2080405a0) = -1 ENOENT (No such file or directory)
[pid 20016] stat("/usr/games/go",  <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... stat resumed> 0xc208040630) = -1 ENOENT (No such file or directory)
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] stat("/home/qboxtest/.rvm/bin/go", 0xc2080406c0) = -1 ENOENT (No such file or directory)
[pid 20016] stat("/usr/local/go/bin/go",  <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... stat resumed> {st_mode=S_IFREG|0755, st_size=9349952, ...}) = 0
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] open("/dev/null", O_RDONLY|O_CLOEXEC <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... open resumed> )        = 3
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] open("/dev/null", O_WRONLY|O_CLOEXEC) = 4
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] open("/dev/null", O_WRONLY|O_CLOEXEC <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] <... open resumed> )        = 5
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] pipe2( <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... pipe2 resumed> [6, 7], O_CLOEXEC) = 0
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] clone( <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
Process 20018 attached (waiting for parent)
[pid 20017] select(0, NULL, NULL, NULL, {0, 20}Process 20018 resumed (parent 20016 ready)
 <unfinished ...>
[pid 20016] <... clone resumed> child_stack=0, flags=|SIGCHLD) = 20018
[pid 20018] dup2(3, 0 <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... dup2 resumed> )        = 0
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] dup2(4, 1 <unfinished ...>
[pid 20016] close(7 <unfinished ...>
[pid 20018] <... dup2 resumed> )        = 1
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20018] dup2(5, 2 <unfinished ...>
[pid 20016] <... close resumed> )       = 0
[pid 20018] <... dup2 resumed> )        = 2
[pid 20018] execve("/usr/local/go/bin/go", ["go", "run", "childProcess.go"], [/* 60 vars */] <unfinished ...>
[pid 20017] mmap(NULL, 131072, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0 <unfinished ...>
[pid 20016] sched_yield( <unfinished ...>
[pid 20017] <... mmap resumed> )        = 0x7f814207e000
[pid 20016] <... sched_yield resumed> ) = 0
[pid 20017] rt_sigprocmask(SIG_SETMASK, ~[],  <unfinished ...>
[pid 20016] futex(0x5850e0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... rt_sigprocmask resumed> [], 8) = 0
[pid 20016] <... futex resumed> )       = 0
[pid 20017] clone( <unfinished ...>
[pid 20018] <... execve resumed> )      = 0
Process 20019 attached (waiting for parent)
Process 20019 resumed (parent 20017 ready)
[pid 20017] <... clone resumed> child_stack=0x7f814207ffa8, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD) = 20019
[pid 20019] gettid( <unfinished ...>
[pid 20017] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20019] <... gettid resumed> )      = 20019
[pid 20017] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20019] arch_prctl(ARCH_SET_FS, 0xc208017270 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... arch_prctl resumed> )  = 0
[pid 20016] futex(0x585870, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20019] sigaltstack({ss_sp=0xc208052000, ss_flags=0, ss_size=32768} <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] <... sigaltstack resumed> , NULL) = 0
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
[pid 20019] futex(0x585870, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] <... futex resumed> )       = 1
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20016] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20018] brk(0 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... brk resumed> )         = 0x194f000
[pid 20016] read(6, "", 8)              = 0
[pid 20018] access("/etc/ld.so.nohwcap", F_OK <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... access resumed> )      = -1 ENOENT (No such file or directory)
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] mmap(NULL, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0 <unfinished ...>
[pid 20016] close(6 <unfinished ...>
[pid 20018] <... mmap resumed> )        = 0x7f33bba7c000
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... close resumed> )       = 0
[pid 20017] rt_sigprocmask(SIG_SETMASK, ~[],  <unfinished ...>
[pid 20018] access("/etc/ld.so.preload", R_OK <unfinished ...>
[pid 20017] <... rt_sigprocmask resumed> [], 8) = 0
[pid 20018] <... access resumed> )      = -1 ENOENT (No such file or directory)
[pid 20017] clone( <unfinished ...>
[pid 20018] open("/usr/local/lib/tls/x86_64/libpthread.so.0", O_RDONLY|O_CLOEXECProcess 20020 attached (waiting for parent)
) = -1 ENOENT (No such file or directory)
Process 20020 resumed (parent 20017 ready)
[pid 20017] <... clone resumed> child_stack=0x7f814209dfa8, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD) = 20020
[pid 20020] gettid( <unfinished ...>
[pid 20018] stat("/usr/local/lib/tls/x86_64",  <unfinished ...>
[pid 20020] <... gettid resumed> )      = 20020
[pid 20018] <... stat resumed> 0x7ffff19b66c0) = -1 ENOENT (No such file or directory)
[pid 20020] arch_prctl(ARCH_SET_FS, 0xc2080176f0 <unfinished ...>
[pid 20018] open("/usr/local/lib/tls/libpthread.so.0", O_RDONLY|O_CLOEXEC <unfinished ...>
[pid 20020] <... arch_prctl resumed> )  = 0
[pid 20018] <... open resumed> )        = -1 ENOENT (No such file or directory)
[pid 20020] sigaltstack({ss_sp=0xc20805a000, ss_flags=0, ss_size=32768} <unfinished ...>
[pid 20018] stat("/usr/local/lib/tls",  <unfinished ...>
[pid 20020] <... sigaltstack resumed> , NULL) = 0
[pid 20018] <... stat resumed> 0x7ffff19b66c0) = -1 ENOENT (No such file or directory)
[pid 20020] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20018] open("/usr/local/lib/x86_64/libpthread.so.0", O_RDONLY|O_CLOEXEC <unfinished ...>
[pid 20020] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20018] <... open resumed> )        = -1 ENOENT (No such file or directory)
[pid 20017] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20020] close(3 <unfinished ...>
[pid 20018] stat("/usr/local/lib/x86_64",  <unfinished ...>
[pid 20020] <... close resumed> )       = 0
[pid 20018] <... stat resumed> 0x7ffff19b66c0) = -1 ENOENT (No such file or directory)
[pid 20020] close(4 <unfinished ...>
[pid 20018] open("/usr/local/lib/libpthread.so.0", O_RDONLY|O_CLOEXEC <unfinished ...>
[pid 20020] <... close resumed> )       = 0
[pid 20018] <... open resumed> )        = -1 ENOENT (No such file or directory)
[pid 20020] close(5 <unfinished ...>
[pid 20018] stat("/usr/local/lib",  <unfinished ...>
[pid 20020] <... close resumed> )       = 0
[pid 20018] <... stat resumed> {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20017] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20018] open("/etc/ld.so.cache", O_RDONLY|O_CLOEXEC <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... open resumed> )        = 3
[pid 20016] futex(0x585870, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] fstat(3,  <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... fstat resumed> {st_mode=S_IFREG|0644, st_size=124566, ...}) = 0
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] mmap(NULL, 124566, PROT_READ, MAP_PRIVATE, 3, 0) = 0x7f33bba5d000
[pid 20018] close(3)                    = 0
[pid 20018] access("/etc/ld.so.nohwcap", F_OK <unfinished ...>
[pid 20020] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... access resumed> )      = -1 ENOENT (No such file or directory)
[pid 20020] <... futex resumed> )       = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20018] open("/lib/x86_64-linux-gnu/libpthread.so.0", O_RDONLY|O_CLOEXEC <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... open resumed> )        = 3
[pid 20018] read(3, "\177ELF\2\1\1\0\0\0\0\0\0\0\0\0\3\0>\0\1\0\0\0\200l\0\0\0\0\0\0"..., 832) = 832
[pid 20020] mmap(NULL, 131072, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0 <unfinished ...>
[pid 20018] fstat(3,  <unfinished ...>
[pid 20020] <... mmap resumed> )        = 0x7f814205e000
[pid 20018] <... fstat resumed> {st_mode=S_IFREG|0755, st_size=135366, ...}) = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20018] mmap(NULL, 2212904, PROT_READ|PROT_EXEC, MAP_PRIVATE|MAP_DENYWRITE, 3, 0 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... mmap resumed> )        = 0x7f33bb63f000
[pid 20018] mprotect(0x7f33bb657000, 2093056, PROT_NONE <unfinished ...>
[pid 20020] futex(0x585870, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... mprotect resumed> )    = 0
[pid 20020] <... futex resumed> )       = 1
[pid 20018] mmap(0x7f33bb856000, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x17000 <unfinished ...>
[pid 20020] futex(0x584ef8, FUTEX_WAIT, 0, {0, 999391764} <unfinished ...>
[pid 20018] <... mmap resumed> )        = 0x7f33bb856000
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20018] mmap(0x7f33bb858000, 13352, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... mmap resumed> )        = 0x7f33bb858000
[pid 20016] <... futex resumed> )       = 0
[pid 20018] close(3 <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... close resumed> )       = 0
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] open("/usr/local/lib/libc.so.6", O_RDONLY|O_CLOEXEC <unfinished ...>
[pid 20016] futex(0x585870, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] <... open resumed> )        = -1 ENOENT (No such file or directory)
[pid 20018] access("/etc/ld.so.nohwcap", F_OK) = -1 ENOENT (No such file or directory)
[pid 20018] open("/lib/x86_64-linux-gnu/libc.so.6", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "\177ELF\2\1\1\0\0\0\0\0\0\0\0\0\3\0>\0\1\0\0\0\200\30\2\0\0\0\0\0"..., 832) = 832
[pid 20018] fstat(3, {st_mode=S_IFREG|0755, st_size=1815224, ...}) = 0
[pid 20018] mmap(NULL, 3929304, PROT_READ|PROT_EXEC, MAP_PRIVATE|MAP_DENYWRITE, 3, 0) = 0x7f33bb27f000
[pid 20018] mprotect(0x7f33bb434000, 2097152, PROT_NONE) = 0
[pid 20018] mmap(0x7f33bb634000, 24576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x1b5000) = 0x7f33bb634000
[pid 20018] mmap(0x7f33bb63a000, 17624, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) = 0x7f33bb63a000
[pid 20018] close(3)                    = 0
[pid 20018] mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f33bba5c000
[pid 20018] mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f33bba5b000
[pid 20018] mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f33bba5a000
[pid 20018] arch_prctl(ARCH_SET_FS, 0x7f33bba5b700) = 0
[pid 20018] mprotect(0x7f33bb634000, 16384, PROT_READ) = 0
[pid 20018] mprotect(0x7f33bb856000, 4096, PROT_READ) = 0
[pid 20018] mprotect(0x7f33bba7e000, 4096, PROT_READ) = 0
[pid 20018] munmap(0x7f33bba5d000, 124566) = 0
[pid 20018] set_tid_address(0x7f33bba5b9d0) = 20018
[pid 20018] set_robust_list(0x7f33bba5b9e0, 0x18) = 0
[pid 20018] futex(0x7ffff19b6fbc, FUTEX_WAIT_BITSET_PRIVATE|FUTEX_CLOCK_REALTIME, 1, NULL, 7f33bba5b700) = -1 EAGAIN (Resource temporarily unavailable)
[pid 20018] rt_sigaction(SIGRTMIN, {0x7f33bb645750, [], SA_RESTORER|SA_SIGINFO, 0x7f33bb64ecb0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_1, {0x7f33bb6457e0, [], SA_RESTORER|SA_RESTART|SA_SIGINFO, 0x7f33bb64ecb0}, NULL, 8) = 0
[pid 20018] rt_sigprocmask(SIG_UNBLOCK, [RTMIN RT_1], NULL, 8) = 0
[pid 20018] getrlimit(RLIMIT_STACK, {rlim_cur=8192*1024, rlim_max=RLIM_INFINITY}) = 0
[pid 20018] sched_getaffinity(0, 128, {f, 0, 0, 0}) = 32
[pid 20018] mmap(0xc000000000, 65536, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0xc000000000
[pid 20018] munmap(0xc000000000, 65536) = 0
[pid 20018] mmap(NULL, 262144, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f33bba1a000
[pid 20018] mmap(0xc208000000, 1048576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0xc208000000
[pid 20018] mmap(0xc207ff0000, 65536, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0xc207ff0000
[pid 20018] mmap(0xc000000000, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0xc000000000
[pid 20018] mmap(NULL, 65536, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f33bba6c000
[pid 20018] mmap(NULL, 1439992, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f33bb8ba000
[pid 20018] mmap(NULL, 131072, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f33bb89a000
[pid 20018] sigaltstack({ss_sp=0xc208006000, ss_flags=0, ss_size=32768}, NULL) = 0
[pid 20018] rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
[pid 20018] rt_sigaction(SIGHUP, NULL, {SIG_DFL, [], 0}, 8) = 0
[pid 20018] rt_sigaction(SIGHUP, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGINT, NULL, {SIG_DFL, [], 0}, 8) = 0
[pid 20018] rt_sigaction(SIGINT, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGQUIT, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGILL, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGTRAP, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGABRT, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGBUS, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGFPE, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGUSR1, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGSEGV, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGUSR2, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGPIPE, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGALRM, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGTERM, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGSTKFLT, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGCHLD, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGURG, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGXCPU, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGXFSZ, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGVTALRM, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGPROF, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGWINCH, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGIO, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGPWR, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGSYS, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_2, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_3, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_4, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_5, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_6, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_7, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_8, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_9, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_10, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_11, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_12, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_13, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_14, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_15, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_16, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_17, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_18, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_19, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_20, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_21, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_22, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_23, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_24, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_25, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_26, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_27, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_28, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_29, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_30, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_31, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] rt_sigaction(SIGRT_32, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20018] brk(0)                      = 0x194f000
[pid 20018] brk(0x1970000)              = 0x1970000
[pid 20018] rt_sigprocmask(SIG_SETMASK, ~[RTMIN RT_1], [], 8) = 0
[pid 20018] mmap(NULL, 8392704, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_STACK, -1, 0) = 0x7f33baa7e000
[pid 20018] mprotect(0x7f33baa7e000, 4096, PROT_NONE) = 0
[pid 20018] clone(Process 20021 attached
child_stack=0x7f33bb27dff0, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLS|CLONE_PARENT_SETTID|CLONE_CHILD_CLEARTID, parent_tidptr=0x7f33bb27e9d0, tls=0x7f33bb27e700, child_tidptr=0x7f33bb27e9d0) = 20021
[pid 20018] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20021] set_robust_list(0x7f33bb27e9e0, 0x18 <unfinished ...>
[pid 20018] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20021] <... set_robust_list resumed> ) = 0
[pid 20021] sigaltstack({ss_sp=0xc20801c000, ss_flags=0, ss_size=32768}, NULL) = 0
[pid 20018] mmap(NULL, 1048576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0 <unfinished ...>
[pid 20021] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20018] <... mmap resumed> )        = 0x7f33ba97e000
[pid 20021] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] futex(0xa80cb8, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xa80cb8, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] futex(0xa80cb8, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xa80cb8, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20018] open("/proc/sys/net/core/somaxconn", O_RDONLY|O_CLOEXEC <unfinished ...>
[pid 20021] open("/sys/devices/system/cpu/online", O_RDONLY|O_CLOEXEC <unfinished ...>
[pid 20018] <... open resumed> )        = 3
[pid 20021] <... open resumed> )        = 4
[pid 20018] futex(0xa81b10, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20021] read(4, "0-3\n", 8192)      = 4
[pid 20021] close(4)                    = 0
[pid 20021] mmap(NULL, 134217728, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_NORESERVE, -1, 0) = 0x7f33b297e000
[pid 20021] munmap(0x7f33b297e000, 23601152) = 0
[pid 20021] munmap(0x7f33b8000000, 43507712) = 0
[pid 20021] mprotect(0x7f33b4000000, 135168, PROT_READ|PROT_WRITE) = 0
[pid 20021] rt_sigprocmask(SIG_SETMASK, ~[RTMIN RT_1], [], 8) = 0
[pid 20021] mmap(NULL, 8392704, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_STACK, -1, 0) = 0x7f33ba17d000
[pid 20021] mprotect(0x7f33ba17d000, 4096, PROT_NONE) = 0
[pid 20021] clone(Process 20022 attached (waiting for parent)
Process 20022 resumed (parent 20021 ready)
child_stack=0x7f33ba97cff0, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLS|CLONE_PARENT_SETTID|CLONE_CHILD_CLEARTID, parent_tidptr=0x7f33ba97d9d0, tls=0x7f33ba97d700, child_tidptr=0x7f33ba97d9d0) = 20022
[pid 20022] set_robust_list(0x7f33ba97d9e0, 0x18 <unfinished ...>
[pid 20021] rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20022] <... set_robust_list resumed> ) = 0
[pid 20022] sigaltstack({ss_sp=0xc208064000, ss_flags=0, ss_size=32768} <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20022] <... sigaltstack resumed> , NULL) = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20022] rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20022] mmap(NULL, 134217728, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_NORESERVE, -1, 0) = 0x7f33ac000000
[pid 20022] munmap(0x7f33b0000000, 67108864 <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20022] <... munmap resumed> )      = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20022] mprotect(0x7f33ac000000, 135168, PROT_READ|PROT_WRITE) = 0
[pid 20022] rt_sigprocmask(SIG_SETMASK, ~[RTMIN RT_1],  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20022] <... rt_sigprocmask resumed> [], 8) = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20022] mmap(NULL, 8392704, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_STACK, -1, 0) = 0x7f33b997c000
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20022] mprotect(0x7f33b997c000, 4096, PROT_NONE <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20022] <... mprotect resumed> )    = 0
[pid 20022] clone( <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
Process 20023 attached
[pid 20022] <... clone resumed> child_stack=0x7f33ba17bff0, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLS|CLONE_PARENT_SETTID|CLONE_CHILD_CLEARTID, parent_tidptr=0x7f33ba17c9d0, tls=0x7f33ba17c700, child_tidptr=0x7f33ba17c9d0) = 20023
[pid 20022] rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
[pid 20023] set_robust_list(0x7f33ba17c9e0, 0x18 <unfinished ...>
[pid 20022] futex(0x7f33bb8b9f60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20023] <... set_robust_list resumed> ) = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20023] sigaltstack({ss_sp=0xc20806c000, ss_flags=0, ss_size=32768}, NULL) = 0
[pid 20023] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20023] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20023] mmap(0x7f33b0000000, 67108864, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_NORESERVE, -1, 0 <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20023] <... mmap resumed> )        = 0x7f33b0000000
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20023] mprotect(0x7f33b0000000, 135168, PROT_READ|PROT_WRITE) = 0
[pid 20023] rt_sigprocmask(SIG_SETMASK, ~[RTMIN RT_1],  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20023] <... rt_sigprocmask resumed> [], 8) = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20023] mmap(NULL, 8392704, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_STACK, -1, 0) = 0x7f33b917b000
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20023] mprotect(0x7f33b917b000, 4096, PROT_NONE <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20023] <... mprotect resumed> )    = 0
[pid 20023] clone( <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
Process 20024 attached (waiting for parent)
Process 20024 resumed (parent 20023 ready)
[pid 20023] <... clone resumed> child_stack=0x7f33b997aff0, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLS|CLONE_PARENT_SETTID|CLONE_CHILD_CLEARTID, parent_tidptr=0x7f33b997b9d0, tls=0x7f33b997b700, child_tidptr=0x7f33b997b9d0) = 20024
[pid 20024] set_robust_list(0x7f33b997b9e0, 0x18 <unfinished ...>
[pid 20023] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20024] <... set_robust_list resumed> ) = 0
[pid 20023] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20024] sigaltstack({ss_sp=0xc208074000, ss_flags=0, ss_size=32768} <unfinished ...>
[pid 20023] futex(0xa808c0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] <... sigaltstack resumed> , NULL) = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] <... futex resumed> )       = 1
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] <... futex resumed> )       = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] read(3,  <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... read resumed> "128\n", 4096) = 4
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] read(3, "", 4092)           = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] close(3)                    = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] socket(PF_INET, SOCK_STREAM, IPPROTO_TCP) = 3
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] close(3 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... close resumed> )       = 0
[pid 20018] socket(PF_INET6, SOCK_STREAM, IPPROTO_TCP <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... socket resumed> )      = 3
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] setsockopt(3, SOL_IPV6, IPV6_V6ONLY, [1], 4) = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] bind(3, {sa_family=AF_INET6, sin6_port=htons(0), inet_pton(AF_INET6, "::1", &sin6_addr), sin6_flowinfo=0, sin6_scope_id=0}, 28 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... bind resumed> )        = 0
[pid 20018] socket(PF_INET6, SOCK_STREAM, IPPROTO_TCP) = 4
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] setsockopt(4, SOL_IPV6, IPV6_V6ONLY, [0], 4 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... setsockopt resumed> )  = 0
[pid 20018] bind(4, {sa_family=AF_INET6, sin6_port=htons(0), inet_pton(AF_INET6, "::ffff:127.0.0.1", &sin6_addr), sin6_flowinfo=0, sin6_scope_id=0}, 28 <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... bind resumed> )        = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] close(4)                    = 0
[pid 20018] close(3 <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... close resumed> )       = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] futex(0xa80cb8, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xa80cb8, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] stat("/home/qboxtest/.rvm/gems/ruby-2.1.3/bin/gccgo",  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... stat resumed> 0xc208044120) = -1 ENOENT (No such file or directory)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] stat("/home/qboxtest/.rvm/gems/ruby-2.1.3@global/bin/gccgo", 0xc2080441b0) = -1 ENOENT (No such file or directory)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] stat("/home/qboxtest/.rvm/rubies/ruby-2.1.3/bin/gccgo",  <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... stat resumed> 0xc208044240) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/home/qboxtest/bin/gccgo",  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... stat resumed> 0xc2080442d0) = -1 ENOENT (No such file or directory)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] stat("/usr/lib/lightdm/lightdm/gccgo", 0xc208044360) = -1 ENOENT (No such file or directory)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] stat("/usr/local/sbin/gccgo",  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20021] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20021] <... futex resumed> )       = 1
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... stat resumed> 0xc2080443f0) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/usr/local/bin/gccgo",  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... stat resumed> 0xc208044480) = -1 ENOENT (No such file or directory)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] stat("/usr/sbin/gccgo", 0xc208044510) = -1 ENOENT (No such file or directory)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] stat("/usr/bin/gccgo",  <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... stat resumed> 0xc2080445a0) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/sbin/gccgo",  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... stat resumed> 0xc208044630) = -1 ENOENT (No such file or directory)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] stat("/bin/gccgo", 0xc2080446c0) = -1 ENOENT (No such file or directory)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] stat("/usr/games/gccgo",  <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... stat resumed> 0xc208044750) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/home/qboxtest/.rvm/bin/gccgo",  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... stat resumed> 0xc2080447e0) = -1 ENOENT (No such file or directory)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] stat("/usr/local/go/bin/gccgo", 0xc208044870) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/home/qboxtest/qbox/ffmpeg/ffmpeg-static/target/bin/gccgo",  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... stat resumed> 0xc208044900) = -1 ENOENT (No such file or directory)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] stat("/home/qboxtest/qbox/base/qiniu/bin/gccgo", 0xc208044990) = -1 ENOENT (No such file or directory)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] stat("/home/qboxtest/qbox/base/biz/bin/gccgo",  <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... stat resumed> 0xc208044a20) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/home/qboxtest/qbox/base/cgo/bin/gccgo", 0xc208044ab0) = -1 ENOENT (No such file or directory)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] stat("/home/qboxtest/qbox/base/com/bin/gccgo",  <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... stat resumed> 0xc208044b40) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/home/qboxtest/qbox/base/mockacc/bin/gccgo", 0xc208044bd0) = -1 ENOENT (No such file or directory)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] stat("/home/qboxtest/qbox/image/fop_ncgo/bin/gccgo",  <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... stat resumed> 0xc208044c60) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/home/qboxtest/qbox/image/fop_cgo/bin/gccgo", 0xc208044cf0) = -1 ENOENT (No such file or directory)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] stat("/home/qboxtest/qbox/ffmpeg/fileop/bin/gccgo",  <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... stat resumed> 0xc208044d80) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/home/qboxtest/qbox/ffmpeg/fop/bin/gccgo", 0xc208044e10) = -1 ENOENT (No such file or directory)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] stat("/home/qboxtest/qbox/fileop/fileop/bin/gccgo",  <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... stat resumed> 0xc208044ea0) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/home/qboxtest/qbox/io/cgo_deps/bin/gccgo", 0xc208044f30) = -1 ENOENT (No such file or directory)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] stat("/home/qboxtest/qbox/io/io/bin/gccgo",  <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... stat resumed> 0xc208044fc0) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/home/qboxtest/qbox/service/service/bin/gccgo",  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... stat resumed> 0xc208045050) = -1 ENOENT (No such file or directory)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] stat("/home/qboxtest/qbox/service/service_cgo/bin/gccgo", 0xc2080450e0) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/home/qboxtest/qbox/fop/bin/gccgo",  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... stat resumed> 0xc208045170) = -1 ENOENT (No such file or directory)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] stat("/home/qboxtest/qbox/devtools/bin/gccgo", 0xc208045200) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/home/qboxtest/qbox/rs/rs/bin/gccgo",  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... stat resumed> 0xc208045290) = -1 ENOENT (No such file or directory)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] stat("/home/qboxtest/qbox/nio/bin/gccgo", 0xc208045320) = -1 ENOENT (No such file or directory)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] stat("/home/qboxtest/qbox/fileop/fileop/bin/gccgo",  <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... stat resumed> 0xc2080453b0) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/home/qboxtest/qbox/biz/bin/gccgo",  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... stat resumed> 0xc208045440) = -1 ENOENT (No such file or directory)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] stat("/home/qboxtest/qbox/office/fop_ncgo/bin/gccgo", 0xc2080454d0) = -1 ENOENT (No such file or directory)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] stat("/home/qboxtest/qbox/office/fop_cgo/bin/gccgo",  <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... stat resumed> 0xc208045560) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/home/qboxtest/qbox/api.qiniu.com/bin/gccgo",  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... stat resumed> 0xc2080455f0) = -1 ENOENT (No such file or directory)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] stat("/home/qboxtest/.rvm/bin/gccgo", 0xc208045680) = -1 ENOENT (No such file or directory)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] getcwd( <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... getcwd resumed> "/home/qboxtest/Desktop/qiniu", 4096) = 29
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... futex resumed> )       = 1
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] futex(0xa81b10, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] <... futex resumed> )       = 1
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] <... futex resumed> )       = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... futex resumed> )       = 0
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... futex resumed> )       = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] <... futex resumed> )       = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] rt_sigprocmask(SIG_SETMASK, ~[RTMIN RT_1],  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... rt_sigprocmask resumed> [], 8) = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] mmap(NULL, 8392704, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_STACK, -1, 0) = 0x7f33b897a000
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] mprotect(0x7f33b897a000, 4096, PROT_NONE <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... mprotect resumed> )    = 0
[pid 20018] clone( <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
Process 20025 attached (waiting for parent)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}Process 20025 resumed (parent 20018 ready)
 <unfinished ...>
[pid 20018] <... clone resumed> child_stack=0x7f33b9179ff0, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLS|CLONE_PARENT_SETTID|CLONE_CHILD_CLEARTID, parent_tidptr=0x7f33b917a9d0, tls=0x7f33b917a700, child_tidptr=0x7f33b917a9d0) = 20025
[pid 20025] set_robust_list(0x7f33b917a9e0, 0x18 <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20025] <... set_robust_list resumed> ) = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20025] sigaltstack({ss_sp=0xc2080b2000, ss_flags=0, ss_size=32768} <unfinished ...>
[pid 20018] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20025] <... sigaltstack resumed> , NULL) = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20025] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20025] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20018] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20025] <... futex resumed> )       = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xa81b10, FUTEX_WAIT, 0, NULL) = -1 EAGAIN (Resource temporarily unavailable)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xa81b10, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 40} <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 80} <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20018] <... futex resumed> )       = 0
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] <... futex resumed> )       = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 160} <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] <... futex resumed> )       = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 320} <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20025] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20025] <... futex resumed> )       = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 640} <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20025] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20025] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20025] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20025] <... futex resumed> )       = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 1280} <unfinished ...>
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1) = 0
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xa81b10, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20018] <... futex resumed> )       = 0
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xa81b10, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1) = 1
[pid 20018] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20025] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20021] select(0, NULL, NULL, NULL, {0, 2560} <unfinished ...>
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20018] <... futex resumed> )       = 0
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 5120} <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20018] <... futex resumed> )       = 0
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20025] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] rt_sigprocmask(SIG_SETMASK, ~[RTMIN RT_1], [], 8) = 0
[pid 20018] mmap(NULL, 8392704, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_STACK, -1, 0) = 0x7f33b8179000
[pid 20018] mprotect(0x7f33b8179000, 4096, PROT_NONE) = 0
[pid 20018] clone(Process 20026 attached (waiting for parent)
Process 20026 resumed (parent 20018 ready)
child_stack=0x7f33b8978ff0, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLS|CLONE_PARENT_SETTID|CLONE_CHILD_CLEARTID, parent_tidptr=0x7f33b89799d0, tls=0x7f33b8979700, child_tidptr=0x7f33b89799d0) = 20026
[pid 20026] set_robust_list(0x7f33b89799e0, 0x18 <unfinished ...>
[pid 20018] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20026] <... set_robust_list resumed> ) = 0
[pid 20018] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20026] sigaltstack({ss_sp=0xc2080ba000, ss_flags=0, ss_size=32768} <unfinished ...>
[pid 20018] futex(0xa81b10, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20026] <... sigaltstack resumed> , NULL) = 0
[pid 20026] rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
[pid 20026] futex(0xa81b10, FUTEX_WAKE, 1) = 1
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20026] futex(0xc20805ae70, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 10000} <unfinished ...>
[pid 20018] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805ae70, FUTEX_WAKE, 1 <unfinished ...>
[pid 20026] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20026] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20026] <... futex resumed> )       = 0
[pid 20025] <... futex resumed> )       = 0
[pid 20026] futex(0xc20805ae70, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20025] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xc20805ae70, FUTEX_WAKE, 1 <unfinished ...>
[pid 20026] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20026] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20026] <... futex resumed> )       = 0
[pid 20025] <... futex resumed> )       = 0
[pid 20026] futex(0xc20805ae70, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20025] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805ae70, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20026] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20026] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20026] <... futex resumed> )       = 0
[pid 20025] <... futex resumed> )       = 0
[pid 20026] futex(0xc20805ae70, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20025] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805ae70, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20026] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20026] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20026] <... futex resumed> )       = 0
[pid 20025] <... futex resumed> )       = 0
[pid 20026] futex(0xc20805ae70, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20025] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805ae70, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20026] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20026] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20026] <... futex resumed> )       = 0
[pid 20025] <... futex resumed> )       = 0
[pid 20026] futex(0xc20805ae70, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20025] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805ae70, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20026] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20026] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20026] <... futex resumed> )       = 0
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20026] futex(0xc20805ae70, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20025] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20025] <... futex resumed> )       = 0
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20025] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20025] <... futex resumed> )       = 0
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1) = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20018] futex(0xa81b10, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20025] <... futex resumed> )       = 0
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20025] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20025] <... futex resumed> )       = 0
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20025] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20025] <... futex resumed> )       = 0
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20025] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20025] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20018] <... futex resumed> )       = 0
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20024] <... futex resumed> )       = 0
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20024] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20018] <... futex resumed> )       = 0
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20025] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20025] <... futex resumed> )       = 0
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20025] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20025] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20018] <... futex resumed> )       = 1
[pid 20025] <... futex resumed> )       = 0
[pid 20018] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20024] futex(0xa81b10, FUTEX_WAKE, 1) = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20018] stat("/usr/local/go", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] getpid()                    = 20018
[pid 20018] mkdir("/tmp/go-build628409689", 0700) = 0
[pid 20018] stat("childProcess.go", {st_mode=S_IFREG|0664, st_size=331, ...}) = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 10000} <unfinished ...>
[pid 20018] lstat("/usr", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go/src", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/vips", 0xc208044a20) = -1 ENOENT (No such file or directory)
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Downloads", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Downloads/mine", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Downloads/mine/UfopExif", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Downloads/mine/UfopExif/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/qiniu", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/qiniu/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/docs", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/docs/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/com", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/com/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/biz", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/biz/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/portal", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/portal/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/cgo", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/cgo/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/mockacc", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/mockacc/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/image", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/image/fop_ncgo", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/image/fop_ncgo/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/image", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/image/fop_cgo", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/image/fop_cgo/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/ffmpeg", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/ffmpeg/fileop", 0xc2080dc870) = -1 ENOENT (No such file or directory)
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/ffmpeg", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/ffmpeg/fop", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/ffmpeg/fop/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/fileop", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/fileop/fileop", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/fileop/fileop/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/io", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/io/simplebd", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/io/simplebd/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/io", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/io/logfs", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/io/logfs/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/io", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/io/cgo_deps", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/io/cgo_deps/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/io", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/io/proxy", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/io/proxy/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/io", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/io/io", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/io/io/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/service", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/service/service", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/service/service/src", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/service", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/service/service_cgo", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/service/service_cgo/src", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/fop", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/fop/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/devtools", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/devtools/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/com", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/com/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/qiniu", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/qiniu/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/biz", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/biz/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/cgo", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/base/cgo/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/rs", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/rs/rs", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/rs/rs/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/nio", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/nio/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/fileop", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/fileop/fileop", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/fileop/fileop/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/biz", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/biz/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox",  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... lstat resumed> {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 10000} <unfinished ...>
[pid 20018] lstat("/home/qboxtest/qbox/office", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/office/blackfriday", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/office/blackfriday/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/office", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/office/fop_ncgo", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/office/fop_ncgo/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/office", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/office/fop_cgo", 0xc2080f23f0) = -1 ENOENT (No such file or directory)
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/api.qiniu.com", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/qbox/api.qiniu.com/src", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] lstat("/home", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest", {st_mode=S_IFDIR|0777, st_size=4096, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop", {st_mode=S_IFDIR|0755, st_size=61440, ...}) = 0
[pid 20018] lstat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/home/qboxtest/Desktop/qiniu/childProcess.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "package main\n\nimport (\n\t\"fmt\"\n\t\""..., 4096) = 331
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/fmt", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/fmt", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 10 entries */, 4096) = 320
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/fmt/stringer_test.go", {st_mode=S_IFREG|0644, st_size=2156, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/fmt/scan_test.go", {st_mode=S_IFREG|0644, st_size=26764, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/fmt/doc.go", {st_mode=S_IFREG|0644, st_size=11299, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/fmt/format.go", {st_mode=S_IFREG|0644, st_size=11646, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/fmt/export_test.go", {st_mode=S_IFREG|0644, st_size=196, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/fmt/print.go", {st_mode=S_IFREG|0644, st_size=30177, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/fmt/scan.go", {st_mode=S_IFREG|0644, st_size=30612, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/fmt/fmt_test.go", {st_mode=S_IFREG|0644, st_size=32299, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/fmt/doc.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] mmap(0xc208100000, 1048576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0xc208100000
[pid 20018] mmap(0xc207fe0000, 65536, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0xc207fe0000
[pid 20018] read(3, "ings or slices in hex (% x, % X)"..., 4096) = 4096
[pid 20018] read(3, ", decorating it with an indicati"..., 4096) = 3107
[pid 20018] read(3, "", 4096)           = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/fmt/export_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 196
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/fmt/fmt_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/fmt/format.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/fmt/print.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/fmt/scan.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/fmt/scan_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/fmt/stringer_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 2156
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/errors", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/errors", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 5 entries */, 4096) = 160
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/errors/errors.go", {st_mode=S_IFREG|0644, st_size=499, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/errors/example_test.go", {st_mode=S_IFREG|0644, st_size=692, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/errors/errors_test.go", {st_mode=S_IFREG|0644, st_size=1271, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/errors/errors.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 499
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/errors/errors_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 1271
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/errors/example_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 692
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/runtime", {st_mode=S_IFDIR|0755, st_size=16384, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 107 entries */, 4096) = 4096
[pid 20018] getdents64(3, /* 109 entries */, 4096) = 4096
[pid 20018] getdents64(3, /* 109 entries */, 4096) = 4072
[pid 20018] getdents64(3, /* 7 entries */, 4096) = 264
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_netbsd_amd64.h", {st_mode=S_IFREG|0644, st_size=3107, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zchan_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=30493, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/lfstack_test.go", {st_mode=S_IFREG|0644, st_size=2781, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zalg_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=15040, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mem_netbsd.c", {st_mode=S_IFREG|0644, st_size=2129, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/runtime_linux_test.go", {st_mode=S_IFREG|0644, st_size=695, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/netpoll.goc", {st_mode=S_IFREG|0644, st_size=12253, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_linux_arm.h", {st_mode=S_IFREG|0644, st_size=1397, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/cgocall.h", {st_mode=S_IFREG|0644, st_size=358, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/malloc.h", {st_mode=S_IFREG|0644, st_size=22646, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_freebsd_amd64.s", {st_mode=S_IFREG|0644, st_size=7457, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_windows.h", {st_mode=S_IFREG|0644, st_size=1203, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/runtime.h", {st_mode=S_IFREG|0644, st_size=32641, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/syscall_windows_test.go", {st_mode=S_IFREG|0644, st_size=10150, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_netbsd_386.s", {st_mode=S_IFREG|0644, st_size=7443, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_nacl_386.s", {st_mode=S_IFREG|0644, st_size=554, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_freebsd.go", {st_mode=S_IFREG|0644, st_size=2977, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_dragonfly.go", {st_mode=S_IFREG|0644, st_size=2681, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_darwin_amd64.h", {st_mode=S_IFREG|0644, st_size=6919, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_netbsd_arm.s", {st_mode=S_IFREG|0644, st_size=8028, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mgc0.go", {st_mode=S_IFREG|0644, st_size=608, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_openbsd_amd64.h", {st_mode=S_IFREG|0644, st_size=1410, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_windows_amd64.s", {st_mode=S_IFREG|0644, st_size=469, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/race0.c", {st_mode=S_IFREG|0644, st_size=1504, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/race", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_nacl_386.h", {st_mode=S_IFREG|0644, st_size=878, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/arch_amd64.h", {st_mode=S_IFREG|0644, st_size=420, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mem_freebsd.c", {st_mode=S_IFREG|0644, st_size=2129, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/cgo", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_arm.c", {st_mode=S_IFREG|0644, st_size=3910, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/traceback_arm.c", {st_mode=S_IFREG|0644, st_size=11373, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/gc_test.go", {st_mode=S_IFREG|0644, st_size=4041, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_nacl_386.h", {st_mode=S_IFREG|0644, st_size=1183, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/proc_test.go", {st_mode=S_IFREG|0644, st_size=8810, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_darwin_amd64.s", {st_mode=S_IFREG|0644, st_size=11153, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mapspeed_test.go", {st_mode=S_IFREG|0644, st_size=6063, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_solaris.h", {st_mode=S_IFREG|0644, st_size=1542, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_linux_386.h", {st_mode=S_IFREG|0644, st_size=3444, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_solaris.c", {st_mode=S_IFREG|0644, st_size=16021, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_freebsd_amd64.h", {st_mode=S_IFREG|0644, st_size=1415, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/chan.h", {st_mode=S_IFREG|0644, st_size=1728, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/memmove_plan9_amd64.s", {st_mode=S_IFREG|0644, st_size=2920, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_solaris_amd64.h", {st_mode=S_IFREG|0644, st_size=4147, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/vlrt_386.c", {st_mode=S_IFREG|0644, st_size=12212, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zslice_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=5563, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_solaris_amd64.h", {st_mode=S_IFREG|0644, st_size=1587, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/stack_gen_test.go", {st_mode=S_IFREG|0644, st_size=130106, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/syscall_windows.goc", {st_mode=S_IFREG|0644, st_size=3115, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_windows_386.s", {st_mode=S_IFREG|0644, st_size=7537, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_linux_arm.h", {st_mode=S_IFREG|0644, st_size=2835, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/env_plan9.c", {st_mode=S_IFREG|0644, st_size=922, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/env_posix.c", {st_mode=S_IFREG|0644, st_size=1508, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/vlrt_arm.c", {st_mode=S_IFREG|0644, st_size=11652, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zmalloc_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=25523, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/race_amd64.s", {st_mode=S_IFREG|0644, st_size=7950, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_darwin_386.s", {st_mode=S_IFREG|0644, st_size=363, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_netbsd_arm.h", {st_mode=S_IFREG|0644, st_size=2900, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_freebsd_arm.s", {st_mode=S_IFREG|0644, st_size=462, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signals_solaris.h", {st_mode=S_IFREG|0644, st_size=3683, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/memmove_arm.s", {st_mode=S_IFREG|0644, st_size=6306, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zsigqueue_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=3257, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_plan9.c", {st_mode=S_IFREG|0644, st_size=7640, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_darwin_amd64.s", {st_mode=S_IFREG|0644, st_size=369, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_netbsd_386.h", {st_mode=S_IFREG|0644, st_size=1136, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signals_freebsd.h", {st_mode=S_IFREG|0644, st_size=1737, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/norace_test.go", {st_mode=S_IFREG|0644, st_size=980, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_amd64x.c", {st_mode=S_IFREG|0644, st_size=5346, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/netpoll_solaris.c", {st_mode=S_IFREG|0644, st_size=8795, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/Makefile", {st_mode=S_IFREG|0644, st_size=181, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zrdebug_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=961, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_dragonfly_386.h", {st_mode=S_IFREG|0644, st_size=3104, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/slice.goc", {st_mode=S_IFREG|0644, st_size=5087, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/string_test.go", {st_mode=S_IFREG|0644, st_size=1563, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/memmove_linux_amd64_test.go", {st_mode=S_IFREG|0644, st_size=1637, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zcpuprof_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=8629, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/softfloat_arm.c", {st_mode=S_IFREG|0644, st_size=14783, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_linux_amd64.h", {st_mode=S_IFREG|0644, st_size=4130, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/type.go", {st_mode=S_IFREG|0644, st_size=1082, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/symtab_test.go", {st_mode=S_IFREG|0644, st_size=1021, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_unix.h", {st_mode=S_IFREG|0644, st_size=463, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/stack_test.go", {st_mode=S_IFREG|0644, st_size=6786, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/futex_test.go", {st_mode=S_IFREG|0644, st_size=2006, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_dragonfly_amd64.h", {st_mode=S_IFREG|0644, st_size=1415, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/typekind.h", {st_mode=S_IFREG|0644, st_size=691, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_netbsd_386.go", {st_mode=S_IFREG|0644, st_size=854, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/time.goc", {st_mode=S_IFREG|0644, st_size=6840, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mallocrand.go", {st_mode=S_IFREG|0644, st_size=2001, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/softfloat64.go", {st_mode=S_IFREG|0644, st_size=9705, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mem.go", {st_mode=S_IFREG|0644, st_size=2353, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/arch_amd64p32.h", {st_mode=S_IFREG|0644, st_size=337, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_plan9_386.c", {st_mode=S_IFREG|0644, st_size=3595, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_plan9_amd64.s", {st_mode=S_IFREG|0644, st_size=377, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_windows_amd64.c", {st_mode=S_IFREG|0644, st_size=4368, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_windows_386.s", {st_mode=S_IFREG|0644, st_size=466, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/malloc_test.go", {st_mode=S_IFREG|0644, st_size=2798, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/hashmap.goc", {st_mode=S_IFREG|0644, st_size=29493, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sigqueue.goc", {st_mode=S_IFREG|0644, st_size=4458, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/memmove_nacl_amd64p32.s", {st_mode=S_IFREG|0644, st_size=642, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signals_netbsd.h", {st_mode=S_IFREG|0644, st_size=1747, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/hashmap_fast.c", {st_mode=S_IFREG|0644, st_size=6037, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zruntime1_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=4443, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/runtime_unix_test.go", {st_mode=S_IFREG|0644, st_size=1226, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_linux_386.h", {st_mode=S_IFREG|0644, st_size=979, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/memmove_amd64.s", {st_mode=S_IFREG|0644, st_size=4688, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_plan9_amd64.c", {st_mode=S_IFREG|0644, st_size=3919, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_dragonfly.h", {st_mode=S_IFREG|0644, st_size=824, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zstring_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=11063, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mfixalloc.c", {st_mode=S_IFREG|0644, st_size=1349, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zruntime_defs_linux_amd64.go", {st_mode=S_IFREG|0644, st_size=10730, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/float.c", {st_mode=S_IFREG|0644, st_size=342, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_plan9_386.h", {st_mode=S_IFREG|0644, st_size=595, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/chan.goc", {st_mode=S_IFREG|0644, st_size=24243, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_linux_386.s", {st_mode=S_IFREG|0644, st_size=528, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zgoos_linux.go", {st_mode=S_IFREG|0644, st_size=76, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/atomic_386.c", {st_mode=S_IFREG|0644, st_size=754, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mprof.goc", {st_mode=S_IFREG|0644, st_size=12227, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_nacl_amd64p32.h", {st_mode=S_IFREG|0644, st_size=1524, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signals_darwin.h", {st_mode=S_IFREG|0644, st_size=1704, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_darwin.h", {st_mode=S_IFREG|0644, st_size=1417, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/string.goc", {st_mode=S_IFREG|0644, st_size=7655, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zsema_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=6341, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_linux_386.c", {st_mode=S_IFREG|0644, st_size=861, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/compiler.go", {st_mode=S_IFREG|0644, st_size=441, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/error.go", {st_mode=S_IFREG|0644, st_size=3053, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zlfstack_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=2001, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mallocrep.go", {st_mode=S_IFREG|0644, st_size=1606, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/memclr_386.s", {st_mode=S_IFREG|0644, st_size=2307, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_openbsd_amd64.h", {st_mode=S_IFREG|0644, st_size=2770, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_dragonfly_386.s", {st_mode=S_IFREG|0644, st_size=7264, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_windows_amd64.s", {st_mode=S_IFREG|0644, st_size=8391, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_freebsd_amd64.s", {st_mode=S_IFREG|0644, st_size=370, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/memmove_386.s", {st_mode=S_IFREG|0644, st_size=3996, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/crash_cgo_test.go", {st_mode=S_IFREG|0644, st_size=1875, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mgc0.h", {st_mode=S_IFREG|0644, st_size=3546, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/netpoll_epoll.c", {st_mode=S_IFREG|0644, st_size=2254, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/noasm_arm.goc", {st_mode=S_IFREG|0644, st_size=1225, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_freebsd_386.h", {st_mode=S_IFREG|0644, st_size=3427, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_solaris_amd64.s", {st_mode=S_IFREG|0644, st_size=438, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/complex.goc", {st_mode=S_IFREG|0644, st_size=1620, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_nacl.h", {st_mode=S_IFREG|0644, st_size=6639, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/malloc1.go", {st_mode=S_IFREG|0644, st_size=478, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_freebsd.h", {st_mode=S_IFREG|0644, st_size=827, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/debug.go", {st_mode=S_IFREG|0644, st_size=7038, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_netbsd_386.h", {st_mode=S_IFREG|0644, st_size=2882, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_linux.go", {st_mode=S_IFREG|0644, st_size=2915, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_x86.c", {st_mode=S_IFREG|0644, st_size=1799, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rdebug.goc", {st_mode=S_IFREG|0644, st_size=624, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_openbsd.c", {st_mode=S_IFREG|0644, st_size=7270, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zgoarch_amd64.go", {st_mode=S_IFREG|0644, st_size=78, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_netbsd_arm.go", {st_mode=S_IFREG|0644, st_size=763, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_netbsd_amd64.s", {st_mode=S_IFREG|0644, st_size=7281, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/hash_test.go", {st_mode=S_IFREG|0644, st_size=11815, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zaexperiment.h", {st_mode=S_IFREG|0644, st_size=59, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_plan9_amd64.s", {st_mode=S_IFREG|0644, st_size=4015, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_openbsd_386.h", {st_mode=S_IFREG|0644, st_size=996, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_nacl_amd64p32.s", {st_mode=S_IFREG|0644, st_size=9028, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/race.c", {st_mode=S_IFREG|0644, st_size=6376, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/vlop_arm_test.go", {st_mode=S_IFREG|0644, st_size=2504, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_nacl_386.s", {st_mode=S_IFREG|0644, st_size=6341, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_plan9_386.s", {st_mode=S_IFREG|0644, st_size=754, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mheap.c", {st_mode=S_IFREG|0644, st_size=25263, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/ztime_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=7403, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_solaris_amd64.go", {st_mode=S_IFREG|0644, st_size=1003, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/runtime.c", {st_mode=S_IFREG|0644, st_size=8222, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/memmove_plan9_386.s", {st_mode=S_IFREG|0644, st_size=2914, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/netpoll_windows.c", {st_mode=S_IFREG|0644, st_size=4436, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_openbsd.go", {st_mode=S_IFREG|0644, st_size=2554, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_darwin_amd64.h", {st_mode=S_IFREG|0644, st_size=1356, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_dragonfly.c", {st_mode=S_IFREG|0644, st_size=6859, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_netbsd_386.c", {st_mode=S_IFREG|0644, st_size=529, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mem_solaris.c", {st_mode=S_IFREG|0644, st_size=2114, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_linux.c", {st_mode=S_IFREG|0644, st_size=8871, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_freebsd_arm.h", {st_mode=S_IFREG|0644, st_size=2987, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/stack.c", {st_mode=S_IFREG|0644, st_size=28481, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/iface.goc", {st_mode=S_IFREG|0644, st_size=12147, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/memclr_plan9_amd64.s", {st_mode=S_IFREG|0644, st_size=785, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/iface_test.go", {st_mode=S_IFREG|0644, st_size=1939, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zhashmap_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=29237, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/time_plan9_386.c", {st_mode=S_IFREG|0644, st_size=1058, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_openbsd_amd64.s", {st_mode=S_IFREG|0644, st_size=370, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs.c", {st_mode=S_IFREG|0644, st_size=401, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/parfor.c", {st_mode=S_IFREG|0644, st_size=4894, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/asm_amd64p32.s", {st_mode=S_IFREG|0644, st_size=22234, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_freebsd.c", {st_mode=S_IFREG|0644, st_size=7250, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/atomic_arm.c", {st_mode=S_IFREG|0644, st_size=2712, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/softfloat64_test.go", {st_mode=S_IFREG|0644, st_size=4141, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zsys_linux_amd64.s", {st_mode=S_IFREG|0644, st_size=35, ...}) = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/cpuprof.goc",  <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 10000} <unfinished ...>
[pid 20018] <... lstat resumed> {st_mode=S_IFREG|0644, st_size=11808, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_unix.c", {st_mode=S_IFREG|0644, st_size=2623, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_netbsd.go", {st_mode=S_IFREG|0644, st_size=2728, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_netbsd.h", {st_mode=S_IFREG|0644, st_size=826, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_dragonfly_amd64.s", {st_mode=S_IFREG|0644, st_size=6500, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/msize.c", {st_mode=S_IFREG|0644, st_size=6317, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/export_futex_test.go", {st_mode=S_IFREG|0644, st_size=367, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mgc0.c", {st_mode=S_IFREG|0644, st_size=78510, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_nacl.c", {st_mode=S_IFREG|0644, st_size=6216, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/heapdump.c", {st_mode=S_IFREG|0644, st_size=22566, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/callback_windows.c", {st_mode=S_IFREG|0644, st_size=2163, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/asm_amd64.s", {st_mode=S_IFREG|0644, st_size=41507, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signals_nacl.h", {st_mode=S_IFREG|0644, st_size=1704, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/export_test.go", {st_mode=S_IFREG|0644, st_size=2153, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_freebsd_arm.s", {st_mode=S_IFREG|0644, st_size=8368, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/netpoll_nacl.c", {st_mode=S_IFREG|0644, st_size=611, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_darwin.go", {st_mode=S_IFREG|0644, st_size=4811, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_linux_arm.s", {st_mode=S_IFREG|0644, st_size=8833, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/ziface_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=16536, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_linux_amd64.s", {st_mode=S_IFREG|0644, st_size=368, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_windows_386.c", {st_mode=S_IFREG|0644, st_size=4083, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/panic.c", {st_mode=S_IFREG|0644, st_size=14794, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/symtab.goc", {st_mode=S_IFREG|0644, st_size=7825, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/closure_test.go", {st_mode=S_IFREG|0644, st_size=936, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mem_windows.c", {st_mode=S_IFREG|0644, st_size=3497, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_windows.c", {st_mode=S_IFREG|0644, st_size=14470, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zcomplex_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=1617, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_freebsd_386.s", {st_mode=S_IFREG|0644, st_size=7105, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_nacl_amd64p32.s", {st_mode=S_IFREG|0644, st_size=875, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/runtime-gdb.py", {st_mode=S_IFREG|0644, st_size=11993, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sema.goc", {st_mode=S_IFREG|0644, st_size=6471, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_windows.go", {st_mode=S_IFREG|0644, st_size=1768, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_plan9_amd64.h", {st_mode=S_IFREG|0644, st_size=523, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/lfstack.goc", {st_mode=S_IFREG|0644, st_size=1927, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_arm.c", {st_mode=S_IFREG|0644, st_size=1137, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_linux_386.s", {st_mode=S_IFREG|0644, st_size=10087, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/syscall_nacl.h", {st_mode=S_IFREG|0644, st_size=1844, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/syscall_solaris.goc", {st_mode=S_IFREG|0644, st_size=8139, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_solaris_amd64.s", {st_mode=S_IFREG|0644, st_size=5909, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mcache.c", {st_mode=S_IFREG|0644, st_size=3214, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/asm_arm.s", {st_mode=S_IFREG|0644, st_size=28207, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/complex_test.go", {st_mode=S_IFREG|0644, st_size=1077, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_darwin_386.h", {st_mode=S_IFREG|0644, st_size=6859, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mfinal_test.go", {st_mode=S_IFREG|0644, st_size=5421, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/arch_arm.h", {st_mode=S_IFREG|0644, st_size=284, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/stack.h", {st_mode=S_IFREG|0644, st_size=4094, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signals_openbsd.h", {st_mode=S_IFREG|0644, st_size=1747, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_netbsd_amd64.s", {st_mode=S_IFREG|0644, st_size=369, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_netbsd.c", {st_mode=S_IFREG|0644, st_size=8268, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_openbsd.h", {st_mode=S_IFREG|0644, st_size=696, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/type.h", {st_mode=S_IFREG|0644, st_size=1523, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signals_dragonfly.h", {st_mode=S_IFREG|0644, st_size=1737, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_darwin_386.s", {st_mode=S_IFREG|0644, st_size=12138, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zmprof_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=13348, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/memclr_plan9_386.s", {st_mode=S_IFREG|0644, st_size=823, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_openbsd_386.s", {st_mode=S_IFREG|0644, st_size=7629, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mem_darwin.c", {st_mode=S_IFREG|0644, st_size=1542, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zversion.go", {st_mode=S_IFREG|0644, st_size=119, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_linux.h", {st_mode=S_IFREG|0644, st_size=1103, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs1_linux.go", {st_mode=S_IFREG|0644, st_size=775, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_nacl_amd64p32.h", {st_mode=S_IFREG|0644, st_size=1268, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_openbsd_386.h", {st_mode=S_IFREG|0644, st_size=2592, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs2_linux.go", {st_mode=S_IFREG|0644, st_size=3502, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/alg.goc", {st_mode=S_IFREG|0644, st_size=11434, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/traceback_x86.c", {st_mode=S_IFREG|0644, st_size=14397, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/race.go", {st_mode=S_IFREG|0644, st_size=818, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_plan9_386.s", {st_mode=S_IFREG|0644, st_size=3466, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/chan_test.go", {st_mode=S_IFREG|0644, st_size=13349, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/memclr_amd64.s", {st_mode=S_IFREG|0644, st_size=2165, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_linux_amd64.h", {st_mode=S_IFREG|0644, st_size=1399, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_dragonfly_386.h", {st_mode=S_IFREG|0644, st_size=1003, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/pprof", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_solaris.go", {st_mode=S_IFREG|0644, st_size=3268, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_arm_linux.go", {st_mode=S_IFREG|0644, st_size=2749, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_freebsd_arm.c", {st_mode=S_IFREG|0644, st_size=644, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/lock_sema.c", {st_mode=S_IFREG|0644, st_size=6392, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/memmove_test.go", {st_mode=S_IFREG|0644, st_size=5968, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/append_test.go", {st_mode=S_IFREG|0644, st_size=4461, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_linux_amd64.s", {st_mode=S_IFREG|0644, st_size=7643, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/print.c", {st_mode=S_IFREG|0644, st_size=6170, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/atomic_amd64x.c", {st_mode=S_IFREG|0644, st_size=528, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mknacl.sh", {st_mode=S_IFREG|0644, st_size=482, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_darwin_386.h", {st_mode=S_IFREG|0644, st_size=968, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mem_plan9.c", {st_mode=S_IFREG|0644, st_size=1585, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/runtime_test.go", {st_mode=S_IFREG|0644, st_size=4952, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/vlop_arm.s", {st_mode=S_IFREG|0644, st_size=8933, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/malloc.goc", {st_mode=S_IFREG|0644, st_size=28485, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/extern.go", {st_mode=S_IFREG|0644, st_size=9138, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/hashmap.h", {st_mode=S_IFREG|0644, st_size=6362, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_freebsd_386.s", {st_mode=S_IFREG|0644, st_size=364, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sqrt.go", {st_mode=S_IFREG|0644, st_size=5014, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_netbsd_arm.c", {st_mode=S_IFREG|0644, st_size=965, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mem_linux.c", {st_mode=S_IFREG|0644, st_size=3908, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_dragonfly_amd64.s", {st_mode=S_IFREG|0644, st_size=372, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/vdso_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=9380, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_openbsd_386.s", {st_mode=S_IFREG|0644, st_size=364, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/arch_386.h", {st_mode=S_IFREG|0644, st_size=337, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signals_plan9.h", {st_mode=S_IFREG|0644, st_size=1783, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/asm_386.s", {st_mode=S_IFREG|0644, st_size=41440, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_netbsd_amd64.go", {st_mode=S_IFREG|0644, st_size=1035, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_freebsd_386.h", {st_mode=S_IFREG|0644, st_size=1003, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signals_windows.h", {st_mode=S_IFREG|0644, st_size=159, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_386.c", {st_mode=S_IFREG|0644, st_size=3594, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_linux_arm.s", {st_mode=S_IFREG|0644, st_size=2552, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/memclr_arm.s", {st_mode=S_IFREG|0644, st_size=2483, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signals_linux.h", {st_mode=S_IFREG|0644, st_size=2606, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/crash_test.go", {st_mode=S_IFREG|0644, st_size=6837, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/funcdata.h", {st_mode=S_IFREG|0644, st_size=979, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/netpoll_stub.c", {st_mode=S_IFREG|0644, st_size=438, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mallocrep1.go", {st_mode=S_IFREG|0644, st_size=2810, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/proc.c", {st_mode=S_IFREG|0644, st_size=83852, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/parfor_test.go", {st_mode=S_IFREG|0644, st_size=3134, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_windows_386.h", {st_mode=S_IFREG|0644, st_size=2272, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/map_test.go", {st_mode=S_IFREG|0644, st_size=10087, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rune.c", {st_mode=S_IFREG|0644, st_size=5469, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/netpoll_kqueue.c", {st_mode=S_IFREG|0644, st_size=2397, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/sys_openbsd_amd64.s", {st_mode=S_IFREG|0644, st_size=6975, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mem_dragonfly.c", {st_mode=S_IFREG|0644, st_size=2436, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/runtime1.goc", {st_mode=S_IFREG|0644, st_size=2937, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_netbsd_amd64.c", {st_mode=S_IFREG|0644, st_size=581, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_netbsd_arm.h", {st_mode=S_IFREG|0644, st_size=1528, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/cgocall.c", {st_mode=S_IFREG|0644, st_size=9325, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zsymtab_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=9665, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/lock_futex.c", {st_mode=S_IFREG|0644, st_size=4842, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/vlop_386.s", {st_mode=S_IFREG|0644, st_size=1972, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_linux_arm.c", {st_mode=S_IFREG|0644, st_size=2425, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/race.h", {st_mode=S_IFREG|0644, st_size=1356, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mcentral.c", {st_mode=S_IFREG|0644, st_size=7756, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/znetpoll_linux_amd64.c", {st_mode=S_IFREG|0644, st_size=11602, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/debug", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_netbsd_amd64.h", {st_mode=S_IFREG|0644, st_size=1630, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_freebsd_amd64.h", {st_mode=S_IFREG|0644, st_size=3610, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/zasm_linux_amd64.h", {st_mode=S_IFREG|0644, st_size=7930, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mem_openbsd.c", {st_mode=S_IFREG|0644, st_size=2129, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_netbsd_386.s", {st_mode=S_IFREG|0644, st_size=363, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/signal_freebsd_arm.h", {st_mode=S_IFREG|0644, st_size=1382, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/mem_nacl.c", {st_mode=S_IFREG|0644, st_size=2741, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_windows_amd64.h", {st_mode=S_IFREG|0644, st_size=2472, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_plan9.h", {st_mode=S_IFREG|0644, st_size=2298, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/os_darwin.c", {st_mode=S_IFREG|0644, st_size=11794, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_netbsd_arm.s", {st_mode=S_IFREG|0644, st_size=383, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/rt0_dragonfly_386.s", {st_mode=S_IFREG|0644, st_size=366, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/runtime/defs_dragonfly_amd64.h", {st_mode=S_IFREG|0644, st_size=3289, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/append_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/arch_amd64.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 420
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/asm_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/atomic_amd64x.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 528
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/cgocall.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/cgocall.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 358
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/chan.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1728
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/chan_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/closure_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 936
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/compiler.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 441
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/complex_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 1077
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/crash_cgo_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 1875
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/crash_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/debug.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/defs.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 401
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/defs1_linux.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 775
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/defs2_linux.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 3502
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/defs_arm_linux.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2749
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/defs_linux.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2915
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/defs_linux_amd64.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Created by cgo -cdefs - DO NO"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/env_posix.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 1508
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/error.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 3053
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/export_futex_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 367
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/export_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 2153
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/extern.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/float.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 342
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/funcdata.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 979
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/futex_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 2006
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/gc_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4041
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/hash_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/hashmap.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/hashmap_fast.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/heapdump.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2014 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/iface_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 1939
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/lfstack_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 2781
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/lock_futex.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/lock_sema.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/malloc.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/malloc1.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 478
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/malloc_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 2798
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/mallocrand.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2001
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/mallocrep.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1606
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/mallocrep1.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2810
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/map_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/mapspeed_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/mcache.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 3214
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/mcentral.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/mem.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2353
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/mem_linux.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 3908
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/memclr_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2014 The Go Authors"..., 4096) = 2165
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/memmove_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Derived from Inferno's libker"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/memmove_linux_amd64_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 1637
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/memmove_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/mfinal_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/mfixalloc.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1349
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/mgc0.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/mgc0.go", O_RDONLY|O_CLOEXEC <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 10000} <unfinished ...>
[pid 20018] <... open resumed> )        = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 608
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/mgc0.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 3546
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/mheap.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/msize.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/netpoll_epoll.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 2254
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/netpoll_kqueue.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 2397
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/netpoll_stub.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 438
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/norace_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 980
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/os_linux.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/os_linux.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1103
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/panic.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/parfor.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/parfor_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 3134
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/print.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/proc.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/proc_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/race.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/race.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 818
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/race.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 1356
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/race0.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 1504
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/race_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/rt0_linux_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 368
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/rune.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "/*\n * The authors of this softwa"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/runtime.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/runtime.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/runtime_linux_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 695
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/runtime_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/runtime_unix_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 1226
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/signal_amd64x.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/signal_linux_amd64.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 1399
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/signal_unix.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 2623
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/signal_unix.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 463
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/signals_linux.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2606
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/softfloat64.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/softfloat64_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/sqrt.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/stack.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/stack.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4094
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/stack_gen_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2014 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/stack_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/string_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 1563
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/symtab_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1021
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/sys_linux_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/sys_x86.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 1799
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/traceback_x86.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/type.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1082
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/type.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1523
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/typekind.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 691
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/vdso_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zaexperiment.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 59
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zalg_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zasm_linux_amd64.h", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zchan_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zcomplex_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 1617
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zcpuprof_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zgoarch_amd64.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 78
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zgoos_linux.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 76
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zhashmap_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/ziface_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zlfstack_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 2001
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zmalloc_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zmprof_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/znetpoll_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zrdebug_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 961
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zruntime1_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zruntime_defs_linux_amd64.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zsema_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zsigqueue_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 3257
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zslice_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zstring_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zsymtab_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zsys_linux_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 35
[pid 20018] read(3, "", 4096)           = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/ztime_linux_amd64.c", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/runtime/zversion.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// auto generated by go tool dis"..., 4096) = 119
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/unsafe", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/unsafe", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 3 entries */, 4096) = 80
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/unsafe/unsafe.go", {st_mode=S_IFREG|0644, st_size=1886, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/unsafe/unsafe.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1886
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/io", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/io", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 9 entries */, 4096) = 280
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/io/ioutil", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/io/multi.go", {st_mode=S_IFREG|0644, st_size=1479, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/io/io.go", {st_mode=S_IFREG|0644, st_size=15473, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/io/pipe_test.go", {st_mode=S_IFREG|0644, st_size=5930, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/io/multi_test.go", {st_mode=S_IFREG|0644, st_size=2822, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/io/pipe.go", {st_mode=S_IFREG|0644, st_size=4570, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/io/io_test.go", {st_mode=S_IFREG|0644, st_size=9531, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/io/io.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/io/io_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/io/multi.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 1479
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/io/multi_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 2822
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/io/pipe.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/io/pipe_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/sync", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/sync", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 21 entries */, 4096) = 704
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/pool_test.go", {st_mode=S_IFREG|0644, st_size=2823, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/cond.go", {st_mode=S_IFREG|0644, st_size=2777, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/cond_test.go", {st_mode=S_IFREG|0644, st_size=3947, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/rwmutex.go", {st_mode=S_IFREG|0644, st_size=3603, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/waitgroup_test.go", {st_mode=S_IFREG|0644, st_size=2704, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/pool.go", {st_mode=S_IFREG|0644, st_size=6513, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/runtime_sema_test.go", {st_mode=S_IFREG|0644, st_size=1318, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/race0.go", {st_mode=S_IFREG|0644, st_size=502, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/mutex.go", {st_mode=S_IFREG|0644, st_size=2698, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/export_test.go", {st_mode=S_IFREG|0644, st_size=286, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/rwmutex_test.go", {st_mode=S_IFREG|0644, st_size=4330, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/example_test.go", {st_mode=S_IFREG|0644, st_size=1165, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/once_test.go", {st_mode=S_IFREG|0644, st_size=1181, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/runtime.go", {st_mode=S_IFREG|0644, st_size=1219, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/race.go", {st_mode=S_IFREG|0644, st_size=691, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/waitgroup.go", {st_mode=S_IFREG|0644, st_size=3821, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/atomic", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/mutex_test.go", {st_mode=S_IFREG|0644, st_size=2295, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/once.go", {st_mode=S_IFREG|0644, st_size=1228, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/cond.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 2777
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/cond_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 3947
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/example_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 1165
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/export_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 286
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/mutex.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2698
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/mutex_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2295
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/once.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1228
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/once_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1181
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/pool.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/pool_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 2823
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/race.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 691
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/race0.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 502
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/runtime.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 1219
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/runtime_sema_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1318
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/rwmutex.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 3603
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/rwmutex_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/waitgroup.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 3821
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/waitgroup_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 2704
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/sync/atomic", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/atomic", O_RDONLY|O_CLOEXEC <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... open resumed> )        = 3
[pid 20021] select(0, NULL, NULL, NULL, {0, 10000} <unfinished ...>
[pid 20018] getdents64(3, /* 15 entries */, 4096) = 536
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/atomic/export_linux_arm_test.go", {st_mode=S_IFREG|0644, st_size=271, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/atomic/64bit_arm.go", {st_mode=S_IFREG|0644, st_size=768, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/atomic/doc.go", {st_mode=S_IFREG|0644, st_size=5672, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/atomic/asm_netbsd_arm.s", {st_mode=S_IFREG|0644, st_size=2243, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/atomic/asm_freebsd_arm.s", {st_mode=S_IFREG|0644, st_size=2244, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/atomic/asm_amd64p32.s", {st_mode=S_IFREG|0644, st_size=3070, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/atomic/asm_linux_arm.s", {st_mode=S_IFREG|0644, st_size=5524, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/atomic/asm_amd64.s", {st_mode=S_IFREG|0644, st_size=2774, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/atomic/atomic_test.go", {st_mode=S_IFREG|0644, st_size=37027, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/atomic/asm_arm.s", {st_mode=S_IFREG|0644, st_size=4306, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/atomic/race.go", {st_mode=S_IFREG|0644, st_size=7040, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/atomic/asm_386.s", {st_mode=S_IFREG|0644, st_size=4400, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sync/atomic/atomic_linux_arm_test.go", {st_mode=S_IFREG|0644, st_size=305, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/atomic/asm_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 2774
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/atomic/atomic_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/atomic/doc.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sync/atomic/race.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/math", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/math", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 121 entries */, 4096) = 4080
[pid 20018] getdents64(3, /* 16 entries */, 4096) = 552
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/atan.go", {st_mode=S_IFREG|0644, st_size=3043, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/expm1_386.s", {st_mode=S_IFREG|0644, st_size=1900, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/asin_amd64.s", {st_mode=S_IFREG|0644, st_size=284, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/mod_arm.s", {st_mode=S_IFREG|0644, st_size=236, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/floor_amd64p32.s", {st_mode=S_IFREG|0644, st_size=186, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/asin.go", {st_mode=S_IFREG|0644, st_size=983, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/log_386.s", {st_mode=S_IFREG|0644, st_size=408, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/sqrt_386.s", {st_mode=S_IFREG|0644, st_size=311, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/abs_amd64.s", {st_mode=S_IFREG|0644, st_size=373, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/tan_amd64p32.s", {st_mode=S_IFREG|0644, st_size=184, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/hypot_386.s", {st_mode=S_IFREG|0644, st_size=1862, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/dim_amd64p32.s", {st_mode=S_IFREG|0644, st_size=184, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/tanh.go", {st_mode=S_IFREG|0644, st_size=2625, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/cmplx", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/hypot_arm.s", {st_mode=S_IFREG|0644, st_size=240, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/sin_amd64.s", {st_mode=S_IFREG|0644, st_size=280, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/log.go", {st_mode=S_IFREG|0644, st_size=3891, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/log_amd64.s", {st_mode=S_IFREG|0644, st_size=3714, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/ldexp_amd64p32.s", {st_mode=S_IFREG|0644, st_size=186, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/dim_arm.s", {st_mode=S_IFREG|0644, st_size=316, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/atan_arm.s", {st_mode=S_IFREG|0644, st_size=238, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/asinh.go", {st_mode=S_IFREG|0644, st_size=1869, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/exp2_arm.s", {st_mode=S_IFREG|0644, st_size=238, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/remainder_386.s", {st_mode=S_IFREG|0644, st_size=556, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/sincos.go", {st_mode=S_IFREG|0644, st_size=1862, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/tan_arm.s", {st_mode=S_IFREG|0644, st_size=236, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/atan_amd64.s", {st_mode=S_IFREG|0644, st_size=240, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/exp2_amd64.s", {st_mode=S_IFREG|0644, st_size=240, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/j1.go", {st_mode=S_IFREG|0644, st_size=13599, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/frexp.go", {st_mode=S_IFREG|0644, st_size=860, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/frexp_386.s", {st_mode=S_IFREG|0644, st_size=703, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/mod.go", {st_mode=S_IFREG|0644, st_size=874, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/abs.go", {st_mode=S_IFREG|0644, st_size=440, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/log1p_arm.s", {st_mode=S_IFREG|0644, st_size=240, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/exp_amd64.s", {st_mode=S_IFREG|0644, st_size=2920, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/atan2_arm.s", {st_mode=S_IFREG|0644, st_size=240, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/floor.go", {st_mode=S_IFREG|0644, st_size=1084, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/frexp_arm.s", {st_mode=S_IFREG|0644, st_size=240, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/ldexp_386.s", {st_mode=S_IFREG|0644, st_size=455, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/gamma.go", {st_mode=S_IFREG|0644, st_size=4957, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/log_arm.s", {st_mode=S_IFREG|0644, st_size=236, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/asin_amd64p32.s", {st_mode=S_IFREG|0644, st_size=185, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/ldexp.go", {st_mode=S_IFREG|0644, st_size=990, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/atanh.go", {st_mode=S_IFREG|0644, st_size=1943, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/ldexp_arm.s", {st_mode=S_IFREG|0644, st_size=240, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/floor_386.s", {st_mode=S_IFREG|0644, st_size=1499, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/remainder_amd64p32.s", {st_mode=S_IFREG|0644, st_size=190, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/abs_amd64p32.s", {st_mode=S_IFREG|0644, st_size=184, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/atan2.go", {st_mode=S_IFREG|0644, st_size=1485, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/modf.go", {st_mode=S_IFREG|0644, st_size=776, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/copysign.go", {st_mode=S_IFREG|0644, st_size=378, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/mod_amd64p32.s", {st_mode=S_IFREG|0644, st_size=184, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/sin_amd64p32.s", {st_mode=S_IFREG|0644, st_size=184, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/exp_386.s", {st_mode=S_IFREG|0644, st_size=1360, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/expm1_amd64.s", {st_mode=S_IFREG|0644, st_size=242, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/exp.go", {st_mode=S_IFREG|0644, st_size=5375, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/exp2_amd64p32.s", {st_mode=S_IFREG|0644, st_size=185, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/jn.go", {st_mode=S_IFREG|0644, st_size=7415, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/sincos_arm.s", {st_mode=S_IFREG|0644, st_size=242, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/sin_arm.s", {st_mode=S_IFREG|0644, st_size=276, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/hypot.go", {st_mode=S_IFREG|0644, st_size=798, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/expm1_amd64p32.s", {st_mode=S_IFREG|0644, st_size=186, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/unsafe.go", {st_mode=S_IFREG|0644, st_size=859, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/mod_386.s", {st_mode=S_IFREG|0644, st_size=544, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/ldexp_amd64.s", {st_mode=S_IFREG|0644, st_size=242, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/sincos_386.s", {st_mode=S_IFREG|0644, st_size=994, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/pow.go", {st_mode=S_IFREG|0644, st_size=2656, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/big", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/sinh.go", {st_mode=S_IFREG|0644, st_size=1535, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/j0.go", {st_mode=S_IFREG|0644, st_size=13919, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/lgamma.go", {st_mode=S_IFREG|0644, st_size=11263, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/log10_amd64.s", {st_mode=S_IFREG|0644, st_size=286, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/modf_arm.s", {st_mode=S_IFREG|0644, st_size=238, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/modf_amd64.s", {st_mode=S_IFREG|0644, st_size=240, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/bits.go", {st_mode=S_IFREG|0644, st_size=1835, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/remainder_amd64.s", {st_mode=S_IFREG|0644, st_size=250, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/acosh.go", {st_mode=S_IFREG|0644, st_size=1735, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/frexp_amd64.s", {st_mode=S_IFREG|0644, st_size=242, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/sincos_amd64p32.s", {st_mode=S_IFREG|0644, st_size=187, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/sqrt_amd64.s", {st_mode=S_IFREG|0644, st_size=302, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/const.go", {st_mode=S_IFREG|0644, st_size=2031, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/log_amd64p32.s", {st_mode=S_IFREG|0644, st_size=184, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/floor_arm.s", {st_mode=S_IFREG|0644, st_size=326, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/tan_amd64.s", {st_mode=S_IFREG|0644, st_size=238, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/atan_amd64p32.s", {st_mode=S_IFREG|0644, st_size=185, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/frexp_amd64p32.s", {st_mode=S_IFREG|0644, st_size=186, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/exp2_386.s", {st_mode=S_IFREG|0644, st_size=1156, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/log1p_386.s", {st_mode=S_IFREG|0644, st_size=874, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/export_test.go", {st_mode=S_IFREG|0644, st_size=289, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/all_test.go", {st_mode=S_IFREG|0644, st_size=63984, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/atan_386.s", {st_mode=S_IFREG|0644, st_size=390, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/exp_amd64p32.s", {st_mode=S_IFREG|0644, st_size=184, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/hypot_amd64.s", {st_mode=S_IFREG|0644, st_size=1086, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/logb.go", {st_mode=S_IFREG|0644, st_size=1014, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/floor_amd64.s", {st_mode=S_IFREG|0644, st_size=2037, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/log1p.go", {st_mode=S_IFREG|0644, st_size=6470, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/atan2_amd64.s", {st_mode=S_IFREG|0644, st_size=242, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/sqrt_arm.s", {st_mode=S_IFREG|0644, st_size=315, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/atan2_386.s", {st_mode=S_IFREG|0644, st_size=410, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/pow10.go", {st_mode=S_IFREG|0644, st_size=828, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/cbrt.go", {st_mode=S_IFREG|0644, st_size=1641, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/tan_386.s", {st_mode=S_IFREG|0644, st_size=955, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/abs_arm.s", {st_mode=S_IFREG|0644, st_size=340, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/expm1_arm.s", {st_mode=S_IFREG|0644, st_size=240, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/asin_arm.s", {st_mode=S_IFREG|0644, st_size=280, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/log1p_amd64.s", {st_mode=S_IFREG|0644, st_size=242, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/log1p_amd64p32.s", {st_mode=S_IFREG|0644, st_size=186, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/modf_amd64p32.s", {st_mode=S_IFREG|0644, st_size=185, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/log10.go", {st_mode=S_IFREG|0644, st_size=575, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/nextafter.go", {st_mode=S_IFREG|0644, st_size=682, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/expm1.go", {st_mode=S_IFREG|0644, st_size=8073, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/mod_amd64.s", {st_mode=S_IFREG|0644, st_size=238, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/asin_386.s", {st_mode=S_IFREG|0644, st_size=1129, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/erf.go", {st_mode=S_IFREG|0644, st_size=11598, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/sqrt.go", {st_mode=S_IFREG|0644, st_size=4758, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/modf_386.s", {st_mode=S_IFREG|0644, st_size=775, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/dim_amd64.s", {st_mode=S_IFREG|0644, st_size=2799, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/log10_arm.s", {st_mode=S_IFREG|0644, st_size=282, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/remainder.go", {st_mode=S_IFREG|0644, st_size=1959, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/tan.go", {st_mode=S_IFREG|0644, st_size=3750, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/dim_386.s", {st_mode=S_IFREG|0644, st_size=322, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/dim.go", {st_mode=S_IFREG|0644, st_size=1371, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/sin.go", {st_mode=S_IFREG|0644, st_size=6425, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/exp_arm.s", {st_mode=S_IFREG|0644, st_size=236, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/atan2_amd64p32.s", {st_mode=S_IFREG|0644, st_size=186, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/rand", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/sin_386.s", {st_mode=S_IFREG|0644, st_size=1535, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/abs_386.s", {st_mode=S_IFREG|0644, st_size=344, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/log10_amd64p32.s", {st_mode=S_IFREG|0644, st_size=186, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/sincos_amd64.s", {st_mode=S_IFREG|0644, st_size=4065, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/hypot_amd64p32.s", {st_mode=S_IFREG|0644, st_size=186, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/signbit.go", {st_mode=S_IFREG|0644, st_size=302, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/log10_386.s", {st_mode=S_IFREG|0644, st_size=610, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/remainder_arm.s", {st_mode=S_IFREG|0644, st_size=248, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/math/sqrt_amd64p32.s", {st_mode=S_IFREG|0644, st_size=185, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/abs.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 440
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/abs_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 373
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/acosh.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 1735
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/all_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/asin.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 983
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/asin_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 284
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/asinh.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 1869
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/atan.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 3043
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/atan2.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1485
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/atan2_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 242
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/atan_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 240
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/atanh.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 1943
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/bits.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1835
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/cbrt.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1641
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/const.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2031
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/copysign.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 378
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/dim.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 1371
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/dim_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 2799
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/erf.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 4096
[pid 20018] read(3, "oximate\n//              g(s)=f(1"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/exp.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/exp2_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 240
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/exp_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 2920
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/expm1.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 4096
[pid 20018] read(3, "0+2.0*(r-E);\n//        (v)   if "..., 4096) = 3977
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/expm1_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 242
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/export_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 289
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/floor.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009-2010 The Go Au"..., 4096) = 1084
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/floor_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 2037
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/frexp.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 860
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/frexp_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 242
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/gamma.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/hypot.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 798
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/hypot_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 1086
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/j0.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/j1.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/jn.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/ldexp.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 990
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/ldexp_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 242
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/lgamma.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/log.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 3891
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/log10.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 575
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/log10_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 286
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/log1p.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/log1p_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 242
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/log_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 3714
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/logb.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 1014
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/mod.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009-2010 The Go Au"..., 4096) = 874
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/mod_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 238
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/modf.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 776
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/modf_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 240
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/nextafter.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 682
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/pow.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2656
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/pow10.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 828
[pid 20018] close(3)                    = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 10000} <unfinished ...>
[pid 20018] open("/usr/local/go/src/pkg/math/remainder.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 1959
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/remainder_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 250
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/signbit.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 302
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/sin.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/sin_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 280
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/sincos.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 1862
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/sincos_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 4065
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/sinh.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1535
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/sqrt.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/sqrt_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 302
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/tan.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 3750
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/tan_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 238
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/tanh.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2625
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/math/unsafe.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 859
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/os", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/os", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 64 entries */, 4096) = 2296
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/env_test.go", {st_mode=S_IFREG|0644, st_size=1415, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/exec", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/stat_solaris.go", {st_mode=S_IFREG|0644, st_size=1400, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/pipe_linux.go", {st_mode=S_IFREG|0644, st_size=990, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/path_plan9.go", {st_mode=S_IFREG|0644, st_size=443, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/file_windows.go", {st_mode=S_IFREG|0644, st_size=12475, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/exec_windows.go", {st_mode=S_IFREG|0644, st_size=3299, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/path_unix.go", {st_mode=S_IFREG|0644, st_size=507, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/path_windows.go", {st_mode=S_IFREG|0644, st_size=488, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/str.go", {st_mode=S_IFREG|0644, st_size=496, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/stat_plan9.go", {st_mode=S_IFREG|0644, st_size=2677, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/file.go", {st_mode=S_IFREG|0644, st_size=7493, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/exec_posix.go", {st_mode=S_IFREG|0644, st_size=2931, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/exec.go", {st_mode=S_IFREG|0644, st_size=2162, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/proc.go", {st_mode=S_IFREG|0644, st_size=1182, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/types.go", {st_mode=S_IFREG|0644, st_size=4044, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/stat_dragonfly.go", {st_mode=S_IFREG|0644, st_size=1400, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/getwd_darwin.go", {st_mode=S_IFREG|0644, st_size=317, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/doc.go", {st_mode=S_IFREG|0644, st_size=4715, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/sys_solaris.go", {st_mode=S_IFREG|0644, st_size=265, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/error_plan9.go", {st_mode=S_IFREG|0644, st_size=1035, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/error_windows.go", {st_mode=S_IFREG|0644, st_size=940, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/sys_darwin.go", {st_mode=S_IFREG|0644, st_size=768, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/dir_windows.go", {st_mode=S_IFREG|0644, st_size=376, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/error.go", {st_mode=S_IFREG|0644, st_size=1937, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/sys_windows.go", {st_mode=S_IFREG|0644, st_size=349, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/os_test.go", {st_mode=S_IFREG|0644, st_size=31197, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/env.go", {st_mode=S_IFREG|0644, st_size=2927, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/os_unix_test.go", {st_mode=S_IFREG|0644, st_size=2095, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/error_windows_test.go", {st_mode=S_IFREG|0644, st_size=972, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/exec_unix.go", {st_mode=S_IFREG|0644, st_size=1544, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/dir_unix.go", {st_mode=S_IFREG|0644, st_size=1185, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/signal", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/path_test.go", {st_mode=S_IFREG|0644, st_size=5416, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/types_notwin.go", {st_mode=S_IFREG|0644, st_size=641, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/stat_openbsd.go", {st_mode=S_IFREG|0644, st_size=1400, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/sys_plan9.go", {st_mode=S_IFREG|0644, st_size=473, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/sys_bsd.go", {st_mode=S_IFREG|0644, st_size=512, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/export_test.go", {st_mode=S_IFREG|0644, st_size=235, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/sys_unix.go", {st_mode=S_IFREG|0644, st_size=339, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/exec_plan9.go", {st_mode=S_IFREG|0644, st_size=2927, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/sys_freebsd.go", {st_mode=S_IFREG|0644, st_size=595, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/user", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/sys_nacl.go", {st_mode=S_IFREG|0644, st_size=290, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/file_posix.go", {st_mode=S_IFREG|0644, st_size=4743, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/file_unix.go", {st_mode=S_IFREG|0644, st_size=8675, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/getwd.go", {st_mode=S_IFREG|0644, st_size=2532, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/stat_linux.go", {st_mode=S_IFREG|0644, st_size=1400, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/file_plan9.go", {st_mode=S_IFREG|0644, st_size=12430, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/error_unix.go", {st_mode=S_IFREG|0644, st_size=914, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/env_unix_test.go", {st_mode=S_IFREG|0644, st_size=652, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/dir_plan9.go", {st_mode=S_IFREG|0644, st_size=1564, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/stat_darwin.go", {st_mode=S_IFREG|0644, st_size=1427, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/sys_linux.go", {st_mode=S_IFREG|0644, st_size=517, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/stat_nacl.go", {st_mode=S_IFREG|0644, st_size=1416, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/types_windows.go", {st_mode=S_IFREG|0644, st_size=2315, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/path.go", {st_mode=S_IFREG|0644, st_size=2646, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/pipe_bsd.go", {st_mode=S_IFREG|0644, st_size=796, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/stat_freebsd.go", {st_mode=S_IFREG|0644, st_size=1410, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/error_test.go", {st_mode=S_IFREG|0644, st_size=3159, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/stat_windows.go", {st_mode=S_IFREG|0644, st_size=3966, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/stat_netbsd.go", {st_mode=S_IFREG|0644, st_size=1410, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/dir_unix.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1185
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/doc.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/env.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 2927
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/env_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 1415
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/env_unix_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 652
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/error.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1937
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/error_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 3159
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/error_unix.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 914
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/exec.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2162
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/exec_posix.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2931
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/exec_unix.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1544
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/export_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 235
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/file.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/file_posix.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/file_unix.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/getwd.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2532
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/os_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/os_unix_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2095
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/path.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2646
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/path_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/path_unix.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 507
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/pipe_bsd.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 796
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/pipe_linux.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 990
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/proc.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1182
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/stat_linux.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1400
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/str.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 496
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/sys_bsd.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 512
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/sys_linux.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 517
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/sys_unix.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2014 The Go Authors"..., 4096) = 339
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/types.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4044
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/types_notwin.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 641
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/syscall", {st_mode=S_IFDIR|0755, st_size=12288, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 97 entries */, 4096) = 4088
[pid 20018] getdents64(3, /* 95 entries */, 4096) = 4064
[pid 20018] getdents64(3, /* 20 entries */, 4096) = 872
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_netbsd_386.go", {st_mode=S_IFREG|0644, st_size=6005, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_windows_386.go", {st_mode=S_IFREG|0644, st_size=40, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_windows_386.s", {st_mode=S_IFREG|0644, st_size=249, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/route_netbsd.go", {st_mode=S_IFREG|0644, st_size=1288, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/env_plan9.go", {st_mode=S_IFREG|0644, st_size=2457, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_windows_test.go", {st_mode=S_IFREG|0644, st_size=1673, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_dragonfly_386.go", {st_mode=S_IFREG|0644, st_size=29794, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_netbsd_arm.go", {st_mode=S_IFREG|0644, st_size=6150, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_openbsd_386.s", {st_mode=S_IFREG|0644, st_size=2718, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/netlink_linux.go", {st_mode=S_IFREG|0644, st_size=4756, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_freebsd_386.go", {st_mode=S_IFREG|0644, st_size=68946, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_nacl_amd64p32.go", {st_mode=S_IFREG|0644, st_size=726, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysctl_openbsd.go", {st_mode=S_IFREG|0644, st_size=11989, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/mksysnum_linux.pl", {st_mode=S_IFREG|0755, st_size=636, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_freebsd_386.go", {st_mode=S_IFREG|0644, st_size=7756, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/types_dragonfly.go", {st_mode=S_IFREG|0644, st_size=5022, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_nacl.go", {st_mode=S_IFREG|0644, st_size=7126, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_windows_amd64.go", {st_mode=S_IFREG|0644, st_size=56039, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_freebsd_amd64.go", {st_mode=S_IFREG|0644, st_size=68991, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_nacl_386.go", {st_mode=S_IFREG|0644, st_size=1390, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/route_dragonfly.go", {st_mode=S_IFREG|0644, st_size=2281, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/mkerrors.sh", {st_mode=S_IFREG|0755, st_size=10352, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/types_linux.go", {st_mode=S_IFREG|0644, st_size=11437, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/types_openbsd.go", {st_mode=S_IFREG|0644, st_size=5015, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_freebsd_arm.go", {st_mode=S_IFREG|0644, st_size=30074, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_openbsd_amd64.go", {st_mode=S_IFREG|0644, st_size=29033, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/mksysnum_openbsd.pl", {st_mode=S_IFREG|0755, st_size=873, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_dragonfly_amd64.go", {st_mode=S_IFREG|0644, st_size=6736, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_linux_amd64.go", {st_mode=S_IFREG|0644, st_size=44690, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_openbsd_amd64.go", {st_mode=S_IFREG|0644, st_size=63818, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall.go", {st_mode=S_IFREG|0644, st_size=2835, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_bsd_test.go", {st_mode=S_IFREG|0644, st_size=652, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/env_windows.go", {st_mode=S_IFREG|0644, st_size=1735, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_dragonfly_amd64.go", {st_mode=S_IFREG|0644, st_size=29642, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/exec_windows.go", {st_mode=S_IFREG|0644, st_size=7462, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_netbsd_386.go", {st_mode=S_IFREG|0644, st_size=28324, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/mksysnum_dragonfly.pl", {st_mode=S_IFREG|0755, st_size=868, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/route_bsd.go", {st_mode=S_IFREG|0644, st_size=5687, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/str.go", {st_mode=S_IFREG|0644, st_size=484, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_plan9_amd64.go", {st_mode=S_IFREG|0644, st_size=1430, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_netbsd_amd64.go", {st_mode=S_IFREG|0644, st_size=6189, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_dragonfly_amd64.go", {st_mode=S_IFREG|0644, st_size=23020, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_solaris_amd64.go", {st_mode=S_IFREG|0644, st_size=5559, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/security_windows.go", {st_mode=S_IFREG|0644, st_size=10899, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/exec_bsd.go", {st_mode=S_IFREG|0644, st_size=6205, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/types_solaris.go", {st_mode=S_IFREG|0644, st_size=4573, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_linux_386.go", {st_mode=S_IFREG|0644, st_size=59688, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_plan9.go", {st_mode=S_IFREG|0644, st_size=7101, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_linux_arm.go", {st_mode=S_IFREG|0644, st_size=6516, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_freebsd_amd64.go", {st_mode=S_IFREG|0644, st_size=7839, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/types_darwin.go", {st_mode=S_IFREG|0644, st_size=5053, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/net_nacl.go", {st_mode=S_IFREG|0644, st_size=17238, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/types_netbsd.go", {st_mode=S_IFREG|0644, st_size=4745, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_no_getwd.go", {st_mode=S_IFREG|0644, st_size=304, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_netbsd.go", {st_mode=S_IFREG|0644, st_size=11608, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_openbsd_amd64.go", {st_mode=S_IFREG|0644, st_size=992, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_freebsd_amd64.go", {st_mode=S_IFREG|0644, st_size=26178, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/mksysnum_plan9.sh", {st_mode=S_IFREG|0755, st_size=616, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_freebsd_amd64.go", {st_mode=S_IFREG|0644, st_size=1442, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_netbsd_386.go", {st_mode=S_IFREG|0644, st_size=69031, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_test.go", {st_mode=S_IFREG|0644, st_size=755, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_darwin_amd64.s", {st_mode=S_IFREG|0644, st_size=2209, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/mksyscall_solaris.pl", {st_mode=S_IFREG|0755, st_size=6925, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_netbsd_arm.go", {st_mode=S_IFREG|0644, st_size=68021, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/dll_windows.go", {st_mode=S_IFREG|0644, st_size=8331, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_freebsd_386.go", {st_mode=S_IFREG|0644, st_size=30049, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_plan9_amd64.go", {st_mode=S_IFREG|0644, st_size=6305, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/time_nacl_386.s", {st_mode=S_IFREG|0644, st_size=314, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/mksysctl_openbsd.pl", {st_mode=S_IFREG|0755, st_size=5162, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_openbsd_386.go", {st_mode=S_IFREG|0644, st_size=6774, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_plan9_amd64.go", {st_mode=S_IFREG|0644, st_size=992, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_darwin_amd64.go", {st_mode=S_IFREG|0644, st_size=56579, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_netbsd_386.go", {st_mode=S_IFREG|0644, st_size=26138, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_unix.go", {st_mode=S_IFREG|0644, st_size=7126, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_netbsd_386.go", {st_mode=S_IFREG|0644, st_size=1021, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_netbsd_amd64.go", {st_mode=S_IFREG|0644, st_size=26138, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/fs_nacl.go", {st_mode=S_IFREG|0644, st_size=17054, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/fd_nacl.go", {st_mode=S_IFREG|0644, st_size=7230, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_solaris_amd64.go", {st_mode=S_IFREG|0644, st_size=22085, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/mksysnum_darwin.pl", {st_mode=S_IFREG|0755, st_size=580, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_netbsd_arm.s", {st_mode=S_IFREG|0644, st_size=2730, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/flock.go", {st_mode=S_IFREG|0644, st_size=698, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_bsd.go", {st_mode=S_IFREG|0644, st_size=14747, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/so_solaris.go", {st_mode=S_IFREG|0644, st_size=7005, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_plan9_386.go", {st_mode=S_IFREG|0644, st_size=6269, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_netbsd_amd64.go", {st_mode=S_IFREG|0644, st_size=1021, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_openbsd_386.go", {st_mode=S_IFREG|0644, st_size=29185, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/srpc_nacl.go", {st_mode=S_IFREG|0644, st_size=18301, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_dragonfly_amd64.go", {st_mode=S_IFREG|0644, st_size=60799, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/mkall_windows.bat", {st_mode=S_IFREG|0644, st_size=573, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_linux_arm.go", {st_mode=S_IFREG|0644, st_size=11806, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_openbsd.go", {st_mode=S_IFREG|0644, st_size=7879, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_nacl_amd64p32.go", {st_mode=S_IFREG|0644, st_size=1390, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/route_freebsd_64bit.go", {st_mode=S_IFREG|0644, st_size=423, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/sockcmsg_linux.go", {st_mode=S_IFREG|0644, st_size=1075, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_windows.go", {st_mode=S_IFREG|0644, st_size=33976, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_dragonfly_386.go", {st_mode=S_IFREG|0644, st_size=23020, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_linux_386.go", {st_mode=S_IFREG|0644, st_size=39628, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/exec_unix.go", {st_mode=S_IFREG|0644, st_size=7213, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_freebsd_arm.s", {st_mode=S_IFREG|0644, st_size=2835, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_windows_386.go", {st_mode=S_IFREG|0644, st_size=177, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/route_freebsd.go", {st_mode=S_IFREG|0644, st_size=2367, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_linux_amd64.go", {st_mode=S_IFREG|0644, st_size=12053, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_netbsd_amd64.go", {st_mode=S_IFREG|0644, st_size=68611, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_netbsd_amd64.s", {st_mode=S_IFREG|0644, st_size=2846, ...}) = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_openbsd_386.go",  <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 10000} <unfinished ...>
[pid 20018] <... lstat resumed> {st_mode=S_IFREG|0644, st_size=63864, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_dragonfly_386.s", {st_mode=S_IFREG|0644, st_size=2601, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/mksyscall_windows.go", {st_mode=S_IFREG|0644, st_size=16003, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_freebsd_386.s", {st_mode=S_IFREG|0644, st_size=2718, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_linux_386.s", {st_mode=S_IFREG|0644, st_size=4575, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_openbsd_amd64.go", {st_mode=S_IFREG|0644, st_size=6933, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/bpf_bsd.go", {st_mode=S_IFREG|0644, st_size=3850, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_windows_amd64.s", {st_mode=S_IFREG|0644, st_size=251, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/race0.go", {st_mode=S_IFREG|0644, st_size=443, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/exec_linux.go", {st_mode=S_IFREG|0644, st_size=7005, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_nacl_amd64p32.s", {st_mode=S_IFREG|0644, st_size=892, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_dragonfly_amd64.go", {st_mode=S_IFREG|0644, st_size=1442, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_windows_amd64.go", {st_mode=S_IFREG|0644, st_size=478, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_dragonfly.go", {st_mode=S_IFREG|0644, st_size=10467, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_windows_386.go", {st_mode=S_IFREG|0644, st_size=478, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_linux_386.go", {st_mode=S_IFREG|0644, st_size=11681, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_openbsd_386.go", {st_mode=S_IFREG|0644, st_size=1020, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/lsf_linux.go", {st_mode=S_IFREG|0644, st_size=1789, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/tables_nacl.go", {st_mode=S_IFREG|0644, st_size=15004, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_solaris_amd64.go", {st_mode=S_IFREG|0644, st_size=264, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_linux_amd64.s", {st_mode=S_IFREG|0644, st_size=2547, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_solaris_amd64.s", {st_mode=S_IFREG|0644, st_size=251, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_windows.go", {st_mode=S_IFREG|0644, st_size=10208, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_linux_arm.s", {st_mode=S_IFREG|0644, st_size=3451, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_windows_386.go", {st_mode=S_IFREG|0644, st_size=177, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/route_freebsd_32bit.go", {st_mode=S_IFREG|0644, st_size=869, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/route_openbsd.go", {st_mode=S_IFREG|0644, st_size=1289, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_windows_amd64.go", {st_mode=S_IFREG|0644, st_size=177, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_netbsd_arm.go", {st_mode=S_IFREG|0644, st_size=1021, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/mkall.sh", {st_mode=S_IFREG|0755, st_size=8459, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/sockcmsg_unix.go", {st_mode=S_IFREG|0644, st_size=2962, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/exec_plan9.go", {st_mode=S_IFREG|0644, st_size=15748, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_freebsd_arm.go", {st_mode=S_IFREG|0644, st_size=1463, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_plan9_amd64.go", {st_mode=S_IFREG|0644, st_size=1298, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_solaris_amd64.go", {st_mode=S_IFREG|0644, st_size=948, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/unzip_nacl.go", {st_mode=S_IFREG|0644, st_size=15156, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/mksysnum_netbsd.pl", {st_mode=S_IFREG|0755, st_size=1037, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/mmap_unix_test.go", {st_mode=S_IFREG|0644, st_size=540, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_openbsd_386.go", {st_mode=S_IFREG|0644, st_size=14323, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_freebsd_arm.go", {st_mode=S_IFREG|0644, st_size=7866, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_freebsd_arm.go", {st_mode=S_IFREG|0644, st_size=68936, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/env_unix.go", {st_mode=S_IFREG|0644, st_size=2148, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_linux_amd64.go", {st_mode=S_IFREG|0644, st_size=10325, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_linux_386.go", {st_mode=S_IFREG|0644, st_size=11515, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_dragonfly_386.go", {st_mode=S_IFREG|0644, st_size=6607, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_darwin_amd64.go", {st_mode=S_IFREG|0644, st_size=7029, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_netbsd_386.s", {st_mode=S_IFREG|0644, st_size=2717, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_plan9_386.go", {st_mode=S_IFREG|0644, st_size=655, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_solaris_amd64.go", {st_mode=S_IFREG|0644, st_size=51909, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/race.go", {st_mode=S_IFREG|0644, st_size=580, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/flock_linux_32bit.go", {st_mode=S_IFREG|0644, st_size=367, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_freebsd_386.go", {st_mode=S_IFREG|0644, st_size=26178, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_windows.go", {st_mode=S_IFREG|0644, st_size=25032, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_linux_arm.go", {st_mode=S_IFREG|0644, st_size=43820, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/types_plan9.c", {st_mode=S_IFREG|0644, st_size=3335, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_darwin_386.go", {st_mode=S_IFREG|0644, st_size=31839, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_plan9_amd64.go", {st_mode=S_IFREG|0644, st_size=391, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_darwin_amd64.go", {st_mode=S_IFREG|0644, st_size=14875, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_darwin_386.s", {st_mode=S_IFREG|0644, st_size=2717, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/route_darwin.go", {st_mode=S_IFREG|0644, st_size=1880, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_freebsd.go", {st_mode=S_IFREG|0644, st_size=10979, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_linux_amd64.go", {st_mode=S_IFREG|0644, st_size=59726, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_windows_amd64.go", {st_mode=S_IFREG|0644, st_size=177, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/time_nacl_amd64p32.s", {st_mode=S_IFREG|0644, st_size=314, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_darwin_amd64.go", {st_mode=S_IFREG|0644, st_size=1751, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/types_freebsd.go", {st_mode=S_IFREG|0644, st_size=7565, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_linux_amd64.go", {st_mode=S_IFREG|0644, st_size=4521, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_freebsd_arm.go", {st_mode=S_IFREG|0644, st_size=26178, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_plan9_386.go", {st_mode=S_IFREG|0644, st_size=971, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_darwin_amd64.go", {st_mode=S_IFREG|0644, st_size=31688, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_windows_amd64.go", {st_mode=S_IFREG|0644, st_size=40, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_unix_test.go", {st_mode=S_IFREG|0644, st_size=8142, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_netbsd_amd64.go", {st_mode=S_IFREG|0644, st_size=28172, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/creds_test.go", {st_mode=S_IFREG|0644, st_size=2925, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_plan9_386.go", {st_mode=S_IFREG|0644, st_size=1430, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/mksyscall.pl", {st_mode=S_IFREG|0755, st_size=7867, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_linux_arm.go", {st_mode=S_IFREG|0644, st_size=60225, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_linux.go", {st_mode=S_IFREG|0644, st_size=26417, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/dir_plan9.go", {st_mode=S_IFREG|0644, st_size=5740, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_darwin_386.go", {st_mode=S_IFREG|0644, st_size=1807, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_dragonfly_amd64.s", {st_mode=S_IFREG|0644, st_size=2777, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_nacl_386.s", {st_mode=S_IFREG|0644, st_size=904, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_dragonfly_386.go", {st_mode=S_IFREG|0644, st_size=60799, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_plan9_386.s", {st_mode=S_IFREG|0644, st_size=2986, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_darwin.go", {st_mode=S_IFREG|0644, st_size=12983, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_solaris.go", {st_mode=S_IFREG|0644, st_size=14946, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_openbsd_amd64.s", {st_mode=S_IFREG|0644, st_size=2847, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/mksysnum_freebsd.pl", {st_mode=S_IFREG|0755, st_size=1382, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_darwin_386.go", {st_mode=S_IFREG|0644, st_size=6813, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_freebsd_amd64.go", {st_mode=S_IFREG|0644, st_size=29898, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_freebsd_386.go", {st_mode=S_IFREG|0644, st_size=1491, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_linux_arm.go", {st_mode=S_IFREG|0644, st_size=11568, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_netbsd_arm.go", {st_mode=S_IFREG|0644, st_size=26138, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_netbsd_arm.go", {st_mode=S_IFREG|0644, st_size=28321, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/ztypes_plan9_386.go", {st_mode=S_IFREG|0644, st_size=1298, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_linux_386.go", {st_mode=S_IFREG|0644, st_size=10239, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_freebsd_amd64.s", {st_mode=S_IFREG|0644, st_size=3090, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_darwin_386.go", {st_mode=S_IFREG|0644, st_size=14875, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsyscall_windows_386.go", {st_mode=S_IFREG|0644, st_size=56037, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_dragonfly_386.go", {st_mode=S_IFREG|0644, st_size=1491, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/asm_plan9_amd64.s", {st_mode=S_IFREG|0644, st_size=3169, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/syscall_nacl_386.go", {st_mode=S_IFREG|0644, st_size=726, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/exec_solaris.go", {st_mode=S_IFREG|0644, st_size=6144, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zerrors_darwin_386.go", {st_mode=S_IFREG|0644, st_size=56579, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/syscall/zsysnum_openbsd_amd64.go", {st_mode=S_IFREG|0644, st_size=14323, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/asm_linux_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2547
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/bpf_bsd.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 3850
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/creds_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 2925
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/env_unix.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 2148
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/exec_bsd.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/exec_linux.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/exec_unix.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/flock.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// +build linux darwin freebsd o"..., 4096) = 698
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/flock_linux_32bit.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// +build linux,386 linux,arm\n\n/"..., 4096) = 367
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/lsf_linux.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 1789
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/mmap_unix_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2014 The Go Authors"..., 4096) = 540
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/netlink_linux.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/race.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 580
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/race0.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 443
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/route_bsd.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/route_freebsd_32bit.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2014 The Go Authors"..., 4096) = 869
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/route_freebsd_64bit.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2014 The Go Authors"..., 4096) = 423
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/sockcmsg_linux.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 1075
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/sockcmsg_unix.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 2962
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/str.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 484
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/syscall.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2835
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/syscall_bsd.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/syscall_bsd_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2014 The Go Authors"..., 4096) = 652
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/syscall_linux.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/syscall_linux_amd64.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/syscall_no_getwd.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 304
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/syscall_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 755
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/syscall_unix.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/syscall_unix_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/types_linux.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/zerrors_linux_amd64.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// mkerrors.sh -m64\n// MACHINE G"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/zsyscall_linux_amd64.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// mksyscall.pl syscall_linux.go"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/zsysnum_linux_amd64.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// mksysnum_linux.pl /usr/includ"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/syscall/ztypes_linux_amd64.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Created by cgo -godefs - DO N"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/time", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/time", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 27 entries */, 4096) = 984
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/zoneinfo_windows_test.go", {st_mode=S_IFREG|0644, st_size=821, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/zoneinfo.go", {st_mode=S_IFREG|0644, st_size=7800, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/Makefile", {st_mode=S_IFREG|0644, st_size=272, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/format.go", {st_mode=S_IFREG|0644, st_size=35278, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/time_test.go", {st_mode=S_IFREG|0644, st_size=31151, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/sys_windows.go", {st_mode=S_IFREG|0644, st_size=1424, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/sleep_test.go", {st_mode=S_IFREG|0644, st_size=8296, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/zoneinfo_read.go", {st_mode=S_IFREG|0644, st_size=7853, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/zoneinfo_plan9.go", {st_mode=S_IFREG|0644, st_size=3222, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/tick_test.go", {st_mode=S_IFREG|0644, st_size=1698, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/sys_plan9.go", {st_mode=S_IFREG|0644, st_size=1451, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/zoneinfo_unix.go", {st_mode=S_IFREG|0644, st_size=2031, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/export_test.go", {st_mode=S_IFREG|0644, st_size=466, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/sys_unix.go", {st_mode=S_IFREG|0644, st_size=1510, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/example_test.go", {st_mode=S_IFREG|0644, st_size=3898, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/zoneinfo_abbrs_windows.go", {st_mode=S_IFREG|0644, st_size=7647, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/tick.go", {st_mode=S_IFREG|0644, st_size=1614, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/format_test.go", {st_mode=S_IFREG|0644, st_size=18049, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/internal_test.go", {st_mode=S_IFREG|0644, st_size=2704, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/sleep.go", {st_mode=S_IFREG|0644, st_size=3249, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/time.go", {st_mode=S_IFREG|0644, st_size=34218, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/zoneinfo_test.go", {st_mode=S_IFREG|0644, st_size=1523, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/genzabbrs.go", {st_mode=S_IFREG|0644, st_size=2830, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/zoneinfo_windows.go", {st_mode=S_IFREG|0644, st_size=7893, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/time/export_windows_test.go", {st_mode=S_IFREG|0644, st_size=263, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/time/example_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 3898
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/time/export_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 466
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/time/format.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/time/format_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/time/genzabbrs.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 2830
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/time/internal_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 2704
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/time/sleep.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 3249
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/time/sleep_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/time/sys_unix.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 1510
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/time/tick.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1614
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/time/tick_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1698
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/time/time.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/time/time_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/time/zoneinfo.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/time/zoneinfo_read.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/time/zoneinfo_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2014 The Go Authors"..., 4096) = 1523
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/time/zoneinfo_unix.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2031
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/reflect", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/reflect", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 15 entries */, 4096) = 496
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/reflect/type.go", {st_mode=S_IFREG|0644, st_size=52801, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/reflect/set_test.go", {st_mode=S_IFREG|0644, st_size=5550, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/reflect/asm_amd64p32.s", {st_mode=S_IFREG|0644, st_size=917, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/reflect/asm_amd64.s", {st_mode=S_IFREG|0644, st_size=919, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/reflect/export_test.go", {st_mode=S_IFREG|0644, st_size=447, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/reflect/all_test.go", {st_mode=S_IFREG|0644, st_size=94237, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/reflect/example_test.go", {st_mode=S_IFREG|0644, st_size=1785, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/reflect/asm_arm.s", {st_mode=S_IFREG|0644, st_size=912, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/reflect/deepequal.go", {st_mode=S_IFREG|0644, st_size=3567, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/reflect/value.go", {st_mode=S_IFREG|0644, st_size=78790, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/reflect/asm_386.s", {st_mode=S_IFREG|0644, st_size=917, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/reflect/makefunc.go", {st_mode=S_IFREG|0644, st_size=4312, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/reflect/tostring_test.go", {st_mode=S_IFREG|0644, st_size=2191, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/reflect/all_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3,  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... read resumed> "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20021] select(0, NULL, NULL, NULL, {0, 10000} <unfinished ...>
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/reflect/asm_amd64.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 919
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/reflect/deepequal.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 3567
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/reflect/example_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 1785
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/reflect/export_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 447
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/reflect/makefunc.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/reflect/set_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/reflect/tostring_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2191
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/reflect/type.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/reflect/value.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/strconv", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 24 entries */, 4096) = 808
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/atob.go", {st_mode=S_IFREG|0644, st_size=983, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/decimal.go", {st_mode=S_IFREG|0644, st_size=8064, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/isprint.go", {st_mode=S_IFREG|0644, st_size=8499, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/quote.go", {st_mode=S_IFREG|0644, st_size=11325, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/decimal_test.go", {st_mode=S_IFREG|0644, st_size=3087, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/atof_test.go", {st_mode=S_IFREG|0644, st_size=12268, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/testdata", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/makeisprint.go", {st_mode=S_IFREG|0644, st_size=3968, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/quote_test.go", {st_mode=S_IFREG|0644, st_size=5941, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/itoa.go", {st_mode=S_IFREG|0644, st_size=3290, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/ftoa.go", {st_mode=S_IFREG|0644, st_size=11360, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/strconv_test.go", {st_mode=S_IFREG|0644, st_size=1768, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/atoi.go", {st_mode=S_IFREG|0644, st_size=4480, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/ftoa_test.go", {st_mode=S_IFREG|0644, st_size=7338, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/internal_test.go", {st_mode=S_IFREG|0644, st_size=385, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/quote_example_test.go", {st_mode=S_IFREG|0644, st_size=755, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/itoa_test.go", {st_mode=S_IFREG|0644, st_size=3924, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/fp_test.go", {st_mode=S_IFREG|0644, st_size=2984, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/atoi_test.go", {st_mode=S_IFREG|0644, st_size=7982, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/atob_test.go", {st_mode=S_IFREG|0644, st_size=1931, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/atof.go", {st_mode=S_IFREG|0644, st_size=11336, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strconv/extfloat.go", {st_mode=S_IFREG|0644, st_size=20253, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/atob.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 983
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/atob_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1931
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/atof.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/atof_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/atoi.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/atoi_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/decimal.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/decimal_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 3087
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/extfloat.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/fp_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 2984
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/ftoa.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/ftoa_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/internal_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 385
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/isprint.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/itoa.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 3290
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/itoa_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 3924
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/makeisprint.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 3968
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/quote.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/quote_example_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 755
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/quote_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strconv/strconv_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 1768
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/unicode/utf8", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/unicode/utf8", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 5 entries */, 4096) = 152
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/unicode/utf8/utf8_test.go", {st_mode=S_IFREG|0644, st_size=10928, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/unicode/utf8/example_test.go", {st_mode=S_IFREG|0644, st_size=3060, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/unicode/utf8/utf8.go", {st_mode=S_IFREG|0644, st_size=10492, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/unicode/utf8/example_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 3060
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/unicode/utf8/utf8.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/unicode/utf8/utf8_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/os/exec", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/os/exec", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 11 entries */, 4096) = 368
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/exec/exec.go", {st_mode=S_IFREG|0644, st_size=12910, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/exec/lp_plan9.go", {st_mode=S_IFREG|0644, st_size=1391, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/exec/lp_unix.go", {st_mode=S_IFREG|0644, st_size=1645, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/exec/lp_windows.go", {st_mode=S_IFREG|0644, st_size=2764, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/exec/exec_test.go", {st_mode=S_IFREG|0644, st_size=19246, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/exec/lp_windows_test.go", {st_mode=S_IFREG|0644, st_size=15169, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/exec/example_test.go", {st_mode=S_IFREG|0644, st_size=1560, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/exec/lp_unix_test.go", {st_mode=S_IFREG|0644, st_size=1163, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/os/exec/lp_test.go", {st_mode=S_IFREG|0644, st_size=737, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/exec/example_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 1560
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/exec/exec.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/exec/exec_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/exec/lp_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 737
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/exec/lp_unix.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 1645
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/os/exec/lp_unix_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 1163
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/bytes", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/bytes", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 14 entries */, 4096) = 496
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/bytes/reader_test.go", {st_mode=S_IFREG|0644, st_size=6146, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/bytes/bytes_test.go", {st_mode=S_IFREG|0644, st_size=31468, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/bytes/bytes.go", {st_mode=S_IFREG|0644, st_size=17977, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/bytes/bytes_decl.go", {st_mode=S_IFREG|0644, st_size=875, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/bytes/export_test.go", {st_mode=S_IFREG|0644, st_size=329, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/bytes/buffer_test.go", {st_mode=S_IFREG|0644, st_size=13454, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/bytes/example_test.go", {st_mode=S_IFREG|0644, st_size=1983, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/bytes/compare_test.go", {st_mode=S_IFREG|0644, st_size=4477, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/bytes/equal_test.go", {st_mode=S_IFREG|0644, st_size=1334, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/bytes/buffer.go", {st_mode=S_IFREG|0644, st_size=13141, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/bytes/reader.go", {st_mode=S_IFREG|0644, st_size=3039, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/bytes/bytes.s", {st_mode=S_IFREG|0644, st_size=213, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/bytes/buffer.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/bytes/buffer_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/bytes/bytes.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/bytes/bytes.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 213
[pid 20018] read(3, "", 4096)           = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/bytes/bytes_decl.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 875
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/bytes/bytes_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/bytes/compare_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/bytes/equal_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 1334
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/bytes/example_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 1983
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/bytes/export_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 329
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/bytes/reader.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 3039
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/bytes/reader_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/unicode", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/unicode", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 15 entries */, 4096) = 504
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/unicode/letter.go", {st_mode=S_IFREG|0644, st_size=9553, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/unicode/letter_test.go", {st_mode=S_IFREG|0644, st_size=12084, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/unicode/utf8", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/unicode/Makefile", {st_mode=S_IFREG|0644, st_size=461, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/unicode/digit_test.go", {st_mode=S_IFREG|0644, st_size=1576, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/unicode/casetables.go", {st_mode=S_IFREG|0644, st_size=755, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/unicode/script_test.go", {st_mode=S_IFREG|0644, st_size=6104, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/unicode/maketables.go", {st_mode=S_IFREG|0644, st_size=33761, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/unicode/utf16", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/unicode/tables.go", {st_mode=S_IFREG|0644, st_size=159224, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/unicode/graphic_test.go", {st_mode=S_IFREG|0644, st_size=2616, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/unicode/digit.go", {st_mode=S_IFREG|0644, st_size=352, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/unicode/graphic.go", {st_mode=S_IFREG|0644, st_size=4458, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/unicode/casetables.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 755
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/unicode/digit.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 352
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/unicode/digit_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 1576
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/unicode/graphic.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/unicode/graphic_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 2616
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/unicode/letter.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/unicode/letter_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/unicode/maketables.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/unicode/script_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/unicode/tables.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/path/filepath", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/path/filepath", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 14 entries */, 4096) = 488
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/path/filepath/symlink_windows.go", {st_mode=S_IFREG|0644, st_size=1675, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/path/filepath/path_plan9.go", {st_mode=S_IFREG|0644, st_size=780, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/path/filepath/example_unix_test.go", {st_mode=S_IFREG|0644, st_size=737, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/path/filepath/path_unix.go", {st_mode=S_IFREG|0644, st_size=818, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/path/filepath/path_windows.go", {st_mode=S_IFREG|0644, st_size=2236, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/path/filepath/path_windows_test.go", {st_mode=S_IFREG|0644, st_size=2223, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/path/filepath/match.go", {st_mode=S_IFREG|0644, st_size=6988, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/path/filepath/symlink.go", {st_mode=S_IFREG|0644, st_size=1342, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/path/filepath/match_test.go", {st_mode=S_IFREG|0644, st_size=4935, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/path/filepath/path_test.go", {st_mode=S_IFREG|0644, st_size=23682, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/path/filepath/export_test.go", {st_mode=S_IFREG|0644, st_size=198, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/path/filepath/path.go", {st_mode=S_IFREG|0644, st_size=13261, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/path/filepath/example_unix_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 737
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/path/filepath/export_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 198
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/path/filepath/match.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/path/filepath/match_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/path/filepath/path.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/path/filepath/path_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/path/filepath/path_unix.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 818
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/path/filepath/symlink.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 1342
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/sort", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/sort", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3,  <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20018] <... getdents64 resumed> /* 12 entries */, 4096) = 448
[pid 20021] select(0, NULL, NULL, NULL, {0, 10000} <unfinished ...>
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sort/search.go", {st_mode=S_IFREG|0644, st_size=4236, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sort/sort_test.go", {st_mode=S_IFREG|0644, st_size=12136, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sort/example_wrapper_test.go", {st_mode=S_IFREG|0644, st_size=1672, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sort/example_keys_test.go", {st_mode=S_IFREG|0644, st_size=2745, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sort/example_interface_test.go", {st_mode=S_IFREG|0644, st_size=896, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sort/search_test.go", {st_mode=S_IFREG|0644, st_size=4349, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sort/export_test.go", {st_mode=S_IFREG|0644, st_size=239, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sort/sort.go", {st_mode=S_IFREG|0644, st_size=13398, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sort/example_test.go", {st_mode=S_IFREG|0644, st_size=481, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/sort/example_multi_test.go", {st_mode=S_IFREG|0644, st_size=4103, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sort/example_interface_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 896
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sort/example_keys_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 2745
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sort/example_multi_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sort/example_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 481
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sort/example_wrapper_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 1672
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sort/export_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 239
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sort/search.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sort/search_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2010 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sort/sort.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/sort/sort_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/src/pkg/strings", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] open("/usr/local/go/src/pkg/strings", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] getdents64(3, /* 14 entries */, 4096) = 488
[pid 20018] getdents64(3, /* 0 entries */, 4096) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strings/reader_test.go", {st_mode=S_IFREG|0644, st_size=3885, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strings/replace_test.go", {st_mode=S_IFREG|0644, st_size=12617, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strings/search.go", {st_mode=S_IFREG|0644, st_size=4289, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strings/strings.s", {st_mode=S_IFREG|0644, st_size=213, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strings/strings_decl.go", {st_mode=S_IFREG|0644, st_size=339, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strings/search_test.go", {st_mode=S_IFREG|0644, st_size=1924, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strings/export_test.go", {st_mode=S_IFREG|0644, st_size=1056, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strings/example_test.go", {st_mode=S_IFREG|0644, st_size=5027, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strings/replace.go", {st_mode=S_IFREG|0644, st_size=13761, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strings/strings_test.go", {st_mode=S_IFREG|0644, st_size=28633, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strings/reader.go", {st_mode=S_IFREG|0644, st_size=3087, ...}) = 0
[pid 20018] lstat("/usr/local/go/src/pkg/strings/strings.go", {st_mode=S_IFREG|0644, st_size=18555, ...}) = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strings/example_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strings/export_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 1056
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strings/reader.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 3087
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strings/reader_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 3885
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strings/replace.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2011 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strings/replace_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strings/search.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strings/search_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2012 The Go Authors"..., 4096) = 1924
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strings/strings.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strings/strings.s", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 213
[pid 20018] read(3, "", 4096)           = 0
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strings/strings_decl.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2013 The Go Authors"..., 4096) = 339
[pid 20018] close(3)                    = 0
[pid 20018] open("/usr/local/go/src/pkg/strings/strings_test.go", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] read(3, "// Copyright 2009 The Go Authors"..., 4096) = 4096
[pid 20018] close(3)                    = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/errors.a", {st_mode=S_IFREG|0644, st_size=5436, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync/atomic.a", {st_mode=S_IFREG|0644, st_size=7762, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync.a", {st_mode=S_IFREG|0644, st_size=82966, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync/atomic.a", {st_mode=S_IFREG|0644, st_size=7762, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/io.a", {st_mode=S_IFREG|0644, st_size=147616, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/errors.a", {st_mode=S_IFREG|0644, st_size=5436, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync.a", {st_mode=S_IFREG|0644, st_size=82966, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync/atomic.a", {st_mode=S_IFREG|0644, st_size=7762, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/math.a", {st_mode=S_IFREG|0644, st_size=186002, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/syscall.a", {st_mode=S_IFREG|0644, st_size=714740, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync.a", {st_mode=S_IFREG|0644, st_size=82966, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync/atomic.a", {st_mode=S_IFREG|0644, st_size=7762, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/time.a", {st_mode=S_IFREG|0644, st_size=464166, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/errors.a", {st_mode=S_IFREG|0644, st_size=5436, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync.a", {st_mode=S_IFREG|0644, st_size=82966, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync/atomic.a", {st_mode=S_IFREG|0644, st_size=7762, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/syscall.a", {st_mode=S_IFREG|0644, st_size=714740, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/os.a", {st_mode=S_IFREG|0644, st_size=313512, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/errors.a", {st_mode=S_IFREG|0644, st_size=5436, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/io.a", {st_mode=S_IFREG|0644, st_size=147616, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync.a", {st_mode=S_IFREG|0644, st_size=82966, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync/atomic.a", {st_mode=S_IFREG|0644, st_size=7762, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/syscall.a", {st_mode=S_IFREG|0644, st_size=714740, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/time.a", {st_mode=S_IFREG|0644, st_size=464166, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/unicode/utf8.a", {st_mode=S_IFREG|0644, st_size=17634, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/strconv.a", {st_mode=S_IFREG|0644, st_size=251324, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/errors.a", {st_mode=S_IFREG|0644, st_size=5436, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/math.a", {st_mode=S_IFREG|0644, st_size=186002, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/unicode/utf8.a", {st_mode=S_IFREG|0644, st_size=17634, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/reflect.a", {st_mode=S_IFREG|0644, st_size=1132820, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/errors.a", {st_mode=S_IFREG|0644, st_size=5436, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/math.a", {st_mode=S_IFREG|0644, st_size=186002, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/strconv.a", {st_mode=S_IFREG|0644, st_size=251324, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync.a", {st_mode=S_IFREG|0644, st_size=82966, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync/atomic.a", {st_mode=S_IFREG|0644, st_size=7762, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/unicode/utf8.a", {st_mode=S_IFREG|0644, st_size=17634, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/fmt.a", {st_mode=S_IFREG|0644, st_size=494232, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/errors.a", {st_mode=S_IFREG|0644, st_size=5436, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/io.a", {st_mode=S_IFREG|0644, st_size=147616, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/math.a", {st_mode=S_IFREG|0644, st_size=186002, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/os.a", {st_mode=S_IFREG|0644, st_size=313512, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/reflect.a", {st_mode=S_IFREG|0644, st_size=1132820, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/strconv.a", {st_mode=S_IFREG|0644, st_size=251324, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync.a", {st_mode=S_IFREG|0644, st_size=82966, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync/atomic.a", {st_mode=S_IFREG|0644, st_size=7762, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/syscall.a", {st_mode=S_IFREG|0644, st_size=714740, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/time.a", {st_mode=S_IFREG|0644, st_size=464166, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/unicode/utf8.a", {st_mode=S_IFREG|0644, st_size=17634, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/unicode.a", {st_mode=S_IFREG|0644, st_size=262606, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/bytes.a", {st_mode=S_IFREG|0644, st_size=173326, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/errors.a", {st_mode=S_IFREG|0644, st_size=5436, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/io.a", {st_mode=S_IFREG|0644, st_size=147616, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync.a", {st_mode=S_IFREG|0644, st_size=82966, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync/atomic.a", {st_mode=S_IFREG|0644, st_size=7762, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/unicode.a", {st_mode=S_IFREG|0644, st_size=262606, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/unicode/utf8.a", {st_mode=S_IFREG|0644, st_size=17634, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sort.a", {st_mode=S_IFREG|0644, st_size=113248, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/strings.a", {st_mode=S_IFREG|0644, st_size=213592, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/errors.a", {st_mode=S_IFREG|0644, st_size=5436, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/io.a", {st_mode=S_IFREG|0644, st_size=147616, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync.a", {st_mode=S_IFREG|0644, st_size=82966, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync/atomic.a", {st_mode=S_IFREG|0644, st_size=7762, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/unicode.a", {st_mode=S_IFREG|0644, st_size=262606, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/unicode/utf8.a", {st_mode=S_IFREG|0644, st_size=17634, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/path/filepath.a", {st_mode=S_IFREG|0644, st_size=90706, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/bytes.a", {st_mode=S_IFREG|0644, st_size=173326, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/errors.a", {st_mode=S_IFREG|0644, st_size=5436, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/io.a", {st_mode=S_IFREG|0644, st_size=147616, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/os.a", {st_mode=S_IFREG|0644, st_size=313512, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sort.a", {st_mode=S_IFREG|0644, st_size=113248, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/strings.a", {st_mode=S_IFREG|0644, st_size=213592, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync.a", {st_mode=S_IFREG|0644, st_size=82966, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync/atomic.a", {st_mode=S_IFREG|0644, st_size=7762, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/syscall.a", {st_mode=S_IFREG|0644, st_size=714740, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/time.a", {st_mode=S_IFREG|0644, st_size=464166, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/unicode.a", {st_mode=S_IFREG|0644, st_size=262606, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/unicode/utf8.a", {st_mode=S_IFREG|0644, st_size=17634, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/os/exec.a", {st_mode=S_IFREG|0644, st_size=250782, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/bytes.a", {st_mode=S_IFREG|0644, st_size=173326, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/errors.a", {st_mode=S_IFREG|0644, st_size=5436, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/io.a", {st_mode=S_IFREG|0644, st_size=147616, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/math.a", {st_mode=S_IFREG|0644, st_size=186002, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/os.a", {st_mode=S_IFREG|0644, st_size=313512, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/path/filepath.a", {st_mode=S_IFREG|0644, st_size=90706, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/runtime.a", {st_mode=S_IFREG|0644, st_size=992958, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sort.a", {st_mode=S_IFREG|0644, st_size=113248, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/strconv.a", {st_mode=S_IFREG|0644, st_size=251324, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/strings.a", {st_mode=S_IFREG|0644, st_size=213592, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync.a", {st_mode=S_IFREG|0644, st_size=82966, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/sync/atomic.a", {st_mode=S_IFREG|0644, st_size=7762, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/syscall.a", {st_mode=S_IFREG|0644, st_size=714740, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/time.a", {st_mode=S_IFREG|0644, st_size=464166, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/unicode.a", {st_mode=S_IFREG|0644, st_size=262606, ...}) = 0
[pid 20018] stat("/usr/local/go/pkg/linux_amd64/unicode/utf8.a", {st_mode=S_IFREG|0644, st_size=17634, ...}) = 0
[pid 20018] stat("/tmp/go-build628409689/command-line-arguments/_obj/", 0xc20810b0e0) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/tmp/go-build628409689/command-line-arguments", 0xc20810b170) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/tmp/go-build628409689", {st_mode=S_IFDIR|0700, st_size=4096, ...}) = 0
[pid 20018] mkdir("/tmp/go-build628409689/command-line-arguments", 0777) = 0
[pid 20018] mkdir("/tmp/go-build628409689/command-line-arguments/_obj/", 0777) = 0
[pid 20018] stat("/tmp/go-build628409689/command-line-arguments/_obj/exe/", 0xc20810b290) = -1 ENOENT (No such file or directory)
[pid 20018] stat("/tmp/go-build628409689/command-line-arguments/_obj", {st_mode=S_IFDIR|0775, st_size=4096, ...}) = 0
[pid 20018] mkdir("/tmp/go-build628409689/command-line-arguments/_obj/exe/", 0777) = 0
[pid 20018] stat("/usr/local/go/pkg/tool/linux_amd64/6g", {st_mode=S_IFREG|0755, st_size=2236110, ...}) = 0
[pid 20018] open("/dev/null", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] pipe2([4, 5], O_CLOEXEC)    = 0
[pid 20018] stat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] pipe2([6, 7], O_CLOEXEC)    = 0
[pid 20018] clone(Process 20027 attached
child_stack=0, flags=|SIGCHLD) = 20027
[pid 20027] chdir("/home/qboxtest/Desktop/qiniu" <unfinished ...>
[pid 20018] close(7 <unfinished ...>
[pid 20027] <... chdir resumed> )       = 0
[pid 20018] <... close resumed> )       = 0
[pid 20027] dup2(3, 0 <unfinished ...>
[pid 20018] read(6,  <unfinished ...>
[pid 20027] <... dup2 resumed> )        = 0
[pid 20027] dup2(5, 1)                  = 1
[pid 20027] dup2(5, 2)                  = 2
[pid 20027] execve("/usr/local/go/pkg/tool/linux_amd64/6g", ["/usr/local/go/pkg/tool/linux_amd"..., "-o", "/tmp/go-build628409689/command-l"..., "-trimpath", "/tmp/go-build628409689", "-p", "command-line-arguments", "-complete", "-D", "_/home/qboxtest/Desktop/qiniu", "-I", "/tmp/go-build628409689", "-pack", "/home/qboxtest/Desktop/qiniu/chi"...], [/* 60 vars */]) = 0
[pid 20018] <... read resumed> "", 8)   = 0
[pid 20018] close(6 <unfinished ...>
[pid 20027] brk(0 <unfinished ...>
[pid 20018] <... close resumed> )       = 0
[pid 20027] <... brk resumed> )         = 0x1dfb000
[pid 20018] close(3 <unfinished ...>
[pid 20027] access("/etc/ld.so.nohwcap", F_OK <unfinished ...>
[pid 20018] <... close resumed> )       = 0
[pid 20027] <... access resumed> )      = -1 ENOENT (No such file or directory)
[pid 20018] close(5 <unfinished ...>
[pid 20027] mmap(NULL, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0 <unfinished ...>
[pid 20018] <... close resumed> )       = 0
[pid 20027] <... mmap resumed> )        = 0x7f0add7eb000
[pid 20018] wait4(20027, Process 20018 suspended
 <unfinished ...>
[pid 20027] access("/etc/ld.so.preload", R_OK) = -1 ENOENT (No such file or directory)
[pid 20027] open("/usr/local/lib/tls/x86_64/libm.so.6", O_RDONLY|O_CLOEXEC) = -1 ENOENT (No such file or directory)
[pid 20027] stat("/usr/local/lib/tls/x86_64", 0x7fff24b35ab0) = -1 ENOENT (No such file or directory)
[pid 20027] open("/usr/local/lib/tls/libm.so.6", O_RDONLY|O_CLOEXEC) = -1 ENOENT (No such file or directory)
[pid 20027] stat("/usr/local/lib/tls", 0x7fff24b35ab0) = -1 ENOENT (No such file or directory)
[pid 20027] open("/usr/local/lib/x86_64/libm.so.6", O_RDONLY|O_CLOEXEC) = -1 ENOENT (No such file or directory)
[pid 20027] stat("/usr/local/lib/x86_64", 0x7fff24b35ab0) = -1 ENOENT (No such file or directory)
[pid 20027] open("/usr/local/lib/libm.so.6", O_RDONLY|O_CLOEXEC) = -1 ENOENT (No such file or directory)
[pid 20027] stat("/usr/local/lib", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20027] open("/etc/ld.so.cache", O_RDONLY|O_CLOEXEC) = 3
[pid 20027] fstat(3, {st_mode=S_IFREG|0644, st_size=124566, ...}) = 0
[pid 20027] mmap(NULL, 124566, PROT_READ, MAP_PRIVATE, 3, 0) = 0x7f0add7cc000
[pid 20027] close(3)                    = 0
[pid 20027] access("/etc/ld.so.nohwcap", F_OK) = -1 ENOENT (No such file or directory)
[pid 20027] open("/lib/x86_64-linux-gnu/libm.so.6", O_RDONLY|O_CLOEXEC) = 3
[pid 20027] read(3, "\177ELF\2\1\1\0\0\0\0\0\0\0\0\0\3\0>\0\1\0\0\0pU\0\0\0\0\0\0"..., 832) = 832
[pid 20027] fstat(3, {st_mode=S_IFREG|0644, st_size=1030512, ...}) = 0
[pid 20027] mmap(NULL, 3125544, PROT_READ|PROT_EXEC, MAP_PRIVATE|MAP_DENYWRITE, 3, 0) = 0x7f0add2cf000
[pid 20027] mprotect(0x7f0add3ca000, 2093056, PROT_NONE) = 0
[pid 20027] mmap(0x7f0add5c9000, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0xfa000) = 0x7f0add5c9000
[pid 20027] close(3)                    = 0
[pid 20027] open("/usr/local/lib/libc.so.6", O_RDONLY|O_CLOEXEC) = -1 ENOENT (No such file or directory)
[pid 20027] access("/etc/ld.so.nohwcap", F_OK) = -1 ENOENT (No such file or directory)
[pid 20027] open("/lib/x86_64-linux-gnu/libc.so.6", O_RDONLY|O_CLOEXEC) = 3
[pid 20027] read(3, "\177ELF\2\1\1\0\0\0\0\0\0\0\0\0\3\0>\0\1\0\0\0\200\30\2\0\0\0\0\0"..., 832) = 832
[pid 20027] fstat(3, {st_mode=S_IFREG|0755, st_size=1815224, ...}) = 0
[pid 20027] mmap(NULL, 3929304, PROT_READ|PROT_EXEC, MAP_PRIVATE|MAP_DENYWRITE, 3, 0) = 0x7f0adcf0f000
[pid 20027] mprotect(0x7f0add0c4000, 2097152, PROT_NONE) = 0
[pid 20027] mmap(0x7f0add2c4000, 24576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x1b5000) = 0x7f0add2c4000
[pid 20027] mmap(0x7f0add2ca000, 17624, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) = 0x7f0add2ca000
[pid 20027] close(3)                    = 0
[pid 20027] mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f0add7cb000
[pid 20027] mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f0add7ca000
[pid 20027] mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f0add7c9000
[pid 20027] arch_prctl(ARCH_SET_FS, 0x7f0add7ca700) = 0
[pid 20027] mprotect(0x7f0add2c4000, 16384, PROT_READ) = 0
[pid 20027] mprotect(0x7f0add5c9000, 4096, PROT_READ) = 0
[pid 20027] mprotect(0x695000, 4096, PROT_READ) = 0
[pid 20027] mprotect(0x7f0add7ed000, 4096, PROT_READ) = 0
[pid 20027] munmap(0x7f0add7cc000, 124566) = 0
[pid 20027] rt_sigaction(SIGBUS, {0x426d10, [BUS], SA_RESTORER|SA_RESTART, 0x7f0adcf454a0}, {SIG_DFL, [], 0}, 8) = 0
[pid 20027] rt_sigaction(SIGSEGV, {0x426d10, [SEGV], SA_RESTORER|SA_RESTART, 0x7f0adcf454a0}, {SIG_DFL, [], 0}, 8) = 0
[pid 20027] mmap(NULL, 802816, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f0add705000
[pid 20027] brk(0)                      = 0x1dfb000
[pid 20027] brk(0x1e1c000)              = 0x1e1c000
[pid 20027] stat("/home/qboxtest/Desktop/qiniu", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20027] stat(".", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20027] open("/home/qboxtest/Desktop/qiniu/childProcess.go", O_RDONLY) = 3
[pid 20027] read(3, "package main\n\nimport (\n\t\"fmt\"\n\t\""..., 8192) = 331
[pid 20027] brk(0x1e43000)              = 0x1e43000
[pid 20027] brk(0x1e68000)              = 0x1e68000
[pid 20027] brk(0x1e8c000)              = 0x1e8c000
[pid 20027] mmap(NULL, 503808, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f0add68a000
[pid 20027] access("/tmp/go-build628409689/fmt.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20027] access("/tmp/go-build628409689/fmt.6", F_OK) = -1 ENOENT (No such file or directory)
[pid 20027] access("/usr/local/go/pkg/linux_amd64/fmt.a", F_OK) = 0
[pid 20027] open("/usr/local/go/pkg/linux_amd64/fmt.a", O_RDONLY) = 4
[pid 20027] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20027] close(4)                    = 0
[pid 20027] access("/tmp/go-build628409689/os/exec.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20027] access("/tmp/go-build628409689/os/exec.6", F_OK) = -1 ENOENT (No such file or directory)
[pid 20027] access("/usr/local/go/pkg/linux_amd64/os/exec.a", F_OK) = 0
[pid 20027] open("/usr/local/go/pkg/linux_amd64/os/exec.a", O_RDONLY) = 4
[pid 20027] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20027] mmap(NULL, 503808, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f0add60f000
[pid 20027] read(4, "@\"time\".t\302\2672 @\"time\".Time \"esc:0"..., 8192) = 8192
[pid 20027] mmap(NULL, 503808, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f0adce94000
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 10000} <unfinished ...>
[pid 20027] read(4, "all\".SysProcAttr; Process *@\"os\""..., 8192) = 8192
[pid 20027] mmap(NULL, 503808, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f0adce19000
[pid 20027] close(4)                    = 0
[pid 20027] access("/tmp/go-build628409689/time.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20027] access("/tmp/go-build628409689/time.6", F_OK) = -1 ENOENT (No such file or directory)
[pid 20027] access("/usr/local/go/pkg/linux_amd64/time.a", F_OK) = 0
[pid 20027] open("/usr/local/go/pkg/linux_amd64/time.a", O_RDONLY) = 4
[pid 20027] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20027] mmap(NULL, 503808, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f0adcd9e000
[pid 20027] read(4, "Timer }\n\tfunc (@\"\".t\302\2671 *@\"\".Tic"..., 8192) = 8192
[pid 20027] close(4)                    = 0
[pid 20027] read(3, "", 8192)           = 0
[pid 20027] close(3)                    = 0
[pid 20027] mmap(NULL, 503808, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f0adcd23000
[pid 20027] open("/tmp/go-build628409689/command-line-arguments.a", O_WRONLY|O_CREAT|O_TRUNC, 0666) = 3
[pid 20027] write(3, "!<arch>\n\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 239) = 239
[pid 20027] write(3, "\0", 1)           = 1
[pid 20027] lseek(3, 8, SEEK_SET)       = 8
[pid 20027] write(3, "__.PKGDEF       0           0   "..., 60) = 60
[pid 20027] lseek(3, 240, SEEK_SET)     = 240
[pid 20027] brk(0x1ead000)              = 0x1ead000
[pid 20027] write(3, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20027] write(3, "\0\0\0\0\0\0\0\0\0\0\0\10\0\20\2\0\0(type..hash.[2]"..., 4883) = 4883
[pid 20027] write(3, "\0", 1)           = 1
[pid 20027] lseek(3, 240, SEEK_SET)     = 240
[pid 20027] write(3, "_go_.6          0           0   "..., 60) = 60
[pid 20027] close(3)                    = 0
[pid 20027] exit_group(0)               = ?
Process 20018 resumed
Process 20027 detached
[pid 20018] <... wait4 resumed> [{WIFEXITED(s) && WEXITSTATUS(s) == 0}], 0, {ru_utime={0, 0}, ru_stime={0, 4000}, ...}) = 20027
[pid 20018] --- SIGCHLD (Child exited) @ 0 (0) ---
[pid 20018] rt_sigreturn(0x11)          = 20027
[pid 20018] read(4, "", 512)            = 0
[pid 20018] close(4)                    = 0
[pid 20018] stat("/usr/local/go/pkg/tool/linux_amd64/6l", {st_mode=S_IFREG|0755, st_size=1040986, ...}) = 0
[pid 20018] open("/dev/null", O_RDONLY|O_CLOEXEC) = 3
[pid 20018] pipe2([4, 5], O_CLOEXEC)    = 0
[pid 20018] stat(".", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20018] pipe2([6, 7], O_CLOEXEC)    = 0
[pid 20018] clone(Process 20028 attached
child_stack=0, flags=|SIGCHLD) = 20028
[pid 20028] chdir("." <unfinished ...>
[pid 20018] close(7 <unfinished ...>
[pid 20028] <... chdir resumed> )       = 0
[pid 20018] <... close resumed> )       = 0
[pid 20028] dup2(3, 0 <unfinished ...>
[pid 20018] read(6,  <unfinished ...>
[pid 20028] <... dup2 resumed> )        = 0
[pid 20028] dup2(5, 1)                  = 1
[pid 20028] dup2(5, 2)                  = 2
[pid 20028] execve("/usr/local/go/pkg/tool/linux_amd64/6l", ["/usr/local/go/pkg/tool/linux_amd"..., "-o", "/tmp/go-build628409689/command-l"..., "-L", "/tmp/go-build628409689", "-w", "-extld=gcc", "/tmp/go-build628409689/command-l"...], [/* 60 vars */]) = 0
[pid 20018] <... read resumed> "", 8)   = 0
[pid 20018] close(6)                    = 0
[pid 20028] brk(0)                      = 0x2006000
[pid 20018] close(3)                    = 0
[pid 20028] access("/etc/ld.so.nohwcap", F_OK <unfinished ...>
[pid 20018] close(5 <unfinished ...>
[pid 20028] <... access resumed> )      = -1 ENOENT (No such file or directory)
[pid 20018] <... close resumed> )       = 0
[pid 20028] mmap(NULL, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7ffbad239000
[pid 20018] wait4(20028, Process 20018 suspended
 <unfinished ...>
[pid 20028] access("/etc/ld.so.preload", R_OK) = -1 ENOENT (No such file or directory)
[pid 20028] open("/usr/local/lib/tls/x86_64/libm.so.6", O_RDONLY|O_CLOEXEC) = -1 ENOENT (No such file or directory)
[pid 20028] stat("/usr/local/lib/tls/x86_64", 0x7fff84f83100) = -1 ENOENT (No such file or directory)
[pid 20028] open("/usr/local/lib/tls/libm.so.6", O_RDONLY|O_CLOEXEC) = -1 ENOENT (No such file or directory)
[pid 20028] stat("/usr/local/lib/tls", 0x7fff84f83100) = -1 ENOENT (No such file or directory)
[pid 20028] open("/usr/local/lib/x86_64/libm.so.6", O_RDONLY|O_CLOEXEC) = -1 ENOENT (No such file or directory)
[pid 20028] stat("/usr/local/lib/x86_64", 0x7fff84f83100) = -1 ENOENT (No such file or directory)
[pid 20028] open("/usr/local/lib/libm.so.6", O_RDONLY|O_CLOEXEC) = -1 ENOENT (No such file or directory)
[pid 20028] stat("/usr/local/lib", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20028] open("/etc/ld.so.cache", O_RDONLY|O_CLOEXEC) = 3
[pid 20028] fstat(3, {st_mode=S_IFREG|0644, st_size=124566, ...}) = 0
[pid 20028] mmap(NULL, 124566, PROT_READ, MAP_PRIVATE, 3, 0) = 0x7ffbad21a000
[pid 20028] close(3)                    = 0
[pid 20028] access("/etc/ld.so.nohwcap", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] open("/lib/x86_64-linux-gnu/libm.so.6", O_RDONLY|O_CLOEXEC) = 3
[pid 20028] read(3, "\177ELF\2\1\1\0\0\0\0\0\0\0\0\0\3\0>\0\1\0\0\0pU\0\0\0\0\0\0"..., 832) = 832
[pid 20028] fstat(3, {st_mode=S_IFREG|0644, st_size=1030512, ...}) = 0
[pid 20028] mmap(NULL, 3125544, PROT_READ|PROT_EXEC, MAP_PRIVATE|MAP_DENYWRITE, 3, 0) = 0x7ffbacd1d000
[pid 20028] mprotect(0x7ffbace18000, 2093056, PROT_NONE) = 0
[pid 20028] mmap(0x7ffbad017000, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0xfa000) = 0x7ffbad017000
[pid 20028] close(3)                    = 0
[pid 20028] open("/usr/local/lib/libc.so.6", O_RDONLY|O_CLOEXEC) = -1 ENOENT (No such file or directory)
[pid 20028] access("/etc/ld.so.nohwcap", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] open("/lib/x86_64-linux-gnu/libc.so.6", O_RDONLY|O_CLOEXEC) = 3
[pid 20028] read(3, "\177ELF\2\1\1\0\0\0\0\0\0\0\0\0\3\0>\0\1\0\0\0\200\30\2\0\0\0\0\0"..., 832) = 832
[pid 20028] fstat(3, {st_mode=S_IFREG|0755, st_size=1815224, ...}) = 0
[pid 20028] mmap(NULL, 3929304, PROT_READ|PROT_EXEC, MAP_PRIVATE|MAP_DENYWRITE, 3, 0) = 0x7ffbac95d000
[pid 20028] mprotect(0x7ffbacb12000, 2097152, PROT_NONE) = 0
[pid 20028] mmap(0x7ffbacd12000, 24576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x1b5000) = 0x7ffbacd12000
[pid 20028] mmap(0x7ffbacd18000, 17624, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) = 0x7ffbacd18000
[pid 20028] close(3)                    = 0
[pid 20028] mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7ffbad219000
[pid 20028] mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7ffbad218000
[pid 20028] mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7ffbad217000
[pid 20028] arch_prctl(ARCH_SET_FS, 0x7ffbad218700) = 0
[pid 20028] mprotect(0x7ffbacd12000, 16384, PROT_READ) = 0
[pid 20028] mprotect(0x7ffbad017000, 4096, PROT_READ) = 0
[pid 20028] mprotect(0x63a000, 4096, PROT_READ) = 0
[pid 20028] mprotect(0x7ffbad23b000, 4096, PROT_READ) = 0
[pid 20028] munmap(0x7ffbad21a000, 124566) = 0
[pid 20028] mmap(NULL, 802816, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7ffbad153000
[pid 20028] brk(0)                      = 0x2006000
[pid 20028] brk(0x2027000)              = 0x2027000
[pid 20028] stat(".", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20028] stat(".", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20028] lstat("/tmp/go-build628409689/command-line-arguments/_obj/exe/childProcess", 0x7fff84f838f0) = -1 ENOENT (No such file or directory)
[pid 20028] unlink("/tmp/go-build628409689/command-line-arguments/_obj/exe/childProcess") = -1 ENOENT (No such file or directory)
[pid 20028] open("/tmp/go-build628409689/command-line-arguments/_obj/exe/childProcess", O_WRONLY|O_CREAT|O_TRUNC, 0775) = 3
[pid 20028] mmap(NULL, 10489856, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7ffbabf5c000
[pid 20028] access("/tmp/go-build628409689/runtime.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/runtime.a", F_OK) = 0
[pid 20028] open("/tmp/go-build628409689/command-line-arguments.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] access("/tmp/go-build628409689/fmt.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/fmt.a", F_OK) = 0
[pid 20028] access("/tmp/go-build628409689/os/exec.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/os/exec.a", F_OK) = 0
[pid 20028] access("/tmp/go-build628409689/time.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/time.a", F_OK) = 0
[pid 20028] read(4, "\16Tgclocals\302\267c34189e3b824b0bbe5cf"..., 8192) = 5124
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/runtime.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] brk(0x2048000)              = 0x2048000
[pid 20028] read(4, "p_0030\0\0\0026type.*\"\".TypeAssertion"..., 8192) = 8192
[pid 20028] read(4, "0_noctxt\0\0\0r\10\2\0\0$\"\".sizeof_C_MSt"..., 8192) = 8192
[pid 20028] read(4, "\0\200H1\353H\211\\$\20\303\0 \0\0\0\4\f\"\".~r1\0\20\4\26type"..., 8192) = 8192
[pid 20028] brk(0x2069000)              = 0x2069000
[pid 20028] read(4, "at64.go\2\376\2\24\"\".fsub64c\0\0\240\1\0\0\206\1dH\213"..., 8192) = 8192
[pid 20028] read(4, "68aa369055d8938d809aa5d80843b\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\2H\203\354\30L\213T$ L\213L$0H\213|$8H1\300H\307\306=\0\0\0H9"..., 8192) = 8192
[pid 20028] brk(0x208a000)              = 0x208a000
[pid 20028] read(4, "H\211\34$H\307D$\10\30\0\0\0H\213\\$0H\211\\$\20H\203|$\20\0\17\204\200"..., 8192) = 8192
[pid 20028] read(4, "\2\0040v/\0010\10/\0010\f\0\10\2\220\1\0\10\4\220\1\0\0\10\0\0\0\0Tgc"..., 8192) = 8192
[pid 20028] read(4, "\3dH\213\f%\0\0\0\0H;!w\7\350\0\0\0\0\353\353H\203\354\30H\213\\$ H"..., 8192) = 8192
[pid 20028] brk(0x20ab000)              = 0x20ab000
[pid 20028] read(4, "H\1\353H\211\\$\20\350\0\0\0\0H\213D$8H\213L$(H\377\301H\213l$ H"..., 8192) = 8192
[pid 20028] read(4, "\354 H\213\\$(H\211\34$H\307D$\10<\0\0\0H\213\\$8H\211\\$\20H\203"..., 8192) = 8192
[pid 20028] read(4, "\2 \0\0 \2\0\0\0\6\0\0\0#\0\0\0#\0\0\0\0\376\16Tgclocal"..., 8192) = 8192
[pid 20028] brk(0x20cc000)              = 0x20cc000
[pid 20028] read(4, "\0\0\0\376,\26\"\".allglock\0\0\20\30type.\"\".loc"..., 8192) = 8192
[pid 20028] read(4, "32]uintptr\0\2 \0\0 \0\1\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x20ed000)              = 0x20ed000
[pid 20028] read(4, "\0\0\200\1\20\2\0\0\26type.uint64\0\0\0\376\16\36type.."..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, " string\0\0\0@\20\2\0\0Xgo.string.\"func("..., 8192) = 8192
[pid 20028] read(4, "er\"\0\0\0\376\16 type.\"\".stringer\0\0\260\2\0\0\260"..., 8192) = 8192
[pid 20028] brk(0x210e000)              = 0x210e000
[pid 20028] read(4, " string\0\0\0p\20\2\0\0\"runtime.zerovalu"..., 8192) = 8192
[pid 20028] read(4, "typelink.[61]struct { Size uint3"..., 8192) = 8192
[pid 20028] read(4, "\2\0\0\34type..gc.uint8\0\0\0@\20\2\0\0\"go.st"..., 8192) = 8192
[pid 20028] brk(0x212f000)              = 0x212f000
[pid 20028] read(4, "untime.algarray\0\0\0000\20\2\0\0ztype..gc"..., 8192) = 8192
[pid 20028] read(4, "[0]\"\".method\0\2 \0\0 \0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0050\0\0\0\0\0\0\0\2068j*\0\10\10\31\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "rovalue\0\0\0\200\1\20\2\0\0\26type.uint64\0\0\0\220"..., 8192) = 8192
[pid 20028] brk(0x2150000)              = 0x2150000
[pid 20028] read(4, "me.algarray\0\0\0000\20\2\0\0\34type..gc.*\"\""..., 8192) = 8192
[pid 20028] read(4, "go.string.\"prev\"\0\0\0\376\16\"go.string."..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x2171000)              = 0x2171000
[pid 20028] read(4, "]uint8\0\2\260\1\0\0\260\1@\0\0\0\0\0\0\0\7\376\355&\0\1\1\221\0\0"..., 8192) = 8192
[pid 20028] read(4, "e.[]uint64\0\0\0\376\16Bgo.typelink.[32]"..., 8192) = 8192
[pid 20028] read(4, "\0\0\30\230x\377\0\10\10\31\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\3\30\0\0\0\0\0\0\0\223\23)\306\0\10\10\31\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x2192000)              = 0x2192000
[pid 20028] read(4, "float64\0\0\0\376\16 type..gc.float64\0\2 "..., 8192) = 8192
[pid 20028] read(4, "type..gc.[]\"\"._1_\0\0\0@\20\2\0\0002go.str"..., 8192) = 8192
[pid 20028] read(4, "type.uint8\0\0\0\360\22\20\2\0\0&go.string.\"b"..., 8192) = 8192
[pid 20028] brk(0x21b3000)              = 0x21b3000
[pid 20028] read(4, "\2\0\0006type.*\"\".BlockProfileRecord\0"..., 8192) = 8192
[pid 20028] read(4, ".*\"\".uncommonType\0\2P\0\0P\10\0\0\0\0\0\0\0\1"..., 8192) = 8192
[pid 20028] read(4, "\20\2\0\0006go.string.\"runtime._string\""..., 8192) = 8192
[pid 20028] read(4, "\20\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\376\16(type..gc.*\"\""..., 8192) = 8192
[pid 20028] brk(0x21d4000)              = 0x21d4000
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\f \20\2\300\2\0 runt"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\10\0\0\0\0\0\0\0waitnote\0\2\0\20\2 \0(go"..., 8192) = 8192
[pid 20028] read(4, "ing.\"runtime.parfor\"\0\0\0`\20\2\0\0\36typ"..., 8192) = 8192
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20028] brk(0x21f5000 <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 10000} <unfinished ...>
[pid 20028] <... brk resumed> )         = 0x21f5000
[pid 20028] read(4, "trtype\0\0\0p\20\2\0\0\"runtime.zerovalue"..., 8192) = 8192
[pid 20028] read(4, "bucket\0\0\0P\20\2\320\2\0,type.\"\".specialp"..., 8192) = 8192
[pid 20028] read(4, "\2 \0\0 \2\0\0\0\f\0\0\0\0\0\0\0\0\n\0\0\0\376\16Tgclocal"..., 8192) = 8192
[pid 20028] read(4, "\0\320\2\20\2\0\0*go.string.\"freelarge\"\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x2216000)              = 0x2216000
[pid 20028] read(4, "\".\0\0\0\220\4\20\2\0\0\34type.*\"\"._type\0\0\0P\20\2"..., 8192) = 8192
[pid 20028] read(4, "o.importpath.\"\".\0\0\0\320\1\20\2\0\0\26type.u"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\1\0\0\0\0\0\0\0008\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x2237000)              = 0x2237000
[pid 20028] read(4, "51a4c4bfeaa840e480961ec6b0b3\0\2\30\0"..., 8192) = 8192
[pid 20028] read(4, "\2\0\0\"go.importpath.\"\".\0\0\0\260\17\20\2\340\17\0\32"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0020\20\2\0\0(type.."..., 8192) = 8192
[pid 20028] brk(0x2258000)              = 0x2258000
[pid 20028] read(4, "..hash.\"\".sigtab\0\0\0\376\16*type..eq.\""..., 8192) = 8192
[pid 20028] read(4, "\0\0\20\0\0\0\0\376*\26_cgo_malloc\0\0\20\0\0\0\0\376*\16."..., 8192) = 8192
[pid 20028] read(4, "\20\2\31 \333\3\37\f\0\10\2\200\4\0h\226\4\31\10\26<A;\6<H\5\20\2%\2\4"..., 8192) = 8192
[pid 20028] brk(0x2279000)              = 0x2279000
[pid 20028] read(4, "\0002runtime.morestack8_noctxt\0\0\0J\10"..., 8192) = 8192
[pid 20028] read(4, "\2\26\0\"\22\5\21\t\22\5\21\v\0\10\0\0\0\0\22gcargs.58\2\26gc"..., 8192) = 8192
[pid 20028] read(4, "\0\0\32runtime.throw\0\0\0\204\1\10\6\0\0002runtim"..., 8192) = 8192
[pid 20028] read(4, "\0t\32H\213D$(H\213JxH\211\10H\213L$(H\211JxH\203\304\30\303H\213B"..., 8192) = 8192
[pid 20028] brk(0x229a000)              = 0x229a000
[pid 20028] read(4, "\0\10\370\377\377\377\0\376\16\24gclocals.5\2\0\10\0\0\10\370\377\377\377\0\376"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0H\203\375\0\17\204\265\r\0\0H\213]\20H\307\301\30\0\0\0H\17\257\313H"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0J\10\2\0\0\10work\2\0\0l\10\6\0\0&runtime.lf"..., 8192) = 8192
[pid 20028] read(4, "lfin\2\0\0\210\4\10\2\0\0\10finc\2\0\0PP\0\0\0\24\2\31P\224\1"..., 8192) = 8192
[pid 20028] brk(0x22bb000)              = 0x22bb000
[pid 20028] read(4, "st\2\0\0\372\1\10\2\240\1\0\10work\2\0\0\214\2\10\2\310\1\0\10work"..., 8192) = 8192
[pid 20028] read(4, "\10\2\240\1\0\10work\2\0\0\350\4\10\2\340\1\0\10work\2\0\0\376\4\10\2"..., 8192) = 8192
[pid 20028] read(4, "\220\205\4\0\32runtime.mheap\0\0\0\0200\0\0\0\16\2\0040c/"..., 8192) = 8192
[pid 20028] read(4, " %D(%D) steal, %D/%D/%D yields\n\0"..., 8192) = 8192
[pid 20028] brk(0x22dc000)              = 0x22dc000
[pid 20028] read(4, "\0\0\0\0\0\0\36\10\6\0\0004runtime.morestack16_"..., 8192) = 8192
[pid 20028] read(4, "B\n8\301t_\17\267B\10H\17\267\300H9\305r\25\17\267B\10H\17\267\300H9\305u@"..., 8192) = 8192
[pid 20028] read(4, "\20\1\0\0\0\0\0\0\0\0\376\16\22gcargs.44\2\0\30\0\0\30\1\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\2\0\0\376\1\10\6\0\0\34runtime.printf\0\0\0 `\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x22fd000)              = 0x22fd000
[pid 20028] read(4, "ring\2\0\360\7\0\0\360\7futexwakeup addr=%p "..., 8192) = 8192
[pid 20028] read(4, "\370\0t`\270\0\0\0\0H\211\4$dH\213\4%\0\0\0\0\213\200\330\0\0\0\211D$\10"..., 8192) = 8192
[pid 20028] read(4, "\213\f%\0\0\0\0H;!w\7\350\0\0\0\0\353\353H\203\354pH\213L$xH\203\301\24"..., 8192) = 8192
[pid 20028] brk(0x231e000)              = 0x231e000
[pid 20028] read(4, "\303\362\17\20\305\362\17^\4%\0\0\0\0\362\17\20\350\362\17\20\4%\0\0\0\0f\17.\350s"..., 8192) = 8192
[pid 20028] read(4, "\22\0\0\0\0\0\0\36\10\6\0\0002runtime.morestack8_"..., 8192) = 8192
[pid 20028] read(4, "\32needaddgcproc\2\0\0\216\1\10\2\0\0\32runtime."..., 8192) = 8192
[pid 20028] brk(0x233f000)              = 0x233f000
[pid 20028] read(4, "\17\205o\1\0\0\203<%\0\0\0\0\0\17\205a\1\0\0\271\0\0\0\0H\211\f$\350\0\0"..., 8192) = 8192
[pid 20028] read(4, "\271\377\377\377\377\211L$\10\350\0\0\0\0\211\301\203\370\0\211D$\24}\22\270\0\0\0\0H\211"..., 8192) = 8192
[pid 20028] read(4, "$\350\0\0\0\0H\211D$\30H\215L$0H\211\f$\350\0\0\0\0H\211D$\20H\213"..., 8192) = 8192
[pid 20028] read(4, "\r\4$\2\23\7\7\0\0026\0\"\22\r\21<\22\5\21\27\22\n\02112\0051\f\22\n\21B"..., 8192) = 8192
[pid 20028] brk(0x2360000)              = 0x2360000
[pid 20028] read(4, "ock\0\0\0\250\2\10\2\20\0\10prof\2\0\0\270\2\10\2\0\0\10prof\2"..., 8192) = 8192
[pid 20028] read(4, "\2\31\200\1\307\4\0\10\2\340\4\0\250\1\320'\31\f\t\2\10\2\2\4\v\2\5\6\10\2\5\2"..., 8192) = 8192
[pid 20028] read(4, "\2\16\2\33\0\2.\0M\22\20\21O\"\t!G\22\t\21\35*\5)\t\22\5\21\33\0\10\0"..., 8192) = 8192
[pid 20028] read(4, "5\2\0\10\0\0\10\360\377\377\377\0\376\16\30gclocals.120\2\0\10\0\0"..., 8192) = 8192
[pid 20028] brk(0x2381000)              = 0x2381000
[pid 20028] read(4, "l/go/src/pkg/runtime/race0.c\2\376\16\24"..., 8192) = 8192
[pid 20028] read(4, "g\2\0\0\244\n\10\6\0\0\32runtime.throw\0\0\0\256\n\10\6\0"..., 8192) = 8192
[pid 20028] read(4, "\2\5\4\22!\27\2\16\2\25\4\5\21\24\10\5\27\5\7\34\2\5O\6\2\r\2\v\2\0160\r"..., 8192) = 8192
[pid 20028] brk(0x23a2000)              = 0x23a2000
[pid 20028] read(4, "\321H\213\270(\1\0\0H\211|$\20\211\336\377\306\203\346\0371\355\203\375\17s/\213\306dH\213"..., 8192) = 8192
[pid 20028] read(4, "\32\4\f\2\6\6\10\35\20L\5'#\2\17\2\"\10\26\20\25\2\v\2\23\2\36\2:\4\10K"..., 8192) = 8192
[pid 20028] read(4, "\4\0\0\0\n\0\0\0\0\376\16\22gcargs.14\2\0\30\0\0\30\1\0\0\0\10"..., 8192) = 8192
[pid 20028] read(4, "\7\1\5\1\7\1\5\7\n\35\n\2\f\2\f\16\5\t\f\2\f\10\21\0\2\330\2\0.\22\r\21"..., 8192) = 8192
[pid 20028] read(4, "\0\20gcargs.8\2\24gclocals.9\2\0\0\0\0\2J/us"..., 8192) = 8192
[pid 20028] brk(0x23c3000)              = 0x23c3000
[pid 20028] read(4, "me.morestack00_noctxt\0\0\0008\10\2\0\0\"ru"..., 8192) = 8192
[pid 20028] read(4, "untime.nohash\0\0\0\320\5\20\2\0\0\36runtime.n"..., 8192) = 8192
[pid 20028] read(4, "$XH\203\3040\3031\311\353\340\4&\10\6\0\0&runtime.getcal"..., 8192) = 8192
[pid 20028] brk(0x23e4000)              = 0x23e4000
[pid 20028] read(4, "ue\2\0\0\234\30\10\6\0\0\22selunlock\2\0\0\304\30\10\6\0\0 r"..., 8192) = 8192
[pid 20028] read(4, "$@\362\17\\\301\362\17^\306\362\17\21D$`\353\235\351:\377\377\377\17\212\360\376\377\377f\17W"..., 8192) = 8192
[pid 20028] read(4, "\4s\3\200\301\4\17\266\301H\211\351H\211\372:\4/\17\204\227\376\377\377\270\0\0\0\0H\211\4"..., 8192) = 8192
[pid 20028] read(4, "\1\0\0\0H\323\340H\377\310H\213L$XH!\301\17\267E\24H\17\267\300H\17\257\301H\213"..., 8192) = 8192
[pid 20028] read(4, "\2R/usr/local/go/src/pkg/runtime/"..., 8192) = 8192
[pid 20028] brk(0x2405000)              = 0x2405000
[pid 20028] read(4, "gcargs.44\2\0\30\0\0\30\1\0\0\0\10\0\0\0\252\0\0\0\0\376\16\22g"..., 8192) = 8192
[pid 20028] read(4, "o/src/pkg/runtime/iface.goc\2\376\2\"r"..., 8192) = 8192
[pid 20028] read(4, "\0\0H\307D$ \0\0\0\0H\203|$\10\0u\23H\307D$\30\0\0\0\0H\307D$"..., 8192) = 8192
[pid 20028] brk(0x2426000)              = 0x2426000
[pid 20028] read(4, "d\0H\22\5\21d\"\n!,\22\24\21E2\0171\202\1\n\n\t8\"\n!22\0171\34"..., 8192) = 8192
[pid 20028] read(4, "D$ H\367\330H\211D$\10\350\0\0\0\0H\213D$\20H\203\304\30\303H\307D$(\10"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0u\26\270\0\0\0\0H\211\4$\350\0\0\0\0H\2134%\0\0\0\0H1\3331"..., 8192) = 8192
[pid 20028] read(4, "\0H\215L$PH\211\f$\350\0\0\0\0H\211D$8H\215L$PH\211\f$\350\0\0"..., 8192) = 8192
[pid 20028] brk(0x2447000)              = 0x2447000
[pid 20028] read(4, "\0\0\0\36\10\6\0\0004runtime.morestack24_noc"..., 8192) = 8192
[pid 20028] read(4, "\2\0 runtime.memStats\0\0\0\316\1\10\6\0\0.run"..., 8192) = 8192
[pid 20028] read(4, "\0`\0\0RdH\213\f%\0\0\0\0H;!w\7\350\0\0\0\0\353\353H\307D$\10\0"..., 8192) = 8192
[pid 20028] brk(0x2468000)              = 0x2468000
[pid 20028] read(4, "\0\0\4\0\0\0\n\0\0\0\0\376\16\20gcargs.6\2\0\30\0\0\30\1\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "$@H\203\372\0u H\213L$8H\215\24%\0\0\0\0H\213\2H\211\1H\213B\10H"..., 8192) = 8192
[pid 20028] read(4, "ck40_noctxt\0\0\0\322\1\10\6\0\0$runtime.run"..., 8192) = 8192
[pid 20028] brk(0x2489000)              = 0x2489000
[pid 20028] read(4, "\30\350\0\0\0\0H\203\304 \303\6\n\10\22\0\0\0\0\0\0\36\10\6\0\0004runti"..., 8192) = 8192
[pid 20028] read(4, "local/go/src/pkg/runtime/time.go"..., 8192) = 8192
[pid 20028] read(4, "\4\2\4\2\5\2\4\2\4\6\t\2\3\2\3\2\2\2\5\2\2\2\t\2\4\2\1\4\2\2\1\2"..., 8192) = 8192
[pid 20028] read(4, "\260\1\0\n\0c\2M\0\10\0\0\0\0$gcargs_reflectcal"..., 8192) = 8192
[pid 20028] brk(0x24aa000)              = 0x24aa000
[pid 20028] read(4, "\t\5\2\5\2\3\2\3\0\0\0\2R/usr/local/go/src/p"..., 8192) = 8192
[pid 20028] read(4, "\306\10H\211\17H\203\307\10H\213\16H\203\306\10H\211\17H\203\307\10H\213\16H\203\306\10H\211"..., 8192) = 8192
[pid 20028] read(4, "ocal/go/src/pkg/runtime/sys_linu"..., 8192) = 1726
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/fmt.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] access("/tmp/go-build628409689/math.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/math.a", F_OK) = 0
[pid 20028] access("/tmp/go-build628409689/strconv.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/strconv.a", F_OK) = 0
[pid 20028] access("/tmp/go-build628409689/unicode/utf8.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/unicode/utf8.a", F_OK) = 0
[pid 20028] access("/tmp/go-build628409689/errors.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/errors.a", F_OK) = 0
[pid 20028] access("/tmp/go-build628409689/io.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/io.a", F_OK) = 0
[pid 20028] access("/tmp/go-build628409689/os.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/os.a", F_OK) = 0
[pid 20028] access("/tmp/go-build628409689/reflect.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/reflect.a", F_OK) = 0
[pid 20028] access("/tmp/go-build628409689/sync.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/sync.a", F_OK) = 0
[pid 20028] read(4, "\0\0\34type.\"\".buffer\0\0\0\204\r\10\6\0\0\"runti"..., 8192) = 8192
[pid 20028] brk(0x24cb000)              = 0x24cb000
[pid 20028] read(4, "\n\2\3\2\25\24j\0024\2\f\2\252\1\2\r\4\6\2\20\2\3\2\20\2\3\2~\2\r\2\f"..., 8192) = 8192
[pid 20028] read(4, "\1\0\0I\203\372\0t[H\211\274$\0\1\0\0I\211:L\211\234$\10\1\0\0M\211Z\10"..., 8192) = 8192
[pid 20028] read(4, "_noctxt\0\0\0\256\1\10\6\0\0*\"\".(*fmt).forma"..., 8192) = 8192
[pid 20028] brk(0x24ec000)              = 0x24ec000
[pid 20028] read(4, "runtime.growslice\0\0\0\272\t\10\6\0\0.unico"..., 8192) = 8192
[pid 20028] read(4, "rc/pkg/fmt/print.go\2\376\2\22\"\".Fprint"..., 8192) = 8192
[pid 20028] brk(0x250d000)              = 0x250d000
[pid 20028] read(4, "tmp_0332\0\0\2\20type.int\0\36\"\".autotmp"..., 8192) = 8192
[pid 20028] read(4, "r/local/go/src/pkg/fmt/print.go\2"..., 8192) = 8192
[pid 20028] read(4, "\313H\203\373\0~XH\307\4$\0\0\0\0H\211\264$\350\1\0\0H\211t$\10L\211D$"..., 8192) = 8192
[pid 20028] read(4, "\0\10\"\".b\0\237\6\2\36type.*\"\".buffer\0\f\"\".e"..., 8192) = 8192
[pid 20028] brk(0x252e000)              = 0x252e000
[pid 20028] read(4, ".autotmp_0578\0\0\2\20type.int\0\36\"\".au"..., 8192) = 8192
[pid 20028] read(4, "\\$xH\211\\$\10H\213\234$\200\0\0\0H\211\\$\20\350\0\0\0\0H\213\264$\350\0"..., 8192) = 8192
[pid 20028] read(4, "$\330\7\0\0H\211\34$H\213\234$H\5\0\0H\211\\$\10H\213\234$P\5\0\0H\211"..., 8192) = 8192
[pid 20028] read(4, "H\203\371\0\17\204\211\10\0\0H\2139H\213q\10H\213Q\20H\213i\30H\211\204$\20\1\0"..., 8192) = 8192
[pid 20028] read(4, "lue.Type\0\0\0\320\r\0\n\0\0\0\0\0\0\322\16\10\6\0\0&refl"..., 8192) = 8192
[pid 20028] brk(0x254f000)              = 0x254f000
[pid 20028] read(4, "\20type.int\0\36\"\".autotmp_0786\0\0\2\30ty"..., 8192) = 8192
[pid 20028] brk(0x2570000)              = 0x2570000
[pid 20028] read(4, "(\30\16\2\354\1[\5\\\16\4\354\1_\5`\16\v\25\v\25\4\272\2\6\5\5\22\t\7\t\25"..., 8192) = 8192
[pid 20028] read(4, "$@H\211\214$\210\0\0\0H\213\254$\220\0\0\0H9\351\17\214\324\1\0\0H\203\370\0\17"..., 8192) = 8192
[pid 20028] read(4, "autotmp_1030\0\0\2\26type.uint64\0\36\"\"."..., 8192) = 8192
[pid 20028] brk(0x2591000)              = 0x2591000
[pid 20028] read(4, "a4c4bfeaa840e480961ec6b0b3\0\0\0\0\0\2"..., 8192) = 8192
[pid 20028] read(4, "o.EOF\0\0\0\314\3\0\n\0\0\0\0\0\0\340\4\10\2\0\0\fio.EOF\0"..., 8192) = 8192
[pid 20028] read(4, "8\0\10\"\".r\0\0\4\"type.*\"\".readRune\0\32\2\31"..., 8192) = 8192
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20028] brk(0x25b2000)              = 0x25b2000
[pid 20021] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20021] <... futex resumed> )       = 1
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] read(4,  <unfinished ...>
[pid 20028] read(4, ").UnreadRune\0\0\0\274\1\10\6\0\0\30\"\".indexRu"..., 8192) = 8192
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20021] futex(0xc20805a9f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20025] <... futex resumed> )       = 0
[pid 20021] <... futex resumed> )       = 1
[pid 20025] futex(0xc20805a9f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20028] read(4, "\0\0\0H\213t$ H\211t$@H\213D$(H\213T$0H\211T$hH\203\370\0"..., 8192) = 8192
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20021] futex(0xa80cb8, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20028] read(4, "\213\\$@H\211\\$\10\350\0\0\0\0\362\17\20D$\20\362\17\21\204$\300\0\0\0H\201\304"..., 8192) = 8192
[pid 20028] brk(0x25d3000)              = 0x25d3000
[pid 20028] read(4, "\t\n\n\2 \2'\4;\v\364\1\7\5\10\20\0\4\36\0\340\1R\5Q\266\1\202\1\22\201\1"..., 8192) = 8192
[pid 20028] read(4, "ime.concatstring2\0\0\0\2600\10\6\0\0(\"\".(*"..., 8192) = 8192
[pid 20028] read(4, "\304X\303H\2114$\350\0\0\0\0H\213L$p\213D$\10\213\\$$9\303t\34H\213\\"..., 8192) = 8192
[pid 20028] brk(0x25f4000)              = 0x25f4000
[pid 20028] read(4, "\0\36runtime.memhash\0\0\0\332\1\10\6\0\0\36runti"..., 8192) = 8192
[pid 20028] read(4, "Tgclocals\302\267ac0a7273fe7c6b3d334b1"..., 8192) = 8192
[pid 20028] read(4, "\376\16Tgclocals\302\267d16406adcea6e8f1fac"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\200\26"..., 8192) = 8192
[pid 20028] brk(0x2615000)              = 0x2615000
[pid 20028] read(4, "\"\0\0\0\376\16\34go.string.\"eE\"\0\0020\0\0&\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "bytes.7\0\0\10\0\0\10map[\0\376*\32\"\"..gobytes"..., 8192) = 8192
[pid 20028] read(4, "Value.Field\0\0\0\376\16*reflect.Value.K"..., 8192) = 8192
[pid 20028] brk(0x2636000)              = 0x2636000
[pid 20028] read(4, "alue.SetInt\302\267f\0\2\20\0\0\20\0\0\0\0\0\0\0\0\2\0\20\2"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0?\370z2\0\10\10\23\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "float32)\0\0\0p\20\2\0\0\"runtime.zeroval"..., 8192) = 8192
[pid 20028] brk(0x2657000)              = 0x2657000
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "]uint8, string)\"\0\0\0\376\0164type.func("..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0!\0\0\0\0\0\0\0func(int64, uint64"..., 8192) = 8192
[pid 20028] brk(0x2678000)              = 0x2678000
[pid 20028] read(4, "g.\"fmt_q\"\0\0\0\220\20\20\2\0\0\"go.importpath"..., 8192) = 8192
[pid 20028] read(4, "value\0\0\0\200\1\20\2\260\1\0\24type.\"\".pp\0\0\0\260\1\20"..., 8192) = 8192
[pid 20028] read(4, "]interface {})\0\0\0\360\1\20\2\0\0\26type.*\"\""..., 8192) = 8192
[pid 20028] read(4, "func(*fmt.pp, string, int32, boo"..., 8192) = 8192
[pid 20028] brk(0x2699000)              = 0x2699000
[pid 20028] read(4, "nc(int32)\0\0\0p\20\2\0\0\"runtime.zerova"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\376\16<go.string.\"func(int64, "..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0fmt.State\0\2\0\20\2 \0*go.strin"..., 8192) = 8192
[pid 20028] brk(0x26ba000)              = 0x26ba000
[pid 20028] read(4, "ror)\0\2\0\20\2 \0rgo.string.\"func(*fmt"..., 8192) = 8192
[pid 20028] read(4, ")\0\0\0\300\1\20\2\200\2\0\"type.func(*\"\".ss)\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "g, bool)\0\2\0\20\2 \0Zgo.string.\"func("..., 8192) = 8192
[pid 20028] brk(0x26db000)              = 0x26db000
[pid 20028] read(4, "\10\23\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\26 \20\2\300\3\0 runtime.algarray\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\24 \20\2\300\3\0 runti"..., 8192) = 8192
[pid 20028] brk(0x26fc000)              = 0x26fc000
[pid 20028] read(4, "g, bool) bool\0\0\0\360\10\20\2\0\0Htype.func"..., 8192) = 8192
[pid 20028] read(4, "h.\"\".\0\0\0\240\2\20\2\320\2\0\"type.\"\".scanErro"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\376\16Bgo.string.\"func() "..., 8192) = 8192
[pid 20028] read(4, "\0\0\2\0\20\2 \0&go.string.\"runtime\"\0\0\0\376"..., 8192) = 2712
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/os/exec.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] lseek(4, 18224, SEEK_SET)   = 18224
[pid 20028] read(4, "_go_.6          0           0   "..., 8192) = 8192
[pid 20028] access("/tmp/go-build628409689/bytes.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/bytes.a", F_OK) = 0
[pid 20028] access("/tmp/go-build628409689/path/filepath.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/path/filepath.a", F_OK) = 0
[pid 20028] access("/tmp/go-build628409689/strings.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/strings.a", F_OK) = 0
[pid 20028] access("/tmp/go-build628409689/syscall.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/syscall.a", F_OK) = 0
[pid 20028] brk(0x271d000)              = 0x271d000
[pid 20028] read(4, "autotmp_0027\0\217\2\2\32type.*os.File\0\n"..., 8192) = 8192
[pid 20028] read(4, "\0\0H\211\204$\20\1\0\0H\201\304\320\0\0\0\303H\203\371\1~iH\203\371\0\17\206\26\2"..., 8192) = 8192
[pid 20028] brk(0x273e000)              = 0x273e000
[pid 20028] read(4, "\0H1\355H9\350t\35H\213T$HH\211\204$\340\0\0\0H\211\224$\350\0\0\0H\201"..., 8192) = 8192
[pid 20028] read(4, "\1}LH\307\4$\0\0\0\0H\211\224$\350\0\0\0H\211T$\10H\211\214$\360\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0\36\"\".autotmp_0315\0\0\2\26type.*uint8"..., 8192) = 8192
[pid 20028] brk(0x275f000)              = 0x275f000
[pid 20028] read(4, "or\0\0\0\334\n\10\6\0\0 runtime.typ2Itab\0\0\0\214"..., 8192) = 8192
[pid 20028] read(4, "\4\20\0\210\0012\n1\36\0\f\0\210\1\2(\0\10\0\0\0\0Tgclocals\302"..., 8192) = 8192
[pid 20028] read(4, "\302\2673280bececceccd33cb74587feedb1f"..., 8192) = 8192
[pid 20028] brk(0x2780000)              = 0x2780000
[pid 20028] read(4, "als\302\267ee79072684efc713cbc89f2036d"..., 8192) = 8192
[pid 20028] read(4, "\302\2674\0P\4\24type.int64\0\20\"\"..this\0\0\4\"t"..., 8192) = 8192
[pid 20028] read(4, "8f76149a707fd2f0025c2178a3\0\2\30\0\0\30"..., 8192) = 8192
[pid 20028] brk(0x27a1000)              = 0x27a1000
[pid 20028] read(4, "os.Pipe\302\267f\0\2\20\0\0\20\0\0\0\0\0\0\0\0\2\0\20\2\0\0\16o"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\1\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "(io.WriteCloser, error)\0\0\0\300\1\20\2\200\2"..., 8192) = 8192
[pid 20028] brk(0x27c2000)              = 0x27c2000
[pid 20028] read(4, ", error)\0\0\0p\20\2\0\0\"runtime.zeroval"..., 8192) = 8192
[pid 20028] read(4, "ng.\"ProcessState\"\0\0\0\360\7\20\2\0\0*type."..., 8192) = 8192
[pid 20028] read(4, "\0P\20\0\0\0\0\0\0\0\1\0\0\0\0\0\0\0\10\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, ") string\0\2\220\2\0\0\220\2\10\0\0\0\0\0\0\0\345\207 \262\0\10\10\23"..., 8192) = 8192
[pid 20028] brk(0x27e3000)              = 0x27e3000
[pid 20028] read(4, "portpath.os.\0\0\0\360\10\20\2\0\0 type.func("..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\220\3\20\2\0\0,\"\".(*ExitError).Exited"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\26 \20\2\300\3"..., 8192) = 8192
[pid 20028] brk(0x2804000)              = 0x2804000
[pid 20028] read(4, " (os.FileInfo, error)\0\0\0\300\1\20\2\200\2\0Z"..., 8192) = 8192
[pid 20028] read(4, "error)\0\0\0\360\1\20\2\0\0\20type.int\0\0\0\200\2\20\2\0"..., 8192) = 8192
[pid 20028] read(4, "ng.\"Stat\"\0\0\0\240\n\20\2\0\0@type.func() ("..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x2825000)              = 0x2825000
[pid 20028] read(4, "unc(\"\".closeOnce, int64) error\0\0"..., 8192) = 8192
[pid 20028] read(4, "ng.\"R\"\0\0\0\376\16Ttype.struct { F uint"..., 8192) = 8192
[pid 20028] read(4, "os.close\302\267f\0\2\20\0\0\20\0\0\0\0\0\0\0\0\2\0\20\2\0\0000"..., 8192) = 3182
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/time.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] lseek(4, 9964, SEEK_SET)    = 9964
[pid 20028] read(4, "_go_.6          0           0   "..., 8192) = 8192
[pid 20028] brk(0x2846000)              = 0x2846000
[pid 20028] read(4, "untime.panicindex\0\0\0\364(\10\2\0\0 go.st"..., 8192) = 8192
[pid 20028] read(4, "\2\26type.string\0\36\"\".autotmp_0107\0\0"..., 8192) = 8192
[pid 20028] brk(0x2867000)              = 0x2867000
[pid 20028] read(4, "\24\10\2\0\0\30type.[]uint8\0\0\0\336\24\10\6\0\0\"runt"..., 8192) = 8192
[pid 20028] read(4, "\20H\213D$HI\270\211\210\210\210\210\210\210\210H\211\303I\367\350H\211\325H\1\335H\301\375\5"..., 8192) = 8192
[pid 20028] read(4, "endUint\0\0\0\3461\10\2\0\0\30type.[]uint8\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x2888000)              = 0x2888000
[pid 20028] read(4, "type.string\0\10\"\".m\0\377\6\2\32type.\"\".Mo"..., 8192) = 8192
[pid 20028] read(4, "2\0@\4\26type.string\0\22\"\".prefix\0 \4\26t"..., 8192) = 8192
[pid 20028] brk(0x28ab000)              = 0x28ab000
[pid 20028] read(4, "\0\0H\203\370\6\17\202\307\0\0\0H\211\317H\203\307\4H\307\305\2\0\0\0H\203\370\6\17\202"..., 8192) = 8192
[pid 20028] read(4, "nicindex\0\0\0\230y\10\6\0\0$runtime.panici"..., 8192) = 8192
[pid 20028] read(4, "\21;\7\37\10\5\v\n\1\r\254\2\5\201\2\r\2\6\2 \376\1\5\367\1F\2\36\6N\2\5"..., 8192) = 8192
[pid 20028] brk(0x28cc000)              = 0x28cc000
[pid 20028] read(4, "\230\1\0\0H\307\204$\240\0\0\0\0\0\0\0H\307\204$\250\0\0\0\0\0\0\0H\307\4$"..., 8192) = 8192
[pid 20028] read(4, "als\302\267ed258766649289fd72399b11905"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0H\213\\$(H\211\234$\230\0\0\0H\201\304\210\0\0\0\303\211\4%\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x28ed000)              = 0x28ed000
[pid 20028] read(4, "L\211\320H)\360H\203\300\7I\271%I\222$I\222$IH\211\305I\367\351H\211\327H\321\377"..., 8192) = 8192
[pid 20028] read(4, "H\211|$`H\203\304(\303\350\0\0\0\0\17\v\350\0\0\0\0\17\vH\307\301\1\0\0\0\353"..., 8192) = 8192
[pid 20028] read(4, "\10\2\0\0\32\"\".daysBefore\0\0\0\252\4\10\6\0\0$runt"..., 8192) = 8192
[pid 20028] brk(0x290e000)              = 0x290e000
[pid 20028] read(4, "time.typ2Itab\0\0\0\210\10\10\2\0\0`go.string"..., 8192) = 8192
[pid 20028] read(4, "src/pkg/time/time.go\2\376\2\16\"\".norm\0"..., 8192) = 8192
[pid 20028] read(4, "1a029c0a2dd9f03ed9dd3411d\0Tgcloc"..., 8192) = 8192
[pid 20028] brk(0x292f000)              = 0x292f000
[pid 20028] read(4, "@\303\350\0\0\0\0\17\v\10\n\10\22\0\0\0\0\0\0\36\10\6\0\0004runtime"..., 8192) = 8192
[pid 20028] read(4, "\"\".autotmp_1345\0\277\7\2\20type.int\0\36\"\""..., 8192) = 8192
[pid 20028] read(4, "$\260\0\0\0\3519\375\377\377\350\0\0\0\0\17\v\350\0\0\0\0\17\v\350\0\0\0\0\17\vH"..., 8192) = 8192
[pid 20028] read(4, "ype.uint64\0\36\"\".autotmp_1468\0\0\2\30t"..., 8192) = 8192
[pid 20028] brk(0x2950000)              = 0x2950000
[pid 20028] read(4, "L\211D$\30\350\0\0\0\0L\213\214$\310\0\0\0L\213\204$\320\0\0\0H\213D$ H"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0Pp\0\0\10\36\"\".autotmp_1578\0\1\2\22t"..., 8192) = 8192
[pid 20028] brk(0x2971000)              = 0x2971000
[pid 20028] read(4, "03e00f7\0Tgclocals\302\2673280bececcecc"..., 8192) = 8192
[pid 20028] read(4, "\1\0\10>\300\1\0\4\n\0nbR\0\n\0n\2R\0\10\0\0\0\0Tgcloca"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0N\10\2\0\0*$f64.0000000000000000\0\0"..., 8192) = 8192
[pid 20028] read(4, "\213D$8H\213L$(H\377\301H\213l$ H9\351|\220H\203\3040\303\211\3\353\314\6"..., 8192) = 8192
[pid 20028] brk(0x2992000)              = 0x2992000
[pid 20028] read(4, "pe.*uintptr\0\24\2\31`\217\1_\1`\7\0\10\2\260\1\0\10\4\260\1"..., 8192) = 8192
[pid 20028] read(4, "me: invalid duration \"\0\2P\0\0P\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0\30\1\0\0\0\f\0\0\0%\0\0\0\0\376\16jgo.string.\"ti"..., 8192) = 8192
[pid 20028] brk(0x29b3000)              = 0x29b3000
[pid 20028] read(4, "\2\260\5\0\0\260\5\f\0\0\0\304\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, " time \"\0\0\0@\20\2 \0 go.string.\" as \""..., 8192) = 8192
[pid 20028] read(4, "\2\0\20\2\0\0\34\"\".Time.IsZero\0\0\0\376\16\34\"\".Ti"..., 8192) = 8192
[pid 20028] brk(0x29d4000)              = 0x29d4000
[pid 20028] read(4, "\376\16\"type..gc.*\"\".zone\0\2P\0\0P\10\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0@type..gc.func(*\"\".Location) b"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x29f5000)              = 0x29f5000
[pid 20028] read(4, ".string.\"Zone\"\0\0\0\376\16Tgclocals\302\26732"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\1\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\1\0\0\0\0\0\0\0\1\0"..., 8192) = 8192
[pid 20028] read(4, "\16<type..gc.func(\"\".Month) string"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0func(*time.Time) *time.Locat"..., 8192) = 8192
[pid 20028] brk(0x2a16000)              = 0x2a16000
[pid 20028] read(4, "\0\2\220\2\0\0\220\2\10\0\0\0\0\0\0\0\2\327\242\224\0\10\10\23\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "() (int, int)\0\2\220\2\0\0\220\2\10\0\0\0\0\0\0\0\204\23I"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x2a37000)              = 0x2a37000
[pid 20028] read(4, "Time) \"\".Time\0\0\0\300\1\20\2\200\2\0004type.fun"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\1\20\2\0\0\30type.[]uint8\0\0\0\376\16>go.typel"..., 8192) = 8192
[pid 20028] read(4, "\0\2\0\20\2 \0 go.string.\"keys\"\0\0\0\376\16$go"..., 8192) = 8192
[pid 20028] brk(0x2a58000)              = 0x2a58000
[pid 20028] read(4, "ration) bool\0\0\0@\20\2\0\0bgo.string.\""..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\376\0160go.string.\"[1"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x2a79000)              = 0x2a79000
[pid 20028] read(4, "Ftype.struct { a string; b float"..., 8192) = 8192
[pid 20028] read(4, "te\0\0\0\376\16(\"\".(*Time).Second\302\267f\0\2\20\0"..., 8192) = 3642
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/math.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] read(4, "dH\213\f%\0\0\0\0H;!w\7\350\0\0\0\0\353\353H\203\354 \362\17\20T$(\362"..., 8192) = 8192
[pid 20028] brk(0x2a9a000)              = 0x2a9a000
[pid 20028] read(4, "\"\".autotmp_0034\0\0\2\30type.float64\0"..., 8192) = 8192
[pid 20028] read(4, "\r\1\0\0H\203\370\0\177\23\362\17\20\f%\0\0\0\0f\17.\314\17\207\364\0\0\0H1\300"..., 8192) = 8192
[pid 20028] read(4, "e.bool\0\f\"\".~r1\0\20\4\30type.float64\0\10"..., 8192) = 8192
[pid 20028] brk(0x2abb000)              = 0x2abb000
[pid 20028] read(4, "&\v\"\4\2\7\"\10\5\27\23\2\30\3\24\3\36\5\f\1\v\1\32\0\4\10\0\360\4\0\16\0"..., 8192) = 8192
[pid 20028] read(4, "\22type.bool\0\36\"\".autotmp_0129\0001\2\22t"..., 8192) = 8192
[pid 20028] read(4, "pe.bool\0\10\"\".x\0\0\4\30type.float64\0\32\2"..., 8192) = 8192
[pid 20028] brk(0x2adc000)              = 0x2adc000
[pid 20028] read(4, "\".s\0O\2\30type.float64\0\f\"\".~r1\0\20\4\30t"..., 8192) = 8192
[pid 20028] read(4, "*$f64.3f2a8c896c257764\0\0\0\364\6\10\2\0\0*"..., 8192) = 8192
[pid 20028] read(4, "\0\0\2\30type.float64\0\36\"\".autotmp_023"..., 8192) = 8192
[pid 20028] brk(0x2afd000)              = 0x2afd000
[pid 20028] read(4, "\17Y\312\353\314H\203\373\7\17\205\22\375\377\377f\17(\321\362\17\20\f%\0\0\0\0\362\17X\313"..., 8192) = 8192
[pid 20028] read(4, ".\312u#z!H\307\4$\377\377\377\377\350\0\0\0\0\362\17\20D$\10\362\17\21\204$\210\0"..., 8192) = 8192
[pid 20028] read(4, "346\0\0\2\22type.bool\0\36\"\".autotmp_034"..., 8192) = 8192
[pid 20028] brk(0x2b1e000)              = 0x2b1e000
[pid 20028] read(4, "\"\".Inf\0\0\0\206\23\10\2\0\0*$f64.7feffffffff"..., 8192) = 8192
[pid 20028] read(4, ".j\0\37\2\24type.int64\0\16\"\".sign\0e\2\22typ"..., 8192) = 8192
[pid 20028] read(4, "tmp_0463\0\0\2\26type.uint64\0\36\"\".auto"..., 8192) = 8192
[pid 20028] brk(0x2b3f000)              = 0x2b3f000
[pid 20028] read(4, "xt\0\0\0\206\2\10\6\0\0\36runtime.f64hash\0\0\0000P"..., 8192) = 8192
[pid 20028] read(4, "\2\20\0\0\20\1\0\0\0\0\0\0\0\0\376\16Tgclocals\302\267f90cf"..., 8192) = 8192
[pid 20028] read(4, "0\25cm\201@m\364\30>\231M\301@f\25\220\16\324\22\342@\0\376,\16\"\".q0S"..., 8192) = 8192
[pid 20028] brk(0x2b60000)              = 0x2b60000
[pid 20028] read(4, "2340f90cf3f379f66fe4b2d382\0\2\30\0\0\30"..., 8192) = 8192
[pid 20028] read(4, "float64\0\2@\0\0@\10\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x2b81000)              = 0x2b81000
[pid 20028] read(4, "\2\0\0*$f64.3ff0000000000000\0\0\0\377\377\377\377"..., 8192) = 5778
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/strconv.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] read(4, "\274\ndH\213\f%\0\0\0\0H;!w\7\350\0\0\0\0\353\353H\203\3540H\213T$8"..., 8192) = 8192
[pid 20028] read(4, "\16\"\".bits\0\37\2\26type.uint64\0\10\"\".n\0O\2"..., 8192) = 8192
[pid 20028] brk(0x2ba2000)              = 0x2ba2000
[pid 20028] read(4, "\0\0H\213+H\211\254$\0\4\0\0H\213k\10H\211\254$\10\4\0\0H\213\234$(\4\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0\236\r\10\6\0\0$runtime.panicindex\0\0\0\336\17"..., 8192) = 8192
[pid 20028] read(4, "ntime.panicslice\0\0\0\312\v\10\6\0\0$runtim"..., 8192) = 8192
[pid 20028] brk(0x2bc3000)              = 0x2bc3000
[pid 20028] read(4, ".autotmp_0278\0\0\2\20type.int\0\36\"\".au"..., 8192) = 8192
[pid 20028] read(4, "\0\0\2\26type.uint64\0\36\"\".autotmp_0329"..., 8192) = 8192
[pid 20028] read(4, "t8H\213\nH\213B\10H\213j\20H\203\370\0v \306\0011H\307B\30\1\0\0\0H\213"..., 8192) = 8192
[pid 20028] brk(0x2be4000)              = 0x2be4000
[pid 20028] read(4, "\0001\300\350\0\0\0\0H1\377I\203\372\0\17\214;\5\0\0M1\311A\200\371\0D\210L$"..., 8192) = 8192
[pid 20028] read(4, "locals\302\2675941ee083e489a3a7d0c562d"..., 8192) = 8192
[pid 20028] read(4, "\16\0\272\1\2\266\21\0\10\0\0\0\0Tgclocals\302\2675941ee08"..., 8192) = 8192
[pid 20028] brk(0x2c05000)              = 0x2c05000
[pid 20028] read(4, "\10\6\0\0004runtime.morestack01_noctxt\0"..., 8192) = 8192
[pid 20028] read(4, "\211T$\20H\211\204$\320\0\0\0H\211D$\30H\307D$ \1\0\0\0\350\0\0\0\0H"..., 8192) = 8192
[pid 20028] read(4, "type.int\0\36\"\".autotmp_0735\0\0\2\20typ"..., 8192) = 8192
[pid 20028] brk(0x2c26000)              = 0x2c26000
[pid 20028] read(4, "ote.go\2\376\2\36\"\".CanBackquote\0\0\340\1\0\0\340"..., 8192) = 8192
[pid 20028] read(4, "[4]uint8\0\36\"\".autotmp_0952\0\0\2\30typ"..., 8192) = 8192
[pid 20028] brk(0x2c47000)              = 0x2c47000
[pid 20028] read(4, "bbe5cf2ca4e567d435\0Tgclocals\302\26732"..., 8192) = 8192
[pid 20028] read(4, "oat32\0\10\"\".s\0\20\4\30type.uintptr\0\n\"\"."..., 8192) = 8192
[pid 20028] read(4, "cab038d51064a089bda21fa03e00f7\0\2"..., 8192) = 8192
[pid 20028] brk(0x2c68000)              = 0x2c68000
[pid 20028] read(4, "\302\2670115f8d53b75c1696444f08ad03251"..., 8192) = 8192
[pid 20028] read(4, "\274\326B\0\0004&\365k\fC\0\200\3407y\303AC\0\240\330\205W4vC\0\310Ngm"..., 8192) = 8192
[pid 20028] read(4, ").ShortestDecimal\0\0\0\376\16*\"\".adjust"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x2c89000)              = 0x2c89000
[pid 20028] read(4, "\0\0\0\0\0\0\376\16*go.string.\"func(int)\"\0\2"..., 8192) = 8192
[pid 20028] read(4, "Info) (\"\".extFloat, \"\".extFloat)"..., 8192) = 8192
[pid 20028] read(4, ".extFloat, *\"\".floatInfo) (uint6"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x2caa000)              = 0x2caa000
[pid 20028] read(4, "untime.zerovalue\0\0\0\200\1\20\2\0\0\24type.u"..., 8192) = 8192
[pid 20028] read(4, ".*[8]\"\".extFloat\0\0\0p\20\2\0\0\"runtime"..., 8192) = 8192
[pid 20028] read(4, "\20\2\0\0\"runtime.zerovalue\0\0\0\200\1\20\2\0\0\26"..., 8192) = 5564
[pid 20028] brk(0x2ccb000)              = 0x2ccb000
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/unicode/utf8.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] read(4, "e.int\0\16\"\".size\0000\4\20type.int\0\10\"\".r"..., 8192) = 8192
[pid 20028] read(4, "en\0\0\0\376\16 \"\".EncodeRune\302\267f\0\2\20\0\0\20\0\0"..., 8192) = 1250
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/errors.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 5436
[pid 20028] brk(0x2cec000)              = 0x2cec000
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/io.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] read(4, "\1\277\1\1\300\1t\0\10\2\320\3\0F\270\4<\2\10\1\3\2\5\0021\4\23\4\216\1\2\v"..., 8192) = 8192
[pid 20028] read(4, "ader\0,\2\36\240\1\206\2\237\1\1\240\1f\237\1\1\240\0010\237\1\4\0\10\2\300\3"..., 8192) = 8192
[pid 20028] read(4, "eturn\0\0\0\334\v\10\6\0\0&runtime.deferretu"..., 8192) = 8192
[pid 20028] brk(0x2d0d000)              = 0x2d0d000
[pid 20028] read(4, "\0H\211\4%\0\0\0\0H\215\34%\0\0\0\0H\215,$H\211\357H\211\336H\245H\245\350"..., 8192) = 8192
[pid 20028] read(4, "faceeq\0\0\0@\320\1\0\0\20\36\"\".autotmp_0154\0"..., 8192) = 8192
[pid 20028] read(4, " H\213D$(H\211T$XH\211L$`H\211D$hdH\213\f%\0\0\0\0\203i"..., 8192) = 8192
[pid 20028] brk(0x2d2e000)              = 0x2d2e000
[pid 20028] read(4, "44038\0\2(\0\0(\3\0\0\0\n\0\0\0\7\0\0\0\7\0\0\0\7\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "ring\"\0\0\0\320\1\20\2\0\0<type.func(string)"..., 8192) = 8192
[pid 20028] read(4, "error)\"\0\0\0`\20\2\0\0Xgo.weak.type.*fu"..., 8192) = 8192
[pid 20028] read(4, " int64) (int, error)\0\2\300\2\0\0\300\2\10\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x2d4f000)              = 0x2d4f000
[pid 20028] read(4, "int8) (int, error)\0\0\0p\20\2\0\0\"runti"..., 8192) = 8192
[pid 20028] read(4, "nc([]uint8) (int, error)\0\0\0\220\2\20\2\0"..., 8192) = 8192
[pid 20028] read(4, "PipeReader) error\0\2@\0\0@\10\0\0\0\0\0\0\0\2"..., 8192) = 8192
[pid 20028] brk(0x2d70000)              = 0x2d70000
[pid 20028] read(4, "ype.func(*\"\".PipeWriter, []uint8"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\10\10\26\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\335\0\10\10\26\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x2d91000)              = 0x2d91000
[pid 20028] read(4, "\267f\0\2\20\0\0\20\0\0\0\0\0\0\0\0\2\0\20\2\0\0000type..has"..., 8192) = 160
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/os.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] lseek(4, 20160, SEEK_SET)   = 20160
[pid 20028] read(4, "_go_.6          0           0   "..., 8192) = 8192
[pid 20028] access("/tmp/go-build628409689/sync/atomic.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/sync/atomic.a", F_OK) = 0
[pid 20028] read(4, "\0\0\0H\213\4%\0\0\0\0H\307D$P\0\0\0\0H\307D$X\0\0\0\0H\307D"..., 8192) = 8192
[pid 20028] read(4, "\0\f\"\".~r2\0@\4\24type.error\0\20\"\".value"..., 8192) = 8192
[pid 20028] brk(0x2db2000)              = 0x2db2000
[pid 20028] read(4, "l.Errno\0\0\0\260\10\10\2\0\0\24type.error\0\0\0\302\10"..., 8192) = 8192
[pid 20028] read(4, "te\0\6\2 \0\6\2 \0\10\244\1 \0\0\10\0\0\0\0Tgclocals\302"..., 8192) = 8192
[pid 20028] read(4, "autotmp_0230\0\0\0020type.*errors.err"..., 8192) = 8192
[pid 20028] brk(0x2dd3000)              = 0x2dd3000
[pid 20028] read(4, "\213.H\203\375\0tvH\215u\10L\215@\20L\211\307H\245H\245H\213l$XH\211h "..., 8192) = 8192
[pid 20028] read(4, "tab.*\"\".LinkError.error\0\0\0\222\4\10\2\0\0"..., 8192) = 8192
[pid 20028] read(4, "string.\"lchown\"\0\0\0\246\3\10\2\0\0006go.itab"..., 8192) = 8192
[pid 20028] brk(0x2df4000)              = 0x2df4000
[pid 20028] read(4, "\215(H\215\34%\0\0\0\0H\211\357H\211\336H\245H\245H\213l$XH\211h\20H\213l"..., 8192) = 8192
[pid 20028] read(4, "tmp_0454\0\277\2\2\30type.*string\0\36\"\".au"..., 8192) = 8192
[pid 20028] read(4, "\1\0\0H\211\234$0\1\0\0H\307D$x\0\0\0\0H\307\204$\200\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x2e15000)              = 0x2e15000
[pid 20028] read(4, "/os/path.go\2\376\2\30\"\".RemoveAll\0\0\200\26\0"..., 8192) = 8192
[pid 20028] read(4, "\211L$hH\211L$XH\307D$0\0\0\0\0H\307D$8\0\0\0\0H\203\370\0H"..., 8192) = 8192
[pid 20028] read(4, "pe.string\0\10\"\".c\0\327\1\2\24type.int32\0\10"..., 8192) = 8192
[pid 20028] brk(0x2e36000)              = 0x2e36000
[pid 20028] read(4, "c0a0e34e\0\0\0\0\0\2H/usr/local/go/src"..., 8192) = 8192
[pid 20028] read(4, "ess\0\2\240\1\0\0\240\1dH\213\f%\0\0\0\0H;!w\7\350\0\0\0\0\353\353"..., 8192) = 8192
[pid 20028] read(4, "\1/\00104\0\10\2\220\2\0\10\4\220\2\0\4&\0A2\0051.2\0051*2\0051*"..., 8192) = 8192
[pid 20028] brk(0x2e57000)              = 0x2e57000
[pid 20028] read(4, "cals\302\267519efd86263089ddb84df3cfe7"..., 8192) = 8192
[pid 20028] read(4, "b76\0\2(\0\0(\3\0\0\0\n\0\0\0\26\0\0\0\26\0\0\0\26\0\0\0\0\376\16"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\376\16Tgclocals\302\267a08e9001cb8f9d"..., 8192) = 8192
[pid 20028] brk(0x2e78000)              = 0x2e78000
[pid 20028] read(4, "\0\0\2\0\20\2\0\0\16\"\".itod\0\0\0\376\0168\"\".(*Proce"..., 8192) = 8192
[pid 20028] read(4, "\2\20\0\0\20\0\0\0\0\0\0\0\0\2\0\20\2\0\0\30syscall.init"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\240\2\20\2\0\0\26type.string\0\0\0\320\2\20\2\0\0&g"..., 8192) = 8192
[pid 20028] brk(0x2e99000)              = 0x2e99000
[pid 20028] read(4, "g.\"func() bool\"\0\2@\0\0008\0\0\0\0\0\0\0\0\v\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\22 \20\2\300\3\0 runtime"..., 8192) = 8192
[pid 20028] read(4, "runtime.zerovalue\0\0\0\200\1\20\2\0\0 type."..., 8192) = 8192
[pid 20028] brk(0x2eba000)              = 0x2eba000
[pid 20028] read(4, "c([]uint8) (int, error)\"\0\0\0`\20\2\0\0"..., 8192) = 8192
[pid 20028] read(4, "\".File\0\0\0\340\1\20\2\0\0\"go.string.\"Chdir"..., 8192) = 8192
[pid 20028] read(4, "nal\"\0\0\0\376\16$go.string.\"Signal\"\0\0020\0"..., 8192) = 8192
[pid 20028] read(4, "@\0\0@\10\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x2edb000)              = 0x2edb000
[pid 20028] read(4, ".ProcessState, error)\"\0\0\0\376\16Jtype"..., 8192) = 8192
[pid 20028] read(4, "or\0\0\0\260\1\20\2\0\0\34go.string.\"Op\"\0\0\0\320\1\20"..., 8192) = 8192
[pid 20028] read(4, "type.**[7]string\0\0\0p\20\2\0\0\"runtime"..., 8192) = 8192
[pid 20028] read(4, "\2\0\0\220\2\10\0\0\0\0\0\0\0FP\255\276\0\10\10\23\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x2efc000)              = 0x2efc000
[pid 20028] read(4, "cfe7fd2992\0\2\30\0\0\30\1\0\0\0\2\0\0\0\2\0\0\0\0\376\16T"..., 8192) = 6632
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/reflect.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] lseek(4, 13486, SEEK_SET)   = 13486
[pid 20028] read(4, "_go_.6          0           0   "..., 8192) = 8192
[pid 20028] read(4, ".autotmp_0068\0\0\2\32type.\"\".Value\0\36"..., 8192) = 8192
[pid 20028] brk(0x2f1d000)              = 0x2f1d000
[pid 20028] read(4, "87feedb1f9f\0\0\0\0\0\2J/usr/local/go/"..., 8192) = 8192
[pid 20028] read(4, "\0H\215\254$\370\0\0\0H\211\357H\211\336\350\0\0\0\0\210\204$P\1\0\0H\201\304\330\0"..., 8192) = 8192
[pid 20028] read(4, "unc\0\0\240\4\0\0\210\4dH\213\f%\0\0\0\0H\215D$\200H;\1w\f\270\200"..., 8192) = 8192
[pid 20028] brk(0x2f3e000)              = 0x2f3e000
[pid 20028] read(4, "5ed09429deb0079b48cbc19ea\0\0\0\0\0\2J"..., 8192) = 8192
[pid 20028] read(4, "H\211\330H\307\204$\20\1\0\0\0\0\0\0H\307\204$\30\1\0\0\0\0\0\0H\307\204$\340"..., 8192) = 8192
[pid 20028] read(4, "\2\20type.int\0\36\"\".autotmp_0360\0\237\2\2J"..., 8192) = 8192
[pid 20028] brk(0x2f5f000)              = 0x2f5f000
[pid 20028] read(4, "Y\17H\203\343\177H\17\266\333H\203\373\24t\n\306D$h\0H\203\304P\303H\211\310H\213Y"..., 8192) = 8192
[pid 20028] read(4, "\377\377\377\350\0\0\0\0\17\vH\211\310\351@\376\377\377\211\6\351\224\376\377\377\350\0\0\0\0\17\v"..., 8192) = 8192
[pid 20028] read(4, "\357H\211\336H\245H\245H\213\274$\200\0\0\0H\213w H\203\376\0\17\204\314\3\0\0H\215"..., 8192) = 8192
[pid 20028] brk(0x2f80000)              = 0x2f80000
[pid 20028] read(4, "\".offset\0\257\2\2\30type.uintptr\0\n\"\".gc"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0H\203\370\0H\211\204$\250\0\0\0t\30H\211\204$p\1\0\0H\211\214$x\1\0"..., 8192) = 8192
[pid 20028] read(4, ".layoutCache\0\0\0\260\10\10\6\0\0.sync.(*RWM"..., 8192) = 8192
[pid 20028] brk(0x2fa1000)              = 0x2fa1000
[pid 20028] read(4, "type.\"\".Value\0\f\"\".~r0\0c\2\22type.bo"..., 8192) = 8192
[pid 20028] read(4, "ng.\"reflect.Value.Addr of unaddr"..., 8192) = 8192
[pid 20028] read(4, ",$H\211\357H\211\336H\245H\245L\211\214$\350\2\0\0L\211L$\20L\211\204$\360\2\0"..., 8192) = 8192
[pid 20028] brk(0x2fc2000)              = 0x2fc2000
[pid 20028] read(4, "\n\"\".in\0`\4\36type.[]\"\".Value\0\n\"\".op"..., 8192) = 8192
[pid 20028] read(4, "t\0\200\1\4\34type.*\"\".rtype\0\26\"\".rcvrtyp"..., 8192) = 8192
[pid 20028] read(4, "ictmp_1038\0\0\0\366\4\10\0020\0\"\"\".statictmp"..., 8192) = 8192
[pid 20028] brk(0x2fe3000)              = 0x2fe3000
[pid 20028] read(4, "068\0\0\2\30type.\"\".flag\0\36\"\".autotmp_"..., 8192) = 8192
[pid 20028] read(4, "\3514\377\377\377H\213(H\17\266]\17H\201\343\200\0\0\0\200\373\0t*H1\311\200\371\0t"..., 8192) = 8192
[pid 20028] read(4, "\0\0Tgclocals\302\2676e70bcebc200316565c"..., 8192) = 8192
[pid 20028] brk(0x3004000)              = 0x3004000
[pid 20028] read(4, "e\0\f\"\".~r0\0\225\1\2\22type.bool\0\f\"\".~r0\0"..., 8192) = 8192
[pid 20028] read(4, "range in SetLen\"\0\0\0\372\1\10\2\0\0\26type.s"..., 8192) = 8192
[pid 20028] read(4, "\377H\307\4$\0\0\0\0\350\0\0\0\0H\213D$\10H\203\370\0t6H\215(H\215\34%"..., 8192) = 8192
[pid 20028] brk(0x3025000)              = 0x3025000
[pid 20028] read(4, "e.string\0\0\0\324\5\10\6\0\0\36runtime.convT2"..., 8192) = 8192
[pid 20028] read(4, "\0\0H\211+H\213\204$\0\2\0\0H\203\370\1\17\205j\5\0\0H\213\254$\10\2\0\0H"..., 8192) = 8192
[pid 20028] read(4, "\"\".statictmp_1440\0\0\0\260\n\10\2\20\0\"\"\".st"..., 8192) = 8192
[pid 20028] read(4, "\4$\0\0\0\0\350\0\0\0\0H\213\\$\10H\211\234$\360\0\0\0H\213\234$\270\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x3046000)              = 0x3046000
[pid 20028] read(4, "ictmp_1536\0\0\0\352\2\10\2\20\0\"\"\".statictmp"..., 8192) = 8192
[pid 20028] read(4, "locals\302\2673280bececceccd33cb74587f"..., 8192) = 8192
[pid 20028] read(4, "(\10\6\0\0004runtime.morestack01_noctxt"..., 8192) = 8192
[pid 20028] brk(0x3067000)              = 0x3067000
[pid 20028] read(4, "utotmp_1632\0\321\1\2\22type.bool\0\36\"\".au"..., 8192) = 8192
[pid 20028] read(4, "\0\353\353H\203\354\20\203A\20\30H\213\\$ H\211\34$H\213\\$\30H\213\233\0\1\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0&go.string.\"reflect\"\0\0\0\210\1\10\2\0\0\"g"..., 8192) = 8192
[pid 20028] brk(0x3088000)              = 0x3088000
[pid 20028] read(4, "tring.\"reflect\"\0\0\0\254\1\10\2\0\0\"go.stri"..., 8192) = 8192
[pid 20028] read(4, "ue).IsNil\0\2\200\3\0\0\342\2dH\213\f%\0\0\0\0H;!w\7\350"..., 8192) = 8192
[pid 20028] read(4, "w\7\350\0\0\0\0\353\353H\203\3540\203A\0208H\213\\$8H1\355H9\353uKH\215"..., 8192) = 8192
[pid 20028] read(4, "%\0\0\0\0H\215l$ H\211\357H\211\336H\245H\245\350\0\0\0\0\17\vH\213t$8"..., 8192) = 8192
[pid 20028] brk(0x30a9000)              = 0x30a9000
[pid 20028] read(4, "\211\336H\245H\245H\215\34%\0\0\0\0H\215l$ H\211\357H\211\336H\245H\245\350\0\0"..., 8192) = 8192
[pid 20028] read(4, "\306\2\0H\203\304H\303H\213H\20H\213x\30H\213F\20H\213v\30H9\310uDH\211D"..., 8192) = 8192
[pid 20028] read(4, "\0t\5\351\0\0\0\0\211\4%\0\0\0\0\353\362\4\20\10\6\314\3\0 runtime"..., 8192) = 8192
[pid 20028] brk(0x30ca000)              = 0x30ca000
[pid 20028] read(4, "\"\".(*interfaceType).Elem\0\2\200\1\0\0dH"..., 8192) = 8192
[pid 20028] read(4, "Align\0\2`\0\0RH\307D$\20\0\0\0\0H\213\\$\10H\211\\$\10H\203"..., 8192) = 8192
[pid 20028] read(4, "|$\10\0t\5\351\0\0\0\0\211\4%\0\0\0\0\353\362\0020\10\16\0\0002\"\".(*"..., 8192) = 8192
[pid 20028] brk(0x30eb000)              = 0x30eb000
[pid 20028] read(4, "\0\0H\213\\$\10H\211\\$\10H\203|$\10\0t\5\351\0\0\0\0\211\4%\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "duffzero\0\0\0(\10\6\324\3\0 runtime.duffze"..., 8192) = 8192
[pid 20028] read(4, "*\"\".uncommonType\0\20\"\"..this\0\0\4 ty"..., 8192) = 8192
[pid 20028] brk(0x310c000)              = 0x310c000
[pid 20028] read(4, "ogenerated>\2\376\2(\"\".(*ptrType).Num"..., 8192) = 8192
[pid 20028] read(4, "sliceType).Elem\0\2\200\1\0\0dH\307D$\20\0\0\0\0H"..., 8192) = 8192
[pid 20028] read(4, "clocals\302\26706cab038d51064a089bda21"..., 8192) = 8192
[pid 20028] brk(0x312d000)              = 0x312d000
[pid 20028] read(4, "\26744568aa369055d8938d809aa5d80843"..., 8192) = 8192
[pid 20028] read(4, "36fc6efe5a7cef63b040e21731\0\2 \0\0 "..., 8192) = 8192
[pid 20028] read(4, "ng.\"chan<- \"\0\0020\0\0000\0\0\0\0\0\0\0\0\7\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "280bececceccd33cb74587feedb1f9f\0"..., 8192) = 8192
[pid 20028] brk(0x314e000)              = 0x314e000
[pid 20028] read(4, "go.string.\"can't Index(i) with i"..., 8192) = 8192
[pid 20028] read(4, "\0\0\26\0\0\0\0\0\0\0reflect.Value.SetFloat"..., 8192) = 8192
[pid 20028] read(4, "\f\0\0\0\22\0\0\0\7\0\0\0\7Z\0\0\7Z\0\0\7\0\0\0\7Z\0\0\7Z\0\0"..., 8192) = 8192
[pid 20028] read(4, "0\0\0\0\0\0\0\0\0\7\0\0\0\0\0\0\0float64\0\2\0\20\2 \0&"..., 8192) = 8192
[pid 20028] brk(0x316f000)              = 0x316f000
[pid 20028] read(4, "\0\0\0\2\0\20\2\0\0\34\"\".Value.Index\0\0\0\376\16\"\"\""..., 8192) = 8192
[pid 20028] read(4, "rted\0\0\0\376\0166\"\".flag.mustBeAssignab"..., 8192) = 8192
[pid 20028] brk(0x3190000)              = 0x3190000
[pid 20028] read(4, "\16Tgclocals\302\2673280bececceccd33cb74"..., 8192) = 8192
[pid 20028] read(4, ".string.\"func() reflect.ChanDir\""..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\360\5\20\2\0\0\"go.importpath.\"\".\0\0\0\200\6"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\23\0\0\0\0\0\0\0func() reflect.Kind\0\2"..., 8192) = 8192
[pid 20028] brk(0x31b1000)              = 0x31b1000
[pid 20028] read(4, "ring.\"func(*reflect.uncommonType"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0d \20\2\0\0$t"..., 8192) = 8192
[pid 20028] read(4, "\20\2\0\0\"runtime.zerovalue\0\0\0\220\1\20\2\360\1\0"..., 8192) = 8192
[pid 20028] read(4, "mplements\"\0\0\0\376\16\34go.string.\"In\"\0\2"..., 8192) = 8192
[pid 20028] brk(0x31d2000)              = 0x31d2000
[pid 20028] read(4, "\0\0\0\240\25\20\2\0\0 go.string.\"Size\"\0\0\0\300\25\20"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\1\0\0\0\0\0\0\0\1\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "edb1f9f\0\2\20\0\0\20\1\0\0\0\0\0\0\0\0\376\16Tgclocal"..., 8192) = 8192
[pid 20028] brk(0x31f3000)              = 0x31f3000
[pid 20028] read(4, "uint8\0\2@\0\0@\10\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "ng) bool) reflect.Value\"\0\0\0`\20\2\0\0"..., 8192) = 8192
[pid 20028] read(4, "lue, float64) bool\0\2\240\2\0\0\240\2\10\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x3214000)              = 0x3214000
[pid 20028] read(4, "ovalue\0\0\0\220\1\20\2\360\1\0006type.func(*\"\".V"..., 8192) = 8192
[pid 20028] read(4, " \"\".Value\0\2@\0\0@\10\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "ntptr\"\0\0\0`\20\2\0\0>go.weak.type.*fun"..., 8192) = 8192
[pid 20028] brk(0x3235000)              = 0x3235000
[pid 20028] read(4, "nt64)\"\0\0\0\376\16 type.func(int64)\0\2\200\2"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\2\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\376\16Xg"..., 8192) = 8192
[pid 20028] read(4, "nc() \"\".Value\0\0\0\220\v\20\2\0\0:type.func"..., 8192) = 8192
[pid 20028] read(4, "\20\2\0\0$go.string.\"mustBe\"\0\0\0\2600\20\2\0\0"..., 8192) = 8192
[pid 20028] brk(0x3256000)              = 0x3256000
[pid 20028] read(4, ") \"\".Value\0\0\0@\20\2\0\0jgo.string.\"fu"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\37\0\0\0\0\0\0\0func(reflect.Valu"..., 8192) = 8192
[pid 20028] read(4, "nc(reflect.Value, unsafe.Pointer"..., 8192) = 8192
[pid 20028] read(4, "lect.Value) []reflect.Value\"\0\0\0\376"..., 8192) = 8192
[pid 20028] brk(0x3277000)              = 0x3277000
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\220%\20\2\0\0>type.func(\"\".Value, co"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "t[reflect.visit]bool\0\2\0\20\2 \0Rgo.s"..., 8192) = 8192
[pid 20028] brk(0x3298000)              = 0x3298000
[pid 20028] read(4, "\0\0\376\16Tgclocals\302\2673280bececceccd33c"..., 8192) = 8192
[pid 20028] read(4, " \0\216\1go.string.\"func(*reflect.fun"..., 8192) = 8192
[pid 20028] read(4, "unc(*\"\".funcType) *\"\".rtype\0\0\0p\20"..., 8192) = 8192
[pid 20028] read(4, "\0\0Btype.func(*\"\".funcType) *\"\".r"..., 8192) = 8192
[pid 20028] brk(0x32b9000)              = 0x32b9000
[pid 20028] read(4, " runtime.algarray\0\0\0000\20\2\0\0002type.."..., 8192) = 8192
[pid 20028] read(4, "\0\1\0\0\0\0\0\0\0\1\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, ".*\"\".interfaceType\0\0\0\200\2\20\2\0\0\26type"..., 8192) = 8192
[pid 20028] read(4, "\0\0\340\f\20\2\0\0.\"\".(*interfaceType).Key"..., 8192) = 8192
[pid 20028] brk(0x32da000)              = 0x32da000
[pid 20028] read(4, "bececceccd33cb74587feedb1f9f\0\2\20\0"..., 8192) = 8192
[pid 20028] read(4, "\0ltype.func(*\"\".chanType, string"..., 8192) = 8192
[pid 20028] read(4, "\0\1\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\22 \20\2\300\3\0"..., 8192) = 8192
[pid 20028] read(4, "ash.\"\".arrayType\0\0\0\20\20\2\0\0*type..e"..., 8192) = 8192
[pid 20028] brk(0x32fb000)              = 0x32fb000
[pid 20028] read(4, "\2\10\0\0\0\0\0\0\0\367Ki\37\0\10\10\23\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\1\0\0\0\0\0\0\0\1\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "bleTo\0\0\0\300\5\20\2\0\0 go.string.\"Elem\"\0"..., 8192) = 8192
[pid 20028] brk(0x331c000)              = 0x331c000
[pid 20028] read(4, "P\20\2\300\4\0\36type.\"\".mapType\0\0\0\300\4\20\2\0\0&"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\300\1\20\2\220\2\0Ttype.func(*\"\".mapType"..., 8192) = 8192
[pid 20028] read(4, "reflect.mapType) uintptr\"\0\2`\0\0^\0"..., 8192) = 8192
[pid 20028] read(4, "mapType).NumOut\0\0\0\340\23\20\2\0\0\36go.stri"..., 8192) = 8192
[pid 20028] brk(0x333d000)              = 0x333d000
[pid 20028] read(4, "c\0\2\30\0\0\30\1\0\0\0\10\0\0\0.\0\0\0\0\376\16Tgclocals\302"..., 8192) = 8192
[pid 20028] read(4, "type..gc.func(*\"\".ptrType) bool\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0P\20\2\220\1\0 type.*\"\".ptrType\0\0\0\260\1\20"..., 8192) = 8192
[pid 20028] read(4, "\0\220\2\20\2\0\0\"go.importpath.\"\".\0\0\0\240\2\20\2"..., 8192) = 8192
[pid 20028] brk(0x335e000)              = 0x335e000
[pid 20028] read(4, "t.sliceType, int) reflect.Struct"..., 8192) = 8192
[pid 20028] read(4, "unc(*\"\".sliceType) string\0\0\0@\20\2\0"..., 8192) = 8192
[pid 20028] read(4, "(go.string.\"NumField\"\0\0\0\200\21\20\2\0\0\36t"..., 8192) = 8192
[pid 20028] brk(0x337f000)              = 0x337f000
[pid 20028] read(4, "00f7\0\2\30\0\0\30\1\0\0\0\4\0\0\0\2\0\0\0\0\376\16Tgcloca"..., 8192) = 8192
[pid 20028] read(4, "g.\"func(*reflect.structType, str"..., 8192) = 8192
[pid 20028] read(4, "Type) *reflect.rtype\"\0\0\0\376\16Ftype."..., 8192) = 8192
[pid 20028] read(4, "\0\260\26\20\2\0\0@type.func(*\"\".structType"..., 8192) = 8192
[pid 20028] brk(0x33a0000)              = 0x33a0000
[pid 20028] read(4, "\0\2\0\20\2 \0\\go.string.\"map.bucket[*r"..., 8192) = 8192
[pid 20028] read(4, "Zgo.string.\"*struct { F uintptr;"..., 8192) = 8192
[pid 20028] read(4, "type..gc.map[*\"\".rtype]*\"\".ptrTy"..., 8192) = 8192
[pid 20028] read(4, "\0\0[8]reflect.cacheKey\0\2\0\20\2 \0>go."..., 8192) = 8192
[pid 20028] brk(0x33c1000)              = 0x33c1000
[pid 20028] read(4, "afe.Pointer\0\0\0\376\16Zgo.typelink.[]u"..., 8192) = 8192
[pid 20028] read(4, "garray\0\0\0000\20\2\0\0*type..gc.\"\".layou"..., 8192) = 8192
[pid 20028] read(4, "\"\".layoutKey]\"\".layoutType\0\0\0\376\16h"..., 8192) = 8192
[pid 20028] brk(0x33e2000)              = 0x33e2000
[pid 20028] read(4, "*[6]string\"\0\0\0\376\16\36type.*[6]string"..., 8192) = 8192
[pid 20028] read(4, "\"reflect.stringHeader\"\0\0\0\376\0160go.s"..., 8192) = 8192
[pid 20028] read(4, "\0\376\0160type..gc.[]\"\".SelectCase\0\2P\0"..., 8192) = 8192
[pid 20028] read(4, "\0\1\0\0\0\0\0\0\0x\0\2\0\20\2 \0\32go.string.\"x\"\0"..., 8192) = 8192
[pid 20028] brk(0x3403000)              = 0x3403000
[pid 20028] read(4, "\0\0\0\2\0\20\2 \0$go.string.\"unsafe\"\0\0\0\376"..., 8192) = 8192
[pid 20028] read(4, ".\"\".visit\302\267f\0\2\20\0\0\20\0\0\0\0\0\0\0\0\2\0\20\2\0\0"..., 8192) = 8192
[pid 20028] brk(0x3424000)              = 0x3424000
[pid 20028] read(4, "ayType).Bits\302\267f\0\2\20\0\0\20\0\0\0\0\0\0\0\0\2\0\20"..., 8192) = 8192
[pid 20028] read(4, "ype).NumMethod\0\0\0\376\0162\"\".(*sliceTy"..., 8192) = 5222
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/sync.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] brk(0x3445000)              = 0x3445000
[pid 20028] read(4, " /\20\2\20\6\17\1\17H\37\1 \23\37\1 0\0\10\2\220\2\0006B\31\2'\2\v\6"..., 8192) = 8192
[pid 20028] read(4, "octxt\0\0\0~\10\6\0\0$runtime.panicindex"..., 8192) = 8192
[pid 20028] read(4, "pe.int32\0\36\"\".autotmp_0101\0\0\2\24typ"..., 8192) = 8192
[pid 20028] brk(0x3466000)              = 0x3466000
[pid 20028] read(4, "7feedb1f9f\0\2\20\0\0\20\1\0\0\0\0\0\0\0\0\376\16Tgclo"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\376\0162go.string.\"sync.syncSema\"\0"..., 8192) = 8192
[pid 20028] read(4, "\"Mutex\"\0\0020\0\0,\0\0\0\0\0\0\0\0\5\0\0\0\0\0\0\0Mut"..., 8192) = 8192
[pid 20028] brk(0x3487000)              = 0x3487000
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\376\16Rgo.string.\"func(*syn"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\5\0\0\0\0\0\0\0\5\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\2\0\20\2 \0&go.string.\"runtime\"\0\0\0\376\16,"..., 8192) = 1046
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/bytes.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] read(4, "644     165003    `\ngo object li"..., 8152) = 8152
[pid 20028] access("/tmp/go-build628409689/unicode.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/unicode.a", F_OK) = 0
[pid 20028] brk(0x34a8000)              = 0x34a8000
[pid 20028] read(4, "\230\10\0\n\0\0\0\0\0\0\366\t\10\2\0\0\fio.EOF\0\0\0\254\n\10\2\0\0"..., 8192) = 8192
[pid 20028] read(4, "\17\v\30\n\10\22\0\0\0\0\0\0\36\10\6\0\0004runtime.morest"..., 8192) = 8192
[pid 20028] read(4, "\10\6\0\0,unicode/utf8.RuneCount\0\0\0\202\6"..., 8192) = 8192
[pid 20028] brk(0x34c9000)              = 0x34c9000
[pid 20028] read(4, "H\211|$pH\211<$H\211L$xH\211L$\10H\211\234$\200\0\0\0H\211\\$\20"..., 8192) = 8192
[pid 20028] read(4, "H\213\234$\270\0\0\0H\211\234$\0\1\0\0H\211\264$\340\0\0\0H\211\360H\211\214$\370"..., 8192) = 8192
[pid 20028] brk(0x34ea000)              = 0x34ea000
[pid 20028] read(4, "db1f9f\0\0\0\0\0\2H/usr/local/go/src/p"..., 8192) = 8192
[pid 20028] read(4, "015dfabaa1c\0Tgclocals\302\2673280becec"..., 8192) = 8192
[pid 20028] read(4, "mp_0554\0\0\2\20type.int\0\36\"\".autotmp_"..., 8192) = 8192
[pid 20028] brk(0x350b000)              = 0x350b000
[pid 20028] read(4, "errors.text\302\2672\0\37\2\26type.string\0\f\""..., 8192) = 8192
[pid 20028] read(4, "\4\16\0NR\5Q-\0\n\0N\0042\0\10\0\0\0\0Tgclocals\302\2670"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\376\16Tgclocals\302\2678648ee1c36be910"..., 8192) = 8192
[pid 20028] brk(0x352c000)              = 0x352c000
[pid 20028] read(4, "\2\20\0\0\20\0\0\0\0\0\0\0\0\2\0\20\2\0\0&runtime.defe"..., 8192) = 8192
[pid 20028] read(4, "dOp\0\0\220\1\0\0\220\1\10\0\0\0\0\0\0\0.\212\260\264\0\10\10\26\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "*\"\".Buffer, io.Reader) (int64, e"..., 8192) = 8192
[pid 20028] brk(0x354d000)              = 0x354d000
[pid 20028] read(4, "t64\0\0\0\240\2\20\2\0\0\24type.error\0\0\0\376\16Dtyp"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\376\16\"go.string.\"Reset\"\0\0020\0\0,\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "nc() (uint8, error)\0\0\0\360\5\20\2\0\0Htyp"..., 8192) = 8192
[pid 20028] read(4, "cialCase }\0\2P\0\0P\10\0\0\0\0\0\0\0\1\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x356e000)              = 0x356e000
[pid 20028] read(4, "Reader, []uint8, int64) (int, er"..., 8192) = 8192
[pid 20028] read(4, "*Reader).UnreadByte\0\0\0\360\6\20\2\0\0.\"\"."..., 8192) = 1334
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/path/filepath.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] lseek(4, 10006, SEEK_SET)   = 10006
[pid 20028] read(4, "_go_.6          0           0   "..., 8192) = 8192
[pid 20028] access("/tmp/go-build628409689/sort.a", F_OK) = -1 ENOENT (No such file or directory)
[pid 20028] access("/usr/local/go/pkg/linux_amd64/sort.a", F_OK) = 0
[pid 20028] read(4, "\0H\211\24$H\211L$\10\350\0\0\0\0H\213D$ H\213\\$(H\211\234$@\1\0"..., 8192) = 8192
[pid 20028] brk(0x358f000)              = 0x358f000
[pid 20028] read(4, "H\211\214$0\1\0\0H\211\204$8\1\0\0H\201\304\30\1\0\0\303L\211\244$\230\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "D$P\0\0\0\0H\307D$X\0\0\0\0H\307D$`\0\0\0\0H\213\\$8H\211"..., 8192) = 8192
[pid 20028] brk(0x35b0000)              = 0x35b0000
[pid 20028] read(4, "nt\0\36\"\".autotmp_0227\0\0\2\20type.int\0"..., 8192) = 8192
[pid 20028] read(4, "ntime.morestack32_noctxt\0\0\0\234\1\10\6\0"..., 8192) = 8192
[pid 20028] read(4, "aJ\"\5!\207\1B\5A\235\3R\5Q\37\2j`\5a\221\1\"\7!\310\1b\5a\367"..., 8192) = 8192
[pid 20028] brk(0x35d1000)              = 0x35d1000
[pid 20028] read(4, "\300\f\0\0\376\16Tgclocals\302\26726c7ef6700d40e4"..., 8192) = 8192
[pid 20028] read(4, "]string\0\0\0 \20\2\0\0 runtime.memprint"..., 8192) = 8192
[pid 20028] read(4, "unc(int) uint8\"\0\0\0`\20\2\0\0:go.weak."..., 8192) = 6972
[pid 20028] brk(0x35f2000)              = 0x35f2000
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/strings.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] read(4, "\211D$ H\211|$`H\211<$H\211T$PH\211T$\10H\211D$\20\350\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\2`\0\4\16\0!\"\5!:\0\n\0!\2?\0\10\0\0\0\0Tgclocals"..., 8192) = 8192
[pid 20028] read(4, "\377H\377\300\351\265\373\377\377\350\0\0\0\0\17\v\350\0\0\0\0\17\vH\203y0\0\17\204\256\1"..., 8192) = 8192
[pid 20028] brk(0x3613000)              = 0x3613000
[pid 20028] read(4, "$\10H\211D$\20\350\0\0\0\0H\213\274$\210\0\0\0H\213T$`H\213D$PH\213"..., 8192) = 8192
[pid 20028] read(4, "time.memmove\0\0\0\340\16\10\2\0\0\30type.[]uin"..., 8192) = 8192
[pid 20028] read(4, "\0\0\376\3\10\2\0\0\30type.[]uint8\0\0\0\234\4\10\6\0\0\"r"..., 8192) = 8192
[pid 20028] brk(0x3634000)              = 0x3634000
[pid 20028] read(4, "\0\0\0\0\0\0\0H\307\204$\270\0\0\0\0\0\0\0H\201\304\210\0\0\0\303H\213\234$\220"..., 8192) = 8192
[pid 20028] read(4, "H\203\373\1uLH\211\360H\377\310H\203\370\0|(H9\360s4H\215\34\7H\17\266\33D"..., 8192) = 8192
[pid 20028] brk(0x3655000)              = 0x3655000
[pid 20028] read(4, "sSpace\302\267f\0\0\0\244\1\10\6\0\0\32\"\".FieldsFunc"..., 8192) = 8192
[pid 20028] read(4, "_0619\0\0\2\30type.[]uint8\0\36\"\".autotm"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0H\213\\$@H\211\34$H\213\\$HH\211\\$\10H\213\\$PH\211\\$\20"..., 8192) = 8192
[pid 20028] brk(0x3676000)              = 0x3676000
[pid 20028] read(4, "H\213\264$\250\0\0\0L9\316\17\202\324\2\0\0H\211\323H\211L$@H\211\312L\211\301H"..., 8192) = 8192
[pid 20028] read(4, "runtime.memhash\0\0\00000\0\0\6\10\"\".p\0 \4\36"..., 8192) = 8192
[pid 20028] read(4, "e0d19febc89\0\2 \0\0 \2\0\0\0\24\0\0\0\0\0\0\0\2\0\0"..., 8192) = 8192
[pid 20028] brk(0x3697000)              = 0x3697000
[pid 20028] read(4, "d33cb74587feedb1f9f\0\2\20\0\0\20\1\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0\376\16&\"\".TrimRightFunc\302\267f\0\2\20\0\0\20\0\0"..., 8192) = 8192
[pid 20028] read(4, "unc(*strings.Reader, int64, int)"..., 8192) = 8192
[pid 20028] brk(0x36b8000)              = 0x36b8000
[pid 20028] read(4, "t64, error)\0\2\0\20\2 \0Tgo.string.\"fu"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\376\16*type..gc.*\"\".repl"..., 8192) = 8192
[pid 20028] read(4, "nder, string) int\0\2\0\20\2 \0fgo.stri"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\5\0\0\0\0\0\0\0\5\0\0"..., 8192) = 8192
[pid 20028] brk(0x36d9000)              = 0x36d9000
[pid 20028] read(4, "\3\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "teReplacer\0\0\0\320\2\20\2\0\0000go.string.\"b"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0\2\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\"\0\0\0\376\16:go.string.\"stringWriterIf"..., 8192) = 8192
[pid 20028] brk(0x36fa000)              = 0x36fa000
[pid 20028] read(4, ".WriteString\0\0\0\376\16Jtype..hash.\"\"."..., 8192) = 600
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/syscall.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] lseek(4, 92024, SEEK_SET)   = 92024
[pid 20028] read(4, "_go_.6          0           0   "..., 8192) = 8192
[pid 20028] read(4, "\224$\200\0\0\0H\211T$hH\2134%\0\0\0\0H\213\24%\0\0\0\0H\213\34%\0"..., 8192) = 8192
[pid 20028] read(4, "\20\2\37\1\7\3\26\0\4\20\0Tb\na\322\1\0\22\0T\2[\2d\1\35\0\10\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x371b000)              = 0x371b000
[pid 20028] read(4, "H\203\370\0H\211D$XvgH\211L$\30\350\0\0\0\0H\213\\$0H\211\\$@H"..., 8192) = 8192
[pid 20028] read(4, "\0\0\36\10\6\0\0004runtime.morestack32_noct"..., 8192) = 8192
[pid 20028] read(4, "ocals\302\267168ae203b9134674893e678c6"..., 8192) = 8192
[pid 20028] brk(0x373c000)              = 0x373c000
[pid 20028] read(4, "|$\10H\213t$`H\203\376\0t)H\245\213\36\211\37H\213\\$@H\211\\$hH\213"..., 8192) = 8192
[pid 20028] read(4, "\4\10\6\0\0 runtime.typ2Itab\0\0\0\206\5\10\2\0\0\24"..., 8192) = 8192
[pid 20028] read(4, "\0\0H\307D$\20\0\0\0\0\350\0\0\0\0H\213D$\30\353\306\307D$ \0\0\0\0\307"..., 8192) = 8192
[pid 20028] brk(0x375d000)              = 0x375d000
[pid 20028] read(4, "\377\300L9\310|\344H\211<$L\211T$8H\203\377\0H\211|$@v#L\211T$\10"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\353\353H\203\3540H\213L$8H\307D$P\0\0\0\0H\307D$X\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0H\215D$\340H;\1w\7\350\0\0\0\0\353\346H\201\354\240\0\0\0H\307\204$\300\0\0"..., 8192) = 8192
[pid 20028] brk(0x377e000)              = 0x377e000
[pid 20028] read(4, "\0\303H\215|$h1\300\350\0\0\0\0H\211T$h\211L$pH\307D$H\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "ta\0 \4\30type.[]uint8\0\16\"\".addr\0\20\4\30t"..., 8192) = 8192
[pid 20028] read(4, "\1\200\3.\0\10\2\320\7\0Z\310\vJ\2\r\2\f\2$\2\27\2\20\1\3\2C\2\r\6\4"..., 8192) = 8192
[pid 20028] brk(0x379f000)              = 0x379f000
[pid 20028] read(4, "type.map[*uint8][]uint8\0\0\0\356\f\10\6\0\0"..., 8192) = 8192
[pid 20028] read(4, "$H\213\\$PH\211\\$\10H\213\\$XH\211\\$\20H\215\\$<H\211\\$\30H"..., 8192) = 8192
[pid 20028] read(4, "go\2\376\2\26\"\".Sendfile\0\0\200\2\0\0\376\1dH\213\f%\0\0"..., 8192) = 8192
[pid 20028] brk(0x37c0000)              = 0x37c0000
[pid 20028] read(4, "\0\0\376\2\10\2\0\0\24type.error\0\0\0\220\3\10\2\0\0,go."..., 8192) = 8192
[pid 20028] read(4, "4runtime.morestack48_noctxt\0\0\0\220\1"..., 8192) = 8192
[pid 20028] read(4, "\0\0\240\1\10\6\0\0\24\"\".Syscall\0\0\0\320\1\10\2\0\0,go."..., 8192) = 8192
[pid 20028] read(4, "\307D$\20\0\0\0\0H\307D$\30\0\0\0\0\350\0\0\0\0H\213\\$ H\211\\$@"..., 8192) = 8192
[pid 20028] brk(0x37e1000)              = 0x37e1000
[pid 20028] read(4, "H\303H\307\4$V\0\0\0H\213\\$@H\211\\$\10H\211D$\20H\307D$\30\0\0"..., 8192) = 8192
[pid 20028] read(4, "d9\0\0\0\0\0\2j/usr/local/go/src/pkg/s"..., 8192) = 8192
[pid 20028] read(4, "H\307D$\20\0\0\0\0\350\0\0\0\0H\213D$\30\353\306\20\n\10\22\0\0\0\0\0\0\36"..., 8192) = 8192
[pid 20028] brk(0x3802000)              = 0x3802000
[pid 20028] read(4, "_noctxt\0\0\0|\10\6\0\0\32\"\".RawSyscall\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\2\0\0,go.itab.\"\".Errno.error\0\0\0\336\2\10"..., 8192) = 8192
[pid 20028] read(4, "\6\0\0004runtime.morestack32_noctxt\0\0"..., 8192) = 8192
[pid 20028] read(4, "H\203\371\0\17\206\260\0\0\0H\211\330H\307\4$\22\0\0\0H\213\\$hH\211\\$\10H"..., 8192) = 8192
[pid 20028] brk(0x3823000)              = 0x3823000
[pid 20028] read(4, "o/src/pkg/syscall/zsyscall_linux"..., 8192) = 8192
[pid 20028] read(4, "r\0\0\0\232\2\10\2\0\0\32type.\"\".Errno\0\0\0\254\2\10\2\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0\20\"\"._zero\0\0\0\220\1\300\1\0\0\20\36\"\".autotmp"..., 8192) = 8192
[pid 20028] brk(0x3844000)              = 0x3844000
[pid 20028] read(4, "al/go/src/pkg/syscall/env_unix.g"..., 8192) = 8192
[pid 20028] read(4, "ocals\302\26706cab038d51064a089bda21fa"..., 8192) = 8192
[pid 20028] read(4, "nix.go\2\376\0020type..eq.\"\".SockaddrUn"..., 8192) = 8192
[pid 20028] brk(0x3865000)              = 0x3865000
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\"\0\0\0\0\0\0\0\0\0\0\0\2\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0008\5\0\0\0\4\0\0\0\0\0\0\0\0\0\0\0\3\0\0\0\0\0\0\0\0\0\0\0\0\376"..., 8192) = 8192
[pid 20028] read(4, "\0\0\35\261\0\0\35\5\0\0\0\376\16Tgclocals\302\2673280bece"..., 8192) = 8192
[pid 20028] brk(0x3886000)              = 0x3886000
[pid 20028] read(4, "f4fe760a82d507fb\0\2\20\0\0\20\3\0\0\0\0\0\0\0\0\376"..., 8192) = 8192
[pid 20028] read(4, "\0\20\2\0\0\0\0\0\0\0\0\376\16Tgclocals\302\2674978b799"..., 8192) = 8192
[pid 20028] read(4, "tl for device\"\0\2`\0\0^\0\0\0\0\0\0\0\0\36\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x38a7000)              = 0x38a7000
[pid 20028] read(4, "\0\0\0\0\0\0\0\36\0\0\0\0\0\0\0protocol wrong ty"..., 8192) = 8192
[pid 20028] read(4, "g timer expired\"\0\2P\0\0P\0\0\0\0\0\0\0\0\27\0"..., 8192) = 8192
[pid 20028] read(4, "go.string.\"connection timed out\""..., 8192) = 8192
[pid 20028] brk(0x38c8000)              = 0x38c8000
[pid 20028] read(4, "rolMessageHeaderAndData\302\267f\0\2\20\0\0\20"..., 8192) = 8192
[pid 20028] read(4, "\2\0\0\"\"\".SetsockoptByte\0\0\0\376\16&\"\".Se"..., 8192) = 8192
[pid 20028] brk(0x38e9000)              = 0x38e9000
[pid 20028] read(4, "ring\0\2\20\0\0\20\0\0\0\0\0\0\0\0\2\0\20\2\0\0\32type.[]"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "ing.\"Signal\"\0\0\0\376\16Tgclocals\302\2673280"..., 8192) = 8192
[pid 20028] brk(0x390a000)              = 0x390a000
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\f \20"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "nt\"\0\0\0\376\16\30type.*[2]int\0\2\220\1\0\0\220\1\10\0\0"..., 8192) = 8192
[pid 20028] read(4, "type.uint8\0\0\0\220\5\20\2\0\0 go.string.\"A"..., 8192) = 8192
[pid 20028] brk(0x392b000)              = 0x392b000
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0\16 \20\2\300\1\0 runtime"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0000type.*\"\".SockaddrNetlink\0\0\0\260\1\20"..., 8192) = 8192
[pid 20028] read(4, "\0\200\2\20\2\0\0\36go.string.\"Uid\"\0\0\0\240\2\20\2\0\0"..., 8192) = 8192
[pid 20028] brk(0x394c000)              = 0x394c000
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\376\0164go.string.\"syscall._C_i"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\210\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x396d000)              = 0x396d000
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\10\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "ddrUnix\"\0\2P\0\0L\0\0\0\0\0\0\0\0\25\0\0\0\0\0\0\0*s"..., 8192) = 8192
[pid 20028] read(4, "\0\200\1\20\2\0\0\34type.\"\".IPMreq\0\0\0\376\16&type"..., 8192) = 8192
[pid 20028] read(4, "P\20\2\220\1\0\34type.*\"\".Iovec\0\0\0\260\1\20\2\340\1\0\34"..., 8192) = 8192
[pid 20028] brk(0x398e000)              = 0x398e000
[pid 20028] read(4, "\0\26type.uint64\0\0\0\360\10\20\2\0\0\36go.string"..., 8192) = 8192
[pid 20028] read(4, "\0[\33359\0\10\10\26\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "or\0\0\0\360\1\20\2\0\0\30type.uintptr\0\0\0\200\2\20\2\0"..., 8192) = 8192
[pid 20028] read(4, "int; cap int }\"\0\0\0\376\16 go.string.\""..., 8192) = 8192
[pid 20028] brk(0x39af000)              = 0x39af000
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0(\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "high\0\2\0\20\2 \0*go.string.\"Totalhigh"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\34\0\0"..., 8192) = 8192
[pid 20028] brk(0x39d0000)              = 0x39d0000
[pid 20028] read(4, "]int64\"\0\0\0\376\16\32type.[4]int64\0\2\260\1\0\0"..., 8192) = 8192
[pid 20028] read(4, "h.\"\".\0\0\0\240\1\20\2\320\1\0 type.\"\"._C_short"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0000\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x39f1000)              = 0x39f1000
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\0\0\0\0\0\0\0\f\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "timeofday_sym\0\0\0 \0\n\0\0\0\0\0\0 \0\2\0\0\6\2"..., 8192) = 124
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/sync/atomic.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 7762
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/unicode.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] lseek(4, 8794, SEEK_SET)    = 8794
[pid 20028] read(4, "_go_.6          0           0   "..., 8192) = 8192
[pid 20028] brk(0x3a12000)              = 0x3a12000
[pid 20028] read(4, "H\213w\20H\211t$PH\211D$HH\203\370\0~HH\211\305H\377\315H\211\323H9\305"..., 8192) = 8192
[pid 20028] read(4, "gclocals\302\2679308e7ef08d2cc2f72ae12"..., 8192) = 8192
[pid 20028] read(4, "H\215\34%\0\0\0\0H\213+H\211l$(H\213k\10H\211l$0H\213\34%\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x3a36000)              = 0x3a36000
[pid 20028] read(4, "]*\"\".RangeTable\0\0\0\204\t\10\2\0\0\32\"\".Cate"..., 8192) = 8192
[pid 20028] read(4, "sign1\0\0\0\366X\10\2\0\0(go.string.\"Gujara"..., 8192) = 8192
[pid 20028] read(4, "pe.map[string]*\"\".RangeTable\0\0\0\272"..., 8192) = 8192
[pid 20028] brk(0x3a57000)              = 0x3a57000
[pid 20028] read(4, "e.mapassign1\0\0\0\262\352\1\10\2\0\0\34go.string"..., 8192) = 8192
[pid 20028] read(4, "f9f\0\2\20\0\0\20\1\0\0\0\0\0\0\0\0\376\16Tgclocals\302\2679"..., 8192) = 8192
[pid 20028] read(4, "anmar\0\2\0\20\2 \0&go.string.\"Myanmar\""..., 8192) = 8192
[pid 20028] brk(0x3a78000)              = 0x3a78000
[pid 20028] read(4, "\0\0\2\0\20\2\0\0\"\"\".statictmp_0064\0\0\0\376*\f"..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\0\2\0\20\2\0\0\"\"\".statictmp_0265\0\0\0\376"..., 8192) = 8192
[pid 20028] brk(0x3a99000)              = 0x3a99000
[pid 20028] read(4, "_0283\0\0\0\376*\22\"\".Lydian\0\0\20&type.*\"\""..., 8192) = 8192
[pid 20028] read(4, "eTable\0\20\0\0\0\0\0\0\0\0\2\0\20\2\0\0\"\"\".static"..., 8192) = 8192
[pid 20028] read(4, "ctmp_0068\0\0\0000\20\2\0\0\"\"\".statictmp_0"..., 8192) = 8192
[pid 20028] read(4, "\0\1\0\v\0\1\0\1\0\0\0\r\0\1\0&\0\1\0\1\0\0\0(\0\1\0:\0\1\0\1"..., 8192) = 8192
[pid 20028] brk(0x3aba000)              = 0x3aba000
[pid 20028] read(4, "\252\1\0\353\253\20\376%R\21\376\26\376\1\0\31\3760\376\27\0E\376F\376\1\0I\376L\376\1"..., 8192) = 8192
[pid 20028] read(4, "\377\1\0\236\377\237\377\1\0\340\377\346\377\1\0\350\377\356\377\1\0\371\377\375\377\1\0\0\376,\"\""..., 8192) = 8192
[pid 20028] read(4, "\"\".statictmp_0269\0\0\0\376,\"\"\".static"..., 8192) = 8192
[pid 20028] read(4, "\"\".RangeTable\0000\0\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0\2"..., 8192) = 8192
[pid 20028] brk(0x3adb000)              = 0x3adb000
[pid 20028] read(4, "p_0431\0\0\0000\20\2\0\0\"\"\".statictmp_0432"..., 8192) = 8192
[pid 20028] read(4, "W\2\0\0003\377\377\377\0\0\0\0003\377\377\377Y\2\0\0Y\2\0\0006\377\377\377\0\0\0\0"..., 8192) = 8192
[pid 20028] read(4, "f\0\2\20\0\0\20\0\0\0\0\0\0\0\0\2\0\20\2\0\0,\"\".Special"..., 8192) = 8192
[pid 20028] read(4, "\0\0`\20\2\0\0&go.weak.type.**\"\".d\0\0\0p\20"..., 8192) = 8192
[pid 20028] brk(0x3afc000)              = 0x3afc000
[pid 20028] read(4, "nt32\0\0\0\360\2\20\2\0\0Ltype.func(\"\".Speci"..., 8192) = 8192
[pid 20028] read(4, "\2\0\0 type.\"\".foldPair\0\0\0\376\16(type.."..., 8192) = 8192
[pid 20028] read(4, "\0\0\0\376\16<go.string.\"[5]unicode.Rang"..., 8192) = 8192
[pid 20028] read(4, "code.Range32\"\0\0\0`\20\2\0\0006go.weak.ty"..., 8192) = 8192
[pid 20028] brk(0x3b1d000)              = 0x3b1d000
[pid 20028] read(4, "4]\"\".Range32\0\0\0@\20\2\0\0>go.string.\""..., 8192) = 8192
[pid 20028] read(4, "ime.algarray\0\0\0000\20\2\0\0.type..gc.[7"..., 8192) = 8192
[pid 20028] read(4, "\0\0\20\0\0\0\0\0\0\0\16 \20\2\0\0 runtime.algarra"..., 8192) = 8052
[pid 20028] brk(0x3b3e000)              = 0x3b3e000
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] open("/usr/local/go/pkg/linux_amd64/sort.a", O_RDONLY) = 4
[pid 20028] read(4, "!<arch>\n__.PKGDEF       0       "..., 8192) = 8192
[pid 20028] read(4, "/go/src/pkg/sort/sort.go\2\376\2\26\"\".h"..., 8192) = 8192
[pid 20028] read(4, "b0f83707\0Tgclocals\302\2673280bececcec"..., 8192) = 8192
[pid 20028] brk(0x3b5f000)              = 0x3b5f000
[pid 20028] read(4, "lice\0\0\0\210\1\10\2\0\0\"type.\"\".Interface\0"..., 8192) = 8192
[pid 20028] read(4, "\202\1\0\10\0\0\0\0Tgclocals\302\2672dc77d960dd3e"..., 8192) = 8192
[pid 20028] read(4, "\10\22\0\0\0\0\0\0\20`\0\0\2\20\"\"..this\0\0\4(type.*"..., 8192) = 8192
[pid 20028] brk(0x3b80000)              = 0x3b80000
[pid 20028] read(4, "\2 \0\0 \2\0\0\0\6\0\0\0\0\0\0\0\7\0\0\0\0\376\16Tgclocal"..., 8192) = 8192
[pid 20028] read(4, "type.[]float64\0\0\0\376\16&type..gc.*[]"..., 8192) = 8192
[pid 20028] read(4, "ce, int, int) bool\0\0\0\360\1\20\2\0\0\"type"..., 8192) = 8192
[pid 20028] brk(0x3ba1000)              = 0x3ba1000
[pid 20028] read(4, "int) int\0\2\0\20\2 \0Pgo.string.\"func("..., 8192) = 8192
[pid 20028] read(4, "t, int)\0\0\0\300\1\20\2\240\2\0Jtype.func(*\"\"."..., 8192) = 8192
[pid 20028] read(4, "477d529823e505c33b388073\0\2\30\0\0\30\1\0"..., 8192) = 8192
[pid 20028] read(4, "tring) int\0\2@\0\0@\10\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] brk(0x3bc2000)              = 0x3bc2000
[pid 20028] read(4, "ring.\"func(*sort.reverse, int, i"..., 8192) = 6752
[pid 20028] read(4, "", 8192)           = 0
[pid 20028] close(4)                    = 0
[pid 20028] brk(0x3be6000)              = 0x3be6000
[pid 20028] brk(0x3be3000)              = 0x3be3000
[pid 20028] mmap(NULL, 135168, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7ffbad132000
[pid 20028] mmap(NULL, 200704, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7ffbad101000
[pid 20028] brk(0x3bcd000)              = 0x3bcd000
[pid 20028] mremap(0x7ffbad132000, 135168, 266240, MREMAP_MAYMOVE) = 0x7ffbad0c0000
[pid 20028] mremap(0x7ffbad101000, 200704, 397312, MREMAP_MAYMOVE) = 0x7ffbad05f000
[pid 20028] mremap(0x7ffbad0c0000, 266240, 528384, MREMAP_MAYMOVE) = 0x7ffbad0c0000
[pid 20028] lseek(3, 3072, SEEK_SET)    = 3072
[pid 20028] write(3, "dH\213\f%\360\377\377\377H;!w\7\350Mn\2\0\353\353H\203\354pH\215\\$PH\203"..., 8192) = 8192
[pid 20028] write(3, "\305Hk\355\30H\1\351\213\32\213)9\353u+H\213Z\10H\213i\10H9\353u\36H\213Z"..., 8192) = 8192
[pid 20028] write(3, " Hc\313H\301\341\4H\213T$0H\215\4\nH\0058\2\0\0H\211\302\213@\10\203\370\0"..., 8192) = 8192
[pid 20028] write(3, "\234\374\377\377H;\214$\370\2\0\0\17\207\216\374\377\377H\213\1H\203\370\0thH\215\4$H"..., 8192) = 8192
[pid 20028] write(3, "\377\377\213\200\270\0\0\0\203\370\0u\30dH\213\4%\370\377\377\377\213\200\254\0\0\0\203\370\0\17"..., 8192) = 8192
[pid 20028] write(3, "H\203\354\30H\213D$ H+\4%\310\303Y\0H\301\370\3H\211\302H\301\350\4H\301\340\3"..., 8192) = 8192
[pid 20028] write(3, "\23H\203\372\0t7\17\267B\10H\17\267\300H9\305u\10\17\266B\n8\301t_\17\267B\10"..., 8192) = 8192
[pid 20028] write(3, "\213\f%\360\377\377\377H\213I H\213I\30H\211H \351*\377\377\377\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "'\270\211>X\0H\211\4$H\213\203\240\0\0\0H\211D$\10H\211\350H\211l$\20\350\215"..., 8192) = 8192
[pid 20028] write(3, "\213\200\270\0\0\0H\211\4$\350q\1\0\0H\213D$(f\307\200\230\0\0\0\1\0H\203\304"..., 8192) = 8192
[pid 20028] write(3, "H\211\4$\350G\367\376\377H\213\204$\200\0\0\0\377\210\254\0\0\0H\203\304X\3031\311\200|"..., 8192) = 8192
[pid 20028] write(3, "c\313H\301\341\4H\211\352H\211l$\30H\215\4\nH\211\4$\211\\$\24Hc\313H\213\24"..., 8192) = 8192
[pid 20028] write(3, "\23\270\322$X\0H\211\4$\350\301g\377\377H\213L$XH\215D$@H\211\4$H\211L"..., 8192) = 8192
[pid 20028] write(3, "$\20\0\0\0\0H\213L$xH\211L$\30\307D$ \0\0\0\0H\307D$(\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\f\371\377\377\210D$PH\203\304(\3031\311\353\340\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "T$0H\213L$ H\211J\20H\211Z\30\17\266K\20\210J2H\213K\30H\211J H"..., 8192) = 8192
[pid 20028] write(3, "H\213\4%\220\202X\0H;\4%\10\203X\0s\10H\211\360H\203\304`\303\307\4$\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "|\200\377\377\211D$$\270HJX\0H\211\4$\350\252\25\376\377\307\4$\0\0\0\0H\213\214"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0H\211\330H9\323H\211T$ sFH\215L$<H\211\f$H\211\\$\10H\211"..., 8192) = 8192
[pid 20028] write(3, "\0\10\0\0\303\0\0\0\0\0\0\0\0\0\0\0dH\213\f%\360\377\377\377H\215\204$\200\360\377"..., 8192) = 8192
[pid 20028] write(3, "H)\337H)\336H\205\333t\303H\203\373\2\17\207Y\377\377\377\212\6\212L\36\377\210\7\210L\37"..., 8192) = 8192
[pid 20028] write(3, "\303\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0dH\213\f%\360\377\377\377H;!w\7\350\215"..., 8192) = 8192
[pid 20028] write(3, "\0\0H)\313H\203\373\0~OH\307\4$ \344J\0H\211t$pH\211t$\10H\211|"..., 8192) = 8192
[pid 20028] write(3, "\20\350\212\300\377\377\351U\377\377\377\211\4%\0\0\0\0\353\336H\211\f$H\203<$\0t\31H"..., 8192) = 8192
[pid 20028] write(3, "~XH\307\4$ \344J\0H\211\264$\10\1\0\0H\211t$\10L\211D$\20H\211L$"..., 8192) = 8192
[pid 20028] write(3, "\0\0H\211\204$(\3\0\0H\213\234$(\3\0\0H\211\34$H\213\234$ \3\0\0H\213"..., 8192) = 8192
[pid 20028] write(3, "\303H)\313H\203\373\1}TH\307\4$ \344J\0H\211\264$`\5\0\0H\211t$\10H"..., 8192) = 8192
[pid 20028] write(3, "\211\7\351\323\363\377\377H\203\370\32\17\204\215\324\377\377\351Z\274\377\377\211\7\351\246\272\377\377\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\214$\200\3\0\0L\213\204$p\3\0\0H\1\372L\211\204$p\3\0\0H\211\224$x\3\0"..., 8192) = 8192
[pid 20028] write(3, "\211\336H\245H\245H\213\\$PH\211\\$\20H\213\\$XH\211\\$\30\350\1\232\376\377H"..., 8192) = 8192
[pid 20028] write(3, "$\350\1\0\0H\211\\$\20H\213\234$\360\1\0\0H\211\\$\30\350\364\0\4\0H\213L$"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\353\216\211\4%\0\0\0\0\351p\377\377\377\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\377YYH\213\204$\200\0\0\0H\213L$`H\203\300\10H\377\301H\213l$XH9\351|"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\0\0H\307\204$(\1\0\0\0\0\0\0H\213\234$\0\1\0\0H\211\34$H\213\264"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\3508.\376\377\353\346H\203\3548\203A\20@H\307\204$\210\0\0\0\0\0\0\0H\307"..., 8192) = 8192
[pid 20028] write(3, "\312 8\321u\34\200\371ar\27\200\371zw\22H\377\300H9\306\177\275\306D$8\1H\203\304"..., 8192) = 8192
[pid 20028] write(3, "\214$`\2\0\0H\211\274$8\2\0\0H\211\274$ \2\0\0H\211\264$@\2\0\0H\211"..., 8192) = 8192
[pid 20028] write(3, "\0H\307\204$0\3\0\0\0\0\0\0H\307\204$8\3\0\0\0\0\0\0H\307\204$@\3\0"..., 8192) = 8192
[pid 20028] write(3, "\0\17\217\v\5\0\0H\201\372\10\1\0\0\17\217\t\2\0\0H\201\372\6\1\0\0\17\217\362\0\0"..., 8192) = 8192
[pid 20028] write(3, "H\321\375H\301\373?H)\335H\211\353Hk\333\7H\211\316H)\336L\211\323H)\363H\203\303"..., 8192) = 8192
[pid 20028] write(3, "\213D$ H9\301u\26H\213\234$\330\0\0\0H\213,%\210\271W\0H\211k\20\353\225H"..., 8192) = 8192
[pid 20028] write(3, "\17\266\33\200\373\0uGH9\306r;H\211T$0H\211\24$H\211D$8H\211D$\10"..., 8192) = 8192
[pid 20028] write(3, "H\203\350\34H\211\254$0\4\0\0H\211\254$\320\3\0\0H\211\224$8\4\0\0H\211\224$"..., 8192) = 8192
[pid 20028] write(3, "H\211\336H\245H\245H\215\34%\200\177N\0H\215l$ H\211\357H\211\336H\245H\245\350m"..., 8192) = 8192
[pid 20028] write(3, "\211\6\353\322\211\7\353\300\0\0\0\0\0\0\0\0dH\213\f%\360\377\377\377H;!w\7\350]"..., 8192) = 8192
[pid 20028] write(3, "$\30H\211L$\10\350\244.\0\0H\213T$XH\213D$PH\213t$0H\213l$\30"..., 8192) = 8192
[pid 20028] write(3, "(H1\311H\213\230 \3\0\0H\203\373\0u'H\215\34%\200oN\0H\213+H\211\254$"..., 8192) = 8192
[pid 20028] write(3, "\314\211\370\367\345\211\320\301\350\3\211\303k\333\n\211\372)\332H\377\316H\215\234$\270\0\0\0H\203"..., 8192) = 8192
[pid 20028] write(3, "dH\213\f%\360\377\377\377H\215\204$\310\371\377\377H;\1w\7\350un\374\377\353\343H\201\354"..., 8192) = 8192
[pid 20028] write(3, "\340\0\0\0H\203\274$\350\0\0\0\0\17\206\4\3\0\0H\17\266\33\300\353\4H\17\266\333H\203"..., 8192) = 8192
[pid 20028] write(3, "$\f\200\373\0u%H\213\34%\200LX\0H\211\234$\220\0\0\0H\213\34%\210LX\0H"..., 8192) = 8192
[pid 20028] write(3, "\307D$0\1\0\0\0\306D$8\0H\203\304\10\303\350I\206\372\377\17\v\307D$(\375\377\0"..., 8192) = 8192
[pid 20028] write(3, "\353\336H\201\354H\1\0\0H\213\224$x\1\0\0H\307\204$\210\1\0\0\0\0\0\0H\307\204"..., 8192) = 8192
[pid 20028] write(3, "H\211D$0H\203\371\0H\211L$(\17\204\273\0\0\0H\307\4$ $M\0\350\177v\373"..., 8192) = 8192
[pid 20028] write(3, "M\0\350\231V\373\377H\213D$\10H\213\254$\200\0\0\0H\211(H\213\254$\210\0\0\0H"..., 8192) = 8192
[pid 20028] write(3, "H\245H\211L$0H\211L$\20H\211D$8H\211D$\30\350\5Z\373\377H\213\\$ "..., 8192) = 8192
[pid 20028] write(3, "\371\377\17\v\350W\346\371\377\17\v\350\200\346\371\377\17\v\350y\346\371\377\17\v\350B\346\371\377\17\v"..., 8192) = 8192
[pid 20028] write(3, " H\211\34$H\211T$\10\350\21\0\0\0H\17\266\\$\20\210\\$0H\203\304\30\303\0\0"..., 8192) = 8192
[pid 20028] write(3, "\\\372\372\377H\213\\$0H\211\234$\200\0\0\0H\213\\$8H\211\234$\210\0\0\0H\203"..., 8192) = 8192
[pid 20028] write(3, "\20H\213D$\30H\307\4$\200\30N\0H\211\214$\320\1\0\0H\211L$\10H\211\204$\330"..., 8192) = 8192
[pid 20028] write(3, "H\211\34$H\213\234$\270\0\0\0H\211\\$\10H\213\234$\300\0\0\0H\211\\$\20H\213"..., 8192) = 8192
[pid 20028] write(3, "\10\30Q\0\362\17Y\302f\17(\320\351R\377\377\377H\203\370\20u\n\306D$x\0H\203\304@"..., 8192) = 8192
[pid 20028] write(3, "\377H\213t$(H\213l$0H\213T$8H\213L$@H\17\266\\$HH\211t$x"..., 8192) = 8192
[pid 20028] write(3, "$\270\0\0\0\0\0\0\0H\307\204$\300\0\0\0\0\0\0\0H\307\204$\310\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "$\20H\213k\30H\211,$\350\221F\377\377dH\213\f%\360\377\377\377\203i\20\20H\203\304\10"..., 8192) = 8192
[pid 20028] write(3, "H\215\34%\240\235N\0H\215l$\20H\211\357H\211\336H\245H\245H\215\34%\0\370N\0H"..., 8192) = 8192
[pid 20028] write(3, "\\$\10H\203|$\10\0t\5\351@\241\376\377\211\4%\0\0\0\0\353\362\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\\$\10H\203|$\10\0t\5\3510\207\376\377\211\4%\0\0\0\0\353\362\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "@\200\376\0@\210t$\32t\3\203\340\375H\211\24$H\203<$\0tg\211L$\34\211L$"..., 8192) = 8192
[pid 20028] write(3, "D$0\0\0\0\0H\307D$8\0\0\0\0H\307@h\0\0\0\0H\211\4$H\307D$"..., 8192) = 8192
[pid 20028] write(3, "JX\0H\211\\$\10H\211\224$\220\0\0\0H\211T$\20H\211\204$\230\0\0\0H\211D"..., 8192) = 8192
[pid 20028] write(3, "H\211t$pH\211D$xH\211\214$\200\0\0\0H\211L$hH\213\224$\220\0\0\0H"..., 8192) = 8192
[pid 20028] write(3, "\200\0\0\0\0\0\0\0H\203\371\0\17\216\241\0\0\0H\213\\$XH\203\371\0\17\206\213\0\0"..., 8192) = 8192
[pid 20028] write(3, "H\211l$0H\213\34%\20\272W\0H\211\\$ H\307\4$\0\323J\0H\213\34%\360J"..., 8192) = 8192
[pid 20028] write(3, "`\274W\0H\211\\$ H\307\4$\0\323J\0H\213\34%\20KX\0H\211\\$\10H\215"..., 6992) = 6992
[pid 20028] lseek(3, 618496, SEEK_SET)  = 618496
[pid 20028] write(3, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\10\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\10\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\10\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\10\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\2\0\0\0\0\0\0\0\20\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "@\364L\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\10\0\0\0\0\0\0\0\250&\331\232\0\10\10\26 \26X\0\0\0\0\0\0\215I\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\20\0\0\0\0\0\0\0\7\0\0\0\0\0\0\0 \0\0\0\0\0\0\0\t\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, " 8M\0\0\0\0\0\0\213J\0\0\0\0\0\36\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\320\0\0\0\0\0\0\0\1\0\0\0\0\0\0\0\10\0\0\0\0\0\0\0\0\320J\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\30\360J\0\0\0\0\0\1\0\0\0\0\0\0\0\1\0\0\0\0\0\0\0\240\276M\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\10\0\0\0\0\0\0\0\304\315\222\207\0\10\10\23`\26X\0\0\0\0\0\200\220I\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0\327O\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0P\262Q\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\0\0\0\0\0008PK\0\0\0\0\0\1\0\0\0\0\0\0\0\1\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, " pK\0\0\0\0\0\1\0\0\0\0\0\0\0\1\0\0\0\0\0\0\0\0\22N\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\240\336J\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\10\0\0\0\0\0\0\0\320\10K\274\0\10\10\23`\26X\0\0\0\0\0\240\330I\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\300\372O\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0P\262Q\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\0\0\0\0\0008\360K\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, " \20L\0\0\0\0\0\2\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0\340LN\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\240\333J\0\0\0\0\0@\207M\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\10\0\0\0\0\0\0\0\370\230x\16\0\10\10\23`\26X\0\0\0\0\0\240\336I\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\200UP\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0P\262Q\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\0\0\0\0\0008\220L\0\0\0\0\0\2\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\0\0\0\0\0\200CL\0\0\0\0\0 \222N\0\0\0\0\0000\261Q\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "@\212J\0\0\0\0\0 \335J\0\0\0\0\0\240\333J\0\0\0\0\0@jL\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\0\0\0\0\0008\360L\0\0\0\0\0\6\0\0\0\0\0\0\0\6\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\30\20M\0\0\0\0\0\2\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0`\322N\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0300M\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\0\0\0\0\0\240\336J\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0\221N\0\0\0\0\0\0\0\0\0\0\0\0\0 \345J\0\0\0\0\0\240\374J\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, " \261Q\0\0\0\0\0 \271J\0\0\0\0\0\0\0\0\0\0\0\0\0Z\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "X\0\0\0\0\0\0\0\340\36O\0\0\0\0\0p\261Q\0\0\0\0\0 \341J\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\340\255F\0\0\0\0\0\340\255F\0\0\0\0\0\300wN\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\240\263N\0\0\0\0\0 \261Q\0\0\0\0\0\0\335K\0\0\0\0\0\300\371K\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\200oH\0\0\0\0\0\200oH\0\0\0\0\0\240\366N\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0\266N\0\0\0\0\0 \261Q\0\0\0\0\0\240\354J\0\0\0\0\0@ K\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "`\240K\0\0\0\0\0@\20L\0\0\0\0\0\360\31H\0\0\0\0\0\360\31H\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\20pN\0\0\0\0\0\3\0\0\0\0\0\0\000125\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\20\220N\0\0\0\0\0\2\0\0\0\0\0\0\0Ps\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\20\260N\0\0\0\0\0\5\0\0\0\0\0\0\0fflag\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\20\320N\0\0\0\0\0\7\0\0\0\0\0\0\0stopped\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\20\360N\0\0\0\0\0\10\0\0\0\0\0\0\0GoString\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\20\20O\0\0\0\0\0\10\0\0\0\0\0\0\0cacheEnd\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0200O\0\0\0\0\0\t\0\0\0\0\0\0\0retOffset\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\20PO\0\0\0\0\0\21\0\0\0\0\0\0\0*syscall.ProcAtt"..., 8192) = 8192
[pid 20028] write(3, "\20pO\0\0\0\0\0\23\0\0\0\0\0\0\0[8]*reflect.ptrT"..., 8192) = 8192
[pid 20028] write(3, "\20\220O\0\0\0\0\0\23\0\0\0\0\0\0\0func(*reflect.fl"..., 8192) = 8192
[pid 20028] write(3, "\20\260O\0\0\0\0\0\22\0\0\0\0\0\0\0reflect.Value.Ty"..., 8192) = 8192
[pid 20028] write(3, "\20\320O\0\0\0\0\0\37\0\0\0\0\0\0\0func(*fmt.pp, co"..., 8192) = 8192
[pid 20028] write(3, "\20\360O\0\0\0\0\0\30\0\0\0\0\0\0\0func(exec.ExitEr"..., 8192) = 8192
[pid 20028] write(3, "\20\20P\0\0\0\0\0\30\0\0\0\0\0\0\0reflect.Value.Se"..., 8192) = 8192
[pid 20028] write(3, "\0200P\0\0\0\0\0#\0\0\0\0\0\0\0func(*reflect.in"..., 8192) = 8192
[pid 20028] write(3, "\20PP\0\0\0\0\0!\0\0\0\0\0\0\0struct { b bool;"..., 8192) = 8192
[pid 20028] write(3, "\20pP\0\0\0\0\0-\0\0\0\0\0\0\0func(reflect.sli"..., 8192) = 8192
[pid 20028] write(3, "ue, reflect.Value) reflect.Value"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\0\0\0\0\0\320\377\377\377\0\0\0\0\0\0\0\0\360\377\377\377\360\377\377\377\370\377\377\377"..., 8192) = 8192
[pid 20028] write(3, "\1\0\0\0\f\0\0\0j\1\0\0\1\0\0\0\16\0\0\0.\0\0\0\1\0\0\0\34\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\0\n\0\0\0\10\0\0\0\2\0\0\0\262\0\0\0\272\0\0\0\272\0\0\0\272\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\0\0\0\0\0\0\0\0\0\2\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "`\16\0\0\0\0\0\0\2\0\0\0\0\0\0\0h\16\0\0\0\0\0\0\2\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "p\36\0\0\0\0\0\0\2\0\0\0\0\0\0\0x\36\0\0\0\0\0\0\2\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "p.\0\0\0\0\0\0\2\0\0\0\0\0\0\0x.\0\0\0\0\0\0\2\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\270\245\370\377\0\0\0\0\5\0\0\0\0\0\0\0\200\16\0\0\0\0\0\0\240\245\370\377\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\320\363G\0\0\0\0\0\260\367G\0\0\0\0\0\20\360G\0\0\0\0\0\0\374G\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\260\\@\0\0\0\0\0\300\272\0\0\0\0\0\0\320]@\0\0\0\0\0\200\273\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, " rB\0\0\0\0\0p\2\2\0\0\0\0\0\320rB\0\0\0\0\0\370\2\2\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "P\277E\0\0\0\0\0008;\3\0\0\0\0\0\340\304E\0\0\0\0\0h<\3\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "pYH\0\0\0\0\0\0w\4\0\0\0\0\0\220YH\0\0\0\0\0xw\4\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\364\317P\0\0\0\0\0runtime.errorCString.Err"..., 8192) = 8192
[pid 20028] write(3, "\310\246\0\0\10\0\0\0000\0\0\0\326\246\0\0\353\246\0\0\357\246\0\0\1\0\0\0\4\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0$\360\2\0\260\25\36\16\6\2\26\2\21\2\30\6\27\2\10\2\v\2\27\2\21\2\6\2\10\2\27"..., 8192) = 8192
[pid 20028] write(3, "\2\7\2\36\2\26\2\4\2+\4\0055\27\n\16\0\0~R\5Q\t\22\n\21\33\22\5\17\17\1"..., 8192) = 8192
[pid 20028] write(3, "\7\2\5\1\10\n\5\t+\22\5\21\t\22\5\21\16\22\5\21/B\5A\17\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\321&\1\0\325&\1\0\1\0\0\0\4\0\0\0\r'\1\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "D\254P\0\0\0\0\0acquirep\0\0\2\0310\207\1/\0010o\0004\220\2\0"..., 8192) = 8192
[pid 20028] write(3, "\2\31\300\1\225\1\277\1\1\300\1\361\5\0>\240\7\0\222\6#\20\10\2\t\2\t\2\25\2\10\222"..., 8192) = 8192
[pid 20028] write(3, "@\274A\0\0\0\0\0\320\206\1\0\30\0\0\0 \0\0\0\346\206\1\0\355\206\1\0\360\206\1\0"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\244\302P\0\0\0\0\0X\260P\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "stkbucket\0\0\2\31`\223\2_\1`\346\1_\1`\212\1_\2\0V\240\5"..., 8192) = 8192
[pid 20028] write(3, "stringiter2\0\0\2\03102/\0010t\0d\300\1\0\274\6\36\4\t\4"..., 8192) = 8192
[pid 20028] write(3, "call268435456\0\0\2@\200\200\200\200\2`\377\377\377\377\1\20\0j\260"..., 8192) = 8192
[pid 20028] write(3, "\300\217B\0\0\0\0\0\250&\2\0 \0\0\0\10\0\0\0\277&\2\0\302&\2\0\305&\2\0"..., 8192) = 8192
[pid 20028] write(3, "\323F\2\0\327F\2\0\2\0\0\0\4\0\0\0\344F\2\0\350F\2\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "P\330P\0\0\0\0\0\0\327P\0\0\0\0\0fmt.(*ss).error\0"..., 8192) = 8192
[pid 20028] write(3, "\0\0)\"\5!\224\1\202\1\r\201\1V2\0051U\"\25!\227\1R\5Q\220\1\0\0)\2"..., 8192) = 8192
[pid 20028] write(3, "\0(0\0\0\0\0\0\200=D\0\0\0\0\0\320\246\2\0\30\0\0\0\10\0\0\0\357\246\2\0"..., 8192) = 8192
[pid 20028] write(3, "\240\0\0\0\315\306\2\0\340\306\2\0\345\306\2\0\2\0\0\0\4\0\0\0\367\306\2\0\375\306\2\0"..., 8192) = 8192
[pid 20028] write(3, "\316\346\2\0\323\346\2\0\2\0\0\0\4\0\0\0\357\346\2\0\363\346\2\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, " \0\0\0\326\6\3\0\337\6\3\0\344\6\3\0\2\0\0\0\4\0\0\0\350\6\3\0\363\6\3\0"..., 8192) = 8192
[pid 20028] write(3, "\1\274\1\277\1\1\300\1\f\277\1\1\300\1\35\0\202\1\200\2\0\4\200\2\0\0\243\1R\5QX"..., 8192) = 8192
[pid 20028] write(3, "8\0\0\0\343F\3\0\361F\3\0\366F\3\0\2\0\0\0\4\0\0\0\24G\3\0\30G\3\0"..., 8192) = 8192
[pid 20028] write(3, "\340WF\0\0\0\0\0\320f\3\0\20\0\0\0\370\0\0\0\341f\3\0\32g\3\0\37g\3\0"..., 8192) = 8192
[pid 20028] write(3, "os.(*Process).setDone\0\0\2\31 &\37\1 \20\0"..., 8192) = 8192
[pid 20028] write(3, "\305\246\3\0\311\246\3\0\0\0\0\0\4\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\220\325P\0\0\0\0\0\220\324P\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "9\2\27\2:\2J\6<\0\0}\"\2!6B\5A\23\"C\0\0K\2j\2[\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0\336\1\300\4\0\216\27&\2\10\2\23\2M\4\21\2\227\1\4\26\2\r\f\177\t \2\24\10\5"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\0\0\0\0\0\10\331P\0\0\0\0\0\310\334P\0\0\0\0\0reflect."..., 8192) = 8192
[pid 20028] write(3, "\0\0\2\31`\223\1_\1`\23\0\340\1\300\1\0\246\1\300\1\0\0nbR\0\0n\2R\0"..., 8192) = 8192
[pid 20028] write(3, "\300\265P\0\0\0\0\0reflect.(*funcType).Impl"..., 8192) = 8192
[pid 20028] write(3, "\345\206\4\0\352\206\4\0\2\0\0\0\4\0\0\0\356\206\4\0\366\206\4\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\343\246\4\0\350\246\4\0\2\0\0\0\4\0\0\0\354\246\4\0\360\246\4\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\300\265P\0\0\0\0\0reflect.structType.uncom"..., 8192) = 8192
[pid 20028] write(3, "Buffer).Reset\0\0\2\31 \33\37\f\0\364\1@\0\230\1@\0\0@"..., 8192) = 8192
[pid 20028] write(3, "\240\1\5\2+\0\0\0\0\356H\0\0\0\0\0\330\6\5\0008\0\0\0@\0\0\0\346\6\5\0"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\0\0\0\0\0D\313P\0\0\0\0\0h\313P\0\0\0\0\0type..eq"..., 8192) = 8192
[pid 20028] write(3, "ir_unix.go\0\0/usr/local/go/src/pk"..., 3793) = 3793
[pid 20028] lseek(3, 1515520, SEEK_SET) = 1515520
[pid 20028] write(3, "\1\0\0\0\0\340\377\370\1\0\0\0\0\330\377\337\1\0\0\0( ( \1\0\0\0) ) "..., 8192) = 8192
[pid 20028] write(3, "\0\6\4\6\1\0\6\6\v\6\1\0\r\6\32\6\1\0\34\6\34\6\1\0\36\6\36\6\1\0 \6"..., 8192) = 8192
[pid 20028] write(3, "\1\0~\3~\3\1\0\205\3\205\3\1\0\207\3\207\3\1\0\211\5\211\5\1\0\f\6\f\6\1\0"..., 8192) = 8192
[pid 20028] write(3, "\275\324\1\0\303\324\1\0\1\0\0\0\305\324\1\0\5\325\1\0\1\0\0\0\7\325\1\0\n\325\1\0"..., 8192) = 8192
[pid 20028] write(3, "\1\0\0\0p\327\1\0\210\327\1\0\1\0\0\0\212\327\1\0\250\327\1\0\1\0\0\0\252\327\1\0"..., 8192) = 8192
[pid 20028] write(3, "\200\277W\0\0\0\0\0\v\0\0\0\0\0\0\0\v\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "execmem for this process.\n\0runti"..., 8192) = 8192
[pid 20028] write(3, "\1\0\0\0\0\0\0\0\252)X\0\0\0\0\0\3\0\0\0\0\0\0\0\277)X\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, " cgo callback before cgo call\n\0_"..., 2144) = 2144
[pid 20028] lseek(3, 1593344, SEEK_SET) = 1593344
[pid 20028] mmap(NULL, 2101248, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7ffbabd5b000
[pid 20028] write(3, "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\1\0\0\0\1\0\2\0"..., 8192) = 8192
[pid 20028] write(3, "\350\"W\0\0\0\0\0\f\0\0\0\0\0\0\0M\17\0\0\21\0\7\0@)W\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "8\0\0\0\0\0\0\0\360-\0\0\21\0\10\0`\321W\0\0\0\0\0008\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\227G\0\0\21\0\t\0 NX\0\0\0\0\0 \0\0\0\0\0\0\0\253G\0\0\21\0\2\0"..., 8192) = 8192
[pid 20028] write(3, "\20 @\0\0\0\0\0\20\1\0\0\0\0\0\0\353`\0\0\22\0\1\0 !@\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "@\0\0\0\0\0\0\0}{\0\0\22\0\1\0@6B\0\0\0\0\0\20\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\235\225\0\0\22\0\1\0\20>D\0\0\0\0\0 \0\0\0\0\0\0\0\275\225\0\0\22\0\1\0"..., 8192) = 8192
[pid 20028] write(3, "`\207F\0\0\0\0\0\320\0\0\0\0\0\0\0\234\262\0\0\22\0\1\0000\210F\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "@\0\0\0\0\0\0\0X\320\0\0\22\0\1\0pGH\0\0\0\0\0@\0\0\0\0\0\0\0"..., 8192) = 8192
[pid 20028] write(3, "\204\364\0\0\22\0\1\0\260\337H\0\0\0\0\0@\0\0\0\0\0\0\0\230\364\0\0\22\0\1\0"..., 2568) = 2568
[pid 20028] write(3, "\0shifts\0masks\0gclocals_reflectca"..., 64954) = 64954
[pid 20028] lseek(3, 0, SEEK_SET)       = 0
[pid 20028] write(3, "\177ELF\2\1\1\0\0\0\0\0\0\0\0\0\2\0>\0\1\0\0\0\260\215B\0\0\0\0\0"..., 1232) = 1232
[pid 20028] exit_group(0)               = ?
Process 20018 resumed
Process 20028 detached
[pid 20024] <... read resumed> "", 512) = 0
[pid 20024] --- SIGCHLD (Child exited) @ 0 (0) ---
[pid 20018] <... wait4 resumed> [{WIFEXITED(s) && WEXITSTATUS(s) == 0}], 0, {ru_utime={0, 92000}, ru_stime={0, 20000}, ...}) = 20028
[pid 20024] rt_sigreturn(0x11 <unfinished ...>
[pid 20018] futex(0xa80cb8, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... rt_sigreturn resumed> ) = 0
[pid 20021] <... futex resumed> )       = 0
[pid 20024] sched_yield( <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20024] <... sched_yield resumed> ) = 0
[pid 20018] <... futex resumed> )       = 1
[pid 20024] futex(0xa80c40, FUTEX_WAIT, 2, NULL <unfinished ...>
[pid 20018] futex(0xa80c40, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20018] <... futex resumed> )       = 0
[pid 20024] futex(0xa80c40, FUTEX_WAKE, 1) = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] close(4)                    = 0
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] pipe2([3, 4], O_CLOEXEC)    = 0
[pid 20018] clone( <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20021] select(0, NULL, NULL, NULL, {0, 20}Process 20029 attached (waiting for parent)
Process 20029 resumed (parent 20018 ready)
 <unfinished ...>
[pid 20018] <... clone resumed> child_stack=0, flags=|SIGCHLD) = 20029
[pid 20029] fcntl(0, F_SETFD, 0 <unfinished ...>
[pid 20018] close(4 <unfinished ...>
[pid 20029] <... fcntl resumed> )       = 0
[pid 20018] <... close resumed> )       = 0
[pid 20029] fcntl(1, F_SETFD, 0 <unfinished ...>
[pid 20018] read(3,  <unfinished ...>
[pid 20029] <... fcntl resumed> )       = 0
[pid 20029] fcntl(2, F_SETFD, 0 <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20029] <... fcntl resumed> )       = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20029] execve("/tmp/go-build628409689/command-line-arguments/_obj/exe/childProcess", ["/tmp/go-build628409689/command-l"...], [/* 60 vars */] <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20021] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20021] <... futex resumed> )       = 1
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20029] <... execve resumed> )      = 0
[pid 20018] <... read resumed> "", 8)   = 0
[pid 20018] close(3 <unfinished ...>
[pid 20029] arch_prctl(ARCH_SET_FS, 0x584f70 <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20029] <... arch_prctl resumed> )  = 0
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20018] <... close resumed> )       = 0
[pid 20029] sched_getaffinity(0, 128, {f, 0, 0, 0}) = 32
[pid 20018] wait4(20029, Process 20018 suspended
 <unfinished ...>
[pid 20029] mmap(0xc000000000, 65536, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0 <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20029] <... mmap resumed> )        = 0xc000000000
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20029] munmap(0xc000000000, 65536) = 0
[pid 20029] mmap(NULL, 262144, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7ffba2bf9000
[pid 20029] mmap(0xc208000000, 1048576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0 <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20029] <... mmap resumed> )        = 0xc208000000
[pid 20021] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20029] mmap(0xc207ff0000, 65536, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0 <unfinished ...>
[pid 20024] <... futex resumed> )       = 0
[pid 20029] <... mmap resumed> )        = 0xc207ff0000
[pid 20024] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20029] mmap(0xc000000000, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0 <unfinished ...>
[pid 20021] <... futex resumed> )       = 1
[pid 20029] <... mmap resumed> )        = 0xc000000000
[pid 20021] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20029] mmap(NULL, 65536, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7ffba2be9000
[pid 20029] mmap(NULL, 1439992, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7ffba2a89000
[pid 20029] mmap(NULL, 131072, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0 <unfinished ...>
[pid 20021] <... select resumed> )      = 0 (Timeout)
[pid 20029] <... mmap resumed> )        = 0x7ffba2a69000
[pid 20021] futex(0xa80cb8, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20029] sigaltstack({ss_sp=0xc208006000, ss_flags=0, ss_size=32768}, NULL) = 0
[pid 20029] rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
[pid 20029] rt_sigaction(SIGHUP, NULL, {SIG_DFL, [], 0}, 8) = 0
[pid 20029] rt_sigaction(SIGHUP, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGINT, NULL, {SIG_DFL, [], 0}, 8) = 0
[pid 20029] rt_sigaction(SIGINT, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGQUIT, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGILL, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGTRAP, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGABRT, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGBUS, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGFPE, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGUSR1, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGSEGV, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGUSR2, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGPIPE, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGALRM, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGTERM, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGSTKFLT, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGCHLD, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGURG, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGXCPU, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGXFSZ, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGVTALRM, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGPROF, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGWINCH, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGIO, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGPWR, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGSYS, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_2, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_3, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_4, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_5, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_6, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_7, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_8, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_9, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_10, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_11, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_12, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_13, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_14, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_15, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_16, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_17, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_18, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_19, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_20, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_21, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_22, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_23, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_24, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_25, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_26, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_27, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_28, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_29, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_30, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_31, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigaction(SIGRT_32, {0x429010, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x429080}, NULL, 8) = 0
[pid 20029] rt_sigprocmask(SIG_SETMASK, ~[], [], 8) = 0
[pid 20029] clone(Process 20030 attached (waiting for parent)
Process 20030 resumed (parent 20029 ready)
child_stack=0x7ffba2a88fa8, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD) = 20030
[pid 20030] gettid( <unfinished ...>
[pid 20029] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20030] <... gettid resumed> )      = 20030
[pid 20029] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20030] arch_prctl(ARCH_SET_FS, 0xc208016070) = 0
[pid 20030] sigaltstack({ss_sp=0xc20801a000, ss_flags=0, ss_size=32768} <unfinished ...>
[pid 20029] mmap(NULL, 1048576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0 <unfinished ...>
[pid 20030] <... sigaltstack resumed> , NULL) = 0
[pid 20029] <... mmap resumed> )        = 0x7ffba2969000
[pid 20030] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20029] stat("/home/qboxtest/.rvm/gems/ruby-2.1.3/bin/go",  <unfinished ...>
[pid 20030] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20029] <... stat resumed> 0xc208040000) = -1 ENOENT (No such file or directory)
[pid 20030] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20029] stat("/home/qboxtest/.rvm/gems/ruby-2.1.3@global/bin/go", 0xc208040090) = -1 ENOENT (No such file or directory)
[pid 20029] stat("/home/qboxtest/.rvm/rubies/ruby-2.1.3/bin/go", 0xc208040120) = -1 ENOENT (No such file or directory)
[pid 20029] stat("/home/qboxtest/bin/go", 0xc2080401b0) = -1 ENOENT (No such file or directory)
[pid 20030] <... select resumed> )      = 0 (Timeout)
[pid 20029] stat("/usr/lib/lightdm/lightdm/go",  <unfinished ...>
[pid 20030] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20029] <... stat resumed> 0xc208040240) = -1 ENOENT (No such file or directory)
[pid 20029] stat("/usr/local/sbin/go", 0xc2080402d0) = -1 ENOENT (No such file or directory)
[pid 20029] stat("/usr/local/bin/go", 0xc208040360) = -1 ENOENT (No such file or directory)
[pid 20029] stat("/usr/sbin/go", 0xc2080403f0) = -1 ENOENT (No such file or directory)
[pid 20029] stat("/usr/bin/go",  <unfinished ...>
[pid 20030] <... select resumed> )      = 0 (Timeout)
[pid 20029] <... stat resumed> 0xc208040480) = -1 ENOENT (No such file or directory)
[pid 20030] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20029] stat("/sbin/go", 0xc208040510) = -1 ENOENT (No such file or directory)
[pid 20029] stat("/bin/go", 0xc2080405a0) = -1 ENOENT (No such file or directory)
[pid 20029] stat("/usr/games/go", 0xc208040630) = -1 ENOENT (No such file or directory)
[pid 20029] stat("/home/qboxtest/.rvm/bin/go", 0xc2080406c0) = -1 ENOENT (No such file or directory)
[pid 20030] <... select resumed> )      = 0 (Timeout)
[pid 20029] stat("/usr/local/go/bin/go",  <unfinished ...>
[pid 20030] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20029] <... stat resumed> {st_mode=S_IFREG|0755, st_size=9349952, ...}) = 0
[pid 20029] open("/dev/null", O_RDONLY|O_CLOEXEC) = 3
[pid 20029] open("/dev/null", O_WRONLY|O_CLOEXEC) = 4
[pid 20030] <... select resumed> )      = 0 (Timeout)
[pid 20029] open("/dev/null", O_WRONLY|O_CLOEXEC <unfinished ...>
[pid 20030] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20029] <... open resumed> )        = 5
[pid 20029] pipe2([6, 7], O_CLOEXEC)    = 0
[pid 20029] clone( <unfinished ...>
[pid 20030] <... select resumed> )      = 0 (Timeout)
Process 20031 attached (waiting for parent)
[pid 20030] select(0, NULL, NULL, NULL, {0, 20}Process 20031 resumed (parent 20029 ready)
 <unfinished ...>
[pid 20029] <... clone resumed> child_stack=0, flags=|SIGCHLD) = 20031
[pid 20031] dup2(3, 0)                  = 0
[pid 20029] close(7 <unfinished ...>
[pid 20031] dup2(4, 1 <unfinished ...>
[pid 20029] <... close resumed> )       = 0
[pid 20031] <... dup2 resumed> )        = 1
[pid 20029] read(6,  <unfinished ...>
[pid 20031] dup2(5, 2)                  = 2
[pid 20031] execve("/usr/local/go/bin/go", ["go", "run", "server2.go"], [/* 60 vars */] <unfinished ...>
[pid 20030] <... select resumed> )      = 0 (Timeout)
[pid 20030] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] <... execve resumed> )      = 0
[pid 20029] <... read resumed> "", 8)   = 0
[pid 20029] close(6)                    = 0
[pid 20031] brk(0 <unfinished ...>
[pid 20029] close(3 <unfinished ...>
[pid 20031] <... brk resumed> )         = 0x1990000
[pid 20030] <... select resumed> )      = 0 (Timeout)
[pid 20029] <... close resumed> )       = 0
[pid 20031] access("/etc/ld.so.nohwcap", F_OK <unfinished ...>
[pid 20030] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] <... access resumed> )      = -1 ENOENT (No such file or directory)
[pid 20029] close(4 <unfinished ...>
[pid 20031] mmap(NULL, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0 <unfinished ...>
[pid 20029] <... close resumed> )       = 0
[pid 20031] <... mmap resumed> )        = 0x7f319cdb8000
[pid 20029] close(5 <unfinished ...>
[pid 20031] access("/etc/ld.so.preload", R_OK <unfinished ...>
[pid 20029] <... close resumed> )       = 0
[pid 20031] <... access resumed> )      = -1 ENOENT (No such file or directory)
[pid 20029] write(1, "child--------- 20031\n", 21 <unfinished ...>
[pid 20031] open("/usr/local/lib/tls/x86_64/libpthread.so.0", O_RDONLY|O_CLOEXEC <unfinished ...>
[pid 20030] <... select resumed> )      = 0 (Timeout)
[pid 20031] <... open resumed> )        = -1 ENOENT (No such file or directory)
[pid 20030] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] stat("/usr/local/lib/tls/x86_64",  <unfinished ...>
[pid 20029] <... write resumed> )       = 21
[pid 20031] <... stat resumed> 0x7fff7dc0c600) = -1 ENOENT (No such file or directory)
[pid 20029] rt_sigprocmask(SIG_SETMASK, ~[],  <unfinished ...>
[pid 20031] open("/usr/local/lib/tls/libpthread.so.0", O_RDONLY|O_CLOEXEC <unfinished ...>
[pid 20029] <... rt_sigprocmask resumed> [], 8) = 0
[pid 20031] <... open resumed> )        = -1 ENOENT (No such file or directory)
[pid 20030] <... select resumed> )      = 0 (Timeout)
[pid 20031] stat("/usr/local/lib/tls",  <unfinished ...>
[pid 20030] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] <... stat resumed> 0x7fff7dc0c600) = -1 ENOENT (No such file or directory)
[pid 20029] clone( <unfinished ...>
[pid 20031] open("/usr/local/lib/x86_64/libpthread.so.0", O_RDONLY|O_CLOEXECProcess 20032 attached (waiting for parent)
) = -1 ENOENT (No such file or directory)
Process 20032 resumed (parent 20029 ready)
[pid 20029] <... clone resumed> child_stack=0x7ffba2a7efa8, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD) = 20032
[pid 20032] gettid( <unfinished ...>
[pid 20031] stat("/usr/local/lib/x86_64",  <unfinished ...>
[pid 20032] <... gettid resumed> )      = 20032
[pid 20031] <... stat resumed> 0x7fff7dc0c600) = -1 ENOENT (No such file or directory)
[pid 20032] arch_prctl(ARCH_SET_FS, 0xc208017270 <unfinished ...>
[pid 20031] open("/usr/local/lib/libpthread.so.0", O_RDONLY|O_CLOEXEC <unfinished ...>
[pid 20032] <... arch_prctl resumed> )  = 0
[pid 20031] <... open resumed> )        = -1 ENOENT (No such file or directory)
[pid 20032] sigaltstack({ss_sp=0xc208056000, ss_flags=0, ss_size=32768} <unfinished ...>
[pid 20031] stat("/usr/local/lib",  <unfinished ...>
[pid 20032] <... sigaltstack resumed> , NULL) = 0
[pid 20031] <... stat resumed> {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20032] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20031] open("/etc/ld.so.cache", O_RDONLY|O_CLOEXEC <unfinished ...>
[pid 20032] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20031] <... open resumed> )        = 3
[pid 20030] <... select resumed> )      = 0 (Timeout)
[pid 20031] fstat(3,  <unfinished ...>
[pid 20030] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20032] mmap(NULL, 131072, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0 <unfinished ...>
[pid 20031] <... fstat resumed> {st_mode=S_IFREG|0644, st_size=124566, ...}) = 0
[pid 20032] <... mmap resumed> )        = 0x7ffba2949000
[pid 20031] mmap(NULL, 124566, PROT_READ, MAP_PRIVATE, 3, 0 <unfinished ...>
[pid 20032] rt_sigprocmask(SIG_SETMASK, ~[],  <unfinished ...>
[pid 20031] <... mmap resumed> )        = 0x7f319cd99000
[pid 20032] <... rt_sigprocmask resumed> [], 8) = 0
[pid 20031] close(3 <unfinished ...>
[pid 20032] clone( <unfinished ...>
[pid 20031] <... close resumed> )       = 0
[pid 20029] rt_sigprocmask(SIG_SETMASK, [], Process 20033 attached
 <unfinished ...>
[pid 20032] <... clone resumed> child_stack=0x7ffba294afa8, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD) = 20033
[pid 20032] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20033] gettid( <unfinished ...>
[pid 20032] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20033] <... gettid resumed> )      = 20033
[pid 20032] futex(0x584ef8, FUTEX_WAIT, 0, {39, 999712088} <unfinished ...>
[pid 20033] arch_prctl(ARCH_SET_FS, 0xc2080176f0 <unfinished ...>
[pid 20031] access("/etc/ld.so.nohwcap", F_OK <unfinished ...>
[pid 20033] <... arch_prctl resumed> )  = 0
[pid 20031] <... access resumed> )      = -1 ENOENT (No such file or directory)
[pid 20033] sigaltstack({ss_sp=0xc20805e000, ss_flags=0, ss_size=32768} <unfinished ...>
[pid 20030] <... select resumed> )      = 0 (Timeout)
[pid 20033] <... sigaltstack resumed> , NULL) = 0
[pid 20031] open("/lib/x86_64-linux-gnu/libpthread.so.0", O_RDONLY|O_CLOEXEC <unfinished ...>
[pid 20033] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20031] <... open resumed> )        = 3
[pid 20033] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20031] read(3,  <unfinished ...>
[pid 20033] futex(0xc208017770, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... read resumed> "\177ELF\2\1\1\0\0\0\0\0\0\0\0\0\3\0>\0\1\0\0\0\200l\0\0\0\0\0\0"..., 832) = 832
[pid 20030] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] fstat(3, {st_mode=S_IFREG|0755, st_size=135366, ...}) = 0
[pid 20029] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20031] mmap(NULL, 2212904, PROT_READ|PROT_EXEC, MAP_PRIVATE|MAP_DENYWRITE, 3, 0 <unfinished ...>
[pid 20029] futex(0x7ffba2a86f60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20031] <... mmap resumed> )        = 0x7f319c97b000
[pid 20031] mprotect(0x7f319c993000, 2093056, PROT_NONE) = 0
[pid 20030] <... select resumed> )      = 0 (Timeout)
[pid 20031] mmap(0x7f319cb92000, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x17000 <unfinished ...>
[pid 20030] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... mmap resumed> )        = 0x7f319cb92000
[pid 20031] mmap(0x7f319cb94000, 13352, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) = 0x7f319cb94000
[pid 20031] close(3)                    = 0
[pid 20031] open("/usr/local/lib/libc.so.6", O_RDONLY|O_CLOEXEC) = -1 ENOENT (No such file or directory)
[pid 20031] access("/etc/ld.so.nohwcap", F_OK) = -1 ENOENT (No such file or directory)
[pid 20031] open("/lib/x86_64-linux-gnu/libc.so.6", O_RDONLY|O_CLOEXEC) = 3
[pid 20031] read(3, "\177ELF\2\1\1\0\0\0\0\0\0\0\0\0\3\0>\0\1\0\0\0\200\30\2\0\0\0\0\0"..., 832) = 832
[pid 20031] fstat(3, {st_mode=S_IFREG|0755, st_size=1815224, ...}) = 0
[pid 20031] mmap(NULL, 3929304, PROT_READ|PROT_EXEC, MAP_PRIVATE|MAP_DENYWRITE, 3, 0) = 0x7f319c5bb000
[pid 20031] mprotect(0x7f319c770000, 2097152, PROT_NONE) = 0
[pid 20031] mmap(0x7f319c970000, 24576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x1b5000) = 0x7f319c970000
[pid 20031] mmap(0x7f319c976000, 17624, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) = 0x7f319c976000
[pid 20031] close(3)                    = 0
[pid 20031] mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f319cd98000
[pid 20031] mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f319cd97000
[pid 20031] mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f319cd96000
[pid 20031] arch_prctl(ARCH_SET_FS, 0x7f319cd97700) = 0
[pid 20031] mprotect(0x7f319c970000, 16384, PROT_READ) = 0
[pid 20031] mprotect(0x7f319cb92000, 4096, PROT_READ) = 0
[pid 20031] mprotect(0x7f319cdba000, 4096, PROT_READ) = 0
[pid 20031] munmap(0x7f319cd99000, 124566) = 0
[pid 20031] set_tid_address(0x7f319cd979d0) = 20031
[pid 20031] set_robust_list(0x7f319cd979e0, 0x18) = 0
[pid 20031] futex(0x7fff7dc0cefc, FUTEX_WAIT_BITSET_PRIVATE|FUTEX_CLOCK_REALTIME, 1, NULL, 7f319cd97700) = -1 EAGAIN (Resource temporarily unavailable)
[pid 20031] rt_sigaction(SIGRTMIN, {0x7f319c981750, [], SA_RESTORER|SA_SIGINFO, 0x7f319c98acb0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_1, {0x7f319c9817e0, [], SA_RESTORER|SA_RESTART|SA_SIGINFO, 0x7f319c98acb0}, NULL, 8) = 0
[pid 20031] rt_sigprocmask(SIG_UNBLOCK, [RTMIN RT_1], NULL, 8) = 0
[pid 20031] getrlimit(RLIMIT_STACK, {rlim_cur=8192*1024, rlim_max=RLIM_INFINITY}) = 0
[pid 20031] sched_getaffinity(0, 128, {f, 0, 0, 0}) = 32
[pid 20031] mmap(0xc000000000, 65536, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0xc000000000
[pid 20031] munmap(0xc000000000, 65536) = 0
[pid 20031] mmap(NULL, 262144, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f319cd56000
[pid 20031] mmap(0xc208000000, 1048576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0xc208000000
[pid 20031] mmap(0xc207ff0000, 65536, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0xc207ff0000
[pid 20031] mmap(0xc000000000, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0xc000000000
[pid 20031] mmap(NULL, 65536, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f319cda8000
[pid 20031] mmap(NULL, 1439992, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f319cbf6000
[pid 20031] mmap(NULL, 131072, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f319cbd6000
[pid 20031] sigaltstack({ss_sp=0xc208006000, ss_flags=0, ss_size=32768}, NULL) = 0
[pid 20031] rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
[pid 20031] rt_sigaction(SIGHUP, NULL, {SIG_DFL, [], 0}, 8) = 0
[pid 20031] rt_sigaction(SIGHUP, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGINT, NULL, {SIG_DFL, [], 0}, 8) = 0
[pid 20031] rt_sigaction(SIGINT, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGQUIT, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGILL, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGTRAP, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGABRT, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGBUS, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGFPE, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGUSR1, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGSEGV, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGUSR2, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGPIPE, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGALRM, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGTERM, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGSTKFLT, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGCHLD, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGURG, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGXCPU, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGXFSZ, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGVTALRM, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGPROF, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGWINCH, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGIO, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGPWR, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGSYS, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_2, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_3, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_4, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_5, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_6, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_7, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_8, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_9, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_10, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_11, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_12, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_13, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_14, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_15, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_16, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_17, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_18, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_19, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_20, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_21, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_22, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_23, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_24, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_25, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_26, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_27, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_28, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_29, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_30, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_31, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] rt_sigaction(SIGRT_32, {0x487460, ~[], SA_RESTORER|SA_STACK|SA_RESTART|SA_SIGINFO, 0x4874d0}, NULL, 8) = 0
[pid 20031] brk(0)                      = 0x1990000
[pid 20031] brk(0x19b1000)              = 0x19b1000
[pid 20031] rt_sigprocmask(SIG_SETMASK, ~[RTMIN RT_1], [], 8) = 0
[pid 20031] mmap(NULL, 8392704, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_STACK, -1, 0) = 0x7f319bdba000
[pid 20031] mprotect(0x7f319bdba000, 4096, PROT_NONE) = 0
[pid 20031] clone(Process 20034 attached (waiting for parent)
Process 20034 resumed (parent 20031 ready)
child_stack=0x7f319c5b9ff0, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLS|CLONE_PARENT_SETTID|CLONE_CHILD_CLEARTID, parent_tidptr=0x7f319c5ba9d0, tls=0x7f319c5ba700, child_tidptr=0x7f319c5ba9d0) = 20034
[pid 20034] set_robust_list(0x7f319c5ba9e0, 0x18 <unfinished ...>
[pid 20031] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20034] <... set_robust_list resumed> ) = 0
[pid 20031] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20034] sigaltstack({ss_sp=0xc20801c000, ss_flags=0, ss_size=32768} <unfinished ...>
[pid 20031] mmap(NULL, 1048576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0 <unfinished ...>
[pid 20034] <... sigaltstack resumed> , NULL) = 0
[pid 20031] <... mmap resumed> )        = 0x7f319bcba000
[pid 20034] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20031] open("/proc/sys/net/core/somaxconn", O_RDONLY|O_CLOEXEC <unfinished ...>
[pid 20034] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20031] <... open resumed> )        = 3
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] read(3, "128\n", 4096)      = 4
[pid 20031] read(3, "", 4092)           = 0
[pid 20031] close(3)                    = 0
[pid 20031] socket(PF_INET, SOCK_STREAM, IPPROTO_TCP <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] <... socket resumed> )      = 3
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] close(3)                    = 0
[pid 20031] socket(PF_INET6, SOCK_STREAM, IPPROTO_TCP) = 3
[pid 20031] setsockopt(3, SOL_IPV6, IPV6_V6ONLY, [1], 4) = 0
[pid 20031] bind(3, {sa_family=AF_INET6, sin6_port=htons(0), inet_pton(AF_INET6, "::1", &sin6_addr), sin6_flowinfo=0, sin6_scope_id=0}, 28 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] <... bind resumed> )        = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] socket(PF_INET6, SOCK_STREAM, IPPROTO_TCP) = 4
[pid 20031] setsockopt(4, SOL_IPV6, IPV6_V6ONLY, [0], 4 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] <... setsockopt resumed> )  = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] bind(4, {sa_family=AF_INET6, sin6_port=htons(0), inet_pton(AF_INET6, "::ffff:127.0.0.1", &sin6_addr), sin6_flowinfo=0, sin6_scope_id=0}, 28) = 0
[pid 20031] close(4)                    = 0
[pid 20031] close(3)                    = 0
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20034] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20034] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] stat("/home/qboxtest/.rvm/gems/ruby-2.1.3/bin/gccgo",  <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] <... stat resumed> 0xc208044240) = -1 ENOENT (No such file or directory)
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] stat("/home/qboxtest/.rvm/gems/ruby-2.1.3@global/bin/gccgo", 0xc2080442d0) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/.rvm/rubies/ruby-2.1.3/bin/gccgo", 0xc208044360) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/bin/gccgo", 0xc2080443f0) = -1 ENOENT (No such file or directory)
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] stat("/usr/lib/lightdm/lightdm/gccgo",  <unfinished ...>
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] <... stat resumed> 0xc208044480) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/usr/local/sbin/gccgo", 0xc208044510) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/usr/local/bin/gccgo", 0xc2080445a0) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/usr/sbin/gccgo", 0xc208044630) = -1 ENOENT (No such file or directory)
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] stat("/usr/bin/gccgo",  <unfinished ...>
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] <... stat resumed> 0xc2080446c0) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/sbin/gccgo", 0xc208044750) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/bin/gccgo", 0xc2080447e0) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/usr/games/gccgo", 0xc208044870) = -1 ENOENT (No such file or directory)
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] stat("/home/qboxtest/.rvm/bin/gccgo",  <unfinished ...>
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] <... stat resumed> 0xc208044900) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/usr/local/go/bin/gccgo", 0xc208044990) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/ffmpeg/ffmpeg-static/target/bin/gccgo", 0xc208044a20) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/base/qiniu/bin/gccgo", 0xc208044ab0) = -1 ENOENT (No such file or directory)
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] stat("/home/qboxtest/qbox/base/biz/bin/gccgo",  <unfinished ...>
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] <... stat resumed> 0xc208044b40) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/base/cgo/bin/gccgo", 0xc208044bd0) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/base/com/bin/gccgo", 0xc208044c60) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/base/mockacc/bin/gccgo", 0xc208044cf0) = -1 ENOENT (No such file or directory)
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] stat("/home/qboxtest/qbox/image/fop_ncgo/bin/gccgo",  <unfinished ...>
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] <... stat resumed> 0xc208044d80) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/image/fop_cgo/bin/gccgo", 0xc208044e10) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/ffmpeg/fileop/bin/gccgo", 0xc208044ea0) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/ffmpeg/fop/bin/gccgo", 0xc208044f30) = -1 ENOENT (No such file or directory)
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] stat("/home/qboxtest/qbox/fileop/fileop/bin/gccgo",  <unfinished ...>
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] <... stat resumed> 0xc208044fc0) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/io/cgo_deps/bin/gccgo", 0xc208045050) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/io/io/bin/gccgo",  <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] <... stat resumed> 0xc2080450e0) = -1 ENOENT (No such file or directory)
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] stat("/home/qboxtest/qbox/service/service/bin/gccgo", 0xc208045170) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/service/service_cgo/bin/gccgo", 0xc208045200) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/fop/bin/gccgo", 0xc208045290) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/devtools/bin/gccgo",  <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] <... stat resumed> 0xc208045320) = -1 ENOENT (No such file or directory)
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] stat("/home/qboxtest/qbox/rs/rs/bin/gccgo", 0xc2080453b0) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/nio/bin/gccgo", 0xc208045440) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/fileop/fileop/bin/gccgo", 0xc2080454d0) = -1 ENOENT (No such file or directory)
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] stat("/home/qboxtest/qbox/biz/bin/gccgo",  <unfinished ...>
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] <... stat resumed> 0xc208045560) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/office/fop_ncgo/bin/gccgo", 0xc2080455f0) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/office/fop_cgo/bin/gccgo", 0xc208045680) = -1 ENOENT (No such file or directory)
[pid 20031] stat("/home/qboxtest/qbox/api.qiniu.com/bin/gccgo",  <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] <... stat resumed> 0xc208045710) = -1 ENOENT (No such file or directory)
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] stat("/home/qboxtest/.rvm/bin/gccgo", 0xc2080457a0) = -1 ENOENT (No such file or directory)
[pid 20031] getcwd("/home/qboxtest/Desktop/qiniu", 4096) = 29
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20034] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20034] futex(0xa80cb8, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] futex(0xa80cb8, FUTEX_WAKE, 1) = 1
[pid 20034] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] rt_sigprocmask(SIG_SETMASK, ~[RTMIN RT_1], [], 8) = 0
[pid 20031] mmap(NULL, 8392704, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_STACK, -1, 0) = 0x7f319b4b9000
[pid 20031] mprotect(0x7f319b4b9000, 4096, PROT_NONE <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] <... mprotect resumed> )    = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] clone(Process 20035 attached
child_stack=0x7f319bcb8ff0, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLS|CLONE_PARENT_SETTID|CLONE_CHILD_CLEARTID, parent_tidptr=0x7f319bcb99d0, tls=0x7f319bcb9700, child_tidptr=0x7f319bcb99d0) = 20035
[pid 20031] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20035] set_robust_list(0x7f319bcb99e0, 0x18 <unfinished ...>
[pid 20031] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20035] <... set_robust_list resumed> ) = 0
[pid 20031] futex(0xa81b10, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20035] sigaltstack({ss_sp=0xc208076000, ss_flags=0, ss_size=32768}, NULL) = 0
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20035] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20035] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20035] open("/sys/devices/system/cpu/online", O_RDONLY|O_CLOEXEC) = 3
[pid 20035] read(3, "0-3\n", 8192)      = 4
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20035] close(3 <unfinished ...>
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20035] <... close resumed> )       = 0
[pid 20035] mmap(NULL, 134217728, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_NORESERVE, -1, 0) = 0x7f31934b9000
[pid 20035] munmap(0x7f31934b9000, 11825152) = 0
[pid 20035] munmap(0x7f3198000000, 55283712) = 0
[pid 20035] mprotect(0x7f3194000000, 135168, PROT_READ|PROT_WRITE <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20035] <... mprotect resumed> )    = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20035] rt_sigprocmask(SIG_SETMASK, ~[RTMIN RT_1], [], 8) = 0
[pid 20035] mmap(NULL, 8392704, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_STACK, -1, 0) = 0x7f319acb8000
[pid 20035] mprotect(0x7f319acb8000, 4096, PROT_NONE) = 0
[pid 20035] clone( <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
Process 20036 attached (waiting for parent)
Process 20036 resumed (parent 20035 ready)
[pid 20035] <... clone resumed> child_stack=0x7f319b4b7ff0, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLS|CLONE_PARENT_SETTID|CLONE_CHILD_CLEARTID, parent_tidptr=0x7f319b4b89d0, tls=0x7f319b4b8700, child_tidptr=0x7f319b4b89d0) = 20036
[pid 20036] set_robust_list(0x7f319b4b89e0, 0x18 <unfinished ...>
[pid 20035] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20036] <... set_robust_list resumed> ) = 0
[pid 20035] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20036] sigaltstack({ss_sp=0xc20808e000, ss_flags=0, ss_size=32768} <unfinished ...>
[pid 20035] futex(0x7f319cbf5f60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20036] <... sigaltstack resumed> , NULL) = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20036] rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
[pid 20036] mmap(NULL, 134217728, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_NORESERVE, -1, 0) = 0x7f318c000000
[pid 20036] munmap(0x7f3190000000, 67108864) = 0
[pid 20036] mprotect(0x7f318c000000, 135168, PROT_READ|PROT_WRITE <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20036] <... mprotect resumed> )    = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20036] rt_sigprocmask(SIG_SETMASK, ~[RTMIN RT_1], [], 8) = 0
[pid 20036] mmap(NULL, 8392704, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_STACK, -1, 0) = 0x7f319a4b7000
[pid 20036] mprotect(0x7f319a4b7000, 4096, PROT_NONE) = 0
[pid 20036] clone( <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
Process 20037 attached (waiting for parent)
Process 20037 resumed (parent 20036 ready)
[pid 20036] <... clone resumed> child_stack=0x7f319acb6ff0, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLS|CLONE_PARENT_SETTID|CLONE_CHILD_CLEARTID, parent_tidptr=0x7f319acb79d0, tls=0x7f319acb7700, child_tidptr=0x7f319acb79d0) = 20037
[pid 20037] set_robust_list(0x7f319acb79e0, 0x18 <unfinished ...>
[pid 20036] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20037] <... set_robust_list resumed> ) = 0
[pid 20036] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20037] sigaltstack({ss_sp=0xc208096000, ss_flags=0, ss_size=32768} <unfinished ...>
[pid 20036] futex(0xa808c0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] <... sigaltstack resumed> , NULL) = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1) = 1
[pid 20031] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20031] <... futex resumed> )       = 0
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1) = 0
[pid 20031] rt_sigprocmask(SIG_SETMASK, ~[RTMIN RT_1],  <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] <... rt_sigprocmask resumed> [], 8) = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] mmap(NULL, 8392704, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_STACK, -1, 0) = 0x7f3199cb6000
[pid 20031] mprotect(0x7f3199cb6000, 4096, PROT_NONE) = 0
[pid 20031] clone(Process 20038 attached (waiting for parent)
Process 20038 resumed (parent 20031 ready)
child_stack=0x7f319a4b5ff0, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLS|CLONE_PARENT_SETTID|CLONE_CHILD_CLEARTID, parent_tidptr=0x7f319a4b69d0, tls=0x7f319a4b6700, child_tidptr=0x7f319a4b69d0) = 20038
[pid 20038] set_robust_list(0x7f319a4b69e0, 0x18 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... set_robust_list resumed> ) = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] sigaltstack({ss_sp=0xc2080a0000, ss_flags=0, ss_size=32768} <unfinished ...>
[pid 20031] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20038] <... sigaltstack resumed> , NULL) = 0
[pid 20031] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20038] rt_sigprocmask(SIG_SETMASK, [],  <unfinished ...>
[pid 20031] futex(0xa81b10, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] <... rt_sigprocmask resumed> NULL, 8) = 0
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 1
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 0
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xa81b10, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1) = 1
[pid 20038] <... futex resumed> )       = 0
[pid 20031] futex(0xa81b10, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] <... futex resumed> )       = 1
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] futex(0xa81b10, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1) = 1
[pid 20031] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20031] <... futex resumed> )       = 0
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20038] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 40} <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 80} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20038] <... futex resumed> )       = 0
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] select(0, NULL, NULL, NULL, {0, 160} <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20031] <... futex resumed> )       = 0
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20038] <... futex resumed> )       = 0
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20038] <... futex resumed> )       = 0
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] select(0, NULL, NULL, NULL, {0, 320} <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20031] <... futex resumed> )       = 0
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20038] <... futex resumed> )       = 0
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20031] <... futex resumed> )       = 0
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20037] <... futex resumed> )       = 0
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] <... futex resumed> )       = 0
[pid 20034] select(0, NULL, NULL, NULL, {0, 640} <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20038] <... futex resumed> )       = 0
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20038] <... futex resumed> )       = 0
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20038] <... futex resumed> )       = 0
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20038] <... futex resumed> )       = 0
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20038] <... futex resumed> )       = 0
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20038] <... futex resumed> )       = 0
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20038] <... futex resumed> )       = 0
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] select(0, NULL, NULL, NULL, {0, 1280} <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20031] <... futex resumed> )       = 0
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20038] <... futex resumed> )       = 0
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20038] <... futex resumed> )       = 0
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20038] <... futex resumed> )       = 0
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] <... futex resumed> )       = 0
[pid 20038] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20037] <... futex resumed> )       = 0
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20037] <... futex resumed> )       = 0
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20037] <... futex resumed> )       = 0
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20037] <... futex resumed> )       = 0
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20037] <... futex resumed> )       = 0
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20037] <... futex resumed> )       = 0
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20037] <... futex resumed> )       = 0
[pid 20031] futex(0xc20809e0f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20038] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] futex(0xc208019bf0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20038] <... futex resumed> )       = 0
[pid 20037] <... futex resumed> )       = 0
[pid 20038] futex(0xc20809e0f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20037] futex(0xa81b10, FUTEX_WAKE, 1 <unfinished ...>
[pid 20031] <... futex resumed> )       = 1
[pid 20037] <... futex resumed> )       = 0
[pid 20037] futex(0xc208019bf0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20034] futex(0xa80cb8, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20031] futex(0xa80cb8, FUTEX_WAKE, 1 <unfinished ...>
[pid 20034] <... futex resumed> )       = 0
[pid 20031] <... futex resumed> )       = 1
[pid 20034] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20034] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20034] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20034] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] stat("/usr/local/go", {st_mode=S_IFDIR|0755, st_size=4096, ...}) = 0
[pid 20031] getpid()                    = 20031
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] mkdir("/tmp/go-build168288847", 0700 <unfinished ...>
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] <... mkdir resumed> )       = 0
[pid 20031] stat("server2.go", 0xc208044480) = -1 ENOENT (No such file or directory)
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20031] write(2, "stat server2.go: no such file or"..., 43 <unfinished ...>
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] <... write resumed> )       = 43
[pid 20031] unlink("/tmp/go-build168288847") = -1 EISDIR (Is a directory)
[pid 20031] rmdir("/tmp/go-build168288847" <unfinished ...>
[pid 20034] <... select resumed> )      = 0 (Timeout)
[pid 20034] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20031] <... rmdir resumed> )       = 0
[pid 20031] exit_group(1)               = ?
Process 20031 detached
[pid 20038] <... ???? resumed> )        = ? <unavailable>
[pid 20029] <... futex resumed> )       = ? ERESTART_RESTARTBLOCK (To be restarted)
[pid 20029] --- SIGCHLD (Child exited) @ 0 (0) ---
[pid 20029] rt_sigreturn(0x11)          = -1 EINTR (Interrupted system call)
[pid 20029] futex(0x7ffba2a86f60, FUTEX_WAIT, 0, {59, 984763686} <unfinished ...>
[pid 20020] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20020] futex(0x585158, FUTEX_WAKE, 1) = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20020] write(1, "1--------- 20018\n", 17) = 17
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20020] kill(20018, SIGKILL <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20026] +++ killed by SIGKILL +++
PANIC: handle_group_exit: 20026 leader 20018
[pid 20025] +++ killed by SIGKILL +++
PANIC: handle_group_exit: 20025 leader 20018
[pid 20024] +++ killed by SIGKILL +++
PANIC: handle_group_exit: 20024 leader 20018
[pid 20023] +++ killed by SIGKILL +++
PANIC: handle_group_exit: 20023 leader 20018
[pid 20022] +++ killed by SIGKILL +++
PANIC: handle_group_exit: 20022 leader 20018
[pid 20021] +++ killed by SIGKILL +++
PANIC: handle_group_exit: 20021 leader 20018
[pid 20020] <... kill resumed> )        = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20020] --- SIGCHLD (Child exited) @ 0 (0) ---
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20020] rt_sigreturn(0x11 <unfinished ...>
[pid 20016] <... futex resumed> )       = ? ERESTARTSYS (To be restarted)
[pid 20020] <... rt_sigreturn resumed> ) = 0
[pid 20016] futex(0x585870, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20020] wait4(20018, Process 20020 suspended
 <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20017] futex(0x585870, FUTEX_WAKE, 1) = 1
[pid 20016] <... futex resumed> )       = 0
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0x585870, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20032] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20032] futex(0x585158, FUTEX_WAKE, 1) = 1
[pid 20030] <... futex resumed> )       = 0
[pid 20032] wait4(20031,  <unfinished ...>
[pid 20030] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20032] <... wait4 resumed> [{WIFEXITED(s) && WEXITSTATUS(s) == 1}], 0, {ru_utime={0, 4000}, ru_stime={0, 0}, ...}) = 20031
[pid 20030] <... select resumed> )      = 0 (Timeout)
[pid 20032] write(1, "child------------ 20031\n", 24 <unfinished ...>
[pid 20030] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20032] <... write resumed> )       = 24
[pid 20032] exit_group(0)               = ?
Process 20032 detached
Process 20018 resumed
[pid 20019] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20019] futex(0x585158, FUTEX_WAKE, 1) = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20019] futex(0x585870, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 1
[pid 20016] <... futex resumed> )       = 0
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] futex(0x585870, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20019] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20019] futex(0x585158, FUTEX_WAKE, 1) = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20019] futex(0x585870, FUTEX_WAKE, 1) = 1
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20016] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] futex(0x585870, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20019] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20019] futex(0x585158, FUTEX_WAKE, 1) = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20019] mmap(NULL, 131072, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... mmap resumed> )        = 0x7f814203e000
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20019] futex(0x585158, FUTEX_WAKE, 1) = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20019] futex(0x585870, FUTEX_WAKE, 1) = 1
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... futex resumed> )       = 0
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... futex resumed> )       = 1
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0}) = -1 ETIMEDOUT (Connection timed out)
[pid 20016] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] <... futex resumed> )       = 1
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0}) = -1 ETIMEDOUT (Connection timed out)
[pid 20016] futex(0x585158, FUTEX_WAKE, 1) = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20019] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20016] sched_yield( <unfinished ...>
[pid 20019] <... futex resumed> )       = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20016] <... sched_yield resumed> ) = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] futex(0x5850e0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] <... futex resumed> )       = 0
[pid 20016] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] madvise(0xc208040000, 8192, MADV_DONTNEED) = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] madvise(0xc20804a000, 32768, MADV_DONTNEED <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] <... madvise resumed> )     = 0
[pid 20016] madvise(0xc208064000, 638976, MADV_DONTNEED) = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20016] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... futex resumed> )       = 1
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0}) = -1 ETIMEDOUT (Connection timed out)
[pid 20016] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... futex resumed> )       = 1
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20016] futex(0x585870, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20019] futex(0x585870, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] <... futex resumed> )       = 0
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20016] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0x585870, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20019] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20019] futex(0x585158, FUTEX_WAKE, 1) = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20019] futex(0x585870, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 1
[pid 20016] <... futex resumed> )       = 0
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20016] futex(0x585870, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20019] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20019] futex(0x585158, FUTEX_WAKE, 1) = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20019] futex(0x585870, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 1
[pid 20016] <... futex resumed> )       = 0
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAKE, 1 <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] <... futex resumed> )       = 1
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] <... futex resumed> )       = 0
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] <... futex resumed> )       = 1
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20016] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... futex resumed> )       = 0
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] <... futex resumed> )       = 1
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] <... futex resumed> )       = 1
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0}) = -1 ETIMEDOUT (Connection timed out)
[pid 20016] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... futex resumed> )       = 1
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20019] futex(0x585158, FUTEX_WAKE, 1) = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 1
[pid 20016] <... futex resumed> )       = 0
[pid 20019] futex(0x585870, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] <... futex resumed> )       = 0
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20016] futex(0x585870, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20017] select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)
[pid 20016] futex(0x585870, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20019] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20019] futex(0x585158, FUTEX_WAKE, 1) = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20019] futex(0x585870, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 1
[pid 20016] <... futex resumed> )       = 0
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] futex(0x585870, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20019] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20019] futex(0x585158, FUTEX_WAKE, 1) = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20019] futex(0x585870, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 1
[pid 20016] <... futex resumed> )       = 0
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAKE, 1) = 0
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... futex resumed> )       = 1
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20016] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1) = 1
[pid 20019] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0}) = -1 ETIMEDOUT (Connection timed out)
[pid 20016] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... futex resumed> )       = 0
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] <... futex resumed> )       = 1
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAKE, 1) = 0
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... futex resumed> )       = 1
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20016] <... futex resumed> )       = 0
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] <... futex resumed> )       = 1
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20016] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1) = 1
[pid 20019] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0}) = -1 ETIMEDOUT (Connection timed out)
[pid 20016] futex(0x585158, FUTEX_WAKE, 1) = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20019] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20016] sched_yield( <unfinished ...>
[pid 20019] <... futex resumed> )       = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20016] <... sched_yield resumed> ) = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] futex(0x5850e0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20016] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] <... futex resumed> )       = 1
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0}) = -1 ETIMEDOUT (Connection timed out)
[pid 20016] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] <... futex resumed> )       = 1
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20019] futex(0x585158, FUTEX_WAKE, 1) = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 1
[pid 20016] <... futex resumed> )       = 0
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL) = -1 EAGAIN (Resource temporarily unavailable)
[pid 20016] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] <... futex resumed> )       = 0
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] <... futex resumed> )       = 0
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] <... futex resumed> )       = 1
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20016] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] <... futex resumed> )       = 1
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0}) = -1 ETIMEDOUT (Connection timed out)
[pid 20016] futex(0x585158, FUTEX_WAKE, 1) = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] <... futex resumed> )       = 1
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] <... futex resumed> )       = 0
[pid 20019] futex(0x585870, FUTEX_WAKE, 1 <unfinished ...>
[pid 20016] futex(0x585870, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] <... futex resumed> )       = -1 EAGAIN (Resource temporarily unavailable)
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] futex(0x585870, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20019] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20019] futex(0x585158, FUTEX_WAKE, 1) = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20019] futex(0x585870, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 1
[pid 20016] <... futex resumed> )       = 0
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] futex(0x585870, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20019] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20019] futex(0x585158, FUTEX_WAKE, 1) = 1
[pid 20017] <... futex resumed> )       = 0
[pid 20019] futex(0x585870, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 1
[pid 20016] <... futex resumed> )       = 0
[pid 20019] futex(0x7f81421bbf60, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAKE, 1 <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] <... futex resumed> )       = 1
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20019] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] <... futex resumed> )       = -1 ETIMEDOUT (Connection timed out)
[pid 20016] futex(0x585158, FUTEX_WAKE, 1 <unfinished ...>
[pid 20017] <... futex resumed> )       = 0
[pid 20016] <... futex resumed> )       = 1
[pid 20017] select(0, NULL, NULL, NULL, {0, 20} <unfinished ...>
[pid 20016] futex(0xc2080172f0, FUTEX_WAKE, 1) = 1
[pid 20019] <... futex resumed> )       = 0
[pid 20017] <... select resumed> )      = 0 (Timeout)
[pid 20019] futex(0xc2080172f0, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20017] futex(0x585158, FUTEX_WAIT, 0, NULL <unfinished ...>
[pid 20016] futex(0x7f81421bbf60, FUTEX_WAIT, 0, {60, 0} <unfinished ...>
Process 20017 detached
Process 20020 resumed
Process 20018 detached
Process 20019 detached
Process 20020 detached

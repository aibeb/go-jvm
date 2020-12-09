package main

// ThreadStack Java虚拟机堆栈
// 每个线程都有一个私有Java虚拟机堆栈，与该线程同时创建。Java虚拟机堆栈存储Frame（第2.6节）。Java虚拟机堆栈类似于常规语言（例如C）的堆栈：它保存局部变量和部分结果，并在方法调用和返回中起作用。因为除了推送和弹出帧外，从不直接操纵Java虚拟机堆栈，所以可以为堆分配帧。Java虚拟机堆栈的内存不必是连续的。
// 实际存储Frame所以不要用指针
// Thread Stack Size (in Kbytes). (0 means use default stack size) [Sparc: 512; Solaris x86: 320 (was 256 prior in 5.0 and earlier); Sparc 64 bit: 1024; Linux amd64: 1024 (was 0 in 5.0 and earlier); all others 0.]

type ThreadStack map[string]Frame

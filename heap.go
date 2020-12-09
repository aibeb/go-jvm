package main

// Heap 线程之间共享的堆
// Java虚拟机具有一个在所有Java虚拟机线程之间共享的堆。堆是运行时数据区，从中分配所有类实例和数组的内存。堆是在虚拟机启动时创建的。
type Heap map[string]interface{}

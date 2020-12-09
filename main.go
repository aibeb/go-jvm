package main

// https://github.com/golang/go/wiki/CodeReviewComments
// https://docs.oracle.com/javase/specs/jvms/se14/html/jvms-2.html#jvms-2.6
// https://docs.oracle.com/javase/specs/jvms/se14/html/jvms-6.html#jvms-6.5
// https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-7.html

import (
	"log"
)

func main() {
	// 日志
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 构造虚拟机
	v := &VirtualMachine{
		ClassLoader: ClassLoader{},
		Metaspace:   Metaspace{},
		Heap:        Heap{},
	}

	// 入口文件.class or .jars
	v.LoadClass("testdata/Sum.class")

	// 虚拟机执行
	v.Start("main")
}

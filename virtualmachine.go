package main

import "os"

// VirtualMachine 定义虚拟机
// Java虚拟机可以一次支持多个执行线程（JLS§17）。每个Java虚拟机线程都有其自己的 pc（程序计数器）寄存器。在任何时候，每个Java虚拟机线程都在执行单个方法的代码，即该线程的当前方法（第2.6节）。如果不是 native，则该pc寄存器包含当前正在执行的Java虚拟机指令的地址。如果线程当前正在执行的方法是native，则Java虚拟机的pc 寄存器值未定义。Java虚拟机的pc寄存器足够宽，可以returnAddress在特定平台上保存一个或本机指针。
type VirtualMachine struct {
	ClassLoader ClassLoader
	Metaspace   Metaspace
	Heap        Heap
}

// LoadClass 类加载方法
// 根据路径类加载，得到ClassFile
func (v *VirtualMachine) LoadClass(path string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	// 读取字节码文件
	classFile := &ClassFile{}
	classFile.read(f)

	// 存储到metaspace
	classInfo := classFile.ConstantPool[classFile.ThisClass-1].(ConstantClassInfo)

	v.Metaspace[string(classFile.ConstantPool[classInfo.NameIndex-1].(ConstantUtf8Info).Bytes)] = *classFile

	return nil
}

// Start 开始虚拟机执行
func (v *VirtualMachine) Start(methodName string) {
	// 构造线程
	t := &Thread{
		ThreadStack: ThreadStack{},
		Heap:        v.Heap,
		Metaspace:   v.Metaspace,
	}
	// 入口方法为main
	if methodName == "" {
		methodName = "main"
	}
	// 定位要执行的classFile
	for _, c := range v.Metaspace {
		t.Invoke(&c, methodName, nil)
	}
}

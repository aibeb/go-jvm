怒了，用Go手撸一个JVM

# 为什么有这个想法？

总有些人喜欢拿着对JVM的一知半解来指点江山。我想让这些人知道你未必真的了解JVM

我们首先从JVM的基本知识开始开始吧

# 我们补充几点JVM对基本

## 字节码文件规范

> class文件由8位字节的流。通过分别读取两个和四个连续的8位字节来构造16位和32位量。多字节数据项始终以高字节顺序存储，高字节在前。就是传说的大端序存储

具体参考 [class文件格式](https://docs.oracle.com/javase/specs/jvms/se14/html/jvms-4.html)

```markdown
ClassFile {
    u4             magic; 4字节 0xCAFEBABE
    u2             minor_version; 2字节
    u2             major_version; 2字节
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
```

## 常量池 constant_pool

> 每个constant_pool表条目的格式 由其第一个“标签”字节指示

详细请参考[详细标签](https://docs.oracle.com/javase/specs/jvms/se14/html/jvms-4.html#jvms-4.4)

CONSTANT_Class	    7	第4.4.1节
CONSTANT_Fieldref	9	第4.4.2条
CONSTANT_Methodref	10	第4.4.2条
CONSTANT_InterfaceMethodref	11	第4.4.2条
CONSTANT_String	    8	§4.4.3
CONSTANT_Integer	3	§4.4.4
CONSTANT_Float	    4	§4.4.4
CONSTANT_Long	    5	第4.4.5节
CONSTANT_Double	    6	第4.4.5节
CONSTANT_NameAndType	12	第4.4.6节
CONSTANT_Utf8	    1	第4.4.7节
CONSTANT_MethodHandle	15	第4.4.8节
CONSTANT_MethodType	16	§4.4.9
CONSTANT_Dynamic	17	§4.4.10
CONSTANT_InvokeDynamic	18	§4.4.10
CONSTANT_Module	    19	§4.4.11
CONSTANT_Package	20

> constant_pool表的索引从1到constant_pool_count-1

存储以下信息

* 各种串常量
* 类和接口名
* 字段名


研读完之后我们直接采用go结构体来定义常量, 并定义好常量池

```go
type ConstantInfo struct {
	Tag              byte   // tag
	NameIndex        uint16 // name_index
	ClassIndex       uint16 // class_index
	NameAndTypeIndex uint16 // name_and_type_index
	StringIndex      uint16 // string_index
	DescriptorIndex  uint16 // descriptor_index
	HighBytes        uint32 //
	LowBytes         uint32 //
	Bytes            []byte // The bytes array contains the bytes of the string
	ReferenceKind    byte
	ReferenceIndex   uint16
}
type ConstantPool []ConstantInfo
```

## JVM内置了汇编指令
> JVM是一个堆栈计算机。每条指令都编码为一个字节，后面可以跟一些其他参数
> Java虚拟机指令操作码（包括保留的操作码（第6.2节））到这些操作码表示的指令的助记符的映射。

例如

```markdown
00 (0x00)    nop
01 (0x01)    aconst_null
02 (0x02)    iconst_m1
03 (0x03)    iconst_0
04 (0x04)    iconst_1
05 (0x05)    iconst_2
06 (0x06)    iconst_3
07 (0x07)    iconst_4
08 (0x08)    iconst_5
09 (0x09)    lconst_0
10 (0x0a)    lconst_1
11 (0x0b)    fconst_0
12 (0x0c)    fconst_1
13 (0x0d)    fconst_2
14 (0x0e)    dconst_0
15 (0x0f)    dconst_1
16 (0x10)    bipush
17 (0x11)    sipush
18 (0x12)    ldc
19 (0x13)    ldc_w
20 (0x14)    ldc2_w
```

[详细指令对照表](https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-7.html)

因为go最终也是编译成了汇编指令，所以我们可以直接采用如下办法处理汇编指令

# 我们开始定义类加载器









# 我们首先创建文件Sum.java

```java
public class Sum {
  public static int add(int a, int b) {
    return a + b;
  }
}
```

# 编译Sum.java得到Sum.class

```shell
javac Sum.java
```

我们打开Sum.class

TODO

```shell
javap -c Sum.class
```

TOOD


# 根据JVM规范划分结构体

```golang

// 常量，静态变量
type Metaspace map[string]ClassFile

// 堆是运行时数据区，从中分配所有类实例和数组的内存。堆是在虚拟机启动时创建的。
type Heap map[string]interface{}

// 虚拟机堆栈
// 每个线程都有一个私有Java虚拟机堆栈，与该线程同时创建。Java虚拟机堆栈存储Frame。
type ThreadStack map[string]Frame

type Frame struct {
	ConstantPool []interface{}  // 常量池引用
	Locals       []interface{}  // 局部变量表
	Stack        Stack          // 操作数栈
}

// 线程
type Thread struct {
	// 管理frame
	ThreadStack ThreadStack
	// 指向共享heap
	Heap Heap
	// 指向Metaspace
	Metaspace Metaspace
}

// 虚拟机
type VirtualMachine struct {
	ClassLoader ClassLoader
	Metaspace   Metaspace
	Heap        Heap
}

```

# 参考连接
https://www.oracle.com/java/technologies/javase/vmoptions-jsp.html
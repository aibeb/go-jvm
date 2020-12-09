package main

import (
	"encoding/binary"
	"fmt"
	"log"
)

// Thread 线程
type Thread struct {
	// 管理frame
	ThreadStack ThreadStack
	// 指向共享heap
	Heap Heap
	// 指向Metaspace
	Metaspace Metaspace
}

// Operate 执行操作码
func (t *Thread) Operate(c *ClassFile, f *Frame, code []byte) (interface{}, error) {
	// 循环执行code
	i := 0
	for {
		op := code[i]
		log.Printf("%02x", op)
		switch op {
		case 0x00: // NOP
		case 0x01: // ACONST_NULL
			f.Stack.Push(nil)
			i = i + 1
		case 0x02: // ICONST_M1
			f.Stack.Push(int32(-1))
			i = i + 1
		case 0x03: // ICONST_0
			f.Stack.Push(int32(0))
			i = i + 1
		case 0x04: // ICONST_1
			f.Stack.Push(int32(1))
			i = i + 1
		case 0x05: // ICONST_2
			f.Stack.Push(int32(2))
			i = i + 1
		case 0x06: // ICONST_3
			f.Stack.Push(int32(3))
			i = i + 1
		case 0x3d: // istore_2
			value := f.Stack.Pop()
			f.Locals[2] = value
			i = i + 1
		case 0xBB: // NEW TODO
			// 定位到运行时常量池
			index := binary.BigEndian.Uint16(code[i+1 : i+3])

			classInfo := (f.ConstantPool)[index-1].(ConstantClassInfo)

			className := string((f.ConstantPool)[classInfo.NameIndex-1].(ConstantUtf8Info).Bytes)

			// 成功解析该类后，如果尚未初始化它，则将其初始化 TODO 调用类加载器

			// 新对象的实例变量被初始化为其默认初始值

			// 绑定对象到heap堆
			(t.Heap)[className] = c

			// 实例的 objectref a reference被压入stack
			f.Stack.Push(c)

			i = i + 3
		case 0x59: // DUP TODO
			value := f.Stack.Pop()
			f.Stack.Push(value)
			f.Stack.Push(value)
			i = i + 1
		case 0x2a: // ALOAD_0
			i = i + 1
		case 0x1b: // ILOAD_1 iload_<n> 局部变量入栈
			f.Stack.Push(f.Locals[1])
			i = i + 1
		case 0x1c: // ILOAD_2 iload_<n> 局部变量入栈
			f.Stack.Push(f.Locals[2])
			i = i + 1
		case 0x60: // iadd
			// 参数出栈，结果入栈
			value1 := f.Stack.Pop()
			value2 := f.Stack.Pop()
			f.Stack.Push(value1.(int32) + value2.(int32))
			log.Println(f.Stack)
			i = i + 1
		case 0xac: // IRETURN 结果出栈，并返回结果
			return f.Stack.Pop(), nil
		case 0xb7: // INVOKESPECIAL 直接调用实例初始化方法以及当前类及其超类型的方法
			index := binary.BigEndian.Uint16(code[i+1 : i+3])

			// 得到需要调用的方法
			methodrefInfo := (f.ConstantPool)[index-1].(ConstantMethodrefInfo)

			// 得到方法所属的类 Sum
			classInfo := (f.ConstantPool)[methodrefInfo.ClassIndex-1].(ConstantClassInfo)

			className := string((f.ConstantPool)[classInfo.NameIndex-1].(ConstantUtf8Info).Bytes)

			// 得到方法名 <init>
			nameAndTypeInfo := (f.ConstantPool)[methodrefInfo.NameAndTypeIndex-1].(ConstantNameAndTypeInfo)

			methodName := string((f.ConstantPool)[nameAndTypeInfo.NameIndex-1].(ConstantUtf8Info).Bytes)

			// 调用方法 TODO 这里需要访问metaspace
			c := (t.Metaspace)[className]
			t.Invoke(&c, methodName, nil)

			// 堆栈清空
			// f.Stack.Empty()

			i = i + 3
		case 0x4c: // ASTORE_1

			// 操作数堆栈顶部的objectref必须为returnAddress或类型reference。它从操作数堆栈中弹出
			objectref := f.Stack.Pop()
			f.Locals[1] = objectref

			i = i + 1
		case 0xb2: // GETSTATIC
			// index := binary.BigEndian.Uint16(code[i+1 : i+3])

			// fieldrefInfo := (f.ConstantPool)[index-1].(ConstantFieldrefInfo)

			// 得到方法所属的类 java/lang/System
			// classInfo := (f.ConstantPool)[fieldrefInfo.ClassIndex-1].(ConstantClassInfo)

			// 得到方法名 out
			// nameAndTypeInfo := (f.ConstantPool)[fieldrefInfo.NameAndTypeIndex-1].(ConstantNameAndTypeInfo)

			i = i + 3
		case 0x2b: // ALOAD_1
			// 局部变量中的objectref被压入操作数堆栈
			f.Stack.Push(f.Locals[1])

			i = i + 1
		case 0xb6: // INVOKEVIRTUAL 所调用基于所述对象的类的方法
			index := binary.BigEndian.Uint16(code[i+1 : i+3])
			// 得到需要调用的方法
			methodrefInfo := (f.ConstantPool)[index-1].(ConstantMethodrefInfo)

			// 得到方法所属的类 Sum java/io/PrintStream
			classInfo := (f.ConstantPool)[methodrefInfo.ClassIndex-1].(ConstantClassInfo)

			className := string((f.ConstantPool)[classInfo.NameIndex-1].(ConstantUtf8Info).Bytes)

			// 得到方法名 add println
			nameAndTypeInfo := (f.ConstantPool)[methodrefInfo.NameAndTypeIndex-1].(ConstantNameAndTypeInfo)

			methodName := string((f.ConstantPool)[nameAndTypeInfo.NameIndex-1].(ConstantUtf8Info).Bytes)

			// 如果要调用的方法不是native，则从操作数堆栈中弹出nargs参数值和objectref。在Java虚拟机堆栈上为要调用的方法创建一个新框架
			// 如果要调用的方法是native并且实现该方法的平台相关代码尚未绑定到Java虚拟机中（第5.6节），则可以完成此操作。所述 NARGS参数值和objectref是从操作数堆栈弹出并作为参数被传递给代码实现该方法
			// 如果该native方法返回值，则平台相关代码的返回值将以实现相关的方式转换为该native方法的返回类型， 并将其压入操作数堆栈

			// heap中查找对象，并执行相应方法
			c := (t.Heap)[className].(*ClassFile)
			value := t.Invoke(c, methodName, f.Stack[len(f.Stack)-len(f.Locals):len(f.Stack)]...)
			// f.Stack.Empty()

			f.Stack.Push(value)

			i = i + 3
		case 0xb1: // RETURN
			return nil, nil
		default:
			return nil, fmt.Errorf("not support code 0x%02x", op)
		}
	}
}

// Invoke 执行
// TODO 查找主线程
func (t *Thread) Invoke(c *ClassFile, methodName string, args ...interface{}) interface{} {
	// 定位要执行的方法
	for _, m := range c.Methods {
		if methodName == string(c.ConstantPool[m.NameIndex-1].(ConstantUtf8Info).Bytes) {
			// 处理未查询到方法的情况 TODO
			for _, a := range m.Attributes {
				// 类型推断
				switch a.(type) {
				case CodeAttribute:
					// 构造frame
					frame := Frame{
						ConstantPool: c.ConstantPool,
						Locals:       make([]interface{}, a.(CodeAttribute).MaxLocals, a.(CodeAttribute).MaxLocals),
					}

					// 存储frame methodName:frame
					t.ThreadStack[methodName] = frame

					for i := 0; i < len(args); i++ {
						frame.Locals[i] = args[i]
					}

					value, _ := t.Operate(c, &frame, a.(CodeAttribute).Code)

					return value

				default:
				}
			}
		}
	}
	return nil
}

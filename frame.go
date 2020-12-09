package main

// Stack 操作数栈
type Stack []interface{}

// Push 压栈
func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

// Pop 出栈
func (s *Stack) Pop() interface{} {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

// Empty 清空
func (s *Stack) Empty() {
	*s = nil
}

// Frame 运行帧
// 每次调用方法时都会创建一个新Frame
// 方法调用完成时，无论该完成是正常的还是突然的（它引发未捕获的异常），它都会被销毁
// 局部变量数组
// 自己的操作数栈
// 当前方法类的运行时常量池的引用
type Frame struct {
	ConstantPool []interface{} // 常量池引用
	Locals       []interface{} // 局部变量表
	Stack        Stack         // 操作数栈
}

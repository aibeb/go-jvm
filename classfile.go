package main

import (
	"encoding/binary"
	"io"
)

// 常量
const (
	CONSTANTClass              = 7
	CONSTANTFieldref           = 9
	CONSTANTMethodref          = 10
	CONSTANTInterfaceMethodref = 11
	CONSTANTString             = 8
	CONSTANTInteger            = 3
	CONSTANTFloat              = 4
	CONSTANTLong               = 5
	CONSTANTDouble             = 6
	CONSTANTNameAndType        = 12
	CONSTANTUtf8               = 1
	CONSTANTMethodHandle       = 15
	CONSTANTMethodType         = 16
	CONSTANTDynamic            = 17
	CONSTANTInvokeDynamic      = 18
	CONSTANTModule             = 19
	CONSTANTPackage            = 20
)

// ConstantClassInfo is used to represent a class or an interface:
type ConstantClassInfo struct {
	Tag       u1
	NameIndex u2
}

// ConstantFieldrefInfo has the value CONSTANT_Fieldref
type ConstantFieldrefInfo struct {
	Tag              u1
	ClassIndex       u2
	NameAndTypeIndex u2
}

// ConstantMethodrefInfo todo
type ConstantMethodrefInfo struct {
	Tag              u1
	ClassIndex       u2
	NameAndTypeIndex u2
}

// ConstantInterfaceMethodrefInfo todo
type ConstantInterfaceMethodrefInfo struct {
	Tag              u1
	ClassIndex       u2
	NameAndTypeIndex u2
}

// ConstantStringInfo todo
type ConstantStringInfo struct {
	Tag         u1
	StringIndex u2
}

// ConstantIntegerInfo todo
type ConstantIntegerInfo struct {
	Tag   u1
	Bytes u4
}

// ConstantFloatInfo todo
type ConstantFloatInfo struct {
	Tag   u1
	Bytes u4
}

// ConstantLongInfo todo
type ConstantLongInfo struct {
	Tag       u1
	HighBytes u4
	LowBytes  u4
}

// ConstantDoubleInfo todo
type ConstantDoubleInfo struct {
	Tag       u1
	HighBytes u4
	LowBytes  u4
}

// ConstantNameAndTypeInfo todo
type ConstantNameAndTypeInfo struct {
	Tag             u1
	NameIndex       u2
	DescriptorIndex u2
}

// ConstantUtf8Info todo
type ConstantUtf8Info struct {
	Tag    u1
	Length u2
	Bytes  []u1
}

// ConstantMethodHandleInfo todo
type ConstantMethodHandleInfo struct {
	Tag            u1
	ReferenceKind  u1
	ReferenceIndex u2
}

// ConstantMethodTypeInfo todo
type ConstantMethodTypeInfo struct {
	Tag             u1
	DescriptorIndex u2
}

// ConstantDynamicInfo todo
type ConstantDynamicInfo struct {
	Tag                      u1
	BootstrapMethodAttrIndex u2
	NameAndTypeIndex         u2
}

// ConstantInvokeDynamicInfo todo
type ConstantInvokeDynamicInfo struct {
	Tag                      u1
	BootstrapMethodAttrIndex u2
	NameAndTypeIndex         u2
}

// ConstantModuleInfo todo
type ConstantModuleInfo struct {
	Tag       u1
	NameIndex u2
}

// ConstantPackageInfo todo
type ConstantPackageInfo struct {
	Tag       u1
	NameIndex u2
}

type u1 = byte
type u2 = uint16
type u4 = uint32

func readU1(r io.Reader) uint8 {
	bytes := make([]byte, 1)
	r.Read(bytes)
	return bytes[0]
}
func readU2(r io.Reader) uint16 {
	bytes := make([]byte, 2)
	r.Read(bytes)
	return binary.BigEndian.Uint16(bytes)
}
func readU4(r io.Reader) uint32 {
	bytes := make([]byte, 4)
	r.Read(bytes)
	return binary.BigEndian.Uint32(bytes)
}
func readU8(r io.Reader) uint64 {
	bytes := make([]byte, 8)
	r.Read(bytes)
	return binary.BigEndian.Uint64(bytes)
}

// ExceptionTable todo
type ExceptionTable struct {
	StartPc   u2
	EndPc     u2
	HandlerPc u2
	CatchType u2
}

// CodeAttribute todo
type CodeAttribute struct {
	AttributeNameIndex   u2
	AttributeLength      u4
	MaxStack             u2   // max_stack给出了该方法执行过程中任何时候该方法的操作数堆栈的最大深度
	MaxLocals            u2   // max_locals给出在调用此方法（第2.6.1节）时分配的局部变量数组中的局部变量数
	CodeLength           u4   // codeLength给出code此方法在数组中的字节数
	Code                 []u1 // 该code数组给出了实现该方法的Java虚拟机代码的实际字节
	ExceptionTableLength u2
	ExceptionTable       []ExceptionTable
	AttributeCount       u2
	Attributes           []ExceptionTable
}

// AttributeInfo todo
type AttributeInfo struct {
	AttributeNameIndex u2 //该attributeNameIndex项目必须是该类的常量池中的有效无符号16位索引
	AttributeLength    u4 //attributeLength指示后续信息的长度（以字节为单位）。该长度不包括包含attributeNameIndex 和attributeLength项的前六个字节
	Info               []u1
}

// MethodInfo todo
type MethodInfo struct {
	AccessFlags     u2 //该accessFlags项目的值是用于表示此方法的访问许可权和属性的标志的掩码
	NameIndex       u2 //该nameIndex项目的值必须是constantPool表中的有效索引。该constantPool索引处的 条目必须是一个constantUtf8Info结构
	DescriptorIndex u2 //该descriptorIndex项目的值必须是constantPool表中的有效索引。该constantPool索引处的 条目必须是constantUtf8Info代表有效方法描述符的 结构
	AttributesCount u2 //项目的值 attributesCount指示此方法的其他属性的数量
	Attributes      []interface{}
}

// FieldInfo todo
type FieldInfo struct {
	AccessFlags     u2
	NameIndex       u2 //该nameIndex项目的值 必须是constantPool表中的有效索引。该constantPool索引处的条目必须是一个constantUtf8Info结构
	DescriptorIndex u2 //该descriptorIndex项目的值 必须是constantPool表中的有效索引。该constantPool索引处的条目必须是代表有效字段描述符constantUtf8Info
	AttributesCount u2
	Attributes      []AttributeInfo
}

// ClassFile todo
type ClassFile struct {
	Magic             u4
	MinorVersion      u2
	MajorVersion      u2
	ConstantPoolCount u2
	ConstantPool      []interface{} //constantPool_count-1
	AccessFlags       u2
	ThisClass         u2 // this_class项目的值必须是constantPool表中的有效索引
	SuperClass        u2 // 该项的值super_class 必须为零或必须是constantPool表中的有效索引
	InterfacesCount   u2
	Interfaces        []u2 // interfaces数组 中的每个值都必须是constantPool表中的有效索引
	FieldsCount       u2
	Fields            []FieldInfo
	MethodsCount      u2
	Methods           []MethodInfo
	AttributesCount   u2
	Attributes        []AttributeInfo
}

// 读取字节码
func (c *ClassFile) read(r io.Reader) {
	c.Magic = readU4(r)
	c.MinorVersion = readU2(r)
	c.MajorVersion = readU2(r)
	c.ConstantPoolCount = readU2(r)

	// 常量池 从1开始迭代 根据规范constantPool_count = len(ConstantPool)+1
	constPool := make([]interface{}, 0)

	for i := uint16(1); i < c.ConstantPoolCount; i++ {
		tag := readU1(r)

		switch tag {
		case CONSTANTUtf8:
			constantInfo := ConstantUtf8Info{
				Tag:    tag,
				Length: readU2(r),
			}
			bytes := make([]byte, constantInfo.Length)
			r.Read(bytes)
			constantInfo.Bytes = bytes
			constPool = append(constPool, constantInfo)

		case CONSTANTClass:
			constantInfo := ConstantClassInfo{
				Tag:       tag,
				NameIndex: readU2(r),
			}
			constPool = append(constPool, constantInfo)
		case CONSTANTString:
			constantInfo := ConstantStringInfo{
				Tag: tag,
				// string_index项目的值必须是constantPool表中的有效索引。该constantPool索引处的 条目必须是一个constantUtf8Info结构
				// c.ConstantPool[string_index-1]
				StringIndex: readU2(r),
			}
			constPool = append(constPool, constantInfo)
		case CONSTANTFieldref:
			constantInfo := ConstantFieldrefInfo{
				Tag:              tag,
				ClassIndex:       readU2(r),
				NameAndTypeIndex: readU2(r),
			}
			constPool = append(constPool, constantInfo)
		case CONSTANTMethodref:
			constantInfo := ConstantMethodrefInfo{
				Tag:              tag,
				ClassIndex:       readU2(r),
				NameAndTypeIndex: readU2(r),
			}
			constPool = append(constPool, constantInfo)
		case CONSTANTNameAndType:
			constantInfo := ConstantNameAndTypeInfo{
				Tag:             tag,
				NameIndex:       readU2(r),
				DescriptorIndex: readU2(r),
			}
			constPool = append(constPool, constantInfo)

		case CONSTANTMethodHandle:

		default:
			// TODO
		}

	}

	c.ConstantPool = constPool

	c.AccessFlags = readU2(r)
	c.ThisClass = readU2(r)
	c.SuperClass = readU2(r)
	c.InterfacesCount = readU2(r)

	for i := uint16(0); i < c.InterfacesCount; i++ {
		c.Interfaces = append(c.Interfaces, readU2(r))
	}

	// 处理字段
	c.FieldsCount = readU2(r)

	for i := uint16(0); i < c.FieldsCount; i++ {

		fieldInfo := FieldInfo{
			AccessFlags:     readU2(r),
			NameIndex:       readU2(r),
			DescriptorIndex: readU2(r),
			AttributesCount: readU2(r),
		}

		for i := uint16(0); i < fieldInfo.AttributesCount; i++ {
			attributeInfo := AttributeInfo{
				AttributeNameIndex: readU2(r),
				AttributeLength:    readU4(r),
			}

			bytes := make([]byte, attributeInfo.AttributeLength)
			r.Read(bytes)
			attributeInfo.Info = bytes
			fieldInfo.Attributes = append(fieldInfo.Attributes, attributeInfo)
		}

		c.Fields = append(c.Fields, fieldInfo)
	}

	// 方法
	c.MethodsCount = readU2(r)

	for i := uint16(0); i < c.MethodsCount; i++ {

		methodInfo := MethodInfo{
			AccessFlags:     readU2(r),
			NameIndex:       readU2(r),
			DescriptorIndex: readU2(r),
			AttributesCount: readU2(r),
		}

		for i := uint16(0); i < methodInfo.AttributesCount; i++ {
			attributeNameIndex := readU2(r)
			attributeLength := readU4(r)

			info := make([]u1, attributeLength)
			r.Read(info)

			if "Code" == string(c.ConstantPool[attributeNameIndex-1].(ConstantUtf8Info).Bytes) {
				i := 0
				codeAttribute := CodeAttribute{
					AttributeNameIndex: attributeNameIndex,
					AttributeLength:    attributeLength,
				}

				codeAttribute.MaxStack = binary.BigEndian.Uint16(info[i : i+2])
				i = i + 2
				codeAttribute.MaxLocals = binary.BigEndian.Uint16(info[i : i+2])
				i = i + 2
				codeAttribute.CodeLength = binary.BigEndian.Uint32(info[i : i+4])
				i = i + 4

				// 每个getfield，putfield，getstatic和 putstatic指令的操作数必须表示表中的有效索引 constantPool。该索引引用的常量池条目必须是种类CONSTANT_Fieldref
				// 所述indexbyte每个操作数 invokespecial和invokestatic指令必须代表一个有效的索引到constantPool表中。如果class文件版本号小于52.0，则该索引引用的常量池条目必须为实物CONSTANT_Methodref；如果class文件版本号为52.0或更高，则该索引引用的常量池条目必须为 CONSTANT_Methodref或CONSTANT_InterfaceMethodref。

				codeAttribute.Code = info[i : uint32(i)+codeAttribute.CodeLength]

				i = int(uint32(i) + codeAttribute.CodeLength)

				codeAttribute.ExceptionTableLength = binary.BigEndian.Uint16(info[i : i+2])

				i = i + 2

				methodInfo.Attributes = append(methodInfo.Attributes, codeAttribute)
			}

		}

		c.Methods = append(c.Methods, methodInfo)
	}

	c.AttributesCount = readU2(r)
	for i := uint16(0); i < c.AttributesCount; i++ {
		attributeInfo := AttributeInfo{
			AttributeNameIndex: readU2(r),
			AttributeLength:    readU4(r),
		}

		bytes := make([]byte, attributeInfo.AttributeLength)
		r.Read(bytes)
		attributeInfo.Info = bytes
		c.Attributes = append(c.Attributes, attributeInfo)
	}
}

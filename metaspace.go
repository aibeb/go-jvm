package main

// Metaspace 存放 ClassFile Structure
// 当一个类被加载时，它的类加载器会负责在 metaspace 中分配空间用于存放这个类的元数据
// 常量、静态变量
type Metaspace map[string]ClassFile

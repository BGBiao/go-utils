package utils
import (
  "reflect"
)
// 检测对象中是否包含某个元素
// func ValueOf(i interface{}) Value 返回一个Value的空结构体
func Contains(obj interface{}, target interface{}) bool {
    targetValue := reflect.ValueOf(target)
    //func ArrayOf(count int, elem Type) Type 返回一个Type的结果体
    //func (k Kind) String() string 获取一个对象的类型
    switch reflect.TypeOf(target).Kind() {

    case reflect.Slice, reflect.Array:
        for i := 0; i < targetValue.Len(); i++ {
            // func (v Value) Index(i int) Value
            // func (v Value) Interface() (i interface{})
            if targetValue.Index(i).Interface() == obj {
                return true
            }
        }
    case reflect.Map:
        // func (v Value) MapIndex(key Value) Value
        // func (v Value) IsValid() bool
        if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
            return true
        }
    }
    return false
}

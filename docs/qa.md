
## mysql因为时区出现误差的问题
使用mysql定义时间类型时，尽量用datetime类型，这样就不会有时区的错误

## 出现interface conversion: interface {} is float64, not int
解决的方法是先将其转为float64
```
uid, ok := m["userId"]
if !ok {
    return 0, errors.New("找不到userid")
}
return uint64(uid.(float64)), nil


// 或者
userId, err := l.ctx.Value("userId").(json.Number).Int64()

```

## 在将int64数字类型unmarshal 转化为interface{} 的时候，默认会先使用浮点数转换，这样，可能会丢失精度
解决的方法是将int64类型转为string类型后存到interface{}中，取的时候再用strconv.ParseUint(tokenId, 10, 64), 见pkg下的jwt包

## 可不可以服务A引用服务B, 服务B引用服务A 
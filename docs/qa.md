
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
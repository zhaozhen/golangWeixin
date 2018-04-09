解析json的两种方式
```
data, err := ioutil.ReadAll(resp.Body)
if err == nil && data != nil {
    err = json.Unmarshal(data, value)
}
```
2
err = json.NewDecoder(resp.Body).Decode(value)
区别：
当来自 io.Reader stream 时就选用 json.Decoder
JSON 数据本来就在内存里，就选用 json.Unmarshal

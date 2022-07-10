## 1.总结几种 socket 粘包的解包方式：

### fix length

- 客户端与服务端约定好请求大小，每次发送/接收固定缓冲区大小数据
- 虽然可以解决粘包，但：当发送数据小于约定大小时容易出现数据内存冗余和浪费；当发送数据大于约定大小，会出现半包问题。

#### client

```go
const (
    BYTE_LENGTH = 1024
)
func client_tcp_fix_length(conn net.Conn) {
  sM := "this is a test msg;"
  tB := []byte(sM)
  for i := 0; i < len(sM)%BYTE_LENGTH+1; i++ {
    sB := make([]byte, BYTE_LENGTH)
    for j := i*BYTE_LENGTH; j < len(tB) && j < BYTE_LENGTH; j++ {
      sB[j] = tB[j]
    }
    _, err := conn.Write(sB)
    if err != nil {
      fmt.Println(err, ", index=", i)
      return
    }
    fmt.Println("send one")
  }
}
```

#### server

```go
const (
    BYTE_LENGTH = 1024
)
func server_tcp_fix_length(conn net.Conn) {
  var err error
  var buf = make([]byte, BYTE_LENGTH)
  for {
    _, err = conn.Read(buf)
    if err != nil {
      fmt.Println(err.Error())
      return
    }
  }
}
```

### delimiter based

- 基于定界符标记数据包边界
- 当数据量过大时，查找定界符会消耗部分性能

#### client

```go
func client_tcp_delimiter(conn net.Conn) {
  rM := "this is a test msg;\n"
  var sM string
  for i := 0; i < 1000; i++ {
    sM=rM+`\n`
    _, err := conn.Write([]byte(sM))
    if err != nil {
      fmt.Println(err, ",err index=", i)
      return
    }
    fmt.Println("send over once")
  }
}
```

#### server

```go
func server_tcp_delimiter(conn net.Conn) {
  r := bufio.NewReader(conn)
  for {
    slice, err := r.ReadSlice('\n')
    if err != nil {
      continue
    }
    fmt.Printf("%s", slice)
  }
}
```

### length field based frame decoder

- 客户端在协议头写入数据长度。服务器接收请求后，根据协议头里面的数据长度来决定接受多少数据。之后按照包长度偏移量对接收数据进行解码，以获取目标消息体数据。

#### client

```go
func client_tcp_delimiter(conn net.Conn) {
  rM := "this is a test msg;\n"
  var sM string
  for i := 0; i < 1000; i++ {
    sM=rM+`\n`
    _, err := conn.Write([]byte(sM))
    if err != nil {
      fmt.Println(err, ",err index=", i)
      return
    }
    fmt.Println("send over once")
  }
}
```

#### server

```go
func server_tcp_frame_decoder(conn net.Conn) {
  var buf = make([]byte, 0)
  var rChan = make(chan []byte, 16)
  go func() {
    select {
    case data := <-rChan:
      fmt.Println("channel=", string(data))
    }
  }()
  rBuf := make([]byte, 1024)
  for {
    n, err := conn.Read(rBuf)
    if err != nil {
      fmt.Println(conn.RemoteAddr().String(), " connection error: ", err.Error())
      return
    }
    protocol.Unpack(append(buf, rBuf[:n]...), rChan)
  }
}
```

## 2.实现一个从 socket connection 中解码出 goim 协议的解码器。

- 客户端样例：main.go
- 协议解码器：protocol\go_im.go
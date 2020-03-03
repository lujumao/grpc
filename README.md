# protobuf grpc


##安装

二进制安装 

地址 https://github.com/protocolbuffers/protobuf/releases


liunx 编译安装

依赖库的安装
```
yum install build-essential
```

```
yum install glibc-headers
```
```
yum install gcc-c++
```

安装protoc
```
git clone https://github.com/google/protobuf
./autogen.sh  # 生成 configure 文件
./configure
make
make check
sudo make install
sudo ldconfig
```

golang的protobuf包
```redshift
go get -u github.com/golang/protobuf/protoc-gen-go
cp protoc-gen-go  /usr/local/bin/
```

编译代码

- --go_out: 生产golang的proto文件
- plugins: 服务端
- grpc: grpc服务
- ../active/: 输出文件的位置
- *.proto: 编译当前目录的所有的 proto文件

编译服务
```redshift
protoc --go_out=plugins=grpc:../test/ *.proto 
```

编译客户
```redshift
protoc --go_out=../test/ *.proto
```








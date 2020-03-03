#!/bin/bash

#cp文件到项目的目录里 -> 在编译这个目录的文件

function ergodic() {
 for file in `ls $1`     #$1是调用时的第一个参数  函数找出所有非目录的文件
  do
  if [ -d "$1/$file" ]
    then
#    echo "$1/$file"
    ergodic "$1/$file"
  else
    cp -avx $1/$file  proto/
    echo $1/$file
  fi
  done
}

ergodic "/mnt/win/docker/data_protocol/"    #需要变了的文件夹

cd proto
protoc --go_out=../test/  *.proto #proto 客户端
#protoc --go_out=plugins=grpc:../test/ *.proto #proto 服务端


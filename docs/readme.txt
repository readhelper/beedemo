#生成根证书
openssl req -new -x509 -days 5480 -keyout root.key -out root.crt -passout pass:1234

#为顶级CA的私钥文件去除保护口令
openssl rsa -in root.key -out root.key

#为应用证书/中级证书生成私钥文件
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr

#使用CA的公私钥文件给 csr 文件签名，生成应用证书，有效期5年
openssl ca -in server.csr -out server.crt -cert root.crt -keyfile root.key -days 1826 -policy policy_anything

使用CA的公私钥文件给 csr 文件签名，生成中级证书，有效期5年
openssl ca -extensions v3_ca -in server.csr -out server.crt -cert root.crt -keyfile root.key -days 1826 -policy policy_anything


#生成个人证书
openssl genrsa -out hao.key 2048
openssl req -new -key hao.key -out hao.csr
openssl ca -in hao.csr -out hao.crt -cert server.crt -keyfile server.key -days 1826 -policy policy_anything

openssl x509 -req -in hao.csr -CA root.crt -CAkey root.key -CAcreateserial -extfile extfile.cnf -out hao.crt -days 5000 
curl --cacert root.crt --cert hao.crt --key hao.key  https://localhost:2379/v2/keys/  (-v)

#生成client证书
openssl genrsa -out client.key 2048
openssl req -new -key client.key -subj "/CN=client" -out client.csr
echo extendedKeyUsage=clientAuth > extfile.cnf
openssl x509 -req -in client.csr -CA server.crt -CAkey server.key -CAcreateserial -extfile extfile.cnf -out client.crt -days 5000 

openssl pkcs12 -export -out hao.pfx -inkey hao.key -in hao.crt 

如果在这步出现错误信息：
[weigw@TEST bin]$ openssl ca -in client.csr -out client.crt -cert ca.crt -keyfile ca.key 
Using configuration from /usr/share/ssl/openssl.cnf I am unable to access the ./demoCA/newcerts directory ./demoCA/newcerts: No such file or directory 
[weigw@TEST bin]$ 
自己手动创建一个CA目录结构：
[weigw@TEST bin]$ mkdir ./demoCA
[weigw@TEST bin]$ mkdir demoCA/newcerts
创建个空文件：
[weigw@TEST bin]$ vi demoCA/index.txt
创建文件：
[weigw@TEST bin]$ vi demoCA/serial
向文件中写入01
并且第二行空一行


#curl 验证命令
curl -cacert ./server.crt https://127.0.0.1:2379/v2/keys/

curl --cacert /path/to/ca.crt --cert /path/to/client.crt --key /path/to/client.key -L https://127.0.0.1:2379/v2/keys/
  
curl --cacert root.crt --cert client.crt --key client.key -L https://127.0.0.1:8081/
curl --cacert root.crt --cert client.crt --key client.key  https://localhost:2379/v2/keys/  (-v)
  
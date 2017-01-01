#���ɸ�֤��
openssl req -new -x509 -days 5480 -keyout root.key -out root.crt -passout pass:1234

#Ϊ����CA��˽Կ�ļ�ȥ����������
openssl rsa -in root.key -out root.key

#ΪӦ��֤��/�м�֤������˽Կ�ļ�
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr

#ʹ��CA�Ĺ�˽Կ�ļ��� csr �ļ�ǩ��������Ӧ��֤�飬��Ч��5��
openssl ca -in server.csr -out server.crt -cert root.crt -keyfile root.key -days 1826 -policy policy_anything

ʹ��CA�Ĺ�˽Կ�ļ��� csr �ļ�ǩ���������м�֤�飬��Ч��5��
openssl ca -extensions v3_ca -in server.csr -out server.crt -cert root.crt -keyfile root.key -days 1826 -policy policy_anything


#���ɸ���֤��
openssl genrsa -out hao.key 2048
openssl req -new -key hao.key -out hao.csr
openssl ca -in hao.csr -out hao.crt -cert server.crt -keyfile server.key -days 1826 -policy policy_anything

openssl x509 -req -in hao.csr -CA root.crt -CAkey root.key -CAcreateserial -extfile extfile.cnf -out hao.crt -days 5000 
curl --cacert root.crt --cert hao.crt --key hao.key  https://localhost:2379/v2/keys/  (-v)

#����client֤��
openssl genrsa -out client.key 2048
openssl req -new -key client.key -subj "/CN=client" -out client.csr
echo extendedKeyUsage=clientAuth > extfile.cnf
openssl x509 -req -in client.csr -CA server.crt -CAkey server.key -CAcreateserial -extfile extfile.cnf -out client.crt -days 5000 

openssl pkcs12 -export -out hao.pfx -inkey hao.key -in hao.crt 

������ⲽ���ִ�����Ϣ��
[weigw@TEST bin]$ openssl ca -in client.csr -out client.crt -cert ca.crt -keyfile ca.key 
Using configuration from /usr/share/ssl/openssl.cnf I am unable to access the ./demoCA/newcerts directory ./demoCA/newcerts: No such file or directory 
[weigw@TEST bin]$ 
�Լ��ֶ�����һ��CAĿ¼�ṹ��
[weigw@TEST bin]$ mkdir ./demoCA
[weigw@TEST bin]$ mkdir demoCA/newcerts
���������ļ���
[weigw@TEST bin]$ vi demoCA/index.txt
�����ļ���
[weigw@TEST bin]$ vi demoCA/serial
���ļ���д��01
���ҵڶ��п�һ��


#curl ��֤����
curl -cacert ./server.crt https://127.0.0.1:2379/v2/keys/

curl --cacert /path/to/ca.crt --cert /path/to/client.crt --key /path/to/client.key -L https://127.0.0.1:2379/v2/keys/
  
curl --cacert root.crt --cert client.crt --key client.key -L https://127.0.0.1:8081/
curl --cacert root.crt --cert client.crt --key client.key  https://localhost:2379/v2/keys/  (-v)
  
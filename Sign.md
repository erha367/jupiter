##### 签名算法

见：https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=4_3

##### 关于编码：

生成待签名字符串时，所有参数都无需进行URL编码；

发送http请求时，如果 Content-Type 为 application/x-www-form-urlencoded，则所有请求参数的值均需要做 URL 编码。
编码采用 RFC 3986，使用 %XY 对特殊字符例如汉字进行百分比编码，其中“X”和“Y”为十六进制字符（0-9 和大写字母 A-F），使用小写将引发错误。参数键和=符号不需要编码。

##### 举例：

1.假如某个接口有如下参数：

|参数名|参数值|
| ------ | ------ |
|timeStamp |  1620299911 |
|source |   1（不参与签名） |
|version |   4.4.3.0（不参与签名） |
|telephone |  123456789 |
|password |   pwd123456 |
|confirmPassword |  pwd123456 |
|verificationCode | 6724 |
|info |  我是中文 字符 |

2.生成待签名串：

按字典序排列参数生成字符串：confirmPassword=pwd123456&info=我是中文 字符&password=pwd123456&telephone=123456789&timeStamp=1620299911&verificationCode=6724

在字符串的最后附加上与客户端约定的密钥(key:classin)：confirmPassword=pwd123456&info=我是中文 字符&password=pwd123456&telephone=123456789&timeStamp=1620299911&verificationCode=6724&key=classin

3.按签名算法(md5)生成签名值: sign=98c79517c47d3d3cd91b11d95e0249b2

4.不参与签名的参数使用 ignoreSign 指出其键值集（使用逗号分隔，ignoreSign为非必传）: ignoreSign=source,version

5.最终请求的参数如下：

|参数名|参数值|
| ------ | ------ |
|timeStamp |  1620299911 |
|source |   1 |
|version |   4.4.3.0 |
|telephone |  123456789 |
|password  |  pwd123456 |
|confirmPassword | pwd123456 |
|verificationCode | 6724 |
|info |   %E6%88%91%E6%98%AF%E4%B8%AD%E6%96%87%20%E5%AD%97%E7%AC%A6 |
sign |    98c79517c47d3d3cd91b11d95e0249b2 |
ignoreSign |  source,version |

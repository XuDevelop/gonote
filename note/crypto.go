package note

import (
	"bufio"
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha512"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"math/big"
	"net"
	"os"

	"github.com/ethereum/go-ethereum/crypto/ecies"
)

//200 密码学
//加密三要素：内容、算法、密钥
//信息传输安全四要素：机密性（靠加密实现）、完整性（不可篡改性：靠单向反列函数【哈希】实现）、可验证性(通过消息验证码实现)、不可否认性(数字签名)

//201 对称加密(只有一个密码，分发密码困难)
//201.1 des 1977
//DES全称为Data Encryption Standard，即数据加密标准，是一种使用密钥加密的块算法.（已经不安全）   几分钟破解
//201.2 3des 1998
//用三个密钥加密三次，还是不安全   几小时破解
//201.3 aes
//密码学中的高级加密标准（Advanced Encryption Standard，AES），又称Rijndael加密法，是美国联邦政府采用的一种区块加密标准。

//基础知识：按位异或 ^  相同则为0，不同则为1
//把明文拆分成相同长度的区块，一般长度为32位。
//使用ECB\CBC模式,最后一个明文分组需要填充,一遍填充为需要删除的长度,如果刚好填满,也多填充一组,方便校验

//201.4 加密模式
//201.4.1 ECB(Electronic Code Book)电子密码本
//将明文依次分为固定长度到块进行加密,明文有规律则密文有规律(不推荐)//明文1>>加密>>密文1, 明文2>>加密>>密文2, 明文3>>加密>>密文3...
//201.4.2 CBC(Cipher Block Chaining)密码块链 (效率低，需要初始化向量 安全)
//通过按位异或遮蔽密文规律,弥补了ECB弊端//明文1与初始化向量按位异或>>加密>>密文1, 明文2与密文1按位异或>>加密>>密文2, 明文3与密文2按位异或>>加密>>密文3...
//201.5 CFB(Cipher FeedBack)密文反馈模式  (效率低，需要加密初始化向量 安全)
//和CBC相似//加密初始化向量>>结果与明文1按位异或>>密文1, 加密密文1>>结果与明文2按位异或>>密文2, 加密密文2>>结果与明文3按位异或>>密文3...
//201.6 OFB(Output FeedBack)输出反馈模式   (效率低，需要加密初始化向量 安全)
//和CFB相似//加密初始化向量>>结果1与明文1按位异或>>密文1, 加密结果1>>结果与明文2按位异或>>密文2, 加密结果2>>结果与明文3按位异或>>密文3...
//201.7 CTR(CounTeR)计数器模式,类似于OFB,但基于随机数  (效率低，需要加密初始化向量 安全)
//加密Seed0随机数>>结果与明文1按位异或>>密文1, 加密Seed1随机数>>结果与明文2按位异或>>密文2, 加密Seed2随机数>>结果与明文3按位异或>>密文3...

//填充最后一个明文分组
func PaddingLastBlock(plainText []byte, blockSize int) []byte {
	paddingNum := blockSize - len(plainText)%blockSize
	return append(plainText, bytes.Repeat([]byte{byte(paddingNum)}, paddingNum)...)
}

//反填充最后一个明文分组
func UnpaddingLastBlock(plainText []byte) []byte {
	return plainText[:len(plainText)-int(plainText[len(plainText)-1])]
}

//不推荐//DES-CBC加密
func DesEncrypt(plainText []byte, key []byte) (cipherText []byte, err error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return
	}
	plainText = PaddingLastBlock(plainText, block.BlockSize())
	iv := []byte("fangpigg") //初始化向量
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText = make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)
	// blockMode.CryptBlocks(plainText,plainText) //这样可以覆盖原文
	return
}

//不推荐//DES-CBC解密
func DesDecrypt(cipherText []byte, key []byte) (plainText []byte, err error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return
	}
	iv := []byte("fangpigg")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainText = make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	// blockMode.CryptBlocks(cipherText,cipherText) //这样可以覆盖原文
	plainText = UnpaddingLastBlock(plainText)
	return
}

func DesTest() {
	srcText := "芳芳是一只猪"
	fmt.Println("原文：", srcText)
	key := []byte("fangpigg")
	cipherText, err := DesEncrypt([]byte(srcText), key)
	if err != nil {
		fmt.Println("加密出错", err)
		return
	}
	fmt.Println("密文：", string(cipherText))
	plainText, err := DesDecrypt(cipherText, key)
	if err != nil {
		fmt.Println("解密出错", err)
		return
	}
	fmt.Println("解密后：", string(plainText))
}

//推荐//AES-CTR加密解密
func AesCTR(src []byte, key []byte) (dst []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	iv := []byte("qazwsxedcrfvtgby") //随机数种子
	stream := cipher.NewCTR(block, iv)
	dst = make([]byte, len(src))
	stream.XORKeyStream(dst, src)
	// blockMode.CryptBlocks(plainText,plainText) //可以这样覆盖原文
	return
}

func AesTest() {
	srcText := "芳芳是只猪"
	fmt.Println("原文：", srcText)
	key := []byte("zaqxswcdevfrbgtn") //必须为16byte
	cipherText, err := AesCTR([]byte(srcText), key)
	if err != nil {
		fmt.Println("加密出错", err)
		return
	}
	fmt.Println("密文：", hex.EncodeToString(cipherText))
	plainText, err := AesCTR(cipherText, key)
	if err != nil {
		fmt.Println("解密出错", err)
		return
	}
	fmt.Println("解密后：", string(plainText))
}

//202 非对称加密(有二个密码，公钥【上锁】加密，私钥解密【开锁】)
//202.1 RSA
//是1977年由罗纳德·李维斯特（Ron Rivest）、阿迪·萨莫尔（Adi Shamir）和伦纳德·阿德曼（Leonard Adleman）一起提出的。
//当时他们三人都在麻省理工学院工作。RSA就是他们三人姓氏开头字母拼在一起组成的.

//202.1.1 产生钥匙对儿
func GenerateRsaKey(keySize int) (err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize) //keySize通常为1024整数倍
	if err != nil {
		return
	}
	bytes := x509.MarshalPKCS1PrivateKey(privateKey) //X.509是密码学里公钥证书的格式标准
	block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: bytes,
	}
	file, err := os.Create("RSAprivateKey.pem")
	if err != nil {
		return
	}
	defer file.Close()
	pem.Encode(file, &block)
	bytes, err = x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return
	}
	block = pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: bytes,
	}
	file, err = os.Create("RSApublicKey.pem")
	if err != nil {
		return
	}
	defer file.Close()
	pem.Encode(file, &block)
	return
}

//202.1.2 RSA加密
func RsaEncrypt(plainText []byte, filePath string) (cipherText []byte, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	fileStat, err := file.Stat()
	if err != nil {
		return
	}
	buf := make([]byte, fileStat.Size())
	file.Read(buf)
	block, _ := pem.Decode(buf)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	publicKey := pubInterface.(*rsa.PublicKey)
	cipherText, err = rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	return
}

//202.1.3 RSA解密
func RsaDecrypt(cipherText []byte, filePath string) (plainText []byte, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	fileStat, err := file.Stat()
	if err != nil {
		return
	}
	buf := make([]byte, fileStat.Size())
	file.Read(buf)
	block, _ := pem.Decode(buf)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	plainText, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	return
}

func RsaTest() {
	//GenerateRsaKey(2048)
	srcText := "芳芳是只猪"
	fmt.Println("原文：", srcText)
	cipherText, err := RsaEncrypt([]byte(srcText), "RSApublicKey.pem")
	if err != nil {
		fmt.Println("加密出错", err)
		return
	}
	fmt.Println("密文：", hex.EncodeToString(cipherText))
	plainText, err := RsaDecrypt(cipherText, "RSAprivateKey.pem")
	if err != nil {
		fmt.Println("解密出错", err)
		return
	}
	fmt.Println("解密后：", string(plainText))
}

//203 ECC (误差校正码)
//是“Error Correcting Code”的简写，ECC是一种能够实现“错误检查和纠正”的技术，
//ECC内存就是应用了这种技术的内存，一般多应用在服务器及图形工作站上，可提高计算机运行的稳定性和增加可靠性。

//203.1 ECC加解密//当前版本p256以上时会报错
func ECCTest() {
	srcText := "芳芳是只猪"
	fmt.Println("原文：", srcText)
	privateKey, err := ecies.GenerateKey(rand.Reader, elliptic.P256(), nil)
	if err != nil {
		panic(err)
	}
	//ecies包可以ImportECDSAPublic()/ImportECDSA()/(&PublicKey/&PrivateKey).ExportECDSA
	cipherText, err := ecies.Encrypt(rand.Reader, &privateKey.PublicKey, []byte(srcText), nil, nil)
	if err != nil {
		fmt.Println("加密出错", err)
		return
	}
	fmt.Println("密文：", hex.EncodeToString(cipherText))
	plainText, err := privateKey.Decrypt(cipherText, nil, nil)
	if err != nil {
		fmt.Println("解密出错", err)
		return
	}
	fmt.Println("解密后：", string(plainText))
}

//204 单项散列函数
//特征：结果定长，抗碰撞，不可逆
//用途：校验数据是否被篡改，验证码，数字签名，伪随机数生成
//算法：MD4/MD5(已经不够安全，散列值长度16byte)，SHA1(已经不够安全，散列值长度20byte)，SHA2/SHA3(尚未被攻破，SHA224-28byte/SHA256-32byte/SHA512-64byte)
func Sha512Test() {
	//方法1
	srcText := "芳芳是只猪"
	fmt.Println("原文：", srcText)
	fingerprint := sha512.Sum512([]byte(srcText))
	fmt.Println("方法1指纹为：", hex.EncodeToString(fingerprint[:]))
	//方法2 //适合大文件
	srcTextbuf := make([]string, 3)
	srcTextbuf[0] = srcText[:2]
	srcTextbuf[1] = srcText[2:5]
	srcTextbuf[2] = srcText[5:]
	sha512Hash := sha512.New()
	for i := 0; i < len(srcTextbuf); i++ {
		sha512Hash.Write([]byte(srcTextbuf[i]))
	}
	res := sha512Hash.Sum(nil)
	fmt.Println("方法2指纹为：", hex.EncodeToString(res))
	//方法2结合文件操作
	file, err := os.OpenFile("note/crypto.go", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("打开文件失败 err=", err)
	}
	defer file.Close()
	buf := make([]byte, 4096)
	sha512Hash = sha512.New()
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		sha512Hash.Write(buf[:n])
	}
	res = sha512Hash.Sum(nil)
	fmt.Println("note/crypto.go 指纹为：", hex.EncodeToString(res))
}

//205 消息验证码
//205.1 生成验证码
//弊端：需要密钥分发，不能第三方校验，不能防止否认
func GenerateHmac(plainText, key []byte) []byte {
	hash := hmac.New(sha1.New, key)
	hash.Write(plainText)
	return hash.Sum(nil)
}

//205.2 校验验证码
func VerifyHmac(plainText, key, hashText []byte) bool {
	hash := hmac.New(sha1.New, key)
	hash.Write(plainText)
	return hmac.Equal(hashText, hash.Sum(nil))
}

func HashText() {
	src := []byte("芳芳是猪")
	key := []byte("pig")
	hmac1 := GenerateHmac(src, key)
	fmt.Println("校验结果为：", VerifyHmac(src, key, hmac1))
}

//206 RSA数字签名
//206.1 RSA前面
func RSASign(plainText []byte, privateKey *rsa.PrivateKey) []byte {
	hash := sha512.New()
	hash.Write(plainText)
	hashText := hash.Sum(nil)
	sigText, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, hashText)
	if err != nil {
		panic(err)
	}
	return sigText
}

//206.2 RSA签名校验
func VerifyRSASign(plainText, sigText []byte, publicKey *rsa.PublicKey) bool {
	hashText := sha512.Sum512(plainText)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA512, hashText[:], sigText)
	return err == nil
}

func RSASignTest() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048) //keySize通常为1024整数倍
	if err != nil {
		return
	}
	srcText := []byte("芳芳是大笨蛋")
	sigText := RSASign(srcText, privateKey)
	fmt.Println("校验结果为：", VerifyRSASign(srcText, sigText, &privateKey.PublicKey))
}

//207.1 ECC签名
func ECCSign(plainText []byte, privateKey *ecdsa.PrivateKey) (rText, sText []byte) {
	hashText := sha512.Sum512(plainText)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashText[:])
	if err != nil {
		panic(err)
	}
	rText, err = r.MarshalText()
	if err != nil {
		panic(err)
	}
	sText, err = s.MarshalText()
	if err != nil {
		panic(err)
	}
	return
}

//207.2 ECC签名验证
func VerifyECCSign(plainText, rText, sText []byte, publicKey *ecdsa.PublicKey) bool {
	hashText := sha512.Sum512(plainText)
	var r, s big.Int
	r.UnmarshalText(rText)
	s.UnmarshalText(sText)
	return ecdsa.Verify(publicKey, hashText[:], &r, &s)
}

func ECCSignTest() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		panic(err)
	}
	srcText := []byte("芳芳是大笨猪")
	r, s := ECCSign(srcText, privateKey)
	fmt.Println("校验结果为：", VerifyECCSign(srcText, r, s, &privateKey.PublicKey))
}

//208 证书
//解决验证签名人是签名人的问题
//单向，双向验证

//通过openssl请求自签名证书
// req -newkey rsa:2048 -nodes -keyout key.pem -x509 -days 365 -out certificate.pem

//209 TLS通讯
//传输层安全性协议（英语：Transport Layer Security，缩写作TLS），及其前身安全套接层（Secure Sockets Layer，缩写作SSL）是一种安全协议，目的是为互联网通信提供安全及数据完整性保障。
//安全传输层协议（TLS）用于在两个通信应用程序之间提供保密性和数据完整性。
//该协议由两层组成： TLS 记录协议（TLS Record）和 TLS 握手协议（TLS Handshake）

//209.1 服务器端
func TLSServerMain() {
	cert, err := tls.LoadX509KeyPair("cert/certificate.pem", "cert/key.pem")
	if err != nil {
		panic(err)
	}
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", ":2055", tlsConfig)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("连接出错", err)
			return
		}
		go func(conn net.Conn) {
			defer conn.Close()
			r := bufio.NewReader(conn)
			for {
				mes, err := r.ReadString('\n')
				if err != nil {
					fmt.Println("ReadString出错", err)
					return
				}
				fmt.Println("mes=", mes)
				n, err := conn.Write([]byte("word\n"))
				if err != nil {
					fmt.Println("conn.Write出错", err)
					return
				}
				fmt.Println("发送了", n, "个字节")
			}
		}(conn)
	}
}

// 209.2 客户端
func TLSClientMain() {
	conf := &tls.Config{
		InsecureSkipVerify: true, //本机测试不验证主机名
	}
	conn, err := tls.Dial("tcp", "localhost:2055", conf)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		panic(err)
	}
	fmt.Println("发送了", n, "个字节")
	buf := make([]byte, 4096)
	n, err = conn.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf[:n]))
}

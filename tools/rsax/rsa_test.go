package rsax

import "testing"

var (
	publicKey = `-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAK12MTd84qkCZzp4iLUj8YSUglaFMsFlv9KlIL4+Xts40PK3+wbsXPEw
cujGeUmdgMeZiK7SLLSz8QeE0v7Vs+cGK4Bs4qLtMGCiO6wEuyt10KsafTyBktFn
dk/+gBLr7B/b+9+HaMIIoJUdsFksdAg3cxTSpwVApe98loFNRfqDAgMBAAE=
-----END RSA PUBLIC KEY-----`

	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCtdjE3fOKpAmc6eIi1I/GElIJWhTLBZb/SpSC+Pl7bONDyt/sG
7FzxMHLoxnlJnYDHmYiu0iy0s/EHhNL+1bPnBiuAbOKi7TBgojusBLsrddCrGn08
gZLRZ3ZP/oAS6+wf2/vfh2jCCKCVHbBZLHQIN3MU0qcFQKXvfJaBTUX6gwIDAQAB
AoGAFwAfEo56t5JcAcLNzccQVVYj2jkbO820G8hNiSxYA5WLD0QaAxcAU/Lqqbb3
ii1aUB0ppJS13NgnU6nnGGdZzUYBG1Hai6EkVyCGrI4amQ93AaVdKncL8gJ4RZAm
YzPPUwSMEESsu24pS1NF1G1Y8C+28b/Wr0oqOsCvL6PhsMECQQDwsPJJoWRx7ZJw
E1K5KLT0cXKyrIpyXY3I6tyA5imCzOzccf3d1vDgB0L9sdSO7bG3ceSwpAeiWEbg
5jGZemPzAkEAuH6U4pEI4AMbWnatpK55Rc235NDgmT3VyIuRaKC02YXAZ+jznFep
XMd4DTli4R9r3j2YVhUpyDVbdQpFH98DMQJAQpOvcU6DSkA80WOG7lCkPTlkUKgJ
Y7kdDwZoF/+SW+vzWMbvQf3CgzV/Ak2+TgrRrbyDVZkJw45HjM4fyiRgoQJBALH/
/qlxgPyQQs3O/s2KQBsm1auAE5IF5MLuVUZ69sF/mBko2hEXSqHnGV645TuKU0pC
Zz12ga9WO3z6gaK0SaECQQDah1pKt9ViBBy4USXK3OWXEloHuTwmyr9AbLqqI5tQ
2eNuH0NkuJYQmnXmHLbKOELoYocldEBXmkzPXSN+X9kV
-----END RSA PRIVATE KEY-----`
)

func TestRsa(t *testing.T) {
	t.Log("公钥加密,私钥解密")
	decode := Encode(publicKey, "hello world")
	s := Decode(privateKey, decode)
	t.Log(s)

	t.Log("私钥加密,公钥解密")
	decode = Encode(privateKey, "hello world")
	s = Decode(publicKey, decode)
	t.Log(s)
}

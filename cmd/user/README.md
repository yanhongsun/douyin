# 随机加盐密码认证

**密码加密的原因**

为了防止数据库意外泄露/破坏和出于保护用户隐私的目的, 不应该在数据空中存入明文密码.

常用做法是通过哈希算法对明文密码加密后存入数据库. 但是人们有使用便于记忆的密码习惯, 并且不同应用也往往使用相同的密码. 因此简单的加密不能应对彩虹表攻击. 仍然存在密码泄露的风险.

可以通过向密码中加盐的方式提高密码的保护等级.

"盐"是一个随机生成的字节数组, 盐与序列化的密码相加后得到加盐密码, 增加了密码的随机性, 再对加盐密码去哈希后存入数据库.

**实现**

本项目中, 基于 `"golang.org/x/crypto/bcrypt"` 实现密码加盐加密.

- `bcrypt.GenerateFromPassword(password, cost)` 函数实现密码加盐后哈希;
- `bcrypt.CompareHashAndPassword(hashedPassword, password)` 函数实现加密密码验证.

**密码串解析**

```text
$2a$10$ESkb/bwSyISLgq1bOH0C2utXdb.hcH9oBQD1hUnfDOzm4bMKK6EX2
$ 为分隔符
2a bcrypt加密版本号
10 Cost值
ESkb/bwSyISLgq1bOH0C2utXdb 盐
hcH9oBQD1hUnfDOzm4bMKK6EX2 密码密文
```

**源码**

```go
// GenerateFromPassword returns the bcrypt hash of the password at the given
// cost. If the cost given is less than MinCost, the cost will be set to
// DefaultCost, instead. Use CompareHashAndPassword, as defined in this package,
// to compare the returned hashed password with its cleartext version.
func GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	p, err := newFromPassword(password, cost)
	if err != nil {
		return nil, err
	}
	return p.Hash(), nil
}

// CompareHashAndPassword compares a bcrypt hashed password with its possible
// plaintext equivalent. Returns nil on success, or an error on failure.
func CompareHashAndPassword(hashedPassword, password []byte) error {
	p, err := newFromHash(hashedPassword)
	if err != nil {
		return err
	}

	otherHash, err := bcrypt(password, p.cost, p.salt)
	if err != nil {
		return err
	}

	otherP := &hashed{otherHash, p.salt, p.cost, p.major, p.minor}
	if subtle.ConstantTimeCompare(p.Hash(), otherP.Hash()) == 1 {
		return nil
	}

	return ErrMismatchedHashAndPassword
}
```

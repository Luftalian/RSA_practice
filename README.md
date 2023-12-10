# RSA暗号

RSA暗号は、公開鍵暗号の一種である。公開鍵暗号とは、暗号化と復号に異なる鍵を用いる暗号のことである。RSA暗号は、1977年にロナルド・リベスト、アディ・シャミア、レナード・アドルマンによって発明された。

## アルゴリズム

RSA暗号は、以下の手順で暗号化と復号を行う。

### 鍵の生成

1. 2つの素数 $p$ と $q$ を選ぶ。
2. $n=pq$とする。
3. $\phi(n)=(p-1)(q-1)$とする。
4. $\phi(n)$と互いに素な $e$ を選ぶ。
5. $ed\equiv 1\pmod{\phi(n)}$となる$d$を求める。
$d\times e = x\times \phi(n) + 1$となる整数 $d,x$ が存在する。

このとき、 $(n,e)$ が公開鍵、 $d$ が秘密鍵である。

### 暗号化

平文を $a$ とし、 $0\leq a< n$ の整数とする。

$a^e\equiv b\pmod{n}$となる $b$ を求める。このときの $b$ が暗号文である。ここで、 $b$ は $0\leq b< n$ を満たす整数である。

### 復号

$b^d\equiv a^{\prime}\pmod{n}$となる $a^{\prime}$ を求める。このとき、 $a^{\prime}=a$ となり、平文に復号できる。この理由については、以下の[証明](https://github.com/Luftalian/RSA_practice#証明)を参照。

## 証明

RSA暗号が完全性を持つことを証明する。

復号化の過程より

$a^{\prime}\equiv b^d \pmod{n} = (a^e)^d = a^{ed} = a^{1+x\phi(n)}$

となる。

(i) $a$と $n$ が互いに素のとき、 $a$ と $p,q$ も互いに素であるから、フェルマーの小定理より

$a^{\phi(n)} = (a^{p-1})^{q-1}\equiv 1^{q-1} = 1\pmod{p}$

$a^{\phi(n)} = (a^{q-1})^{p-1}\equiv 1^{p-1} = 1\pmod{q}$

となる。

よって、 $a^{\phi(n)}\equiv 1\pmod{n}$ である。

以上より

$a^{\prime}\equiv a^{1+x\phi(n)}\equiv a\times (a^{\phi(n)})^x\equiv a\times 1^x\equiv a\pmod{n}$

となり、 $0 \leq a^{\prime}, a < n$ であるから、 $a^{\prime}=a$ である。

(ii) $a$と $n$ が互いに素でないとき、 $a$ と $p,q$ も互いに素でない。よって、 $a$ は $p$ の倍数または $q$ の倍数である。つまり、 $a = kp$ または $a = kq$ となる整数 $k$ が存在する。

(ii-1) ここで、 $a$ が $p$ の倍数だが、 $q$ の倍数でないとする。このとき、 $a\equiv 0\pmod{p}$ であり、 $a\equiv a\pmod{q}$ である。

また、 $a$ と $q$ は互いに素であるから、[フェルマーの小定理](https://ja.wikipedia.org/wiki/フェルマーの小定理)より

$a^{\prime} = a\times (a^{p-1})^{(q-1)x}\equiv a\times 1^{(q-1)x} = a\pmod{q}$

となる。さらに、

$a^{\prime} = a\times (kp)^{x\phi(n)}\equiv a\times 0^{p-1} = 0\pmod{p}$

となる。以上の結果より[中国の剰余定理](https://ja.wikipedia.org/wiki/中国の剰余定理)の条件を満たし、 $a^{\prime}$ と $a$ が $n(=pq)$ を法として一意に決まることがわかるので、 $a^{\prime}=a$ である。

(ii-2) $a$が $q$ の倍数だが、 $p$ の倍数でないとする。このとき、 $a\equiv a\pmod{p}$ であり、 $a\equiv 0\pmod{q}$ である。

(ii-1)と同様に、$a^{\prime}=a$である。

(ii-3) $a$が $p$ の倍数かつ $q$ の倍数であるとする。このとき、 $a=kpq$ となる整数 $k$ が存在する。つまり、 $a\equiv 0\pmod{n}$ である。

このとき、

$a^{\prime} = a\times (a^{p-1})^{(q-1)x}\equiv a\times 0^{(q-1)x} = 0\pmod{q}$

$a^{\prime} = a\times (a^{q-1})^{(p-1)x}\equiv a\times 0^{(p-1)x} = 0\pmod{p}$

であるから、 $a^{\prime}\equiv 0\pmod{n}$ である。よって、 $a^{\prime}=a$ である。

以上より、 $a^{\prime}=a$ である。

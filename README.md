# feistel

[![codecov](https://codecov.io/gh/rogercoll/feistelnetwork/branch/master/graph/badge.svg)](https://codecov.io/gh/rogercoll/feistelnetwork)

In cryptography, a Feistel cipher is a symmetric structure used in the construction of block ciphers,
named after the German-born physicist and cryptographer Horst Feistel who did pioneering research while working for IBM (USA); // part. To reconstruct the original secret, a minimum number of parts is required. In the threshold
it is also commonly known as a Feistel network. A large proportion of block ciphers use the scheme, including the Data Encryption Standard (DES).
The Feistel structure has the advantage that encryption and decryption operations are very similar, even identical in some cases,
requiring only a reversal of the key schedule.
Therefore, the size of the code or circuitry required to implement such a cipher is nearly halved.
See [wiki page](https://en.wikipedia.org/wiki/Feistel_cipher).

## Examples

### New

```golang
ceasar := func(text []rune) []rune {
    //for quick test we set the same shift = 3
    result := make([]rune, len(text))
    _ = make([]rune, 1024)
    for i, t := range text {
        s := int(t) + 3
        if s > 'z' {
            s = s - 26
        } else if s < 'a' {
            s = s + 26
        }
        result[i] = rune(s)
    }
    return result
}
p, err := New("hell", 9)
if err != nil {
    log.Fatal(err)
}
encrypted, err := p.Run(ceasar)
if err != nil {
    log.Fatal(err)
}
message, err := p.Reverse(ceasar)
if err != nil {
    log.Fatal(err)
}
fmt.Println(encrypted)
fmt.Println(message)
```

 Output:

```
[102 22 8 55]
hell
```

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)

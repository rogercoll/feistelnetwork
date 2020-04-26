# feistel

See [wiki page](https://en.wikipedia.org/wiki/Feistel_cipher).

#### Examples

##### Run

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
p, err := New("hell", 1)
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
[7 10 108 108]
hell

```


---

Created by [goreadme](https://github.com/apps/goreadme)

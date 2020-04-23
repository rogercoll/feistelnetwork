# feistel

See [wiki page](https://en.wikipedia.org/wiki/Feistel_cipher).

#### Examples

##### Run

```golang
ceasar := func(text []rune) []rune {
    //for quick test we set the same shift = 3
    result := make([]rune, len(text))
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
encrypted, err := Run("hell", 1, ceasar)
if err != nil {
    log.Fatal(err)
}
fmt.Println(encrypted)
```

 Output:

```
[7 10 108 108]

```


---

Created by [goreadme](https://github.com/apps/goreadme)

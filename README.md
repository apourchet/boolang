# Boolang
Boolang is a library for parsing your own boolean expressions. It takes care of operation precedence and tokenization,
leaving you with only statement evaluation logic to write.

## Example
```go
tree, _ := boolang.Parse("(isBoolangAwesome && isBoolangSimple) || !hasLazyEvaluation")
fn := func(l *boolang.Leaf, _ ...interface{}) (bool, error) {
    content := strings.Trim(l.Content, " ")
    if content == "isBoolangAwesome" {
        return true, nil
    } else if content == "isBoolangSimple" {
        return true, nil
    } else if content == "hasLazyEvaluation" {
        return true, nil
    }
    return false, nil
}
result, _ := tree.Eval(fn) // result == true
```

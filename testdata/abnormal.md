# this is test


## include syntax error

```hcl
resource "aws_vpc" "main" {
  cidr_block       = "10.0.0.0/16"
        instance_tenancy = "dedicated"

tags = {
    Name = "main"
  }
  hoge
}
```

## code block (not fenced code)

`a:=1+1`

`println!("Hello, world!");`

`Point::new(1.0, 1.0)`

## other language (no format)

```go
func main(){
fmt.Println("test!!")
}
```

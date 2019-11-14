## fenced code(hcl)

```hcl
resource "aws_vpc" "main" {
  cidr_block       = "10.0.0.0/16"
  instance_tenancy = "dedicated"

  tags = {
    Name = "main"
  }
}
```

## fenced code(hcl-terraform)

```hcl-terraform
resource "aws_vpc" "main" {
  cidr_block       = "10.0.0.0/16"
  instance_tenancy = "dedicated"

  tags = {
    Name = "main"
  }
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
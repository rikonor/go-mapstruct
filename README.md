mapstruct
---

Like [mitchellh/mapstructure](https://github.com/mitchellh/mapstructure) but based on a simple round-trip conversion to json.

### Usage

```go
type Person struct {
  Name string `json:"name"`
}

// We have a map that's compatible with the struct
data := map[string]interface{}{
  "name": "Johnny",
}

// We want to convert the map to the struct
var p Person
if err := Decode(data, &p); err != nil {
  log.Fatal(err)
}
```

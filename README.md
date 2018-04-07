# zippy

Zippy is intended to be a simple way to let you compress and decompress []byte streams in different compression formats.

It is useful for giving you the option to switch compression types by changing a very simple config.  E.g. if you need 
more disk space for your redis caching, but don't want to pay for a bigger AWS instance... you could easily swap a less 
CPU-performant compression format for a more disk-efficient format with a simple config change.

Eventual goal is to support all the compression formats, but for now it supports:
   * snappy
   * gzip
   * none

## Getting Started

Initialise zippy by choosing your compression format:
```
    zpy := zippy.New(zippy.Config{
        CompressionFormat: "snappy",
        //CompressionFormat: "gzip",
        //CompressionFormat: "none",
    })
    
    //consts such as zippy.COMPRESSION_SNAPPY exist if you prefer that.
```

Take your []byte stream, and compress it like so:
```
    content := []byte("test string for compression")
    cmpr := zpy.Compress(content)
```

Now you can decompress the compressed content like so:
```
    dcmp, _ := zpy.Decompress(cmpr)
    fmt.Println("Decompressed string: ", string(dcmp[:]))

```

A more practical example would be if you were to say, change your redigo GET/SET wrappers.

So if your redigo SET methods looked something like this: 
```
    _, err := conn.Do("SET", key, content, "EX", 300)
```
    
You would swap out the content []byte stream like so:
``` 
    cmprContent, _ := zpy.Compress(content)
    _, err := conn.Do("SET", key, cmprContent, "EX", 300)
```
      
Likewise, for your  "GET" redigo wrappers, you would use zpy.Decompress() instead of zpy.Compress().
    
Simple!  You can `redis-cli --bigkeys` afterwards to compare disk space before and after. 

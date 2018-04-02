# zippy

Zippy is intended as a simple way to let you switch between different compression / decompression formats for []byte
streams.

I started it as my use case required a CPU performant way to reduce redis disk size, which required experimenting with
the different formats such as Snappy and Gzip to find the best fit. Hopefully making a module that does this for you
makes life a little easier for you.

Eventual goal is to support all the compression formats, but for now it supports:
    - snappy
    - gzip
    - none    

USAGE:    

	content := []byte("test string for compression")

	zippy := zippy.New(zippy.Config{
		CompressionFormat: "snappy",
	})
		
	//Zipping
	zp := zippy.Zip(content)

	//Unzipping
	uz, _ := zippy.Unzip(zp)

	fmt.Println("unzipped string:", string(uz[:]))

    ________________________________________________
    

    A more practical example would be if you were to say, change your redigo wrappers 
    E.g. your SET methods could be changed like so:
    //BEFORE
    _, err := conn.Do("SET", key, content, "EX", 300)
    

    //AFTER!
    zippy := zippy.New(zippy.Config{
        CompressionFormat: "snappy",
    })
    	
    _, err := conn.Do("SET", key, zippy.Zip(content), "EX", 300)
     
    And you would correspondingly replace "GET" redis commands to use zippy.Unzip().
    The only difference is that unzipping also returns a second err param, so that 
    needs to be handled first.
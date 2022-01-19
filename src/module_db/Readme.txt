For debugging post request:
https://httpbin.org/post

// ====================================================


1. Create Aquila db:
-----------------------
    
	{
		"data": {
			"schema": {
				"description": "this is my database",
				"unique": "r8and0mseEd901",
				"encoder": "strn:msmarco-distilbert-base-tas-b",
				"codelen": 768,
				"metadata": {
					"name": "string",
					"age": "number"
				}
			}
		},
		"signature": "secret"
	}

Example struct:
---------------------

createDb := &CreateDbStruct{
    Data: DataStructCreateDb{
        Schema: SchemaStruct{
            Description: "this is my database",
            Unique:      "r8and0mseEd901",
            Encoder:     "strn:msmarco-distilbert-base-tas-b",
            Codelen:     768,
            Metadata: MetadataStructCreateDb{
                Name: "string",
                Age:  "number",
            },
        },
    },
    Signature: "secret",
}

data, err := json.Marshal(createDb)
if err != nil {
    log.Fatal(err)
}

// another way to pass reader to post
// reader := bytes.NewReader(data)


// ====================================================





// ====================================================
// ====================================================
// ====================================================
// ====================================================
// ====================================================
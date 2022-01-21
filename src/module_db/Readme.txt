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

2. Doc Insert:

{
   "data": {
       "docs": [
           {
           "payload":
               {
                   "metadata": {
                       "name":"name1",
                       "age": 20
                   },
                   "code": [0.1, 0.2, 0.3]
               }
           },
           {
           "payload":
               {
                   "metadata": {
                       "name":"name2",
                       "age": 30
                   },
                   "code": [0.4, 0.5, 0.6]
               }
           }
       ],
       "database_name": "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7"
   },
   "signature": "secret"
}

Example struct:
-------------------

docInsert := &DocInsertStruct{
		Data: DatatDocInsertStruct{
			Docs: []DocsStruct{
				{
					Payload: PayloadStruct{
						Metadata: MetadataStructDocInsert{
							Name: "name1",
							Age:  20,
						},
						Code: []float32{0.1, 0.2, 0.3},
					},
				},
				{
					Payload: PayloadStruct{
						Metadata: MetadataStructDocInsert{
							Name: "name2",
							Age:  20,
						},
						Code: []float32{0.4, 0.5, 0.6},
					},
				},
			},
			DatabaseName: "BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7",
		},
		Signature: "secret",
	}

	data, err := json.Marshal(docInsert)
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
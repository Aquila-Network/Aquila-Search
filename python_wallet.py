from Crypto.Hash import SHA384
from Crypto.PublicKey import RSA
from Crypto.Signature import pkcs1_15
import base58
import bson

priv_key = None
key_file = "/home/dev/aquilax/ossl/private_unencrypted.pem"
# key_file = "private_unencrypted.pem"
with open(key_file, "r") as pkf:
    k = pkf.read()
    priv_key = RSA.import_key(k)

data_ = {
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
        }

data_bson = bson.dumps(data_)

# generate hash
hash = SHA384.new()
hash.update(data_bson)

# Sign with pvt key
signer = pkcs1_15.new(priv_key)
signature = signer.sign(hash)
signature = base58.b58encode(signature).decode("utf-8")

print(data_bson)
print("")
print(signature)

# run 
# $ python3 python_json_to_bson.py
#
# install tool
# $ sudo apt install python3-pip
# 
# install libraries:
# $ pip3 install pycryptodome
# $ pip3 install base58
# $ pip3 install bson

# Response
# 5JNkAvWL6AdtbYmNQrBUxWubEDQp6bZzw3zf8djxSRzg8pywE5wgs8ufj7TrJ6uK71wDZTWEcTJYT4bDU5v5nQgY5aopwvRySmBmu1HYLYBGTNRQTQ81QRpjF1pvxBtHn1NrWipTo6kxVvJZURY9kfwv9oweTvv2Xt6FAPx9VaZgo8g6YTxEBLyVnrhvsKE7QmZuUd4sSgXQ3zm9fYB3f6v5RdprtCuUqUEAzs7ZecGSvDzc8kNMF7DkQajJnjbcYPQAiNMgaPtnrrArofYNjwxsecZngA7eTU1d5XQwnWAgKvpJ9xMXNRg41uz4CRpUyDs8YFNtvMQbpLkc5MJA2UgrBcxscc
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

data_ = {"matrix":[[-0.01806008443236351,-0.17380790412425995,0.03992759436368942,0.43514639139175415]],"k":10,"r":0,"database_name":"BN4Bik3RbaY5mzJS94u8SvjZd1keyjTWaDNF36TjYzj7"}

data_bson = bson.dumps(data_)

# generate hash
hash = SHA384.new()
hash.update(data_bson)

# Sign with pvt key
signer = pkcs1_15.new(priv_key)
signature = signer.sign(hash)
signature = base58.b58encode(signature).decode("utf-8")

print(data_bson.hex()) # print in hexidecimal
print("")
print(signature)

# run 
# $ python3 python_wallet.py
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
# Golang-TLS-Example
Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)

openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

RSA recommendation key ≥ 2048-bit

openssl req -x509 -nodes -newkey rsa:2048 -keyout server.rsa.key -out server.rsa.crt -days 3650

ln -sf server.rsa.key server.key

ln -sf server.rsa.crt server.crt

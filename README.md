Simpe socks5 proxy

# Build
`docker build -t gilaz/simple-socks5 .`

# Start
`docker run -d -e PROXY_USER=gilaz -e PROXY_PASSWORD=secret -p 1080:1080 gilaz/simple-socks5`

# Test
`curl --socks5 127.0.0.1:1080 -U gilaz:secret https://yandex.ru`

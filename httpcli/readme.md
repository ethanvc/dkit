使用http库的场景：
1. 使用json协议发送请求，对body进行签名，签名信息放到header中，返回信息也要验证签名。
3. 使用query-string携带参数，返回是json格式的数据。
4. 路径中携带参数，返回xml格式的数据。
5. 路径中携带参数，get请求，返回文件。
6. 发送表单（form-data 或 x-www-form-urlencoded）数据，常用于文件上传或表单提交，返回 json 或文本格式。
7. 需要自定义请求头（如 User-Agent、Authorization、Cookie 等）进行身份认证或模拟浏览器行为。
8. 处理重定向（3xx 响应），自动跟随或自定义重定向策略。
9. 需要设置超时、重试机制，保证请求的健壮性。
10. 进行并发/批量请求，提升数据抓取或接口测试效率。
11. 处理大文件下载或分块上传，支持流式处理。
12. 需要处理 HTTPS 证书校验、忽略证书错误或自定义 CA 证书。
13. 需要代理（如 HTTP/HTTPS/SOCKS5）访问目标服务。
14. 需要记录和调试请求/响应的详细日志，便于排查问题。
15. 处理多部分响应（如 chunked transfer encoding）或 SSE（Server-Sent Events）等流式数据。
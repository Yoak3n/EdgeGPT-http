<p align="center" >
    <h2 align="center">EdgeGPT-http</h2>
    <h3 align="center">a http server use go<br/>
    </h3>
    <p align="center">just a practice work for self-entertainment. </p>
<p>

### Usage
Find the configuration file ```config.example.yaml``` and rename it as ```config.yaml```, then modify the configuration inside.

1. Start the service by using the following command:
```bash
go run .\main.go
```

2. Client post the ```url/chat``` with a json data like the following:
```json
{
  "name": "test",
  "style": "bing-c",
  "question": "hello bing"
}
```
One ```name```means one conversation (currently there is no function to reset conversation)

Then the client will get a json response:
```json
{
  "status":  "success",
  "style":   "bing-c",
  "message": "bing's answer"
}
```

### Problem
- [x] multiple sessions
- [ ] commands for conversation like ```reset```or more
- [ ] source auto release
- [ ] high concurrency

### Reference
At the very beginning I use the Python lib [EdgeGPT](https://github.com/acheong08/EdgeGPT) in my [bing-qqbot](https://github.com/Yoak3n/bing-qqbot),it's too complicated to start.So I planed to convert it to golang at some point,hover thank [billikeu](https://github.com/billikeu/billikeu) that his repository [Go-EdgeGPT](https://github.com/billikeu/Go-EdgeGPT) really helped me(btw my plan is also named ```go-EdgeGPT```)lol   
For the convenience of development, it is temporarily included as part of this project, and may be submitted as a PR in the future

And the idea came from the http service of [chatgpt-mirai-qq-bot](https://github.com/lss233/chatgpt-mirai-qq-bot)

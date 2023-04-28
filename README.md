<p align="center" >
    <h2 align="center">EdgeGPT-http</h2>
    <h3 align="center">a http server based on go<br/>
    </h3>
    <p align="center">just a practice work for self-entertainment. </p>
<p>

### Download
```bash
git clone https://github.com/Yoak3n/EdgeGPT-http.git
```
or

[Download ZIP](https://github.com/Yoak3n/EdgeGPT-http/archive/refs/heads/main.zip) 

### Usage

#### Basic(manually)

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
One ```name```means one conversation  
Then the client will get a json response:
```json
{
  "status":  "success",
  "style":   "bing-c",
  "message": "bing's answer",
  "count": {
    "currentNum":1,
    "maxNum":20
  }
}
```

<details>
<summary>Docker</summary>

**Must installed  ```Docker```**

1. Pull the image by:
```bash
docker pull yoaken/edgegpt-http:latest 
```
2.Run the image:
```bash
docker run -v 'your absolute path of config.yml:/app/config.yml' -v'your absolute path of cookies.json:/app/cookies.json' -p "8080:8080" yoaken/edgegpt-http:latest 
```

</details>

<details>
<summary>docker-compose<b>(recommended)</b></summary>

**Must installed ```Docker``` and ```docker-compose```**

1. Copy the file ```docker-compose.yml``` to the dir path as you like
2. Create your ```config.yml``` and ```cookies.json``` in the path the same to the above
3. Change the work dir where these files exist and run with a single command:
```bash
docker-compose up -d  
```
</details>

### Problem
- [x] multiple sessions
- [x] commands for conversation like ```reset```or more
- [ ] source auto release
- [x] high concurrency(perhaps ```gin``` already supported)

### Reference
At the very beginning I use the Python lib [EdgeGPT](https://github.com/acheong08/EdgeGPT) in my [bing-qqbot](https://github.com/Yoak3n/bing-qqbot),it's too complicated to start.So I planed to convert it to golang at some point,however thank [billikeu](https://github.com/billikeu) that his repository [Go-EdgeGPT](https://github.com/billikeu/Go-EdgeGPT) really helped me(btw my plan is also named ```go-EdgeGPT```)XD   
For the convenience of development, it is temporarily included as part of this project, and may be submitted as a PR in the future

And the idea came from the http service of [chatgpt-mirai-qq-bot](https://github.com/lss233/chatgpt-mirai-qq-bot)

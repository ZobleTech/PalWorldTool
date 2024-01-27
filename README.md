# **使用方法**

Star！Plz！

## 传统方式

### **启动WEB**服务

> 在启动前端项目之前需要先安装[Node.JS](https://nodejs.org/en)

1、进入`PalworldTools\Web` 目录

2、运行`npm install`

```js
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      "/api": {
        target: "http://localhost:8080", // 修改为你当前的后端API的地址
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ""),
      },
    }
  }
})
```

3、运行`npm run dev`



### 启动后端服务器

1、进入`PalworldTools\Server` 目录

2、运行 `go run .`

> 如果运行失败则可能先运行 `go mod tidy`



## 在Docker中使用

```dockerfile
version: '3'
services:
  web:
    image: registry.cn-hangzhou.aliyuncs.com/zobel/palworldtools:web-v1.0.2
    container_name: web
    ports:
      - "80:80"
    networks:
      - main # 填入一个docker网络名称

  server:
    image: registry.cn-hangzhou.aliyuncs.com/zobel/palworldtools:server-v1.0.1
    container_name: server
    ports:
      - "8080:8080"
    networks:
      - main # 填入一个docker网络名称
    depends_on:
      - web

networks:
  main: # 需要新建一个外部docker网络
    external: true
```


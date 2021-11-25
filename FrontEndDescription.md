# Ant Design Pro

This project is initialized with [Ant Design Pro](https://pro.ant.design). Follow is the quick guide for how to use.

## Environment Prepare

Install `node_modules`:

```bash
npm install
```

or

```bash
yarn
```

## Provided Scripts

Ant Design Pro provides some useful script to help you quick start and build with web project, code style check and test.

Scripts provided in `package.json`. It's safe to modify or add additional script:

### Start project

```bash
npm start
```

### Build project

```bash
npm run build
```

### Check code style

```bash
npm run lint
```

You can also use script to auto fix some lint error:

```bash
npm run lint:fix
```

### Test code

```bash
npm test
```

## More

You can view full document on our [official website](https://pro.ant.design). And welcome any feedback in our [github](https://github.com/ant-design/ant-design-pro).

## APIS

**though may not likely be used, pls add a "msg" param in the response as the comment of itslef**

- Video upload: when a project's video rendered, upload its url via (https://qcmt57.fn.thelarkcloud.com/createVideo),not commented para pls keep it as it is.

```json
{
  "controller": "1612779773437",
  "userId": "1612779773437",
  "videoName": "搞笑视频2", //视频名称
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxNjEyNzc5NzczNDM3IiwiaWF0IjoxNjEyNzgwMDE0fQ.J27ujArwYmr2b7Muv2wI3FEs1YbXO8Ce2llju6dMzjo",
  "url": "https://sf1-hscdn-tos.pstatp.com/obj/media-fe/xgplayer_doc_video/mp4/xgplayer-demo-360p.mp4" //视屏地址
}
```

- Video Clip effect: give video clip a effect, and return its pulic resource link:
  - post:{file:[buffer],effect:[string]}
  - respond:

````json
   {
     "isOnline": true,
     "isFx": true,
     "effect": "magic",
     "onLineId": "1", //use this to identify fxed clip on the the server
     "onLineUrl": "https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.webm",
     "coverUrl": "https://s3-lc.thelarkcloud.com/obj/larkcloud-mgcloud/baas/qcmt57/0bb0008f264e7ffb_1613939141653.jpeg"
   }
   ```

````

- project video render
  - upload all clips->respond with url(no effect needed) ->client post the project timeline->combine, render,upoload,and return the final film url post:  
    { file:[buffer],effect:"none" } respond:

```json
{
  "isOnline": true,
  "isFx": false,
  "effect": "none",
  "onLineId": "1", //use this to identify fxed clip on the the server
  "onLineUrl": "https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.webm",
  "coverUrl": ""
}
```

finnal post:

```json
{
  "project": "video's name",
  "timeline": [
    {
      "index": 0,
      "onLineId": "1"
    }
  ]
}
```

response:  
 { "msg": "rendered and posted", "onLineUrl": "https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.webm" }

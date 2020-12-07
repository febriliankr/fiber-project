# Fiber (Go) and NextTS

## Deployment

https://dev.to/heavykenny/how-to-deploy-a-golang-app-to-heroku-5g1j


https://medium.com/jaysonmulwa/deploying-a-go-fiber-go-react-app-to-heroku-using-docker-7379ed47e0fc

```
heroku login
heroku create app-name
heroku container:login OR docker login
heroku container:push web --app [YOUR_APP_NAME]
```

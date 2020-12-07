# Fiber (Go) and NextTS

Fiber is a golang framework, and the frontend is deployed with nextjs with typescript

## Demo

https://fiber-golang.herokuapp.com/

## Todo

api/getUser/:param is working
api/getAllUsers is not working because mysql has not been deployed

## Setting up the database

https://devcenter.heroku.com/articles/cleardb#:~:text=Provisioning%20the%20dedicated%20MySQL%20add%2Don,-Existing%20ClearDB%20shared&text=Once%20in%20the%20ClearDB%20portal,to%20build%20your%20new%20DATABASE_URL.

### Nyampah Code

INSERT INTO `users` (`id`, `name`, `email`, `password`) VALUES
(1, 'Febrilian Kristiawan', 'febrilian@outlook.com', 'febrilian');

# Golang REST API application for management of movie ratings and personal movie lists (favorites, wishlist) with data storage in PostgreSQL

## Stack

 - Go 1.18 
 - PostgreSQL 
 - Docker
 - JWT
 - Redis

 # API endpoints

These endpoints allow you to handle Auth and Films management.

## POST
`SECURED:false` /auth/signup

`SECURED:false` /auth/signin

`SECURED:true` /auth/signout

`SECURED:true` /films/film
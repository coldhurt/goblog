# goblog 

A simple blog under development using gin + mongodb

# curl test

## auth
```bash
# login
curl -X POST -H "content-type: application/json" -d '{"username":"admin","password":"pwd"}' http://localhost:8080/admin/login

# update password
curl -X POST -H "content-type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjI5MjA3MjEyLCJpYXQiOjE2MjkwMzQ0MTIsImlzcyI6IkJpa2FzaCJ9.OtP8G8peOPkZjm5i-7rmRLYod9y4aBmfZ0Ywd0OWQ8c" -d '{"password":"12345678", "new_password":"12345678"}' http://localhost:8080/admin/update
```
## article
```bash
# create article
curl -X POST -H "content-type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjI5MjA3MjEyLCJpYXQiOjE2MjkwMzQ0MTIsImlzcyI6IkJpa2FzaCJ9.OtP8G8peOPkZjm5i-7rmRLYod9y4aBmfZ0Ywd0OWQ8c" -d '{"title":"test","content":"test"}' http://localhost:8080/article/create

# list article
curl -X POST -H "content-type: application/json" -H "Authorization: token" -d '{"page_no":1,"page_size":10}' http://localhost:8080/article/list

# update article
curl -X PUT -H "content-type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjI5MjA3MjEyLCJpYXQiOjE2MjkwMzQ0MTIsImlzcyI6IkJpa2FzaCJ9.OtP8G8peOPkZjm5i-7rmRLYod9y4aBmfZ0Ywd0OWQ8c" -d '{"id":"611a7a5fdb720e9bdc95c1e3", "title":"test","content":"test"}' http://localhost:8080/article/edit

# delete article
curl -X DELETE -H "content-type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjI5MjA3MjEyLCJpYXQiOjE2MjkwMzQ0MTIsImlzcyI6IkJpa2FzaCJ9.OtP8G8peOPkZjm5i-7rmRLYod9y4aBmfZ0Ywd0OWQ8c" -d '{"id":"611a7a5fdb720e9bdc95c1e3"}' http://localhost:8080/article/remove
```


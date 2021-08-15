# goblog 

gin + mongodb 

# curl test

```bash
# login
curl -X POST -H "content-type: application/json" -d '{"username":"admin","password":"pwd"}' http://localhost:8080/admin/login

# create article
curl -X POST -H "content-type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjI5MjA3MjEyLCJpYXQiOjE2MjkwMzQ0MTIsImlzcyI6IkJpa2FzaCJ9.OtP8G8peOPkZjm5i-7rmRLYod9y4aBmfZ0Ywd0OWQ8c" -d '{"title":"test","content":"test"}' http://localhost:8080/article/create

# list article
curl -X POST -H "content-type: application/json" -H "Authorization: token" -d '{"page_no":1,"page_size":10}' http://localhost:8080/article/list
```
const host = 'http://localhost:8080'

const Post = (url: string, init: { body: Object }) => {
  return fetch(`${host}/${url}`, {
    ...init,
    body: JSON.stringify(init.body || {}),
    method: 'post',
    headers: {
      'content-type': 'application/json',
    },
  }).then((res) => res.json())
}

const Get = (url: string, init = {}) => {
  return fetch(`${host}/${url}`, {
    ...init,
    method: 'get',
    headers: {
      'content-type': 'application/json',
    },
  }).then((res) => res.json())
}

export { Post, Get }

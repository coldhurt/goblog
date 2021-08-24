const host = 'http://localhost:8080'

const Post = (url: string, init: RequestInit) => {
  return fetch(`${host}/${url}`, {
    ...init,
    method: 'post',
    headers: {
      'content-type': 'application/json',
    },
  })
}

export { Post }

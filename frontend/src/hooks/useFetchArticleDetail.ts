import { Get } from '../utils/fetch'
import { Article } from '../store/articleSlice'
import { useEffect, useState } from 'react'

const useFetchArticleDetail = (id: string) => {
  const [loading, setLoading] = useState(false)
  const [article, setArticle] = useState<Article | null>(null)
  const refresh = () => {
    setLoading(true)
    Get(`article/detail/${id}`)
      .then((res) => {
        if (res.data) {
          setArticle(res.data)
        }
      })
      .finally(() => setLoading(false))
  }

  useEffect(() => {
    refresh()
  }, [])

  return {
    loading,
    article,
    refresh,
  }
}

export default useFetchArticleDetail

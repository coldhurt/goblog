import { useDispatch, useSelector } from 'react-redux'
import { Post } from '../utils/fetch'
import { RootState } from '../store'
import { startLoading, setArticles } from '../store/articleSlice'
import { useEffect } from 'react'

const useFetchArticles = () => {
  const article = useSelector((state: RootState) => state.article)
  const dispatch = useDispatch()
  const refresh = () => {
    dispatch(startLoading())
    Post('article/list', {
      body: {},
    }).then((res) => {
      if (res.data) {
        dispatch(setArticles({ articles: res.data }))
      }
    })
  }

  useEffect(() => {
    refresh()
  }, [])

  return {
    article,
    refresh,
  }
}

export default useFetchArticles

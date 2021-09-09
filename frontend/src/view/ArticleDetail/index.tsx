import { Box, Container, CssBaseline } from '@material-ui/core'
import React from 'react'
import { useParams } from 'react-router-dom'
import useFetchArticleDetail from '../../hooks/useFetchArticleDetail'

const ArticleDetail = () => {
  const { id } = useParams<{ id: string }>()
  const { article, loading } = useFetchArticleDetail(id)
  return (
    <Container>
      <CssBaseline />
      <Box
        sx={{
          marginTop: 8,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        {loading ? (
          <div>loading</div>
        ) : article ? (
          <div>
            <div>{article.title}</div>
            <div>{article.content}</div>
          </div>
        ) : (
          <div>no data</div>
        )}
      </Box>
    </Container>
  )
}

export default ArticleDetail

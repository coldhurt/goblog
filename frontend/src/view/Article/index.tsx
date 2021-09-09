import { Box, Container, CssBaseline, List, ListItem, ListItemText } from '@material-ui/core'
import React from 'react'
import useFetchArticles from '../../hooks/useFetchArticles'
import { Link } from 'react-router-dom'
import { Article } from '../../store/articleSlice'
import { HourglassEmpty } from '@material-ui/icons'
import Spin from '../../components/Spin'

const ListItemLink: React.FC<Article> = ({ title, id }) => {
  const CustomLink: React.FC<{ to: string }> = (props) => <Link {...props} />

  return (
    <li>
      <ListItem button component={CustomLink} to={`/article/${id}`}>
        <ListItemText primary={title} />
      </ListItem>
    </li>
  )
}

const ArticleList = () => {
  const { article } = useFetchArticles()
  const articles = article.articles
  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <Box
        sx={{
          marginTop: 8,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        <List>
          {article.loading ? (
            <Spin />
          ) : articles.length > 0 ? (
            articles.map((item) => <ListItemLink key={item.id} {...item} />)
          ) : (
            <HourglassEmpty />
          )}
        </List>
      </Box>
    </Container>
  )
}

export default ArticleList

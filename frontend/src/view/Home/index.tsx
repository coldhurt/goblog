import React from 'react'
import { Helmet } from 'react-helmet'
import ArticleList from '../Article'

const Home = () => {
  return (
    <div>
      <Helmet>
        <title>Goblog</title>
      </Helmet>
      <ArticleList />
    </div>
  )
}
export default Home

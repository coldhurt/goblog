import React, { lazy, Suspense } from 'react'
import './App.css'
import { BrowserRouter as Router, Switch, Route, useLocation } from 'react-router-dom'
import { Container, createTheme, CssBaseline, ThemeProvider } from '@material-ui/core'
import Spin from './components/Spin'

const Home = lazy(() => import('./view/Home'))
const Login = lazy(() => import('./view/Login'))
const Article = lazy(() => import('./view/Article'))
const ArticleDetail = lazy(() => import('./view/ArticleDetail'))

function NoMatch() {
  let location = useLocation()
  return (
    <div>
      <h3>
        No match for <code>{location.pathname}</code>
      </h3>
    </div>
  )
}

const theme = createTheme({
  palette: {
    mode: 'dark',
  },
})

function App() {
  return (
    <ThemeProvider theme={theme}>
      <Container>
        <CssBaseline />
        <Suspense fallback={<Spin />}>
          <Router>
            <Switch>
              <Route path="/" exact>
                <Home />
              </Route>
              <Route path="/article" exact>
                <Article />
              </Route>
              <Route path="/article/:id" exact>
                <ArticleDetail />
              </Route>
              <Route path="/login" exact>
                <Login />
              </Route>
              <Route path="*">
                <NoMatch />
              </Route>
            </Switch>
          </Router>
        </Suspense>
      </Container>
    </ThemeProvider>
  )
}

export default App

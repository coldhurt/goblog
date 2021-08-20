import React, { lazy, Suspense } from 'react'
import './App.css'
import { BrowserRouter as Router, Switch, Route, useLocation } from 'react-router-dom'

const Home = lazy(() => import('./view/Home'))
const Login = lazy(() => import('./view/Login'))
const Article = lazy(() => import('./view/Article'))

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

function App() {
  return (
    <Suspense fallback={<div>loading</div>}>
      <Router>
        <Switch>
          <Route path="/" exact>
            <Home />
          </Route>
          <Route path="/article" exact>
            <Article />
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
  )
}

export default App

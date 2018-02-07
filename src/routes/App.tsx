import * as React from 'react'
import { connect } from 'react-redux'

import {
  BrowserRouter as Router,
  Route
} from 'react-router-dom'

import Home from './Home'

import '../styles/core/'

class App extends React.Component {
  render() {
    return (
      <Router>
        <div>
          <Route path="/" component={Home} />
        </div>
      </Router>
    )
  }
}

export default connect(
  state => ({
    state
  })
)(App as React.ComponentClass<any>)

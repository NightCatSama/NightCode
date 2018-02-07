import * as React from 'react'
import * as ReactDOM from 'react-dom'
import App from './routes/App'
import registerServiceWorker from './registerServiceWorker'

import { Middleware } from 'redux'
import { Provider } from 'react-redux'
import { ConnectedRouter } from 'react-router-redux'
import createHistory from 'history/createBrowserHistory'

import { routerMiddleware } from 'react-router-redux'

import configureStore from './stores'

const history = createHistory()
const historyMiddleWare: Middleware = routerMiddleware(history)

const store = configureStore(historyMiddleWare)

ReactDOM.render(
  <Provider store={store}>
    <ConnectedRouter history={history}>
      <App />
    </ConnectedRouter>
  </Provider>,
  document.getElementById('root') as HTMLElement
)
registerServiceWorker()

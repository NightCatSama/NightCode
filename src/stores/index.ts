import { createStore, combineReducers, applyMiddleware, Middleware, compose } from 'redux'
import { routerReducer } from 'react-router-redux'
import thunkMiddleware from 'redux-thunk'
import reducers from '../reducers'

export default (historyMiddleware: Middleware) => {
  const middlewares: Middleware[] = [
    thunkMiddleware, 
    historyMiddleware
  ]
  return createStore(
    combineReducers({
      ...reducers,
      router: routerReducer
    }),
    compose(
      applyMiddleware(...middlewares),
      (window as any).__REDUX_DEVTOOLS_EXTENSION__ && (window as any).__REDUX_DEVTOOLS_EXTENSION__()
    )
  )  
}

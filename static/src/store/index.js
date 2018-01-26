import {createStore, applyMiddleware, combineReducers} from 'redux'
import thunk from 'redux-thunk'
import * as auth from './auth'

export default createStore(combineReducers({
  auth: auth.reducer
}), applyMiddleware(thunk))
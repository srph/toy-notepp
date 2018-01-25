import {createStore, combineReducers} from 'redux';
import * as auth from './auth';

export default createStore(combineReducers({
  auth: auth.reducer
}))
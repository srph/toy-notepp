import {createStore, combineReducers} from 'redux';
import auth from './auth';

export default createStore(combineReducers({
  auth
}))
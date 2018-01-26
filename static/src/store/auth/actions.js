import axios from 'axios'
import history from '@/history'

export function init() {
  return (dispatch, getState) => {
    return axios.get('/me')
      .then(data => {
        dispatch({
          type: 'auth:init!',
          payload: data
        })

        history.push('/')
      })
      .catch(err => {
        dispatch({
          type: 'auth:init',
          payload: null
        })
      })
  }
}

export function login(credentials) {
  return (dispatch, getState) => {
    if (getState().auth.credentials) {
      return
    }

    dispatch({
      type: 'auth:login!',
      payload: credentials
    })

    return axios.post('/login', credentials)
      .then(data => {
        dispatch({
          type: 'auth:login.success',
          payload: data
        })
      })
      .catch(err => {
        dispatch({
          type: 'auth:logout.success',
          payload: { error: err.response.message }
        })
      })
  }
}

export function logout() {
  return (dispatch, getState) => {
    dispatch({ type: 'auth:logout' })
    history.push('/login')
  }
}
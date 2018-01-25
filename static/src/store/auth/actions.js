import axios from 'axios'

export function login(credentials) {
  return function(getState, dispatch) {
    dispatch({
      type: 'auth:login!',
      payload: credentials
    })

    axios.post('/api/login', credentails)
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
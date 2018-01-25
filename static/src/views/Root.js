import React, {Component} from 'react';
import Helmet from 'react-helmet';

import Router from 'react-router-dom';
import routes from '@/routes';
import history from '@/history';

import {Provider} from 'react-redux';
import store from '@/store';

class Root extends Component {
  render() {
    return (
      <Provider store={store}>
        <div>
          <Helmet titleTemplate="%s | Note++" />
          <Router history={history}>
            {routes}
          </Router>
        </div>
      </Provider>
    );
  }
}

export default Root;

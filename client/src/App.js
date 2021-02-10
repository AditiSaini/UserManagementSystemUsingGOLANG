import React, { useState } from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import './App.css';
import Login from './pages/login';
import ProfilePage from './pages/profile';
import InfoPage from './pages/info';
import PicturePage from './pages/picture';
import { createBrowserHistory } from 'history';


function App() {
  const history = createBrowserHistory();

  return (
    <Router >
      <div className="wrapper" >
        <Switch>
          <Route exact path="/">
            <Login history={history} />
          </Route>
          <Route path="/profile">
            <ProfilePage history={history} />
          </Route>
          <Route path="/info">
            <InfoPage history={history} />
          </Route>
          <Route path="/picture">
            <PicturePage history={history} />
          </Route>
        </Switch>
      </div>
    </Router >
  );
}

export default App;

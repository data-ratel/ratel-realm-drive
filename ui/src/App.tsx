import React from 'react';
import './App.css';
import './pages/page.css';
import 'bootstrap/dist/css/bootstrap.min.css';

import Home from './pages/Home';
import Login from './pages/Login';
import NotFound from './pages/NotFound';

import {
  HashRouter as Router,
  Route,
  NavLink
} from 'react-router-dom';

import { CSSTransition } from 'react-transition-group'
import { Navbar, Nav } from 'react-bootstrap'

function App() {
  const routes = [
    { path: '/', name: 'Home', Component: Home },
    { path: '/login', name: 'Login', Component: Login },
    { path: '/404', name: 'NotFound', Component: NotFound}
  ];

  return (
    <div className="App">
      <Router>
        <>
          <Navbar bg="light">
            <Nav className="mx-auto">
              {routes.filter(route => route.path !== '/404').map(route => (
                <Nav.Link
                  key={route.path}
                  as={NavLink}
                  to={route.path}
                  activeClassName="active"
                  exact
                >
                  {route.name}
                </Nav.Link>
              ))}
            </Nav>
          </Navbar>
          
          <div>
            {routes.map(({ path, Component }) => (
              <Route key={path} exact path={path}>
                {({ match }) => (
                  <CSSTransition
                    in={match != null}
                    timeout={300}
                    classNames='page'
                    unmountOnExit
                  >
                    <div className='page'>
                      <Component />
                    </div>
                  </CSSTransition>
                )}
              </Route>
            ))}
          </div>
        </>
      </Router>
    </div>
  );
}

export default App;
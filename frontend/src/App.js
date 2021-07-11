import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import { Home } from './pages/Home';
import { Create } from './pages/Create';
import { Register } from './pages/Register';
import { Events } from './pages/Events';
import { NoMatch } from './pages/NoMatch';
import { Layout } from './components/Layout';
import { NavBar } from './components/Nav';

class App extends Component {
  render() {
    return (
      <React.Fragment>
        <NavBar/>
        <Layout>
          <Router>
            <Switch>
              <Route exact path="/" component={Home} />
              <Route exact path="/create" component={Create} />
              <Route exact path="/register" component={Register} />
              <Route exact path="/events" component={Events} />
              <Route component={NoMatch} />
            </Switch>
          </Router>
        </Layout>
      </React.Fragment>
    )
  }
}

export default App;

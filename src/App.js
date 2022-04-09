import React, {Fragment} from 'react';
import { HashRouter as Router, Switch, Route, Link } from 'react-router-dom';
// import Income from './components/Income';
// import Expenses from './components/Expenses';
import Home from './components/Home';
import EditIncome from './components/EditIncome';
import EditExpenses from './components/EditExpenses';
// import './App.css';


export default function App() {
  return (
    <Router>
      <div className='container'>
        <div className='row'>
          <h1 className='mt-3'>
            Go Watch Movie!
          </h1>
          <hr className='mb-3'>
            
          </hr>
        </div>

        <div className='row'>
          <div className='col-md-2'>
            <nav>
              <ul className='list-group'>
                <li className='list-group-item'>
                  <Link to="/">Home</Link>
                </li>
                <li className='list-group-item'>
                  <Link to="/income/income/0">Income</Link>
                </li>
                {/* <li className='list-group-item'>
                  <Link to="/income/add">Income 1</Link>
                </li> */}
                <li className='list-group-item'>
                  <Link to="/expenses/add">Expenses</Link>
                </li>
              </ul>
            </nav>
          </div>
          <div className='col-md-10'>
            <Switch>

              <Route path="/income/income/:id" component={EditIncome} />
              {/* <Route path="/income/">
                <Income />
              </Route> */}

              <Route path="/expenses/add" component={EditExpenses} />
              {/* <Route path="/expenses">
                <Expenses />
              </Route> */}
              
              <Route path="/">
                <Home />
              </Route>
            </Switch>
          </div>
        </div>
      </div>
    </Router>
  );
}
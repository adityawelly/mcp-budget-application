import React, {Fragment} from 'react';
import { HashRouter as Router, Switch, Route, Link } from 'react-router-dom';
// import Income from './components/Income';
// import Expenses from './components/Expenses';
import Home from './components/Home';
import EditIncome from './components/EditIncome';
import EditExpenses from './components/EditExpenses';
import './App.css';


export default function App() {
  return (
    <Router>
      <nav>
        <label class="logo"> BudgetX </label>
        <ul>
          <li><Link to="/">Home</Link></li>
          <li><Link to="/income/income/0">Income</Link></li>
          <li><Link to="/expenses/add">Expenses</Link></li>
          <li><Link to="/login">Login</Link></li>
        </ul>
      </nav>
      
      <div className='box'>
        <Switch>
          <Route path="/income/income/:id" component={EditIncome} />
          <Route path="/expenses/add" component={EditExpenses} />
          <Route path="/"><Home /></Route>
        </Switch>
      </div>
    </Router>
  );
}
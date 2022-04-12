import React, { Component, Fragment } from "react";

export default class Home extends Component {

  state = {
    income: []
  };

  componentDidMount(){
    fetch("http://localhost:4000/v1/home")
    .then((response) => response.json())
    .then((json) => {
      this.setState({
        income: json.income,
        isLoaded: true,
      })
    })
  }

  componentWillUnmount() {
    this.setState({
      income: [
        {id: 1, jumlah_in: "20000"},
        {id: 2, jumlah_in: "12000"},
        {id: 3, jumlah_in: "31000"},
      ]
    })
  }

  render() {
    // const { income, isLoaded } = this.state;
    // if(!isLoaded) {
      // return <p>Loading...</p>
    // } else {
      return (
        <Fragment>
          <h1>asd</h1>

          <ul>
            {this.state.income.map( (m) => (
              <li key={m.id}>
                {m.jumlah_in}
              </li>
            ))}
          </ul>
        </Fragment>
        // <>
        // <h2 className="home-balance">Your Balance &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; : &nbsp; 287.000</h2>
        // <br/>
        // <h2 className="home-income">Income &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; : &nbsp; 500.000</h2>
        // <br/>
        // <h2 className="home-expenses">Expenses &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; : &nbsp; 213.000</h2>
        // </>
      );
    // }
  }
}
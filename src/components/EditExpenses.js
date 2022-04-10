import React, { Component, Fragment } from 'react'
import Input from './form-components/input';
import '../App.css';
import Alert from "./ui-components/Alert";

export default class EditExpenses extends Component {

  constructor(props) {
    super(props);
    this.state = {
      expenses: {
        id: 0,
        tanggal_ex: "",
        jumlah_ex: "",
        catatan_ex: "",
      },
      isLoaded: false,
      error: null,
      alert: {
        type: "d-none",
        message: "",
      }
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSubmit = (evt) => {
    evt.preventDefault();
    // console.log("Form was submitted");

    const data = new FormData(evt.target);
    const payload = Object.fromEntries(data.entries());
    console.log(payload);

    const requestOptions = {
      method: 'POST',
      body : JSON.stringify(payload)
    }

    fetch('http://localhost:4000/v1/expenses/editExpenses', requestOptions)
    .then(response => response.json())
    .then(data => {
      if (data.error) {
        this.setState({
          alert: { type: "alert-danger", message: data.error.message },
        });
      } else {
        this.setState({
          alert: { type: "alert-success", message: "Changes saved!" },
        });
      }
    })
  };

  handleChange = (evt) => {
    let value = evt.target.value;
    let name = evt.target.name;
    this.setState((prevState) => ({
      expenses: {
        ...prevState.expenses,
        [name]: value,
      }
    }))
  }

  componentDidMount() {
  //   this.setState({
  //     expenses: {
  //       title: "Jajan",
  //       amount: "20.000",
  //     }
  //   })

  const id = this.props.match.params.id;
  if (id > 0) {
    fetch("http://localhost:4000/v1/expenses/" + id)
      .then((response) => {
        if (response.status !== "200") {
          let err = Error;
          err.Message = "Invalid response code: " + response.status;
          this.setState({error: err});
        }
        return response.json();
      })
      .then((json) => {
        const tanggal = new Date(json.expenses.tanggal_in);

        this.setState(
          {
            expenses: {
              id: id,
              tanggal_in: tanggal.toISOString().split("T")[0],
              jumlah_in: json.expenses.jumlah_in,
              catatan_in: json.expenses.catatan_in,
            },
            isLoaded: true,
          },
          (error) => {
            this.setState({
              isLoaded: true,
              error,
            })
          }
        )
      })
  } else {
    this.setState({ isLoaded: true});
  }
  }

  render() {
    let {expenses, isLoaded, error} = this.state;

    if(error) {
      return <div>Error: {error.message}</div>
    } else if (!isLoaded) {
      return <p>Loading...</p>
    } else {

    return(
      <Fragment>
        <h1 className='expenses'>Add / Edit Expenses</h1>
        <Alert alertType={this.state.alert.type} alertMessage={this.state.alert.message} />
        <form onSubmit={this.handleSubmit}>
          <input type="hidden" name='id' id='id' value={expenses.id} onChange={this.handleChange}/>
          
          <div className='textbox'>
            <label>Tanggal</label>
            <Input type={'date'} name={'tanggal_ex'} value={expenses.tanggal_ex} handleChange={this.handleChange} />
          </div>
          
          <div className='textbox'>
            <label>Jumlah</label>
            <Input type={'text'} name={'jumlah_ex'} value={expenses.jumlah_ex} handleChange={this.handleChange} />
          </div>

          <div className='textbox'>
            <label>Catatan</label>
            <Input type={'text'} name={'catatan_ex'} value={expenses.catatan_ex} handleChange={this.handleChange} />
          </div>
          
          <button className='btn'>Simpan</button>
        </form>
      </Fragment>
    );
        }
  }
}
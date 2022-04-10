import React, { Component, Fragment } from 'react'
import Input from './form-components/input';
import '../App.css';
import Alert from "./ui-components/Alert";

export default class EditIncome extends Component {

  constructor(props) {
    super(props);
    this.state = {
      income: {
        id: 0,
        tanggal_in: "",
        jumlah_in: "",
        catatan_in: "",
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

    fetch('http://localhost:4000/v1/income/editIncome', requestOptions)
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
      income: {
        ...prevState.income,
        [name]: value,
      }
    }))
  }

  componentDidMount() {
    const id = this.props.match.params.id;
    if (id > 0) {
      fetch("http://localhost:4000/v1/income/" + id)
        .then((response) => {
          if (response.status !== "200") {
            let err = Error;
            err.Message = "Invalid response code: " + response.status;
            this.setState({error: err});
          }
          return response.json();
        })
        .then((json) => {
          const tanggal = new Date(json.income.tanggal_in);

          this.setState(
            {
              income: {
                id: id,
                tanggal_in: tanggal.toISOString().split("T")[0],
                jumlah_in: json.income.jumlah_in,
                catatan_in: json.income.catatan_in,
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
    let {income, isLoaded, error} = this.state;

    if(error) {
      return <div>Error: {error.message}</div>
    } else if (!isLoaded) {
      return <p>Loading...</p>
    } else {

    return(
      <Fragment>
        <h1 className='income'>Add / Edit Income</h1>
        <Alert alertType={this.state.alert.type} alertMessage={this.state.alert.message} />
        <form onSubmit={this.handleSubmit}>
          <input type="hidden" name='id' id='id' value={income.id} onChange={this.handleChange}/>
          
          <div className='textbox'>
            <label>Tanggal</label>
            <Input type={'date'} name={'tanggal_in'} value={income.tanggal_in} handleChange={this.handleChange} />
          </div>
          
          <div className='textbox'>
            <label>Jumlah</label>
            <Input type={'text'} name={'jumlah_in'} value={income.jumlah_in} handleChange={this.handleChange} />
          </div>

          <div className='textbox'>
            <label>Catatan</label>
            <Input type={'text'} name={'catatan_in'} value={income.catatan_in} handleChange={this.handleChange} />
          </div>
          
          <button className='btn'>Simpan</button>
        </form>
      </Fragment>
    );
        }
  }
}
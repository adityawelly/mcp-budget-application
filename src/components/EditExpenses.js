import React, { Component, Fragment } from 'react'

export default class EditExpenses extends Component {
  state = {
    expenses: {},
    isLoaded: false,
    error: null,
  }

  // componentDidMount() {
  //   this.setState({
  //     income: {
  //       title: "Jajan",
  //       amount: "20.000",
  //     }
  //   })
  // }

  render() {
    let {expenses} = this.state;

    return(
      <Fragment>
        <h2>Add/Edit Expenses</h2>
        <hr />
        <form method='post'>

          <div className='mb-3'>
            <label for='tanggal_ex' className='form-label'>
              Tanggal
            </label>
            <input type='text' className='form-control' id='tanggal_ex' name='tanggal_ex' value={expenses.tanggal_ex}/>
          </div>

          <div className='mb-3'>
            <label for='jumlah_ex' className='form-label'>
              Jumlah
            </label>
            <input type='jumlah_ex' className='form-control' id='jumlah_ex' name='jumlah_ex' value={expenses.jumlah_ex}/>
          </div>

          <div className='mb-3'>
            <label for='catatan_ex' className='form-label'>
              Catatan
            </label>
            <input type='text' className='form-control' id='catatan_ex' name='catatan_ex' value={expenses.catatan_ex}/>
          </div>
          <hr />
          <button className='btn'>Simpan</button>
        </form>
      </Fragment>
    )
  }
}
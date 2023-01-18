import React, { Component } from 'react'
import { Breadcrumb, Button, Col, FormControl, FormGroup, FormSelect, Row, Form, Table } from 'react-bootstrap';

export default class Pencairan extends Component {
  constructor(props) {
    super(props)
    this.state = {
      currentDate: new Date().toISOString().split('T')[0],
    };
  }

  handleSubmit = (e) => {
    e.preventDefault();
    console.log("handlesubmit")
  }

  render() {
    console.log(this.state.currentDate)
    return (
      <div>
        <h2 className="ps-4 pb-2">Checklist Pencairan</h2>
        <Breadcrumb className='ms-4 breadcrumb'>
          <Breadcrumb.Item href="/">Halaman Utama</Breadcrumb.Item>
          <Breadcrumb.Item active>Transaksi - Checklist Pencairan</Breadcrumb.Item>
        </Breadcrumb>

        <div className="container-fluid">
          <div className="status-container">
            <Form onSubmit={(e) => this.handleSubmit(e)} className="filterContainer">
              <Row className="d-flex align-items-center justify-content-center">
                <Col className="d-flex align-items-center gap-2 justify-content mb-3">
                  <label>Branch:</label>
                  <FormGroup>
                    <FormSelect name='branch'>
                      <option className='d-none'>Pilih...</option>
                      Cabang
                    </FormSelect>
                  </FormGroup>

                  <label>Company:</label>
                  <FormGroup>
                    <FormSelect name='company'>
                      <option className='d-none'>Pilih...</option>
                      Nama Perusahaan
                    </FormSelect>
                  </FormGroup>

                  <label>Tanggal Mulai: </label>
                  <FormGroup>
                    <FormControl type='date' name='startDate' defaultValue={this.state.currentDate}></FormControl>
                  </FormGroup>

                  <label>Tanggal Akhir:</label>
                  <FormGroup>
                    <FormControl type='date' name='endDate' defaultValue={this.state.currentDate}></FormControl>
                  </FormGroup>
                </Col>
              </Row>
              <Button type='submit' className="filterBtn">Submit</Button>
            </Form>

            <Table striped bordered hover className="table">
              <thead>
                <tr>
                  <th>#</th>
                  <th>PPK</th>
                  <th>Name</th>
                  <th>OTR</th>
                  <th>Loan Amount</th>
                  <th>Drawdown Date</th>
                  <th>Periode</th>
                  <th>Effective Rate</th>
                  <th>Angsuran</th>
                  <th>Branch</th>
                  <th>Check</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>1</td>
                  <td>Mark</td>
                  <td>Otto</td>
                  <td>@mdo</td>
                  <td>@mdo</td>
                  <td>@mdo</td>
                  <td>@mdo</td>
                  <td>@mdo</td>
                  <td>@mdo</td>
                  <td>@mdo</td>
                  <td>@mdo</td>
                </tr>
              </tbody>
            </Table>
          </div>
        </div>

      </div>
    )
  }
}

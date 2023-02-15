import React, { Component } from 'react'
import { Breadcrumb, Button, Col, FormControl, FormGroup, FormSelect, Row, Form, Table, Modal } from 'react-bootstrap';
import axios from "axios";

export default class Laporan extends Component {
  constructor(props) {
    super(props)
    this.state = {
      currentDate: new Date().toISOString().split('T')[0],
      drawdownReport: [],
      branchList: [],
      companyList: [],
      show: false,
      modalMessage: "",
    };
  }

  handleClose = () => this.setState({ show: false, showGreen: false })
  handleShow = () => this.setState({ show: true })
  setAlertMessage = (modalMessage) => this.setState({ modalMessage })

  componentDidMount() {
    this.getBranchList()
    this.getCompanyList()
    this.getDrawdownReport()
  }

  getDrawdownReport = () => {
    axios.get(`http://localhost:8080/getDrawdownReport`)
      .then(res => {
        const drawdownReport = res.data.customer_data_tab;
        this.setState({ drawdownReport });
      })
  }

  getPengajuanbyFilter = (branch, company, startDate, endDate) => {
    axios.get(`http://localhost:8080/getDrawdownReportFiltered?branch=` + branch + `&company=` + company + `&start_date=` + startDate + `&end_date=` + endDate)
      .then(res => {
        const drawdownReport = res.data.customer_data_tab;
        this.setState({ drawdownReport });
      })
  }

  getBranchList = () => {
    axios.get(`http://localhost:8080/getBranchList`)
      .then(res => {
        const branchList = res.data.branch_tab;
        this.setState({ branchList });
      })
  }

  getCompanyList = () => {
    axios.get(`http://localhost:8080/getCompanyList`)
      .then(res => {
        const companyList = res.data.mst_company_tab;
        this.setState({ companyList });
      })
  }

  handleSubmit = (e) => {
    e.preventDefault()
    const formData = new FormData(e.currentTarget)

    if (formData.get('branch') === "Pilih...") {
      this.setAlertMessage("Pilih cabang dahulu !")
      this.handleShow()
    }
    else if (formData.get('company') === "Pilih...") {
      this.setAlertMessage("Pilih company dahulu !")
      this.handleShow()
    }
    else {
      this.getPengajuanbyFilter(formData.get('branch'), formData.get('company'), formData.get('startDate'), formData.get('endDate'))
    }
  }

  render() {
    if (this.state.drawdownReport != null) {
      var dataPengajuan = this.state.drawdownReport.map(
        (dataPengajuan, index) => (
          <tr>
            <td>{index + 1}</td>
            <td>{dataPengajuan.ppk}</td>
            <td>{dataPengajuan.name}</td>
            <td>{dataPengajuan.channeling_company}</td>
            <td>{dataPengajuan.drawdown_date}</td>
            <td>{dataPengajuan.LoanDataTab.loan_amount}</td>
            <td>{dataPengajuan.LoanDataTab.loan_period}</td>
            <td>{dataPengajuan.LoanDataTab.interest_effective}%</td>
          </tr>
        )
      )
    }

    if (this.state.branchList != null) {
      var branch = this.state.branchList.map(
        branch => (
          <option value={branch.code}>{branch.code} - {branch.description}</option>
        )
      )
    }

    if (this.state.companyList != null) {
      var company = this.state.companyList.map(
        company => (
          <option value={company.company_name}>{company.company_name}</option>
        )
      )
    }

    return (
      <div>
        <h2 className="ps-4 pb-2">Laporan</h2>
        <Breadcrumb className='ms-4 breadcrumb'>
          <Breadcrumb.Item href="/">Halaman Utama</Breadcrumb.Item>
          <Breadcrumb.Item active>Laporan</Breadcrumb.Item>
        </Breadcrumb>

        <div className="container-fluid">
          <div className="status-container">
            <Form onSubmit={(e) => this.handleSubmit(e)} className="formContainer">
              <Row className="d-flex align-items-center justify-content-center">
                <Col className="d-flex align-items-center gap-2 justify-content mb-3">
                  <label>Branch:</label>
                  <FormGroup>
                    <FormSelect name='branch'>
                      <option>Pilih...</option>
                      {branch}
                    </FormSelect>
                  </FormGroup>

                  <label>Company:</label>
                  <FormGroup>
                    <FormSelect name='company'>
                      <option>Pilih...</option>
                      {company}
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
              <Button type='submit' className="defaultBtn">Submit</Button>
            </Form>

            <Table striped bordered hover className="table">
              <thead>
                <tr>
                  <th>#</th>
                  <th>PPK</th>
                  <th>Name</th>
                  <th>Company</th>
                  <th>Drawdown Date</th>
                  <th>Loan Amount</th>
                  <th>Periode</th>
                  <th>Effective Rate</th>
                </tr>
              </thead>
              <tbody>
                {dataPengajuan}
              </tbody>
            </Table>
          </div>
        </div>

        <Modal centered show={this.state.show} onHide={this.handleClose}>
          <Modal.Header closeButton className="backgroundRed">
            <Modal.Title className="text-white"></Modal.Title>
          </Modal.Header>
          <Modal.Body>{this.state.modalMessage}</Modal.Body>
        </Modal>
      </div>

    )
  }
}



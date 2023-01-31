import React, { Component } from 'react'
import { Breadcrumb, Button, Col, FormControl, FormGroup, FormSelect, Row, Form, Table, Modal } from 'react-bootstrap';
import axios from "axios";

export default class Pencairan extends Component {
  constructor(props) {
    super(props)
    this.state = {
      currentDate: new Date().toISOString().split('T')[0],
      checklistPengajuan: [],
      branchList: [],
      companyList: [],
      dataChecked: [],
      show: false,
      modalMessage: "Ini Validasi",
    };
  }

  handleClose = () => this.setState({ show: false })
  handleShow = () => this.setState({ show: true })
  setAlertMessage = (modalMessage) => this.setState({ modalMessage })

  componentDidMount() {
    this.getBranchList()
    this.getCompanyList()
    this.getChecklistPengajuan()
  }

  getChecklistPengajuan = () => {
    axios.get(`http://localhost:8080/getChecklistPengajuan?approval_status=9`)
      .then(res => {
        const checklistPengajuan = res.data.customer_data_tab;
        this.setState({ checklistPengajuan });
        console.log(this.state.checklistPengajuan)
      })
  }

  getPengajuanbyFilter = (branch, company, startDate, endDate) => {
    axios.get(`http://localhost:8080/getChecklistPengajuanFiltered?approval_status=9&branch=` + branch + `&company=` + company + `&start_date=` + startDate + `&end_date=` + endDate)
      .then(res => {
        const checklistPengajuan = res.data.customer_data_tab;
        this.setState({ checklistPengajuan });
        console.log(this.state.checklistPengajuan)
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

  dataChecked = (ppk, event) => {
    if (event.target.checked) {
      var dataChecked = [...this.state.dataChecked, { ppk: ppk }]
      this.setState({ dataChecked })
    }
    else {
      var data = this.state.dataChecked
      data = data.filter((j) => j.ppk !== ppk)
      this.setState({ dataChecked: data })
    }

  }

  handleSubmit = async (e) => {
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
    console.log(this.state.currentDate)
    if (this.state.checklistPengajuan != null) {
      var dataPengajuan = this.state.checklistPengajuan.map(
        (dataPengajuan, index) => (
          <tr>
            <td>{index + 1}</td>
            <td>{dataPengajuan.ppk}</td>
            <td>{dataPengajuan.name}</td>
            <td>{dataPengajuan.LoanDataTab.otr}</td>
            <td>{dataPengajuan.LoanDataTab.loan_amount}</td>
            <td>{dataPengajuan.drawdown_date}</td>
            <td>{dataPengajuan.LoanDataTab.loan_period}</td>
            <td>{dataPengajuan.LoanDataTab.interest_effective}%</td>
            <td>{dataPengajuan.LoanDataTab.monthly_payment}</td>
            <td>{dataPengajuan.channeling_company}</td>
            <td>{dataPengajuan.LoanDataTab.branch}</td>
            <td><input type={"checkbox"} onChange={(e) => this.dataChecked(dataPengajuan.ppk, e)}></input></td>
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
                  <th>Company</th>
                  <th>Branch</th>
                  <th>Check</th>
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

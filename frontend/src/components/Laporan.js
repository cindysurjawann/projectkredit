import React, { Component } from 'react'
import { Breadcrumb } from 'react-bootstrap'

export default class Laporan extends Component {
  render() {
    return (
      <div>
        <h2 className="ps-4 pb-2">Laporan</h2>
        <Breadcrumb className='ms-4 breadcrumb'>
          <Breadcrumb.Item href="/">Halaman Utama</Breadcrumb.Item>
          <Breadcrumb.Item active>Laporan</Breadcrumb.Item>
        </Breadcrumb>
      </div>
    )
  }
}

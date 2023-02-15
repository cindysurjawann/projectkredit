import React, { Component } from 'react'
import { Breadcrumb, Button, Form, FormControl, FormGroup, Modal } from 'react-bootstrap'
import axios from "axios";

export default class ChangePassword extends Component {
    constructor(props) {
        super(props)
        this.state = {
            show: false,
            showGreen: false,
            modalMessage: "",
            passwordTrue: false,
            passwordMatch: false,
            backToHome: false,
        };
    }

    handleClose = () => {
        this.setState({ show: false, showGreen: false })
        if (this.state.backToHome) {
            window.location.href = "/"
            this.setState({ backToHome: false })
        }
    }
    handleShow = () => this.setState({ show: true })
    handleShowGreen = () => this.setState({ showGreen: true })
    setAlertMessage = (modalMessage) => this.setState({ modalMessage })

    checkMatchNewPassword = (e) => {
        e.preventDefault()
        const formData = new FormData(e.currentTarget)

        if (formData.get('passwordBaru') === formData.get('konfirmPasswordBaru')) {
            this.setState({ passwordMatch: true })
            this.handleSubmit(formData)
        } else {
            this.setState({ passwordMatch: false })
            this.setAlertMessage("Password Baru dan Konfirmasi Password Baru tidak sesuai")
            this.handleShow()
        }
    }

    updatePassword = (formData) => {
        axios.patch(`http://localhost:8080/updatePassword`, {
            user_id: localStorage.getItem("user_id"),
            password: formData.get('passwordBaru')
        }).then((response) => {
            if (response.data.message == "success") {
                this.setAlertMessage("Ubah Password Berhasil")
                this.handleShowGreen()
                this.setState({ backToHome: true })
            } else {
                this.setAlertMessage("Ubah Password Gagal")
                this.handleShow()
            }
        })
    }

    handleSubmit = (formData) => {
        //check password to database & update when match
        axios.post('http://localhost:8080/matchPassword', {
            user_id: localStorage.getItem("user_id"),
            password: formData.get('passwordLama'),
        })
            .then((response) => {
                if (response.data.message !== "password not match") {
                    this.setState({ passwordTrue: true })
                    this.updatePassword(formData)
                } else {
                    this.setState({ passwordTrue: false })
                    this.setAlertMessage("Password lama salah")
                    this.handleShow()
                }
            })
    }

    render() {
        return (
            <div>
                <h2 className="ps-4 pb-2">Ubah Password</h2>
                <Breadcrumb className='ms-4 breadcrumb'>
                    <Breadcrumb.Item href="/">Halaman Utama</Breadcrumb.Item>
                    <Breadcrumb.Item active>Ubah Password</Breadcrumb.Item>
                </Breadcrumb>

                <div className="container-fluid">
                    <div className="status-container">
                        <Form onSubmit={(e) => this.checkMatchNewPassword(e)} className="formContainer">
                            <FormGroup className="gap-5 justify-content mb-3">
                                <label>Password Lama: </label>
                                <FormControl type="password" name="passwordLama" required></FormControl>
                            </FormGroup>

                            <FormGroup className="gap-5 justify-content mb-3">
                                <label>Password Baru: </label>
                                <FormControl type="password" name="passwordBaru" required></FormControl>
                            </FormGroup>

                            <FormGroup className="gap-5 justify-content mb-3">
                                <label>Konfirmasi Password Baru: </label>
                                <FormControl type="password" name="konfirmPasswordBaru" required></FormControl>
                            </FormGroup>

                            <Button type='submit' className="defaultBtn">Submit</Button>
                        </Form>
                    </div>
                </div>

                <Modal centered show={this.state.show} onHide={this.handleClose}>
                    <Modal.Header closeButton className="backgroundRed">
                        <Modal.Title className="text-white"></Modal.Title>
                    </Modal.Header>
                    <Modal.Body>{this.state.modalMessage}</Modal.Body>
                </Modal>

                <Modal centered show={this.state.showGreen} onHide={this.handleClose}>
                    <Modal.Header closeButton className="backgroundGreen">
                        <Modal.Title className="text-white"></Modal.Title>
                    </Modal.Header>
                    <Modal.Body>{this.state.modalMessage}</Modal.Body>
                </Modal>
            </div>
        )
    }
}

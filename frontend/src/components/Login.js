import React, { Component } from 'react'
import logo from "../assets/logobsim.png"
import axios from "axios";

export default class Login extends Component {
  constructor(props) {
    super(props);
  }

  handleLogin = (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget)


    const alertWrongInput = document.getElementById('alertWrongInput')
    const alertWrongConnection = document.getElementById('alertWrongConnection')

    axios.post('http://localhost:8080/login', {
      user_id: formData.get("user_id"),
      password: formData.get("password"),
    })
      .then((response) => {
        if (response.data.length !== 0) {
          localStorage.setItem("info", "true")
          localStorage.setItem("user_id", formData.get("user_id"))
          window.location.href = "/"
        } else {
          alertWrongInput.classList.remove("d-none");
        }
      }).catch(function (error) {
        console.log(error)
        if (error.response.data.message == "password is wrong") {
          alertWrongInput.classList.remove("d-none");
        }
        else {
          alertWrongConnection.classList.remove("d-none");
        }
      })
  }

  render() {
    return (
      <div>
        <div className="login-container">
          <div className="row justify-content-center">
            <div className="wrapper p-3">
              <div className="col">
                <div className="logoImg col-lg-12 d-flex justify-content-center">
                  <img src={logo}></img>
                </div>
                <div className="d-flex justify-content-center">
                  <div className="titleLogin">
                    <hr></hr>
                    <h3>Pencairan Kredit</h3>
                  </div>
                </div>
              </div>
              <div className="col-lg-12 d-flex justify-content-center">
                <form onSubmit={this.handleLogin} className="p-2">
                  <div id="alertWrongInput" className="alert alert-danger d-none" role="alert">
                    UserID atau Password salah
                  </div>
                  <div id="alertWrongConnection" className="alert alert-danger d-none" role="alert">
                    Terjadi error koneksi. Hubungi IT
                  </div>

                  <div className="input-group mb-4 mt-4">
                    <input type="text" className="input" name="user_id" />
                    <label className="placeholder">User ID</label>
                  </div>

                  <div className="input-group mb-4">
                    <input type="password" className="input" name="password" />
                    <label className="placeholder">Password</label>
                  </div>

                  <div className="d-flex justify-content-around">
                    <button type="submit" className="btn btn-pertama w-50">
                      Sign in
                    </button>
                  </div>
                </form>
              </div>
            </div>
          </div>
        </div>
      </div>
    )
  }
}
